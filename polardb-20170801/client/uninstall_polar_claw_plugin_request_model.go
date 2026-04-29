// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPolarClawPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UninstallPolarClawPluginRequest
	GetApplicationId() *string
	SetPluginId(v string) *UninstallPolarClawPluginRequest
	GetPluginId() *string
	SetRestart(v bool) *UninstallPolarClawPluginRequest
	GetRestart() *bool
}

type UninstallPolarClawPluginRequest struct {
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
	// openclaw-lark
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
}

func (s UninstallPolarClawPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallPolarClawPluginRequest) GoString() string {
	return s.String()
}

func (s *UninstallPolarClawPluginRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UninstallPolarClawPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *UninstallPolarClawPluginRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UninstallPolarClawPluginRequest) SetApplicationId(v string) *UninstallPolarClawPluginRequest {
	s.ApplicationId = &v
	return s
}

func (s *UninstallPolarClawPluginRequest) SetPluginId(v string) *UninstallPolarClawPluginRequest {
	s.PluginId = &v
	return s
}

func (s *UninstallPolarClawPluginRequest) SetRestart(v bool) *UninstallPolarClawPluginRequest {
	s.Restart = &v
	return s
}

func (s *UninstallPolarClawPluginRequest) Validate() error {
	return dara.Validate(s)
}
