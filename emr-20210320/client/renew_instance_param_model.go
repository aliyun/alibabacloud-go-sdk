// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceParam interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RenewInstanceParam
	GetInstanceId() *string
	SetRenewDuration(v int64) *RenewInstanceParam
	GetRenewDuration() *int64
	SetRenewDurationUnit(v string) *RenewInstanceParam
	GetRenewDurationUnit() *string
}

type RenewInstanceParam struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RenewDuration     *int64  `json:"RenewDuration,omitempty" xml:"RenewDuration,omitempty"`
	RenewDurationUnit *string `json:"RenewDurationUnit,omitempty" xml:"RenewDurationUnit,omitempty"`
}

func (s RenewInstanceParam) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceParam) GoString() string {
	return s.String()
}

func (s *RenewInstanceParam) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewInstanceParam) GetRenewDuration() *int64 {
	return s.RenewDuration
}

func (s *RenewInstanceParam) GetRenewDurationUnit() *string {
	return s.RenewDurationUnit
}

func (s *RenewInstanceParam) SetInstanceId(v string) *RenewInstanceParam {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceParam) SetRenewDuration(v int64) *RenewInstanceParam {
	s.RenewDuration = &v
	return s
}

func (s *RenewInstanceParam) SetRenewDurationUnit(v string) *RenewInstanceParam {
	s.RenewDurationUnit = &v
	return s
}

func (s *RenewInstanceParam) Validate() error {
	return dara.Validate(s)
}
