// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkSQLEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillSparkSQLEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillSparkSQLEngineResponse
	GetStatusCode() *int32
	SetBody(v *KillSparkSQLEngineResponseBody) *KillSparkSQLEngineResponse
	GetBody() *KillSparkSQLEngineResponseBody
}

type KillSparkSQLEngineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSparkSQLEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillSparkSQLEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s KillSparkSQLEngineResponse) GoString() string {
	return s.String()
}

func (s *KillSparkSQLEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillSparkSQLEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillSparkSQLEngineResponse) GetBody() *KillSparkSQLEngineResponseBody {
	return s.Body
}

func (s *KillSparkSQLEngineResponse) SetHeaders(v map[string]*string) *KillSparkSQLEngineResponse {
	s.Headers = v
	return s
}

func (s *KillSparkSQLEngineResponse) SetStatusCode(v int32) *KillSparkSQLEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSparkSQLEngineResponse) SetBody(v *KillSparkSQLEngineResponseBody) *KillSparkSQLEngineResponse {
	s.Body = v
	return s
}

func (s *KillSparkSQLEngineResponse) Validate() error {
	return dara.Validate(s)
}
