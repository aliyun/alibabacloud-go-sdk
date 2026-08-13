// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdminKnowledgeBasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAdminKnowledgeBasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAdminKnowledgeBasesResponse
	GetStatusCode() *int32
	SetBody(v *ListAdminKnowledgeBasesResponseBody) *ListAdminKnowledgeBasesResponse
	GetBody() *ListAdminKnowledgeBasesResponseBody
}

type ListAdminKnowledgeBasesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdminKnowledgeBasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdminKnowledgeBasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesResponse) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAdminKnowledgeBasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAdminKnowledgeBasesResponse) GetBody() *ListAdminKnowledgeBasesResponseBody {
	return s.Body
}

func (s *ListAdminKnowledgeBasesResponse) SetHeaders(v map[string]*string) *ListAdminKnowledgeBasesResponse {
	s.Headers = v
	return s
}

func (s *ListAdminKnowledgeBasesResponse) SetStatusCode(v int32) *ListAdminKnowledgeBasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdminKnowledgeBasesResponse) SetBody(v *ListAdminKnowledgeBasesResponseBody) *ListAdminKnowledgeBasesResponse {
	s.Body = v
	return s
}

func (s *ListAdminKnowledgeBasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
