// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *QueryFaceGroupRequest
	GetCurrentPage() *int64
	SetGroupCode(v string) *QueryFaceGroupRequest
	GetGroupCode() *string
	SetMaxResults(v int32) *QueryFaceGroupRequest
	GetMaxResults() *int32
	SetName(v string) *QueryFaceGroupRequest
	GetName() *string
	SetNextToken(v string) *QueryFaceGroupRequest
	GetNextToken() *string
	SetPageSize(v int32) *QueryFaceGroupRequest
	GetPageSize() *int32
}

type QueryFaceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// groupCode001
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// test008
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// WpY9RBGa5Vrzxi3+mp2Cdw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryFaceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceGroupRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryFaceGroupRequest) GetGroupCode() *string {
	return s.GroupCode
}

func (s *QueryFaceGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryFaceGroupRequest) GetName() *string {
	return s.Name
}

func (s *QueryFaceGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryFaceGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceGroupRequest) SetCurrentPage(v int64) *QueryFaceGroupRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryFaceGroupRequest) SetGroupCode(v string) *QueryFaceGroupRequest {
	s.GroupCode = &v
	return s
}

func (s *QueryFaceGroupRequest) SetMaxResults(v int32) *QueryFaceGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryFaceGroupRequest) SetName(v string) *QueryFaceGroupRequest {
	s.Name = &v
	return s
}

func (s *QueryFaceGroupRequest) SetNextToken(v string) *QueryFaceGroupRequest {
	s.NextToken = &v
	return s
}

func (s *QueryFaceGroupRequest) SetPageSize(v int32) *QueryFaceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *QueryFaceGroupRequest) Validate() error {
	return dara.Validate(s)
}
