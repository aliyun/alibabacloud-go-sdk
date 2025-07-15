// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemParametersResponseBody) *DescribeSystemParametersResponse
	GetBody() *DescribeSystemParametersResponseBody
}

type DescribeSystemParametersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemParametersResponse) GetBody() *DescribeSystemParametersResponseBody {
	return s.Body
}

func (s *DescribeSystemParametersResponse) SetHeaders(v map[string]*string) *DescribeSystemParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemParametersResponse) SetStatusCode(v int32) *DescribeSystemParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemParametersResponse) SetBody(v *DescribeSystemParametersResponseBody) *DescribeSystemParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemParametersResponse) Validate() error {
	return dara.Validate(s)
}
