// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRetcodeAppByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *SearchRetcodeAppByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRetcodeAppByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *SearchRetcodeAppByPageRequest
	GetRegionId() *string
	SetRetcodeAppName(v string) *SearchRetcodeAppByPageRequest
	GetRetcodeAppName() *string
}

type SearchRetcodeAppByPageRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
}

func (s SearchRetcodeAppByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRetcodeAppByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRetcodeAppByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchRetcodeAppByPageRequest) GetRetcodeAppName() *string {
	return s.RetcodeAppName
}

func (s *SearchRetcodeAppByPageRequest) SetPageNumber(v int32) *SearchRetcodeAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetPageSize(v int32) *SearchRetcodeAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRegionId(v string) *SearchRetcodeAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRetcodeAppName(v string) *SearchRetcodeAppByPageRequest {
	s.RetcodeAppName = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) Validate() error {
	return dara.Validate(s)
}
