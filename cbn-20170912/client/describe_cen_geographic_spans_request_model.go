// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenGeographicSpansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGeographicSpanId(v string) *DescribeCenGeographicSpansRequest
	GetGeographicSpanId() *string
	SetOwnerAccount(v string) *DescribeCenGeographicSpansRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCenGeographicSpansRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCenGeographicSpansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenGeographicSpansRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeCenGeographicSpansRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCenGeographicSpansRequest
	GetResourceOwnerId() *int64
}

type DescribeCenGeographicSpansRequest struct {
	// The ID of the geographic span.
	//
	// > If you do not specify this parameter, the system queries all geographic spans that are supported by Cloud Enterprise Network (CEN).
	//
	// example:
	//
	// china_asia-pacific
	GeographicSpanId *string `json:"GeographicSpanId,omitempty" xml:"GeographicSpanId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCenGeographicSpansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpansRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansRequest) GetGeographicSpanId() *string {
	return s.GeographicSpanId
}

func (s *DescribeCenGeographicSpansRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCenGeographicSpansRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCenGeographicSpansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenGeographicSpansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenGeographicSpansRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCenGeographicSpansRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCenGeographicSpansRequest) SetGeographicSpanId(v string) *DescribeCenGeographicSpansRequest {
	s.GeographicSpanId = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetOwnerAccount(v string) *DescribeCenGeographicSpansRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetOwnerId(v int64) *DescribeCenGeographicSpansRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetPageNumber(v int32) *DescribeCenGeographicSpansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetPageSize(v int32) *DescribeCenGeographicSpansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetResourceOwnerAccount(v string) *DescribeCenGeographicSpansRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) SetResourceOwnerId(v int64) *DescribeCenGeographicSpansRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCenGeographicSpansRequest) Validate() error {
	return dara.Validate(s)
}
