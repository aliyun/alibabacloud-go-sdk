// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIntegrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIntegrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIntegrationPolicyResponseBody) *UpdateIntegrationPolicyResponse
	GetBody() *UpdateIntegrationPolicyResponseBody
}

type UpdateIntegrationPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntegrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIntegrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIntegrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIntegrationPolicyResponse) GetBody() *UpdateIntegrationPolicyResponseBody {
	return s.Body
}

func (s *UpdateIntegrationPolicyResponse) SetHeaders(v map[string]*string) *UpdateIntegrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntegrationPolicyResponse) SetStatusCode(v int32) *UpdateIntegrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntegrationPolicyResponse) SetBody(v *UpdateIntegrationPolicyResponseBody) *UpdateIntegrationPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateIntegrationPolicyResponse) Validate() error {
	return dara.Validate(s)
}
