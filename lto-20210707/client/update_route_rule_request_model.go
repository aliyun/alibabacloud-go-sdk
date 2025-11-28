// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *UpdateRouteRuleRequest
	GetBizChainId() *string
	SetContractName(v string) *UpdateRouteRuleRequest
	GetContractName() *string
	SetContractTemplateId(v string) *UpdateRouteRuleRequest
	GetContractTemplateId() *string
	SetInvokeType(v string) *UpdateRouteRuleRequest
	GetInvokeType() *string
	SetPrivacyRuleId(v string) *UpdateRouteRuleRequest
	GetPrivacyRuleId() *string
	SetRegionId(v string) *UpdateRouteRuleRequest
	GetRegionId() *string
	SetRemark(v string) *UpdateRouteRuleRequest
	GetRemark() *string
	SetRouteRuleId(v string) *UpdateRouteRuleRequest
	GetRouteRuleId() *string
}

type UpdateRouteRuleRequest struct {
	// This parameter is required.
	BizChainId         *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	ContractName       *string `json:"ContractName,omitempty" xml:"ContractName,omitempty"`
	ContractTemplateId *string `json:"ContractTemplateId,omitempty" xml:"ContractTemplateId,omitempty"`
	// This parameter is required.
	InvokeType    *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	RouteRuleId *string `json:"RouteRuleId,omitempty" xml:"RouteRuleId,omitempty"`
}

func (s UpdateRouteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteRuleRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *UpdateRouteRuleRequest) GetContractName() *string {
	return s.ContractName
}

func (s *UpdateRouteRuleRequest) GetContractTemplateId() *string {
	return s.ContractTemplateId
}

func (s *UpdateRouteRuleRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *UpdateRouteRuleRequest) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *UpdateRouteRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateRouteRuleRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateRouteRuleRequest) GetRouteRuleId() *string {
	return s.RouteRuleId
}

func (s *UpdateRouteRuleRequest) SetBizChainId(v string) *UpdateRouteRuleRequest {
	s.BizChainId = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetContractName(v string) *UpdateRouteRuleRequest {
	s.ContractName = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetContractTemplateId(v string) *UpdateRouteRuleRequest {
	s.ContractTemplateId = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetInvokeType(v string) *UpdateRouteRuleRequest {
	s.InvokeType = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetPrivacyRuleId(v string) *UpdateRouteRuleRequest {
	s.PrivacyRuleId = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetRegionId(v string) *UpdateRouteRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetRemark(v string) *UpdateRouteRuleRequest {
	s.Remark = &v
	return s
}

func (s *UpdateRouteRuleRequest) SetRouteRuleId(v string) *UpdateRouteRuleRequest {
	s.RouteRuleId = &v
	return s
}

func (s *UpdateRouteRuleRequest) Validate() error {
	return dara.Validate(s)
}
