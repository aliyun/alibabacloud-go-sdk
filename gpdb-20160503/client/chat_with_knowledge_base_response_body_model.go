// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatCompletion(v *ChatWithKnowledgeBaseResponseBodyChatCompletion) *ChatWithKnowledgeBaseResponseBody
	GetChatCompletion() *ChatWithKnowledgeBaseResponseBodyChatCompletion
	SetMessage(v string) *ChatWithKnowledgeBaseResponseBody
	GetMessage() *string
	SetMultiCollectionRecallResult(v *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) *ChatWithKnowledgeBaseResponseBody
	GetMultiCollectionRecallResult() *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult
	SetRequestId(v string) *ChatWithKnowledgeBaseResponseBody
	GetRequestId() *string
	SetStatus(v string) *ChatWithKnowledgeBaseResponseBody
	GetStatus() *string
}

type ChatWithKnowledgeBaseResponseBody struct {
	ChatCompletion *ChatWithKnowledgeBaseResponseBodyChatCompletion `json:"ChatCompletion,omitempty" xml:"ChatCompletion,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message                     *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	MultiCollectionRecallResult *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult `json:"MultiCollectionRecallResult,omitempty" xml:"MultiCollectionRecallResult,omitempty" type:"Struct"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ChatWithKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBody) GetChatCompletion() *ChatWithKnowledgeBaseResponseBodyChatCompletion {
	return s.ChatCompletion
}

func (s *ChatWithKnowledgeBaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatWithKnowledgeBaseResponseBody) GetMultiCollectionRecallResult() *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	return s.MultiCollectionRecallResult
}

func (s *ChatWithKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatWithKnowledgeBaseResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ChatWithKnowledgeBaseResponseBody) SetChatCompletion(v *ChatWithKnowledgeBaseResponseBodyChatCompletion) *ChatWithKnowledgeBaseResponseBody {
	s.ChatCompletion = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBody) SetMessage(v string) *ChatWithKnowledgeBaseResponseBody {
	s.Message = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBody) SetMultiCollectionRecallResult(v *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) *ChatWithKnowledgeBaseResponseBody {
	s.MultiCollectionRecallResult = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBody) SetRequestId(v string) *ChatWithKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBody) SetStatus(v string) *ChatWithKnowledgeBaseResponseBody {
	s.Status = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyChatCompletion struct {
	Choices []*ChatWithKnowledgeBaseResponseBodyChatCompletionChoices `json:"Choices,omitempty" xml:"Choices,omitempty" type:"Repeated"`
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
	Model *string                                               `json:"Model,omitempty" xml:"Model,omitempty"`
	Usage *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletion) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletion) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) GetChoices() []*ChatWithKnowledgeBaseResponseBodyChatCompletionChoices {
	return s.Choices
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) GetCreated() *int64 {
	return s.Created
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) GetId() *string {
	return s.Id
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) GetModel() *string {
	return s.Model
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) GetUsage() *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage {
	return s.Usage
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) SetChoices(v []*ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) *ChatWithKnowledgeBaseResponseBodyChatCompletion {
	s.Choices = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) SetCreated(v int64) *ChatWithKnowledgeBaseResponseBodyChatCompletion {
	s.Created = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) SetId(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletion {
	s.Id = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) SetModel(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletion {
	s.Model = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) SetUsage(v *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) *ChatWithKnowledgeBaseResponseBodyChatCompletion {
	s.Usage = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletion) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoices struct {
	// example:
	//
	// finish
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int64                                                         `json:"Index,omitempty" xml:"Index,omitempty"`
	Message *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) GetIndex() *int64 {
	return s.Index
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) GetMessage() *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage {
	return s.Message
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) SetFinishReason(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices {
	s.FinishReason = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) SetIndex(v int64) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices {
	s.Index = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) SetMessage(v *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices {
	s.Message = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoices) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// user
	Role      *string                                                                   `json:"Role,omitempty" xml:"Role,omitempty"`
	ToolCalls []*ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls `json:"ToolCalls,omitempty" xml:"ToolCalls,omitempty" type:"Repeated"`
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) GetContent() *string {
	return s.Content
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) GetRole() *string {
	return s.Role
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) GetToolCalls() []*ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls {
	return s.ToolCalls
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) SetContent(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage {
	s.Content = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) SetRole(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage {
	s.Role = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) SetToolCalls(v []*ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage {
	s.ToolCalls = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls struct {
	Function *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
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

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) GetFunction() *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction {
	return s.Function
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) GetId() *string {
	return s.Id
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) GetIndex() *int64 {
	return s.Index
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) SetFunction(v *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls {
	s.Function = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) SetId(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls {
	s.Id = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) SetIndex(v int64) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls {
	s.Index = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction struct {
	// example:
	//
	// {"city":"hangzhou"}
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// example:
	//
	// "get_weather"
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) GetArguments() *string {
	return s.Arguments
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) SetArguments(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction {
	s.Arguments = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) SetName(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionUsage struct {
	// example:
	//
	// 42
	CompletionTokens *int64 `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// example:
	//
	// 42
	PromptTokens        *int64                                                                   `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	PromptTokensDetails *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails `json:"PromptTokensDetails,omitempty" xml:"PromptTokensDetails,omitempty" type:"Struct"`
	// example:
	//
	// 42
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) GetCompletionTokens() *int64 {
	return s.CompletionTokens
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) GetPromptTokens() *int64 {
	return s.PromptTokens
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) GetPromptTokensDetails() *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails {
	return s.PromptTokensDetails
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) SetCompletionTokens(v int64) *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage {
	s.CompletionTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) SetPromptTokens(v int64) *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage {
	s.PromptTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) SetPromptTokensDetails(v *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails) *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage {
	s.PromptTokensDetails = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) SetTotalTokens(v int64) *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage {
	s.TotalTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsage) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails struct {
	// example:
	//
	// 24
	CachedTokens *int64 `json:"CachedTokens,omitempty" xml:"CachedTokens,omitempty"`
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails) GetCachedTokens() *int64 {
	return s.CachedTokens
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails) SetCachedTokens(v int64) *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails {
	s.CachedTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult struct {
	Entities  []*string                                                              `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	Matches   []*ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Repeated"`
	Relations []*string                                                              `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Repeated"`
	// example:
	//
	// 6B9E3255-4543-5B3B-9E00-6490CA64742B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 42
	Tokens *int64                                                             `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
	Usage  *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GetEntities() []*string {
	return s.Entities
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GetMatches() []*ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	return s.Matches
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GetRelations() []*string {
	return s.Relations
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GetStatus() *string {
	return s.Status
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GetTokens() *int64 {
	return s.Tokens
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) GetUsage() *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage {
	return s.Usage
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) SetEntities(v []*string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	s.Entities = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) SetMatches(v []*ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	s.Matches = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) SetRelations(v []*string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	s.Relations = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) SetRequestId(v string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	s.RequestId = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) SetStatus(v string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	s.Status = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) SetTokens(v int64) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	s.Tokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) SetUsage(v *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult {
	s.Usage = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// process_info_19b9df4dc9ad4bf2b30eb2faa4a9a987.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// http://viapi-customer-pop.oss-cn-shanghai.aliyuncs.com/b4d8_207196811002111319_570c0e199f03428f812ab21fcc00dd6a
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// example:
	//
	// 273e3fc7-8f56-4167-a1bb-d35d2f3b9043
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {"page":1}
	LoaderMetadata interface{}                                                                  `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	Metadata       *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// example:
	//
	// 0.1
	RerankScore *float64 `json:"RerankScore,omitempty" xml:"RerankScore,omitempty"`
	// example:
	//
	// 3
	RetrievalSource *int64 `json:"RetrievalSource,omitempty" xml:"RetrievalSource,omitempty"`
	// example:
	//
	// 12
	Score  *float64   `json:"Score,omitempty" xml:"Score,omitempty"`
	Vector []*float64 `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetContent() *string {
	return s.Content
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetFileName() *string {
	return s.FileName
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetFileURL() *string {
	return s.FileURL
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetId() *string {
	return s.Id
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetLoaderMetadata() interface{} {
	return s.LoaderMetadata
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetMetadata() *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata {
	return s.Metadata
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetRerankScore() *float64 {
	return s.RerankScore
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetRetrievalSource() *int64 {
	return s.RetrievalSource
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetScore() *float64 {
	return s.Score
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) GetVector() []*float64 {
	return s.Vector
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetContent(v string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.Content = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetFileName(v string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.FileName = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetFileURL(v string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.FileURL = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetId(v string) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.Id = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetLoaderMetadata(v interface{}) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.LoaderMetadata = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetMetadata(v *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.Metadata = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetRerankScore(v float64) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.RerankScore = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetRetrievalSource(v int64) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.RetrievalSource = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetScore(v float64) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.Score = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) SetVector(v []*float64) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches {
	s.Vector = v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata struct {
	// example:
	//
	// 1
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata) GetSource() *int64 {
	return s.Source
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata) SetSource(v int64) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata {
	s.Source = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage struct {
	// example:
	//
	// 21
	EmbeddingTokens *int64 `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage) GetEmbeddingTokens() *int64 {
	return s.EmbeddingTokens
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage) SetEmbeddingTokens(v int64) *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage {
	s.EmbeddingTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage) Validate() error {
	return dara.Validate(s)
}
