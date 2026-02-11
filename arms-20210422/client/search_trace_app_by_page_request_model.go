// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *SearchTraceAppByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchTraceAppByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *SearchTraceAppByPageRequest
	GetRegionId() *string
	SetTraceAppName(v string) *SearchTraceAppByPageRequest
	GetTraceAppName() *string
}

type SearchTraceAppByPageRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTraceAppByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTraceAppByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByPageRequest) GetTraceAppName() *string {
	return s.TraceAppName
}

func (s *SearchTraceAppByPageRequest) SetPageNumber(v int32) *SearchTraceAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetPageSize(v int32) *SearchTraceAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetRegionId(v string) *SearchTraceAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetTraceAppName(v string) *SearchTraceAppByPageRequest {
	s.TraceAppName = &v
	return s
}

func (s *SearchTraceAppByPageRequest) Validate() error {
	return dara.Validate(s)
}
