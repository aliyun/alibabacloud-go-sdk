// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSourceResponseBody) *DescribeDataSourceResponse
	GetBody() *DescribeDataSourceResponseBody
}

type DescribeDataSourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSourceResponse) GetBody() *DescribeDataSourceResponseBody {
	return s.Body
}

func (s *DescribeDataSourceResponse) SetHeaders(v map[string]*string) *DescribeDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourceResponse) SetStatusCode(v int32) *DescribeDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourceResponse) SetBody(v *DescribeDataSourceResponseBody) *DescribeDataSourceResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
