// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnassociateRouteTableRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *UnassociateRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateRouteTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassociateRouteTableRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociateRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateRouteTableRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *UnassociateRouteTableRequest
	GetRouteTableId() *string
	SetVSwitchId(v string) *UnassociateRouteTableRequest
	GetVSwitchId() *string
}

type UnassociateRouteTableRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual private cloud (VPC) to which the custom route table belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The ID of the vSwitch from which you want to disassociate the route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-25naue4****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UnassociateRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateRouteTableRequest) GoString() string {
	return s.String()
}

func (s *UnassociateRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassociateRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateRouteTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateRouteTableRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *UnassociateRouteTableRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UnassociateRouteTableRequest) SetClientToken(v string) *UnassociateRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociateRouteTableRequest) SetOwnerAccount(v string) *UnassociateRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateRouteTableRequest) SetOwnerId(v int64) *UnassociateRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateRouteTableRequest) SetRegionId(v string) *UnassociateRouteTableRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateRouteTableRequest) SetResourceOwnerAccount(v string) *UnassociateRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateRouteTableRequest) SetResourceOwnerId(v int64) *UnassociateRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateRouteTableRequest) SetRouteTableId(v string) *UnassociateRouteTableRequest {
	s.RouteTableId = &v
	return s
}

func (s *UnassociateRouteTableRequest) SetVSwitchId(v string) *UnassociateRouteTableRequest {
	s.VSwitchId = &v
	return s
}

func (s *UnassociateRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
