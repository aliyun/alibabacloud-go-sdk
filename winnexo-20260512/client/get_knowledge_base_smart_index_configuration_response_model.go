// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseSmartIndexConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseSmartIndexConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseSmartIndexConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBaseSmartIndexConfigurationResponseBody) *GetKnowledgeBaseSmartIndexConfigurationResponse
	GetBody() *GetKnowledgeBaseSmartIndexConfigurationResponseBody
}

type GetKnowledgeBaseSmartIndexConfigurationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBaseSmartIndexConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseSmartIndexConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponse) GetBody() *GetKnowledgeBaseSmartIndexConfigurationResponseBody {
	return s.Body
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseSmartIndexConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponse) SetStatusCode(v int32) *GetKnowledgeBaseSmartIndexConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponse) SetBody(v *GetKnowledgeBaseSmartIndexConfigurationResponseBody) *GetKnowledgeBaseSmartIndexConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseSmartIndexConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
