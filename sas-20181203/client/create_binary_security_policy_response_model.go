// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBinarySecurityPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBinarySecurityPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBinarySecurityPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateBinarySecurityPolicyResponseBody) *CreateBinarySecurityPolicyResponse
	GetBody() *CreateBinarySecurityPolicyResponseBody
}

type CreateBinarySecurityPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBinarySecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBinarySecurityPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBinarySecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateBinarySecurityPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBinarySecurityPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBinarySecurityPolicyResponse) GetBody() *CreateBinarySecurityPolicyResponseBody {
	return s.Body
}

func (s *CreateBinarySecurityPolicyResponse) SetHeaders(v map[string]*string) *CreateBinarySecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateBinarySecurityPolicyResponse) SetStatusCode(v int32) *CreateBinarySecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBinarySecurityPolicyResponse) SetBody(v *CreateBinarySecurityPolicyResponseBody) *CreateBinarySecurityPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateBinarySecurityPolicyResponse) Validate() error {
	return dara.Validate(s)
}
