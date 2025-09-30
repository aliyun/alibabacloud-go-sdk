// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatCompletion(v *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) *ChatWithKnowledgeBaseStreamResponseBody
	GetChatCompletion() *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion
	SetMessage(v string) *ChatWithKnowledgeBaseStreamResponseBody
	GetMessage() *string
	SetMultiCollectionRecallResult(v *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) *ChatWithKnowledgeBaseStreamResponseBody
	GetMultiCollectionRecallResult() *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult
	SetRequestId(v string) *ChatWithKnowledgeBaseStreamResponseBody
	GetRequestId() *string
	SetStatus(v string) *ChatWithKnowledgeBaseStreamResponseBody
	GetStatus() *string
}

type ChatWithKnowledgeBaseStreamResponseBody struct {
	ChatCompletion *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion `json:"ChatCompletion,omitempty" xml:"ChatCompletion,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message                     *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	MultiCollectionRecallResult *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult `json:"MultiCollectionRecallResult,omitempty" xml:"MultiCollectionRecallResult,omitempty" type:"Struct"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) GetChatCompletion() *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion {
	return s.ChatCompletion
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) GetMultiCollectionRecallResult() *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	return s.MultiCollectionRecallResult
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) SetChatCompletion(v *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) *ChatWithKnowledgeBaseStreamResponseBody {
	s.ChatCompletion = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) SetMessage(v string) *ChatWithKnowledgeBaseStreamResponseBody {
	s.Message = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) SetMultiCollectionRecallResult(v *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) *ChatWithKnowledgeBaseStreamResponseBody {
	s.MultiCollectionRecallResult = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) SetRequestId(v string) *ChatWithKnowledgeBaseStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) SetStatus(v string) *ChatWithKnowledgeBaseStreamResponseBody {
	s.Status = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletion struct {
	Choices []*ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices `json:"Choices,omitempty" xml:"Choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1758529748
	Created *int64 `json:"Created,omitempty" xml:"Created,omitempty"`
	// example:
	//
	// 273e3fc7-8f56-4167-a1bb-d35d2f3b9043
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string                                                     `json:"Model,omitempty" xml:"Model,omitempty"`
	Usage *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) GetChoices() []*ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices {
	return s.Choices
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) GetCreated() *int64 {
	return s.Created
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) GetId() *string {
	return s.Id
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) GetModel() *string {
	return s.Model
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) GetUsage() *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage {
	return s.Usage
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) SetChoices(v []*ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion {
	s.Choices = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) SetCreated(v int64) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion {
	s.Created = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) SetId(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion {
	s.Id = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) SetModel(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion {
	s.Model = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) SetUsage(v *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion {
	s.Usage = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices struct {
	// example:
	//
	// finish
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int64                                                               `json:"Index,omitempty" xml:"Index,omitempty"`
	Message *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) GetIndex() *int64 {
	return s.Index
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) GetMessage() *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage {
	return s.Message
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) SetFinishReason(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices {
	s.FinishReason = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) SetIndex(v int64) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices {
	s.Index = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) SetMessage(v *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices {
	s.Message = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// user
	Role      *string                                                                         `json:"Role,omitempty" xml:"Role,omitempty"`
	ToolCalls []*ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls `json:"ToolCalls,omitempty" xml:"ToolCalls,omitempty" type:"Repeated"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) GetContent() *string {
	return s.Content
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) GetRole() *string {
	return s.Role
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) GetToolCalls() []*ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls {
	return s.ToolCalls
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) SetContent(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage {
	s.Content = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) SetRole(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage {
	s.Role = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) SetToolCalls(v []*ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage {
	s.ToolCalls = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls struct {
	Function *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// ID
	//
	// example:
	//
	// "chatcmpl-c1bebafa-cc48-44e2-88c6-1a3572952f8e"
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) GetFunction() *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction {
	return s.Function
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) GetId() *string {
	return s.Id
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) GetIndex() *int64 {
	return s.Index
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) SetFunction(v *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls {
	s.Function = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) SetId(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls {
	s.Id = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) SetIndex(v int64) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls {
	s.Index = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction struct {
	// example:
	//
	// {"city":"hangzhou"}
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// example:
	//
	// "get_weather"
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) GetArguments() *string {
	return s.Arguments
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) SetArguments(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction {
	s.Arguments = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) SetName(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage struct {
	// example:
	//
	// 42
	CompletionTokens *int64 `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// example:
	//
	// 42
	PromptTokens        *int64                                                                         `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	PromptTokensDetails *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails `json:"PromptTokensDetails,omitempty" xml:"PromptTokensDetails,omitempty" type:"Struct"`
	// example:
	//
	// 42
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) GetCompletionTokens() *int64 {
	return s.CompletionTokens
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) GetPromptTokens() *int64 {
	return s.PromptTokens
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) GetPromptTokensDetails() *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails {
	return s.PromptTokensDetails
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) SetCompletionTokens(v int64) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage {
	s.CompletionTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) SetPromptTokens(v int64) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage {
	s.PromptTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) SetPromptTokensDetails(v *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage {
	s.PromptTokensDetails = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) SetTotalTokens(v int64) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage {
	s.TotalTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsage) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails struct {
	// example:
	//
	// 24
	CachedTokens *int64 `json:"CachedTokens,omitempty" xml:"CachedTokens,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails) GetCachedTokens() *int64 {
	return s.CachedTokens
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails) SetCachedTokens(v int64) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails {
	s.CachedTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult struct {
	Entities  []*string                                                                    `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	Matches   []*ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Repeated"`
	Relations []*string                                                                    `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Repeated"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 42
	Tokens *int64                                                                   `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
	Usage  *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GetEntities() []*string {
	return s.Entities
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GetMatches() []*ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	return s.Matches
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GetRelations() []*string {
	return s.Relations
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GetStatus() *string {
	return s.Status
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GetTokens() *int64 {
	return s.Tokens
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) GetUsage() *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage {
	return s.Usage
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) SetEntities(v []*string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	s.Entities = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) SetMatches(v []*ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	s.Matches = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) SetRelations(v []*string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	s.Relations = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) SetRequestId(v string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	s.RequestId = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) SetStatus(v string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	s.Status = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) SetTokens(v int64) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	s.Tokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) SetUsage(v *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult {
	s.Usage = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// a14b0221-e3f2-4cf2-96cd-b3c293510770.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// http://dailyshort-sh.oss-cn-shanghai.aliyuncs.com/vod-8efba5/f06147795c6c71f080605420848d0302/0ca34d5743a84bf7c68f489a60715dac-ld.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// example:
	//
	// 273e3fc7-8f56-4167-a1bb-d35d2f3b9043
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {"page":1}
	LoaderMetadata interface{}                                                                        `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	Metadata       *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// example:
	//
	// 0.12
	RerankScore *float64 `json:"RerankScore,omitempty" xml:"RerankScore,omitempty"`
	// example:
	//
	// 0.12
	RetrievalSource *int64 `json:"RetrievalSource,omitempty" xml:"RetrievalSource,omitempty"`
	// example:
	//
	// 10
	Score  *float64   `json:"Score,omitempty" xml:"Score,omitempty"`
	Vector []*float64 `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetContent() *string {
	return s.Content
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetFileName() *string {
	return s.FileName
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetFileURL() *string {
	return s.FileURL
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetId() *string {
	return s.Id
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetLoaderMetadata() interface{} {
	return s.LoaderMetadata
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetMetadata() *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata {
	return s.Metadata
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetRerankScore() *float64 {
	return s.RerankScore
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetRetrievalSource() *int64 {
	return s.RetrievalSource
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetScore() *float64 {
	return s.Score
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetVector() []*float64 {
	return s.Vector
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetContent(v string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.Content = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetFileName(v string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.FileName = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetFileURL(v string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.FileURL = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetId(v string) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.Id = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetLoaderMetadata(v interface{}) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.LoaderMetadata = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetMetadata(v *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.Metadata = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetRerankScore(v float64) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.RerankScore = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetRetrievalSource(v int64) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.RetrievalSource = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetScore(v float64) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.Score = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetVector(v []*float64) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
	s.Vector = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata struct {
	// example:
	//
	// 1
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata) GetSource() *int64 {
	return s.Source
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata) SetSource(v int64) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata {
	s.Source = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatchesMetadata) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage struct {
	// example:
	//
	// 158
	EmbeddingTokens *int64 `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage) GetEmbeddingTokens() *int64 {
	return s.EmbeddingTokens
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage) SetEmbeddingTokens(v int64) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage {
	s.EmbeddingTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage) Validate() error {
	return dara.Validate(s)
}
