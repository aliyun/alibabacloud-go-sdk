// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnSslServerLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnSslServerLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnSslServerLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnSslServerLogsResponseBody) *DescribeVpnSslServerLogsResponse
	GetBody() *DescribeVpnSslServerLogsResponseBody
}

type DescribeVpnSslServerLogsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnSslServerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnSslServerLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnSslServerLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnSslServerLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnSslServerLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnSslServerLogsResponse) GetBody() *DescribeVpnSslServerLogsResponseBody {
	return s.Body
}

func (s *DescribeVpnSslServerLogsResponse) SetHeaders(v map[string]*string) *DescribeVpnSslServerLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnSslServerLogsResponse) SetStatusCode(v int32) *DescribeVpnSslServerLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnSslServerLogsResponse) SetBody(v *DescribeVpnSslServerLogsResponseBody) *DescribeVpnSslServerLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnSslServerLogsResponse) Validate() error {
	return dara.Validate(s)
}
