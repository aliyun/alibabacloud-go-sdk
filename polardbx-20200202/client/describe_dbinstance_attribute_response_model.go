// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceAttributeResponseBody) *DescribeDBInstanceAttributeResponse
	GetBody() *DescribeDBInstanceAttributeResponseBody
}

type DescribeDBInstanceAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceAttributeResponse) GetBody() *DescribeDBInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetStatusCode(v int32) *DescribeDBInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetBody(v *DescribeDBInstanceAttributeResponseBody) *DescribeDBInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
