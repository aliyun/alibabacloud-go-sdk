// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupNotifyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorGroupNotifyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorGroupNotifyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMonitorGroupNotifyPolicyResponseBody) *DeleteMonitorGroupNotifyPolicyResponse
	GetBody() *DeleteMonitorGroupNotifyPolicyResponseBody
}

type DeleteMonitorGroupNotifyPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMonitorGroupNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMonitorGroupNotifyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) GetBody() *DeleteMonitorGroupNotifyPolicyResponseBody {
	return s.Body
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) SetStatusCode(v int32) *DeleteMonitorGroupNotifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) SetBody(v *DeleteMonitorGroupNotifyPolicyResponseBody) *DeleteMonitorGroupNotifyPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) Validate() error {
	return dara.Validate(s)
}
