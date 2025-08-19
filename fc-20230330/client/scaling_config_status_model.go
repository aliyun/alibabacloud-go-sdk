// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingConfigStatus interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ScalingConfigStatus
	GetFunctionName() *string
	SetQualifier(v string) *ScalingConfigStatus
	GetQualifier() *string
	SetResidentConfig(v *ResidentConfig) *ScalingConfigStatus
	GetResidentConfig() *ResidentConfig
	SetResourceType(v string) *ScalingConfigStatus
	GetResourceType() *string
	SetScalingStatus(v *ScalingStatus) *ScalingConfigStatus
	GetScalingStatus() *ScalingStatus
}

type ScalingConfigStatus struct {
	FunctionName   *string         `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Qualifier      *string         `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	ResidentConfig *ResidentConfig `json:"residentConfig,omitempty" xml:"residentConfig,omitempty"`
	ResourceType   *string         `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	ScalingStatus  *ScalingStatus  `json:"scalingStatus,omitempty" xml:"scalingStatus,omitempty"`
}

func (s ScalingConfigStatus) String() string {
	return dara.Prettify(s)
}

func (s ScalingConfigStatus) GoString() string {
	return s.String()
}

func (s *ScalingConfigStatus) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ScalingConfigStatus) GetQualifier() *string {
	return s.Qualifier
}

func (s *ScalingConfigStatus) GetResidentConfig() *ResidentConfig {
	return s.ResidentConfig
}

func (s *ScalingConfigStatus) GetResourceType() *string {
	return s.ResourceType
}

func (s *ScalingConfigStatus) GetScalingStatus() *ScalingStatus {
	return s.ScalingStatus
}

func (s *ScalingConfigStatus) SetFunctionName(v string) *ScalingConfigStatus {
	s.FunctionName = &v
	return s
}

func (s *ScalingConfigStatus) SetQualifier(v string) *ScalingConfigStatus {
	s.Qualifier = &v
	return s
}

func (s *ScalingConfigStatus) SetResidentConfig(v *ResidentConfig) *ScalingConfigStatus {
	s.ResidentConfig = v
	return s
}

func (s *ScalingConfigStatus) SetResourceType(v string) *ScalingConfigStatus {
	s.ResourceType = &v
	return s
}

func (s *ScalingConfigStatus) SetScalingStatus(v *ScalingStatus) *ScalingConfigStatus {
	s.ScalingStatus = v
	return s
}

func (s *ScalingConfigStatus) Validate() error {
	return dara.Validate(s)
}
