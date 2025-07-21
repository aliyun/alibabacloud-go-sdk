// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecretPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecretPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetSecretPolicyResponseBody) *GetSecretPolicyResponse
	GetBody() *GetSecretPolicyResponseBody
}

type GetSecretPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecretPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetSecretPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecretPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecretPolicyResponse) GetBody() *GetSecretPolicyResponseBody {
	return s.Body
}

func (s *GetSecretPolicyResponse) SetHeaders(v map[string]*string) *GetSecretPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetSecretPolicyResponse) SetStatusCode(v int32) *GetSecretPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretPolicyResponse) SetBody(v *GetSecretPolicyResponseBody) *GetSecretPolicyResponse {
	s.Body = v
	return s
}

func (s *GetSecretPolicyResponse) Validate() error {
	return dara.Validate(s)
}
