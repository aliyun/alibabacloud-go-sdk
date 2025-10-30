// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJDBCDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJDBCDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJDBCDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJDBCDataSourceResponseBody) *DescribeJDBCDataSourceResponse
	GetBody() *DescribeJDBCDataSourceResponseBody
}

type DescribeJDBCDataSourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJDBCDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJDBCDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJDBCDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeJDBCDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJDBCDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJDBCDataSourceResponse) GetBody() *DescribeJDBCDataSourceResponseBody {
	return s.Body
}

func (s *DescribeJDBCDataSourceResponse) SetHeaders(v map[string]*string) *DescribeJDBCDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeJDBCDataSourceResponse) SetStatusCode(v int32) *DescribeJDBCDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJDBCDataSourceResponse) SetBody(v *DescribeJDBCDataSourceResponseBody) *DescribeJDBCDataSourceResponse {
	s.Body = v
	return s
}

func (s *DescribeJDBCDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
