// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunServiceScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunServiceScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunServiceScheduleResponse
	GetStatusCode() *int32
	SetBody(v *RunServiceScheduleResponseBody) *RunServiceScheduleResponse
	GetBody() *RunServiceScheduleResponseBody
}

type RunServiceScheduleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunServiceScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunServiceScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s RunServiceScheduleResponse) GoString() string {
	return s.String()
}

func (s *RunServiceScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunServiceScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunServiceScheduleResponse) GetBody() *RunServiceScheduleResponseBody {
	return s.Body
}

func (s *RunServiceScheduleResponse) SetHeaders(v map[string]*string) *RunServiceScheduleResponse {
	s.Headers = v
	return s
}

func (s *RunServiceScheduleResponse) SetStatusCode(v int32) *RunServiceScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *RunServiceScheduleResponse) SetBody(v *RunServiceScheduleResponseBody) *RunServiceScheduleResponse {
	s.Body = v
	return s
}

func (s *RunServiceScheduleResponse) Validate() error {
	return dara.Validate(s)
}
