// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCallInConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCallInConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCallInConfigResponse
	GetStatusCode() *int32
	SetBody(v *StopCallInConfigResponseBody) *StopCallInConfigResponse
	GetBody() *StopCallInConfigResponseBody
}

type StopCallInConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCallInConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCallInConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCallInConfigResponse) GoString() string {
	return s.String()
}

func (s *StopCallInConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCallInConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCallInConfigResponse) GetBody() *StopCallInConfigResponseBody {
	return s.Body
}

func (s *StopCallInConfigResponse) SetHeaders(v map[string]*string) *StopCallInConfigResponse {
	s.Headers = v
	return s
}

func (s *StopCallInConfigResponse) SetStatusCode(v int32) *StopCallInConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCallInConfigResponse) SetBody(v *StopCallInConfigResponseBody) *StopCallInConfigResponse {
	s.Body = v
	return s
}

func (s *StopCallInConfigResponse) Validate() error {
	return dara.Validate(s)
}
