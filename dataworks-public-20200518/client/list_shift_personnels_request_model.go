// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShiftPersonnelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListShiftPersonnelsRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListShiftPersonnelsRequest
	GetEndTime() *int64
	SetShiftPersonUID(v string) *ListShiftPersonnelsRequest
	GetShiftPersonUID() *string
	SetShiftScheduleIdentifier(v string) *ListShiftPersonnelsRequest
	GetShiftScheduleIdentifier() *string
	SetUserType(v string) *ListShiftPersonnelsRequest
	GetUserType() *string
}

type ListShiftPersonnelsRequest struct {
	// The time when the on-duty engineer starts a shift. Set the value to a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1593950832000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The time when the on-duty engineer ends a shift. Set the value to a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1593950832000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud account. You can log on to the DataWorks console and move the pointer over the profile picture in the upper-right corner to view the ID.
	//
	// example:
	//
	// 1933790****551
	ShiftPersonUID *string `json:"ShiftPersonUID,omitempty" xml:"ShiftPersonUID,omitempty"`
	// The unique identifier of the shift schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab6456ada634b2f938ee******9b45b
	ShiftScheduleIdentifier *string `json:"ShiftScheduleIdentifier,omitempty" xml:"ShiftScheduleIdentifier,omitempty"`
	// The type of the on-duty engineer that you want to query. Valid values: ALL, PRIMARY, BACKUP, and DESIGNATED_USER.
	//
	// example:
	//
	// ALL
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListShiftPersonnelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListShiftPersonnelsRequest) GoString() string {
	return s.String()
}

func (s *ListShiftPersonnelsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListShiftPersonnelsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListShiftPersonnelsRequest) GetShiftPersonUID() *string {
	return s.ShiftPersonUID
}

func (s *ListShiftPersonnelsRequest) GetShiftScheduleIdentifier() *string {
	return s.ShiftScheduleIdentifier
}

func (s *ListShiftPersonnelsRequest) GetUserType() *string {
	return s.UserType
}

func (s *ListShiftPersonnelsRequest) SetBeginTime(v int64) *ListShiftPersonnelsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListShiftPersonnelsRequest) SetEndTime(v int64) *ListShiftPersonnelsRequest {
	s.EndTime = &v
	return s
}

func (s *ListShiftPersonnelsRequest) SetShiftPersonUID(v string) *ListShiftPersonnelsRequest {
	s.ShiftPersonUID = &v
	return s
}

func (s *ListShiftPersonnelsRequest) SetShiftScheduleIdentifier(v string) *ListShiftPersonnelsRequest {
	s.ShiftScheduleIdentifier = &v
	return s
}

func (s *ListShiftPersonnelsRequest) SetUserType(v string) *ListShiftPersonnelsRequest {
	s.UserType = &v
	return s
}

func (s *ListShiftPersonnelsRequest) Validate() error {
	return dara.Validate(s)
}
