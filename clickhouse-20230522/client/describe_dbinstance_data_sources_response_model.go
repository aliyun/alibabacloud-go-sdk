// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceDataSourcesResponseBody) *DescribeDBInstanceDataSourcesResponse
	GetBody() *DescribeDBInstanceDataSourcesResponseBody
}

type DescribeDBInstanceDataSourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceDataSourcesResponse) GetBody() *DescribeDBInstanceDataSourcesResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceDataSourcesResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponse) SetStatusCode(v int32) *DescribeDBInstanceDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponse) SetBody(v *DescribeDBInstanceDataSourcesResponseBody) *DescribeDBInstanceDataSourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
