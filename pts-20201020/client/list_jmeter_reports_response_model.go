// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJMeterReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJMeterReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJMeterReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListJMeterReportsResponseBody) *ListJMeterReportsResponse
	GetBody() *ListJMeterReportsResponseBody
}

type ListJMeterReportsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJMeterReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJMeterReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJMeterReportsResponse) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJMeterReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJMeterReportsResponse) GetBody() *ListJMeterReportsResponseBody {
	return s.Body
}

func (s *ListJMeterReportsResponse) SetHeaders(v map[string]*string) *ListJMeterReportsResponse {
	s.Headers = v
	return s
}

func (s *ListJMeterReportsResponse) SetStatusCode(v int32) *ListJMeterReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJMeterReportsResponse) SetBody(v *ListJMeterReportsResponseBody) *ListJMeterReportsResponse {
	s.Body = v
	return s
}

func (s *ListJMeterReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
