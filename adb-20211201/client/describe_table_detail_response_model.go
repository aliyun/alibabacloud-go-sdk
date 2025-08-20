// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTableDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTableDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTableDetailResponseBody) *DescribeTableDetailResponse
	GetBody() *DescribeTableDetailResponseBody
}

type DescribeTableDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTableDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTableDetailResponse) GetBody() *DescribeTableDetailResponseBody {
	return s.Body
}

func (s *DescribeTableDetailResponse) SetHeaders(v map[string]*string) *DescribeTableDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableDetailResponse) SetStatusCode(v int32) *DescribeTableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableDetailResponse) SetBody(v *DescribeTableDetailResponseBody) *DescribeTableDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeTableDetailResponse) Validate() error {
	return dara.Validate(s)
}
