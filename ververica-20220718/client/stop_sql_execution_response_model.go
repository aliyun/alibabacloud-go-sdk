// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSqlExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSqlExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSqlExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StopSqlExecutionResponseBody) *StopSqlExecutionResponse
	GetBody() *StopSqlExecutionResponseBody
}

type StopSqlExecutionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSqlExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSqlExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSqlExecutionResponse) GoString() string {
	return s.String()
}

func (s *StopSqlExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSqlExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSqlExecutionResponse) GetBody() *StopSqlExecutionResponseBody {
	return s.Body
}

func (s *StopSqlExecutionResponse) SetHeaders(v map[string]*string) *StopSqlExecutionResponse {
	s.Headers = v
	return s
}

func (s *StopSqlExecutionResponse) SetStatusCode(v int32) *StopSqlExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSqlExecutionResponse) SetBody(v *StopSqlExecutionResponseBody) *StopSqlExecutionResponse {
	s.Body = v
	return s
}

func (s *StopSqlExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
