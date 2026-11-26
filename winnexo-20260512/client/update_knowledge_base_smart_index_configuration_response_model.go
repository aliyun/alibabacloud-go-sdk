// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSmartIndexConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseSmartIndexConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseSmartIndexConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) *UpdateKnowledgeBaseSmartIndexConfigurationResponse
	GetBody() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody
}

type UpdateKnowledgeBaseSmartIndexConfigurationResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponse) GetBody() *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseSmartIndexConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseSmartIndexConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponse) SetBody(v *UpdateKnowledgeBaseSmartIndexConfigurationResponseBody) *UpdateKnowledgeBaseSmartIndexConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
