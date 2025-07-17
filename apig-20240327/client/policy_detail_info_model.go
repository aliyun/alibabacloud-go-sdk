// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolicyDetailInfo interface {
	dara.Model
	String() string
	GoString() string
	SetClassId(v string) *PolicyDetailInfo
	GetClassId() *string
	SetClassName(v string) *PolicyDetailInfo
	GetClassName() *string
	SetConfig(v string) *PolicyDetailInfo
	GetConfig() *string
	SetDescription(v string) *PolicyDetailInfo
	GetDescription() *string
	SetName(v string) *PolicyDetailInfo
	GetName() *string
	SetPolicyId(v string) *PolicyDetailInfo
	GetPolicyId() *string
}

type PolicyDetailInfo struct {
	ClassId     *string `json:"classId,omitempty" xml:"classId,omitempty"`
	ClassName   *string `json:"className,omitempty" xml:"className,omitempty"`
	Config      *string `json:"config,omitempty" xml:"config,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	PolicyId    *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s PolicyDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s PolicyDetailInfo) GoString() string {
	return s.String()
}

func (s *PolicyDetailInfo) GetClassId() *string {
	return s.ClassId
}

func (s *PolicyDetailInfo) GetClassName() *string {
	return s.ClassName
}

func (s *PolicyDetailInfo) GetConfig() *string {
	return s.Config
}

func (s *PolicyDetailInfo) GetDescription() *string {
	return s.Description
}

func (s *PolicyDetailInfo) GetName() *string {
	return s.Name
}

func (s *PolicyDetailInfo) GetPolicyId() *string {
	return s.PolicyId
}

func (s *PolicyDetailInfo) SetClassId(v string) *PolicyDetailInfo {
	s.ClassId = &v
	return s
}

func (s *PolicyDetailInfo) SetClassName(v string) *PolicyDetailInfo {
	s.ClassName = &v
	return s
}

func (s *PolicyDetailInfo) SetConfig(v string) *PolicyDetailInfo {
	s.Config = &v
	return s
}

func (s *PolicyDetailInfo) SetDescription(v string) *PolicyDetailInfo {
	s.Description = &v
	return s
}

func (s *PolicyDetailInfo) SetName(v string) *PolicyDetailInfo {
	s.Name = &v
	return s
}

func (s *PolicyDetailInfo) SetPolicyId(v string) *PolicyDetailInfo {
	s.PolicyId = &v
	return s
}

func (s *PolicyDetailInfo) Validate() error {
	return dara.Validate(s)
}
