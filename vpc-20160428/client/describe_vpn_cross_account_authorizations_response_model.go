// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnCrossAccountAuthorizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnCrossAccountAuthorizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnCrossAccountAuthorizationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnCrossAccountAuthorizationsResponseBody) *DescribeVpnCrossAccountAuthorizationsResponse
	GetBody() *DescribeVpnCrossAccountAuthorizationsResponseBody
}

type DescribeVpnCrossAccountAuthorizationsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnCrossAccountAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnCrossAccountAuthorizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnCrossAccountAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnCrossAccountAuthorizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnCrossAccountAuthorizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnCrossAccountAuthorizationsResponse) GetBody() *DescribeVpnCrossAccountAuthorizationsResponseBody {
	return s.Body
}

func (s *DescribeVpnCrossAccountAuthorizationsResponse) SetHeaders(v map[string]*string) *DescribeVpnCrossAccountAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponse) SetStatusCode(v int32) *DescribeVpnCrossAccountAuthorizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponse) SetBody(v *DescribeVpnCrossAccountAuthorizationsResponseBody) *DescribeVpnCrossAccountAuthorizationsResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponse) Validate() error {
	return dara.Validate(s)
}
