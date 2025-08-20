// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsJobDetailResponseBody) *DescribeApsJobDetailResponse
	GetBody() *DescribeApsJobDetailResponseBody
}

type DescribeApsJobDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsJobDetailResponse) GetBody() *DescribeApsJobDetailResponseBody {
	return s.Body
}

func (s *DescribeApsJobDetailResponse) SetHeaders(v map[string]*string) *DescribeApsJobDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsJobDetailResponse) SetStatusCode(v int32) *DescribeApsJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsJobDetailResponse) SetBody(v *DescribeApsJobDetailResponseBody) *DescribeApsJobDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeApsJobDetailResponse) Validate() error {
	return dara.Validate(s)
}
