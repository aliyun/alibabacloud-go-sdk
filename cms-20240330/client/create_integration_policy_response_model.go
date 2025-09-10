// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntegrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIntegrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIntegrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateIntegrationPolicyResponseBody) *CreateIntegrationPolicyResponse
	GetBody() *CreateIntegrationPolicyResponseBody
}

type CreateIntegrationPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntegrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntegrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateIntegrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIntegrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIntegrationPolicyResponse) GetBody() *CreateIntegrationPolicyResponseBody {
	return s.Body
}

func (s *CreateIntegrationPolicyResponse) SetHeaders(v map[string]*string) *CreateIntegrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateIntegrationPolicyResponse) SetStatusCode(v int32) *CreateIntegrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntegrationPolicyResponse) SetBody(v *CreateIntegrationPolicyResponseBody) *CreateIntegrationPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateIntegrationPolicyResponse) Validate() error {
	return dara.Validate(s)
}
