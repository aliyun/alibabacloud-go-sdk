// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstancePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstancePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstancePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstancePolicyResponseBody) *CheckInstancePolicyResponse
	GetBody() *CheckInstancePolicyResponseBody
}

type CheckInstancePolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstancePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *CheckInstancePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstancePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstancePolicyResponse) GetBody() *CheckInstancePolicyResponseBody {
	return s.Body
}

func (s *CheckInstancePolicyResponse) SetHeaders(v map[string]*string) *CheckInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *CheckInstancePolicyResponse) SetStatusCode(v int32) *CheckInstancePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstancePolicyResponse) SetBody(v *CheckInstancePolicyResponseBody) *CheckInstancePolicyResponse {
	s.Body = v
	return s
}

func (s *CheckInstancePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
