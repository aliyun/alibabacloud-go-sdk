// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeBaseAdministratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddKnowledgeBaseAdministratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddKnowledgeBaseAdministratorResponse
	GetStatusCode() *int32
	SetBody(v *AddKnowledgeBaseAdministratorResponseBody) *AddKnowledgeBaseAdministratorResponse
	GetBody() *AddKnowledgeBaseAdministratorResponseBody
}

type AddKnowledgeBaseAdministratorResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKnowledgeBaseAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKnowledgeBaseAdministratorResponse) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeBaseAdministratorResponse) GoString() string {
	return s.String()
}

func (s *AddKnowledgeBaseAdministratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddKnowledgeBaseAdministratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddKnowledgeBaseAdministratorResponse) GetBody() *AddKnowledgeBaseAdministratorResponseBody {
	return s.Body
}

func (s *AddKnowledgeBaseAdministratorResponse) SetHeaders(v map[string]*string) *AddKnowledgeBaseAdministratorResponse {
	s.Headers = v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponse) SetStatusCode(v int32) *AddKnowledgeBaseAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponse) SetBody(v *AddKnowledgeBaseAdministratorResponseBody) *AddKnowledgeBaseAdministratorResponse {
	s.Body = v
	return s
}

func (s *AddKnowledgeBaseAdministratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
