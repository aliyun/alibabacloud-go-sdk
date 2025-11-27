// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePromoteActivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstancePromoteActivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstancePromoteActivityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstancePromoteActivityResponseBody) *DescribeDBInstancePromoteActivityResponse
	GetBody() *DescribeDBInstancePromoteActivityResponseBody
}

type DescribeDBInstancePromoteActivityResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstancePromoteActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstancePromoteActivityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePromoteActivityResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePromoteActivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstancePromoteActivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstancePromoteActivityResponse) GetBody() *DescribeDBInstancePromoteActivityResponseBody {
	return s.Body
}

func (s *DescribeDBInstancePromoteActivityResponse) SetHeaders(v map[string]*string) *DescribeDBInstancePromoteActivityResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponse) SetStatusCode(v int32) *DescribeDBInstancePromoteActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponse) SetBody(v *DescribeDBInstancePromoteActivityResponseBody) *DescribeDBInstancePromoteActivityResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
