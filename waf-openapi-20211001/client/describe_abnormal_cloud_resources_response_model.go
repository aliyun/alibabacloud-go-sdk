// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbnormalCloudResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAbnormalCloudResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAbnormalCloudResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAbnormalCloudResourcesResponseBody) *DescribeAbnormalCloudResourcesResponse
	GetBody() *DescribeAbnormalCloudResourcesResponseBody
}

type DescribeAbnormalCloudResourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAbnormalCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAbnormalCloudResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalCloudResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAbnormalCloudResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAbnormalCloudResourcesResponse) GetBody() *DescribeAbnormalCloudResourcesResponseBody {
	return s.Body
}

func (s *DescribeAbnormalCloudResourcesResponse) SetHeaders(v map[string]*string) *DescribeAbnormalCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponse) SetStatusCode(v int32) *DescribeAbnormalCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponse) SetBody(v *DescribeAbnormalCloudResourcesResponseBody) *DescribeAbnormalCloudResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponse) Validate() error {
	return dara.Validate(s)
}
