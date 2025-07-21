// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecretPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSecretPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSecretPolicyResponse
	GetStatusCode() *int32
	SetBody(v *SetSecretPolicyResponseBody) *SetSecretPolicyResponse
	GetBody() *SetSecretPolicyResponseBody
}

type SetSecretPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSecretPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSecretPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSecretPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetSecretPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSecretPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSecretPolicyResponse) GetBody() *SetSecretPolicyResponseBody {
	return s.Body
}

func (s *SetSecretPolicyResponse) SetHeaders(v map[string]*string) *SetSecretPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetSecretPolicyResponse) SetStatusCode(v int32) *SetSecretPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSecretPolicyResponse) SetBody(v *SetSecretPolicyResponseBody) *SetSecretPolicyResponse {
	s.Body = v
	return s
}

func (s *SetSecretPolicyResponse) Validate() error {
	return dara.Validate(s)
}
