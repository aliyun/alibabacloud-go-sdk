// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRouteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *AddRouteRuleRequest
	GetBizChainId() *string
	SetChainUpMode(v string) *AddRouteRuleRequest
	GetChainUpMode() *string
	SetContractName(v string) *AddRouteRuleRequest
	GetContractName() *string
	SetContractTemplateId(v string) *AddRouteRuleRequest
	GetContractTemplateId() *string
	SetDeviceGroupId(v string) *AddRouteRuleRequest
	GetDeviceGroupId() *string
	SetInvokeType(v string) *AddRouteRuleRequest
	GetInvokeType() *string
	SetPrivacyRuleId(v string) *AddRouteRuleRequest
	GetPrivacyRuleId() *string
	SetRegionId(v string) *AddRouteRuleRequest
	GetRegionId() *string
	SetRemark(v string) *AddRouteRuleRequest
	GetRemark() *string
}

type AddRouteRuleRequest struct {
	// This parameter is required.
	BizChainId         *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	ChainUpMode        *string `json:"ChainUpMode,omitempty" xml:"ChainUpMode,omitempty"`
	ContractName       *string `json:"ContractName,omitempty" xml:"ContractName,omitempty"`
	ContractTemplateId *string `json:"ContractTemplateId,omitempty" xml:"ContractTemplateId,omitempty"`
	DeviceGroupId      *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// This parameter is required.
	InvokeType    *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddRouteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *AddRouteRuleRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *AddRouteRuleRequest) GetChainUpMode() *string {
	return s.ChainUpMode
}

func (s *AddRouteRuleRequest) GetContractName() *string {
	return s.ContractName
}

func (s *AddRouteRuleRequest) GetContractTemplateId() *string {
	return s.ContractTemplateId
}

func (s *AddRouteRuleRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *AddRouteRuleRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *AddRouteRuleRequest) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *AddRouteRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddRouteRuleRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddRouteRuleRequest) SetBizChainId(v string) *AddRouteRuleRequest {
	s.BizChainId = &v
	return s
}

func (s *AddRouteRuleRequest) SetChainUpMode(v string) *AddRouteRuleRequest {
	s.ChainUpMode = &v
	return s
}

func (s *AddRouteRuleRequest) SetContractName(v string) *AddRouteRuleRequest {
	s.ContractName = &v
	return s
}

func (s *AddRouteRuleRequest) SetContractTemplateId(v string) *AddRouteRuleRequest {
	s.ContractTemplateId = &v
	return s
}

func (s *AddRouteRuleRequest) SetDeviceGroupId(v string) *AddRouteRuleRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddRouteRuleRequest) SetInvokeType(v string) *AddRouteRuleRequest {
	s.InvokeType = &v
	return s
}

func (s *AddRouteRuleRequest) SetPrivacyRuleId(v string) *AddRouteRuleRequest {
	s.PrivacyRuleId = &v
	return s
}

func (s *AddRouteRuleRequest) SetRegionId(v string) *AddRouteRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AddRouteRuleRequest) SetRemark(v string) *AddRouteRuleRequest {
	s.Remark = &v
	return s
}

func (s *AddRouteRuleRequest) Validate() error {
	return dara.Validate(s)
}
