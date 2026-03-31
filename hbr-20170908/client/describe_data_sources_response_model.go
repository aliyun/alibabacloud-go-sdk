// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSourcesResponseBody) *DescribeDataSourcesResponse
	GetBody() *DescribeDataSourcesResponseBody
}

type DescribeDataSourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSourcesResponse) GetBody() *DescribeDataSourcesResponseBody {
	return s.Body
}

func (s *DescribeDataSourcesResponse) SetHeaders(v map[string]*string) *DescribeDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourcesResponse) SetStatusCode(v int32) *DescribeDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourcesResponse) SetBody(v *DescribeDataSourcesResponseBody) *DescribeDataSourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
