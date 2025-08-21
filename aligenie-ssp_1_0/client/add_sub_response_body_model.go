// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddSubResponseBody
	GetCode() *int32
	SetMessage(v string) *AddSubResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSubResponseBody
	GetRequestId() *string
	SetResult(v *AddSubResponseBodyResult) *AddSubResponseBody
	GetResult() *AddSubResponseBodyResult
}

type AddSubResponseBody struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B99D27ED-4E12-1414-9FDE-599C57C4B204
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddSubResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s AddSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSubResponseBody) GoString() string {
	return s.String()
}

func (s *AddSubResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddSubResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSubResponseBody) GetResult() *AddSubResponseBodyResult {
	return s.Result
}

func (s *AddSubResponseBody) SetCode(v int32) *AddSubResponseBody {
	s.Code = &v
	return s
}

func (s *AddSubResponseBody) SetMessage(v string) *AddSubResponseBody {
	s.Message = &v
	return s
}

func (s *AddSubResponseBody) SetRequestId(v string) *AddSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSubResponseBody) SetResult(v *AddSubResponseBodyResult) *AddSubResponseBody {
	s.Result = v
	return s
}

func (s *AddSubResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddSubResponseBodyResult struct {
	// example:
	//
	// 51999575
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// example:
	//
	// 2
	DailyStudyCnt *int32 `json:"DailyStudyCnt,omitempty" xml:"DailyStudyCnt,omitempty"`
	// example:
	//
	// 5874DBCCA3038FAA1A70A8060F07F26D
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 81
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sequence
	PlayMode     *string                               `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	ScheduleInfo *AddSubResponseBodyResultScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1152893538998276761
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddSubResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddSubResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddSubResponseBodyResult) GetAlbumId() *string {
	return s.AlbumId
}

func (s *AddSubResponseBodyResult) GetDailyStudyCnt() *int32 {
	return s.DailyStudyCnt
}

func (s *AddSubResponseBodyResult) GetDeviceId() *string {
	return s.DeviceId
}

func (s *AddSubResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *AddSubResponseBodyResult) GetPlayMode() *string {
	return s.PlayMode
}

func (s *AddSubResponseBodyResult) GetScheduleInfo() *AddSubResponseBodyResultScheduleInfo {
	return s.ScheduleInfo
}

func (s *AddSubResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *AddSubResponseBodyResult) SetAlbumId(v string) *AddSubResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *AddSubResponseBodyResult) SetDailyStudyCnt(v int32) *AddSubResponseBodyResult {
	s.DailyStudyCnt = &v
	return s
}

func (s *AddSubResponseBodyResult) SetDeviceId(v string) *AddSubResponseBodyResult {
	s.DeviceId = &v
	return s
}

func (s *AddSubResponseBodyResult) SetId(v int64) *AddSubResponseBodyResult {
	s.Id = &v
	return s
}

func (s *AddSubResponseBodyResult) SetPlayMode(v string) *AddSubResponseBodyResult {
	s.PlayMode = &v
	return s
}

func (s *AddSubResponseBodyResult) SetScheduleInfo(v *AddSubResponseBodyResultScheduleInfo) *AddSubResponseBodyResult {
	s.ScheduleInfo = v
	return s
}

func (s *AddSubResponseBodyResult) SetUserId(v string) *AddSubResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *AddSubResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type AddSubResponseBodyResultScheduleInfo struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 23
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
}

func (s AddSubResponseBodyResultScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s AddSubResponseBodyResultScheduleInfo) GoString() string {
	return s.String()
}

func (s *AddSubResponseBodyResultScheduleInfo) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *AddSubResponseBodyResultScheduleInfo) GetHour() *int32 {
	return s.Hour
}

func (s *AddSubResponseBodyResultScheduleInfo) GetMinute() *int32 {
	return s.Minute
}

func (s *AddSubResponseBodyResultScheduleInfo) SetDaysOfWeek(v []*int32) *AddSubResponseBodyResultScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *AddSubResponseBodyResultScheduleInfo) SetHour(v int32) *AddSubResponseBodyResultScheduleInfo {
	s.Hour = &v
	return s
}

func (s *AddSubResponseBodyResultScheduleInfo) SetMinute(v int32) *AddSubResponseBodyResultScheduleInfo {
	s.Minute = &v
	return s
}

func (s *AddSubResponseBodyResultScheduleInfo) Validate() error {
	return dara.Validate(s)
}
