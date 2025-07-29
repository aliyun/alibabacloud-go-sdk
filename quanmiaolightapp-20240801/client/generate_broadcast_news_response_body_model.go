// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateBroadcastNewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateBroadcastNewsResponseBody
	GetCode() *string
	SetData(v *GenerateBroadcastNewsResponseBodyData) *GenerateBroadcastNewsResponseBody
	GetData() *GenerateBroadcastNewsResponseBodyData
	SetHttpStatusCode(v int32) *GenerateBroadcastNewsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateBroadcastNewsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateBroadcastNewsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateBroadcastNewsResponseBody
	GetSuccess() *bool
}

type GenerateBroadcastNewsResponseBody struct {
	// example:
	//
	// xx
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *GenerateBroadcastNewsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GenerateBroadcastNewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateBroadcastNewsResponseBody) GetData() *GenerateBroadcastNewsResponseBodyData {
	return s.Data
}

func (s *GenerateBroadcastNewsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateBroadcastNewsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateBroadcastNewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateBroadcastNewsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateBroadcastNewsResponseBody) SetCode(v string) *GenerateBroadcastNewsResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetData(v *GenerateBroadcastNewsResponseBodyData) *GenerateBroadcastNewsResponseBody {
	s.Data = v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetHttpStatusCode(v int32) *GenerateBroadcastNewsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetMessage(v string) *GenerateBroadcastNewsResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetRequestId(v string) *GenerateBroadcastNewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) SetSuccess(v bool) *GenerateBroadcastNewsResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateBroadcastNewsResponseBodyData struct {
	HotTopicSummaries []*GenerateBroadcastNewsResponseBodyDataHotTopicSummaries `json:"hotTopicSummaries,omitempty" xml:"hotTopicSummaries,omitempty" type:"Repeated"`
	// example:
	//
	// 2bb0ea82dafd48a8817fadc4c90e2b52
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string                                     `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Text   *string                                     `json:"text,omitempty" xml:"text,omitempty"`
	Usage  *GenerateBroadcastNewsResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GenerateBroadcastNewsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyData) GetHotTopicSummaries() []*GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	return s.HotTopicSummaries
}

func (s *GenerateBroadcastNewsResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GenerateBroadcastNewsResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GenerateBroadcastNewsResponseBodyData) GetText() *string {
	return s.Text
}

func (s *GenerateBroadcastNewsResponseBodyData) GetUsage() *GenerateBroadcastNewsResponseBodyDataUsage {
	return s.Usage
}

func (s *GenerateBroadcastNewsResponseBodyData) SetHotTopicSummaries(v []*GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) *GenerateBroadcastNewsResponseBodyData {
	s.HotTopicSummaries = v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetSessionId(v string) *GenerateBroadcastNewsResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetTaskId(v string) *GenerateBroadcastNewsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetText(v string) *GenerateBroadcastNewsResponseBodyData {
	s.Text = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) SetUsage(v *GenerateBroadcastNewsResponseBodyDataUsage) *GenerateBroadcastNewsResponseBodyData {
	s.Usage = v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GenerateBroadcastNewsResponseBodyDataHotTopicSummaries struct {
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_08
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 1000000
	HotValue *float64 `json:"hotValue,omitempty" xml:"hotValue,omitempty"`
	// example:
	//
	// 1458tb3bjo7531kap42a
	Id     *string                                                         `json:"id,omitempty" xml:"id,omitempty"`
	Images []*GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// example:
	//
	// xxx
	TextSummary *string `json:"textSummary,omitempty" xml:"textSummary,omitempty"`
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) String() string {
	return dara.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GetCategory() *string {
	return s.Category
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GetHotTopic() *string {
	return s.HotTopic
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GetHotValue() *float64 {
	return s.HotValue
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GetId() *string {
	return s.Id
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GetImages() []*GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages {
	return s.Images
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) GetTextSummary() *string {
	return s.TextSummary
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetCategory(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.Category = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetHotTopic(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.HotTopic = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetHotTopicVersion(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.HotTopicVersion = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetHotValue(v float64) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.HotValue = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetId(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.Id = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetImages(v []*GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.Images = v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) SetTextSummary(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries {
	s.TextSummary = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummaries) Validate() error {
	return dara.Validate(s)
}

type GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages struct {
	// example:
	//
	// http://xxx.com/xxx.jpeg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) String() string {
	return dara.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) GetUrl() *string {
	return s.Url
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) SetUrl(v string) *GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages {
	s.Url = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataHotTopicSummariesImages) Validate() error {
	return dara.Validate(s)
}

type GenerateBroadcastNewsResponseBodyDataUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 2
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 3
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GenerateBroadcastNewsResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s GenerateBroadcastNewsResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) SetInputTokens(v int64) *GenerateBroadcastNewsResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) SetOutputTokens(v int64) *GenerateBroadcastNewsResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) SetTotalTokens(v int64) *GenerateBroadcastNewsResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *GenerateBroadcastNewsResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
