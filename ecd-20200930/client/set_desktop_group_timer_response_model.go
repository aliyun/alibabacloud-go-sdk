// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDesktopGroupTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDesktopGroupTimerResponse
	GetStatusCode() *int32
	SetBody(v *SetDesktopGroupTimerResponseBody) *SetDesktopGroupTimerResponse
	GetBody() *SetDesktopGroupTimerResponseBody
}

type SetDesktopGroupTimerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDesktopGroupTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDesktopGroupTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupTimerResponse) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDesktopGroupTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDesktopGroupTimerResponse) GetBody() *SetDesktopGroupTimerResponseBody {
	return s.Body
}

func (s *SetDesktopGroupTimerResponse) SetHeaders(v map[string]*string) *SetDesktopGroupTimerResponse {
	s.Headers = v
	return s
}

func (s *SetDesktopGroupTimerResponse) SetStatusCode(v int32) *SetDesktopGroupTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDesktopGroupTimerResponse) SetBody(v *SetDesktopGroupTimerResponseBody) *SetDesktopGroupTimerResponse {
	s.Body = v
	return s
}

func (s *SetDesktopGroupTimerResponse) Validate() error {
	return dara.Validate(s)
}
