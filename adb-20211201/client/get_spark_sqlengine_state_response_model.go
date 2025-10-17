// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkSQLEngineStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkSQLEngineStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkSQLEngineStateResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkSQLEngineStateResponseBody) *GetSparkSQLEngineStateResponse
	GetBody() *GetSparkSQLEngineStateResponseBody
}

type GetSparkSQLEngineStateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkSQLEngineStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkSQLEngineStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkSQLEngineStateResponse) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkSQLEngineStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkSQLEngineStateResponse) GetBody() *GetSparkSQLEngineStateResponseBody {
	return s.Body
}

func (s *GetSparkSQLEngineStateResponse) SetHeaders(v map[string]*string) *GetSparkSQLEngineStateResponse {
	s.Headers = v
	return s
}

func (s *GetSparkSQLEngineStateResponse) SetStatusCode(v int32) *GetSparkSQLEngineStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkSQLEngineStateResponse) SetBody(v *GetSparkSQLEngineStateResponseBody) *GetSparkSQLEngineStateResponse {
	s.Body = v
	return s
}

func (s *GetSparkSQLEngineStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
