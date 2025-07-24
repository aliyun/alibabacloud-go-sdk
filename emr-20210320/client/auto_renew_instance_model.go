// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoRenewInstance interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *AutoRenewInstance
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *AutoRenewInstance
	GetAutoRenewDuration() *int32
	SetAutoRenewDurationUnit(v string) *AutoRenewInstance
	GetAutoRenewDurationUnit() *string
	SetEmrAutoRenewDuration(v int32) *AutoRenewInstance
	GetEmrAutoRenewDuration() *int32
	SetEmrAutoRenewDurationUnit(v string) *AutoRenewInstance
	GetEmrAutoRenewDurationUnit() *string
	SetInstanceId(v string) *AutoRenewInstance
	GetInstanceId() *string
}

type AutoRenewInstance struct {
	// 自动续费。
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// 自动续费时长。
	//
	// example:
	//
	// 12
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// 自动付费时长单位。
	//
	// example:
	//
	// Month
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	// emr实例自动续费时长。
	//
	// example:
	//
	// 12
	EmrAutoRenewDuration *int32 `json:"EmrAutoRenewDuration,omitempty" xml:"EmrAutoRenewDuration,omitempty"`
	// emr实例自动续费时长单位。
	//
	// example:
	//
	// Month
	EmrAutoRenewDurationUnit *string `json:"EmrAutoRenewDurationUnit,omitempty" xml:"EmrAutoRenewDurationUnit,omitempty"`
	// 节点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1cudc25w2bfwl5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AutoRenewInstance) String() string {
	return dara.Prettify(s)
}

func (s AutoRenewInstance) GoString() string {
	return s.String()
}

func (s *AutoRenewInstance) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *AutoRenewInstance) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *AutoRenewInstance) GetAutoRenewDurationUnit() *string {
	return s.AutoRenewDurationUnit
}

func (s *AutoRenewInstance) GetEmrAutoRenewDuration() *int32 {
	return s.EmrAutoRenewDuration
}

func (s *AutoRenewInstance) GetEmrAutoRenewDurationUnit() *string {
	return s.EmrAutoRenewDurationUnit
}

func (s *AutoRenewInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AutoRenewInstance) SetAutoRenew(v bool) *AutoRenewInstance {
	s.AutoRenew = &v
	return s
}

func (s *AutoRenewInstance) SetAutoRenewDuration(v int32) *AutoRenewInstance {
	s.AutoRenewDuration = &v
	return s
}

func (s *AutoRenewInstance) SetAutoRenewDurationUnit(v string) *AutoRenewInstance {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *AutoRenewInstance) SetEmrAutoRenewDuration(v int32) *AutoRenewInstance {
	s.EmrAutoRenewDuration = &v
	return s
}

func (s *AutoRenewInstance) SetEmrAutoRenewDurationUnit(v string) *AutoRenewInstance {
	s.EmrAutoRenewDurationUnit = &v
	return s
}

func (s *AutoRenewInstance) SetInstanceId(v string) *AutoRenewInstance {
	s.InstanceId = &v
	return s
}

func (s *AutoRenewInstance) Validate() error {
	return dara.Validate(s)
}
