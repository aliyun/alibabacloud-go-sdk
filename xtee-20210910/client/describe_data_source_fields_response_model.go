// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSourceFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSourceFieldsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSourceFieldsResponseBody) *DescribeDataSourceFieldsResponse
	GetBody() *DescribeDataSourceFieldsResponseBody
}

type DescribeDataSourceFieldsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourceFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourceFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceFieldsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSourceFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSourceFieldsResponse) GetBody() *DescribeDataSourceFieldsResponseBody {
	return s.Body
}

func (s *DescribeDataSourceFieldsResponse) SetHeaders(v map[string]*string) *DescribeDataSourceFieldsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourceFieldsResponse) SetStatusCode(v int32) *DescribeDataSourceFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourceFieldsResponse) SetBody(v *DescribeDataSourceFieldsResponseBody) *DescribeDataSourceFieldsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSourceFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
