// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGadInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGadInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGadInstancesResponseBody) *DescribeGadInstancesResponse
	GetBody() *DescribeGadInstancesResponseBody
}

type DescribeGadInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGadInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGadInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGadInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGadInstancesResponse) GetBody() *DescribeGadInstancesResponseBody {
	return s.Body
}

func (s *DescribeGadInstancesResponse) SetHeaders(v map[string]*string) *DescribeGadInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGadInstancesResponse) SetStatusCode(v int32) *DescribeGadInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGadInstancesResponse) SetBody(v *DescribeGadInstancesResponseBody) *DescribeGadInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeGadInstancesResponse) Validate() error {
	return dara.Validate(s)
}
