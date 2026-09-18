// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBasesResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeBasesResponseBodyData) *ListKnowledgeBasesResponseBody
	GetData() *ListKnowledgeBasesResponseBodyData
	SetMessage(v string) *ListKnowledgeBasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeBasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeBasesResponseBody
	GetSuccess() *bool
}

type ListKnowledgeBasesResponseBody struct {
	// The response code. A value of Success indicates that the call was successful. If the call failed, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result of the knowledge base list, including knowledge base entries and pagination information.
	Data *ListKnowledgeBasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value Operation success is returned if the call was successful. A specific error description is returned if the call failed.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBasesResponseBody) GetData() *ListKnowledgeBasesResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeBasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeBasesResponseBody) SetCode(v string) *ListKnowledgeBasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetData(v *ListKnowledgeBasesResponseBodyData) *ListKnowledgeBasesResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetMessage(v string) *ListKnowledgeBasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetRequestId(v string) *ListKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetSuccess(v bool) *ListKnowledgeBasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBasesResponseBodyData struct {
	// The list of knowledge base entries.
	//
	// example:
	//
	// [{"KnowledgeBaseName":"my-knowledge-base","Status":"ACTIVE"}]
	KnowledgeBases []*KnowledgeBase `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty" type:"Repeated"`
	// The maximum number of results per page that was applied to this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token (an opaque string) for the next page. Pass this value as the NextToken parameter in the next request to retrieve the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of knowledge bases that match the filter conditions. Use an empty NextToken value as the termination condition for pagination.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKnowledgeBasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponseBodyData) GetKnowledgeBases() []*KnowledgeBase {
	return s.KnowledgeBases
}

func (s *ListKnowledgeBasesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKnowledgeBasesResponseBodyData) SetKnowledgeBases(v []*KnowledgeBase) *ListKnowledgeBasesResponseBodyData {
	s.KnowledgeBases = v
	return s
}

func (s *ListKnowledgeBasesResponseBodyData) SetMaxResults(v int32) *ListKnowledgeBasesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyData) SetNextToken(v string) *ListKnowledgeBasesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyData) SetTotalCount(v int32) *ListKnowledgeBasesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListKnowledgeBasesResponseBodyData) Validate() error {
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
