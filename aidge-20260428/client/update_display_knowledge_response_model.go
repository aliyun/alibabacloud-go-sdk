// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisplayKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDisplayKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDisplayKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDisplayKnowledgeResponseBody) *UpdateDisplayKnowledgeResponse
	GetBody() *UpdateDisplayKnowledgeResponseBody
}

type UpdateDisplayKnowledgeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDisplayKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDisplayKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisplayKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDisplayKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDisplayKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDisplayKnowledgeResponse) GetBody() *UpdateDisplayKnowledgeResponseBody {
	return s.Body
}

func (s *UpdateDisplayKnowledgeResponse) SetHeaders(v map[string]*string) *UpdateDisplayKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDisplayKnowledgeResponse) SetStatusCode(v int32) *UpdateDisplayKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDisplayKnowledgeResponse) SetBody(v *UpdateDisplayKnowledgeResponseBody) *UpdateDisplayKnowledgeResponse {
	s.Body = v
	return s
}

func (s *UpdateDisplayKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
