// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoRenewInstanceParam interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *AutoRenewInstanceParam
	GetAutoRenew() *string
	SetAutoRenewDuration(v string) *AutoRenewInstanceParam
	GetAutoRenewDuration() *string
	SetAutoRenewDurationUnit(v string) *AutoRenewInstanceParam
	GetAutoRenewDurationUnit() *string
	SetInstanceId(v string) *AutoRenewInstanceParam
	GetInstanceId() *string
}

type AutoRenewInstanceParam struct {
	AutoRenew             *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewDuration     *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AutoRenewInstanceParam) String() string {
	return dara.Prettify(s)
}

func (s AutoRenewInstanceParam) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceParam) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *AutoRenewInstanceParam) GetAutoRenewDuration() *string {
	return s.AutoRenewDuration
}

func (s *AutoRenewInstanceParam) GetAutoRenewDurationUnit() *string {
	return s.AutoRenewDurationUnit
}

func (s *AutoRenewInstanceParam) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AutoRenewInstanceParam) SetAutoRenew(v string) *AutoRenewInstanceParam {
	s.AutoRenew = &v
	return s
}

func (s *AutoRenewInstanceParam) SetAutoRenewDuration(v string) *AutoRenewInstanceParam {
	s.AutoRenewDuration = &v
	return s
}

func (s *AutoRenewInstanceParam) SetAutoRenewDurationUnit(v string) *AutoRenewInstanceParam {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *AutoRenewInstanceParam) SetInstanceId(v string) *AutoRenewInstanceParam {
	s.InstanceId = &v
	return s
}

func (s *AutoRenewInstanceParam) Validate() error {
	return dara.Validate(s)
}
