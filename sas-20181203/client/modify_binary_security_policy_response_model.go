// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBinarySecurityPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBinarySecurityPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBinarySecurityPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBinarySecurityPolicyResponseBody) *ModifyBinarySecurityPolicyResponse
	GetBody() *ModifyBinarySecurityPolicyResponseBody
}

type ModifyBinarySecurityPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBinarySecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBinarySecurityPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBinarySecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBinarySecurityPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBinarySecurityPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBinarySecurityPolicyResponse) GetBody() *ModifyBinarySecurityPolicyResponseBody {
	return s.Body
}

func (s *ModifyBinarySecurityPolicyResponse) SetHeaders(v map[string]*string) *ModifyBinarySecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBinarySecurityPolicyResponse) SetStatusCode(v int32) *ModifyBinarySecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBinarySecurityPolicyResponse) SetBody(v *ModifyBinarySecurityPolicyResponseBody) *ModifyBinarySecurityPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyBinarySecurityPolicyResponse) Validate() error {
	return dara.Validate(s)
}
