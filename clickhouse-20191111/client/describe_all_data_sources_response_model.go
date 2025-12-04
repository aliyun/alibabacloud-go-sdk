// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllDataSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllDataSourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllDataSourcesResponseBody) *DescribeAllDataSourcesResponse
	GetBody() *DescribeAllDataSourcesResponseBody
}

type DescribeAllDataSourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllDataSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllDataSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllDataSourcesResponse) GetBody() *DescribeAllDataSourcesResponseBody {
	return s.Body
}

func (s *DescribeAllDataSourcesResponse) SetHeaders(v map[string]*string) *DescribeAllDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllDataSourcesResponse) SetStatusCode(v int32) *DescribeAllDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllDataSourcesResponse) SetBody(v *DescribeAllDataSourcesResponseBody) *DescribeAllDataSourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeAllDataSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
