// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTairAIGatewayParametersRequest
	GetInstanceId() *string
	SetPluginType(v string) *DescribeTairAIGatewayParametersRequest
	GetPluginType() *string
	SetRegionId(v string) *DescribeTairAIGatewayParametersRequest
	GetRegionId() *string
	SetSecurityToken(v string) *DescribeTairAIGatewayParametersRequest
	GetSecurityToken() *string
}

type DescribeTairAIGatewayParametersRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tg-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The plug-in type. You can use this parameter to filter the parameters of a specific plug-in.
	//
	// example:
	//
	// ai_cache
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61012.html) operation to obtain the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeTairAIGatewayParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayParametersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairAIGatewayParametersRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *DescribeTairAIGatewayParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairAIGatewayParametersRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairAIGatewayParametersRequest) SetInstanceId(v string) *DescribeTairAIGatewayParametersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairAIGatewayParametersRequest) SetPluginType(v string) *DescribeTairAIGatewayParametersRequest {
	s.PluginType = &v
	return s
}

func (s *DescribeTairAIGatewayParametersRequest) SetRegionId(v string) *DescribeTairAIGatewayParametersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTairAIGatewayParametersRequest) SetSecurityToken(v string) *DescribeTairAIGatewayParametersRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairAIGatewayParametersRequest) Validate() error {
	return dara.Validate(s)
}
