// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAppPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *InstallAppPluginRequest
	GetPluginId() *string
	SetPluginVersion(v string) *InstallAppPluginRequest
	GetPluginVersion() *string
}

type InstallAppPluginRequest struct {
	// The gateway plug-in ID.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The locked version. Leave empty to follow the latest version.
	//
	// example:
	//
	// 1.0.7
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
}

func (s InstallAppPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAppPluginRequest) GoString() string {
	return s.String()
}

func (s *InstallAppPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *InstallAppPluginRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *InstallAppPluginRequest) SetPluginId(v string) *InstallAppPluginRequest {
	s.PluginId = &v
	return s
}

func (s *InstallAppPluginRequest) SetPluginVersion(v string) *InstallAppPluginRequest {
	s.PluginVersion = &v
	return s
}

func (s *InstallAppPluginRequest) Validate() error {
	return dara.Validate(s)
}
