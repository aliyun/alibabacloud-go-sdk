// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingActivityResult interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ScalingActivityResult
	GetInstanceId() *string
}

type ScalingActivityResult struct {
	// 实例ID。
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ScalingActivityResult) String() string {
	return dara.Prettify(s)
}

func (s ScalingActivityResult) GoString() string {
	return s.String()
}

func (s *ScalingActivityResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ScalingActivityResult) SetInstanceId(v string) *ScalingActivityResult {
	s.InstanceId = &v
	return s
}

func (s *ScalingActivityResult) Validate() error {
	return dara.Validate(s)
}
