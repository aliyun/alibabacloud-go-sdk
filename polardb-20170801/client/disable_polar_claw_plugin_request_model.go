// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisablePolarClawPluginRequest
	GetApplicationId() *string
	SetPluginId(v string) *DisablePolarClawPluginRequest
	GetPluginId() *string
	SetRestart(v bool) *DisablePolarClawPluginRequest
	GetRestart() *bool
}

type DisablePolarClawPluginRequest struct {
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

func (s DisablePolarClawPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawPluginRequest) GoString() string {
	return s.String()
}

func (s *DisablePolarClawPluginRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisablePolarClawPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DisablePolarClawPluginRequest) GetRestart() *bool {
	return s.Restart
}

func (s *DisablePolarClawPluginRequest) SetApplicationId(v string) *DisablePolarClawPluginRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisablePolarClawPluginRequest) SetPluginId(v string) *DisablePolarClawPluginRequest {
	s.PluginId = &v
	return s
}

func (s *DisablePolarClawPluginRequest) SetRestart(v bool) *DisablePolarClawPluginRequest {
	s.Restart = &v
	return s
}

func (s *DisablePolarClawPluginRequest) Validate() error {
	return dara.Validate(s)
}
