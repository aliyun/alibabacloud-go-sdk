// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientTreeDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ClientTreeDTO
	GetAddress() *string
	SetAllowedModels(v string) *ClientTreeDTO
	GetAllowedModels() *string
	SetBalance(v *ClientBalanceDTO) *ClientTreeDTO
	GetBalance() *ClientBalanceDTO
	SetChildren(v []*ClientTreeDTO) *ClientTreeDTO
	GetChildren() []*ClientTreeDTO
	SetClientUuid(v string) *ClientTreeDTO
	GetClientUuid() *string
	SetContact(v string) *ClientTreeDTO
	GetContact() *string
	SetDeleteTag(v int32) *ClientTreeDTO
	GetDeleteTag() *int32
	SetDiscount(v float64) *ClientTreeDTO
	GetDiscount() *float64
	SetGmtCreate(v string) *ClientTreeDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ClientTreeDTO
	GetGmtModified() *string
	SetId(v int64) *ClientTreeDTO
	GetId() *int64
	SetLevel(v int32) *ClientTreeDTO
	GetLevel() *int32
	SetMain(v int32) *ClientTreeDTO
	GetMain() *int32
	SetName(v string) *ClientTreeDTO
	GetName() *string
	SetParentId(v int64) *ClientTreeDTO
	GetParentId() *int64
	SetRemark(v string) *ClientTreeDTO
	GetRemark() *string
}

type ClientTreeDTO struct {
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
	// null
	Balance *ClientBalanceDTO `json:"balance,omitempty" xml:"balance,omitempty"`
	// example:
	//
	// []
	Children []*ClientTreeDTO `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
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
	// 1.0
	Discount *float64 `json:"discount,omitempty" xml:"discount,omitempty"`
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
	// 1
	Level *int32 `json:"level,omitempty" xml:"level,omitempty"`
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
	// null
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s ClientTreeDTO) String() string {
	return dara.Prettify(s)
}

func (s ClientTreeDTO) GoString() string {
	return s.String()
}

func (s *ClientTreeDTO) GetAddress() *string {
	return s.Address
}

func (s *ClientTreeDTO) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ClientTreeDTO) GetBalance() *ClientBalanceDTO {
	return s.Balance
}

func (s *ClientTreeDTO) GetChildren() []*ClientTreeDTO {
	return s.Children
}

func (s *ClientTreeDTO) GetClientUuid() *string {
	return s.ClientUuid
}

func (s *ClientTreeDTO) GetContact() *string {
	return s.Contact
}

func (s *ClientTreeDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ClientTreeDTO) GetDiscount() *float64 {
	return s.Discount
}

func (s *ClientTreeDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ClientTreeDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ClientTreeDTO) GetId() *int64 {
	return s.Id
}

func (s *ClientTreeDTO) GetLevel() *int32 {
	return s.Level
}

func (s *ClientTreeDTO) GetMain() *int32 {
	return s.Main
}

func (s *ClientTreeDTO) GetName() *string {
	return s.Name
}

func (s *ClientTreeDTO) GetParentId() *int64 {
	return s.ParentId
}

func (s *ClientTreeDTO) GetRemark() *string {
	return s.Remark
}

func (s *ClientTreeDTO) SetAddress(v string) *ClientTreeDTO {
	s.Address = &v
	return s
}

func (s *ClientTreeDTO) SetAllowedModels(v string) *ClientTreeDTO {
	s.AllowedModels = &v
	return s
}

func (s *ClientTreeDTO) SetBalance(v *ClientBalanceDTO) *ClientTreeDTO {
	s.Balance = v
	return s
}

func (s *ClientTreeDTO) SetChildren(v []*ClientTreeDTO) *ClientTreeDTO {
	s.Children = v
	return s
}

func (s *ClientTreeDTO) SetClientUuid(v string) *ClientTreeDTO {
	s.ClientUuid = &v
	return s
}

func (s *ClientTreeDTO) SetContact(v string) *ClientTreeDTO {
	s.Contact = &v
	return s
}

func (s *ClientTreeDTO) SetDeleteTag(v int32) *ClientTreeDTO {
	s.DeleteTag = &v
	return s
}

func (s *ClientTreeDTO) SetDiscount(v float64) *ClientTreeDTO {
	s.Discount = &v
	return s
}

func (s *ClientTreeDTO) SetGmtCreate(v string) *ClientTreeDTO {
	s.GmtCreate = &v
	return s
}

func (s *ClientTreeDTO) SetGmtModified(v string) *ClientTreeDTO {
	s.GmtModified = &v
	return s
}

func (s *ClientTreeDTO) SetId(v int64) *ClientTreeDTO {
	s.Id = &v
	return s
}

func (s *ClientTreeDTO) SetLevel(v int32) *ClientTreeDTO {
	s.Level = &v
	return s
}

func (s *ClientTreeDTO) SetMain(v int32) *ClientTreeDTO {
	s.Main = &v
	return s
}

func (s *ClientTreeDTO) SetName(v string) *ClientTreeDTO {
	s.Name = &v
	return s
}

func (s *ClientTreeDTO) SetParentId(v int64) *ClientTreeDTO {
	s.ParentId = &v
	return s
}

func (s *ClientTreeDTO) SetRemark(v string) *ClientTreeDTO {
	s.Remark = &v
	return s
}

func (s *ClientTreeDTO) Validate() error {
	if s.Balance != nil {
		if err := s.Balance.Validate(); err != nil {
			return err
		}
	}
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
