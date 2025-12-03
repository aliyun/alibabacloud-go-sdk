// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancePacketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstancePacketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstancePacketsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstancePacketsResponseBody) *DescribeInstancePacketsResponse
	GetBody() *DescribeInstancePacketsResponseBody
}

type DescribeInstancePacketsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancePacketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancePacketsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePacketsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePacketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstancePacketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstancePacketsResponse) GetBody() *DescribeInstancePacketsResponseBody {
	return s.Body
}

func (s *DescribeInstancePacketsResponse) SetHeaders(v map[string]*string) *DescribeInstancePacketsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancePacketsResponse) SetStatusCode(v int32) *DescribeInstancePacketsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancePacketsResponse) SetBody(v *DescribeInstancePacketsResponseBody) *DescribeInstancePacketsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstancePacketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
