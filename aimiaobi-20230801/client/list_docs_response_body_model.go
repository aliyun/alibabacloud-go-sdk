// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDocsResponseBody
	GetCode() *string
	SetData(v []*ListDocsResponseBodyData) *ListDocsResponseBody
	GetData() []*ListDocsResponseBodyData
	SetHttpStatusCode(v int32) *ListDocsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListDocsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListDocsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListDocsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDocsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDocsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDocsResponseBody
	GetTotalCount() *int32
}

type ListDocsResponseBody struct {
	// example:
	//
	// successful
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListDocsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJANEQ4mYAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM4NzA3MjZjN2E2NDYyNzUzODMxMzY3ODM0NmIzNTZkNjc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 70
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDocsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDocsResponseBody) GetData() []*ListDocsResponseBodyData {
	return s.Data
}

func (s *ListDocsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDocsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDocsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDocsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDocsResponseBody) SetCode(v string) *ListDocsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDocsResponseBody) SetData(v []*ListDocsResponseBodyData) *ListDocsResponseBody {
	s.Data = v
	return s
}

func (s *ListDocsResponseBody) SetHttpStatusCode(v int32) *ListDocsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDocsResponseBody) SetMaxResults(v int32) *ListDocsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDocsResponseBody) SetMessage(v string) *ListDocsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDocsResponseBody) SetNextToken(v string) *ListDocsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDocsResponseBody) SetRequestId(v string) *ListDocsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocsResponseBody) SetSuccess(v bool) *ListDocsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDocsResponseBody) SetTotalCount(v int32) *ListDocsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDocsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDocsResponseBodyData struct {
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 12345
	DocId   *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// 0
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ListDocsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDocsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDocsResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListDocsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDocsResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *ListDocsResponseBodyData) GetDocName() *string {
	return s.DocName
}

func (s *ListDocsResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *ListDocsResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListDocsResponseBodyData) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListDocsResponseBodyData) SetCategoryId(v string) *ListDocsResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *ListDocsResponseBodyData) SetCreateTime(v int64) *ListDocsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDocsResponseBodyData) SetDocId(v string) *ListDocsResponseBodyData {
	s.DocId = &v
	return s
}

func (s *ListDocsResponseBodyData) SetDocName(v string) *ListDocsResponseBodyData {
	s.DocName = &v
	return s
}

func (s *ListDocsResponseBodyData) SetDocType(v string) *ListDocsResponseBodyData {
	s.DocType = &v
	return s
}

func (s *ListDocsResponseBodyData) SetStatus(v int32) *ListDocsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDocsResponseBodyData) SetStatusMessage(v string) *ListDocsResponseBodyData {
	s.StatusMessage = &v
	return s
}

func (s *ListDocsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
