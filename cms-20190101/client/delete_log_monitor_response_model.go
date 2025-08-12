// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogMonitorResponseBody) *DeleteLogMonitorResponse
	GetBody() *DeleteLogMonitorResponseBody
}

type DeleteLogMonitorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogMonitorResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogMonitorResponse) GetBody() *DeleteLogMonitorResponseBody {
	return s.Body
}

func (s *DeleteLogMonitorResponse) SetHeaders(v map[string]*string) *DeleteLogMonitorResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogMonitorResponse) SetStatusCode(v int32) *DeleteLogMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogMonitorResponse) SetBody(v *DeleteLogMonitorResponseBody) *DeleteLogMonitorResponse {
	s.Body = v
	return s
}

func (s *DeleteLogMonitorResponse) Validate() error {
	return dara.Validate(s)
}
