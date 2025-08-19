// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DetachControlPolicyResponseBody) *DetachControlPolicyResponse
	GetBody() *DetachControlPolicyResponseBody
}

type DetachControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachControlPolicyResponse) GetBody() *DetachControlPolicyResponseBody {
	return s.Body
}

func (s *DetachControlPolicyResponse) SetHeaders(v map[string]*string) *DetachControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachControlPolicyResponse) SetStatusCode(v int32) *DetachControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachControlPolicyResponse) SetBody(v *DetachControlPolicyResponseBody) *DetachControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DetachControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
