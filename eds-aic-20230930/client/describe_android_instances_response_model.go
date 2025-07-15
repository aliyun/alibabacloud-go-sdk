// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAndroidInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAndroidInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAndroidInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAndroidInstancesResponseBody) *DescribeAndroidInstancesResponse
	GetBody() *DescribeAndroidInstancesResponseBody
}

type DescribeAndroidInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAndroidInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAndroidInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAndroidInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAndroidInstancesResponse) GetBody() *DescribeAndroidInstancesResponseBody {
	return s.Body
}

func (s *DescribeAndroidInstancesResponse) SetHeaders(v map[string]*string) *DescribeAndroidInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAndroidInstancesResponse) SetStatusCode(v int32) *DescribeAndroidInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAndroidInstancesResponse) SetBody(v *DescribeAndroidInstancesResponseBody) *DescribeAndroidInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeAndroidInstancesResponse) Validate() error {
	return dara.Validate(s)
}
