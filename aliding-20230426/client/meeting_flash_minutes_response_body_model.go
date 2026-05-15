// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBasicInfo(v *MeetingFlashMinutesResponseBodyBasicInfo) *MeetingFlashMinutesResponseBody
	GetBasicInfo() *MeetingFlashMinutesResponseBodyBasicInfo
	SetFlashMinutesUrl(v string) *MeetingFlashMinutesResponseBody
	GetFlashMinutesUrl() *string
	SetFullSummary(v string) *MeetingFlashMinutesResponseBody
	GetFullSummary() *string
	SetRequestId(v string) *MeetingFlashMinutesResponseBody
	GetRequestId() *string
	SetTodos(v *MeetingFlashMinutesResponseBodyTodos) *MeetingFlashMinutesResponseBody
	GetTodos() *MeetingFlashMinutesResponseBodyTodos
	SetVendorRequestId(v string) *MeetingFlashMinutesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *MeetingFlashMinutesResponseBody
	GetVendorType() *string
}

type MeetingFlashMinutesResponseBody struct {
	BasicInfo *MeetingFlashMinutesResponseBodyBasicInfo `json:"basicInfo,omitempty" xml:"basicInfo,omitempty" type:"Struct"`
	// example:
	//
	// https://shanji.dingtalk.com/app/transcribes/76XXX
	FlashMinutesUrl *string `json:"flashMinutesUrl,omitempty" xml:"flashMinutesUrl,omitempty"`
	// example:
	//
	// XXX
	FullSummary *string `json:"fullSummary,omitempty" xml:"fullSummary,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Todos     *MeetingFlashMinutesResponseBodyTodos `json:"todos,omitempty" xml:"todos,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s MeetingFlashMinutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesResponseBody) GetBasicInfo() *MeetingFlashMinutesResponseBodyBasicInfo {
	return s.BasicInfo
}

func (s *MeetingFlashMinutesResponseBody) GetFlashMinutesUrl() *string {
	return s.FlashMinutesUrl
}

func (s *MeetingFlashMinutesResponseBody) GetFullSummary() *string {
	return s.FullSummary
}

func (s *MeetingFlashMinutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MeetingFlashMinutesResponseBody) GetTodos() *MeetingFlashMinutesResponseBodyTodos {
	return s.Todos
}

func (s *MeetingFlashMinutesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *MeetingFlashMinutesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *MeetingFlashMinutesResponseBody) SetBasicInfo(v *MeetingFlashMinutesResponseBodyBasicInfo) *MeetingFlashMinutesResponseBody {
	s.BasicInfo = v
	return s
}

func (s *MeetingFlashMinutesResponseBody) SetFlashMinutesUrl(v string) *MeetingFlashMinutesResponseBody {
	s.FlashMinutesUrl = &v
	return s
}

func (s *MeetingFlashMinutesResponseBody) SetFullSummary(v string) *MeetingFlashMinutesResponseBody {
	s.FullSummary = &v
	return s
}

func (s *MeetingFlashMinutesResponseBody) SetRequestId(v string) *MeetingFlashMinutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *MeetingFlashMinutesResponseBody) SetTodos(v *MeetingFlashMinutesResponseBodyTodos) *MeetingFlashMinutesResponseBody {
	s.Todos = v
	return s
}

func (s *MeetingFlashMinutesResponseBody) SetVendorRequestId(v string) *MeetingFlashMinutesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *MeetingFlashMinutesResponseBody) SetVendorType(v string) *MeetingFlashMinutesResponseBody {
	s.VendorType = &v
	return s
}

func (s *MeetingFlashMinutesResponseBody) Validate() error {
	if s.BasicInfo != nil {
		if err := s.BasicInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Todos != nil {
		if err := s.Todos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MeetingFlashMinutesResponseBodyBasicInfo struct {
	// example:
	//
	// 500529
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1778490089000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1778490089000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// u0VGeOiPUBSVMypV3Hylp7wXXX
	TaskCreator *string `json:"taskCreator,omitempty" xml:"taskCreator,omitempty"`
	// example:
	//
	// XXX
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://shanji.dingtalk.com/app/transcribes/76XXX
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MeetingFlashMinutesResponseBodyBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesResponseBodyBasicInfo) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) GetTaskCreator() *string {
	return s.TaskCreator
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) GetTitle() *string {
	return s.Title
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) GetUrl() *string {
	return s.Url
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) SetDuration(v int64) *MeetingFlashMinutesResponseBodyBasicInfo {
	s.Duration = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) SetEndTime(v int64) *MeetingFlashMinutesResponseBodyBasicInfo {
	s.EndTime = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) SetStartTime(v int64) *MeetingFlashMinutesResponseBodyBasicInfo {
	s.StartTime = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) SetTaskCreator(v string) *MeetingFlashMinutesResponseBodyBasicInfo {
	s.TaskCreator = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) SetTitle(v string) *MeetingFlashMinutesResponseBodyBasicInfo {
	s.Title = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) SetUrl(v string) *MeetingFlashMinutesResponseBodyBasicInfo {
	s.Url = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyBasicInfo) Validate() error {
	return dara.Validate(s)
}

type MeetingFlashMinutesResponseBodyTodos struct {
	Actions          []*string                                               `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	DingtalkTodoList []*MeetingFlashMinutesResponseBodyTodosDingtalkTodoList `json:"dingtalkTodoList,omitempty" xml:"dingtalkTodoList,omitempty" type:"Repeated"`
}

func (s MeetingFlashMinutesResponseBodyTodos) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesResponseBodyTodos) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesResponseBodyTodos) GetActions() []*string {
	return s.Actions
}

func (s *MeetingFlashMinutesResponseBodyTodos) GetDingtalkTodoList() []*MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	return s.DingtalkTodoList
}

func (s *MeetingFlashMinutesResponseBodyTodos) SetActions(v []*string) *MeetingFlashMinutesResponseBodyTodos {
	s.Actions = v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodos) SetDingtalkTodoList(v []*MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) *MeetingFlashMinutesResponseBodyTodos {
	s.DingtalkTodoList = v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodos) Validate() error {
	if s.DingtalkTodoList != nil {
		for _, item := range s.DingtalkTodoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MeetingFlashMinutesResponseBodyTodosDingtalkTodoList struct {
	// example:
	//
	// 1778490089000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// u0VGeOiPUBSVMypV3Hylp7wXXX
	CreatorUnionId *string `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	// example:
	//
	// deadline
	Deadline *string `json:"deadline,omitempty" xml:"deadline,omitempty"`
	// example:
	//
	// dingtalkTodoId
	DingtalkTodoId *string                                                             `json:"dingtalkTodoId,omitempty" xml:"dingtalkTodoId,omitempty"`
	ExecutorList   []*MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList `json:"executorList,omitempty" xml:"executorList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// example:
	//
	// minutesTodoId
	MinutesTodoId *string `json:"minutesTodoId,omitempty" xml:"minutesTodoId,omitempty"`
	// example:
	//
	// XXX
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetCreatorUnionId() *string {
	return s.CreatorUnionId
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetDeadline() *string {
	return s.Deadline
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetDingtalkTodoId() *string {
	return s.DingtalkTodoId
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetExecutorList() []*MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList {
	return s.ExecutorList
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetIsDone() *bool {
	return s.IsDone
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetMinutesTodoId() *string {
	return s.MinutesTodoId
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) GetTitle() *string {
	return s.Title
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetCreatedTime(v int64) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.CreatedTime = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetCreatorUnionId(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.CreatorUnionId = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetDeadline(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.Deadline = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetDingtalkTodoId(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.DingtalkTodoId = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetExecutorList(v []*MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.ExecutorList = v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetIsDone(v bool) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.IsDone = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetMinutesTodoId(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.MinutesTodoId = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) SetTitle(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList {
	s.Title = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoList) Validate() error {
	if s.ExecutorList != nil {
		for _, item := range s.ExecutorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList struct {
	// example:
	//
	// https://XXX221rNAbjNAbg_440_440.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// XXX
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// example:
	//
	// u0VGeOiPUBSVMypV3Hylp7wXXX
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) GetAvatar() *string {
	return s.Avatar
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) GetNick() *string {
	return s.Nick
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) GetUnionId() *string {
	return s.UnionId
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) SetAvatar(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList {
	s.Avatar = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) SetNick(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList {
	s.Nick = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) SetUnionId(v string) *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList {
	s.UnionId = &v
	return s
}

func (s *MeetingFlashMinutesResponseBodyTodosDingtalkTodoListExecutorList) Validate() error {
	return dara.Validate(s)
}
