// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterKnowledgeBaseFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterKnowledgeBaseFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterKnowledgeBaseFileResponse
	GetStatusCode() *int32
	SetBody(v *RegisterKnowledgeBaseFileResponseBody) *RegisterKnowledgeBaseFileResponse
	GetBody() *RegisterKnowledgeBaseFileResponseBody
}

type RegisterKnowledgeBaseFileResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterKnowledgeBaseFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterKnowledgeBaseFileResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterKnowledgeBaseFileResponse) GoString() string {
	return s.String()
}

func (s *RegisterKnowledgeBaseFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterKnowledgeBaseFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterKnowledgeBaseFileResponse) GetBody() *RegisterKnowledgeBaseFileResponseBody {
	return s.Body
}

func (s *RegisterKnowledgeBaseFileResponse) SetHeaders(v map[string]*string) *RegisterKnowledgeBaseFileResponse {
	s.Headers = v
	return s
}

func (s *RegisterKnowledgeBaseFileResponse) SetStatusCode(v int32) *RegisterKnowledgeBaseFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterKnowledgeBaseFileResponse) SetBody(v *RegisterKnowledgeBaseFileResponseBody) *RegisterKnowledgeBaseFileResponse {
	s.Body = v
	return s
}

func (s *RegisterKnowledgeBaseFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
