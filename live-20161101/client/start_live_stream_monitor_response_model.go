// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveStreamMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartLiveStreamMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartLiveStreamMonitorResponse
	GetStatusCode() *int32
	SetBody(v *StartLiveStreamMonitorResponseBody) *StartLiveStreamMonitorResponse
	GetBody() *StartLiveStreamMonitorResponseBody
}

type StartLiveStreamMonitorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLiveStreamMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLiveStreamMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s StartLiveStreamMonitorResponse) GoString() string {
	return s.String()
}

func (s *StartLiveStreamMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartLiveStreamMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartLiveStreamMonitorResponse) GetBody() *StartLiveStreamMonitorResponseBody {
	return s.Body
}

func (s *StartLiveStreamMonitorResponse) SetHeaders(v map[string]*string) *StartLiveStreamMonitorResponse {
	s.Headers = v
	return s
}

func (s *StartLiveStreamMonitorResponse) SetStatusCode(v int32) *StartLiveStreamMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLiveStreamMonitorResponse) SetBody(v *StartLiveStreamMonitorResponseBody) *StartLiveStreamMonitorResponse {
	s.Body = v
	return s
}

func (s *StartLiveStreamMonitorResponse) Validate() error {
	return dara.Validate(s)
}
