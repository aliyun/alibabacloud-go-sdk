// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnRouteEntriesResponseBody) *DescribeVpnRouteEntriesResponse
	GetBody() *DescribeVpnRouteEntriesResponseBody
}

type DescribeVpnRouteEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnRouteEntriesResponse) GetBody() *DescribeVpnRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribeVpnRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeVpnRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnRouteEntriesResponse) SetStatusCode(v int32) *DescribeVpnRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponse) SetBody(v *DescribeVpnRouteEntriesResponseBody) *DescribeVpnRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
