// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaKnowledgeBasesResponseBody
	GetCode() *string
	SetData(v *ListLumaKnowledgeBasesResponseBodyData) *ListLumaKnowledgeBasesResponseBody
	GetData() *ListLumaKnowledgeBasesResponseBodyData
	SetMessage(v string) *ListLumaKnowledgeBasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaKnowledgeBasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaKnowledgeBasesResponseBody
	GetSuccess() *bool
}

type ListLumaKnowledgeBasesResponseBody struct {
	// The response code. A value of Success indicates that the call was successful. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of knowledge bases bound to the agent, including entries and pagination information.
	Data *ListLumaKnowledgeBasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value Operation success is returned if the call was successful. A specific error description is returned if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. Use this ID for troubleshooting and when you submit a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates that the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaKnowledgeBasesResponseBody) GetData() *ListLumaKnowledgeBasesResponseBodyData {
	return s.Data
}

func (s *ListLumaKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaKnowledgeBasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaKnowledgeBasesResponseBody) SetCode(v string) *ListLumaKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetData(v *ListLumaKnowledgeBasesResponseBodyData) *ListLumaKnowledgeBasesResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetMessage(v string) *ListLumaKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetRequestId(v string) *ListLumaKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) SetSuccess(v bool) *ListLumaKnowledgeBasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaKnowledgeBasesResponseBodyData struct {
	// The list of knowledge bases bound to the agent.
	//
	// example:
	//
	// [{"KnowledgeBaseName":"my-knowledge-base"}]
	KnowledgeBases []*KnowledgeBase `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty" type:"Repeated"`
	// The maximum number of results per page that takes effect for this request. If MaxResults is not specified, this value is the server default. If the specified value exceeds the upper limit, this value is the adjusted value.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page (an opaque string). Pass this value as the NextToken parameter in the next request to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of knowledge bases bound to the agent, regardless of the number of entries returned on the current page.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLumaKnowledgeBasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaKnowledgeBasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaKnowledgeBasesResponseBodyData) GetKnowledgeBases() []*KnowledgeBase {
	return s.KnowledgeBases
}

func (s *ListLumaKnowledgeBasesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLumaKnowledgeBasesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaKnowledgeBasesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLumaKnowledgeBasesResponseBodyData) SetKnowledgeBases(v []*KnowledgeBase) *ListLumaKnowledgeBasesResponseBodyData {
	s.KnowledgeBases = v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBodyData) SetMaxResults(v int32) *ListLumaKnowledgeBasesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBodyData) SetNextToken(v string) *ListLumaKnowledgeBasesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBodyData) SetTotalCount(v int32) *ListLumaKnowledgeBasesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponseBodyData) Validate() error {
	if s.KnowledgeBases != nil {
		for _, item := range s.KnowledgeBases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
