// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeGrantRulesToCenRequest
	GetClientToken() *string
	SetInstanceId(v string) *DescribeGrantRulesToCenRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeGrantRulesToCenRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *DescribeGrantRulesToCenRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGrantRulesToCenRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeGrantRulesToCenRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantRulesToCenRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGrantRulesToCenRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGrantRulesToCenRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeGrantRulesToCenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGrantRulesToCenRequest
	GetResourceOwnerId() *int64
}

type DescribeGrantRulesToCenRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the network instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnvc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VBR**
	//
	// 	- **CCN**
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the network instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the network instance belongs.
	//
	// example:
	//
	// rg-acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeGrantRulesToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToCenRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToCenRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeGrantRulesToCenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGrantRulesToCenRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGrantRulesToCenRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGrantRulesToCenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGrantRulesToCenRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantRulesToCenRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantRulesToCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantRulesToCenRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGrantRulesToCenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGrantRulesToCenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGrantRulesToCenRequest) SetClientToken(v string) *DescribeGrantRulesToCenRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetInstanceId(v string) *DescribeGrantRulesToCenRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetInstanceType(v string) *DescribeGrantRulesToCenRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerAccount(v string) *DescribeGrantRulesToCenRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetPageNumber(v int32) *DescribeGrantRulesToCenRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetPageSize(v int32) *DescribeGrantRulesToCenRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetRegionId(v string) *DescribeGrantRulesToCenRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetResourceGroupId(v string) *DescribeGrantRulesToCenRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetResourceOwnerAccount(v string) *DescribeGrantRulesToCenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) SetResourceOwnerId(v int64) *DescribeGrantRulesToCenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGrantRulesToCenRequest) Validate() error {
	return dara.Validate(s)
}
