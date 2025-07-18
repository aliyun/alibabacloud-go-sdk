// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRegistrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRegistrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRegistrationPolicyResponseBody) *UpdateRegistrationPolicyResponse
	GetBody() *UpdateRegistrationPolicyResponseBody
}

type UpdateRegistrationPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRegistrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRegistrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRegistrationPolicyResponse) GetBody() *UpdateRegistrationPolicyResponseBody {
	return s.Body
}

func (s *UpdateRegistrationPolicyResponse) SetHeaders(v map[string]*string) *UpdateRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateRegistrationPolicyResponse) SetStatusCode(v int32) *UpdateRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRegistrationPolicyResponse) SetBody(v *UpdateRegistrationPolicyResponseBody) *UpdateRegistrationPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateRegistrationPolicyResponse) Validate() error {
	return dara.Validate(s)
}
