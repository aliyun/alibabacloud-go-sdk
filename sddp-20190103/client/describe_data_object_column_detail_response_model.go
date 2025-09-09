// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectColumnDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataObjectColumnDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataObjectColumnDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataObjectColumnDetailResponseBody) *DescribeDataObjectColumnDetailResponse
	GetBody() *DescribeDataObjectColumnDetailResponseBody
}

type DescribeDataObjectColumnDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataObjectColumnDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataObjectColumnDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataObjectColumnDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataObjectColumnDetailResponse) GetBody() *DescribeDataObjectColumnDetailResponseBody {
	return s.Body
}

func (s *DescribeDataObjectColumnDetailResponse) SetHeaders(v map[string]*string) *DescribeDataObjectColumnDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponse) SetStatusCode(v int32) *DescribeDataObjectColumnDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataObjectColumnDetailResponse) SetBody(v *DescribeDataObjectColumnDetailResponseBody) *DescribeDataObjectColumnDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDataObjectColumnDetailResponse) Validate() error {
	return dara.Validate(s)
}
