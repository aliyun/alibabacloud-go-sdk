// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmInstancesResponseBody) *DescribeGtmInstancesResponse
	GetBody() *DescribeGtmInstancesResponseBody
}

type DescribeGtmInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmInstancesResponse) GetBody() *DescribeGtmInstancesResponseBody {
	return s.Body
}

func (s *DescribeGtmInstancesResponse) SetHeaders(v map[string]*string) *DescribeGtmInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstancesResponse) SetStatusCode(v int32) *DescribeGtmInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmInstancesResponse) SetBody(v *DescribeGtmInstancesResponseBody) *DescribeGtmInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmInstancesResponse) Validate() error {
	return dara.Validate(s)
}
