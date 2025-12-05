// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretRotationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecretRotationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecretRotationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecretRotationPolicyResponseBody) *UpdateSecretRotationPolicyResponse
	GetBody() *UpdateSecretRotationPolicyResponseBody
}

type UpdateSecretRotationPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretRotationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretRotationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRotationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecretRotationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecretRotationPolicyResponse) GetBody() *UpdateSecretRotationPolicyResponseBody {
	return s.Body
}

func (s *UpdateSecretRotationPolicyResponse) SetHeaders(v map[string]*string) *UpdateSecretRotationPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretRotationPolicyResponse) SetStatusCode(v int32) *UpdateSecretRotationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretRotationPolicyResponse) SetBody(v *UpdateSecretRotationPolicyResponseBody) *UpdateSecretRotationPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateSecretRotationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
