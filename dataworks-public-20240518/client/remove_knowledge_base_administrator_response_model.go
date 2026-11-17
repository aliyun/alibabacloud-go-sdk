// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeBaseAdministratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveKnowledgeBaseAdministratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveKnowledgeBaseAdministratorResponse
	GetStatusCode() *int32
	SetBody(v *RemoveKnowledgeBaseAdministratorResponseBody) *RemoveKnowledgeBaseAdministratorResponse
	GetBody() *RemoveKnowledgeBaseAdministratorResponseBody
}

type RemoveKnowledgeBaseAdministratorResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveKnowledgeBaseAdministratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveKnowledgeBaseAdministratorResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeBaseAdministratorResponse) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeBaseAdministratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveKnowledgeBaseAdministratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveKnowledgeBaseAdministratorResponse) GetBody() *RemoveKnowledgeBaseAdministratorResponseBody {
	return s.Body
}

func (s *RemoveKnowledgeBaseAdministratorResponse) SetHeaders(v map[string]*string) *RemoveKnowledgeBaseAdministratorResponse {
	s.Headers = v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponse) SetStatusCode(v int32) *RemoveKnowledgeBaseAdministratorResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponse) SetBody(v *RemoveKnowledgeBaseAdministratorResponseBody) *RemoveKnowledgeBaseAdministratorResponse {
	s.Body = v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
