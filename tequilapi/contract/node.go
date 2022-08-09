/*
 * Copyright (C) 2021 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package contract

import (
	"github.com/mysteriumnetwork/node/core/node"
)

// NodeStatusResponse a node status reflects monitoring agent POV on node availability
// swagger:model NodeStatusResponse
type NodeStatusResponse struct {
	Status node.MonitoringStatus `json:"status"`
}

// MonitoringAgentResponse reflects amount of connectivity statuses for each service_type.
// swagger:model MonitoringAgentResponse
type MonitoringAgentResponse struct {
	Statuses node.MonitoringAgentStatuses `json:"statuses"`
	Error    string                       `json:"error,omitempty"`
}

// MonitoringAgentSessionsResponse reflects a list of sessions metrics during a period of time.
// swagger:model MonitoringAgentSessionsResponse
type MonitoringAgentSessionsResponse struct {
	Statuses node.MonitoringAgentSessions `json:"sessions"`
	Error    string                       `json:"error,omitempty"`
}
