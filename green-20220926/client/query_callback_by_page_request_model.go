// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallbackByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryCallbackByPageRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryCallbackByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *QueryCallbackByPageRequest
	GetRegionId() *string
}

type QueryCallbackByPageRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryCallbackByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCallbackByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryCallbackByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCallbackByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCallbackByPageRequest) SetCurrentPage(v int32) *QueryCallbackByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCallbackByPageRequest) SetPageSize(v int32) *QueryCallbackByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCallbackByPageRequest) SetRegionId(v string) *QueryCallbackByPageRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCallbackByPageRequest) Validate() error {
	return dara.Validate(s)
}
