// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *TaskListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *TaskListResponseBody
	GetCode() *int64
	SetMessage(v string) *TaskListResponseBody
	GetMessage() *string
	SetModel(v []*TaskListResponseBodyModel) *TaskListResponseBody
	GetModel() []*TaskListResponseBodyModel
	SetRequestId(v string) *TaskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TaskListResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *TaskListResponseBody
	GetTimestamp() *int64
}

type TaskListResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 30
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*TaskListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 80
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskListResponseBody) GoString() string {
	return s.String()
}

func (s *TaskListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *TaskListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TaskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskListResponseBody) GetModel() []*TaskListResponseBodyModel {
	return s.Model
}

func (s *TaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TaskListResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TaskListResponseBody) SetAccessDeniedDetail(v string) *TaskListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TaskListResponseBody) SetCode(v int64) *TaskListResponseBody {
	s.Code = &v
	return s
}

func (s *TaskListResponseBody) SetMessage(v string) *TaskListResponseBody {
	s.Message = &v
	return s
}

func (s *TaskListResponseBody) SetModel(v []*TaskListResponseBodyModel) *TaskListResponseBody {
	s.Model = v
	return s
}

func (s *TaskListResponseBody) SetRequestId(v string) *TaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskListResponseBody) SetSuccess(v bool) *TaskListResponseBody {
	s.Success = &v
	return s
}

func (s *TaskListResponseBody) SetTimestamp(v int64) *TaskListResponseBody {
	s.Timestamp = &v
	return s
}

func (s *TaskListResponseBody) Validate() error {
	if s.Model != nil {
		for _, item := range s.Model {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TaskListResponseBodyModel struct {
	// 外呼时间段
	//
	// example:
	//
	// “8:00~20:30”
	AllowCallTime *string `json:"AllowCallTime,omitempty" xml:"AllowCallTime,omitempty"`
	// 外呼时间段格式化
	//
	// example:
	//
	// “8:00 ~ 20:00”
	AllowCallTimeFormat *string `json:"AllowCallTimeFormat,omitempty" xml:"AllowCallTimeFormat,omitempty"`
	// 外呼时间
	//
	// example:
	//
	// “1,2,3”
	AllowDayOfWeek *string `json:"AllowDayOfWeek,omitempty" xml:"AllowDayOfWeek,omitempty"`
	// 外呼类型
	//
	// example:
	//
	// 95
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 闪信模板id
	//
	// example:
	//
	// 99
	FlashSmsTemplateId *int64 `json:"FlashSmsTemplateId,omitempty" xml:"FlashSmsTemplateId,omitempty"`
	// 闪信模板名称
	//
	// example:
	//
	// 示例值示例值示例值
	FlashSmsTemplateName *string `json:"FlashSmsTemplateName,omitempty" xml:"FlashSmsTemplateName,omitempty"`
	// 发送闪信配置，可选0，1；0表示否，1表示是
	//
	// example:
	//
	// 1
	FlashSmsType *int64 `json:"FlashSmsType,omitempty" xml:"FlashSmsType,omitempty"`
	// 最近导入时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	ImportTime *string `json:"ImportTime,omitempty" xml:"ImportTime,omitempty"`
	// 意向标签列表
	IntentTags []*TaskListResponseBodyModelIntentTags `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 接通重呼次数
	//
	// example:
	//
	// 31
	InvalidReCall *int64 `json:"InvalidReCall,omitempty" xml:"InvalidReCall,omitempty"`
	// 最后外呼时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	LastCallTime *string `json:"LastCallTime,omitempty" xml:"LastCallTime,omitempty"`
	// 最大并发数
	//
	// example:
	//
	// 95
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// 个性标签列表
	PersonalityTags []*string `json:"PersonalityTags,omitempty" xml:"PersonalityTags,omitempty" type:"Repeated"`
	// 优先任务
	//
	// example:
	//
	// 66
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 任务所需参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 自动重呼
	//
	// example:
	//
	// 20
	RecallType *int64 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// 发送短信
	//
	// example:
	//
	// 39
	SendSms *int64 `json:"SendSms,omitempty" xml:"SendSms,omitempty"`
	// 短信模板
	//
	// example:
	//
	// 示例值示例值
	SmsName *string `json:"SmsName,omitempty" xml:"SmsName,omitempty"`
	// 任务状态
	//
	// example:
	//
	// 79
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 68
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务名称
	//
	// example:
	//
	// 示例值示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 话术模板Id
	//
	// example:
	//
	// 示例值示例值示例值
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板名称
	//
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s TaskListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s TaskListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskListResponseBodyModel) GetAllowCallTime() *string {
	return s.AllowCallTime
}

func (s *TaskListResponseBodyModel) GetAllowCallTimeFormat() *string {
	return s.AllowCallTimeFormat
}

func (s *TaskListResponseBodyModel) GetAllowDayOfWeek() *string {
	return s.AllowDayOfWeek
}

func (s *TaskListResponseBodyModel) GetCallType() *int64 {
	return s.CallType
}

func (s *TaskListResponseBodyModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *TaskListResponseBodyModel) GetFlashSmsTemplateId() *int64 {
	return s.FlashSmsTemplateId
}

func (s *TaskListResponseBodyModel) GetFlashSmsTemplateName() *string {
	return s.FlashSmsTemplateName
}

func (s *TaskListResponseBodyModel) GetFlashSmsType() *int64 {
	return s.FlashSmsType
}

func (s *TaskListResponseBodyModel) GetImportTime() *string {
	return s.ImportTime
}

func (s *TaskListResponseBodyModel) GetIntentTags() []*TaskListResponseBodyModelIntentTags {
	return s.IntentTags
}

func (s *TaskListResponseBodyModel) GetInvalidReCall() *int64 {
	return s.InvalidReCall
}

func (s *TaskListResponseBodyModel) GetLastCallTime() *string {
	return s.LastCallTime
}

func (s *TaskListResponseBodyModel) GetMaxConcurrency() *int64 {
	return s.MaxConcurrency
}

func (s *TaskListResponseBodyModel) GetPersonalityTags() []*string {
	return s.PersonalityTags
}

func (s *TaskListResponseBodyModel) GetPriority() *int64 {
	return s.Priority
}

func (s *TaskListResponseBodyModel) GetProperties() *string {
	return s.Properties
}

func (s *TaskListResponseBodyModel) GetRecallType() *int64 {
	return s.RecallType
}

func (s *TaskListResponseBodyModel) GetSendSms() *int64 {
	return s.SendSms
}

func (s *TaskListResponseBodyModel) GetSmsName() *string {
	return s.SmsName
}

func (s *TaskListResponseBodyModel) GetStatus() *int64 {
	return s.Status
}

func (s *TaskListResponseBodyModel) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskListResponseBodyModel) GetTaskName() *string {
	return s.TaskName
}

func (s *TaskListResponseBodyModel) GetTemplateId() *string {
	return s.TemplateId
}

func (s *TaskListResponseBodyModel) GetTemplateName() *string {
	return s.TemplateName
}

func (s *TaskListResponseBodyModel) SetAllowCallTime(v string) *TaskListResponseBodyModel {
	s.AllowCallTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetAllowCallTimeFormat(v string) *TaskListResponseBodyModel {
	s.AllowCallTimeFormat = &v
	return s
}

func (s *TaskListResponseBodyModel) SetAllowDayOfWeek(v string) *TaskListResponseBodyModel {
	s.AllowDayOfWeek = &v
	return s
}

func (s *TaskListResponseBodyModel) SetCallType(v int64) *TaskListResponseBodyModel {
	s.CallType = &v
	return s
}

func (s *TaskListResponseBodyModel) SetCreateTime(v string) *TaskListResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetFlashSmsTemplateId(v int64) *TaskListResponseBodyModel {
	s.FlashSmsTemplateId = &v
	return s
}

func (s *TaskListResponseBodyModel) SetFlashSmsTemplateName(v string) *TaskListResponseBodyModel {
	s.FlashSmsTemplateName = &v
	return s
}

func (s *TaskListResponseBodyModel) SetFlashSmsType(v int64) *TaskListResponseBodyModel {
	s.FlashSmsType = &v
	return s
}

func (s *TaskListResponseBodyModel) SetImportTime(v string) *TaskListResponseBodyModel {
	s.ImportTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetIntentTags(v []*TaskListResponseBodyModelIntentTags) *TaskListResponseBodyModel {
	s.IntentTags = v
	return s
}

func (s *TaskListResponseBodyModel) SetInvalidReCall(v int64) *TaskListResponseBodyModel {
	s.InvalidReCall = &v
	return s
}

func (s *TaskListResponseBodyModel) SetLastCallTime(v string) *TaskListResponseBodyModel {
	s.LastCallTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetMaxConcurrency(v int64) *TaskListResponseBodyModel {
	s.MaxConcurrency = &v
	return s
}

func (s *TaskListResponseBodyModel) SetPersonalityTags(v []*string) *TaskListResponseBodyModel {
	s.PersonalityTags = v
	return s
}

func (s *TaskListResponseBodyModel) SetPriority(v int64) *TaskListResponseBodyModel {
	s.Priority = &v
	return s
}

func (s *TaskListResponseBodyModel) SetProperties(v string) *TaskListResponseBodyModel {
	s.Properties = &v
	return s
}

func (s *TaskListResponseBodyModel) SetRecallType(v int64) *TaskListResponseBodyModel {
	s.RecallType = &v
	return s
}

func (s *TaskListResponseBodyModel) SetSendSms(v int64) *TaskListResponseBodyModel {
	s.SendSms = &v
	return s
}

func (s *TaskListResponseBodyModel) SetSmsName(v string) *TaskListResponseBodyModel {
	s.SmsName = &v
	return s
}

func (s *TaskListResponseBodyModel) SetStatus(v int64) *TaskListResponseBodyModel {
	s.Status = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTaskId(v int64) *TaskListResponseBodyModel {
	s.TaskId = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTaskName(v string) *TaskListResponseBodyModel {
	s.TaskName = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTemplateId(v string) *TaskListResponseBodyModel {
	s.TemplateId = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTemplateName(v string) *TaskListResponseBodyModel {
	s.TemplateName = &v
	return s
}

func (s *TaskListResponseBodyModel) Validate() error {
	if s.IntentTags != nil {
		for _, item := range s.IntentTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TaskListResponseBodyModelIntentTags struct {
	// 意向标签描述
	//
	// example:
	//
	// 示例值示例值
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// 意向标签ID
	//
	// example:
	//
	// 示例值示例值
	IntentTag *string `json:"IntentTag,omitempty" xml:"IntentTag,omitempty"`
}

func (s TaskListResponseBodyModelIntentTags) String() string {
	return dara.Prettify(s)
}

func (s TaskListResponseBodyModelIntentTags) GoString() string {
	return s.String()
}

func (s *TaskListResponseBodyModelIntentTags) GetIntentDescription() *string {
	return s.IntentDescription
}

func (s *TaskListResponseBodyModelIntentTags) GetIntentTag() *string {
	return s.IntentTag
}

func (s *TaskListResponseBodyModelIntentTags) SetIntentDescription(v string) *TaskListResponseBodyModelIntentTags {
	s.IntentDescription = &v
	return s
}

func (s *TaskListResponseBodyModelIntentTags) SetIntentTag(v string) *TaskListResponseBodyModelIntentTags {
	s.IntentTag = &v
	return s
}

func (s *TaskListResponseBodyModelIntentTags) Validate() error {
	return dara.Validate(s)
}
