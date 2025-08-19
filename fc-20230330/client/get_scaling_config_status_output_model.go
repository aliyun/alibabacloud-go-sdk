// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScalingConfigStatusOutput interface {
	dara.Model
	String() string
	GoString() string
	SetScalingConfigStatus(v *ScalingConfigStatus) *GetScalingConfigStatusOutput
	GetScalingConfigStatus() *ScalingConfigStatus
}

type GetScalingConfigStatusOutput struct {
	ScalingConfigStatus *ScalingConfigStatus `json:"scalingConfigStatus,omitempty" xml:"scalingConfigStatus,omitempty"`
}

func (s GetScalingConfigStatusOutput) String() string {
	return dara.Prettify(s)
}

func (s GetScalingConfigStatusOutput) GoString() string {
	return s.String()
}

func (s *GetScalingConfigStatusOutput) GetScalingConfigStatus() *ScalingConfigStatus {
	return s.ScalingConfigStatus
}

func (s *GetScalingConfigStatusOutput) SetScalingConfigStatus(v *ScalingConfigStatus) *GetScalingConfigStatusOutput {
	s.ScalingConfigStatus = v
	return s
}

func (s *GetScalingConfigStatusOutput) Validate() error {
	return dara.Validate(s)
}
