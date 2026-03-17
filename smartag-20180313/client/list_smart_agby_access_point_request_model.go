// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartAGByAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v int32) *ListSmartAGByAccessPointRequest
	GetAccessPointId() *int32
	SetOwnerAccount(v string) *ListSmartAGByAccessPointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListSmartAGByAccessPointRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListSmartAGByAccessPointRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSmartAGByAccessPointRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListSmartAGByAccessPointRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListSmartAGByAccessPointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListSmartAGByAccessPointRequest
	GetResourceOwnerId() *int64
	SetSmartAGStatus(v string) *ListSmartAGByAccessPointRequest
	GetSmartAGStatus() *string
}

type ListSmartAGByAccessPointRequest struct {
	// The ID of the access point.
	//
	// This parameter is required.
	//
	// example:
	//
	// 238
	AccessPointId *int32  `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **50*	- .
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// A region contains one or more access points. You can call the [ListAccessPoints](https://help.aliyun.com/document_detail/183876.html) operation to query access points in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the SAG instance. Valid values:
	//
	// 	- **Active**: The SAG device is connected to Alibaba Cloud.
	//
	// 	- **offline**: The SAG device is disconnected from Alibaba Cloud.
	//
	// example:
	//
	// Active
	SmartAGStatus *string `json:"SmartAGStatus,omitempty" xml:"SmartAGStatus,omitempty"`
}

func (s ListSmartAGByAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGByAccessPointRequest) GoString() string {
	return s.String()
}

func (s *ListSmartAGByAccessPointRequest) GetAccessPointId() *int32 {
	return s.AccessPointId
}

func (s *ListSmartAGByAccessPointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListSmartAGByAccessPointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSmartAGByAccessPointRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSmartAGByAccessPointRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSmartAGByAccessPointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSmartAGByAccessPointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListSmartAGByAccessPointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListSmartAGByAccessPointRequest) GetSmartAGStatus() *string {
	return s.SmartAGStatus
}

func (s *ListSmartAGByAccessPointRequest) SetAccessPointId(v int32) *ListSmartAGByAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetOwnerAccount(v string) *ListSmartAGByAccessPointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetOwnerId(v int64) *ListSmartAGByAccessPointRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetPageNumber(v int32) *ListSmartAGByAccessPointRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetPageSize(v int32) *ListSmartAGByAccessPointRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetRegionId(v string) *ListSmartAGByAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetResourceOwnerAccount(v string) *ListSmartAGByAccessPointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetResourceOwnerId(v int64) *ListSmartAGByAccessPointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) SetSmartAGStatus(v string) *ListSmartAGByAccessPointRequest {
	s.SmartAGStatus = &v
	return s
}

func (s *ListSmartAGByAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
