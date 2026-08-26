// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *CreateKnowledgeBaseResponseBody
	GetKnowledgeBaseId() *string
	SetRequestId(v string) *CreateKnowledgeBaseResponseBody
	GetRequestId() *string
}

type CreateKnowledgeBaseResponseBody struct {
	// The unique identifier of the knowledge base.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseResponseBody) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseResponseBody) SetKnowledgeBaseId(v string) *CreateKnowledgeBaseResponseBody {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) SetRequestId(v string) *CreateKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
