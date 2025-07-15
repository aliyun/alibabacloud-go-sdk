// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlanningProposalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPlanningProposalResponseBody
	GetCode() *string
	SetData(v []*ListPlanningProposalResponseBodyData) *ListPlanningProposalResponseBody
	GetData() []*ListPlanningProposalResponseBodyData
	SetHttpStatusCode(v int32) *ListPlanningProposalResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListPlanningProposalResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListPlanningProposalResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListPlanningProposalResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPlanningProposalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPlanningProposalResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListPlanningProposalResponseBody
	GetTotalCount() *int32
}

type ListPlanningProposalResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListPlanningProposalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 77
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
	// 80
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPlanningProposalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPlanningProposalResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPlanningProposalResponseBody) GetData() []*ListPlanningProposalResponseBodyData {
	return s.Data
}

func (s *ListPlanningProposalResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPlanningProposalResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPlanningProposalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPlanningProposalResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPlanningProposalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPlanningProposalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPlanningProposalResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPlanningProposalResponseBody) SetCode(v string) *ListPlanningProposalResponseBody {
	s.Code = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetData(v []*ListPlanningProposalResponseBodyData) *ListPlanningProposalResponseBody {
	s.Data = v
	return s
}

func (s *ListPlanningProposalResponseBody) SetHttpStatusCode(v int32) *ListPlanningProposalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetMaxResults(v int32) *ListPlanningProposalResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetMessage(v string) *ListPlanningProposalResponseBody {
	s.Message = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetNextToken(v string) *ListPlanningProposalResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetRequestId(v string) *ListPlanningProposalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetSuccess(v bool) *ListPlanningProposalResponseBody {
	s.Success = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetTotalCount(v int32) *ListPlanningProposalResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPlanningProposalResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPlanningProposalResponseBodyData struct {
	Outlines []*ListPlanningProposalResponseBodyDataOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	Summary  *string                                         `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title    *string                                         `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListPlanningProposalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPlanningProposalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponseBodyData) GetOutlines() []*ListPlanningProposalResponseBodyDataOutlines {
	return s.Outlines
}

func (s *ListPlanningProposalResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *ListPlanningProposalResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListPlanningProposalResponseBodyData) SetOutlines(v []*ListPlanningProposalResponseBodyDataOutlines) *ListPlanningProposalResponseBodyData {
	s.Outlines = v
	return s
}

func (s *ListPlanningProposalResponseBodyData) SetSummary(v string) *ListPlanningProposalResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListPlanningProposalResponseBodyData) SetTitle(v string) *ListPlanningProposalResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListPlanningProposalResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListPlanningProposalResponseBodyDataOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListPlanningProposalResponseBodyDataOutlines) String() string {
	return dara.Prettify(s)
}

func (s ListPlanningProposalResponseBodyDataOutlines) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponseBodyDataOutlines) GetOutline() *string {
	return s.Outline
}

func (s *ListPlanningProposalResponseBodyDataOutlines) GetSummary() *string {
	return s.Summary
}

func (s *ListPlanningProposalResponseBodyDataOutlines) SetOutline(v string) *ListPlanningProposalResponseBodyDataOutlines {
	s.Outline = &v
	return s
}

func (s *ListPlanningProposalResponseBodyDataOutlines) SetSummary(v string) *ListPlanningProposalResponseBodyDataOutlines {
	s.Summary = &v
	return s
}

func (s *ListPlanningProposalResponseBodyDataOutlines) Validate() error {
	return dara.Validate(s)
}
