// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduleTaskResponseBody) *GetScheduleTaskResponse
	GetBody() *GetScheduleTaskResponseBody
}

type GetScheduleTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduleTaskResponse) GetBody() *GetScheduleTaskResponseBody {
	return s.Body
}

func (s *GetScheduleTaskResponse) SetHeaders(v map[string]*string) *GetScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleTaskResponse) SetStatusCode(v int32) *GetScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduleTaskResponse) SetBody(v *GetScheduleTaskResponseBody) *GetScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *GetScheduleTaskResponse) Validate() error {
	return dara.Validate(s)
}
