// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMeteringDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreate(v string) *PushMeteringDataRequest
	GetGmtCreate() *string
	SetMeteringData(v []*PushMeteringDataRequestMeteringData) *PushMeteringDataRequest
	GetMeteringData() []*PushMeteringDataRequestMeteringData
}

type PushMeteringDataRequest struct {
	// example:
	//
	// 2023-01-11 10:31:00
	GmtCreate    *string                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MeteringData []*PushMeteringDataRequestMeteringData `json:"MeteringData,omitempty" xml:"MeteringData,omitempty" type:"Repeated"`
}

func (s PushMeteringDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PushMeteringDataRequest) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequest) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PushMeteringDataRequest) GetMeteringData() []*PushMeteringDataRequestMeteringData {
	return s.MeteringData
}

func (s *PushMeteringDataRequest) SetGmtCreate(v string) *PushMeteringDataRequest {
	s.GmtCreate = &v
	return s
}

func (s *PushMeteringDataRequest) SetMeteringData(v []*PushMeteringDataRequestMeteringData) *PushMeteringDataRequest {
	s.MeteringData = v
	return s
}

func (s *PushMeteringDataRequest) Validate() error {
	return dara.Validate(s)
}

type PushMeteringDataRequestMeteringData struct {
	// example:
	//
	// 1666854480406
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// gtm-cn-20p314k5h05
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test001
	MeteringAssist *string `json:"MeteringAssist,omitempty" xml:"MeteringAssist,omitempty"`
	// example:
	//
	// {"VirtualCpu":10}
	MeteringEntity *string `json:"MeteringEntity,omitempty" xml:"MeteringEntity,omitempty"`
	// example:
	//
	// ORDER20231001
	PushOrderBizId *string `json:"PushOrderBizId,omitempty" xml:"PushOrderBizId,omitempty"`
	// example:
	//
	// 1662284820000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PushMeteringDataRequestMeteringData) String() string {
	return dara.Prettify(s)
}

func (s PushMeteringDataRequestMeteringData) GoString() string {
	return s.String()
}

func (s *PushMeteringDataRequestMeteringData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *PushMeteringDataRequestMeteringData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PushMeteringDataRequestMeteringData) GetMeteringAssist() *string {
	return s.MeteringAssist
}

func (s *PushMeteringDataRequestMeteringData) GetMeteringEntity() *string {
	return s.MeteringEntity
}

func (s *PushMeteringDataRequestMeteringData) GetPushOrderBizId() *string {
	return s.PushOrderBizId
}

func (s *PushMeteringDataRequestMeteringData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *PushMeteringDataRequestMeteringData) SetEndTime(v int64) *PushMeteringDataRequestMeteringData {
	s.EndTime = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetInstanceId(v string) *PushMeteringDataRequestMeteringData {
	s.InstanceId = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetMeteringAssist(v string) *PushMeteringDataRequestMeteringData {
	s.MeteringAssist = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetMeteringEntity(v string) *PushMeteringDataRequestMeteringData {
	s.MeteringEntity = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetPushOrderBizId(v string) *PushMeteringDataRequestMeteringData {
	s.PushOrderBizId = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) SetStartTime(v int64) *PushMeteringDataRequestMeteringData {
	s.StartTime = &v
	return s
}

func (s *PushMeteringDataRequestMeteringData) Validate() error {
	return dara.Validate(s)
}
