// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestAgenticDataSourceConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestAgenticDataSourceConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestAgenticDataSourceConnectionResponse
	GetStatusCode() *int32
	SetBody(v *TestAgenticDataSourceConnectionResponseBody) *TestAgenticDataSourceConnectionResponse
	GetBody() *TestAgenticDataSourceConnectionResponseBody
}

type TestAgenticDataSourceConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestAgenticDataSourceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestAgenticDataSourceConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s TestAgenticDataSourceConnectionResponse) GoString() string {
	return s.String()
}

func (s *TestAgenticDataSourceConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestAgenticDataSourceConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestAgenticDataSourceConnectionResponse) GetBody() *TestAgenticDataSourceConnectionResponseBody {
	return s.Body
}

func (s *TestAgenticDataSourceConnectionResponse) SetHeaders(v map[string]*string) *TestAgenticDataSourceConnectionResponse {
	s.Headers = v
	return s
}

func (s *TestAgenticDataSourceConnectionResponse) SetStatusCode(v int32) *TestAgenticDataSourceConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *TestAgenticDataSourceConnectionResponse) SetBody(v *TestAgenticDataSourceConnectionResponseBody) *TestAgenticDataSourceConnectionResponse {
	s.Body = v
	return s
}

func (s *TestAgenticDataSourceConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
