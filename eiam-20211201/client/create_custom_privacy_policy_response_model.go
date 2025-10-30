// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomPrivacyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomPrivacyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomPrivacyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomPrivacyPolicyResponseBody) *CreateCustomPrivacyPolicyResponse
	GetBody() *CreateCustomPrivacyPolicyResponseBody
}

type CreateCustomPrivacyPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomPrivacyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomPrivacyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomPrivacyPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomPrivacyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomPrivacyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomPrivacyPolicyResponse) GetBody() *CreateCustomPrivacyPolicyResponseBody {
	return s.Body
}

func (s *CreateCustomPrivacyPolicyResponse) SetHeaders(v map[string]*string) *CreateCustomPrivacyPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomPrivacyPolicyResponse) SetStatusCode(v int32) *CreateCustomPrivacyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomPrivacyPolicyResponse) SetBody(v *CreateCustomPrivacyPolicyResponseBody) *CreateCustomPrivacyPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateCustomPrivacyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
