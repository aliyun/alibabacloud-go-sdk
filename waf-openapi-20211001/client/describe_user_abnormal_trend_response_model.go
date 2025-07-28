// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAbnormalTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserAbnormalTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserAbnormalTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserAbnormalTrendResponseBody) *DescribeUserAbnormalTrendResponse
	GetBody() *DescribeUserAbnormalTrendResponseBody
}

type DescribeUserAbnormalTrendResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAbnormalTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAbnormalTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAbnormalTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAbnormalTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserAbnormalTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserAbnormalTrendResponse) GetBody() *DescribeUserAbnormalTrendResponseBody {
	return s.Body
}

func (s *DescribeUserAbnormalTrendResponse) SetHeaders(v map[string]*string) *DescribeUserAbnormalTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAbnormalTrendResponse) SetStatusCode(v int32) *DescribeUserAbnormalTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAbnormalTrendResponse) SetBody(v *DescribeUserAbnormalTrendResponseBody) *DescribeUserAbnormalTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeUserAbnormalTrendResponse) Validate() error {
	return dara.Validate(s)
}
