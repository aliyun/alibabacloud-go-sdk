// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtcMPUEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRtcMPUEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRtcMPUEventSubResponse
	GetStatusCode() *int32
	SetBody(v *CreateRtcMPUEventSubResponseBody) *CreateRtcMPUEventSubResponse
	GetBody() *CreateRtcMPUEventSubResponseBody
}

type CreateRtcMPUEventSubResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRtcMPUEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRtcMPUEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRtcMPUEventSubResponse) GoString() string {
	return s.String()
}

func (s *CreateRtcMPUEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRtcMPUEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRtcMPUEventSubResponse) GetBody() *CreateRtcMPUEventSubResponseBody {
	return s.Body
}

func (s *CreateRtcMPUEventSubResponse) SetHeaders(v map[string]*string) *CreateRtcMPUEventSubResponse {
	s.Headers = v
	return s
}

func (s *CreateRtcMPUEventSubResponse) SetStatusCode(v int32) *CreateRtcMPUEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRtcMPUEventSubResponse) SetBody(v *CreateRtcMPUEventSubResponseBody) *CreateRtcMPUEventSubResponse {
	s.Body = v
	return s
}

func (s *CreateRtcMPUEventSubResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
