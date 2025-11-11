// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotTopicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHotTopicsResponseBody
	GetCode() *string
	SetData(v []*ListHotTopicsResponseBodyData) *ListHotTopicsResponseBody
	GetData() []*ListHotTopicsResponseBodyData
	SetHttpStatusCode(v int32) *ListHotTopicsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListHotTopicsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListHotTopicsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListHotTopicsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListHotTopicsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListHotTopicsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListHotTopicsResponseBody
	GetTotalCount() *int32
}

type ListHotTopicsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListHotTopicsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 94
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHotTopicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHotTopicsResponseBody) GetData() []*ListHotTopicsResponseBodyData {
	return s.Data
}

func (s *ListHotTopicsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHotTopicsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotTopicsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotTopicsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotTopicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotTopicsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListHotTopicsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHotTopicsResponseBody) SetCode(v string) *ListHotTopicsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetData(v []*ListHotTopicsResponseBodyData) *ListHotTopicsResponseBody {
	s.Data = v
	return s
}

func (s *ListHotTopicsResponseBody) SetHttpStatusCode(v int32) *ListHotTopicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetMaxResults(v int32) *ListHotTopicsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetMessage(v string) *ListHotTopicsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetNextToken(v string) *ListHotTopicsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetRequestId(v string) *ListHotTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetSuccess(v bool) *ListHotTopicsResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetTotalCount(v int32) *ListHotTopicsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHotTopicsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotTopicsResponseBodyData struct {
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
	// 61
	HotValue *int64 `json:"HotValue,omitempty" xml:"HotValue,omitempty"`
	// example:
	//
	// 热榜ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// FAILED
	Status           *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureSummary []*ListHotTopicsResponseBodyDataStructureSummary `json:"StructureSummary,omitempty" xml:"StructureSummary,omitempty" type:"Repeated"`
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
	// 26
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

func (s ListHotTopicsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBodyData) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *ListHotTopicsResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListHotTopicsResponseBodyData) GetHotValue() *int64 {
	return s.HotValue
}

func (s *ListHotTopicsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListHotTopicsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListHotTopicsResponseBodyData) GetStructureSummary() []*ListHotTopicsResponseBodyDataStructureSummary {
	return s.StructureSummary
}

func (s *ListHotTopicsResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *ListHotTopicsResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *ListHotTopicsResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListHotTopicsResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *ListHotTopicsResponseBodyData) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListHotTopicsResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListHotTopicsResponseBodyData) SetAsyncTaskId(v string) *ListHotTopicsResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetCreateUser(v string) *ListHotTopicsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetHotValue(v int64) *ListHotTopicsResponseBodyData {
	s.HotValue = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetId(v string) *ListHotTopicsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetStatus(v string) *ListHotTopicsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetStructureSummary(v []*ListHotTopicsResponseBodyDataStructureSummary) *ListHotTopicsResponseBodyData {
	s.StructureSummary = v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetSummary(v string) *ListHotTopicsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTaskErrorMessage(v string) *ListHotTopicsResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTaskStatus(v int32) *ListHotTopicsResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTopic(v string) *ListHotTopicsResponseBodyData {
	s.Topic = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTopicSource(v string) *ListHotTopicsResponseBodyData {
	s.TopicSource = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetVersion(v string) *ListHotTopicsResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) Validate() error {
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

type ListHotTopicsResponseBodyDataStructureSummary struct {
	DocList []*ListHotTopicsResponseBodyDataStructureSummaryDocList `json:"DocList,omitempty" xml:"DocList,omitempty" type:"Repeated"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListHotTopicsResponseBodyDataStructureSummary) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicsResponseBodyDataStructureSummary) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) GetDocList() []*ListHotTopicsResponseBodyDataStructureSummaryDocList {
	return s.DocList
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) GetSummary() *string {
	return s.Summary
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) GetTitle() *string {
	return s.Title
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) SetDocList(v []*ListHotTopicsResponseBodyDataStructureSummaryDocList) *ListHotTopicsResponseBodyDataStructureSummary {
	s.DocList = v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) SetSummary(v string) *ListHotTopicsResponseBodyDataStructureSummary {
	s.Summary = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) SetTitle(v string) *ListHotTopicsResponseBodyDataStructureSummary {
	s.Title = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) Validate() error {
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

type ListHotTopicsResponseBodyDataStructureSummaryDocList struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// xxxxx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotTopicsResponseBodyDataStructureSummaryDocList) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicsResponseBodyDataStructureSummaryDocList) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) GetSource() *string {
	return s.Source
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) GetTitle() *string {
	return s.Title
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) GetUrl() *string {
	return s.Url
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) SetSource(v string) *ListHotTopicsResponseBodyDataStructureSummaryDocList {
	s.Source = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) SetTitle(v string) *ListHotTopicsResponseBodyDataStructureSummaryDocList {
	s.Title = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) SetUrl(v string) *ListHotTopicsResponseBodyDataStructureSummaryDocList {
	s.Url = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) Validate() error {
	return dara.Validate(s)
}
