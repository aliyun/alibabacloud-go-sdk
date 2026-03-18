// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceListResponseBody) *DescribeInstanceListResponse
	GetBody() *DescribeInstanceListResponseBody
}

type DescribeInstanceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceListResponse) GetBody() *DescribeInstanceListResponseBody {
	return s.Body
}

func (s *DescribeInstanceListResponse) SetHeaders(v map[string]*string) *DescribeInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceListResponse) SetStatusCode(v int32) *DescribeInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceListResponse) SetBody(v *DescribeInstanceListResponseBody) *DescribeInstanceListResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
