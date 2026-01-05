// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceStatus interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstanceStatus
	GetCode() *string
	SetMessage(v string) *InstanceStatus
	GetMessage() *string
	SetPhase(v string) *InstanceStatus
	GetPhase() *string
	SetSlotNum(v int32) *InstanceStatus
	GetSlotNum() *int32
	SetUsedCapacity(v string) *InstanceStatus
	GetUsedCapacity() *string
}

type InstanceStatus struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Init Succeed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// 10
	SlotNum *int32 `json:"SlotNum,omitempty" xml:"SlotNum,omitempty"`
	// example:
	//
	// 20.0G
	UsedCapacity *string `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s InstanceStatus) String() string {
	return dara.Prettify(s)
}

func (s InstanceStatus) GoString() string {
	return s.String()
}

func (s *InstanceStatus) GetCode() *string {
	return s.Code
}

func (s *InstanceStatus) GetMessage() *string {
	return s.Message
}

func (s *InstanceStatus) GetPhase() *string {
	return s.Phase
}

func (s *InstanceStatus) GetSlotNum() *int32 {
	return s.SlotNum
}

func (s *InstanceStatus) GetUsedCapacity() *string {
	return s.UsedCapacity
}

func (s *InstanceStatus) SetCode(v string) *InstanceStatus {
	s.Code = &v
	return s
}

func (s *InstanceStatus) SetMessage(v string) *InstanceStatus {
	s.Message = &v
	return s
}

func (s *InstanceStatus) SetPhase(v string) *InstanceStatus {
	s.Phase = &v
	return s
}

func (s *InstanceStatus) SetSlotNum(v int32) *InstanceStatus {
	s.SlotNum = &v
	return s
}

func (s *InstanceStatus) SetUsedCapacity(v string) *InstanceStatus {
	s.UsedCapacity = &v
	return s
}

func (s *InstanceStatus) Validate() error {
	return dara.Validate(s)
}
