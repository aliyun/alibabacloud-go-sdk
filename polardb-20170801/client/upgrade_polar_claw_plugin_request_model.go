// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpgradePolarClawPluginRequest
	GetApplicationId() *string
	SetNpmPackage(v string) *UpgradePolarClawPluginRequest
	GetNpmPackage() *string
	SetPluginId(v string) *UpgradePolarClawPluginRequest
	GetPluginId() *string
	SetRestart(v bool) *UpgradePolarClawPluginRequest
	GetRestart() *bool
}

type UpgradePolarClawPluginRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-********************
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

func (s UpgradePolarClawPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawPluginRequest) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawPluginRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpgradePolarClawPluginRequest) GetNpmPackage() *string {
	return s.NpmPackage
}

func (s *UpgradePolarClawPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *UpgradePolarClawPluginRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpgradePolarClawPluginRequest) SetApplicationId(v string) *UpgradePolarClawPluginRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpgradePolarClawPluginRequest) SetNpmPackage(v string) *UpgradePolarClawPluginRequest {
	s.NpmPackage = &v
	return s
}

func (s *UpgradePolarClawPluginRequest) SetPluginId(v string) *UpgradePolarClawPluginRequest {
	s.PluginId = &v
	return s
}

func (s *UpgradePolarClawPluginRequest) SetRestart(v bool) *UpgradePolarClawPluginRequest {
	s.Restart = &v
	return s
}

func (s *UpgradePolarClawPluginRequest) Validate() error {
	return dara.Validate(s)
}
