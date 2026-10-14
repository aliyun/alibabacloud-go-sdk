// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAgentConfigSupportedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAgentConfigSupportedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAgentConfigSupportedResponse
	GetStatusCode() *int32
	SetBody(v *CheckAgentConfigSupportedResponseBody) *CheckAgentConfigSupportedResponse
	GetBody() *CheckAgentConfigSupportedResponseBody
}

type CheckAgentConfigSupportedResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAgentConfigSupportedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAgentConfigSupportedResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAgentConfigSupportedResponse) GoString() string {
	return s.String()
}

func (s *CheckAgentConfigSupportedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAgentConfigSupportedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAgentConfigSupportedResponse) GetBody() *CheckAgentConfigSupportedResponseBody {
	return s.Body
}

func (s *CheckAgentConfigSupportedResponse) SetHeaders(v map[string]*string) *CheckAgentConfigSupportedResponse {
	s.Headers = v
	return s
}

func (s *CheckAgentConfigSupportedResponse) SetStatusCode(v int32) *CheckAgentConfigSupportedResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAgentConfigSupportedResponse) SetBody(v *CheckAgentConfigSupportedResponseBody) *CheckAgentConfigSupportedResponse {
	s.Body = v
	return s
}

func (s *CheckAgentConfigSupportedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
