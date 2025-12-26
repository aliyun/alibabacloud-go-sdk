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
	SetWorkspaceId(v string) *CreateKnowledgeBaseResponseBody
	GetWorkspaceId() *string
}

type CreateKnowledgeBaseResponseBody struct {
	// example:
	//
	// d-ksicx823d
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// example:
	//
	// 48E6392E-C3C9-5212-9FAD-13256ABD9AF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *CreateKnowledgeBaseResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateKnowledgeBaseResponseBody) SetKnowledgeBaseId(v string) *CreateKnowledgeBaseResponseBody {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) SetRequestId(v string) *CreateKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) SetWorkspaceId(v string) *CreateKnowledgeBaseResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}
