// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataLimitDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataLimitDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataLimitDetailResponseBody) *DescribeDataLimitDetailResponse
	GetBody() *DescribeDataLimitDetailResponseBody
}

type DescribeDataLimitDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataLimitDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataLimitDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataLimitDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataLimitDetailResponse) GetBody() *DescribeDataLimitDetailResponseBody {
	return s.Body
}

func (s *DescribeDataLimitDetailResponse) SetHeaders(v map[string]*string) *DescribeDataLimitDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataLimitDetailResponse) SetStatusCode(v int32) *DescribeDataLimitDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataLimitDetailResponse) SetBody(v *DescribeDataLimitDetailResponseBody) *DescribeDataLimitDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDataLimitDetailResponse) Validate() error {
	return dara.Validate(s)
}
