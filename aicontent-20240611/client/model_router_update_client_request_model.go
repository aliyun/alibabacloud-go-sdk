// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ModelRouterUpdateClientRequest
	GetAddress() *string
	SetAllowedModels(v string) *ModelRouterUpdateClientRequest
	GetAllowedModels() *string
	SetContact(v string) *ModelRouterUpdateClientRequest
	GetContact() *string
	SetName(v string) *ModelRouterUpdateClientRequest
	GetName() *string
	SetRemark(v string) *ModelRouterUpdateClientRequest
	GetRemark() *string
	SetStatus(v int32) *ModelRouterUpdateClientRequest
	GetStatus() *int32
}

type ModelRouterUpdateClientRequest struct {
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

func (s ModelRouterUpdateClientRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateClientRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateClientRequest) GetAddress() *string {
	return s.Address
}

func (s *ModelRouterUpdateClientRequest) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ModelRouterUpdateClientRequest) GetContact() *string {
	return s.Contact
}

func (s *ModelRouterUpdateClientRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterUpdateClientRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModelRouterUpdateClientRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModelRouterUpdateClientRequest) SetAddress(v string) *ModelRouterUpdateClientRequest {
	s.Address = &v
	return s
}

func (s *ModelRouterUpdateClientRequest) SetAllowedModels(v string) *ModelRouterUpdateClientRequest {
	s.AllowedModels = &v
	return s
}

func (s *ModelRouterUpdateClientRequest) SetContact(v string) *ModelRouterUpdateClientRequest {
	s.Contact = &v
	return s
}

func (s *ModelRouterUpdateClientRequest) SetName(v string) *ModelRouterUpdateClientRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterUpdateClientRequest) SetRemark(v string) *ModelRouterUpdateClientRequest {
	s.Remark = &v
	return s
}

func (s *ModelRouterUpdateClientRequest) SetStatus(v int32) *ModelRouterUpdateClientRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterUpdateClientRequest) Validate() error {
	return dara.Validate(s)
}
