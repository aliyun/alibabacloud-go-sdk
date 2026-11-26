// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKnowledgeBaseVisibilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetKnowledgeBaseVisibilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetKnowledgeBaseVisibilityResponse
	GetStatusCode() *int32
	SetBody(v *SetKnowledgeBaseVisibilityResponseBody) *SetKnowledgeBaseVisibilityResponse
	GetBody() *SetKnowledgeBaseVisibilityResponseBody
}

type SetKnowledgeBaseVisibilityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetKnowledgeBaseVisibilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetKnowledgeBaseVisibilityResponse) String() string {
	return dara.Prettify(s)
}

func (s SetKnowledgeBaseVisibilityResponse) GoString() string {
	return s.String()
}

func (s *SetKnowledgeBaseVisibilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetKnowledgeBaseVisibilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetKnowledgeBaseVisibilityResponse) GetBody() *SetKnowledgeBaseVisibilityResponseBody {
	return s.Body
}

func (s *SetKnowledgeBaseVisibilityResponse) SetHeaders(v map[string]*string) *SetKnowledgeBaseVisibilityResponse {
	s.Headers = v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponse) SetStatusCode(v int32) *SetKnowledgeBaseVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponse) SetBody(v *SetKnowledgeBaseVisibilityResponseBody) *SetKnowledgeBaseVisibilityResponse {
	s.Body = v
	return s
}

func (s *SetKnowledgeBaseVisibilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
