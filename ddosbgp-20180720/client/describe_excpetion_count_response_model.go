// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcpetionCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExcpetionCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExcpetionCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExcpetionCountResponseBody) *DescribeExcpetionCountResponse
	GetBody() *DescribeExcpetionCountResponseBody
}

type DescribeExcpetionCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExcpetionCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExcpetionCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcpetionCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExcpetionCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExcpetionCountResponse) GetBody() *DescribeExcpetionCountResponseBody {
	return s.Body
}

func (s *DescribeExcpetionCountResponse) SetHeaders(v map[string]*string) *DescribeExcpetionCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeExcpetionCountResponse) SetStatusCode(v int32) *DescribeExcpetionCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExcpetionCountResponse) SetBody(v *DescribeExcpetionCountResponseBody) *DescribeExcpetionCountResponse {
	s.Body = v
	return s
}

func (s *DescribeExcpetionCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
