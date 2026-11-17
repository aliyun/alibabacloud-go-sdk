// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeSourceResponseBody) *GetKnowledgeSourceResponse
	GetBody() *GetKnowledgeSourceResponseBody
}

type GetKnowledgeSourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeSourceResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeSourceResponse) GetBody() *GetKnowledgeSourceResponseBody {
	return s.Body
}

func (s *GetKnowledgeSourceResponse) SetHeaders(v map[string]*string) *GetKnowledgeSourceResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeSourceResponse) SetStatusCode(v int32) *GetKnowledgeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeSourceResponse) SetBody(v *GetKnowledgeSourceResponseBody) *GetKnowledgeSourceResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
