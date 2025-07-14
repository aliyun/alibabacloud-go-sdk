// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdasContainersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdasContainersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdasContainersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdasContainersResponseBody) *DescribeEdasContainersResponse
	GetBody() *DescribeEdasContainersResponseBody
}

type DescribeEdasContainersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdasContainersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdasContainersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdasContainersResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdasContainersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdasContainersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdasContainersResponse) GetBody() *DescribeEdasContainersResponseBody {
	return s.Body
}

func (s *DescribeEdasContainersResponse) SetHeaders(v map[string]*string) *DescribeEdasContainersResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdasContainersResponse) SetStatusCode(v int32) *DescribeEdasContainersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdasContainersResponse) SetBody(v *DescribeEdasContainersResponseBody) *DescribeEdasContainersResponse {
	s.Body = v
	return s
}

func (s *DescribeEdasContainersResponse) Validate() error {
	return dara.Validate(s)
}
