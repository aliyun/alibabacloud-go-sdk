// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *WorkspaceSpec
	GetCode() *string
	SetCodeType(v string) *WorkspaceSpec
	GetCodeType() *string
	SetIsGuaranteedValid(v bool) *WorkspaceSpec
	GetIsGuaranteedValid() *bool
	SetIsOverSoldValid(v bool) *WorkspaceSpec
	GetIsOverSoldValid() *bool
	SetReason(v string) *WorkspaceSpec
	GetReason() *string
	SetSpec(v *ResourceAmount) *WorkspaceSpec
	GetSpec() *ResourceAmount
	SetSpecName(v string) *WorkspaceSpec
	GetSpecName() *string
}

type WorkspaceSpec struct {
	Code              *string         `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeType          *string         `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	IsGuaranteedValid *bool           `json:"IsGuaranteedValid,omitempty" xml:"IsGuaranteedValid,omitempty"`
	IsOverSoldValid   *bool           `json:"IsOverSoldValid,omitempty" xml:"IsOverSoldValid,omitempty"`
	Reason            *string         `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Spec              *ResourceAmount `json:"Spec,omitempty" xml:"Spec,omitempty"`
	SpecName          *string         `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s WorkspaceSpec) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceSpec) GoString() string {
	return s.String()
}

func (s *WorkspaceSpec) GetCode() *string {
	return s.Code
}

func (s *WorkspaceSpec) GetCodeType() *string {
	return s.CodeType
}

func (s *WorkspaceSpec) GetIsGuaranteedValid() *bool {
	return s.IsGuaranteedValid
}

func (s *WorkspaceSpec) GetIsOverSoldValid() *bool {
	return s.IsOverSoldValid
}

func (s *WorkspaceSpec) GetReason() *string {
	return s.Reason
}

func (s *WorkspaceSpec) GetSpec() *ResourceAmount {
	return s.Spec
}

func (s *WorkspaceSpec) GetSpecName() *string {
	return s.SpecName
}

func (s *WorkspaceSpec) SetCode(v string) *WorkspaceSpec {
	s.Code = &v
	return s
}

func (s *WorkspaceSpec) SetCodeType(v string) *WorkspaceSpec {
	s.CodeType = &v
	return s
}

func (s *WorkspaceSpec) SetIsGuaranteedValid(v bool) *WorkspaceSpec {
	s.IsGuaranteedValid = &v
	return s
}

func (s *WorkspaceSpec) SetIsOverSoldValid(v bool) *WorkspaceSpec {
	s.IsOverSoldValid = &v
	return s
}

func (s *WorkspaceSpec) SetReason(v string) *WorkspaceSpec {
	s.Reason = &v
	return s
}

func (s *WorkspaceSpec) SetSpec(v *ResourceAmount) *WorkspaceSpec {
	s.Spec = v
	return s
}

func (s *WorkspaceSpec) SetSpecName(v string) *WorkspaceSpec {
	s.SpecName = &v
	return s
}

func (s *WorkspaceSpec) Validate() error {
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
