// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcMPUEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRtcMPUEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRtcMPUEventSubResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRtcMPUEventSubResponseBody) *UpdateRtcMPUEventSubResponse
	GetBody() *UpdateRtcMPUEventSubResponseBody
}

type UpdateRtcMPUEventSubResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRtcMPUEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRtcMPUEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcMPUEventSubResponse) GoString() string {
	return s.String()
}

func (s *UpdateRtcMPUEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRtcMPUEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRtcMPUEventSubResponse) GetBody() *UpdateRtcMPUEventSubResponseBody {
	return s.Body
}

func (s *UpdateRtcMPUEventSubResponse) SetHeaders(v map[string]*string) *UpdateRtcMPUEventSubResponse {
	s.Headers = v
	return s
}

func (s *UpdateRtcMPUEventSubResponse) SetStatusCode(v int32) *UpdateRtcMPUEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRtcMPUEventSubResponse) SetBody(v *UpdateRtcMPUEventSubResponseBody) *UpdateRtcMPUEventSubResponse {
	s.Body = v
	return s
}

func (s *UpdateRtcMPUEventSubResponse) Validate() error {
	return dara.Validate(s)
}
