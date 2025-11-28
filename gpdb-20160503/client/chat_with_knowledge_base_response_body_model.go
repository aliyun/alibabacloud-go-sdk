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
	// model response.
	ChatCompletion *ChatWithKnowledgeBaseResponseBodyChatCompletion `json:"ChatCompletion,omitempty" xml:"ChatCompletion,omitempty" type:"Struct"`
	// The returned information.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Retrieve information from multiple knowledge bases.
	MultiCollectionRecallResult *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResult `json:"MultiCollectionRecallResult,omitempty" xml:"MultiCollectionRecallResult,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
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
	if s.ChatCompletion != nil {
		if err := s.ChatCompletion.Validate(); err != nil {
			return err
		}
	}
	if s.MultiCollectionRecallResult != nil {
		if err := s.MultiCollectionRecallResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyChatCompletion struct {
	// Text content generated in real time.
	Choices []*ChatWithKnowledgeBaseResponseBodyChatCompletionChoices `json:"Choices,omitempty" xml:"Choices,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 1758529748
	Created *int64 `json:"Created,omitempty" xml:"Created,omitempty"`
	// The ID of the response.
	//
	// example:
	//
	// 273e3fc7-8f56-4167-a1bb-d35d2f3b9043
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of tokens used in LLM output.
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
	if s.Choices != nil {
		for _, item := range s.Choices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoices struct {
	// Finish reason.
	//
	// example:
	//
	// finish
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// The sequence number of the reply.
	//
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// LLM response.
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
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage struct {
	// The content of the document.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 模型思维内容
	//
	// example:
	//
	// 逻辑推理过程
	ReasoningContent *string `json:"ReasoningContent,omitempty" xml:"ReasoningContent,omitempty"`
	// Message role:
	//
	// 	- system
	//
	// 	- user
	//
	// 	- assistant
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Tool call response.
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

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) GetReasoningContent() *string {
	return s.ReasoningContent
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

func (s *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage) SetReasoningContent(v string) *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessage {
	s.ReasoningContent = &v
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
	if s.ToolCalls != nil {
		for _, item := range s.ToolCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCalls struct {
	// Function call information.
	Function *ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// ID
	//
	// example:
	//
	// "chatcmpl-c1bebafa-cc48-44e2-88c6-1a3572952f8e"
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The position of this tool in the \\"input\\" request parameter.
	//
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
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionChoicesMessageToolCallsFunction struct {
	// Arguments of the called function.
	//
	// example:
	//
	// {"city":"hangzhou"}
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// The name of the called function.
	//
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
	// The number of tokens consumed by the generated content.
	//
	// example:
	//
	// 42
	CompletionTokens *int64 `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// The number of tokens consumed by the prompt.
	//
	// example:
	//
	// 42
	PromptTokens *int64 `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	// The details about the prompt token.
	PromptTokensDetails *ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails `json:"PromptTokensDetails,omitempty" xml:"PromptTokensDetails,omitempty" type:"Struct"`
	// The total number of tokens.
	//
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
	if s.PromptTokensDetails != nil {
		if err := s.PromptTokensDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyChatCompletionUsagePromptTokensDetails struct {
	// The number of tokens from cache hits.
	//
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
	// The details of the entity.
	Entities []*string `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The retrieved items.
	Matches []*ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Repeated"`
	// The name of the file.
	Relations []*string `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 6B9E3255-4543-5B3B-9E00-6490CA64742B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of tokens that are consumed.
	//
	// example:
	//
	// 42
	Tokens *int64 `json:"Tokens,omitempty" xml:"Tokens,omitempty"`
	// The number of tokens that are consumed during document understanding or embedding.
	Usage *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
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
	if s.Matches != nil {
		for _, item := range s.Matches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatches struct {
	// The content of the document.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The file name.
	//
	// example:
	//
	// process_info_19b9df4dc9ad4bf2b30eb2faa4a9a987.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The URL of the image result. By default, the URL is valid for 2 hours.
	//
	// You can use the UrlExpiration parameter to specify a validity period.
	//
	// example:
	//
	// http://viapi-customer-pop.oss-cn-shanghai.aliyuncs.com/b4d8_207196811002111319_570c0e199f03428f812ab21fcc00dd6a
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The unique ID of the vector data.
	//
	// example:
	//
	// 273e3fc7-8f56-4167-a1bb-d35d2f3b9043
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Document loader metadata.
	//
	// example:
	//
	// {"page":1}
	LoaderMetadata interface{} `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	// The metadata.
	Metadata *ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// The rerank score.
	//
	// example:
	//
	// 0.1
	RerankScore *float64 `json:"RerankScore,omitempty" xml:"RerankScore,omitempty"`
	// The source of the retrieved results. 1 indicates vector retrieval, 2 indicates full-text retrieval, and 3 indicates dual-path retrieval.
	//
	// example:
	//
	// 3
	RetrievalSource *int64 `json:"RetrievalSource,omitempty" xml:"RetrievalSource,omitempty"`
	// The similarity score of the data. It is related to the `l2, ip, or cosine` algorithm that is specified when you create an index.
	//
	// example:
	//
	// 12
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The vector data.
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
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseResponseBodyMultiCollectionRecallResultMatchesMetadata struct {
	// The source of the document.
	//
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
	// The number of tokens that are used during vectorization.
	//
	// >  A token is the minimum unit for splitting text. A token can be a word, phrase, punctuation, or character.
	//
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
