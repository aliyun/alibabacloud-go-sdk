// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestDataSourceConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestDataSourceConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestDataSourceConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *TestDataSourceConnectivityResponseBody) *TestDataSourceConnectivityResponse
	GetBody() *TestDataSourceConnectivityResponseBody
}

type TestDataSourceConnectivityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestDataSourceConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestDataSourceConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s TestDataSourceConnectivityResponse) GoString() string {
	return s.String()
}

func (s *TestDataSourceConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestDataSourceConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestDataSourceConnectivityResponse) GetBody() *TestDataSourceConnectivityResponseBody {
	return s.Body
}

func (s *TestDataSourceConnectivityResponse) SetHeaders(v map[string]*string) *TestDataSourceConnectivityResponse {
	s.Headers = v
	return s
}

func (s *TestDataSourceConnectivityResponse) SetStatusCode(v int32) *TestDataSourceConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *TestDataSourceConnectivityResponse) SetBody(v *TestDataSourceConnectivityResponseBody) *TestDataSourceConnectivityResponse {
	s.Body = v
	return s
}

func (s *TestDataSourceConnectivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
