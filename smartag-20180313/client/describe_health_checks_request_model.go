// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthChecksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHcInstanceId(v string) *DescribeHealthChecksRequest
	GetHcInstanceId() *string
	SetName(v string) *DescribeHealthChecksRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeHealthChecksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHealthChecksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeHealthChecksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHealthChecksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHealthChecksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHealthChecksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHealthChecksRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeHealthChecksRequest
	GetSmartAGId() *string
}

type DescribeHealthChecksRequest struct {
	// The ID of the health check.
	//
	// example:
	//
	// hc-ifflm5ygj3diwiu****
	HcInstanceId *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	// The name of the health check.
	//
	// The name must be 2 to 100 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// ddd3333
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-qi0p07ld5q86k3****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeHealthChecksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthChecksRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthChecksRequest) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *DescribeHealthChecksRequest) GetName() *string {
	return s.Name
}

func (s *DescribeHealthChecksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHealthChecksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHealthChecksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHealthChecksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHealthChecksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHealthChecksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHealthChecksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHealthChecksRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeHealthChecksRequest) SetHcInstanceId(v string) *DescribeHealthChecksRequest {
	s.HcInstanceId = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetName(v string) *DescribeHealthChecksRequest {
	s.Name = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetOwnerAccount(v string) *DescribeHealthChecksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetOwnerId(v int64) *DescribeHealthChecksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetPageNumber(v int32) *DescribeHealthChecksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetPageSize(v int32) *DescribeHealthChecksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetRegionId(v string) *DescribeHealthChecksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetResourceOwnerAccount(v string) *DescribeHealthChecksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetResourceOwnerId(v int64) *DescribeHealthChecksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHealthChecksRequest) SetSmartAGId(v string) *DescribeHealthChecksRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeHealthChecksRequest) Validate() error {
	return dara.Validate(s)
}
