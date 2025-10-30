// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomPrivacyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCustomPrivacyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCustomPrivacyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DisableCustomPrivacyPolicyResponseBody) *DisableCustomPrivacyPolicyResponse
	GetBody() *DisableCustomPrivacyPolicyResponseBody
}

type DisableCustomPrivacyPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCustomPrivacyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCustomPrivacyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomPrivacyPolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableCustomPrivacyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCustomPrivacyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCustomPrivacyPolicyResponse) GetBody() *DisableCustomPrivacyPolicyResponseBody {
	return s.Body
}

func (s *DisableCustomPrivacyPolicyResponse) SetHeaders(v map[string]*string) *DisableCustomPrivacyPolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableCustomPrivacyPolicyResponse) SetStatusCode(v int32) *DisableCustomPrivacyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCustomPrivacyPolicyResponse) SetBody(v *DisableCustomPrivacyPolicyResponseBody) *DisableCustomPrivacyPolicyResponse {
	s.Body = v
	return s
}

func (s *DisableCustomPrivacyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
