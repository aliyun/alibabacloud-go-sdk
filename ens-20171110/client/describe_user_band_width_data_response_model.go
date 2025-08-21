// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBandWidthDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserBandWidthDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserBandWidthDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserBandWidthDataResponseBody) *DescribeUserBandWidthDataResponse
	GetBody() *DescribeUserBandWidthDataResponseBody
}

type DescribeUserBandWidthDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBandWidthDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBandWidthDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBandWidthDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserBandWidthDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserBandWidthDataResponse) GetBody() *DescribeUserBandWidthDataResponseBody {
	return s.Body
}

func (s *DescribeUserBandWidthDataResponse) SetHeaders(v map[string]*string) *DescribeUserBandWidthDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBandWidthDataResponse) SetStatusCode(v int32) *DescribeUserBandWidthDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBandWidthDataResponse) SetBody(v *DescribeUserBandWidthDataResponseBody) *DescribeUserBandWidthDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUserBandWidthDataResponse) Validate() error {
	return dara.Validate(s)
}
