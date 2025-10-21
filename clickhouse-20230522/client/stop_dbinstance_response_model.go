// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopDBInstanceResponseBody) *StopDBInstanceResponse
	GetBody() *StopDBInstanceResponseBody
}

type StopDBInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDBInstanceResponse) GetBody() *StopDBInstanceResponseBody {
	return s.Body
}

func (s *StopDBInstanceResponse) SetHeaders(v map[string]*string) *StopDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopDBInstanceResponse) SetStatusCode(v int32) *StopDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDBInstanceResponse) SetBody(v *StopDBInstanceResponseBody) *StopDBInstanceResponse {
	s.Body = v
	return s
}

func (s *StopDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
