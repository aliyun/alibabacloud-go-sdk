// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllEndEntityInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAllEndEntityInstanceResponseBody
	GetCurrentPage() *int32
	SetList(v []map[string]interface{}) *ListAllEndEntityInstanceResponseBody
	GetList() []map[string]interface{}
	SetMaxResults(v int32) *ListAllEndEntityInstanceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAllEndEntityInstanceResponseBody
	GetNextToken() *string
	SetPageCount(v int32) *ListAllEndEntityInstanceResponseBody
	GetPageCount() *int32
	SetRequestId(v string) *ListAllEndEntityInstanceResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListAllEndEntityInstanceResponseBody
	GetShowSize() *int32
	SetTotalCount(v int32) *ListAllEndEntityInstanceResponseBody
	GetTotalCount() *int32
}

type ListAllEndEntityInstanceResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []map[string]interface{} `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 443C05A8-6C16-52B5-BB97-5D8798F7A49A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAllEndEntityInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllEndEntityInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllEndEntityInstanceResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAllEndEntityInstanceResponseBody) GetList() []map[string]interface{} {
	return s.List
}

func (s *ListAllEndEntityInstanceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAllEndEntityInstanceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAllEndEntityInstanceResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *ListAllEndEntityInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllEndEntityInstanceResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListAllEndEntityInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAllEndEntityInstanceResponseBody) SetCurrentPage(v int32) *ListAllEndEntityInstanceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) SetList(v []map[string]interface{}) *ListAllEndEntityInstanceResponseBody {
	s.List = v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) SetMaxResults(v int32) *ListAllEndEntityInstanceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) SetNextToken(v string) *ListAllEndEntityInstanceResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) SetPageCount(v int32) *ListAllEndEntityInstanceResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) SetRequestId(v string) *ListAllEndEntityInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) SetShowSize(v int32) *ListAllEndEntityInstanceResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) SetTotalCount(v int32) *ListAllEndEntityInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAllEndEntityInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
