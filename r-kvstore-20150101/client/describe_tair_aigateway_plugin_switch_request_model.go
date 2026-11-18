// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayPluginSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairAIGatewayPluginSwitchRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeTairAIGatewayPluginSwitchRequest
	GetRegionId() *string
	SetSecurityToken(v string) *DescribeTairAIGatewayPluginSwitchRequest
	GetSecurityToken() *string
}

type DescribeTairAIGatewayPluginSwitchRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tg-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeTairAIGatewayPluginSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayPluginSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayPluginSwitchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairAIGatewayPluginSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairAIGatewayPluginSwitchRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairAIGatewayPluginSwitchRequest) SetInstanceId(v string) *DescribeTairAIGatewayPluginSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchRequest) SetRegionId(v string) *DescribeTairAIGatewayPluginSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchRequest) SetSecurityToken(v string) *DescribeTairAIGatewayPluginSwitchRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairAIGatewayPluginSwitchRequest) Validate() error {
	return dara.Validate(s)
}
