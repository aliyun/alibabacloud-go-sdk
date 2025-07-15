// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateRouteTableRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *AssociateRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateRouteTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateRouteTableRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateRouteTableRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *AssociateRouteTableRequest
	GetRouteTableId() *string
	SetVSwitchId(v string) *AssociateRouteTableRequest
	GetVSwitchId() *string
}

type AssociateRouteTableRequest struct {
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
	// The region ID of the VPC to which the route table belongs.
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
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-25ncdvfaue4****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s AssociateRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTableRequest) GoString() string {
	return s.String()
}

func (s *AssociateRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateRouteTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateRouteTableRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *AssociateRouteTableRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AssociateRouteTableRequest) SetClientToken(v string) *AssociateRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateRouteTableRequest) SetOwnerAccount(v string) *AssociateRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateRouteTableRequest) SetOwnerId(v int64) *AssociateRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateRouteTableRequest) SetRegionId(v string) *AssociateRouteTableRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateRouteTableRequest) SetResourceOwnerAccount(v string) *AssociateRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateRouteTableRequest) SetResourceOwnerId(v int64) *AssociateRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateRouteTableRequest) SetRouteTableId(v string) *AssociateRouteTableRequest {
	s.RouteTableId = &v
	return s
}

func (s *AssociateRouteTableRequest) SetVSwitchId(v string) *AssociateRouteTableRequest {
	s.VSwitchId = &v
	return s
}

func (s *AssociateRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
