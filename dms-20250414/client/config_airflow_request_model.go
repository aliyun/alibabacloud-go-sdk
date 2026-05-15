// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAirflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirflowId(v string) *ConfigAirflowRequest
	GetAirflowId() *string
	SetCustomAirflowCfg(v []*string) *ConfigAirflowRequest
	GetCustomAirflowCfg() []*string
	SetWorkspaceId(v string) *ConfigAirflowRequest
	GetWorkspaceId() *string
}

type ConfigAirflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af-b3a7f110a6vmvn7xxxxxx
	AirflowId *string `json:"AirflowId,omitempty" xml:"AirflowId,omitempty"`
	// This parameter is required.
	CustomAirflowCfg []*string `json:"CustomAirflowCfg,omitempty" xml:"CustomAirflowCfg,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ConfigAirflowRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigAirflowRequest) GoString() string {
	return s.String()
}

func (s *ConfigAirflowRequest) GetAirflowId() *string {
	return s.AirflowId
}

func (s *ConfigAirflowRequest) GetCustomAirflowCfg() []*string {
	return s.CustomAirflowCfg
}

func (s *ConfigAirflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ConfigAirflowRequest) SetAirflowId(v string) *ConfigAirflowRequest {
	s.AirflowId = &v
	return s
}

func (s *ConfigAirflowRequest) SetCustomAirflowCfg(v []*string) *ConfigAirflowRequest {
	s.CustomAirflowCfg = v
	return s
}

func (s *ConfigAirflowRequest) SetWorkspaceId(v string) *ConfigAirflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ConfigAirflowRequest) Validate() error {
	return dara.Validate(s)
}
