// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutScalingConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetResidentConfig(v *ResidentConfig) *PutScalingConfigInput
	GetResidentConfig() *ResidentConfig
	SetResourceType(v string) *PutScalingConfigInput
	GetResourceType() *string
}

type PutScalingConfigInput struct {
	ResidentConfig *ResidentConfig `json:"residentConfig,omitempty" xml:"residentConfig,omitempty"`
	ResourceType   *string         `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s PutScalingConfigInput) String() string {
	return dara.Prettify(s)
}

func (s PutScalingConfigInput) GoString() string {
	return s.String()
}

func (s *PutScalingConfigInput) GetResidentConfig() *ResidentConfig {
	return s.ResidentConfig
}

func (s *PutScalingConfigInput) GetResourceType() *string {
	return s.ResourceType
}

func (s *PutScalingConfigInput) SetResidentConfig(v *ResidentConfig) *PutScalingConfigInput {
	s.ResidentConfig = v
	return s
}

func (s *PutScalingConfigInput) SetResourceType(v string) *PutScalingConfigInput {
	s.ResourceType = &v
	return s
}

func (s *PutScalingConfigInput) Validate() error {
	return dara.Validate(s)
}
