// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckColumnResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ListDataCheckColumnResultsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckColumnResultsRequest
	GetPageSize() *int32
	SetResultId(v string) *ListDataCheckColumnResultsRequest
	GetResultId() *string
}

type ListDataCheckColumnResultsRequest struct {
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The validation result ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001
	ResultId *string `json:"resultId,omitempty" xml:"resultId,omitempty"`
}

func (s ListDataCheckColumnResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckColumnResultsRequest) GoString() string {
	return s.String()
}

func (s *ListDataCheckColumnResultsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckColumnResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckColumnResultsRequest) GetResultId() *string {
	return s.ResultId
}

func (s *ListDataCheckColumnResultsRequest) SetPageIndex(v int32) *ListDataCheckColumnResultsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckColumnResultsRequest) SetPageSize(v int32) *ListDataCheckColumnResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckColumnResultsRequest) SetResultId(v string) *ListDataCheckColumnResultsRequest {
	s.ResultId = &v
	return s
}

func (s *ListDataCheckColumnResultsRequest) Validate() error {
	return dara.Validate(s)
}
