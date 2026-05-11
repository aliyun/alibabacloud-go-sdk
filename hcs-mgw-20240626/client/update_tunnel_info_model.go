// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTunnelInfo interface {
	dara.Model
	String() string
	GoString() string
	SetTags(v string) *UpdateTunnelInfo
	GetTags() *string
	SetTunnelQos(v *TunnelQos) *UpdateTunnelInfo
	GetTunnelQos() *TunnelQos
}

type UpdateTunnelInfo struct {
	// The tags.
	//
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The throttling settings of the tunnel.
	TunnelQos *TunnelQos `json:"TunnelQos,omitempty" xml:"TunnelQos,omitempty"`
}

func (s UpdateTunnelInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTunnelInfo) GoString() string {
	return s.String()
}

func (s *UpdateTunnelInfo) GetTags() *string {
	return s.Tags
}

func (s *UpdateTunnelInfo) GetTunnelQos() *TunnelQos {
	return s.TunnelQos
}

func (s *UpdateTunnelInfo) SetTags(v string) *UpdateTunnelInfo {
	s.Tags = &v
	return s
}

func (s *UpdateTunnelInfo) SetTunnelQos(v *TunnelQos) *UpdateTunnelInfo {
	s.TunnelQos = v
	return s
}

func (s *UpdateTunnelInfo) Validate() error {
	if s.TunnelQos != nil {
		if err := s.TunnelQos.Validate(); err != nil {
			return err
		}
	}
	return nil
}
