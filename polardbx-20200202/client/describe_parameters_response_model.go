// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse
	GetBody() *DescribeParametersResponseBody
}

type DescribeParametersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParametersResponse) GetBody() *DescribeParametersResponseBody {
	return s.Body
}

func (s *DescribeParametersResponse) SetHeaders(v map[string]*string) *DescribeParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersResponse) SetStatusCode(v int32) *DescribeParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeParametersResponse) Validate() error {
	return dara.Validate(s)
}
