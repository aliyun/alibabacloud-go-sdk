// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTunnelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportTunnel(v *GetTunnelResp) *GetTunnelResponseBody
	GetImportTunnel() *GetTunnelResp
}

type GetTunnelResponseBody struct {
	// The details for obtaining the details of the tunnel.
	ImportTunnel *GetTunnelResp `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty"`
}

func (s GetTunnelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTunnelResponseBody) GoString() string {
	return s.String()
}

func (s *GetTunnelResponseBody) GetImportTunnel() *GetTunnelResp {
	return s.ImportTunnel
}

func (s *GetTunnelResponseBody) SetImportTunnel(v *GetTunnelResp) *GetTunnelResponseBody {
	s.ImportTunnel = v
	return s
}

func (s *GetTunnelResponseBody) Validate() error {
	if s.ImportTunnel != nil {
		if err := s.ImportTunnel.Validate(); err != nil {
			return err
		}
	}
	return nil
}
