// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAlertResponse
	GetStatusCode() *int32
	SetBody(v *StopAlertResponseBody) *StopAlertResponse
	GetBody() *StopAlertResponseBody
}

type StopAlertResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAlertResponse) GoString() string {
	return s.String()
}

func (s *StopAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAlertResponse) GetBody() *StopAlertResponseBody {
	return s.Body
}

func (s *StopAlertResponse) SetHeaders(v map[string]*string) *StopAlertResponse {
	s.Headers = v
	return s
}

func (s *StopAlertResponse) SetStatusCode(v int32) *StopAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAlertResponse) SetBody(v *StopAlertResponseBody) *StopAlertResponse {
	s.Body = v
	return s
}

func (s *StopAlertResponse) Validate() error {
	return dara.Validate(s)
}
