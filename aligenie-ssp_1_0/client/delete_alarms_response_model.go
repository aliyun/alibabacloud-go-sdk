// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlarmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlarmsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlarmsResponseBody) *DeleteAlarmsResponse
	GetBody() *DeleteAlarmsResponseBody
}

type DeleteAlarmsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlarmsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlarmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlarmsResponse) GetBody() *DeleteAlarmsResponseBody {
	return s.Body
}

func (s *DeleteAlarmsResponse) SetHeaders(v map[string]*string) *DeleteAlarmsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlarmsResponse) SetStatusCode(v int32) *DeleteAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlarmsResponse) SetBody(v *DeleteAlarmsResponseBody) *DeleteAlarmsResponse {
	s.Body = v
	return s
}

func (s *DeleteAlarmsResponse) Validate() error {
	return dara.Validate(s)
}
