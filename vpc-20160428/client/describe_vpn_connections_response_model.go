// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnConnectionsResponseBody) *DescribeVpnConnectionsResponse
	GetBody() *DescribeVpnConnectionsResponseBody
}

type DescribeVpnConnectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnConnectionsResponse) GetBody() *DescribeVpnConnectionsResponseBody {
	return s.Body
}

func (s *DescribeVpnConnectionsResponse) SetHeaders(v map[string]*string) *DescribeVpnConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnConnectionsResponse) SetStatusCode(v int32) *DescribeVpnConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnConnectionsResponse) SetBody(v *DescribeVpnConnectionsResponseBody) *DescribeVpnConnectionsResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
