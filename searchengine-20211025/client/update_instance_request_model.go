// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v []*UpdateInstanceRequestComponents) *UpdateInstanceRequest
	GetComponents() []*UpdateInstanceRequestComponents
	SetDescription(v string) *UpdateInstanceRequest
	GetDescription() *string
	SetOrderType(v string) *UpdateInstanceRequest
	GetOrderType() *string
}

type UpdateInstanceRequest struct {
	// The information about the instance specification.
	Components []*UpdateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The description of the instance.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the order. Valid values: UPGRADE and DOWNGRADE. UPGRADE upgrades the instance specifications. DOWNGRADE: downgrades the instance specifications.
	//
	// example:
	//
	// ""
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetComponents() []*UpdateInstanceRequestComponents {
	return s.Components
}

func (s *UpdateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *UpdateInstanceRequest) SetComponents(v []*UpdateInstanceRequestComponents) *UpdateInstanceRequest {
	s.Components = v
	return s
}

func (s *UpdateInstanceRequest) SetDescription(v string) *UpdateInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceRequest) SetOrderType(v string) *UpdateInstanceRequest {
	s.OrderType = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceRequestComponents struct {
	// The code of the specification, which must be consistent with the value that you specify on the buy page.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The value of the specification.
	//
	// example:
	//
	// ""
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateInstanceRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestComponents) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceRequestComponents) GetValue() *string {
	return s.Value
}

func (s *UpdateInstanceRequestComponents) SetCode(v string) *UpdateInstanceRequestComponents {
	s.Code = &v
	return s
}

func (s *UpdateInstanceRequestComponents) SetValue(v string) *UpdateInstanceRequestComponents {
	s.Value = &v
	return s
}

func (s *UpdateInstanceRequestComponents) Validate() error {
	return dara.Validate(s)
}
