// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourcePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBResourcePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBResourcePoolResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBResourcePoolResponseBody) *DescribeDBResourcePoolResponse
	GetBody() *DescribeDBResourcePoolResponseBody
}

type DescribeDBResourcePoolResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBResourcePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBResourcePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourcePoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBResourcePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBResourcePoolResponse) GetBody() *DescribeDBResourcePoolResponseBody {
	return s.Body
}

func (s *DescribeDBResourcePoolResponse) SetHeaders(v map[string]*string) *DescribeDBResourcePoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBResourcePoolResponse) SetStatusCode(v int32) *DescribeDBResourcePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBResourcePoolResponse) SetBody(v *DescribeDBResourcePoolResponseBody) *DescribeDBResourcePoolResponse {
	s.Body = v
	return s
}

func (s *DescribeDBResourcePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
