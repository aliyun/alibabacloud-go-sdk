// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvServiceMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnvServiceMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnvServiceMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnvServiceMonitorResponseBody) *DeleteEnvServiceMonitorResponse
	GetBody() *DeleteEnvServiceMonitorResponseBody
}

type DeleteEnvServiceMonitorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnvServiceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnvServiceMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvServiceMonitorResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvServiceMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnvServiceMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnvServiceMonitorResponse) GetBody() *DeleteEnvServiceMonitorResponseBody {
	return s.Body
}

func (s *DeleteEnvServiceMonitorResponse) SetHeaders(v map[string]*string) *DeleteEnvServiceMonitorResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvServiceMonitorResponse) SetStatusCode(v int32) *DeleteEnvServiceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnvServiceMonitorResponse) SetBody(v *DeleteEnvServiceMonitorResponseBody) *DeleteEnvServiceMonitorResponse {
	s.Body = v
	return s
}

func (s *DeleteEnvServiceMonitorResponse) Validate() error {
	return dara.Validate(s)
}
