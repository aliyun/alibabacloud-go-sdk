// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveHttpsDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeLiveHttpsDomainListRequest
	GetKeyword() *string
	SetOwnerId(v int64) *DescribeLiveHttpsDomainListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLiveHttpsDomainListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveHttpsDomainListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLiveHttpsDomainListRequest
	GetRegionId() *string
}

type DescribeLiveHttpsDomainListRequest struct {
	// The accelerated domain name.
	//
	// example:
	//
	// demo.aliyun.com
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: **1 to 10000**.
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveHttpsDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveHttpsDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveHttpsDomainListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeLiveHttpsDomainListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveHttpsDomainListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveHttpsDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveHttpsDomainListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveHttpsDomainListRequest) SetKeyword(v string) *DescribeLiveHttpsDomainListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeLiveHttpsDomainListRequest) SetOwnerId(v int64) *DescribeLiveHttpsDomainListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveHttpsDomainListRequest) SetPageNumber(v int32) *DescribeLiveHttpsDomainListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveHttpsDomainListRequest) SetPageSize(v int32) *DescribeLiveHttpsDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveHttpsDomainListRequest) SetRegionId(v string) *DescribeLiveHttpsDomainListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveHttpsDomainListRequest) Validate() error {
	return dara.Validate(s)
}
