// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestConnectivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestConnectivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestConnectivityResponse
	GetStatusCode() *int32
	SetBody(v *TestConnectivityResponseBody) *TestConnectivityResponse
	GetBody() *TestConnectivityResponseBody
}

type TestConnectivityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestConnectivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestConnectivityResponse) String() string {
	return dara.Prettify(s)
}

func (s TestConnectivityResponse) GoString() string {
	return s.String()
}

func (s *TestConnectivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestConnectivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestConnectivityResponse) GetBody() *TestConnectivityResponseBody {
	return s.Body
}

func (s *TestConnectivityResponse) SetHeaders(v map[string]*string) *TestConnectivityResponse {
	s.Headers = v
	return s
}

func (s *TestConnectivityResponse) SetStatusCode(v int32) *TestConnectivityResponse {
	s.StatusCode = &v
	return s
}

func (s *TestConnectivityResponse) SetBody(v *TestConnectivityResponseBody) *TestConnectivityResponse {
	s.Body = v
	return s
}

func (s *TestConnectivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
