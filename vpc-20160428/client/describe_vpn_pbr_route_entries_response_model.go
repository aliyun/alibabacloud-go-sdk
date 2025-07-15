// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnPbrRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnPbrRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnPbrRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnPbrRouteEntriesResponseBody) *DescribeVpnPbrRouteEntriesResponse
	GetBody() *DescribeVpnPbrRouteEntriesResponseBody
}

type DescribeVpnPbrRouteEntriesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnPbrRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnPbrRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnPbrRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnPbrRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnPbrRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnPbrRouteEntriesResponse) GetBody() *DescribeVpnPbrRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribeVpnPbrRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeVpnPbrRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponse) SetStatusCode(v int32) *DescribeVpnPbrRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponse) SetBody(v *DescribeVpnPbrRouteEntriesResponseBody) *DescribeVpnPbrRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
