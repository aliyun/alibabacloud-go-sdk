// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataSkewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceDataSkewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceDataSkewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceDataSkewResponseBody) *DescribeDBInstanceDataSkewResponse
	GetBody() *DescribeDBInstanceDataSkewResponseBody
}

type DescribeDBInstanceDataSkewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceDataSkewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceDataSkewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceDataSkewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceDataSkewResponse) GetBody() *DescribeDBInstanceDataSkewResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceDataSkewResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataSkewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataSkewResponse) SetStatusCode(v int32) *DescribeDBInstanceDataSkewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponse) SetBody(v *DescribeDBInstanceDataSkewResponseBody) *DescribeDBInstanceDataSkewResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceDataSkewResponse) Validate() error {
	return dara.Validate(s)
}
