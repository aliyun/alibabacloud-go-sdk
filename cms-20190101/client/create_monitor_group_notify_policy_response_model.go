// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupNotifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorGroupNotifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorGroupNotifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitorGroupNotifyPolicyResponseBody) *CreateMonitorGroupNotifyPolicyResponse
	GetBody() *CreateMonitorGroupNotifyPolicyResponseBody
}

type CreateMonitorGroupNotifyPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorGroupNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorGroupNotifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupNotifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorGroupNotifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorGroupNotifyPolicyResponse) GetBody() *CreateMonitorGroupNotifyPolicyResponseBody {
	return s.Body
}

func (s *CreateMonitorGroupNotifyPolicyResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponse) SetStatusCode(v int32) *CreateMonitorGroupNotifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponse) SetBody(v *CreateMonitorGroupNotifyPolicyResponseBody) *CreateMonitorGroupNotifyPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
