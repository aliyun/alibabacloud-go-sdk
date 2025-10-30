// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistrationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRegistrationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRegistrationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateRegistrationPolicyResponseBody) *CreateRegistrationPolicyResponse
	GetBody() *CreateRegistrationPolicyResponseBody
}

type CreateRegistrationPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRegistrationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRegistrationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRegistrationPolicyResponse) GetBody() *CreateRegistrationPolicyResponseBody {
	return s.Body
}

func (s *CreateRegistrationPolicyResponse) SetHeaders(v map[string]*string) *CreateRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateRegistrationPolicyResponse) SetStatusCode(v int32) *CreateRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRegistrationPolicyResponse) SetBody(v *CreateRegistrationPolicyResponseBody) *CreateRegistrationPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateRegistrationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
