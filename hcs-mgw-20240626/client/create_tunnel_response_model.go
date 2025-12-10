// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTunnelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTunnelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTunnelResponse
	GetStatusCode() *int32
}

type CreateTunnelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateTunnelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTunnelResponse) GoString() string {
	return s.String()
}

func (s *CreateTunnelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTunnelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTunnelResponse) SetHeaders(v map[string]*string) *CreateTunnelResponse {
	s.Headers = v
	return s
}

func (s *CreateTunnelResponse) SetStatusCode(v int32) *CreateTunnelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTunnelResponse) Validate() error {
	return dara.Validate(s)
}
