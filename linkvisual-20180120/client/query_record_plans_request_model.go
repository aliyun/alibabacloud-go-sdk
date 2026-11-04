// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryRecordPlansRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryRecordPlansRequest
	GetPageSize() *int32
}

type QueryRecordPlansRequest struct {
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRecordPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlansRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryRecordPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRecordPlansRequest) SetCurrentPage(v int32) *QueryRecordPlansRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRecordPlansRequest) SetPageSize(v int32) *QueryRecordPlansRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlansRequest) Validate() error {
	return dara.Validate(s)
}
