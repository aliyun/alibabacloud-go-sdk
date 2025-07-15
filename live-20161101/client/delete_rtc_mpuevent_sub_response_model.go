// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtcMPUEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRtcMPUEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRtcMPUEventSubResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRtcMPUEventSubResponseBody) *DeleteRtcMPUEventSubResponse
	GetBody() *DeleteRtcMPUEventSubResponseBody
}

type DeleteRtcMPUEventSubResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRtcMPUEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRtcMPUEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtcMPUEventSubResponse) GoString() string {
	return s.String()
}

func (s *DeleteRtcMPUEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRtcMPUEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRtcMPUEventSubResponse) GetBody() *DeleteRtcMPUEventSubResponseBody {
	return s.Body
}

func (s *DeleteRtcMPUEventSubResponse) SetHeaders(v map[string]*string) *DeleteRtcMPUEventSubResponse {
	s.Headers = v
	return s
}

func (s *DeleteRtcMPUEventSubResponse) SetStatusCode(v int32) *DeleteRtcMPUEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRtcMPUEventSubResponse) SetBody(v *DeleteRtcMPUEventSubResponseBody) *DeleteRtcMPUEventSubResponse {
	s.Body = v
	return s
}

func (s *DeleteRtcMPUEventSubResponse) Validate() error {
	return dara.Validate(s)
}
