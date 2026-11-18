// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairAIGatewayPluginSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyTairAIGatewayPluginSwitchRequest
	GetInstanceId() *string
	SetParameters(v string) *ModifyTairAIGatewayPluginSwitchRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyTairAIGatewayPluginSwitchRequest
	GetRegionId() *string
	SetSecurityToken(v string) *ModifyTairAIGatewayPluginSwitchRequest
	GetSecurityToken() *string
}

type ModifyTairAIGatewayPluginSwitchRequest struct {
	// The Tair AI gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tg-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The plug-in switch parameters for the Tair AI gateway instance. Format: {"PluginName1":"true/false","PluginName2":"true/false",...}. The specified values overwrite the existing values.
	//
	// example:
	//
	// {\\"lang_cache\\":\\"false\\"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyTairAIGatewayPluginSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairAIGatewayPluginSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) SetInstanceId(v string) *ModifyTairAIGatewayPluginSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) SetParameters(v string) *ModifyTairAIGatewayPluginSwitchRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) SetRegionId(v string) *ModifyTairAIGatewayPluginSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) SetSecurityToken(v string) *ModifyTairAIGatewayPluginSwitchRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchRequest) Validate() error {
	return dara.Validate(s)
}
