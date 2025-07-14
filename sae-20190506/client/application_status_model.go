// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplicationStatus interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCount(v int64) *ApplicationStatus
	GetInstanceCount() *int64
	SetScaleConfig(v *ScaleConfig) *ApplicationStatus
	GetScaleConfig() *ScaleConfig
}

type ApplicationStatus struct {
	InstanceCount *int64       `json:"instanceCount,omitempty" xml:"instanceCount,omitempty"`
	ScaleConfig   *ScaleConfig `json:"scaleConfig,omitempty" xml:"scaleConfig,omitempty"`
}

func (s ApplicationStatus) String() string {
	return dara.Prettify(s)
}

func (s ApplicationStatus) GoString() string {
	return s.String()
}

func (s *ApplicationStatus) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ApplicationStatus) GetScaleConfig() *ScaleConfig {
	return s.ScaleConfig
}

func (s *ApplicationStatus) SetInstanceCount(v int64) *ApplicationStatus {
	s.InstanceCount = &v
	return s
}

func (s *ApplicationStatus) SetScaleConfig(v *ScaleConfig) *ApplicationStatus {
	s.ScaleConfig = v
	return s
}

func (s *ApplicationStatus) Validate() error {
	return dara.Validate(s)
}
