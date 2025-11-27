// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceDetailResponseBody) *DescribeDBInstanceDetailResponse
	GetBody() *DescribeDBInstanceDetailResponseBody
}

type DescribeDBInstanceDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceDetailResponse) GetBody() *DescribeDBInstanceDetailResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceDetailResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDetailResponse) SetStatusCode(v int32) *DescribeDBInstanceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDetailResponse) SetBody(v *DescribeDBInstanceDetailResponseBody) *DescribeDBInstanceDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
