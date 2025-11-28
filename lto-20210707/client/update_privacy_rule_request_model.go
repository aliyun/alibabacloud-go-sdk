// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivacyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgImpl(v string) *UpdatePrivacyRuleRequest
	GetAlgImpl() *string
	SetAlgType(v string) *UpdatePrivacyRuleRequest
	GetAlgType() *string
	SetName(v string) *UpdatePrivacyRuleRequest
	GetName() *string
	SetPrivacyRuleId(v string) *UpdatePrivacyRuleRequest
	GetPrivacyRuleId() *string
	SetRegionId(v string) *UpdatePrivacyRuleRequest
	GetRegionId() *string
	SetRemark(v string) *UpdatePrivacyRuleRequest
	GetRemark() *string
}

type UpdatePrivacyRuleRequest struct {
	AlgImpl *string `json:"AlgImpl,omitempty" xml:"AlgImpl,omitempty"`
	// This parameter is required.
	AlgType *string `json:"AlgType,omitempty" xml:"AlgType,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdatePrivacyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivacyRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivacyRuleRequest) GetAlgImpl() *string {
	return s.AlgImpl
}

func (s *UpdatePrivacyRuleRequest) GetAlgType() *string {
	return s.AlgType
}

func (s *UpdatePrivacyRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePrivacyRuleRequest) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *UpdatePrivacyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrivacyRuleRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdatePrivacyRuleRequest) SetAlgImpl(v string) *UpdatePrivacyRuleRequest {
	s.AlgImpl = &v
	return s
}

func (s *UpdatePrivacyRuleRequest) SetAlgType(v string) *UpdatePrivacyRuleRequest {
	s.AlgType = &v
	return s
}

func (s *UpdatePrivacyRuleRequest) SetName(v string) *UpdatePrivacyRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdatePrivacyRuleRequest) SetPrivacyRuleId(v string) *UpdatePrivacyRuleRequest {
	s.PrivacyRuleId = &v
	return s
}

func (s *UpdatePrivacyRuleRequest) SetRegionId(v string) *UpdatePrivacyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrivacyRuleRequest) SetRemark(v string) *UpdatePrivacyRuleRequest {
	s.Remark = &v
	return s
}

func (s *UpdatePrivacyRuleRequest) Validate() error {
	return dara.Validate(s)
}
