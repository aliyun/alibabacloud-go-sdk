// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotSourcesResponseBody
	GetCode() *string
	SetData(v []*ListHotSourcesResponseBodyData) *ListHotSourcesResponseBody
	GetData() []*ListHotSourcesResponseBodyData
	SetHttpStatusCode(v int32) *ListHotSourcesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListHotSourcesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListHotSourcesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListHotSourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHotSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHotSourcesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListHotSourcesResponseBody
	GetTotalCount() *int32
}

type ListHotSourcesResponseBody struct {
	// Status code
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Third-party source list
	Data []*ListHotSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Maximum number of results returned
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Fault description
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Next page token
	//
	// example:
	//
	// xxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request UUID
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded: true for success, false for failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHotSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotSourcesResponseBody) GetData() []*ListHotSourcesResponseBodyData {
	return s.Data
}

func (s *ListHotSourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHotSourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotSourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHotSourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHotSourcesResponseBody) SetCode(v string) *ListHotSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetData(v []*ListHotSourcesResponseBodyData) *ListHotSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListHotSourcesResponseBody) SetHttpStatusCode(v int32) *ListHotSourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetMaxResults(v int32) *ListHotSourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetMessage(v string) *ListHotSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetNextToken(v string) *ListHotSourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetRequestId(v string) *ListHotSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetSuccess(v bool) *ListHotSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetTotalCount(v int32) *ListHotSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHotSourcesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotSourcesResponseBodyData struct {
	// Hot ranking source description
	//
	// example:
	//
	// 热榜源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether to display in the console
	//
	// example:
	//
	// true
	Show *bool `json:"Show,omitempty" xml:"Show,omitempty"`
	// Sorting value
	//
	// example:
	//
	// 86
	Sort *int32 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Hot ranking source identity
	//
	// example:
	//
	// 热榜源标识
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListHotSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHotSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotSourcesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListHotSourcesResponseBodyData) GetShow() *bool {
	return s.Show
}

func (s *ListHotSourcesResponseBodyData) GetSort() *int32 {
	return s.Sort
}

func (s *ListHotSourcesResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListHotSourcesResponseBodyData) SetDescription(v string) *ListHotSourcesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListHotSourcesResponseBodyData) SetShow(v bool) *ListHotSourcesResponseBodyData {
	s.Show = &v
	return s
}

func (s *ListHotSourcesResponseBodyData) SetSort(v int32) *ListHotSourcesResponseBodyData {
	s.Sort = &v
	return s
}

func (s *ListHotSourcesResponseBodyData) SetSource(v string) *ListHotSourcesResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListHotSourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
