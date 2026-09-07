// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeUploadUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeUploadUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeUploadUserResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeUploadUserResponseBody) *ListKnowledgeUploadUserResponse
	GetBody() *ListKnowledgeUploadUserResponseBody
}

type ListKnowledgeUploadUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeUploadUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeUploadUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeUploadUserResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeUploadUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeUploadUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeUploadUserResponse) GetBody() *ListKnowledgeUploadUserResponseBody {
	return s.Body
}

func (s *ListKnowledgeUploadUserResponse) SetHeaders(v map[string]*string) *ListKnowledgeUploadUserResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeUploadUserResponse) SetStatusCode(v int32) *ListKnowledgeUploadUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeUploadUserResponse) SetBody(v *ListKnowledgeUploadUserResponseBody) *ListKnowledgeUploadUserResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeUploadUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
