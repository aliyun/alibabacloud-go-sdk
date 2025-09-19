// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExposedInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExposedInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExposedInstanceListResponseBody) *DescribeExposedInstanceListResponse
	GetBody() *DescribeExposedInstanceListResponseBody
}

type DescribeExposedInstanceListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExposedInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExposedInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExposedInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExposedInstanceListResponse) GetBody() *DescribeExposedInstanceListResponseBody {
	return s.Body
}

func (s *DescribeExposedInstanceListResponse) SetHeaders(v map[string]*string) *DescribeExposedInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedInstanceListResponse) SetStatusCode(v int32) *DescribeExposedInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExposedInstanceListResponse) SetBody(v *DescribeExposedInstanceListResponseBody) *DescribeExposedInstanceListResponse {
	s.Body = v
	return s
}

func (s *DescribeExposedInstanceListResponse) Validate() error {
	return dara.Validate(s)
}
