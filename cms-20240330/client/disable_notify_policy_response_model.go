// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNotifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableNotifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableNotifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DisableNotifyPolicyResponseBody) *DisableNotifyPolicyResponse
	GetBody() *DisableNotifyPolicyResponseBody
}

type DisableNotifyPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableNotifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableNotifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableNotifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableNotifyPolicyResponse) GetBody() *DisableNotifyPolicyResponseBody {
	return s.Body
}

func (s *DisableNotifyPolicyResponse) SetHeaders(v map[string]*string) *DisableNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableNotifyPolicyResponse) SetStatusCode(v int32) *DisableNotifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableNotifyPolicyResponse) SetBody(v *DisableNotifyPolicyResponseBody) *DisableNotifyPolicyResponse {
	s.Body = v
	return s
}

func (s *DisableNotifyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
