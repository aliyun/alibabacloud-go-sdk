// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultHttpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefaultHttpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefaultHttpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefaultHttpsResponseBody) *DescribeDefaultHttpsResponse
	GetBody() *DescribeDefaultHttpsResponseBody
}

type DescribeDefaultHttpsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefaultHttpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefaultHttpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultHttpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefaultHttpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefaultHttpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefaultHttpsResponse) GetBody() *DescribeDefaultHttpsResponseBody {
	return s.Body
}

func (s *DescribeDefaultHttpsResponse) SetHeaders(v map[string]*string) *DescribeDefaultHttpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefaultHttpsResponse) SetStatusCode(v int32) *DescribeDefaultHttpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefaultHttpsResponse) SetBody(v *DescribeDefaultHttpsResponseBody) *DescribeDefaultHttpsResponse {
	s.Body = v
	return s
}

func (s *DescribeDefaultHttpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
