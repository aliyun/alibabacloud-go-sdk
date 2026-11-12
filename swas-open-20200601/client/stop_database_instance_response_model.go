// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDatabaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDatabaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDatabaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopDatabaseInstanceResponseBody) *StopDatabaseInstanceResponse
	GetBody() *StopDatabaseInstanceResponseBody
}

type StopDatabaseInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDatabaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDatabaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDatabaseInstanceResponse) GetBody() *StopDatabaseInstanceResponseBody {
	return s.Body
}

func (s *StopDatabaseInstanceResponse) SetHeaders(v map[string]*string) *StopDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopDatabaseInstanceResponse) SetStatusCode(v int32) *StopDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDatabaseInstanceResponse) SetBody(v *StopDatabaseInstanceResponseBody) *StopDatabaseInstanceResponse {
	s.Body = v
	return s
}

func (s *StopDatabaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
