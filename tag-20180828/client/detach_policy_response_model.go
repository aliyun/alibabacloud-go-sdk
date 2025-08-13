// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DetachPolicyResponseBody) *DetachPolicyResponse
	GetBody() *DetachPolicyResponseBody
}

type DetachPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachPolicyResponse) GetBody() *DetachPolicyResponseBody {
	return s.Body
}

func (s *DetachPolicyResponse) SetHeaders(v map[string]*string) *DetachPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyResponse) SetStatusCode(v int32) *DetachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyResponse) SetBody(v *DetachPolicyResponseBody) *DetachPolicyResponse {
	s.Body = v
	return s
}

func (s *DetachPolicyResponse) Validate() error {
	return dara.Validate(s)
}
