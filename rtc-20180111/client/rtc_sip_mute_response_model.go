// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSipMuteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RtcSipMuteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RtcSipMuteResponse
	GetStatusCode() *int32
	SetBody(v *RtcSipMuteResponseBody) *RtcSipMuteResponse
	GetBody() *RtcSipMuteResponseBody
}

type RtcSipMuteResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RtcSipMuteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RtcSipMuteResponse) String() string {
	return dara.Prettify(s)
}

func (s RtcSipMuteResponse) GoString() string {
	return s.String()
}

func (s *RtcSipMuteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RtcSipMuteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RtcSipMuteResponse) GetBody() *RtcSipMuteResponseBody {
	return s.Body
}

func (s *RtcSipMuteResponse) SetHeaders(v map[string]*string) *RtcSipMuteResponse {
	s.Headers = v
	return s
}

func (s *RtcSipMuteResponse) SetStatusCode(v int32) *RtcSipMuteResponse {
	s.StatusCode = &v
	return s
}

func (s *RtcSipMuteResponse) SetBody(v *RtcSipMuteResponseBody) *RtcSipMuteResponse {
	s.Body = v
	return s
}

func (s *RtcSipMuteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
