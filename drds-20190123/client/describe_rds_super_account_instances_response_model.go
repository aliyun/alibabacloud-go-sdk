// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsSuperAccountInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsSuperAccountInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsSuperAccountInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsSuperAccountInstancesResponseBody) *DescribeRdsSuperAccountInstancesResponse
	GetBody() *DescribeRdsSuperAccountInstancesResponseBody
}

type DescribeRdsSuperAccountInstancesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsSuperAccountInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsSuperAccountInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsSuperAccountInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsSuperAccountInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsSuperAccountInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsSuperAccountInstancesResponse) GetBody() *DescribeRdsSuperAccountInstancesResponseBody {
	return s.Body
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetHeaders(v map[string]*string) *DescribeRdsSuperAccountInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetStatusCode(v int32) *DescribeRdsSuperAccountInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponse) SetBody(v *DescribeRdsSuperAccountInstancesResponseBody) *DescribeRdsSuperAccountInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsSuperAccountInstancesResponse) Validate() error {
	return dara.Validate(s)
}
