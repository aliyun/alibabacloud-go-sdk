// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBaseConfigurationResponseBody) *GetKnowledgeBaseConfigurationResponse
	GetBody() *GetKnowledgeBaseConfigurationResponseBody
}

type GetKnowledgeBaseConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBaseConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseConfigurationResponse) GetBody() *GetKnowledgeBaseConfigurationResponseBody {
	return s.Body
}

func (s *GetKnowledgeBaseConfigurationResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponse) SetStatusCode(v int32) *GetKnowledgeBaseConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponse) SetBody(v *GetKnowledgeBaseConfigurationResponseBody) *GetKnowledgeBaseConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
