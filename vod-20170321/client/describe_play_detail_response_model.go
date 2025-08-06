// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayDetailResponseBody) *DescribePlayDetailResponse
	GetBody() *DescribePlayDetailResponseBody
}

type DescribePlayDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayDetailResponse) GetBody() *DescribePlayDetailResponseBody {
	return s.Body
}

func (s *DescribePlayDetailResponse) SetHeaders(v map[string]*string) *DescribePlayDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayDetailResponse) SetStatusCode(v int32) *DescribePlayDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayDetailResponse) SetBody(v *DescribePlayDetailResponseBody) *DescribePlayDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePlayDetailResponse) Validate() error {
	return dara.Validate(s)
}
