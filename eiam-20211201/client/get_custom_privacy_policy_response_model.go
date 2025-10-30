// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomPrivacyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomPrivacyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomPrivacyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomPrivacyPolicyResponseBody) *GetCustomPrivacyPolicyResponse
	GetBody() *GetCustomPrivacyPolicyResponseBody
}

type GetCustomPrivacyPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomPrivacyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomPrivacyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomPrivacyPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetCustomPrivacyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomPrivacyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomPrivacyPolicyResponse) GetBody() *GetCustomPrivacyPolicyResponseBody {
	return s.Body
}

func (s *GetCustomPrivacyPolicyResponse) SetHeaders(v map[string]*string) *GetCustomPrivacyPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetCustomPrivacyPolicyResponse) SetStatusCode(v int32) *GetCustomPrivacyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomPrivacyPolicyResponse) SetBody(v *GetCustomPrivacyPolicyResponseBody) *GetCustomPrivacyPolicyResponse {
	s.Body = v
	return s
}

func (s *GetCustomPrivacyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
