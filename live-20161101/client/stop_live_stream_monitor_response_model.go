// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveStreamMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLiveStreamMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLiveStreamMonitorResponse
	GetStatusCode() *int32
	SetBody(v *StopLiveStreamMonitorResponseBody) *StopLiveStreamMonitorResponse
	GetBody() *StopLiveStreamMonitorResponseBody
}

type StopLiveStreamMonitorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLiveStreamMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLiveStreamMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLiveStreamMonitorResponse) GoString() string {
	return s.String()
}

func (s *StopLiveStreamMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLiveStreamMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLiveStreamMonitorResponse) GetBody() *StopLiveStreamMonitorResponseBody {
	return s.Body
}

func (s *StopLiveStreamMonitorResponse) SetHeaders(v map[string]*string) *StopLiveStreamMonitorResponse {
	s.Headers = v
	return s
}

func (s *StopLiveStreamMonitorResponse) SetStatusCode(v int32) *StopLiveStreamMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLiveStreamMonitorResponse) SetBody(v *StopLiveStreamMonitorResponseBody) *StopLiveStreamMonitorResponse {
	s.Body = v
	return s
}

func (s *StopLiveStreamMonitorResponse) Validate() error {
	return dara.Validate(s)
}
