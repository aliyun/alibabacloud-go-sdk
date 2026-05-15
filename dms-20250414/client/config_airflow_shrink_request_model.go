// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAirflowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *ConfigAirflowShrinkRequest
	GetAirflowId() *string
	SetCustomAirflowCfgShrink(v string) *ConfigAirflowShrinkRequest
	GetCustomAirflowCfgShrink() *string
	SetWorkspaceId(v string) *ConfigAirflowShrinkRequest
	GetWorkspaceId() *string
}

type ConfigAirflowShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af-b3a7f110a6vmvn7xxxxxx
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// This parameter is required.
	CustomAirflowCfgShrink *string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ConfigAirflowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigAirflowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigAirflowShrinkRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *ConfigAirflowShrinkRequest) GetCustomAirflowCfgShrink() *string {
	return s.CustomAirflowCfgShrink
}

func (s *ConfigAirflowShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ConfigAirflowShrinkRequest) SetAirflowId(v string) *ConfigAirflowShrinkRequest {
	s.AirflowId = &v
	return s
}

func (s *ConfigAirflowShrinkRequest) SetCustomAirflowCfgShrink(v string) *ConfigAirflowShrinkRequest {
	s.CustomAirflowCfgShrink = &v
	return s
}

func (s *ConfigAirflowShrinkRequest) SetWorkspaceId(v string) *ConfigAirflowShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ConfigAirflowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
