// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachFromPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachFromPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachFromPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DetachFromPolicyResponseBody) *DetachFromPolicyResponse
	GetBody() *DetachFromPolicyResponseBody
}

type DetachFromPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachFromPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachFromPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachFromPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachFromPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachFromPolicyResponse) GetBody() *DetachFromPolicyResponseBody {
	return s.Body
}

func (s *DetachFromPolicyResponse) SetHeaders(v map[string]*string) *DetachFromPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachFromPolicyResponse) SetStatusCode(v int32) *DetachFromPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachFromPolicyResponse) SetBody(v *DetachFromPolicyResponseBody) *DetachFromPolicyResponse {
	s.Body = v
	return s
}

func (s *DetachFromPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
