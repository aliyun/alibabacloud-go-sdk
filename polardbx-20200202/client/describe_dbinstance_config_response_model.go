// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceConfigResponseBody) *DescribeDBInstanceConfigResponse
	GetBody() *DescribeDBInstanceConfigResponseBody
}

type DescribeDBInstanceConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceConfigResponse) GetBody() *DescribeDBInstanceConfigResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceConfigResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceConfigResponse) SetStatusCode(v int32) *DescribeDBInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceConfigResponse) SetBody(v *DescribeDBInstanceConfigResponseBody) *DescribeDBInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceConfigResponse) Validate() error {
	return dara.Validate(s)
}
