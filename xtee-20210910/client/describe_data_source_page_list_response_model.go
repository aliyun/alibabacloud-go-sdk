// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourcePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSourcePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSourcePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSourcePageListResponseBody) *DescribeDataSourcePageListResponse
	GetBody() *DescribeDataSourcePageListResponseBody
}

type DescribeDataSourcePageListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourcePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourcePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSourcePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSourcePageListResponse) GetBody() *DescribeDataSourcePageListResponseBody {
	return s.Body
}

func (s *DescribeDataSourcePageListResponse) SetHeaders(v map[string]*string) *DescribeDataSourcePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourcePageListResponse) SetStatusCode(v int32) *DescribeDataSourcePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourcePageListResponse) SetBody(v *DescribeDataSourcePageListResponseBody) *DescribeDataSourcePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSourcePageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
