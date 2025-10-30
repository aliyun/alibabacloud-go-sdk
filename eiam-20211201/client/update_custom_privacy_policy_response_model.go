// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomPrivacyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomPrivacyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomPrivacyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomPrivacyPolicyResponseBody) *UpdateCustomPrivacyPolicyResponse
	GetBody() *UpdateCustomPrivacyPolicyResponseBody
}

type UpdateCustomPrivacyPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomPrivacyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomPrivacyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomPrivacyPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomPrivacyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomPrivacyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomPrivacyPolicyResponse) GetBody() *UpdateCustomPrivacyPolicyResponseBody {
	return s.Body
}

func (s *UpdateCustomPrivacyPolicyResponse) SetHeaders(v map[string]*string) *UpdateCustomPrivacyPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomPrivacyPolicyResponse) SetStatusCode(v int32) *UpdateCustomPrivacyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomPrivacyPolicyResponse) SetBody(v *UpdateCustomPrivacyPolicyResponseBody) *UpdateCustomPrivacyPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomPrivacyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
