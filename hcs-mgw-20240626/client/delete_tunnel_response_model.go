// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTunnelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTunnelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTunnelResponse
	GetStatusCode() *int32
}

type DeleteTunnelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteTunnelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTunnelResponse) GoString() string {
	return s.String()
}

func (s *DeleteTunnelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTunnelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTunnelResponse) SetHeaders(v map[string]*string) *DeleteTunnelResponse {
	s.Headers = v
	return s
}

func (s *DeleteTunnelResponse) SetStatusCode(v int32) *DeleteTunnelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTunnelResponse) Validate() error {
	return dara.Validate(s)
}
