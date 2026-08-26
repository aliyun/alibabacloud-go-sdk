// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseResponseBody
	GetKnowledgeBaseId() *string
	SetRequestId(v string) *DeleteKnowledgeBaseResponseBody
	GetRequestId() *string
}

type DeleteKnowledgeBaseResponseBody struct {
	// The unique ID of the knowledge base.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseResponseBody) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKnowledgeBaseResponseBody) SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseResponseBody {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) SetRequestId(v string) *DeleteKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
