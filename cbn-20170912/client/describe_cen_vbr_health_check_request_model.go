// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenVbrHealthCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeCenVbrHealthCheckRequest
	GetCenId() *string
	SetOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenVbrHealthCheckRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenVbrHealthCheckRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenVbrHealthCheckRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenVbrHealthCheckRequest
	GetResourceOwnerId() *int64
	SetVbrInstanceId(v string) *DescribeCenVbrHealthCheckRequest
	GetVbrInstanceId() *string
	SetVbrInstanceOwnerId(v int64) *DescribeCenVbrHealthCheckRequest
	GetVbrInstanceOwnerId() *int64
	SetVbrInstanceRegionId(v string) *DescribeCenVbrHealthCheckRequest
	GetVbrInstanceRegionId() *string
}

type DescribeCenVbrHealthCheckRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-6hpdgj7ni6pz1k****
	CenId        *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-bp1kznorjeembsuhl****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VBRs.
	//
	// example:
	//
	// 1250123456123456
	VbrInstanceOwnerId *int64 `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
	// The ID of the region where the VBRs are deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrInstanceRegionId *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s DescribeCenVbrHealthCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenVbrHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenVbrHealthCheckRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenVbrHealthCheckRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenVbrHealthCheckRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenVbrHealthCheckRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenVbrHealthCheckRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenVbrHealthCheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenVbrHealthCheckRequest) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *DescribeCenVbrHealthCheckRequest) GetVbrInstanceOwnerId() *int64 {
	return s.VbrInstanceOwnerId
}

func (s *DescribeCenVbrHealthCheckRequest) GetVbrInstanceRegionId() *string {
	return s.VbrInstanceRegionId
}

func (s *DescribeCenVbrHealthCheckRequest) SetCenId(v string) *DescribeCenVbrHealthCheckRequest {
	s.CenId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetPageNumber(v int32) *DescribeCenVbrHealthCheckRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetPageSize(v int32) *DescribeCenVbrHealthCheckRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetResourceOwnerAccount(v string) *DescribeCenVbrHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetResourceOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *DescribeCenVbrHealthCheckRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *DescribeCenVbrHealthCheckRequest {
	s.VbrInstanceOwnerId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *DescribeCenVbrHealthCheckRequest {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckRequest) Validate() error {
	return dara.Validate(s)
}
