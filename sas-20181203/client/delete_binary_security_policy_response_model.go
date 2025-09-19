// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBinarySecurityPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBinarySecurityPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBinarySecurityPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBinarySecurityPolicyResponseBody) *DeleteBinarySecurityPolicyResponse
	GetBody() *DeleteBinarySecurityPolicyResponseBody
}

type DeleteBinarySecurityPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBinarySecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBinarySecurityPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBinarySecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteBinarySecurityPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBinarySecurityPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBinarySecurityPolicyResponse) GetBody() *DeleteBinarySecurityPolicyResponseBody {
	return s.Body
}

func (s *DeleteBinarySecurityPolicyResponse) SetHeaders(v map[string]*string) *DeleteBinarySecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteBinarySecurityPolicyResponse) SetStatusCode(v int32) *DeleteBinarySecurityPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBinarySecurityPolicyResponse) SetBody(v *DeleteBinarySecurityPolicyResponseBody) *DeleteBinarySecurityPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteBinarySecurityPolicyResponse) Validate() error {
	return dara.Validate(s)
}
