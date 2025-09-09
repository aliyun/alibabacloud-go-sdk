// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcPeerConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcPeerConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcPeerConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcPeerConnectionsResponseBody) *ListVpcPeerConnectionsResponse
	GetBody() *ListVpcPeerConnectionsResponseBody
}

type ListVpcPeerConnectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcPeerConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcPeerConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcPeerConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcPeerConnectionsResponse) GetBody() *ListVpcPeerConnectionsResponseBody {
	return s.Body
}

func (s *ListVpcPeerConnectionsResponse) SetHeaders(v map[string]*string) *ListVpcPeerConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcPeerConnectionsResponse) SetStatusCode(v int32) *ListVpcPeerConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcPeerConnectionsResponse) SetBody(v *ListVpcPeerConnectionsResponseBody) *ListVpcPeerConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListVpcPeerConnectionsResponse) Validate() error {
	return dara.Validate(s)
}
