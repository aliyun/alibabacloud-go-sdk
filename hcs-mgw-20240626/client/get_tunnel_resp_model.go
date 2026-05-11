// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTunnelResp interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetTunnelResp
	GetCreateTime() *string
	SetModifyTime(v string) *GetTunnelResp
	GetModifyTime() *string
	SetOwner(v string) *GetTunnelResp
	GetOwner() *string
	SetTags(v string) *GetTunnelResp
	GetTags() *string
	SetTunnelId(v string) *GetTunnelResp
	GetTunnelId() *string
	SetTunnelQos(v *TunnelQos) *GetTunnelResp
	GetTunnelQos() *TunnelQos
}

type GetTunnelResp struct {
	// The time when the tunnel was created.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the tunnel was last modified.
	//
	// example:
	//
	// 2024-05-01 12:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The tunnel owner.
	//
	// example:
	//
	// test_owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The tags.
	//
	// example:
	//
	// key1:value1,key2:value2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The tunnel ID.
	//
	// example:
	//
	// test_tunnel_id
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The throttling settings of the tunnel.
	TunnelQos *TunnelQos `json:"TunnelQos,omitempty" xml:"TunnelQos,omitempty"`
}

func (s GetTunnelResp) String() string {
	return dara.Prettify(s)
}

func (s GetTunnelResp) GoString() string {
	return s.String()
}

func (s *GetTunnelResp) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTunnelResp) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetTunnelResp) GetOwner() *string {
	return s.Owner
}

func (s *GetTunnelResp) GetTags() *string {
	return s.Tags
}

func (s *GetTunnelResp) GetTunnelId() *string {
	return s.TunnelId
}

func (s *GetTunnelResp) GetTunnelQos() *TunnelQos {
	return s.TunnelQos
}

func (s *GetTunnelResp) SetCreateTime(v string) *GetTunnelResp {
	s.CreateTime = &v
	return s
}

func (s *GetTunnelResp) SetModifyTime(v string) *GetTunnelResp {
	s.ModifyTime = &v
	return s
}

func (s *GetTunnelResp) SetOwner(v string) *GetTunnelResp {
	s.Owner = &v
	return s
}

func (s *GetTunnelResp) SetTags(v string) *GetTunnelResp {
	s.Tags = &v
	return s
}

func (s *GetTunnelResp) SetTunnelId(v string) *GetTunnelResp {
	s.TunnelId = &v
	return s
}

func (s *GetTunnelResp) SetTunnelQos(v *TunnelQos) *GetTunnelResp {
	s.TunnelQos = v
	return s
}

func (s *GetTunnelResp) Validate() error {
	if s.TunnelQos != nil {
		if err := s.TunnelQos.Validate(); err != nil {
			return err
		}
	}
	return nil
}
