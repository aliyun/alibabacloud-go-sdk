// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnConnectionLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnConnectionLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnConnectionLogsResponseBody) *DescribeVpnConnectionLogsResponse
	GetBody() *DescribeVpnConnectionLogsResponseBody
}

type DescribeVpnConnectionLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnConnectionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnConnectionLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnConnectionLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnConnectionLogsResponse) GetBody() *DescribeVpnConnectionLogsResponseBody {
	return s.Body
}

func (s *DescribeVpnConnectionLogsResponse) SetHeaders(v map[string]*string) *DescribeVpnConnectionLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnConnectionLogsResponse) SetStatusCode(v int32) *DescribeVpnConnectionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnConnectionLogsResponse) SetBody(v *DescribeVpnConnectionLogsResponseBody) *DescribeVpnConnectionLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnConnectionLogsResponse) Validate() error {
	return dara.Validate(s)
}
