// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingActivityResultDTO interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ScalingActivityResultDTO
	GetInstanceId() *string
}

type ScalingActivityResultDTO struct {
	// 实例ID。
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ScalingActivityResultDTO) String() string {
	return dara.Prettify(s)
}

func (s ScalingActivityResultDTO) GoString() string {
	return s.String()
}

func (s *ScalingActivityResultDTO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ScalingActivityResultDTO) SetInstanceId(v string) *ScalingActivityResultDTO {
	s.InstanceId = &v
	return s
}

func (s *ScalingActivityResultDTO) Validate() error {
	return dara.Validate(s)
}
