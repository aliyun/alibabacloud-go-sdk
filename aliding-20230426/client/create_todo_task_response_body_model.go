// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizTag(v string) *CreateTodoTaskResponseBody
	GetBizTag() *string
	SetContentFieldList(v []*CreateTodoTaskResponseBodyContentFieldList) *CreateTodoTaskResponseBody
	GetContentFieldList() []*CreateTodoTaskResponseBodyContentFieldList
	SetCreatedTime(v int64) *CreateTodoTaskResponseBody
	GetCreatedTime() *int64
	SetCreatorId(v string) *CreateTodoTaskResponseBody
	GetCreatorId() *string
	SetDescription(v string) *CreateTodoTaskResponseBody
	GetDescription() *string
	SetDetailUrl(v *CreateTodoTaskResponseBodyDetailUrl) *CreateTodoTaskResponseBody
	GetDetailUrl() *CreateTodoTaskResponseBodyDetailUrl
	SetDone(v bool) *CreateTodoTaskResponseBody
	GetDone() *bool
	SetDueTime(v int64) *CreateTodoTaskResponseBody
	GetDueTime() *int64
	SetExecutorIds(v []*string) *CreateTodoTaskResponseBody
	GetExecutorIds() []*string
	SetFinishTime(v int64) *CreateTodoTaskResponseBody
	GetFinishTime() *int64
	SetId(v string) *CreateTodoTaskResponseBody
	GetId() *string
	SetIsOnlyShowExecutor(v bool) *CreateTodoTaskResponseBody
	GetIsOnlyShowExecutor() *bool
	SetModifiedTime(v int64) *CreateTodoTaskResponseBody
	GetModifiedTime() *int64
	SetModifierId(v string) *CreateTodoTaskResponseBody
	GetModifierId() *string
	SetNotifyConfigs(v *CreateTodoTaskResponseBodyNotifyConfigs) *CreateTodoTaskResponseBody
	GetNotifyConfigs() *CreateTodoTaskResponseBodyNotifyConfigs
	SetParticipantIds(v []*string) *CreateTodoTaskResponseBody
	GetParticipantIds() []*string
	SetPriority(v int32) *CreateTodoTaskResponseBody
	GetPriority() *int32
	SetRequestId(v string) *CreateTodoTaskResponseBody
	GetRequestId() *string
	SetSource(v string) *CreateTodoTaskResponseBody
	GetSource() *string
	SetSourceId(v string) *CreateTodoTaskResponseBody
	GetSourceId() *string
	SetStartTime(v int64) *CreateTodoTaskResponseBody
	GetStartTime() *int64
	SetSubject(v string) *CreateTodoTaskResponseBody
	GetSubject() *string
}

type CreateTodoTaskResponseBody struct {
	// example:
	//
	// isv_dingtalkTodo
	BizTag           *string                                       `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	ContentFieldList []*CreateTodoTaskResponseBodyContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// example:
	//
	// 1617675200000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// PUoiinWIpa2yH2ymhiiGiP6g
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 应用可以调用该接口发起一个钉钉待办任务，该待办事项会出现在钉钉客户端“待办”页面，需要注意的是，通过开放接口发起的待办，目前仅支持直接跳转ISV应用详情页（ISV在调该接口时需传入自身应用详情页链接）。
	Description *string                              `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl   *CreateTodoTaskResponseBodyDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// example:
	//
	// false
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// example:
	//
	// 1617675100000
	DueTime     *int64    `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1617675200000
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// OPJpwtwPVNGIFKURjrzd
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsOnlyShowExecutor *bool `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	// example:
	//
	// 1617675200000
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// PUoiinWIpa2yH2ymhiiGiP6g
	ModifierId     *string                                  `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	NotifyConfigs  *CreateTodoTaskResponseBodyNotifyConfigs `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	ParticipantIds []*string                                `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// isv_dingtalkTodo
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// isv_dingtalkTodo1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 1617675000000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 接入钉钉待办
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateTodoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBody) GetBizTag() *string {
	return s.BizTag
}

func (s *CreateTodoTaskResponseBody) GetContentFieldList() []*CreateTodoTaskResponseBodyContentFieldList {
	return s.ContentFieldList
}

func (s *CreateTodoTaskResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *CreateTodoTaskResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateTodoTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateTodoTaskResponseBody) GetDetailUrl() *CreateTodoTaskResponseBodyDetailUrl {
	return s.DetailUrl
}

func (s *CreateTodoTaskResponseBody) GetDone() *bool {
	return s.Done
}

func (s *CreateTodoTaskResponseBody) GetDueTime() *int64 {
	return s.DueTime
}

func (s *CreateTodoTaskResponseBody) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *CreateTodoTaskResponseBody) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *CreateTodoTaskResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateTodoTaskResponseBody) GetIsOnlyShowExecutor() *bool {
	return s.IsOnlyShowExecutor
}

func (s *CreateTodoTaskResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *CreateTodoTaskResponseBody) GetModifierId() *string {
	return s.ModifierId
}

func (s *CreateTodoTaskResponseBody) GetNotifyConfigs() *CreateTodoTaskResponseBodyNotifyConfigs {
	return s.NotifyConfigs
}

func (s *CreateTodoTaskResponseBody) GetParticipantIds() []*string {
	return s.ParticipantIds
}

func (s *CreateTodoTaskResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTodoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTodoTaskResponseBody) GetSource() *string {
	return s.Source
}

func (s *CreateTodoTaskResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateTodoTaskResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateTodoTaskResponseBody) GetSubject() *string {
	return s.Subject
}

func (s *CreateTodoTaskResponseBody) SetBizTag(v string) *CreateTodoTaskResponseBody {
	s.BizTag = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetContentFieldList(v []*CreateTodoTaskResponseBodyContentFieldList) *CreateTodoTaskResponseBody {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatedTime(v int64) *CreateTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetCreatorId(v string) *CreateTodoTaskResponseBody {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDescription(v string) *CreateTodoTaskResponseBody {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDetailUrl(v *CreateTodoTaskResponseBodyDetailUrl) *CreateTodoTaskResponseBody {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDone(v bool) *CreateTodoTaskResponseBody {
	s.Done = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetDueTime(v int64) *CreateTodoTaskResponseBody {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetExecutorIds(v []*string) *CreateTodoTaskResponseBody {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetFinishTime(v int64) *CreateTodoTaskResponseBody {
	s.FinishTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetId(v string) *CreateTodoTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskResponseBody {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifiedTime(v int64) *CreateTodoTaskResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetModifierId(v string) *CreateTodoTaskResponseBody {
	s.ModifierId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetNotifyConfigs(v *CreateTodoTaskResponseBodyNotifyConfigs) *CreateTodoTaskResponseBody {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetParticipantIds(v []*string) *CreateTodoTaskResponseBody {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskResponseBody) SetPriority(v int32) *CreateTodoTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetRequestId(v string) *CreateTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSource(v string) *CreateTodoTaskResponseBody {
	s.Source = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSourceId(v string) *CreateTodoTaskResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetStartTime(v int64) *CreateTodoTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateTodoTaskResponseBody) SetSubject(v string) *CreateTodoTaskResponseBody {
	s.Subject = &v
	return s
}

func (s *CreateTodoTaskResponseBody) Validate() error {
	if s.ContentFieldList != nil {
		for _, item := range s.ContentFieldList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DetailUrl != nil {
		if err := s.DetailUrl.Validate(); err != nil {
			return err
		}
	}
	if s.NotifyConfigs != nil {
		if err := s.NotifyConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTodoTaskResponseBodyContentFieldList struct {
	// fieldKey
	//
	// example:
	//
	// fieldKey
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// fieldValue
	//
	// example:
	//
	// fieldValue
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateTodoTaskResponseBodyContentFieldList) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskResponseBodyContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyContentFieldList) GetFieldKey() *string {
	return s.FieldKey
}

func (s *CreateTodoTaskResponseBodyContentFieldList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateTodoTaskResponseBodyContentFieldList) SetFieldKey(v string) *CreateTodoTaskResponseBodyContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTaskResponseBodyContentFieldList) SetFieldValue(v string) *CreateTodoTaskResponseBodyContentFieldList {
	s.FieldValue = &v
	return s
}

func (s *CreateTodoTaskResponseBodyContentFieldList) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskResponseBodyDetailUrl struct {
	// example:
	//
	// https://www.dingtalk.com
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s CreateTodoTaskResponseBodyDetailUrl) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskResponseBodyDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyDetailUrl) GetAppUrl() *string {
	return s.AppUrl
}

func (s *CreateTodoTaskResponseBodyDetailUrl) GetPcUrl() *string {
	return s.PcUrl
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetAppUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateTodoTaskResponseBodyDetailUrl) SetPcUrl(v string) *CreateTodoTaskResponseBodyDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *CreateTodoTaskResponseBodyDetailUrl) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskResponseBodyNotifyConfigs struct {
	// example:
	//
	// 1
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskResponseBodyNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) GetDingNotify() *string {
	return s.DingNotify
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskResponseBodyNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreateTodoTaskResponseBodyNotifyConfigs) Validate() error {
	return dara.Validate(s)
}
