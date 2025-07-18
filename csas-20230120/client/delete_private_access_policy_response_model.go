// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateAccessPolicyResponseBody) *DeletePrivateAccessPolicyResponse
	GetBody() *DeletePrivateAccessPolicyResponseBody
}

type DeletePrivateAccessPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateAccessPolicyResponse) GetBody() *DeletePrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *DeletePrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *DeletePrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateAccessPolicyResponse) SetStatusCode(v int32) *DeletePrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateAccessPolicyResponse) SetBody(v *DeletePrivateAccessPolicyResponseBody) *DeletePrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
