// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDesktopTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDesktopTimerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDesktopTimerResponseBody) *ModifyDesktopTimerResponse
	GetBody() *ModifyDesktopTimerResponseBody
}

type ModifyDesktopTimerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDesktopTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDesktopTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopTimerResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDesktopTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDesktopTimerResponse) GetBody() *ModifyDesktopTimerResponseBody {
	return s.Body
}

func (s *ModifyDesktopTimerResponse) SetHeaders(v map[string]*string) *ModifyDesktopTimerResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopTimerResponse) SetStatusCode(v int32) *ModifyDesktopTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDesktopTimerResponse) SetBody(v *ModifyDesktopTimerResponseBody) *ModifyDesktopTimerResponse {
	s.Body = v
	return s
}

func (s *ModifyDesktopTimerResponse) Validate() error {
	return dara.Validate(s)
}
