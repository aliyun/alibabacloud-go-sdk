// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryId(v string) *AnswerKnowledgeBaseResponseBody
	GetQueryId() *string
	SetRequestId(v string) *AnswerKnowledgeBaseResponseBody
	GetRequestId() *string
	SetStatus(v string) *AnswerKnowledgeBaseResponseBody
	GetStatus() *string
}

type AnswerKnowledgeBaseResponseBody struct {
	// The unique ID of the Q&A task.
	//
	// example:
	//
	// R3BGbnBqcXN******.2a5a23c9-******-179970533d30
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the Q&A task.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AnswerKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnswerKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerKnowledgeBaseResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *AnswerKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnswerKnowledgeBaseResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AnswerKnowledgeBaseResponseBody) SetQueryId(v string) *AnswerKnowledgeBaseResponseBody {
	s.QueryId = &v
	return s
}

func (s *AnswerKnowledgeBaseResponseBody) SetRequestId(v string) *AnswerKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerKnowledgeBaseResponseBody) SetStatus(v string) *AnswerKnowledgeBaseResponseBody {
	s.Status = &v
	return s
}

func (s *AnswerKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
