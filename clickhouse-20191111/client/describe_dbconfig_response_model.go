// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBConfigResponseBody) *DescribeDBConfigResponse
	GetBody() *DescribeDBConfigResponseBody
}

type DescribeDBConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBConfigResponse) GetBody() *DescribeDBConfigResponseBody {
	return s.Body
}

func (s *DescribeDBConfigResponse) SetHeaders(v map[string]*string) *DescribeDBConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBConfigResponse) SetStatusCode(v int32) *DescribeDBConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBConfigResponse) SetBody(v *DescribeDBConfigResponseBody) *DescribeDBConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDBConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
