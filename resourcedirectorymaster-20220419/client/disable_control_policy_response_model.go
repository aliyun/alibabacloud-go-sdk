// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DisableControlPolicyResponseBody) *DisableControlPolicyResponse
	GetBody() *DisableControlPolicyResponseBody
}

type DisableControlPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableControlPolicyResponse) GetBody() *DisableControlPolicyResponseBody {
	return s.Body
}

func (s *DisableControlPolicyResponse) SetHeaders(v map[string]*string) *DisableControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableControlPolicyResponse) SetStatusCode(v int32) *DisableControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableControlPolicyResponse) SetBody(v *DisableControlPolicyResponseBody) *DisableControlPolicyResponse {
	s.Body = v
	return s
}

func (s *DisableControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
