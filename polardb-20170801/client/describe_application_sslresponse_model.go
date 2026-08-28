// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationSSLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationSSLResponseBody) *DescribeApplicationSSLResponse
	GetBody() *DescribeApplicationSSLResponseBody
}

type DescribeApplicationSSLResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationSSLResponse) GetBody() *DescribeApplicationSSLResponseBody {
	return s.Body
}

func (s *DescribeApplicationSSLResponse) SetHeaders(v map[string]*string) *DescribeApplicationSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationSSLResponse) SetStatusCode(v int32) *DescribeApplicationSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationSSLResponse) SetBody(v *DescribeApplicationSSLResponseBody) *DescribeApplicationSSLResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationSSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
