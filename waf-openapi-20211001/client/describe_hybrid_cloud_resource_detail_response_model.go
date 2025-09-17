// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudResourceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudResourceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudResourceDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudResourceDetailResponseBody) *DescribeHybridCloudResourceDetailResponse
	GetBody() *DescribeHybridCloudResourceDetailResponseBody
}

type DescribeHybridCloudResourceDetailResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudResourceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudResourceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudResourceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudResourceDetailResponse) GetBody() *DescribeHybridCloudResourceDetailResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudResourceDetailResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudResourceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponse) SetStatusCode(v int32) *DescribeHybridCloudResourceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponse) SetBody(v *DescribeHybridCloudResourceDetailResponseBody) *DescribeHybridCloudResourceDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponse) Validate() error {
	return dara.Validate(s)
}
