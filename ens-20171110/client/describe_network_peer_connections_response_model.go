// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkPeerConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkPeerConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkPeerConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkPeerConnectionsResponseBody) *DescribeNetworkPeerConnectionsResponse
	GetBody() *DescribeNetworkPeerConnectionsResponseBody
}

type DescribeNetworkPeerConnectionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkPeerConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkPeerConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPeerConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPeerConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkPeerConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkPeerConnectionsResponse) GetBody() *DescribeNetworkPeerConnectionsResponseBody {
	return s.Body
}

func (s *DescribeNetworkPeerConnectionsResponse) SetHeaders(v map[string]*string) *DescribeNetworkPeerConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponse) SetStatusCode(v int32) *DescribeNetworkPeerConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponse) SetBody(v *DescribeNetworkPeerConnectionsResponseBody) *DescribeNetworkPeerConnectionsResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkPeerConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
