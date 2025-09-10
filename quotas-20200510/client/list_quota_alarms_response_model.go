// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaAlarmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaAlarmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaAlarmsResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaAlarmsResponseBody) *ListQuotaAlarmsResponse
	GetBody() *ListQuotaAlarmsResponseBody
}

type ListQuotaAlarmsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaAlarmsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaAlarmsResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaAlarmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaAlarmsResponse) GetBody() *ListQuotaAlarmsResponseBody {
	return s.Body
}

func (s *ListQuotaAlarmsResponse) SetHeaders(v map[string]*string) *ListQuotaAlarmsResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaAlarmsResponse) SetStatusCode(v int32) *ListQuotaAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaAlarmsResponse) SetBody(v *ListQuotaAlarmsResponseBody) *ListQuotaAlarmsResponse {
	s.Body = v
	return s
}

func (s *ListQuotaAlarmsResponse) Validate() error {
	return dara.Validate(s)
}
