// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVpcResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateVpcResponseBody
	GetResourceGroupId() *string
	SetRouteTableId(v string) *CreateVpcResponseBody
	GetRouteTableId() *string
	SetVRouterId(v string) *CreateVpcResponseBody
	GetVRouterId() *string
	SetVpcId(v string) *CreateVpcResponseBody
	GetVpcId() *string
}

type CreateVpcResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the route table that is automatically created by the system after the VPC is created.
	//
	// example:
	//
	// vtb-bp145q7glnuzdv****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The ID of the vRouter that is automatically created by the system after the VPC is created.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen8****
	VRouterId *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	// The ID of the created VPC.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcResponseBody) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateVpcResponseBody) GetVRouterId() *string {
	return s.VRouterId
}

func (s *CreateVpcResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcResponseBody) SetRequestId(v string) *CreateVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcResponseBody) SetResourceGroupId(v string) *CreateVpcResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcResponseBody) SetRouteTableId(v string) *CreateVpcResponseBody {
	s.RouteTableId = &v
	return s
}

func (s *CreateVpcResponseBody) SetVRouterId(v string) *CreateVpcResponseBody {
	s.VRouterId = &v
	return s
}

func (s *CreateVpcResponseBody) SetVpcId(v string) *CreateVpcResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
