// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTimeTopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetTimeTopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetTimeTopResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetTimeTopResponseBody) *DescribeInternetTimeTopResponse
	GetBody() *DescribeInternetTimeTopResponseBody
}

type DescribeInternetTimeTopResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetTimeTopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetTimeTopResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTimeTopResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetTimeTopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetTimeTopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetTimeTopResponse) GetBody() *DescribeInternetTimeTopResponseBody {
	return s.Body
}

func (s *DescribeInternetTimeTopResponse) SetHeaders(v map[string]*string) *DescribeInternetTimeTopResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetTimeTopResponse) SetStatusCode(v int32) *DescribeInternetTimeTopResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetTimeTopResponse) SetBody(v *DescribeInternetTimeTopResponseBody) *DescribeInternetTimeTopResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetTimeTopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
