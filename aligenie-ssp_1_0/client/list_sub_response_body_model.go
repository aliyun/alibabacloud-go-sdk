// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSubResponseBody
	GetCode() *int32
	SetMessage(v string) *ListSubResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubResponseBody
	GetRequestId() *string
	SetResult(v *ListSubResponseBodyResult) *ListSubResponseBody
	GetResult() *ListSubResponseBodyResult
}

type ListSubResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0D0C09C2-ADC1-198B-964D-24F4FAD967DB
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListSubResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSubResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubResponseBody) GetResult() *ListSubResponseBodyResult {
	return s.Result
}

func (s *ListSubResponseBody) SetCode(v int32) *ListSubResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubResponseBody) SetMessage(v string) *ListSubResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubResponseBody) SetRequestId(v string) *ListSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubResponseBody) SetResult(v *ListSubResponseBodyResult) *ListSubResponseBody {
	s.Result = v
	return s
}

func (s *ListSubResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSubResponseBodyResult struct {
	DataList []*ListSubResponseBodyResultDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	HasNext  *bool                                `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListSubResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubResponseBodyResult) GetDataList() []*ListSubResponseBodyResultDataList {
	return s.DataList
}

func (s *ListSubResponseBodyResult) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListSubResponseBodyResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSubResponseBodyResult) GetTotalPageCount() *int32 {
	return s.TotalPageCount
}

func (s *ListSubResponseBodyResult) SetDataList(v []*ListSubResponseBodyResultDataList) *ListSubResponseBodyResult {
	s.DataList = v
	return s
}

func (s *ListSubResponseBodyResult) SetHasNext(v bool) *ListSubResponseBodyResult {
	s.HasNext = &v
	return s
}

func (s *ListSubResponseBodyResult) SetTotalCount(v int64) *ListSubResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListSubResponseBodyResult) SetTotalPageCount(v int32) *ListSubResponseBodyResult {
	s.TotalPageCount = &v
	return s
}

func (s *ListSubResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListSubResponseBodyResultDataList struct {
	// example:
	//
	// 51999575
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// example:
	//
	// https://ailabs.alibabausercontent.com/images/8838/1600839452498.jpg
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
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
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sequence
	PlayMode     *string                                        `json:"PlayMode,omitempty" xml:"PlayMode,omitempty"`
	ScheduleInfo *ListSubResponseBodyResultDataListScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// 小科学家探索
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1152893538998276761
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSubResponseBodyResultDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSubResponseBodyResultDataList) GoString() string {
	return s.String()
}

func (s *ListSubResponseBodyResultDataList) GetAlbumId() *string {
	return s.AlbumId
}

func (s *ListSubResponseBodyResultDataList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListSubResponseBodyResultDataList) GetDailyStudyCnt() *int32 {
	return s.DailyStudyCnt
}

func (s *ListSubResponseBodyResultDataList) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListSubResponseBodyResultDataList) GetId() *int64 {
	return s.Id
}

func (s *ListSubResponseBodyResultDataList) GetPlayMode() *string {
	return s.PlayMode
}

func (s *ListSubResponseBodyResultDataList) GetScheduleInfo() *ListSubResponseBodyResultDataListScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListSubResponseBodyResultDataList) GetTitle() *string {
	return s.Title
}

func (s *ListSubResponseBodyResultDataList) GetUserId() *int64 {
	return s.UserId
}

func (s *ListSubResponseBodyResultDataList) SetAlbumId(v string) *ListSubResponseBodyResultDataList {
	s.AlbumId = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetCoverUrl(v string) *ListSubResponseBodyResultDataList {
	s.CoverUrl = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetDailyStudyCnt(v int32) *ListSubResponseBodyResultDataList {
	s.DailyStudyCnt = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetDeviceId(v string) *ListSubResponseBodyResultDataList {
	s.DeviceId = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetId(v int64) *ListSubResponseBodyResultDataList {
	s.Id = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetPlayMode(v string) *ListSubResponseBodyResultDataList {
	s.PlayMode = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetScheduleInfo(v *ListSubResponseBodyResultDataListScheduleInfo) *ListSubResponseBodyResultDataList {
	s.ScheduleInfo = v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetTitle(v string) *ListSubResponseBodyResultDataList {
	s.Title = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) SetUserId(v int64) *ListSubResponseBodyResultDataList {
	s.UserId = &v
	return s
}

func (s *ListSubResponseBodyResultDataList) Validate() error {
	return dara.Validate(s)
}

type ListSubResponseBodyResultDataListScheduleInfo struct {
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

func (s ListSubResponseBodyResultDataListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubResponseBodyResultDataListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) GetHour() *int32 {
	return s.Hour
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) GetMinute() *int32 {
	return s.Minute
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) SetDaysOfWeek(v []*int32) *ListSubResponseBodyResultDataListScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) SetHour(v int32) *ListSubResponseBodyResultDataListScheduleInfo {
	s.Hour = &v
	return s
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) SetMinute(v int32) *ListSubResponseBodyResultDataListScheduleInfo {
	s.Minute = &v
	return s
}

func (s *ListSubResponseBodyResultDataListScheduleInfo) Validate() error {
	return dara.Validate(s)
}
