// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDefaultVpcResponseBody
	GetRequestId() *string
	SetRouteTableId(v string) *CreateDefaultVpcResponseBody
	GetRouteTableId() *string
	SetVRouterId(v string) *CreateDefaultVpcResponseBody
	GetVRouterId() *string
	SetVpcId(v string) *CreateDefaultVpcResponseBody
	GetVpcId() *string
}

type CreateDefaultVpcResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the route table that the system automatically creates after the default VPC is created.
	//
	// example:
	//
	// vtb-bp1q1uirugzb1x32m****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The ID of the vRouter that the system automatically creates after the default VPC is created.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen8****
	VRouterId *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	// The ID of the default VPC that is created.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateDefaultVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultVpcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefaultVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDefaultVpcResponseBody) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateDefaultVpcResponseBody) GetVRouterId() *string {
	return s.VRouterId
}

func (s *CreateDefaultVpcResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDefaultVpcResponseBody) SetRequestId(v string) *CreateDefaultVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefaultVpcResponseBody) SetRouteTableId(v string) *CreateDefaultVpcResponseBody {
	s.RouteTableId = &v
	return s
}

func (s *CreateDefaultVpcResponseBody) SetVRouterId(v string) *CreateDefaultVpcResponseBody {
	s.VRouterId = &v
	return s
}

func (s *CreateDefaultVpcResponseBody) SetVpcId(v string) *CreateDefaultVpcResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateDefaultVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
