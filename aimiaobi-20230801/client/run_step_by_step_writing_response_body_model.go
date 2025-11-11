// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStepByStepWritingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunStepByStepWritingResponseBodyHeader) *RunStepByStepWritingResponseBody
	GetHeader() *RunStepByStepWritingResponseBodyHeader
	SetPayload(v *RunStepByStepWritingResponseBodyPayload) *RunStepByStepWritingResponseBody
	GetPayload() *RunStepByStepWritingResponseBodyPayload
	SetRequestId(v string) *RunStepByStepWritingResponseBody
	GetRequestId() *string
}

type RunStepByStepWritingResponseBody struct {
	Header  *RunStepByStepWritingResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunStepByStepWritingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunStepByStepWritingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBody) GetHeader() *RunStepByStepWritingResponseBodyHeader {
	return s.Header
}

func (s *RunStepByStepWritingResponseBody) GetPayload() *RunStepByStepWritingResponseBodyPayload {
	return s.Payload
}

func (s *RunStepByStepWritingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunStepByStepWritingResponseBody) SetHeader(v *RunStepByStepWritingResponseBodyHeader) *RunStepByStepWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunStepByStepWritingResponseBody) SetPayload(v *RunStepByStepWritingResponseBodyPayload) *RunStepByStepWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunStepByStepWritingResponseBody) SetRequestId(v string) *RunStepByStepWritingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunStepByStepWritingResponseBody) Validate() error {
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

type RunStepByStepWritingResponseBodyHeader struct {
	// example:
	//
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-started
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 全链路ID
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunStepByStepWritingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunStepByStepWritingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunStepByStepWritingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunStepByStepWritingResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunStepByStepWritingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunStepByStepWritingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunStepByStepWritingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunStepByStepWritingResponseBodyHeader) SetErrorCode(v string) *RunStepByStepWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyHeader) SetErrorMessage(v string) *RunStepByStepWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyHeader) SetEvent(v string) *RunStepByStepWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyHeader) SetOriginSessionId(v string) *RunStepByStepWritingResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyHeader) SetSessionId(v string) *RunStepByStepWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyHeader) SetTaskId(v string) *RunStepByStepWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyHeader) SetTraceId(v string) *RunStepByStepWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingResponseBodyPayload struct {
	Output *RunStepByStepWritingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunStepByStepWritingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunStepByStepWritingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyPayload) GetOutput() *RunStepByStepWritingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunStepByStepWritingResponseBodyPayload) GetUsage() *RunStepByStepWritingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunStepByStepWritingResponseBodyPayload) SetOutput(v *RunStepByStepWritingResponseBodyPayloadOutput) *RunStepByStepWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayload) SetUsage(v *RunStepByStepWritingResponseBodyPayloadUsage) *RunStepByStepWritingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayload) Validate() error {
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

type RunStepByStepWritingResponseBodyPayloadOutput struct {
	Articles    []*RunStepByStepWritingResponseBodyPayloadOutputArticles  `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	ExtraOutput *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput `json:"ExtraOutput,omitempty" xml:"ExtraOutput,omitempty" type:"Struct"`
	// example:
	//
	// 文章精排之后的片段
	MiniDoc []*string `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
	// example:
	//
	// 大模型改变世界
	SearchQuery *string `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunStepByStepWritingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) GetArticles() []*RunStepByStepWritingResponseBodyPayloadOutputArticles {
	return s.Articles
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) GetExtraOutput() *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput {
	return s.ExtraOutput
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) GetMiniDoc() []*string {
	return s.MiniDoc
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) SetArticles(v []*RunStepByStepWritingResponseBodyPayloadOutputArticles) *RunStepByStepWritingResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) SetExtraOutput(v *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) *RunStepByStepWritingResponseBodyPayloadOutput {
	s.ExtraOutput = v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) SetMiniDoc(v []*string) *RunStepByStepWritingResponseBodyPayloadOutput {
	s.MiniDoc = v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) SetSearchQuery(v string) *RunStepByStepWritingResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) SetText(v string) *RunStepByStepWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutput) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExtraOutput != nil {
		if err := s.ExtraOutput.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunStepByStepWritingResponseBodyPayloadOutputArticles struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文档-自定义的唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// f1da53894e784759946d22e2cb2b522a
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// http://www.example.com
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// example:
	//
	// 2024-09-10 14:17:53
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 央视网
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标签
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunStepByStepWritingResponseBodyPayloadOutputArticles) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetAuthor() *string {
	return s.Author
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetContent() *string {
	return s.Content
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetDocId() *string {
	return s.DocId
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetMediaUrl() *string {
	return s.MediaUrl
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetSource() *string {
	return s.Source
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetTag() *string {
	return s.Tag
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetTitle() *string {
	return s.Title
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) GetUrl() *string {
	return s.Url
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetAuthor(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.Author = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetContent(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetDocId(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.DocId = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetDocUuid(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.DocUuid = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetMediaUrl(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.MediaUrl = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetSource(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetSummary(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetTag(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.Tag = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetTitle(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) SetUrl(v string) *RunStepByStepWritingResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputArticles) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingResponseBodyPayloadOutputExtraOutput struct {
	Summarization []*string `json:"summarization,omitempty" xml:"summarization,omitempty" type:"Repeated"`
}

func (s RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) GetSummarization() []*string {
	return s.Summarization
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) SetSummarization(v []*string) *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput {
	s.Summarization = v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) Validate() error {
	return dara.Validate(s)
}

type RunStepByStepWritingResponseBodyPayloadUsage struct {
	// example:
	//
	// 65
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 80
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 32
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunStepByStepWritingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunStepByStepWritingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunStepByStepWritingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunStepByStepWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunStepByStepWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunStepByStepWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunStepByStepWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
