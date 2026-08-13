// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBaseSourceResponseBody) *GetKnowledgeBaseSourceResponse
	GetBody() *GetKnowledgeBaseSourceResponseBody
}

type GetKnowledgeBaseSourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBaseSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSourceResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseSourceResponse) GetBody() *GetKnowledgeBaseSourceResponseBody {
	return s.Body
}

func (s *GetKnowledgeBaseSourceResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseSourceResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseSourceResponse) SetStatusCode(v int32) *GetKnowledgeBaseSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseSourceResponse) SetBody(v *GetKnowledgeBaseSourceResponseBody) *GetKnowledgeBaseSourceResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
