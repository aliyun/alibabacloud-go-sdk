// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPolarClawPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *InstallPolarClawPluginRequest
	GetApplicationId() *string
	SetNpmPackage(v string) *InstallPolarClawPluginRequest
	GetNpmPackage() *string
	SetPluginId(v string) *InstallPolarClawPluginRequest
	GetPluginId() *string
	SetRestart(v bool) *InstallPolarClawPluginRequest
	GetRestart() *bool
}

type InstallPolarClawPluginRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @larksuite/openclaw-lark@2026.4.7
	NpmPackage *string `json:"NpmPackage,omitempty" xml:"NpmPackage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s InstallPolarClawPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallPolarClawPluginRequest) GoString() string {
	return s.String()
}

func (s *InstallPolarClawPluginRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *InstallPolarClawPluginRequest) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *InstallPolarClawPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *InstallPolarClawPluginRequest) GetRestart() *bool {
	return s.Restart
}

func (s *InstallPolarClawPluginRequest) SetApplicationId(v string) *InstallPolarClawPluginRequest {
	s.ApplicationId = &v
	return s
}

func (s *InstallPolarClawPluginRequest) SetNpmPackage(v string) *InstallPolarClawPluginRequest {
	s.NpmPackage = &v
	return s
}

func (s *InstallPolarClawPluginRequest) SetPluginId(v string) *InstallPolarClawPluginRequest {
	s.PluginId = &v
	return s
}

func (s *InstallPolarClawPluginRequest) SetRestart(v bool) *InstallPolarClawPluginRequest {
	s.Restart = &v
	return s
}

func (s *InstallPolarClawPluginRequest) Validate() error {
	return dara.Validate(s)
}
