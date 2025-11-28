// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivacyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgImpl(v string) *AddPrivacyRuleRequest
	GetAlgImpl() *string
	SetAlgType(v string) *AddPrivacyRuleRequest
	GetAlgType() *string
	SetName(v string) *AddPrivacyRuleRequest
	GetName() *string
	SetRegionId(v string) *AddPrivacyRuleRequest
	GetRegionId() *string
	SetRemark(v string) *AddPrivacyRuleRequest
	GetRemark() *string
}

type AddPrivacyRuleRequest struct {
	AlgImpl *string `json:"AlgImpl,omitempty" xml:"AlgImpl,omitempty"`
	// This parameter is required.
	AlgType *string `json:"AlgType,omitempty" xml:"AlgType,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddPrivacyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPrivacyRuleRequest) GoString() string {
	return s.String()
}

func (s *AddPrivacyRuleRequest) GetAlgImpl() *string {
	return s.AlgImpl
}

func (s *AddPrivacyRuleRequest) GetAlgType() *string {
	return s.AlgType
}

func (s *AddPrivacyRuleRequest) GetName() *string {
	return s.Name
}

func (s *AddPrivacyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPrivacyRuleRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddPrivacyRuleRequest) SetAlgImpl(v string) *AddPrivacyRuleRequest {
	s.AlgImpl = &v
	return s
}

func (s *AddPrivacyRuleRequest) SetAlgType(v string) *AddPrivacyRuleRequest {
	s.AlgType = &v
	return s
}

func (s *AddPrivacyRuleRequest) SetName(v string) *AddPrivacyRuleRequest {
	s.Name = &v
	return s
}

func (s *AddPrivacyRuleRequest) SetRegionId(v string) *AddPrivacyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AddPrivacyRuleRequest) SetRemark(v string) *AddPrivacyRuleRequest {
	s.Remark = &v
	return s
}

func (s *AddPrivacyRuleRequest) Validate() error {
	return dara.Validate(s)
}
