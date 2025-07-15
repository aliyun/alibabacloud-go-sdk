// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcGatewayEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *CreateVpcGatewayEndpointResponseBody
	GetCreationTime() *string
	SetEndpointId(v string) *CreateVpcGatewayEndpointResponseBody
	GetEndpointId() *string
	SetEndpointName(v string) *CreateVpcGatewayEndpointResponseBody
	GetEndpointName() *string
	SetRequestId(v string) *CreateVpcGatewayEndpointResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateVpcGatewayEndpointResponseBody
	GetResourceGroupId() *string
	SetServiceName(v string) *CreateVpcGatewayEndpointResponseBody
	GetServiceName() *string
}

type CreateVpcGatewayEndpointResponseBody struct {
	// The time when the gateway endpoint was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-27T01:58:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the gateway endpoint.
	//
	// example:
	//
	// vpce-bp1w1dmdqjpwul0v3****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the gateway endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 62CFC815-E08A-5CF4-92D1-54273EC9E406
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the gateway endpoint belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyun.cn-hangzhou.oss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateVpcGatewayEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcGatewayEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcGatewayEndpointResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *CreateVpcGatewayEndpointResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateVpcGatewayEndpointResponseBody) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateVpcGatewayEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcGatewayEndpointResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcGatewayEndpointResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateVpcGatewayEndpointResponseBody) SetCreationTime(v string) *CreateVpcGatewayEndpointResponseBody {
	s.CreationTime = &v
	return s
}

func (s *CreateVpcGatewayEndpointResponseBody) SetEndpointId(v string) *CreateVpcGatewayEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateVpcGatewayEndpointResponseBody) SetEndpointName(v string) *CreateVpcGatewayEndpointResponseBody {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcGatewayEndpointResponseBody) SetRequestId(v string) *CreateVpcGatewayEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcGatewayEndpointResponseBody) SetResourceGroupId(v string) *CreateVpcGatewayEndpointResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcGatewayEndpointResponseBody) SetServiceName(v string) *CreateVpcGatewayEndpointResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcGatewayEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
