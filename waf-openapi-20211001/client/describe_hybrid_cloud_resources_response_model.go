// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudResourcesResponseBody) *DescribeHybridCloudResourcesResponse
	GetBody() *DescribeHybridCloudResourcesResponseBody
}

type DescribeHybridCloudResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudResourcesResponse) GetBody() *DescribeHybridCloudResourcesResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudResourcesResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudResourcesResponse) SetStatusCode(v int32) *DescribeHybridCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponse) SetBody(v *DescribeHybridCloudResourcesResponseBody) *DescribeHybridCloudResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudResourcesResponse) Validate() error {
	return dara.Validate(s)
}
