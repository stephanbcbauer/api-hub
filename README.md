# API Hub

Welcome to the **API Hub** repository, a centralized location for hosting and viewing API documentation for the Tractus-X organization. This repository automates the collection of OpenAPI specifications from GitHub releases, generates Swagger UI documentation, and publishes it on GitHub Pages.

## Overview

The API Hub repository is designed to streamline the process of documenting APIs across multiple repositories within the Tractus-X organization. By automatically identifying OpenAPI specifications from GitHub releases, this repository ensures that the Swagger UI documentation is always up-to-date and accessible via GitHub Pages.

## Directory Structure

```bash
api-hub/
├── .github/
│ └── workflows/
│ └── publish_api.yml # GitHub Actions workflow file
├── docs/ # Directory for storing OpenAPI spec files & generated Swagger UI documentation
└── src/api-collector # Go code collecting API specs
```

## Contact

- [Mailing List](https://accounts.eclipse.org/mailing-list/tractusx-dev)

## NOTICE

This work is licensed under the [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0).

- SPDX-License-Identifier: Apache-2.0
- SPDX-FileCopyrightText: 2024 Contributors to the Eclipse Foundation
- Source URL: https://github.com/eclipse-tractusx/api-hub