// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNotifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNotifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNotifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNotifyPolicyResponseBody) *UpdateNotifyPolicyResponse
	GetBody() *UpdateNotifyPolicyResponseBody
}

type UpdateNotifyPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNotifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateNotifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNotifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNotifyPolicyResponse) GetBody() *UpdateNotifyPolicyResponseBody {
	return s.Body
}

func (s *UpdateNotifyPolicyResponse) SetHeaders(v map[string]*string) *UpdateNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateNotifyPolicyResponse) SetStatusCode(v int32) *UpdateNotifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNotifyPolicyResponse) SetBody(v *UpdateNotifyPolicyResponseBody) *UpdateNotifyPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateNotifyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
