// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocQaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunDocQaResponseBodyHeader) *RunDocQaResponseBody
	GetHeader() *RunDocQaResponseBodyHeader
	SetPayload(v *RunDocQaResponseBodyPayload) *RunDocQaResponseBody
	GetPayload() *RunDocQaResponseBodyPayload
	SetRequestId(v string) *RunDocQaResponseBody
	GetRequestId() *string
}

type RunDocQaResponseBody struct {
	// response header
	Header *RunDocQaResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// response body
	Payload *RunDocQaResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunDocQaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBody) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBody) GetHeader() *RunDocQaResponseBodyHeader {
	return s.Header
}

func (s *RunDocQaResponseBody) GetPayload() *RunDocQaResponseBodyPayload {
	return s.Payload
}

func (s *RunDocQaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocQaResponseBody) SetHeader(v *RunDocQaResponseBodyHeader) *RunDocQaResponseBody {
	s.Header = v
	return s
}

func (s *RunDocQaResponseBody) SetPayload(v *RunDocQaResponseBodyPayload) *RunDocQaResponseBody {
	s.Payload = v
	return s
}

func (s *RunDocQaResponseBody) SetRequestId(v string) *RunDocQaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDocQaResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunDocQaResponseBodyHeader struct {
	// error code
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// error message
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// management event
	//
	// example:
	//
	// task-started
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// Description of the management event
	//
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// Session ID
	//
	// example:
	//
	// f5517ee8-dbec-4dc8-bd0a-af084b5e3db1
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Job ID
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// End-to-end trace ID
	//
	// example:
	//
	// 46e5c2b5-0877-4f09-bd91-ab0cf314e48b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDocQaResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunDocQaResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunDocQaResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDocQaResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunDocQaResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocQaResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDocQaResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDocQaResponseBodyHeader) SetErrorCode(v string) *RunDocQaResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunDocQaResponseBodyHeader) SetErrorMessage(v string) *RunDocQaResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunDocQaResponseBodyHeader) SetEvent(v string) *RunDocQaResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDocQaResponseBodyHeader) SetEventInfo(v string) *RunDocQaResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunDocQaResponseBodyHeader) SetSessionId(v string) *RunDocQaResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDocQaResponseBodyHeader) SetTaskId(v string) *RunDocQaResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDocQaResponseBodyHeader) SetTraceId(v string) *RunDocQaResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDocQaResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDocQaResponseBodyPayload struct {
	// Outputs
	Output *RunDocQaResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// token usage
	Usage *RunDocQaResponseBodyPayloadUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunDocQaResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyPayload) GetOutput() *RunDocQaResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDocQaResponseBodyPayload) GetUsage() *RunDocQaResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunDocQaResponseBodyPayload) SetOutput(v *RunDocQaResponseBodyPayloadOutput) *RunDocQaResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDocQaResponseBodyPayload) SetUsage(v *RunDocQaResponseBodyPayloadUsage) *RunDocQaResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunDocQaResponseBodyPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunDocQaResponseBodyPayloadOutput struct {
	// Content of the response
	//
	// example:
	//
	// 回答内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Response content after intervention
	//
	// example:
	//
	// 干预后的回答内容
	InterveneContent *string `json:"InterveneContent,omitempty" xml:"InterveneContent,omitempty"`
	// Indicates whether the request is rejected
	//
	// example:
	//
	// false
	IsReject *bool `json:"IsReject,omitempty" xml:"IsReject,omitempty"`
	// List of multimodal resource information
	MediaUrlList []*RunDocQaResponseBodyPayloadOutputMediaUrlList `json:"MediaUrlList,omitempty" xml:"MediaUrlList,omitempty" type:"Repeated"`
	// Array of recommended content
	Recommends []*RunDocQaResponseBodyPayloadOutputRecommends `json:"Recommends,omitempty" xml:"Recommends,omitempty" type:"Repeated"`
	// Array of sources for the response content
	References []*RunDocQaResponseBodyPayloadOutputReferences `json:"References,omitempty" xml:"References,omitempty" type:"Repeated"`
}

func (s RunDocQaResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunDocQaResponseBodyPayloadOutput) GetInterveneContent() *string {
	return s.InterveneContent
}

func (s *RunDocQaResponseBodyPayloadOutput) GetIsReject() *bool {
	return s.IsReject
}

func (s *RunDocQaResponseBodyPayloadOutput) GetMediaUrlList() []*RunDocQaResponseBodyPayloadOutputMediaUrlList {
	return s.MediaUrlList
}

func (s *RunDocQaResponseBodyPayloadOutput) GetRecommends() []*RunDocQaResponseBodyPayloadOutputRecommends {
	return s.Recommends
}

func (s *RunDocQaResponseBodyPayloadOutput) GetReferences() []*RunDocQaResponseBodyPayloadOutputReferences {
	return s.References
}

func (s *RunDocQaResponseBodyPayloadOutput) SetContent(v string) *RunDocQaResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutput) SetInterveneContent(v string) *RunDocQaResponseBodyPayloadOutput {
	s.InterveneContent = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutput) SetIsReject(v bool) *RunDocQaResponseBodyPayloadOutput {
	s.IsReject = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutput) SetMediaUrlList(v []*RunDocQaResponseBodyPayloadOutputMediaUrlList) *RunDocQaResponseBodyPayloadOutput {
	s.MediaUrlList = v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutput) SetRecommends(v []*RunDocQaResponseBodyPayloadOutputRecommends) *RunDocQaResponseBodyPayloadOutput {
	s.Recommends = v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutput) SetReferences(v []*RunDocQaResponseBodyPayloadOutputReferences) *RunDocQaResponseBodyPayloadOutput {
	s.References = v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutput) Validate() error {
	if s.MediaUrlList != nil {
		for _, item := range s.MediaUrlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Recommends != nil {
		for _, item := range s.Recommends {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.References != nil {
		for _, item := range s.References {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunDocQaResponseBodyPayloadOutputMediaUrlList struct {
	// Array of related video time information
	ClipInfos []*RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos `json:"ClipInfos,omitempty" xml:"ClipInfos,omitempty" type:"Repeated"`
	// File URL
	//
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/2775676850/O1CN01kdeffE20TM0E7wvpq_!!2775676850.jpg_q60.jpg
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Media asset type
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s RunDocQaResponseBodyPayloadOutputMediaUrlList) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyPayloadOutputMediaUrlList) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlList) GetClipInfos() []*RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos {
	return s.ClipInfos
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlList) GetMediaType() *string {
	return s.MediaType
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlList) SetClipInfos(v []*RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) *RunDocQaResponseBodyPayloadOutputMediaUrlList {
	s.ClipInfos = v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlList) SetFileUrl(v string) *RunDocQaResponseBodyPayloadOutputMediaUrlList {
	s.FileUrl = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlList) SetMediaType(v string) *RunDocQaResponseBodyPayloadOutputMediaUrlList {
	s.MediaType = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlList) Validate() error {
	if s.ClipInfos != nil {
		for _, item := range s.ClipInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos struct {
	// Start time of the video segment
	//
	// example:
	//
	// 0
	From *float64 `json:"From,omitempty" xml:"From,omitempty"`
	// End time of the video segment
	//
	// example:
	//
	// 30
	To *float64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) GetFrom() *float64 {
	return s.From
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) GetTo() *float64 {
	return s.To
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) SetFrom(v float64) *RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos {
	s.From = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) SetTo(v float64) *RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos {
	s.To = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputMediaUrlListClipInfos) Validate() error {
	return dara.Validate(s)
}

type RunDocQaResponseBodyPayloadOutputRecommends struct {
	// Title of the recommended content
	//
	// example:
	//
	// 标题内容
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// URL of the recommended content
	//
	// example:
	//
	// 推荐内容url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunDocQaResponseBodyPayloadOutputRecommends) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyPayloadOutputRecommends) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyPayloadOutputRecommends) GetTitle() *string {
	return s.Title
}

func (s *RunDocQaResponseBodyPayloadOutputRecommends) GetUrl() *string {
	return s.Url
}

func (s *RunDocQaResponseBodyPayloadOutputRecommends) SetTitle(v string) *RunDocQaResponseBodyPayloadOutputRecommends {
	s.Title = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputRecommends) SetUrl(v string) *RunDocQaResponseBodyPayloadOutputRecommends {
	s.Url = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputRecommends) Validate() error {
	return dara.Validate(s)
}

type RunDocQaResponseBodyPayloadOutputReferences struct {
	// Published At
	//
	// example:
	//
	// 2024-10-08 18:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// Source
	//
	// example:
	//
	// 新浪新闻
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Source docId
	//
	// example:
	//
	// 123456
	SourceDocId *string `json:"SourceDocId,omitempty" xml:"SourceDocId,omitempty"`
	// Title of the associated content
	//
	// example:
	//
	// 标题内容
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Article URL
	//
	// example:
	//
	// http://xxxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunDocQaResponseBodyPayloadOutputReferences) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyPayloadOutputReferences) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) GetPubTime() *string {
	return s.PubTime
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) GetSource() *string {
	return s.Source
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) GetSourceDocId() *string {
	return s.SourceDocId
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) GetTitle() *string {
	return s.Title
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) GetUrl() *string {
	return s.Url
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) SetPubTime(v string) *RunDocQaResponseBodyPayloadOutputReferences {
	s.PubTime = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) SetSource(v string) *RunDocQaResponseBodyPayloadOutputReferences {
	s.Source = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) SetSourceDocId(v string) *RunDocQaResponseBodyPayloadOutputReferences {
	s.SourceDocId = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) SetTitle(v string) *RunDocQaResponseBodyPayloadOutputReferences {
	s.Title = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) SetUrl(v string) *RunDocQaResponseBodyPayloadOutputReferences {
	s.Url = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadOutputReferences) Validate() error {
	return dara.Validate(s)
}

type RunDocQaResponseBodyPayloadUsage struct {
	// Quantity of input tokens
	//
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// Number of tokens used in the output
	//
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// Total number of tokens
	//
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunDocQaResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunDocQaResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunDocQaResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunDocQaResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunDocQaResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunDocQaResponseBodyPayloadUsage) SetInputTokens(v int64) *RunDocQaResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunDocQaResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunDocQaResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunDocQaResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
