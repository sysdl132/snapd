// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package client

import (
	"golang.org/x/xerrors"

	"github.com/snapcore/snapd/snap"
)

// SystemModelData contains information about the model
type SystemModelData struct {
	// Model as the model assertion
	Model string `json:"model,omitempty"`
	// BrandID corresponds to brand-id in the model assertion
	BrandID string `json:"brand-id,omitempty"`
	// DisplayName is human friendly name, corresponds to display-name in
	// the model assertion
	DisplayName string `json:"display-name,omitempty"`
}

type System struct {
	// Current is true when the system running now was installed from that
	// recovery seed
	Current bool `json:"current,omitempty"`
	// Label of the recovery system
	Label string `json:"label,omitempty"`
	// Model information
	Model SystemModelData `json:"model,omitempty"`
	// Brand information
	Brand snap.StoreAccount `json:"brand,omitempty"`
	// Actions available for this system
	Actions []SystemAction `json:"actions,omitempty"`
}

type SystemAction struct {
	// Title is a user presentable action description
	Title string `json:"title,omitempty"`
	// Mode given action can be executed in
	Mode string `json:"mode,omitempty"`
}

// ListSystems list all systems available for seeding or recovery.
func (client *Client) ListSystems() ([]System, error) {
	type systemsResponse struct {
		Systems []System `json:"systems,omitempty"`
	}

	var rsp systemsResponse

	if _, err := client.doSync("GET", "/v2/systems", nil, nil, nil, &rsp); err != nil {
		return nil, xerrors.Errorf("cannot list recovery systems: %v", err)
	}
	return rsp.Systems, nil
}
