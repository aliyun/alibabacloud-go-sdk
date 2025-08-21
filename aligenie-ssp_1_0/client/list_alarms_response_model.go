// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlarmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlarmsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlarmsResponseBody) *ListAlarmsResponse
	GetBody() *ListAlarmsResponseBody
}

type ListAlarmsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlarmsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlarmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlarmsResponse) GetBody() *ListAlarmsResponseBody {
	return s.Body
}

func (s *ListAlarmsResponse) SetHeaders(v map[string]*string) *ListAlarmsResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmsResponse) SetStatusCode(v int32) *ListAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlarmsResponse) SetBody(v *ListAlarmsResponseBody) *ListAlarmsResponse {
	s.Body = v
	return s
}

func (s *ListAlarmsResponse) Validate() error {
	return dara.Validate(s)
}
