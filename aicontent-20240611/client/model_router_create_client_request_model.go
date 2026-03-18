// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ModelRouterCreateClientRequest
	GetAddress() *string
	SetAllowedModels(v string) *ModelRouterCreateClientRequest
	GetAllowedModels() *string
	SetContact(v string) *ModelRouterCreateClientRequest
	GetContact() *string
	SetName(v string) *ModelRouterCreateClientRequest
	GetName() *string
	SetRemark(v string) *ModelRouterCreateClientRequest
	GetRemark() *string
}

type ModelRouterCreateClientRequest struct {
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
}

func (s ModelRouterCreateClientRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateClientRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateClientRequest) GetAddress() *string {
	return s.Address
}

func (s *ModelRouterCreateClientRequest) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ModelRouterCreateClientRequest) GetContact() *string {
	return s.Contact
}

func (s *ModelRouterCreateClientRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterCreateClientRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModelRouterCreateClientRequest) SetAddress(v string) *ModelRouterCreateClientRequest {
	s.Address = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetAllowedModels(v string) *ModelRouterCreateClientRequest {
	s.AllowedModels = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetContact(v string) *ModelRouterCreateClientRequest {
	s.Contact = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetName(v string) *ModelRouterCreateClientRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetRemark(v string) *ModelRouterCreateClientRequest {
	s.Remark = &v
	return s
}

func (s *ModelRouterCreateClientRequest) Validate() error {
	return dara.Validate(s)
}
