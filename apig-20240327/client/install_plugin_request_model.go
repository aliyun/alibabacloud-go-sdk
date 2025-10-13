// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayIds(v []*string) *InstallPluginRequest
	GetGatewayIds() []*string
	SetPluginClassId(v string) *InstallPluginRequest
	GetPluginClassId() *string
}

type InstallPluginRequest struct {
	GatewayIds []*string `json:"gatewayIds,omitempty" xml:"gatewayIds,omitempty" type:"Repeated"`
	// example:
	//
	// pls-csqmjndlhtguk0loef21
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
}

func (s InstallPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallPluginRequest) GoString() string {
	return s.String()
}

func (s *InstallPluginRequest) GetGatewayIds() []*string {
	return s.GatewayIds
}

func (s *InstallPluginRequest) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *InstallPluginRequest) SetGatewayIds(v []*string) *InstallPluginRequest {
	s.GatewayIds = v
	return s
}

func (s *InstallPluginRequest) SetPluginClassId(v string) *InstallPluginRequest {
	s.PluginClassId = &v
	return s
}

func (s *InstallPluginRequest) Validate() error {
	return dara.Validate(s)
}
