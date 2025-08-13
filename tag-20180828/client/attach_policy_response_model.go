// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AttachPolicyResponseBody) *AttachPolicyResponse
	GetBody() *AttachPolicyResponseBody
}

type AttachPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachPolicyResponse) GetBody() *AttachPolicyResponseBody {
	return s.Body
}

func (s *AttachPolicyResponse) SetHeaders(v map[string]*string) *AttachPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyResponse) SetStatusCode(v int32) *AttachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyResponse) SetBody(v *AttachPolicyResponseBody) *AttachPolicyResponse {
	s.Body = v
	return s
}

func (s *AttachPolicyResponse) Validate() error {
	return dara.Validate(s)
}
