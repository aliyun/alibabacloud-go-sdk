// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeUploadUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddKnowledgeUploadUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddKnowledgeUploadUserResponse
	GetStatusCode() *int32
	SetBody(v *AddKnowledgeUploadUserResponseBody) *AddKnowledgeUploadUserResponse
	GetBody() *AddKnowledgeUploadUserResponseBody
}

type AddKnowledgeUploadUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKnowledgeUploadUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKnowledgeUploadUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeUploadUserResponse) GoString() string {
	return s.String()
}

func (s *AddKnowledgeUploadUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddKnowledgeUploadUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddKnowledgeUploadUserResponse) GetBody() *AddKnowledgeUploadUserResponseBody {
	return s.Body
}

func (s *AddKnowledgeUploadUserResponse) SetHeaders(v map[string]*string) *AddKnowledgeUploadUserResponse {
	s.Headers = v
	return s
}

func (s *AddKnowledgeUploadUserResponse) SetStatusCode(v int32) *AddKnowledgeUploadUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKnowledgeUploadUserResponse) SetBody(v *AddKnowledgeUploadUserResponseBody) *AddKnowledgeUploadUserResponse {
	s.Body = v
	return s
}

func (s *AddKnowledgeUploadUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
