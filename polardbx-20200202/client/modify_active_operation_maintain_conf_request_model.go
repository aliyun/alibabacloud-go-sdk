// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintainConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycleTime(v string) *ModifyActiveOperationMaintainConfRequest
	GetCycleTime() *string
	SetCycleType(v string) *ModifyActiveOperationMaintainConfRequest
	GetCycleType() *string
	SetMaintainEndTime(v string) *ModifyActiveOperationMaintainConfRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyActiveOperationMaintainConfRequest
	GetMaintainStartTime() *string
	SetRegionId(v string) *ModifyActiveOperationMaintainConfRequest
	GetRegionId() *string
	SetStatus(v int32) *ModifyActiveOperationMaintainConfRequest
	GetStatus() *int32
}

type ModifyActiveOperationMaintainConfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3,4,5,6,7
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 02:00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyActiveOperationMaintainConfRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfRequest) GetCycleTime() *string {
	return s.CycleTime
}

func (s *ModifyActiveOperationMaintainConfRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *ModifyActiveOperationMaintainConfRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyActiveOperationMaintainConfRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyActiveOperationMaintainConfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyActiveOperationMaintainConfRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyActiveOperationMaintainConfRequest) SetCycleTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.CycleTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetCycleType(v string) *ModifyActiveOperationMaintainConfRequest {
	s.CycleType = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetMaintainEndTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetMaintainStartTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetRegionId(v string) *ModifyActiveOperationMaintainConfRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetStatus(v int32) *ModifyActiveOperationMaintainConfRequest {
	s.Status = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) Validate() error {
	return dara.Validate(s)
}
