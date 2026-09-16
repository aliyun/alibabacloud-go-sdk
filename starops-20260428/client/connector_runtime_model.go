// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorRuntime interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *ConnectorRuntime
	GetMode() *string
	SetPluginId(v string) *ConnectorRuntime
	GetPluginId() *string
	SetSatelliteId(v string) *ConnectorRuntime
	GetSatelliteId() *string
}

type ConnectorRuntime struct {
	// Runtime mode
	//
	// This parameter is required.
	//
	// example:
	//
	// STAROPS_MANAGED
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// Plugin ID
	//
	// example:
	//
	// gitlab
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// Satellite ID
	//
	// example:
	//
	// satellite-private-runtime
	SatelliteId *string `json:"satelliteId,omitempty" xml:"satelliteId,omitempty"`
}

func (s ConnectorRuntime) String() string {
	return dara.Prettify(s)
}

func (s ConnectorRuntime) GoString() string {
	return s.String()
}

func (s *ConnectorRuntime) GetMode() *string {
	return s.Mode
}

func (s *ConnectorRuntime) GetPluginId() *string {
	return s.PluginId
}

func (s *ConnectorRuntime) GetSatelliteId() *string {
	return s.SatelliteId
}

func (s *ConnectorRuntime) SetMode(v string) *ConnectorRuntime {
	s.Mode = &v
	return s
}

func (s *ConnectorRuntime) SetPluginId(v string) *ConnectorRuntime {
	s.PluginId = &v
	return s
}

func (s *ConnectorRuntime) SetSatelliteId(v string) *ConnectorRuntime {
	s.SatelliteId = &v
	return s
}

func (s *ConnectorRuntime) Validate() error {
	return dara.Validate(s)
}
