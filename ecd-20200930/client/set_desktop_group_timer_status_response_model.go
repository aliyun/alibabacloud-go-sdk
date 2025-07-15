// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupTimerStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDesktopGroupTimerStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDesktopGroupTimerStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDesktopGroupTimerStatusResponseBody) *SetDesktopGroupTimerStatusResponse
	GetBody() *SetDesktopGroupTimerStatusResponseBody
}

type SetDesktopGroupTimerStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDesktopGroupTimerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDesktopGroupTimerStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupTimerStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupTimerStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDesktopGroupTimerStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDesktopGroupTimerStatusResponse) GetBody() *SetDesktopGroupTimerStatusResponseBody {
	return s.Body
}

func (s *SetDesktopGroupTimerStatusResponse) SetHeaders(v map[string]*string) *SetDesktopGroupTimerStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDesktopGroupTimerStatusResponse) SetStatusCode(v int32) *SetDesktopGroupTimerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDesktopGroupTimerStatusResponse) SetBody(v *SetDesktopGroupTimerStatusResponseBody) *SetDesktopGroupTimerStatusResponse {
	s.Body = v
	return s
}

func (s *SetDesktopGroupTimerStatusResponse) Validate() error {
	return dara.Validate(s)
}
