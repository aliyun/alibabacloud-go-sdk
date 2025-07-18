// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIPArrayListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceIPArrayListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceIPArrayListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceIPArrayListResponseBody) *DescribeDBInstanceIPArrayListResponse
	GetBody() *DescribeDBInstanceIPArrayListResponseBody
}

type DescribeDBInstanceIPArrayListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceIPArrayListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceIPArrayListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceIPArrayListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceIPArrayListResponse) GetBody() *DescribeDBInstanceIPArrayListResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceIPArrayListResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceIPArrayListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) SetStatusCode(v int32) *DescribeDBInstanceIPArrayListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) SetBody(v *DescribeDBInstanceIPArrayListResponseBody) *DescribeDBInstanceIPArrayListResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) Validate() error {
	return dara.Validate(s)
}
