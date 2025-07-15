// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupScaleTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDesktopGroupScaleTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDesktopGroupScaleTimerResponse
	GetStatusCode() *int32
	SetBody(v *SetDesktopGroupScaleTimerResponseBody) *SetDesktopGroupScaleTimerResponse
	GetBody() *SetDesktopGroupScaleTimerResponseBody
}

type SetDesktopGroupScaleTimerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDesktopGroupScaleTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDesktopGroupScaleTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupScaleTimerResponse) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupScaleTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDesktopGroupScaleTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDesktopGroupScaleTimerResponse) GetBody() *SetDesktopGroupScaleTimerResponseBody {
	return s.Body
}

func (s *SetDesktopGroupScaleTimerResponse) SetHeaders(v map[string]*string) *SetDesktopGroupScaleTimerResponse {
	s.Headers = v
	return s
}

func (s *SetDesktopGroupScaleTimerResponse) SetStatusCode(v int32) *SetDesktopGroupScaleTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDesktopGroupScaleTimerResponse) SetBody(v *SetDesktopGroupScaleTimerResponseBody) *SetDesktopGroupScaleTimerResponse {
	s.Body = v
	return s
}

func (s *SetDesktopGroupScaleTimerResponse) Validate() error {
	return dara.Validate(s)
}
