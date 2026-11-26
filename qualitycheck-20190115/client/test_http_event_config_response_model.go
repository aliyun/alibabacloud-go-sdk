// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestHttpEventConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestHttpEventConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestHttpEventConfigResponse
	GetStatusCode() *int32
	SetBody(v *TestHttpEventConfigResponseBody) *TestHttpEventConfigResponse
	GetBody() *TestHttpEventConfigResponseBody
}

type TestHttpEventConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestHttpEventConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestHttpEventConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s TestHttpEventConfigResponse) GoString() string {
	return s.String()
}

func (s *TestHttpEventConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestHttpEventConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestHttpEventConfigResponse) GetBody() *TestHttpEventConfigResponseBody {
	return s.Body
}

func (s *TestHttpEventConfigResponse) SetHeaders(v map[string]*string) *TestHttpEventConfigResponse {
	s.Headers = v
	return s
}

func (s *TestHttpEventConfigResponse) SetStatusCode(v int32) *TestHttpEventConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *TestHttpEventConfigResponse) SetBody(v *TestHttpEventConfigResponseBody) *TestHttpEventConfigResponse {
	s.Body = v
	return s
}

func (s *TestHttpEventConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
