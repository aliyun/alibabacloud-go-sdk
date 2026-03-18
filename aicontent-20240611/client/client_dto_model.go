// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ClientDTO
	GetAddress() *string
	SetAllowedModels(v string) *ClientDTO
	GetAllowedModels() *string
	SetClientUuid(v string) *ClientDTO
	GetClientUuid() *string
	SetContact(v string) *ClientDTO
	GetContact() *string
	SetDeleteTag(v int32) *ClientDTO
	GetDeleteTag() *int32
	SetGmtCreate(v string) *ClientDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ClientDTO
	GetGmtModified() *string
	SetId(v int64) *ClientDTO
	GetId() *int64
	SetMain(v int32) *ClientDTO
	GetMain() *int32
	SetName(v string) *ClientDTO
	GetName() *string
	SetRemark(v string) *ClientDTO
	GetRemark() *string
}

type ClientDTO struct {
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
	// st-xxxx
	ClientUuid *string `json:"clientUuid,omitempty" xml:"clientUuid,omitempty"`
	// example:
	//
	// 13800138000
	Contact *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	Main *int32 `json:"main,omitempty" xml:"main,omitempty"`
	// example:
	//
	// 我的客户
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s ClientDTO) String() string {
	return dara.Prettify(s)
}

func (s ClientDTO) GoString() string {
	return s.String()
}

func (s *ClientDTO) GetAddress() *string {
	return s.Address
}

func (s *ClientDTO) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ClientDTO) GetClientUuid() *string {
	return s.ClientUuid
}

func (s *ClientDTO) GetContact() *string {
	return s.Contact
}

func (s *ClientDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ClientDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ClientDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ClientDTO) GetId() *int64 {
	return s.Id
}

func (s *ClientDTO) GetMain() *int32 {
	return s.Main
}

func (s *ClientDTO) GetName() *string {
	return s.Name
}

func (s *ClientDTO) GetRemark() *string {
	return s.Remark
}

func (s *ClientDTO) SetAddress(v string) *ClientDTO {
	s.Address = &v
	return s
}

func (s *ClientDTO) SetAllowedModels(v string) *ClientDTO {
	s.AllowedModels = &v
	return s
}

func (s *ClientDTO) SetClientUuid(v string) *ClientDTO {
	s.ClientUuid = &v
	return s
}

func (s *ClientDTO) SetContact(v string) *ClientDTO {
	s.Contact = &v
	return s
}

func (s *ClientDTO) SetDeleteTag(v int32) *ClientDTO {
	s.DeleteTag = &v
	return s
}

func (s *ClientDTO) SetGmtCreate(v string) *ClientDTO {
	s.GmtCreate = &v
	return s
}

func (s *ClientDTO) SetGmtModified(v string) *ClientDTO {
	s.GmtModified = &v
	return s
}

func (s *ClientDTO) SetId(v int64) *ClientDTO {
	s.Id = &v
	return s
}

func (s *ClientDTO) SetMain(v int32) *ClientDTO {
	s.Main = &v
	return s
}

func (s *ClientDTO) SetName(v string) *ClientDTO {
	s.Name = &v
	return s
}

func (s *ClientDTO) SetRemark(v string) *ClientDTO {
	s.Remark = &v
	return s
}

func (s *ClientDTO) Validate() error {
	return dara.Validate(s)
}
