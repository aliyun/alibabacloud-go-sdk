// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSqlExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSqlExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSqlExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StartSqlExecutionResponseBody) *StartSqlExecutionResponse
	GetBody() *StartSqlExecutionResponseBody
}

type StartSqlExecutionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSqlExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSqlExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSqlExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartSqlExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSqlExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSqlExecutionResponse) GetBody() *StartSqlExecutionResponseBody {
	return s.Body
}

func (s *StartSqlExecutionResponse) SetHeaders(v map[string]*string) *StartSqlExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartSqlExecutionResponse) SetStatusCode(v int32) *StartSqlExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSqlExecutionResponse) SetBody(v *StartSqlExecutionResponseBody) *StartSqlExecutionResponse {
	s.Body = v
	return s
}

func (s *StartSqlExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
