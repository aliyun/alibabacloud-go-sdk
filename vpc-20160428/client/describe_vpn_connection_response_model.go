// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnConnectionResponseBody) *DescribeVpnConnectionResponse
	GetBody() *DescribeVpnConnectionResponseBody
}

type DescribeVpnConnectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnConnectionResponse) GetBody() *DescribeVpnConnectionResponseBody {
	return s.Body
}

func (s *DescribeVpnConnectionResponse) SetHeaders(v map[string]*string) *DescribeVpnConnectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnConnectionResponse) SetStatusCode(v int32) *DescribeVpnConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnConnectionResponse) SetBody(v *DescribeVpnConnectionResponseBody) *DescribeVpnConnectionResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
