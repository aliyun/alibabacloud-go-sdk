// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterParametersResponseBody) *DescribeDBClusterParametersResponse
	GetBody() *DescribeDBClusterParametersResponseBody
}

type DescribeDBClusterParametersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterParametersResponse) GetBody() *DescribeDBClusterParametersResponseBody {
	return s.Body
}

func (s *DescribeDBClusterParametersResponse) SetHeaders(v map[string]*string) *DescribeDBClusterParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterParametersResponse) SetStatusCode(v int32) *DescribeDBClusterParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterParametersResponse) SetBody(v *DescribeDBClusterParametersResponseBody) *DescribeDBClusterParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterParametersResponse) Validate() error {
	return dara.Validate(s)
}
