// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamWatermarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeLiveStreamWatermarksRequest
	GetDomain() *string
	SetKeyWord(v string) *DescribeLiveStreamWatermarksRequest
	GetKeyWord() *string
	SetOwnerId(v int64) *DescribeLiveStreamWatermarksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLiveStreamWatermarksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveStreamWatermarksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLiveStreamWatermarksRequest
	GetRegionId() *string
}

type DescribeLiveStreamWatermarksRequest struct {
	// The name of the main streaming domain.
	//
	// example:
	//
	// live.yourdomain.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The keyword used in the query. You can specify the template ID or name. Fuzzy search is supported for the name.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of templates per page. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveStreamWatermarksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarksRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarksRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveStreamWatermarksRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeLiveStreamWatermarksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamWatermarksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveStreamWatermarksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamWatermarksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamWatermarksRequest) SetDomain(v string) *DescribeLiveStreamWatermarksRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLiveStreamWatermarksRequest) SetKeyWord(v string) *DescribeLiveStreamWatermarksRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeLiveStreamWatermarksRequest) SetOwnerId(v int64) *DescribeLiveStreamWatermarksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamWatermarksRequest) SetPageNumber(v int32) *DescribeLiveStreamWatermarksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveStreamWatermarksRequest) SetPageSize(v int32) *DescribeLiveStreamWatermarksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamWatermarksRequest) SetRegionId(v string) *DescribeLiveStreamWatermarksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamWatermarksRequest) Validate() error {
	return dara.Validate(s)
}
