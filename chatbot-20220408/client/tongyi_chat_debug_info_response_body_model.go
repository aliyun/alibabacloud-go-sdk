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
	SetSessionId(v string) *TongyiChatDebugInfoResponseBody
	GetSessionId() *string
}

type TongyiChatDebugInfoResponseBody struct {
	// The answer information corresponding to the conversation.
	AnswerInfo *TongyiChatDebugInfoResponseBodyAnswerInfo `json:"AnswerInfo,omitempty" xml:"AnswerInfo,omitempty" type:"Struct"`
	// The ID of the response message for the current session.
	//
	// example:
	//
	// 2828708A-2C7A-1BAE-B810-87DB9DA9C661
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The information about the entire Q&A process.
	//
	// example:
	//
	// For examples, see the child elements
	Pipeline []*TongyiChatDebugInfoResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// E3E5C779-A630-45AC-B0F2-A4506A4212F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the session to which the current response message belongs.
	//
	// example:
	//
	// 19856284-0ba1-4da3-9a5a-05db2bdf296e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
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

func (s *TongyiChatDebugInfoResponseBody) GetSessionId() *string {
	return s.SessionId
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

func (s *TongyiChatDebugInfoResponseBody) SetSessionId(v string) *TongyiChatDebugInfoResponseBody {
	s.SessionId = &v
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
	// The referenced knowledge information of the matched answer.
	AnswerReferenceInfo *TongyiChatDebugInfoResponseBodyAnswerInfoAnswerReferenceInfo `json:"AnswerReferenceInfo,omitempty" xml:"AnswerReferenceInfo,omitempty" type:"Struct"`
	// The message body.
	MessageBody *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBody `json:"MessageBody,omitempty" xml:"MessageBody,omitempty" type:"Struct"`
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
	// The list of referenced knowledge entries.
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
	// The type of the referenced text.
	//
	// example:
	//
	// PLAIN_TEXT/RICH_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The source of the referenced knowledge.
	//
	// example:
	//
	// FLOW/DOC/FAQ/FAQ_DIRECT
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The ID of the referenced knowledge.
	//
	// example:
	//
	// 611926
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number that corresponds one-to-one with the ReferNumber in AnswerInfo.MessageBody.DirectMessageBody.SentenceList.
	//
	// example:
	//
	// 1
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// The extension information specific to the knowledge type.
	//
	// example:
	//
	// {
	//
	//             "CollectedParams": {
	//
	//               "ReturnMaster": "true"
	//
	//             },
	//
	//             "Edges": [],
	//
	//             "Nodes": [
	//
	//               {
	//
	//                 "InvokedApis": [
	//
	//                   {
	//
	//                     "ApiName": "intentionRepeatCallQry"
	//
	//                   }
	//
	//                 ],
	//
	//                 "OptionalParams": [
	//
	//                   "mobile",
	//
	//                   "intention"
	//
	//                 ],
	//
	//                 "Id": "22874cb7-c253-4832-a38d-7113b15ef7fc",
	//
	//                 "RequiredParams": [],
	//
	//                 "Name": "Check for repeated contacts"
	//
	//               },
	//
	//               {
	//
	//                 "InvokedApis": [
	//
	//                   {
	//
	//                     "ApiName": "complaintAllChannelLatest"
	//
	//                   }
	//
	//                 ],
	//
	//                 "OptionalParams": [
	//
	//                   "mobile",
	//
	//                   "billCode"
	//
	//                 ],
	//
	//                 "Id": "071b082e-ce76-47f9-8b8d-f29e0dd1015c",
	//
	//                 "RequiredParams": [],
	//
	//                 "Name": "Query ticket information"
	//
	//               },
	//
	//               {
	//
	//                 "Id": "db02f361-61ee-4e30-bda0-2fd62cf4ca46",
	//
	//                 "Name": "Logical judgment_35a6"
	//
	//               },
	//
	//               {
	//
	//                 "OptionalParams": [],
	//
	//                 "Id": "97f68b60-2973-4d10-8e42-1cbfcbb15cf6",
	//
	//                 "RequiredParams": [
	//
	//                   "returnMaster"
	//
	//                 ],
	//
	//                 "Name": "Whether to return to the main flow"
	//
	//               },
	//
	//               {
	//
	//                 "Id": "8167e4ec-a589-42de-97c2-271672b26ffe",
	//
	//                 "Name": "Return to the main flow"
	//
	//               }
	//
	//             ]
	//
	//           }
	ReferenceExt interface{} `json:"ReferenceExt,omitempty" xml:"ReferenceExt,omitempty"`
	// The title of the referenced knowledge.
	//
	// example:
	//
	// Is it possible to make the payment when the contract entity and the paying entity are different?
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// The command parameters, such as the skill group for the transfer-to-agent command.
	//
	// example:
	//
	// {
	//
	//         "SysToAgent": "{\\"skillGroup\\":\\"\\",\\"ext\\":\\"\\",\\"toAgentReason\\":\\"HitKeywords\\"}"
	//
	//       }
	Commands interface{} `json:"Commands,omitempty" xml:"Commands,omitempty"`
	// The message body that contains the Q&A results, including the answer, text type, recommended questions, and matched system type.
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
	// The answer type.
	//
	// example:
	//
	// RICH_TEXT/PLAIN_TEXT
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The matched system Q&A policy.
	//
	// example:
	//
	// SENSITIVE_WORD/NO_ANSWER/COMMANDS
	HitSystemAskConfig *string `json:"HitSystemAskConfig,omitempty" xml:"HitSystemAskConfig,omitempty"`
	// The response time (RT) debugging information for the current Q&A.
	RtDetail *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail `json:"RtDetail,omitempty" xml:"RtDetail,omitempty" type:"Struct"`
	// The list of transition phrases.
	TransitionList []*string `json:"TransitionList,omitempty" xml:"TransitionList,omitempty" type:"Repeated"`
	// The list of similar questions.
	RelatedQuestionList []*string `json:"relatedQuestionList,omitempty" xml:"relatedQuestionList,omitempty" type:"Repeated"`
	// The list of formal answers.
	SentenceList []*TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
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

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) GetHitSystemAskConfig() *string {
	return s.HitSystemAskConfig
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) GetRtDetail() *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail {
	return s.RtDetail
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

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) SetHitSystemAskConfig(v string) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody {
	s.HitSystemAskConfig = &v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody) SetRtDetail(v *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBody {
	s.RtDetail = v
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
	if s.RtDetail != nil {
		if err := s.RtDetail.Validate(); err != nil {
			return err
		}
	}
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

type TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail struct {
	// The list of RTs for each segment from the time the user question is received on the text side to the sequential generation of the formal answer.
	//
	// example:
	//
	// [200, 256, 300, 430]
	SentenceRtList interface{} `json:"SentenceRtList,omitempty" xml:"SentenceRtList,omitempty"`
	// The list of RTs for each segment from the time the user question is received on the text side to the sequential generation of the transition phrases.
	//
	// example:
	//
	// [130]
	TransitionRtList interface{} `json:"TransitionRtList,omitempty" xml:"TransitionRtList,omitempty"`
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) GetSentenceRtList() interface{} {
	return s.SentenceRtList
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) GetTransitionRtList() interface{} {
	return s.TransitionRtList
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) SetSentenceRtList(v interface{}) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail {
	s.SentenceRtList = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) SetTransitionRtList(v interface{}) *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail {
	s.TransitionRtList = v
	return s
}

func (s *TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodyRtDetail) Validate() error {
	return dara.Validate(s)
}

type TongyiChatDebugInfoResponseBodyAnswerInfoMessageBodyDirectMessageBodySentenceList struct {
	// The formal answer entry.
	//
	// example:
	//
	// The weather is sunny today
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The number that corresponds one-to-one with the Number in AnswerReferenceInfo.ItemList for generating this answer.
	//
	// example:
	//
	// 1
	ReferNumber *int32 `json:"ReferNumber,omitempty" xml:"ReferNumber,omitempty"`
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
	// The debugging input information.
	//
	// example:
	//
	// User question: Transfer to agent
	//
	// Rule hits: [transfer] agent [customer service|service|representative]
	Input interface{} `json:"Input,omitempty" xml:"Input,omitempty"`
	// The policy name. Valid values:
	//
	// - High-frequency Q&A direct output
	//
	// - Keyword-based transfer to agent
	//
	// - Global sensitive words
	//
	// This field is available only when NodeType is system_strategy.
	//
	// example:
	//
	// Keyword-based transfer to agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node type. Valid values:
	//
	// 	- **system_strategy**: system policy.
	//
	// 	- **rewrite_query**: retrieved question.
	//
	// 	- **invoke_llm**: LLM invocation.
	//
	// 	- **invoke_tools**: tool calling.
	//
	// example:
	//
	// system_strategy
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The output information.
	//
	// example:
	//
	// commands:{"sysToAgent":"{\\"skillGroup\\":\\"\\",\\"ext\\":\\"\\",\\"toAgentReason\\":\\"HitKeywords\\"}"}
	//
	// response:Transferring you to a human agent
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
