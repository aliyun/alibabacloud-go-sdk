// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTunnelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportTunnelList(v *ListTunnelResp) *ListTunnelResponseBody
	GetImportTunnelList() *ListTunnelResp
}

type ListTunnelResponseBody struct {
	// The details of the tunnels.
	ImportTunnelList *ListTunnelResp `json:"ImportTunnelList,omitempty" xml:"ImportTunnelList,omitempty"`
}

func (s ListTunnelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelResponseBody) GoString() string {
	return s.String()
}

func (s *ListTunnelResponseBody) GetImportTunnelList() *ListTunnelResp {
	return s.ImportTunnelList
}

func (s *ListTunnelResponseBody) SetImportTunnelList(v *ListTunnelResp) *ListTunnelResponseBody {
	s.ImportTunnelList = v
	return s
}

func (s *ListTunnelResponseBody) Validate() error {
	if s.ImportTunnelList != nil {
		if err := s.ImportTunnelList.Validate(); err != nil {
			return err
		}
	}
	return nil
}
