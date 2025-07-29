// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisTagDetailByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetCode() *string
	SetData(v []*ListAnalysisTagDetailByTaskIdResponseBodyData) *ListAnalysisTagDetailByTaskIdResponseBody
	GetData() []*ListAnalysisTagDetailByTaskIdResponseBodyData
	SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAnalysisTagDetailByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAnalysisTagDetailByTaskIdResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListAnalysisTagDetailByTaskIdResponseBody
	GetTotalCount() *int32
}

type ListAnalysisTagDetailByTaskIdResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListAnalysisTagDetailByTaskIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 成功
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAnalysisTagDetailByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetData() []*ListAnalysisTagDetailByTaskIdResponseBodyData {
	return s.Data
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetCode(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetData(v []*ListAnalysisTagDetailByTaskIdResponseBodyData) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetMessage(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetNextToken(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetRequestId(v string) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetSuccess(v bool) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) SetTotalCount(v int32) *ListAnalysisTagDetailByTaskIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAnalysisTagDetailByTaskIdResponseBodyData struct {
	// example:
	//
	// xxxx
	Content        *string                                                     `json:"content,omitempty" xml:"content,omitempty"`
	ContentTags    []*ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags `json:"contentTags,omitempty" xml:"contentTags,omitempty" type:"Repeated"`
	OriginResponse *string                                                     `json:"originResponse,omitempty" xml:"originResponse,omitempty"`
	SourceList     []*string                                                   `json:"sourceList,omitempty" xml:"sourceList,omitempty" type:"Repeated"`
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetContentTags() []*ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags {
	return s.ContentTags
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetOriginResponse() *string {
	return s.OriginResponse
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) GetSourceList() []*string {
	return s.SourceList
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetContent(v string) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetContentTags(v []*ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.ContentTags = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetOriginResponse(v string) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.OriginResponse = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) SetSourceList(v []*string) *ListAnalysisTagDetailByTaskIdResponseBodyData {
	s.SourceList = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags struct {
	TagName *string   `json:"tagName,omitempty" xml:"tagName,omitempty"`
	Tags    []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) GetTagName() *string {
	return s.TagName
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) GetTags() []*string {
	return s.Tags
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) SetTagName(v string) *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags {
	s.TagName = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) SetTags(v []*string) *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags {
	s.Tags = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponseBodyDataContentTags) Validate() error {
	return dara.Validate(s)
}
