// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSourceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSourceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSourceInstanceResponseBody) *DescribeDataSourceInstanceResponse
	GetBody() *DescribeDataSourceInstanceResponseBody
}

type DescribeDataSourceInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSourceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSourceInstanceResponse) GetBody() *DescribeDataSourceInstanceResponseBody {
	return s.Body
}

func (s *DescribeDataSourceInstanceResponse) SetHeaders(v map[string]*string) *DescribeDataSourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourceInstanceResponse) SetStatusCode(v int32) *DescribeDataSourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourceInstanceResponse) SetBody(v *DescribeDataSourceInstanceResponseBody) *DescribeDataSourceInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSourceInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
