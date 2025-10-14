// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnsSaleControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *UpdateEnsSaleControlRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *UpdateEnsSaleControlRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *UpdateEnsSaleControlRequest
	GetCustomAccount() *string
	SetSaleControls(v []*UpdateEnsSaleControlRequestSaleControls) *UpdateEnsSaleControlRequest
	GetSaleControls() []*UpdateEnsSaleControlRequestSaleControls
}

type UpdateEnsSaleControlRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControls []*UpdateEnsSaleControlRequestSaleControls `json:"SaleControls,omitempty" xml:"SaleControls,omitempty" type:"Repeated"`
}

func (s UpdateEnsSaleControlRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnsSaleControlRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnsSaleControlRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *UpdateEnsSaleControlRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *UpdateEnsSaleControlRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *UpdateEnsSaleControlRequest) GetSaleControls() []*UpdateEnsSaleControlRequestSaleControls {
	return s.SaleControls
}

func (s *UpdateEnsSaleControlRequest) SetAliUidAccount(v string) *UpdateEnsSaleControlRequest {
	s.AliUidAccount = &v
	return s
}

func (s *UpdateEnsSaleControlRequest) SetCommodityCode(v string) *UpdateEnsSaleControlRequest {
	s.CommodityCode = &v
	return s
}

func (s *UpdateEnsSaleControlRequest) SetCustomAccount(v string) *UpdateEnsSaleControlRequest {
	s.CustomAccount = &v
	return s
}

func (s *UpdateEnsSaleControlRequest) SetSaleControls(v []*UpdateEnsSaleControlRequestSaleControls) *UpdateEnsSaleControlRequest {
	s.SaleControls = v
	return s
}

func (s *UpdateEnsSaleControlRequest) Validate() error {
	if s.SaleControls != nil {
		for _, item := range s.SaleControls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEnsSaleControlRequestSaleControls struct {
	ConditionControls []*UpdateEnsSaleControlRequestSaleControlsConditionControls `json:"ConditionControls,omitempty" xml:"ConditionControls,omitempty" type:"Repeated"`
	Description       *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// This parameter is required.
	ModuleValue *UpdateEnsSaleControlRequestSaleControlsModuleValue `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Struct"`
	// This parameter is required.
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s UpdateEnsSaleControlRequestSaleControls) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnsSaleControlRequestSaleControls) GoString() string {
	return s.String()
}

func (s *UpdateEnsSaleControlRequestSaleControls) GetConditionControls() []*UpdateEnsSaleControlRequestSaleControlsConditionControls {
	return s.ConditionControls
}

func (s *UpdateEnsSaleControlRequestSaleControls) GetDescription() *string {
	return s.Description
}

func (s *UpdateEnsSaleControlRequestSaleControls) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *UpdateEnsSaleControlRequestSaleControls) GetModuleValue() *UpdateEnsSaleControlRequestSaleControlsModuleValue {
	return s.ModuleValue
}

func (s *UpdateEnsSaleControlRequestSaleControls) GetOperator() *string {
	return s.Operator
}

func (s *UpdateEnsSaleControlRequestSaleControls) GetOrderType() *string {
	return s.OrderType
}

func (s *UpdateEnsSaleControlRequestSaleControls) SetConditionControls(v []*UpdateEnsSaleControlRequestSaleControlsConditionControls) *UpdateEnsSaleControlRequestSaleControls {
	s.ConditionControls = v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControls) SetDescription(v string) *UpdateEnsSaleControlRequestSaleControls {
	s.Description = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControls) SetModuleCode(v string) *UpdateEnsSaleControlRequestSaleControls {
	s.ModuleCode = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControls) SetModuleValue(v *UpdateEnsSaleControlRequestSaleControlsModuleValue) *UpdateEnsSaleControlRequestSaleControls {
	s.ModuleValue = v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControls) SetOperator(v string) *UpdateEnsSaleControlRequestSaleControls {
	s.Operator = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControls) SetOrderType(v string) *UpdateEnsSaleControlRequestSaleControls {
	s.OrderType = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControls) Validate() error {
	if s.ConditionControls != nil {
		for _, item := range s.ConditionControls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModuleValue != nil {
		if err := s.ModuleValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEnsSaleControlRequestSaleControlsConditionControls struct {
	ConditionControlModuleCode  *string `json:"ConditionControlModuleCode,omitempty" xml:"ConditionControlModuleCode,omitempty"`
	ConditionControlModuleValue *string `json:"ConditionControlModuleValue,omitempty" xml:"ConditionControlModuleValue,omitempty"`
}

func (s UpdateEnsSaleControlRequestSaleControlsConditionControls) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnsSaleControlRequestSaleControlsConditionControls) GoString() string {
	return s.String()
}

func (s *UpdateEnsSaleControlRequestSaleControlsConditionControls) GetConditionControlModuleCode() *string {
	return s.ConditionControlModuleCode
}

func (s *UpdateEnsSaleControlRequestSaleControlsConditionControls) GetConditionControlModuleValue() *string {
	return s.ConditionControlModuleValue
}

func (s *UpdateEnsSaleControlRequestSaleControlsConditionControls) SetConditionControlModuleCode(v string) *UpdateEnsSaleControlRequestSaleControlsConditionControls {
	s.ConditionControlModuleCode = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControlsConditionControls) SetConditionControlModuleValue(v string) *UpdateEnsSaleControlRequestSaleControlsConditionControls {
	s.ConditionControlModuleValue = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControlsConditionControls) Validate() error {
	return dara.Validate(s)
}

type UpdateEnsSaleControlRequestSaleControlsModuleValue struct {
	ModuleMaxValue *string   `json:"ModuleMaxValue,omitempty" xml:"ModuleMaxValue,omitempty"`
	ModuleMinValue *string   `json:"ModuleMinValue,omitempty" xml:"ModuleMinValue,omitempty"`
	ModuleValue    []*string `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Repeated"`
}

func (s UpdateEnsSaleControlRequestSaleControlsModuleValue) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnsSaleControlRequestSaleControlsModuleValue) GoString() string {
	return s.String()
}

func (s *UpdateEnsSaleControlRequestSaleControlsModuleValue) GetModuleMaxValue() *string {
	return s.ModuleMaxValue
}

func (s *UpdateEnsSaleControlRequestSaleControlsModuleValue) GetModuleMinValue() *string {
	return s.ModuleMinValue
}

func (s *UpdateEnsSaleControlRequestSaleControlsModuleValue) GetModuleValue() []*string {
	return s.ModuleValue
}

func (s *UpdateEnsSaleControlRequestSaleControlsModuleValue) SetModuleMaxValue(v string) *UpdateEnsSaleControlRequestSaleControlsModuleValue {
	s.ModuleMaxValue = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControlsModuleValue) SetModuleMinValue(v string) *UpdateEnsSaleControlRequestSaleControlsModuleValue {
	s.ModuleMinValue = &v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControlsModuleValue) SetModuleValue(v []*string) *UpdateEnsSaleControlRequestSaleControlsModuleValue {
	s.ModuleValue = v
	return s
}

func (s *UpdateEnsSaleControlRequestSaleControlsModuleValue) Validate() error {
	return dara.Validate(s)
}
