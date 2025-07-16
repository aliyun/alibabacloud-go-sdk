// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceTDEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceTDEResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceTDEResponseBody) *DescribeDBInstanceTDEResponse
	GetBody() *DescribeDBInstanceTDEResponseBody
}

type DescribeDBInstanceTDEResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceTDEResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceTDEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceTDEResponse) GetBody() *DescribeDBInstanceTDEResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceTDEResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceTDEResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceTDEResponse) SetStatusCode(v int32) *DescribeDBInstanceTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceTDEResponse) SetBody(v *DescribeDBInstanceTDEResponseBody) *DescribeDBInstanceTDEResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceTDEResponse) Validate() error {
	return dara.Validate(s)
}
