// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSparkSQLEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSparkSQLEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSparkSQLEngineResponse
	GetStatusCode() *int32
	SetBody(v *StartSparkSQLEngineResponseBody) *StartSparkSQLEngineResponse
	GetBody() *StartSparkSQLEngineResponseBody
}

type StartSparkSQLEngineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSparkSQLEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSparkSQLEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSparkSQLEngineResponse) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSparkSQLEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSparkSQLEngineResponse) GetBody() *StartSparkSQLEngineResponseBody {
	return s.Body
}

func (s *StartSparkSQLEngineResponse) SetHeaders(v map[string]*string) *StartSparkSQLEngineResponse {
	s.Headers = v
	return s
}

func (s *StartSparkSQLEngineResponse) SetStatusCode(v int32) *StartSparkSQLEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSparkSQLEngineResponse) SetBody(v *StartSparkSQLEngineResponseBody) *StartSparkSQLEngineResponse {
	s.Body = v
	return s
}

func (s *StartSparkSQLEngineResponse) Validate() error {
	return dara.Validate(s)
}
