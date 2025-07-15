// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociateType(v string) *CreateRouteTableRequest
	GetAssociateType() *string
	SetClientToken(v string) *CreateRouteTableRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRouteTableRequest
	GetDescription() *string
	SetOwnerAccount(v string) *CreateRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateRouteTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateRouteTableRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateRouteTableRequest
	GetResourceOwnerId() *int64
	SetRouteTableName(v string) *CreateRouteTableRequest
	GetRouteTableName() *string
	SetTag(v []*CreateRouteTableRequestTag) *CreateRouteTableRequest
	GetTag() []*CreateRouteTableRequestTag
	SetVpcId(v string) *CreateRouteTableRequest
	GetVpcId() *string
}

type CreateRouteTableRequest struct {
	// The type of the route table. Valid values:
	//
	// 	- **VSwitch*	- (default): vSwitch route table
	//
	// 	- **Gateway**: gateway route table
	//
	// example:
	//
	// VSwitch
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- in each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the route table.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the route table.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// myRouteTable
	RouteTableName *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	// The tags of the resource.
	Tag []*CreateRouteTableRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the custom route table belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1qpo0kug3a20qqe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTableRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteTableRequest) GetAssociateType() *string {
	return s.AssociateType
}

func (s *CreateRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteTableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRouteTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateRouteTableRequest) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *CreateRouteTableRequest) GetTag() []*CreateRouteTableRequestTag {
	return s.Tag
}

func (s *CreateRouteTableRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateRouteTableRequest) SetAssociateType(v string) *CreateRouteTableRequest {
	s.AssociateType = &v
	return s
}

func (s *CreateRouteTableRequest) SetClientToken(v string) *CreateRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteTableRequest) SetDescription(v string) *CreateRouteTableRequest {
	s.Description = &v
	return s
}

func (s *CreateRouteTableRequest) SetOwnerAccount(v string) *CreateRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateRouteTableRequest) SetOwnerId(v int64) *CreateRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRouteTableRequest) SetRegionId(v string) *CreateRouteTableRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteTableRequest) SetResourceOwnerAccount(v string) *CreateRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRouteTableRequest) SetResourceOwnerId(v int64) *CreateRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRouteTableRequest) SetRouteTableName(v string) *CreateRouteTableRequest {
	s.RouteTableName = &v
	return s
}

func (s *CreateRouteTableRequest) SetTag(v []*CreateRouteTableRequestTag) *CreateRouteTableRequest {
	s.Tag = v
	return s
}

func (s *CreateRouteTableRequest) SetVpcId(v string) *CreateRouteTableRequest {
	s.VpcId = &v
	return s
}

func (s *CreateRouteTableRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRouteTableRequestTag struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRouteTableRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTableRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRouteTableRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRouteTableRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRouteTableRequestTag) SetKey(v string) *CreateRouteTableRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRouteTableRequestTag) SetValue(v string) *CreateRouteTableRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRouteTableRequestTag) Validate() error {
	return dara.Validate(s)
}
