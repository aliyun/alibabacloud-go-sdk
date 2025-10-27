// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScheduleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScheduleResponseBody) *DeleteScheduleResponse
	GetBody() *DeleteScheduleResponseBody
}

type DeleteScheduleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScheduleResponse) GetBody() *DeleteScheduleResponseBody {
	return s.Body
}

func (s *DeleteScheduleResponse) SetHeaders(v map[string]*string) *DeleteScheduleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleResponse) SetStatusCode(v int32) *DeleteScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduleResponse) SetBody(v *DeleteScheduleResponseBody) *DeleteScheduleResponse {
	s.Body = v
	return s
}

func (s *DeleteScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
