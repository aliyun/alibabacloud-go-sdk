// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAdHocTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAdHocTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAdHocTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopAdHocTaskResponseBody) *StopAdHocTaskResponse
	GetBody() *StopAdHocTaskResponseBody
}

type StopAdHocTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAdHocTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAdHocTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAdHocTaskResponse) GoString() string {
	return s.String()
}

func (s *StopAdHocTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAdHocTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAdHocTaskResponse) GetBody() *StopAdHocTaskResponseBody {
	return s.Body
}

func (s *StopAdHocTaskResponse) SetHeaders(v map[string]*string) *StopAdHocTaskResponse {
	s.Headers = v
	return s
}

func (s *StopAdHocTaskResponse) SetStatusCode(v int32) *StopAdHocTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAdHocTaskResponse) SetBody(v *StopAdHocTaskResponseBody) *StopAdHocTaskResponse {
	s.Body = v
	return s
}

func (s *StopAdHocTaskResponse) Validate() error {
	return dara.Validate(s)
}
