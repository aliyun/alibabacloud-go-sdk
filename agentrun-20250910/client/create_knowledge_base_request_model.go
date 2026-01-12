// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateKnowledgeBaseInput) *CreateKnowledgeBaseRequest
	GetBody() *CreateKnowledgeBaseInput
}

type CreateKnowledgeBaseRequest struct {
	Body *CreateKnowledgeBaseInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequest) GetBody() *CreateKnowledgeBaseInput {
	return s.Body
}

func (s *CreateKnowledgeBaseRequest) SetBody(v *CreateKnowledgeBaseInput) *CreateKnowledgeBaseRequest {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
