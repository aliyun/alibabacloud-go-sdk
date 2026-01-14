// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBases(v []*KnowledgeBase) *ListKnowledgeBasesResponseBody
	GetKnowledgeBases() []*KnowledgeBase
	SetMaxResults(v int32) *ListKnowledgeBasesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBasesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListKnowledgeBasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKnowledgeBasesResponseBody
	GetTotalCount() *int32
}

type ListKnowledgeBasesResponseBody struct {
	KnowledgeBases []*KnowledgeBase `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 11
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 35686626-8D83-5ADE-B400-08A6613A6057
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponseBody) GetKnowledgeBases() []*KnowledgeBase {
	return s.KnowledgeBases
}

func (s *ListKnowledgeBasesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBasesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKnowledgeBasesResponseBody) SetKnowledgeBases(v []*KnowledgeBase) *ListKnowledgeBasesResponseBody {
	s.KnowledgeBases = v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetMaxResults(v int32) *ListKnowledgeBasesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetNextToken(v string) *ListKnowledgeBasesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetRequestId(v string) *ListKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetTotalCount(v int32) *ListKnowledgeBasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) Validate() error {
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
