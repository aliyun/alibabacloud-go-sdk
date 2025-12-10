// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTunnelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImportTunnel(v *UpdateTunnelInfo) *UpdateTunnelRequest
	GetImportTunnel() *UpdateTunnelInfo
}

type UpdateTunnelRequest struct {
	// The details for updating the tunnel.
	ImportTunnel *UpdateTunnelInfo `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty"`
}

func (s UpdateTunnelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelRequest) GoString() string {
	return s.String()
}

func (s *UpdateTunnelRequest) GetImportTunnel() *UpdateTunnelInfo {
	return s.ImportTunnel
}

func (s *UpdateTunnelRequest) SetImportTunnel(v *UpdateTunnelInfo) *UpdateTunnelRequest {
	s.ImportTunnel = v
	return s
}

func (s *UpdateTunnelRequest) Validate() error {
	if s.ImportTunnel != nil {
		if err := s.ImportTunnel.Validate(); err != nil {
			return err
		}
	}
	return nil
}
