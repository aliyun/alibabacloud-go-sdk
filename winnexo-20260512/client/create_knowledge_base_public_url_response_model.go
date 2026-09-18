// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBasePublicUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBasePublicUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBasePublicUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBasePublicUrlResponseBody) *CreateKnowledgeBasePublicUrlResponse
	GetBody() *CreateKnowledgeBasePublicUrlResponseBody
}

type CreateKnowledgeBasePublicUrlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBasePublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBasePublicUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBasePublicUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBasePublicUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBasePublicUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBasePublicUrlResponse) GetBody() *CreateKnowledgeBasePublicUrlResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBasePublicUrlResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBasePublicUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponse) SetStatusCode(v int32) *CreateKnowledgeBasePublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponse) SetBody(v *CreateKnowledgeBasePublicUrlResponseBody) *CreateKnowledgeBasePublicUrlResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBasePublicUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
