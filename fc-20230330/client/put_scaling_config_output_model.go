// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutScalingConfigOutput interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *PutScalingConfigOutput
	GetFunctionName() *string
	SetQualifier(v string) *PutScalingConfigOutput
	GetQualifier() *string
	SetResidentConfig(v *ResidentConfig) *PutScalingConfigOutput
	GetResidentConfig() *ResidentConfig
	SetResourceType(v string) *PutScalingConfigOutput
	GetResourceType() *string
}

type PutScalingConfigOutput struct {
	FunctionName   *string         `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Qualifier      *string         `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	ResidentConfig *ResidentConfig `json:"residentConfig,omitempty" xml:"residentConfig,omitempty"`
	ResourceType   *string         `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s PutScalingConfigOutput) String() string {
	return dara.Prettify(s)
}

func (s PutScalingConfigOutput) GoString() string {
	return s.String()
}

func (s *PutScalingConfigOutput) GetFunctionName() *string {
	return s.FunctionName
}

func (s *PutScalingConfigOutput) GetQualifier() *string {
	return s.Qualifier
}

func (s *PutScalingConfigOutput) GetResidentConfig() *ResidentConfig {
	return s.ResidentConfig
}

func (s *PutScalingConfigOutput) GetResourceType() *string {
	return s.ResourceType
}

func (s *PutScalingConfigOutput) SetFunctionName(v string) *PutScalingConfigOutput {
	s.FunctionName = &v
	return s
}

func (s *PutScalingConfigOutput) SetQualifier(v string) *PutScalingConfigOutput {
	s.Qualifier = &v
	return s
}

func (s *PutScalingConfigOutput) SetResidentConfig(v *ResidentConfig) *PutScalingConfigOutput {
	s.ResidentConfig = v
	return s
}

func (s *PutScalingConfigOutput) SetResourceType(v string) *PutScalingConfigOutput {
	s.ResourceType = &v
	return s
}

func (s *PutScalingConfigOutput) Validate() error {
	return dara.Validate(s)
}
