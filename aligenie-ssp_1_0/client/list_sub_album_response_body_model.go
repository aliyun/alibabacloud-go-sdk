// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAlbumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSubAlbumResponseBody
	GetCode() *int32
	SetMessage(v string) *ListSubAlbumResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubAlbumResponseBody
	GetRequestId() *string
	SetResult(v *ListSubAlbumResponseBodyResult) *ListSubAlbumResponseBody
	GetResult() *ListSubAlbumResponseBodyResult
}

type ListSubAlbumResponseBody struct {
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
	// CB13B8D7-37FB-1B3E-8EB9-65BB413267E1
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListSubAlbumResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListSubAlbumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSubAlbumResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubAlbumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubAlbumResponseBody) GetResult() *ListSubAlbumResponseBodyResult {
	return s.Result
}

func (s *ListSubAlbumResponseBody) SetCode(v int32) *ListSubAlbumResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubAlbumResponseBody) SetMessage(v string) *ListSubAlbumResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubAlbumResponseBody) SetRequestId(v string) *ListSubAlbumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubAlbumResponseBody) SetResult(v *ListSubAlbumResponseBodyResult) *ListSubAlbumResponseBody {
	s.Result = v
	return s
}

func (s *ListSubAlbumResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSubAlbumResponseBodyResult struct {
	DataList []*ListSubAlbumResponseBodyResultDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	HasNext  *bool                                     `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPageCount *int32 `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListSubAlbumResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBodyResult) GetDataList() []*ListSubAlbumResponseBodyResultDataList {
	return s.DataList
}

func (s *ListSubAlbumResponseBodyResult) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListSubAlbumResponseBodyResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSubAlbumResponseBodyResult) GetTotalPageCount() *int32 {
	return s.TotalPageCount
}

func (s *ListSubAlbumResponseBodyResult) SetDataList(v []*ListSubAlbumResponseBodyResultDataList) *ListSubAlbumResponseBodyResult {
	s.DataList = v
	return s
}

func (s *ListSubAlbumResponseBodyResult) SetHasNext(v bool) *ListSubAlbumResponseBodyResult {
	s.HasNext = &v
	return s
}

func (s *ListSubAlbumResponseBodyResult) SetTotalCount(v int32) *ListSubAlbumResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListSubAlbumResponseBodyResult) SetTotalPageCount(v int32) *ListSubAlbumResponseBodyResult {
	s.TotalPageCount = &v
	return s
}

func (s *ListSubAlbumResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListSubAlbumResponseBodyResultDataList struct {
	// example:
	//
	// 4476001
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// example:
	//
	// 80011
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// https://ailabs.alibabausercontent.com/images/17825/jknoamc2.jpg
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// example:
	//
	// 1
	Id           *int64                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	IsAdded      *bool                                               `json:"IsAdded,omitempty" xml:"IsAdded,omitempty"`
	ScheduleInfo *ListSubAlbumResponseBodyResultDataListScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Sequence *int64 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// 睡前故事
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 23
	TotalEpisode *int32 `json:"TotalEpisode,omitempty" xml:"TotalEpisode,omitempty"`
}

func (s ListSubAlbumResponseBodyResultDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumResponseBodyResultDataList) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBodyResultDataList) GetAlbumId() *string {
	return s.AlbumId
}

func (s *ListSubAlbumResponseBodyResultDataList) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *ListSubAlbumResponseBodyResultDataList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListSubAlbumResponseBodyResultDataList) GetId() *int64 {
	return s.Id
}

func (s *ListSubAlbumResponseBodyResultDataList) GetIsAdded() *bool {
	return s.IsAdded
}

func (s *ListSubAlbumResponseBodyResultDataList) GetScheduleInfo() *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	return s.ScheduleInfo
}

func (s *ListSubAlbumResponseBodyResultDataList) GetSequence() *int64 {
	return s.Sequence
}

func (s *ListSubAlbumResponseBodyResultDataList) GetTitle() *string {
	return s.Title
}

func (s *ListSubAlbumResponseBodyResultDataList) GetTotalEpisode() *int32 {
	return s.TotalEpisode
}

func (s *ListSubAlbumResponseBodyResultDataList) SetAlbumId(v string) *ListSubAlbumResponseBodyResultDataList {
	s.AlbumId = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetCategoryId(v int32) *ListSubAlbumResponseBodyResultDataList {
	s.CategoryId = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetCoverUrl(v string) *ListSubAlbumResponseBodyResultDataList {
	s.CoverUrl = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetId(v int64) *ListSubAlbumResponseBodyResultDataList {
	s.Id = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetIsAdded(v bool) *ListSubAlbumResponseBodyResultDataList {
	s.IsAdded = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetScheduleInfo(v *ListSubAlbumResponseBodyResultDataListScheduleInfo) *ListSubAlbumResponseBodyResultDataList {
	s.ScheduleInfo = v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetSequence(v int64) *ListSubAlbumResponseBodyResultDataList {
	s.Sequence = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetTitle(v string) *ListSubAlbumResponseBodyResultDataList {
	s.Title = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) SetTotalEpisode(v int32) *ListSubAlbumResponseBodyResultDataList {
	s.TotalEpisode = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataList) Validate() error {
	return dara.Validate(s)
}

type ListSubAlbumResponseBodyResultDataListScheduleInfo struct {
	DaysOfWeek []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// 23
	Minute     *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	ScheduleId *int64 `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
}

func (s ListSubAlbumResponseBodyResultDataListScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumResponseBodyResultDataListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) GetDaysOfWeek() []*int32 {
	return s.DaysOfWeek
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) GetHour() *int32 {
	return s.Hour
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) GetMinute() *int32 {
	return s.Minute
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) GetScheduleId() *int64 {
	return s.ScheduleId
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetDaysOfWeek(v []*int32) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.DaysOfWeek = v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetHour(v int32) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.Hour = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetMinute(v int32) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.Minute = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) SetScheduleId(v int64) *ListSubAlbumResponseBodyResultDataListScheduleInfo {
	s.ScheduleId = &v
	return s
}

func (s *ListSubAlbumResponseBodyResultDataListScheduleInfo) Validate() error {
	return dara.Validate(s)
}
