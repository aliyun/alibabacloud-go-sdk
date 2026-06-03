// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemAlarmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemAlarmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemAlarmsResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemAlarmsResponseBody) *ListSystemAlarmsResponse
	GetBody() *ListSystemAlarmsResponseBody
}

type ListSystemAlarmsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemAlarmsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAlarmsResponse) GoString() string {
	return s.String()
}

func (s *ListSystemAlarmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemAlarmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemAlarmsResponse) GetBody() *ListSystemAlarmsResponseBody {
	return s.Body
}

func (s *ListSystemAlarmsResponse) SetHeaders(v map[string]*string) *ListSystemAlarmsResponse {
	s.Headers = v
	return s
}

func (s *ListSystemAlarmsResponse) SetStatusCode(v int32) *ListSystemAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemAlarmsResponse) SetBody(v *ListSystemAlarmsResponseBody) *ListSystemAlarmsResponse {
	s.Body = v
	return s
}

func (s *ListSystemAlarmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
