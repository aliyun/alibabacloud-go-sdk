// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeObjectUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeObjectUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeObjectUploadResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeObjectUploadResponseBody) *CreateKnowledgeObjectUploadResponse
	GetBody() *CreateKnowledgeObjectUploadResponseBody
}

type CreateKnowledgeObjectUploadResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeObjectUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeObjectUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectUploadResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeObjectUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeObjectUploadResponse) GetBody() *CreateKnowledgeObjectUploadResponseBody {
	return s.Body
}

func (s *CreateKnowledgeObjectUploadResponse) SetHeaders(v map[string]*string) *CreateKnowledgeObjectUploadResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeObjectUploadResponse) SetStatusCode(v int32) *CreateKnowledgeObjectUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponse) SetBody(v *CreateKnowledgeObjectUploadResponseBody) *CreateKnowledgeObjectUploadResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeObjectUploadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
