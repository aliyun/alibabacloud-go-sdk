// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkAppLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkAppLogResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkAppLogResponseBody) *GetSparkAppLogResponse
	GetBody() *GetSparkAppLogResponseBody
}

type GetSparkAppLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppLogResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkAppLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkAppLogResponse) GetBody() *GetSparkAppLogResponseBody {
	return s.Body
}

func (s *GetSparkAppLogResponse) SetHeaders(v map[string]*string) *GetSparkAppLogResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppLogResponse) SetStatusCode(v int32) *GetSparkAppLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppLogResponse) SetBody(v *GetSparkAppLogResponseBody) *GetSparkAppLogResponse {
	s.Body = v
	return s
}

func (s *GetSparkAppLogResponse) Validate() error {
	return dara.Validate(s)
}
