// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTunnelInfo interface {
	dara.Model
	String() string
	GoString() string
	SetTags(v string) *CreateTunnelInfo
	GetTags() *string
	SetTunnelQos(v *TunnelQos) *CreateTunnelInfo
	GetTunnelQos() *TunnelQos
}

type CreateTunnelInfo struct {
	// example:
	//
	// K1:V1,K2:V2
	Tags      *string    `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TunnelQos *TunnelQos `json:"TunnelQos,omitempty" xml:"TunnelQos,omitempty"`
}

func (s CreateTunnelInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTunnelInfo) GoString() string {
	return s.String()
}

func (s *CreateTunnelInfo) GetTags() *string {
	return s.Tags
}

func (s *CreateTunnelInfo) GetTunnelQos() *TunnelQos {
	return s.TunnelQos
}

func (s *CreateTunnelInfo) SetTags(v string) *CreateTunnelInfo {
	s.Tags = &v
	return s
}

func (s *CreateTunnelInfo) SetTunnelQos(v *TunnelQos) *CreateTunnelInfo {
	s.TunnelQos = v
	return s
}

func (s *CreateTunnelInfo) Validate() error {
	if s.TunnelQos != nil {
		if err := s.TunnelQos.Validate(); err != nil {
			return err
		}
	}
	return nil
}
