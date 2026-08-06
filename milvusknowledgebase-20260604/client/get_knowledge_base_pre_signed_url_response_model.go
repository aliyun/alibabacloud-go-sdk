// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBasePreSignedUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBasePreSignedUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBasePreSignedUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBasePreSignedUrlResponseBody) *GetKnowledgeBasePreSignedUrlResponse
	GetBody() *GetKnowledgeBasePreSignedUrlResponseBody
}

type GetKnowledgeBasePreSignedUrlResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBasePreSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBasePreSignedUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBasePreSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBasePreSignedUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBasePreSignedUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBasePreSignedUrlResponse) GetBody() *GetKnowledgeBasePreSignedUrlResponseBody {
	return s.Body
}

func (s *GetKnowledgeBasePreSignedUrlResponse) SetHeaders(v map[string]*string) *GetKnowledgeBasePreSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponse) SetStatusCode(v int32) *GetKnowledgeBasePreSignedUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponse) SetBody(v *GetKnowledgeBasePreSignedUrlResponseBody) *GetKnowledgeBasePreSignedUrlResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
