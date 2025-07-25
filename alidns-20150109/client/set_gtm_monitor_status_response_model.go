// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGtmMonitorStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetGtmMonitorStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetGtmMonitorStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetGtmMonitorStatusResponseBody) *SetGtmMonitorStatusResponse
	GetBody() *SetGtmMonitorStatusResponseBody
}

type SetGtmMonitorStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetGtmMonitorStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetGtmMonitorStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetGtmMonitorStatusResponse) GoString() string {
	return s.String()
}

func (s *SetGtmMonitorStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetGtmMonitorStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetGtmMonitorStatusResponse) GetBody() *SetGtmMonitorStatusResponseBody {
	return s.Body
}

func (s *SetGtmMonitorStatusResponse) SetHeaders(v map[string]*string) *SetGtmMonitorStatusResponse {
	s.Headers = v
	return s
}

func (s *SetGtmMonitorStatusResponse) SetStatusCode(v int32) *SetGtmMonitorStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGtmMonitorStatusResponse) SetBody(v *SetGtmMonitorStatusResponseBody) *SetGtmMonitorStatusResponse {
	s.Body = v
	return s
}

func (s *SetGtmMonitorStatusResponse) Validate() error {
	return dara.Validate(s)
}
