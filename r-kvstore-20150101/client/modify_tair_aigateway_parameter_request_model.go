// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairAIGatewayParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyTairAIGatewayParameterRequest
	GetInstanceId() *string
	SetParameters(v string) *ModifyTairAIGatewayParameterRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyTairAIGatewayParameterRequest
	GetRegionId() *string
	SetSecurityToken(v string) *ModifyTairAIGatewayParameterRequest
	GetSecurityToken() *string
}

type ModifyTairAIGatewayParameterRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tg-2zeb1ce76fee****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The JSON-formatted text of parameters and their corresponding values. Format: {"Parameter1":"Value1","Parameter2":"Value2",...}. The specified values overwrite the existing values.
	//
	// > The parameters available for different versions are subject to the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"access-db-id":"1","cu-limit":"10"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyTairAIGatewayParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairAIGatewayParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyTairAIGatewayParameterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTairAIGatewayParameterRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyTairAIGatewayParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTairAIGatewayParameterRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyTairAIGatewayParameterRequest) SetInstanceId(v string) *ModifyTairAIGatewayParameterRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTairAIGatewayParameterRequest) SetParameters(v string) *ModifyTairAIGatewayParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyTairAIGatewayParameterRequest) SetRegionId(v string) *ModifyTairAIGatewayParameterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTairAIGatewayParameterRequest) SetSecurityToken(v string) *ModifyTairAIGatewayParameterRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyTairAIGatewayParameterRequest) Validate() error {
	return dara.Validate(s)
}
