// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ListAccessPointsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAccessPointsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListAccessPointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAccessPointsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListAccessPointsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListAccessPointsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAccessPointsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ListAccessPointsRequest
	GetSmartAGId() *string
}

type ListAccessPointsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// Valid values: **1 to 50**.
	//
	// Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the queried access points reside. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-far8v6owtdxlua****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ListAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessPointsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAccessPointsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAccessPointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAccessPointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccessPointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccessPointsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAccessPointsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAccessPointsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ListAccessPointsRequest) SetOwnerAccount(v string) *ListAccessPointsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAccessPointsRequest) SetOwnerId(v int64) *ListAccessPointsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAccessPointsRequest) SetPageNumber(v int32) *ListAccessPointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccessPointsRequest) SetPageSize(v int32) *ListAccessPointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccessPointsRequest) SetRegionId(v string) *ListAccessPointsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccessPointsRequest) SetResourceOwnerAccount(v string) *ListAccessPointsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAccessPointsRequest) SetResourceOwnerId(v int64) *ListAccessPointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAccessPointsRequest) SetSmartAGId(v string) *ListAccessPointsRequest {
	s.SmartAGId = &v
	return s
}

func (s *ListAccessPointsRequest) Validate() error {
	return dara.Validate(s)
}
