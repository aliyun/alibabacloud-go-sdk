// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateKnowledgeBaseInput) *UpdateKnowledgeBaseRequest
	GetBody() *UpdateKnowledgeBaseInput
}

type UpdateKnowledgeBaseRequest struct {
	Body *UpdateKnowledgeBaseInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequest) GetBody() *UpdateKnowledgeBaseInput {
	return s.Body
}

func (s *UpdateKnowledgeBaseRequest) SetBody(v *UpdateKnowledgeBaseInput) *UpdateKnowledgeBaseRequest {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
