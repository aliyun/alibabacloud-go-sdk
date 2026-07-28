// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutopilotPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutopilotPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutopilotPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutopilotPolicyResponseBody) *UpdateAutopilotPolicyResponse
	GetBody() *UpdateAutopilotPolicyResponseBody
}

type UpdateAutopilotPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutopilotPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutopilotPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutopilotPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutopilotPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutopilotPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutopilotPolicyResponse) GetBody() *UpdateAutopilotPolicyResponseBody {
	return s.Body
}

func (s *UpdateAutopilotPolicyResponse) SetHeaders(v map[string]*string) *UpdateAutopilotPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutopilotPolicyResponse) SetStatusCode(v int32) *UpdateAutopilotPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutopilotPolicyResponse) SetBody(v *UpdateAutopilotPolicyResponseBody) *UpdateAutopilotPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateAutopilotPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
