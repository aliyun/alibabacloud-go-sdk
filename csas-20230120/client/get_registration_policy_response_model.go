// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegistrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegistrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetRegistrationPolicyResponseBody) *GetRegistrationPolicyResponse
	GetBody() *GetRegistrationPolicyResponseBody
}

type GetRegistrationPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegistrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegistrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegistrationPolicyResponse) GetBody() *GetRegistrationPolicyResponseBody {
	return s.Body
}

func (s *GetRegistrationPolicyResponse) SetHeaders(v map[string]*string) *GetRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetRegistrationPolicyResponse) SetStatusCode(v int32) *GetRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegistrationPolicyResponse) SetBody(v *GetRegistrationPolicyResponseBody) *GetRegistrationPolicyResponse {
	s.Body = v
	return s
}

func (s *GetRegistrationPolicyResponse) Validate() error {
	return dara.Validate(s)
}
