// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVersionPageListRequest
	GetLang() *string
	SetMaxResults(v int32) *DescribeVersionPageListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVersionPageListRequest
	GetNextToken() *string
	SetCurrentPage(v int32) *DescribeVersionPageListRequest
	GetCurrentPage() *int32
	SetObjectCode(v string) *DescribeVersionPageListRequest
	GetObjectCode() *string
	SetObjectId(v int64) *DescribeVersionPageListRequest
	GetObjectId() *int64
	SetPageSize(v int32) *DescribeVersionPageListRequest
	GetPageSize() *int32
	SetPaging(v bool) *DescribeVersionPageListRequest
	GetPaging() *bool
	SetRegId(v string) *DescribeVersionPageListRequest
	GetRegId() *string
	SetType(v string) *DescribeVersionPageListRequest
	GetType() *string
}

type DescribeVersionPageListRequest struct {
	// Sets the language type for the request and response messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Maximum number of results to be read in this call, with a default value of 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Used to mark the starting position for reading. An empty value indicates starting from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Name of the variable.
	//
	// example:
	//
	// ex_OERlw0Zqfb23
	ObjectCode *string `json:"objectCode,omitempty" xml:"objectCode,omitempty"`
	// Primary key ID of the variable.
	//
	// example:
	//
	// 392023
	ObjectId *int64 `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// Number of items per page, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Pagination flag, with a default value of true.
	//
	// example:
	//
	// true
	Paging *bool `json:"paging,omitempty" xml:"paging,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Type.
	//
	// example:
	//
	// EXPRESSION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeVersionPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVersionPageListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVersionPageListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVersionPageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVersionPageListRequest) GetObjectCode() *string {
	return s.ObjectCode
}

func (s *DescribeVersionPageListRequest) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *DescribeVersionPageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVersionPageListRequest) GetPaging() *bool {
	return s.Paging
}

func (s *DescribeVersionPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVersionPageListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVersionPageListRequest) SetLang(v string) *DescribeVersionPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetMaxResults(v int32) *DescribeVersionPageListRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetNextToken(v string) *DescribeVersionPageListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetCurrentPage(v int32) *DescribeVersionPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetObjectCode(v string) *DescribeVersionPageListRequest {
	s.ObjectCode = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetObjectId(v int64) *DescribeVersionPageListRequest {
	s.ObjectId = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetPageSize(v int32) *DescribeVersionPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetPaging(v bool) *DescribeVersionPageListRequest {
	s.Paging = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetRegId(v string) *DescribeVersionPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVersionPageListRequest) SetType(v string) *DescribeVersionPageListRequest {
	s.Type = &v
	return s
}

func (s *DescribeVersionPageListRequest) Validate() error {
	return dara.Validate(s)
}
