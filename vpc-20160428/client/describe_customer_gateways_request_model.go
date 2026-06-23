// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerGatewayId(v string) *DescribeCustomerGatewaysRequest
	GetCustomerGatewayId() *string
	SetOwnerAccount(v string) *DescribeCustomerGatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCustomerGatewaysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCustomerGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCustomerGatewaysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCustomerGatewaysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCustomerGatewaysRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCustomerGatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCustomerGatewaysRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeCustomerGatewaysRequestTag) *DescribeCustomerGatewaysRequest
	GetTag() []*DescribeCustomerGatewaysRequestTag
}

type DescribeCustomerGatewaysRequest struct {
	// The instance ID of the customer gateway.
	//
	// > If you do not specify the instance ID of a customer gateway, the system queries information about all customer gateways in the current region by default.
	//
	// example:
	//
	// cgw-bp1pvpl9r9adju6l5****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the customer gateway.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the customer gateway belongs.
	//
	// You can call [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) to query resource group IDs.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags bound to the customer gateway.
	Tag []*DescribeCustomerGatewaysRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCustomerGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeCustomerGatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCustomerGatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCustomerGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCustomerGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomerGatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomerGatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCustomerGatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCustomerGatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCustomerGatewaysRequest) GetTag() []*DescribeCustomerGatewaysRequestTag {
	return s.Tag
}

func (s *DescribeCustomerGatewaysRequest) SetCustomerGatewayId(v string) *DescribeCustomerGatewaysRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetOwnerAccount(v string) *DescribeCustomerGatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetOwnerId(v int64) *DescribeCustomerGatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetPageNumber(v int32) *DescribeCustomerGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetPageSize(v int32) *DescribeCustomerGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetRegionId(v string) *DescribeCustomerGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetResourceGroupId(v string) *DescribeCustomerGatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetResourceOwnerAccount(v string) *DescribeCustomerGatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetResourceOwnerId(v int64) *DescribeCustomerGatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCustomerGatewaysRequest) SetTag(v []*DescribeCustomerGatewaysRequestTag) *DescribeCustomerGatewaysRequest {
	s.Tag = v
	return s
}

func (s *DescribeCustomerGatewaysRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomerGatewaysRequestTag struct {
	// The tag key. If you specify this parameter, the value cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length and can be an empty string. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCustomerGatewaysRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCustomerGatewaysRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCustomerGatewaysRequestTag) SetKey(v string) *DescribeCustomerGatewaysRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCustomerGatewaysRequestTag) SetValue(v string) *DescribeCustomerGatewaysRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCustomerGatewaysRequestTag) Validate() error {
	return dara.Validate(s)
}
