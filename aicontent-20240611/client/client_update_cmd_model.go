// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ClientUpdateCmd
	GetAddress() *string
	SetAllowedModels(v string) *ClientUpdateCmd
	GetAllowedModels() *string
	SetContact(v string) *ClientUpdateCmd
	GetContact() *string
	SetName(v string) *ClientUpdateCmd
	GetName() *string
	SetRemark(v string) *ClientUpdateCmd
	GetRemark() *string
	SetStatus(v int32) *ClientUpdateCmd
	GetStatus() *int32
}

type ClientUpdateCmd struct {
	// example:
	//
	// 杭州市
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 1,2,3
	AllowedModels *string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty"`
	// example:
	//
	// 13800138000
	Contact *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// example:
	//
	// 我的客户
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ClientUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ClientUpdateCmd) GoString() string {
	return s.String()
}

func (s *ClientUpdateCmd) GetAddress() *string {
	return s.Address
}

func (s *ClientUpdateCmd) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ClientUpdateCmd) GetContact() *string {
	return s.Contact
}

func (s *ClientUpdateCmd) GetName() *string {
	return s.Name
}

func (s *ClientUpdateCmd) GetRemark() *string {
	return s.Remark
}

func (s *ClientUpdateCmd) GetStatus() *int32 {
	return s.Status
}

func (s *ClientUpdateCmd) SetAddress(v string) *ClientUpdateCmd {
	s.Address = &v
	return s
}

func (s *ClientUpdateCmd) SetAllowedModels(v string) *ClientUpdateCmd {
	s.AllowedModels = &v
	return s
}

func (s *ClientUpdateCmd) SetContact(v string) *ClientUpdateCmd {
	s.Contact = &v
	return s
}

func (s *ClientUpdateCmd) SetName(v string) *ClientUpdateCmd {
	s.Name = &v
	return s
}

func (s *ClientUpdateCmd) SetRemark(v string) *ClientUpdateCmd {
	s.Remark = &v
	return s
}

func (s *ClientUpdateCmd) SetStatus(v int32) *ClientUpdateCmd {
	s.Status = &v
	return s
}

func (s *ClientUpdateCmd) Validate() error {
	return dara.Validate(s)
}
