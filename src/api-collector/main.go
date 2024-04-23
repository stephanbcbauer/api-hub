/*******************************************************************************
 * Copyright (c) 2024 Contributors to the Eclipse Foundation
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information regarding copyright ownership.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache License, Version 2.0 which is available at
 * https://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 ******************************************************************************/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

func main() {
	owner, token := getArgs()
	ctx := context.Background()
	client := getAuthenticatedClient(ctx, token)
	repos, err := getOrgRepos(ctx, owner, client)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(repos)
}

func getArgs() (string, string) {

	owner := flag.String("owner", "", "Specify GitHub User or Organization")
	token := flag.String("token", "", "Specify GitHub Token")
	flag.Parse()

	if *owner == "" || *token == "" {
		fmt.Println("Missing required arguments, please specify -owner [owner] and -token [token]")
		os.Exit(1)
	}
	return *owner, *token
}

func getAuthenticatedClient(ctx context.Context, gitToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getOrgRepos(ctx context.Context, gitOwner string, client *github.Client) ([]*github.Repository, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{},
	}

	var allRepos []*github.Repository

	for {
		repos, response, err := client.Repositories.ListByOrg(ctx, gitOwner, opt)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if response.NextPage == 0 {
			break
		}
		opt.Page = response.NextPage
	}

	return allRepos, nil
}
