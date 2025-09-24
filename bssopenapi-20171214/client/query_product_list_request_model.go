// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *QueryProductListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryProductListRequest
	GetPageSize() *int32
	SetQueryTotalCount(v bool) *QueryProductListRequest
	GetQueryTotalCount() *bool
}

type QueryProductListRequest struct {
	// The page number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to return the total number of services. Default value: false.
	//
	// example:
	//
	// true
	QueryTotalCount *bool `json:"QueryTotalCount,omitempty" xml:"QueryTotalCount,omitempty"`
}

func (s QueryProductListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryProductListRequest) GoString() string {
	return s.String()
}

func (s *QueryProductListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryProductListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryProductListRequest) GetQueryTotalCount() *bool {
	return s.QueryTotalCount
}

func (s *QueryProductListRequest) SetPageNum(v int32) *QueryProductListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryProductListRequest) SetPageSize(v int32) *QueryProductListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryProductListRequest) SetQueryTotalCount(v bool) *QueryProductListRequest {
	s.QueryTotalCount = &v
	return s
}

func (s *QueryProductListRequest) Validate() error {
	return dara.Validate(s)
}
