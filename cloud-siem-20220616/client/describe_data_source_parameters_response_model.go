// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourceParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataSourceParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataSourceParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataSourceParametersResponseBody) *DescribeDataSourceParametersResponse
	GetBody() *DescribeDataSourceParametersResponseBody
}

type DescribeDataSourceParametersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataSourceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataSourceParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourceParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataSourceParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataSourceParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataSourceParametersResponse) GetBody() *DescribeDataSourceParametersResponseBody {
	return s.Body
}

func (s *DescribeDataSourceParametersResponse) SetHeaders(v map[string]*string) *DescribeDataSourceParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataSourceParametersResponse) SetStatusCode(v int32) *DescribeDataSourceParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataSourceParametersResponse) SetBody(v *DescribeDataSourceParametersResponseBody) *DescribeDataSourceParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeDataSourceParametersResponse) Validate() error {
	return dara.Validate(s)
}
