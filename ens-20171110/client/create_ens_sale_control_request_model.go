// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsSaleControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *CreateEnsSaleControlRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *CreateEnsSaleControlRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *CreateEnsSaleControlRequest
	GetCustomAccount() *string
	SetSaleControls(v []*CreateEnsSaleControlRequestSaleControls) *CreateEnsSaleControlRequest
	GetSaleControls() []*CreateEnsSaleControlRequestSaleControls
}

type CreateEnsSaleControlRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControls []*CreateEnsSaleControlRequestSaleControls `json:"SaleControls,omitempty" xml:"SaleControls,omitempty" type:"Repeated"`
}

func (s CreateEnsSaleControlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsSaleControlRequest) GoString() string {
	return s.String()
}

func (s *CreateEnsSaleControlRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *CreateEnsSaleControlRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateEnsSaleControlRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *CreateEnsSaleControlRequest) GetSaleControls() []*CreateEnsSaleControlRequestSaleControls {
	return s.SaleControls
}

func (s *CreateEnsSaleControlRequest) SetAliUidAccount(v string) *CreateEnsSaleControlRequest {
	s.AliUidAccount = &v
	return s
}

func (s *CreateEnsSaleControlRequest) SetCommodityCode(v string) *CreateEnsSaleControlRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateEnsSaleControlRequest) SetCustomAccount(v string) *CreateEnsSaleControlRequest {
	s.CustomAccount = &v
	return s
}

func (s *CreateEnsSaleControlRequest) SetSaleControls(v []*CreateEnsSaleControlRequestSaleControls) *CreateEnsSaleControlRequest {
	s.SaleControls = v
	return s
}

func (s *CreateEnsSaleControlRequest) Validate() error {
	return dara.Validate(s)
}

type CreateEnsSaleControlRequestSaleControls struct {
	ConditionControls []*CreateEnsSaleControlRequestSaleControlsConditionControls `json:"ConditionControls,omitempty" xml:"ConditionControls,omitempty" type:"Repeated"`
	Description       *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// This parameter is required.
	ModuleValue *CreateEnsSaleControlRequestSaleControlsModuleValue `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Struct"`
	// This parameter is required.
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s CreateEnsSaleControlRequestSaleControls) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsSaleControlRequestSaleControls) GoString() string {
	return s.String()
}

func (s *CreateEnsSaleControlRequestSaleControls) GetConditionControls() []*CreateEnsSaleControlRequestSaleControlsConditionControls {
	return s.ConditionControls
}

func (s *CreateEnsSaleControlRequestSaleControls) GetDescription() *string {
	return s.Description
}

func (s *CreateEnsSaleControlRequestSaleControls) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *CreateEnsSaleControlRequestSaleControls) GetModuleValue() *CreateEnsSaleControlRequestSaleControlsModuleValue {
	return s.ModuleValue
}

func (s *CreateEnsSaleControlRequestSaleControls) GetOperator() *string {
	return s.Operator
}

func (s *CreateEnsSaleControlRequestSaleControls) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateEnsSaleControlRequestSaleControls) SetConditionControls(v []*CreateEnsSaleControlRequestSaleControlsConditionControls) *CreateEnsSaleControlRequestSaleControls {
	s.ConditionControls = v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControls) SetDescription(v string) *CreateEnsSaleControlRequestSaleControls {
	s.Description = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControls) SetModuleCode(v string) *CreateEnsSaleControlRequestSaleControls {
	s.ModuleCode = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControls) SetModuleValue(v *CreateEnsSaleControlRequestSaleControlsModuleValue) *CreateEnsSaleControlRequestSaleControls {
	s.ModuleValue = v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControls) SetOperator(v string) *CreateEnsSaleControlRequestSaleControls {
	s.Operator = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControls) SetOrderType(v string) *CreateEnsSaleControlRequestSaleControls {
	s.OrderType = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControls) Validate() error {
	return dara.Validate(s)
}

type CreateEnsSaleControlRequestSaleControlsConditionControls struct {
	ConditionControlModuleCode  *string `json:"ConditionControlModuleCode,omitempty" xml:"ConditionControlModuleCode,omitempty"`
	ConditionControlModuleValue *string `json:"ConditionControlModuleValue,omitempty" xml:"ConditionControlModuleValue,omitempty"`
}

func (s CreateEnsSaleControlRequestSaleControlsConditionControls) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsSaleControlRequestSaleControlsConditionControls) GoString() string {
	return s.String()
}

func (s *CreateEnsSaleControlRequestSaleControlsConditionControls) GetConditionControlModuleCode() *string {
	return s.ConditionControlModuleCode
}

func (s *CreateEnsSaleControlRequestSaleControlsConditionControls) GetConditionControlModuleValue() *string {
	return s.ConditionControlModuleValue
}

func (s *CreateEnsSaleControlRequestSaleControlsConditionControls) SetConditionControlModuleCode(v string) *CreateEnsSaleControlRequestSaleControlsConditionControls {
	s.ConditionControlModuleCode = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControlsConditionControls) SetConditionControlModuleValue(v string) *CreateEnsSaleControlRequestSaleControlsConditionControls {
	s.ConditionControlModuleValue = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControlsConditionControls) Validate() error {
	return dara.Validate(s)
}

type CreateEnsSaleControlRequestSaleControlsModuleValue struct {
	ModuleMaxValue *string   `json:"ModuleMaxValue,omitempty" xml:"ModuleMaxValue,omitempty"`
	ModuleMinValue *string   `json:"ModuleMinValue,omitempty" xml:"ModuleMinValue,omitempty"`
	ModuleValue    []*string `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Repeated"`
}

func (s CreateEnsSaleControlRequestSaleControlsModuleValue) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsSaleControlRequestSaleControlsModuleValue) GoString() string {
	return s.String()
}

func (s *CreateEnsSaleControlRequestSaleControlsModuleValue) GetModuleMaxValue() *string {
	return s.ModuleMaxValue
}

func (s *CreateEnsSaleControlRequestSaleControlsModuleValue) GetModuleMinValue() *string {
	return s.ModuleMinValue
}

func (s *CreateEnsSaleControlRequestSaleControlsModuleValue) GetModuleValue() []*string {
	return s.ModuleValue
}

func (s *CreateEnsSaleControlRequestSaleControlsModuleValue) SetModuleMaxValue(v string) *CreateEnsSaleControlRequestSaleControlsModuleValue {
	s.ModuleMaxValue = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControlsModuleValue) SetModuleMinValue(v string) *CreateEnsSaleControlRequestSaleControlsModuleValue {
	s.ModuleMinValue = &v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControlsModuleValue) SetModuleValue(v []*string) *CreateEnsSaleControlRequestSaleControlsModuleValue {
	s.ModuleValue = v
	return s
}

func (s *CreateEnsSaleControlRequestSaleControlsModuleValue) Validate() error {
	return dara.Validate(s)
}
