// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeSourceResponseBody) *UpdateKnowledgeSourceResponse
	GetBody() *UpdateKnowledgeSourceResponseBody
}

type UpdateKnowledgeSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeSourceResponse) GetBody() *UpdateKnowledgeSourceResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeSourceResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeSourceResponse) SetStatusCode(v int32) *UpdateKnowledgeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeSourceResponse) SetBody(v *UpdateKnowledgeSourceResponseBody) *UpdateKnowledgeSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
