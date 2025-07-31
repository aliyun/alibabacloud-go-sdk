// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewalAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetDBInstanceId() *string
	SetDBInstanceType(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetDBInstanceType() *string
	SetOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeInstanceAutoRenewalAttributeRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInstanceAutoRenewalAttributeRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceAutoRenewalAttributeRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// dds-bp567b****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The category of the instance. Valid values:
	//
	// 	- **replicate**: the standalone or replica set instance
	//
	// 	- **sharding**: the sharded cluster instance
	//
	// Default value: **replicate**.
	//
	// example:
	//
	// replicate
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be a positive integer that does not exceed the maximum value of the Integer parameter. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**.
	//
	// >  Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetDBInstanceId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetDBInstanceType(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageNumber(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageSize(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetRegionId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) Validate() error {
	return dara.Validate(s)
}
