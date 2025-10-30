// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomPrivacyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomPrivacyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomPrivacyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomPrivacyPolicyResponseBody) *DeleteCustomPrivacyPolicyResponse
	GetBody() *DeleteCustomPrivacyPolicyResponseBody
}

type DeleteCustomPrivacyPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomPrivacyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomPrivacyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomPrivacyPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomPrivacyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomPrivacyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomPrivacyPolicyResponse) GetBody() *DeleteCustomPrivacyPolicyResponseBody {
	return s.Body
}

func (s *DeleteCustomPrivacyPolicyResponse) SetHeaders(v map[string]*string) *DeleteCustomPrivacyPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomPrivacyPolicyResponse) SetStatusCode(v int32) *DeleteCustomPrivacyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomPrivacyPolicyResponse) SetBody(v *DeleteCustomPrivacyPolicyResponseBody) *DeleteCustomPrivacyPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomPrivacyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
