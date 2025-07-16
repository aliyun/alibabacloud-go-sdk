// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceTopologyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceTopologyResponseBody) *DescribeDBInstanceTopologyResponse
	GetBody() *DescribeDBInstanceTopologyResponseBody
}

type DescribeDBInstanceTopologyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceTopologyResponse) GetBody() *DescribeDBInstanceTopologyResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceTopologyResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceTopologyResponse) SetStatusCode(v int32) *DescribeDBInstanceTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponse) SetBody(v *DescribeDBInstanceTopologyResponseBody) *DescribeDBInstanceTopologyResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceTopologyResponse) Validate() error {
	return dara.Validate(s)
}
