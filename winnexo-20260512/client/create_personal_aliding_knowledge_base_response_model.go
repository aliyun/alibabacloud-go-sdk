// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalAlidingKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalAlidingKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalAlidingKnowledgeBaseResponseBody) *CreatePersonalAlidingKnowledgeBaseResponse
	GetBody() *CreatePersonalAlidingKnowledgeBaseResponseBody
}

type CreatePersonalAlidingKnowledgeBaseResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalAlidingKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalAlidingKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalAlidingKnowledgeBaseResponse) GetBody() *CreatePersonalAlidingKnowledgeBaseResponseBody {
	return s.Body
}

func (s *CreatePersonalAlidingKnowledgeBaseResponse) SetHeaders(v map[string]*string) *CreatePersonalAlidingKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponse) SetStatusCode(v int32) *CreatePersonalAlidingKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponse) SetBody(v *CreatePersonalAlidingKnowledgeBaseResponseBody) *CreatePersonalAlidingKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
