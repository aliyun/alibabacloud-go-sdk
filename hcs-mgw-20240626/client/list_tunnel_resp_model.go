// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTunnelResp interface {
	dara.Model
	String() string
	GoString() string
	SetImportTunnel(v []*GetTunnelResp) *ListTunnelResp
	GetImportTunnel() []*GetTunnelResp
	SetNextMarker(v string) *ListTunnelResp
	GetNextMarker() *string
	SetTruncated(v bool) *ListTunnelResp
	GetTruncated() *bool
}

type ListTunnelResp struct {
	ImportTunnel []*GetTunnelResp `json:"ImportTunnel,omitempty" xml:"ImportTunnel,omitempty" type:"Repeated"`
	NextMarker   *string          `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Truncated    *bool            `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListTunnelResp) String() string {
	return dara.Prettify(s)
}

func (s ListTunnelResp) GoString() string {
	return s.String()
}

func (s *ListTunnelResp) GetImportTunnel() []*GetTunnelResp {
	return s.ImportTunnel
}

func (s *ListTunnelResp) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListTunnelResp) GetTruncated() *bool {
	return s.Truncated
}

func (s *ListTunnelResp) SetImportTunnel(v []*GetTunnelResp) *ListTunnelResp {
	s.ImportTunnel = v
	return s
}

func (s *ListTunnelResp) SetNextMarker(v string) *ListTunnelResp {
	s.NextMarker = &v
	return s
}

func (s *ListTunnelResp) SetTruncated(v bool) *ListTunnelResp {
	s.Truncated = &v
	return s
}

func (s *ListTunnelResp) Validate() error {
	if s.ImportTunnel != nil {
		for _, item := range s.ImportTunnel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
