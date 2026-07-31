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
	SetAllowedModelGroupConfig(v string) *ModelRouterCreateClientRequest
	GetAllowedModelGroupConfig() *string
	SetAllowedModels(v string) *ModelRouterCreateClientRequest
	GetAllowedModels() *string
	SetContact(v string) *ModelRouterCreateClientRequest
	GetContact() *string
	SetDiscount(v float64) *ModelRouterCreateClientRequest
	GetDiscount() *float64
	SetName(v string) *ModelRouterCreateClientRequest
	GetName() *string
	SetParentId(v int64) *ModelRouterCreateClientRequest
	GetParentId() *int64
	SetRemark(v string) *ModelRouterCreateClientRequest
	GetRemark() *string
}

type ModelRouterCreateClientRequest struct {
	// The company address.
	//
	// example:
	//
	// Hangzhou
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The allowed model group configuration in JSON string format: {"model_ids":[101],"group_ids":["mg_xxx"]}. If both this field and allowedModels are specified, this field takes precedence.
	//
	// example:
	//
	// {"model_ids":[101],"group_ids":["mg_xxx"]}
	AllowedModelGroupConfig *string `json:"allowedModelGroupConfig,omitempty" xml:"allowedModelGroupConfig,omitempty"`
	// The list of allowed model IDs, separated by commas. An empty value indicates all models are allowed.
	//
	// example:
	//
	// 1,2,3
	AllowedModels *string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty"`
	// The contact information.
	//
	// example:
	//
	// 13800138000
	Contact *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// The discount coefficient. A value of 1.0 indicates no discount, and 0.8 indicates a 20% discount. Default value: 1.0.
	//
	// example:
	//
	// 1.0
	Discount *float64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// The customer name.
	//
	// example:
	//
	// MyCustomer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the parent department. If not specified, a top-level department is created.
	//
	// example:
	//
	// 292090
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Remarks
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

func (s *ModelRouterCreateClientRequest) GetAllowedModelGroupConfig() *string {
	return s.AllowedModelGroupConfig
}

func (s *ModelRouterCreateClientRequest) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ModelRouterCreateClientRequest) GetContact() *string {
	return s.Contact
}

func (s *ModelRouterCreateClientRequest) GetDiscount() *float64 {
	return s.Discount
}

func (s *ModelRouterCreateClientRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterCreateClientRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *ModelRouterCreateClientRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModelRouterCreateClientRequest) SetAddress(v string) *ModelRouterCreateClientRequest {
	s.Address = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetAllowedModelGroupConfig(v string) *ModelRouterCreateClientRequest {
	s.AllowedModelGroupConfig = &v
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

func (s *ModelRouterCreateClientRequest) SetDiscount(v float64) *ModelRouterCreateClientRequest {
	s.Discount = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetName(v string) *ModelRouterCreateClientRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetParentId(v int64) *ModelRouterCreateClientRequest {
	s.ParentId = &v
	return s
}

func (s *ModelRouterCreateClientRequest) SetRemark(v string) *ModelRouterCreateClientRequest {
	s.Remark = &v
	return s
}

func (s *ModelRouterCreateClientRequest) Validate() error {
	return dara.Validate(s)
}
