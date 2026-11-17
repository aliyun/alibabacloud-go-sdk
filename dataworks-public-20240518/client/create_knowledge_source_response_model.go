// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeSourceResponseBody) *CreateKnowledgeSourceResponse
	GetBody() *CreateKnowledgeSourceResponseBody
}

type CreateKnowledgeSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeSourceResponse) GetBody() *CreateKnowledgeSourceResponseBody {
	return s.Body
}

func (s *CreateKnowledgeSourceResponse) SetHeaders(v map[string]*string) *CreateKnowledgeSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeSourceResponse) SetStatusCode(v int32) *CreateKnowledgeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeSourceResponse) SetBody(v *CreateKnowledgeSourceResponseBody) *CreateKnowledgeSourceResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
