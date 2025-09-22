// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSumRecordFlowPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *AddSumRecordFlowPopRequest
	GetActivityId() *string
	SetCode(v string) *AddSumRecordFlowPopRequest
	GetCode() *string
	SetConferenceName(v string) *AddSumRecordFlowPopRequest
	GetConferenceName() *string
	SetDeviceId(v string) *AddSumRecordFlowPopRequest
	GetDeviceId() *string
	SetEntryName(v string) *AddSumRecordFlowPopRequest
	GetEntryName() *string
	SetIdcard(v string) *AddSumRecordFlowPopRequest
	GetIdcard() *string
	SetSignTime(v string) *AddSumRecordFlowPopRequest
	GetSignTime() *string
	SetType(v int32) *AddSumRecordFlowPopRequest
	GetType() *int32
}

type AddSumRecordFlowPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 会议名称
	ConferenceName *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Z10
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 入口名称
	EntryName *string `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	// example:
	//
	// 429005111100000
	Idcard *string `json:"Idcard,omitempty" xml:"Idcard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-25 14:11
	SignTime *string `json:"SignTime,omitempty" xml:"SignTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddSumRecordFlowPopRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSumRecordFlowPopRequest) GoString() string {
	return s.String()
}

func (s *AddSumRecordFlowPopRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *AddSumRecordFlowPopRequest) GetCode() *string {
	return s.Code
}

func (s *AddSumRecordFlowPopRequest) GetConferenceName() *string {
	return s.ConferenceName
}

func (s *AddSumRecordFlowPopRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *AddSumRecordFlowPopRequest) GetEntryName() *string {
	return s.EntryName
}

func (s *AddSumRecordFlowPopRequest) GetIdcard() *string {
	return s.Idcard
}

func (s *AddSumRecordFlowPopRequest) GetSignTime() *string {
	return s.SignTime
}

func (s *AddSumRecordFlowPopRequest) GetType() *int32 {
	return s.Type
}

func (s *AddSumRecordFlowPopRequest) SetActivityId(v string) *AddSumRecordFlowPopRequest {
	s.ActivityId = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetCode(v string) *AddSumRecordFlowPopRequest {
	s.Code = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetConferenceName(v string) *AddSumRecordFlowPopRequest {
	s.ConferenceName = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetDeviceId(v string) *AddSumRecordFlowPopRequest {
	s.DeviceId = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetEntryName(v string) *AddSumRecordFlowPopRequest {
	s.EntryName = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetIdcard(v string) *AddSumRecordFlowPopRequest {
	s.Idcard = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetSignTime(v string) *AddSumRecordFlowPopRequest {
	s.SignTime = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) SetType(v int32) *AddSumRecordFlowPopRequest {
	s.Type = &v
	return s
}

func (s *AddSumRecordFlowPopRequest) Validate() error {
	return dara.Validate(s)
}
