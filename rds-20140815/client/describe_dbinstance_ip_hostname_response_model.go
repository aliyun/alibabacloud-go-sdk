// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIpHostnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceIpHostnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceIpHostnameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceIpHostnameResponseBody) *DescribeDBInstanceIpHostnameResponse
	GetBody() *DescribeDBInstanceIpHostnameResponseBody
}

type DescribeDBInstanceIpHostnameResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceIpHostnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceIpHostnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIpHostnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIpHostnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceIpHostnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceIpHostnameResponse) GetBody() *DescribeDBInstanceIpHostnameResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceIpHostnameResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceIpHostnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceIpHostnameResponse) SetStatusCode(v int32) *DescribeDBInstanceIpHostnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameResponse) SetBody(v *DescribeDBInstanceIpHostnameResponseBody) *DescribeDBInstanceIpHostnameResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceIpHostnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
