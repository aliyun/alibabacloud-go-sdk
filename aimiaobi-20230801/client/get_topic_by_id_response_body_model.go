// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTopicByIdResponseBody
	GetCode() *string
	SetData(v *GetTopicByIdResponseBodyData) *GetTopicByIdResponseBody
	GetData() *GetTopicByIdResponseBodyData
	SetHttpStatusCode(v int32) *GetTopicByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTopicByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTopicByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicByIdResponseBody
	GetSuccess() *bool
}

type GetTopicByIdResponseBody struct {
	// example:
	//
	// NoData
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTopicByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTopicByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTopicByIdResponseBody) GetData() *GetTopicByIdResponseBodyData {
	return s.Data
}

func (s *GetTopicByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTopicByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicByIdResponseBody) SetCode(v string) *GetTopicByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetData(v *GetTopicByIdResponseBodyData) *GetTopicByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicByIdResponseBody) SetHttpStatusCode(v int32) *GetTopicByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetMessage(v string) *GetTopicByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetRequestId(v string) *GetTopicByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetSuccess(v bool) *GetTopicByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTopicByIdResponseBodyData struct {
	// example:
	//
	// 异步任务ID（自定义主题场景下使用）
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 创建用户ID（自定义主题场景下使用）
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 43
	HotValue *int64 `json:"HotValue,omitempty" xml:"HotValue,omitempty"`
	// example:
	//
	// 热榜ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// PENDING
	Status           *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureSummary []*GetTopicByIdResponseBodyDataStructureSummary `json:"StructureSummary,omitempty" xml:"StructureSummary,omitempty" type:"Repeated"`
	// example:
	//
	// 热榜摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 异步任务失败错误信息
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 14
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 主题唯一名称
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 热榜源，目前支持的热榜源: Toutiao：头条、Quark：夸克、Baidu：百度、Sina：新浪。Custom：自定义、Aggregation：热点话题榜
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 数据版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetTopicByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTopicByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBodyData) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *GetTopicByIdResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetTopicByIdResponseBodyData) GetHotValue() *int64 {
	return s.HotValue
}

func (s *GetTopicByIdResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetTopicByIdResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTopicByIdResponseBodyData) GetStructureSummary() []*GetTopicByIdResponseBodyDataStructureSummary {
	return s.StructureSummary
}

func (s *GetTopicByIdResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicByIdResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *GetTopicByIdResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *GetTopicByIdResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicByIdResponseBodyData) GetTopicSource() *string {
	return s.TopicSource
}

func (s *GetTopicByIdResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetTopicByIdResponseBodyData) SetAsyncTaskId(v string) *GetTopicByIdResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetCreateUser(v string) *GetTopicByIdResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetHotValue(v int64) *GetTopicByIdResponseBodyData {
	s.HotValue = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetId(v string) *GetTopicByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetStatus(v string) *GetTopicByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetStructureSummary(v []*GetTopicByIdResponseBodyDataStructureSummary) *GetTopicByIdResponseBodyData {
	s.StructureSummary = v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetSummary(v string) *GetTopicByIdResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTaskErrorMessage(v string) *GetTopicByIdResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTaskStatus(v int32) *GetTopicByIdResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTopic(v string) *GetTopicByIdResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTopicSource(v string) *GetTopicByIdResponseBodyData {
	s.TopicSource = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetVersion(v string) *GetTopicByIdResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) Validate() error {
	if s.StructureSummary != nil {
		for _, item := range s.StructureSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicByIdResponseBodyDataStructureSummary struct {
	DocList []*GetTopicByIdResponseBodyDataStructureSummaryDocList `json:"DocList,omitempty" xml:"DocList,omitempty" type:"Repeated"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTopicByIdResponseBodyDataStructureSummary) String() string {
	return dara.Prettify(s)
}

func (s GetTopicByIdResponseBodyDataStructureSummary) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) GetDocList() []*GetTopicByIdResponseBodyDataStructureSummaryDocList {
	return s.DocList
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) GetSummary() *string {
	return s.Summary
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) GetTitle() *string {
	return s.Title
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) SetDocList(v []*GetTopicByIdResponseBodyDataStructureSummaryDocList) *GetTopicByIdResponseBodyDataStructureSummary {
	s.DocList = v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) SetSummary(v string) *GetTopicByIdResponseBodyDataStructureSummary {
	s.Summary = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) SetTitle(v string) *GetTopicByIdResponseBodyDataStructureSummary {
	s.Title = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) Validate() error {
	if s.DocList != nil {
		for _, item := range s.DocList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTopicByIdResponseBodyDataStructureSummaryDocList struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTopicByIdResponseBodyDataStructureSummaryDocList) String() string {
	return dara.Prettify(s)
}

func (s GetTopicByIdResponseBodyDataStructureSummaryDocList) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) GetSource() *string {
	return s.Source
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) GetTitle() *string {
	return s.Title
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) GetUrl() *string {
	return s.Url
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) SetSource(v string) *GetTopicByIdResponseBodyDataStructureSummaryDocList {
	s.Source = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) SetTitle(v string) *GetTopicByIdResponseBodyDataStructureSummaryDocList {
	s.Title = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) SetUrl(v string) *GetTopicByIdResponseBodyDataStructureSummaryDocList {
	s.Url = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) Validate() error {
	return dara.Validate(s)
}
