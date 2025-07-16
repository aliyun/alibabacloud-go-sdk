// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduleResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduleResponseBody) *GetScheduleResponse
	GetBody() *GetScheduleResponseBody
}

type GetScheduleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduleResponse) GetBody() *GetScheduleResponseBody {
	return s.Body
}

func (s *GetScheduleResponse) SetHeaders(v map[string]*string) *GetScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetScheduleResponse) SetStatusCode(v int32) *GetScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduleResponse) SetBody(v *GetScheduleResponseBody) *GetScheduleResponse {
	s.Body = v
	return s
}

func (s *GetScheduleResponse) Validate() error {
	return dara.Validate(s)
}
