// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceProxyConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceProxyConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceProxyConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceProxyConfigurationResponseBody) *DescribeDBInstanceProxyConfigurationResponse
	GetBody() *DescribeDBInstanceProxyConfigurationResponseBody
}

type DescribeDBInstanceProxyConfigurationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceProxyConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceProxyConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceProxyConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceProxyConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceProxyConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceProxyConfigurationResponse) GetBody() *DescribeDBInstanceProxyConfigurationResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceProxyConfigurationResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceProxyConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationResponse) SetStatusCode(v int32) *DescribeDBInstanceProxyConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationResponse) SetBody(v *DescribeDBInstanceProxyConfigurationResponseBody) *DescribeDBInstanceProxyConfigurationResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceProxyConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
