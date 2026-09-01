// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseResponseBody
	GetKnowledgeBaseId() *string
	SetRequestId(v string) *UpdateKnowledgeBaseResponseBody
	GetRequestId() *string
}

type UpdateKnowledgeBaseResponseBody struct {
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
	// EB07CFF0-D8A4-5C76-AED7-D00E26FC2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseResponseBody) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseResponseBody) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseResponseBody {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
