// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsCommodityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsCommodityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsCommodityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsCommodityResponseBody) *DescribeRdsCommodityResponse
	GetBody() *DescribeRdsCommodityResponseBody
}

type DescribeRdsCommodityResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsCommodityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsCommodityResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsCommodityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsCommodityResponse) GetBody() *DescribeRdsCommodityResponseBody {
	return s.Body
}

func (s *DescribeRdsCommodityResponse) SetHeaders(v map[string]*string) *DescribeRdsCommodityResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsCommodityResponse) SetStatusCode(v int32) *DescribeRdsCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsCommodityResponse) SetBody(v *DescribeRdsCommodityResponseBody) *DescribeRdsCommodityResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsCommodityResponse) Validate() error {
	return dara.Validate(s)
}
