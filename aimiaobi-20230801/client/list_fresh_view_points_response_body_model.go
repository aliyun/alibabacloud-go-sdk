// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreshViewPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFreshViewPointsResponseBody
	GetCode() *string
	SetData(v []*ListFreshViewPointsResponseBodyData) *ListFreshViewPointsResponseBody
	GetData() []*ListFreshViewPointsResponseBodyData
	SetHttpStatusCode(v int32) *ListFreshViewPointsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListFreshViewPointsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListFreshViewPointsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListFreshViewPointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFreshViewPointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListFreshViewPointsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListFreshViewPointsResponseBody
	GetTotalCount() *int32
}

type ListFreshViewPointsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListFreshViewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 94
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
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
	// 26
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFreshViewPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFreshViewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFreshViewPointsResponseBody) GetData() []*ListFreshViewPointsResponseBodyData {
	return s.Data
}

func (s *ListFreshViewPointsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFreshViewPointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFreshViewPointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFreshViewPointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreshViewPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFreshViewPointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFreshViewPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFreshViewPointsResponseBody) SetCode(v string) *ListFreshViewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetData(v []*ListFreshViewPointsResponseBodyData) *ListFreshViewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetHttpStatusCode(v int32) *ListFreshViewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetMaxResults(v int32) *ListFreshViewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetMessage(v string) *ListFreshViewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetNextToken(v string) *ListFreshViewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetRequestId(v string) *ListFreshViewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetSuccess(v bool) *ListFreshViewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetTotalCount(v int32) *ListFreshViewPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFreshViewPointsResponseBodyData struct {
	Outlines []*ListFreshViewPointsResponseBodyDataOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListFreshViewPointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFreshViewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponseBodyData) GetOutlines() []*ListFreshViewPointsResponseBodyDataOutlines {
	return s.Outlines
}

func (s *ListFreshViewPointsResponseBodyData) GetPoint() *string {
	return s.Point
}

func (s *ListFreshViewPointsResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *ListFreshViewPointsResponseBodyData) SetOutlines(v []*ListFreshViewPointsResponseBodyDataOutlines) *ListFreshViewPointsResponseBodyData {
	s.Outlines = v
	return s
}

func (s *ListFreshViewPointsResponseBodyData) SetPoint(v string) *ListFreshViewPointsResponseBodyData {
	s.Point = &v
	return s
}

func (s *ListFreshViewPointsResponseBodyData) SetSummary(v string) *ListFreshViewPointsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListFreshViewPointsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListFreshViewPointsResponseBodyDataOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListFreshViewPointsResponseBodyDataOutlines) String() string {
	return dara.Prettify(s)
}

func (s ListFreshViewPointsResponseBodyDataOutlines) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponseBodyDataOutlines) GetOutline() *string {
	return s.Outline
}

func (s *ListFreshViewPointsResponseBodyDataOutlines) GetSummary() *string {
	return s.Summary
}

func (s *ListFreshViewPointsResponseBodyDataOutlines) SetOutline(v string) *ListFreshViewPointsResponseBodyDataOutlines {
	s.Outline = &v
	return s
}

func (s *ListFreshViewPointsResponseBodyDataOutlines) SetSummary(v string) *ListFreshViewPointsResponseBodyDataOutlines {
	s.Summary = &v
	return s
}

func (s *ListFreshViewPointsResponseBodyDataOutlines) Validate() error {
	return dara.Validate(s)
}
