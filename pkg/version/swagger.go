// Vikunja is a to-do list application to facilitate your life.
// Copyright 2018-2023 Vikunja and contributors. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public Licensee as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public Licensee for more details.
//
// You should have received a copy of the GNU Affero General Public Licensee
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//go:build swagger
// +build swagger

package version

import "code.vikunja.io/api/pkg/swagger"

func init() {
	// Additional swagger information
	swagger.SwaggerInfo.Version = Version
}
