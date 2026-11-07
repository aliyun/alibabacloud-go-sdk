// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateLlmTemplateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateLlmTemplateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateLlmTemplateConfigResponse
	GetStatusCode() *int32
	SetBody(v *ValidateLlmTemplateConfigResponseBody) *ValidateLlmTemplateConfigResponse
	GetBody() *ValidateLlmTemplateConfigResponseBody
}

type ValidateLlmTemplateConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateLlmTemplateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateLlmTemplateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateLlmTemplateConfigResponse) GoString() string {
	return s.String()
}

func (s *ValidateLlmTemplateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateLlmTemplateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateLlmTemplateConfigResponse) GetBody() *ValidateLlmTemplateConfigResponseBody {
	return s.Body
}

func (s *ValidateLlmTemplateConfigResponse) SetHeaders(v map[string]*string) *ValidateLlmTemplateConfigResponse {
	s.Headers = v
	return s
}

func (s *ValidateLlmTemplateConfigResponse) SetStatusCode(v int32) *ValidateLlmTemplateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateLlmTemplateConfigResponse) SetBody(v *ValidateLlmTemplateConfigResponseBody) *ValidateLlmTemplateConfigResponse {
	s.Body = v
	return s
}

func (s *ValidateLlmTemplateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
