// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniSupportRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUniSupportRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUniSupportRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUniSupportRegionResponseBody) *DescribeUniSupportRegionResponse
	GetBody() *DescribeUniSupportRegionResponseBody
}

type DescribeUniSupportRegionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUniSupportRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUniSupportRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniSupportRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeUniSupportRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUniSupportRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUniSupportRegionResponse) GetBody() *DescribeUniSupportRegionResponseBody {
	return s.Body
}

func (s *DescribeUniSupportRegionResponse) SetHeaders(v map[string]*string) *DescribeUniSupportRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeUniSupportRegionResponse) SetStatusCode(v int32) *DescribeUniSupportRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUniSupportRegionResponse) SetBody(v *DescribeUniSupportRegionResponseBody) *DescribeUniSupportRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeUniSupportRegionResponse) Validate() error {
	return dara.Validate(s)
}
