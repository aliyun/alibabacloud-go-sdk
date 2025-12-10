// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTunnelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportTunnel(v *CreateTunnelInfo) *CreateTunnelRequest
	GetImportTunnel() *CreateTunnelInfo
}

type CreateTunnelRequest struct {
	// The details for creating the tunnel.
	ImportTunnel *CreateTunnelInfo `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty"`
}

func (s CreateTunnelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTunnelRequest) GoString() string {
	return s.String()
}

func (s *CreateTunnelRequest) GetImportTunnel() *CreateTunnelInfo {
	return s.ImportTunnel
}

func (s *CreateTunnelRequest) SetImportTunnel(v *CreateTunnelInfo) *CreateTunnelRequest {
	s.ImportTunnel = v
	return s
}

func (s *CreateTunnelRequest) Validate() error {
	if s.ImportTunnel != nil {
		if err := s.ImportTunnel.Validate(); err != nil {
			return err
		}
	}
	return nil
}
