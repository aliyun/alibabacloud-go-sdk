// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMinutesSummaryResponseBody
	GetRequestId() *string
	SetSummary(v *QueryMinutesSummaryResponseBodySummary) *QueryMinutesSummaryResponseBody
	GetSummary() *QueryMinutesSummaryResponseBodySummary
	SetVendorRequestId(v string) *QueryMinutesSummaryResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryMinutesSummaryResponseBody
	GetVendorType() *string
}

type QueryMinutesSummaryResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Summary   *QueryMinutesSummaryResponseBodySummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryMinutesSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMinutesSummaryResponseBody) GetSummary() *QueryMinutesSummaryResponseBodySummary {
	return s.Summary
}

func (s *QueryMinutesSummaryResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryMinutesSummaryResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryMinutesSummaryResponseBody) SetRequestId(v string) *QueryMinutesSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBody) SetSummary(v *QueryMinutesSummaryResponseBodySummary) *QueryMinutesSummaryResponseBody {
	s.Summary = v
	return s
}

func (s *QueryMinutesSummaryResponseBody) SetVendorRequestId(v string) *QueryMinutesSummaryResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBody) SetVendorType(v string) *QueryMinutesSummaryResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryMinutesSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesSummaryResponseBodySummary struct {
	// example:
	//
	// {}
	Actions *QueryMinutesSummaryResponseBodySummaryActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Struct"`
	// example:
	//
	// []
	AutoChapters []*QueryMinutesSummaryResponseBodySummaryAutoChapters `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty" type:"Repeated"`
	// example:
	//
	// []
	ConversationalSummary []*QueryMinutesSummaryResponseBodySummaryConversationalSummary `json:"ConversationalSummary,omitempty" xml:"ConversationalSummary,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	KeySentences *QueryMinutesSummaryResponseBodySummaryKeySentences `json:"KeySentences,omitempty" xml:"KeySentences,omitempty" type:"Struct"`
	// example:
	//
	// []
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// example:
	//
	// 全文摘要结果，全文摘要通过两三百字的篇幅将最重要的信息呈现出来，用于快速了解记录内容与主旨
	ParagraphSummary *string `json:"ParagraphSummary,omitempty" xml:"ParagraphSummary,omitempty"`
	// example:
	//
	// []
	QuestionsAnsweringSummary []*QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary `json:"QuestionsAnsweringSummary,omitempty" xml:"QuestionsAnsweringSummary,omitempty" type:"Repeated"`
}

func (s QueryMinutesSummaryResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummary) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummary) GetActions() *QueryMinutesSummaryResponseBodySummaryActions {
	return s.Actions
}

func (s *QueryMinutesSummaryResponseBodySummary) GetAutoChapters() []*QueryMinutesSummaryResponseBodySummaryAutoChapters {
	return s.AutoChapters
}

func (s *QueryMinutesSummaryResponseBodySummary) GetConversationalSummary() []*QueryMinutesSummaryResponseBodySummaryConversationalSummary {
	return s.ConversationalSummary
}

func (s *QueryMinutesSummaryResponseBodySummary) GetKeySentences() *QueryMinutesSummaryResponseBodySummaryKeySentences {
	return s.KeySentences
}

func (s *QueryMinutesSummaryResponseBodySummary) GetKeywords() []*string {
	return s.Keywords
}

func (s *QueryMinutesSummaryResponseBodySummary) GetParagraphSummary() *string {
	return s.ParagraphSummary
}

func (s *QueryMinutesSummaryResponseBodySummary) GetQuestionsAnsweringSummary() []*QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	return s.QuestionsAnsweringSummary
}

func (s *QueryMinutesSummaryResponseBodySummary) SetActions(v *QueryMinutesSummaryResponseBodySummaryActions) *QueryMinutesSummaryResponseBodySummary {
	s.Actions = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetAutoChapters(v []*QueryMinutesSummaryResponseBodySummaryAutoChapters) *QueryMinutesSummaryResponseBodySummary {
	s.AutoChapters = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetConversationalSummary(v []*QueryMinutesSummaryResponseBodySummaryConversationalSummary) *QueryMinutesSummaryResponseBodySummary {
	s.ConversationalSummary = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetKeySentences(v *QueryMinutesSummaryResponseBodySummaryKeySentences) *QueryMinutesSummaryResponseBodySummary {
	s.KeySentences = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetKeywords(v []*string) *QueryMinutesSummaryResponseBodySummary {
	s.Keywords = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetParagraphSummary(v string) *QueryMinutesSummaryResponseBodySummary {
	s.ParagraphSummary = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) SetQuestionsAnsweringSummary(v []*QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) *QueryMinutesSummaryResponseBodySummary {
	s.QuestionsAnsweringSummary = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummary) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesSummaryResponseBodySummaryActions struct {
	// example:
	//
	// 7910000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2
	SentenceId *int64 `json:"SentenceId,omitempty" xml:"SentenceId,omitempty"`
	// example:
	//
	// 7901100
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// 内容
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryActions) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryActions) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) GetEnd() *int64 {
	return s.End
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) GetId() *int64 {
	return s.Id
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) GetSentenceId() *int64 {
	return s.SentenceId
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) GetStart() *int64 {
	return s.Start
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) GetText() *string {
	return s.Text
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetEnd(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.End = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetId(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.Id = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetSentenceId(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.SentenceId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetStart(v int64) *QueryMinutesSummaryResponseBodySummaryActions {
	s.Start = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) SetText(v string) *QueryMinutesSummaryResponseBodySummaryActions {
	s.Text = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryActions) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesSummaryResponseBodySummaryAutoChapters struct {
	// example:
	//
	// 7910000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 章节的一句话标题
	Headline *string `json:"Headline,omitempty" xml:"Headline,omitempty"`
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 7901100
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// 章节总结
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryAutoChapters) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryAutoChapters) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) GetEnd() *int64 {
	return s.End
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) GetHeadline() *string {
	return s.Headline
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) GetId() *int64 {
	return s.Id
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) GetStart() *int64 {
	return s.Start
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) GetSummary() *string {
	return s.Summary
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetEnd(v int64) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.End = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetHeadline(v string) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Headline = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetId(v int64) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Id = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetStart(v int64) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Start = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) SetSummary(v string) *QueryMinutesSummaryResponseBodySummaryAutoChapters {
	s.Summary = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryAutoChapters) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesSummaryResponseBodySummaryConversationalSummary struct {
	// example:
	//
	// 012345
	SpeakerId *string `json:"SpeakerId,omitempty" xml:"SpeakerId,omitempty"`
	// example:
	//
	// 发言人姓名
	SpeakerName *string `json:"SpeakerName,omitempty" xml:"SpeakerName,omitempty"`
	// example:
	//
	// 发言人对应的总结
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryConversationalSummary) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryConversationalSummary) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) GetSpeakerId() *string {
	return s.SpeakerId
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) GetSpeakerName() *string {
	return s.SpeakerName
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) GetSummary() *string {
	return s.Summary
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) SetSpeakerId(v string) *QueryMinutesSummaryResponseBodySummaryConversationalSummary {
	s.SpeakerId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) SetSpeakerName(v string) *QueryMinutesSummaryResponseBodySummaryConversationalSummary {
	s.SpeakerName = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) SetSummary(v string) *QueryMinutesSummaryResponseBodySummaryConversationalSummary {
	s.Summary = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryConversationalSummary) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesSummaryResponseBodySummaryKeySentences struct {
	// example:
	//
	// 7910000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2
	SentenceId *int64 `json:"SentenceId,omitempty" xml:"SentenceId,omitempty"`
	// example:
	//
	// 7901100
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// 内容
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QueryMinutesSummaryResponseBodySummaryKeySentences) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryKeySentences) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) GetEnd() *int64 {
	return s.End
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) GetId() *int64 {
	return s.Id
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) GetSentenceId() *int64 {
	return s.SentenceId
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) GetStart() *int64 {
	return s.Start
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) GetText() *string {
	return s.Text
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetEnd(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.End = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetId(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.Id = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetSentenceId(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.SentenceId = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetStart(v int64) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.Start = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) SetText(v string) *QueryMinutesSummaryResponseBodySummaryKeySentences {
	s.Text = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryKeySentences) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary struct {
	// example:
	//
	// 问题
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 回答
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// []
	SentenceIdsOfAnswer []*int64 `json:"SentenceIdsOfAnswer,omitempty" xml:"SentenceIdsOfAnswer,omitempty" type:"Repeated"`
	// example:
	//
	// []
	SentenceIdsOfQuestion []*int64 `json:"SentenceIdsOfQuestion,omitempty" xml:"SentenceIdsOfQuestion,omitempty" type:"Repeated"`
}

func (s QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) GetAnswer() *string {
	return s.Answer
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) GetQuestion() *string {
	return s.Question
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) GetSentenceIdsOfAnswer() []*int64 {
	return s.SentenceIdsOfAnswer
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) GetSentenceIdsOfQuestion() []*int64 {
	return s.SentenceIdsOfQuestion
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetAnswer(v string) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.Answer = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetQuestion(v string) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.Question = &v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetSentenceIdsOfAnswer(v []*int64) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.SentenceIdsOfAnswer = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) SetSentenceIdsOfQuestion(v []*int64) *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary {
	s.SentenceIdsOfQuestion = v
	return s
}

func (s *QueryMinutesSummaryResponseBodySummaryQuestionsAnsweringSummary) Validate() error {
	return dara.Validate(s)
}
