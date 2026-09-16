// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeUploadUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveKnowledgeUploadUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveKnowledgeUploadUserResponse
	GetStatusCode() *int32
	SetBody(v *RemoveKnowledgeUploadUserResponseBody) *RemoveKnowledgeUploadUserResponse
	GetBody() *RemoveKnowledgeUploadUserResponseBody
}

type RemoveKnowledgeUploadUserResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveKnowledgeUploadUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveKnowledgeUploadUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeUploadUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeUploadUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveKnowledgeUploadUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveKnowledgeUploadUserResponse) GetBody() *RemoveKnowledgeUploadUserResponseBody {
	return s.Body
}

func (s *RemoveKnowledgeUploadUserResponse) SetHeaders(v map[string]*string) *RemoveKnowledgeUploadUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveKnowledgeUploadUserResponse) SetStatusCode(v int32) *RemoveKnowledgeUploadUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveKnowledgeUploadUserResponse) SetBody(v *RemoveKnowledgeUploadUserResponseBody) *RemoveKnowledgeUploadUserResponse {
	s.Body = v
	return s
}

func (s *RemoveKnowledgeUploadUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
