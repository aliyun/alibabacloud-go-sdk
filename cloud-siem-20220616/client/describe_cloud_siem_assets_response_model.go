// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudSiemAssetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudSiemAssetsResponseBody) *DescribeCloudSiemAssetsResponse
	GetBody() *DescribeCloudSiemAssetsResponseBody
}

type DescribeCloudSiemAssetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudSiemAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudSiemAssetsResponse) GetBody() *DescribeCloudSiemAssetsResponseBody {
	return s.Body
}

func (s *DescribeCloudSiemAssetsResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemAssetsResponse) SetStatusCode(v int32) *DescribeCloudSiemAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponse) SetBody(v *DescribeCloudSiemAssetsResponseBody) *DescribeCloudSiemAssetsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudSiemAssetsResponse) Validate() error {
	return dara.Validate(s)
}
