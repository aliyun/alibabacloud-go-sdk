// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpcZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpcZoneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpcZoneResponseBody) *DescribeVpcZoneResponse
	GetBody() *DescribeVpcZoneResponseBody
}

type DescribeVpcZoneResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcZoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpcZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpcZoneResponse) GetBody() *DescribeVpcZoneResponseBody {
	return s.Body
}

func (s *DescribeVpcZoneResponse) SetHeaders(v map[string]*string) *DescribeVpcZoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcZoneResponse) SetStatusCode(v int32) *DescribeVpcZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcZoneResponse) SetBody(v *DescribeVpcZoneResponseBody) *DescribeVpcZoneResponse {
	s.Body = v
	return s
}

func (s *DescribeVpcZoneResponse) Validate() error {
	return dara.Validate(s)
}
