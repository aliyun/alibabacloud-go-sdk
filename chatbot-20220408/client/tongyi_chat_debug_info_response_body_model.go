// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTongyiChatDebugInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswerInfo(v *TongyiChatDebugInfoResponseBodyAnswerInfo) *TongyiChatDebugInfoResponseBody
	GetAnswerInfo() *TongyiChatDebugInfoResponseBodyAnswerInfo
	SetMessageId(v string) *TongyiChatDebugInfoResponseBody
	GetMessageId() *string
	SetPipeline(v []*TongyiChatDebugInfoResponseBodyPipeline) *TongyiChatDebugInfoResponseBody
	GetPipeline() []*TongyiChatDebugInfoResponseBodyPipeline
	SetRequestId(v string) *TongyiChatDebugInfoResponseBody
	GetRequestId() *string
}

type TongyiChatDebugInfoResponseBody struct {
	AnswerInfo *TongyiChatDebugInfoResponseBodyAnswerInfo `json:"AnswerInfo,omitempty" xml:"AnswerInfo,omitempty" type:"Struct"`
	// The ID of the response message in the current session.
	//
	// example:
	//
	// 2828708A-2C7A-1BAE-B810-87DB9DA9C661
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The array of nodes that constitute the Q\\&A workflow.
	Pipeline []*TongyiChatDebugInfoResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E3E5C779-A630-45AC-B0F2-A4506A4212F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TongyiChatDebugInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBody) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBody) GetAnswerInfo() *TongyiChatDebugInfoResponseBodyAnswerInfo {
	return s.AnswerInfo
}

func (s *TongyiChatDebugInfoResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *TongyiChatDebugInfoResponseBody) GetPipeline() []*TongyiChatDebugInfoResponseBodyPipeline {
	return s.Pipeline
}

func (s *TongyiChatDebugInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TongyiChatDebugInfoResponseBody) SetAnswerInfo(v *TongyiChatDebugInfoResponseBodyAnswerInfo) *TongyiChatDebugInfoResponseBody {
	s.AnswerInfo = v
	return s
}

func (s *TongyiChatDebugInfoResponseBody) SetMessageId(v string) *TongyiChatDebugInfoResponseBody {
	s.MessageId = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBody) SetPipeline(v []*TongyiChatDebugInfoResponseBodyPipeline) *TongyiChatDebugInfoResponseBody {
	s.Pipeline = v
	return s
}

func (s *TongyiChatDebugInfoResponseBody) SetRequestId(v string) *TongyiChatDebugInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBody) Validate() error {
	if s.AnswerInfo != nil {
		if err := s.AnswerInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Pipeline != nil {
		for _, item := range s.Pipeline {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TongyiChatDebugInfoResponseBodyAnswerInfo struct {
	AnswerReferenceInfo *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo `json:"AnswerReferenceInfo,omitempty" xml:"AnswerReferenceInfo,omitempty" type:"Struct"`
	MessageBody         *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody         `json:"MessageBody,omitempty" xml:"MessageBody,omitempty" type:"Struct"`
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfo) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfo) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfo) GetAnswerReferenceInfo() *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo {
	return s.AnswerReferenceInfo
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfo) GetMessageBody() *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody {
	return s.MessageBody
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfo) SetAnswerReferenceInfo(v *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo) *TongyiChatDebugInfoResponseBodyAnswerInfo {
	s.AnswerReferenceInfo = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfo) SetMessageBody(v *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) *TongyiChatDebugInfoResponseBodyAnswerInfo {
	s.MessageBody = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfo) Validate() error {
	if s.AnswerReferenceInfo != nil {
		if err := s.AnswerReferenceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MessageBody != nil {
		if err := s.MessageBody.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo struct {
	ItemList []*TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo) GetItemList() []*TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList {
	return s.ItemList
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo) SetItemList(v []*TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo {
	s.ItemList = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo) Validate() error {
	if s.ItemList != nil {
		for _, item := range s.ItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList struct {
	ContentType  *string     `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	DataSource   *string     `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	Id           *string     `json:"Id,omitempty" xml:"Id,omitempty"`
	Number       *int32      `json:"Number,omitempty" xml:"Number,omitempty"`
	ReferenceExt interface{} `json:"ReferenceExt,omitempty" xml:"ReferenceExt,omitempty"`
	Title        *string     `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) GetContentType() *string {
	return s.ContentType
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) GetDataSource() *string {
	return s.DataSource
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) GetId() *string {
	return s.Id
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) GetNumber() *int32 {
	return s.Number
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) GetReferenceExt() interface{} {
	return s.ReferenceExt
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) GetTitle() *string {
	return s.Title
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) SetContentType(v string) *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList {
	s.ContentType = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) SetDataSource(v string) *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList {
	s.DataSource = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) SetId(v string) *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList {
	s.Id = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) SetNumber(v int32) *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList {
	s.Number = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) SetReferenceExt(v interface{}) *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList {
	s.ReferenceExt = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) SetTitle(v string) *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList {
	s.Title = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfoItemList) Validate() error {
	return dara.Validate(s)
}

type TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody struct {
	Commands          interface{}                                                            `json:"Commands,omitempty" xml:"Commands,omitempty"`
	DirectMessageBody *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody `json:"DirectMessageBody,omitempty" xml:"DirectMessageBody,omitempty" type:"Struct"`
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) GetCommands() interface{} {
	return s.Commands
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) GetDirectMessageBody() *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody {
	return s.DirectMessageBody
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) SetCommands(v interface{}) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody {
	s.Commands = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) SetDirectMessageBody(v *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody {
	s.DirectMessageBody = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody) Validate() error {
	if s.DirectMessageBody != nil {
		if err := s.DirectMessageBody.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody struct {
	ContentType         *string                                                                              `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	TransitionList      []*string                                                                            `json:"TransitionList,omitempty" xml:"TransitionList,omitempty" type:"Repeated"`
	RelatedQuestionList []*string                                                                            `json:"relatedQuestionList,omitempty" xml:"relatedQuestionList,omitempty" type:"Repeated"`
	SentenceList        []*TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) GetContentType() *string {
	return s.ContentType
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) GetTransitionList() []*string {
	return s.TransitionList
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) GetRelatedQuestionList() []*string {
	return s.RelatedQuestionList
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) GetSentenceList() []*TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList {
	return s.SentenceList
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) SetContentType(v string) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody {
	s.ContentType = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) SetTransitionList(v []*string) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody {
	s.TransitionList = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) SetRelatedQuestionList(v []*string) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody {
	s.RelatedQuestionList = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) SetSentenceList(v []*TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody {
	s.SentenceList = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) Validate() error {
	if s.SentenceList != nil {
		for _, item := range s.SentenceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ReferNumber *int32  `json:"ReferNumber,omitempty" xml:"ReferNumber,omitempty"`
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) GetContent() *string {
	return s.Content
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) GetReferNumber() *int32 {
	return s.ReferNumber
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) SetContent(v string) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList {
	s.Content = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) SetReferNumber(v int32) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList {
	s.ReferNumber = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList) Validate() error {
	return dara.Validate(s)
}

type TongyiChatDebugInfoResponseBodyPipeline struct {
	// The input data for the node.
	//
	// example:
	//
	// 用户问句：转人工\\n命中规则：[转]人工[客服|服务|坐席]
	Input interface{} `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the strategy. Possible values include:
	//
	// - FAQ
	//
	// - Hit Keywords
	//
	// - Global Sensitive Words
	//
	// This parameter is returned only when `NodeType` is set to `system_strategy`.
	//
	// example:
	//
	// 关键词转人工
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node type. Valid values:
	//
	// - **system_strategy**: system strategy.
	//
	// - **rewrite_query**: retrieval query.
	//
	// - **invoke_llm**: LLM invocation.
	//
	// - **invoke_tools**: tool invocation.
	//
	// example:
	//
	// system_strategy
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The output data from the node.
	//
	// example:
	//
	// commands:{\\"sysToAgent\\":\\"{\\\\\\"skillGroup\\\\\\":\\\\\\"\\\\\\",\\\\\\"ext\\\\\\":\\\\\\"\\\\\\",\\\\\\"toAgentReason\\\\\\":\\\\\\"HitKeywords\\\\\\"}\\"}\\nresponse:正在为您转接人工客服
	Output interface{} `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s TongyiChatDebugInfoResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetInput() interface{} {
	return s.Input
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetName() *string {
	return s.Name
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetNodeType() *string {
	return s.NodeType
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) GetOutput() interface{} {
	return s.Output
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetInput(v interface{}) *TongyiChatDebugInfoResponseBodyPipeline {
	s.Input = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetName(v string) *TongyiChatDebugInfoResponseBodyPipeline {
	s.Name = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetNodeType(v string) *TongyiChatDebugInfoResponseBodyPipeline {
	s.NodeType = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) SetOutput(v interface{}) *TongyiChatDebugInfoResponseBodyPipeline {
	s.Output = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyPipeline) Validate() error {
	return dara.Validate(s)
}
