// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstance interface {
	dara.Model
	String() string
	GoString() string
	SetEmrRenewDuration(v int32) *RenewInstance
	GetEmrRenewDuration() *int32
	SetEmrRenewDurationUnit(v string) *RenewInstance
	GetEmrRenewDurationUnit() *string
	SetInstanceId(v string) *RenewInstance
	GetInstanceId() *string
	SetRenewDuration(v int32) *RenewInstance
	GetRenewDuration() *int32
	SetRenewDurationUnit(v string) *RenewInstance
	GetRenewDurationUnit() *string
}

type RenewInstance struct {
	// emr实例续费时长。
	//
	// example:
	//
	// 12
	EmrRenewDuration *int32 `json:"EmrRenewDuration,omitempty" xml:"EmrRenewDuration,omitempty"`
	// emr实例续费时长单位。
	//
	// example:
	//
	// Month
	EmrRenewDurationUnit *string `json:"EmrRenewDurationUnit,omitempty" xml:"EmrRenewDurationUnit,omitempty"`
	// 节点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 续费时长。
	//
	// example:
	//
	// 12
	RenewDuration *int32 `json:"RenewDuration,omitempty" xml:"RenewDuration,omitempty"`
	// 付费时长单位。
	//
	// example:
	//
	// Month
	RenewDurationUnit *string `json:"RenewDurationUnit,omitempty" xml:"RenewDurationUnit,omitempty"`
}

func (s RenewInstance) String() string {
	return dara.Prettify(s)
}

func (s RenewInstance) GoString() string {
	return s.String()
}

func (s *RenewInstance) GetEmrRenewDuration() *int32 {
	return s.EmrRenewDuration
}

func (s *RenewInstance) GetEmrRenewDurationUnit() *string {
	return s.EmrRenewDurationUnit
}

func (s *RenewInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstance) GetRenewDuration() *int32 {
	return s.RenewDuration
}

func (s *RenewInstance) GetRenewDurationUnit() *string {
	return s.RenewDurationUnit
}

func (s *RenewInstance) SetEmrRenewDuration(v int32) *RenewInstance {
	s.EmrRenewDuration = &v
	return s
}

func (s *RenewInstance) SetEmrRenewDurationUnit(v string) *RenewInstance {
	s.EmrRenewDurationUnit = &v
	return s
}

func (s *RenewInstance) SetInstanceId(v string) *RenewInstance {
	s.InstanceId = &v
	return s
}

func (s *RenewInstance) SetRenewDuration(v int32) *RenewInstance {
	s.RenewDuration = &v
	return s
}

func (s *RenewInstance) SetRenewDurationUnit(v string) *RenewInstance {
	s.RenewDurationUnit = &v
	return s
}

func (s *RenewInstance) Validate() error {
	return dara.Validate(s)
}
