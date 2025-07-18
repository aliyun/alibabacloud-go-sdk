// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceErrorLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceErrorLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceErrorLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceErrorLogResponseBody) *DescribeDBInstanceErrorLogResponse
	GetBody() *DescribeDBInstanceErrorLogResponseBody
}

type DescribeDBInstanceErrorLogResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceErrorLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceErrorLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceErrorLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceErrorLogResponse) GetBody() *DescribeDBInstanceErrorLogResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceErrorLogResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceErrorLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceErrorLogResponse) SetStatusCode(v int32) *DescribeDBInstanceErrorLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponse) SetBody(v *DescribeDBInstanceErrorLogResponseBody) *DescribeDBInstanceErrorLogResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceErrorLogResponse) Validate() error {
	return dara.Validate(s)
}
