// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTunnelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTunnelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTunnelResponse
	GetStatusCode() *int32
}

type UpdateTunnelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateTunnelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelResponse) GoString() string {
	return s.String()
}

func (s *UpdateTunnelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTunnelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTunnelResponse) SetHeaders(v map[string]*string) *UpdateTunnelResponse {
	s.Headers = v
	return s
}

func (s *UpdateTunnelResponse) SetStatusCode(v int32) *UpdateTunnelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTunnelResponse) Validate() error {
	return dara.Validate(s)
}
