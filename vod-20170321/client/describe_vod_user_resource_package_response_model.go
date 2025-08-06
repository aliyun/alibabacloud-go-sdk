// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserResourcePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodUserResourcePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodUserResourcePackageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodUserResourcePackageResponseBody) *DescribeVodUserResourcePackageResponse
	GetBody() *DescribeVodUserResourcePackageResponseBody
}

type DescribeVodUserResourcePackageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodUserResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodUserResourcePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserResourcePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodUserResourcePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodUserResourcePackageResponse) GetBody() *DescribeVodUserResourcePackageResponseBody {
	return s.Body
}

func (s *DescribeVodUserResourcePackageResponse) SetHeaders(v map[string]*string) *DescribeVodUserResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserResourcePackageResponse) SetStatusCode(v int32) *DescribeVodUserResourcePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodUserResourcePackageResponse) SetBody(v *DescribeVodUserResourcePackageResponseBody) *DescribeVodUserResourcePackageResponse {
	s.Body = v
	return s
}

func (s *DescribeVodUserResourcePackageResponse) Validate() error {
	return dara.Validate(s)
}
