// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseAdministratorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseAdministratorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseAdministratorsResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseAdministratorsResponseBody) *ListKnowledgeBaseAdministratorsResponse
	GetBody() *ListKnowledgeBaseAdministratorsResponseBody
}

type ListKnowledgeBaseAdministratorsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseAdministratorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseAdministratorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseAdministratorsResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseAdministratorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseAdministratorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseAdministratorsResponse) GetBody() *ListKnowledgeBaseAdministratorsResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseAdministratorsResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseAdministratorsResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponse) SetStatusCode(v int32) *ListKnowledgeBaseAdministratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponse) SetBody(v *ListKnowledgeBaseAdministratorsResponseBody) *ListKnowledgeBaseAdministratorsResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseAdministratorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
