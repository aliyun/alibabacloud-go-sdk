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
	// model response.
	ChatCompletion *ChatWithKnowledgeBaseStreamResponseBodyChatCompletion `json:"ChatCompletion,omitempty" xml:"ChatCompletion,omitempty" type:"Struct"`
	// The returned information.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Retrieve information from multiple knowledge bases.
	MultiCollectionRecallResult *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResult `json:"MultiCollectionRecallResult,omitempty" xml:"MultiCollectionRecallResult,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**.
	//
	// 	- **fail**.
	//
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

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletion struct {
	// Text content generated in real time.
	Choices []*ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices `json:"Choices,omitempty" xml:"Choices,omitempty" type:"Repeated"`
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

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoices struct {
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
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage struct {
	// The content of the document.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Model reasoning chain content.
	//
	// example:
	//
	// Logical reasoning process
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

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) GetReasoningContent() *string {
	return s.ReasoningContent
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

func (s *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage) SetReasoningContent(v string) *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessage {
	s.ReasoningContent = &v
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

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCalls struct {
	// Function call information.
	Function *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// The ID.
	//
	// example:
	//
	// "chatcmpl-c1bebafa-cc48-44e2-88c6-1a3572952f8e"
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The position of this tool in the \\"input\\" request parameter, which starts from 0.
	//
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
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionChoicesMessageToolCallsFunction struct {
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
	PromptTokensDetails *ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails `json:"PromptTokensDetails,omitempty" xml:"PromptTokensDetails,omitempty" type:"Struct"`
	// The total number of tokens.
	//
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
	if s.PromptTokensDetails != nil {
		if err := s.PromptTokensDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamResponseBodyChatCompletionUsagePromptTokensDetails struct {
	// The number of tokens from cache hits.
	//
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
	// The details of the entity.
	Entities []*string `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The retrieved item.
	Matches []*ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Repeated"`
	// The relationship name.
	Relations []*string `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Repeated"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**.
	//
	// 	- **fail**.
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
	Usage *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
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

type ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches struct {
	// The content of the document.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The file name.
	//
	// example:
	//
	// a14b0221-e3f2-4cf2-96cd-b3c293510770.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The URL of the image result. By default, the URL is valid for 2 hours.
	//
	// You can use the UrlExpiration parameter to specify a validity period.
	//
	// example:
	//
	// http://dailyshort-sh.oss-cn-shanghai.aliyuncs.com/vod-8efba5/f06147795c6c71f080605420848d0302/0ca34d5743a84bf7c68f489a60715dac-ld.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The unique ID of the vector data.
	//
	// >  If you leave this parameter empty, a unique ID is automatically generated by using uuidgen. If it is not empty and conflicts with an existing ID in the database, the value in the database will be updated with the data from the API.
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
	// Metadata.
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The rerank score.
	//
	// example:
	//
	// 0.12
	RerankScore *float64 `json:"RerankScore,omitempty" xml:"RerankScore,omitempty"`
	// The source of the retrieved results. 1 indicates vector retrieval, 2 indicates full-text retrieval, and 3 indicates dual-path retrieval.
	//
	// example:
	//
	// 0.12
	RetrievalSource *int64 `json:"RetrievalSource,omitempty" xml:"RetrievalSource,omitempty"`
	// The similarity score of the data. It is related to the `l2, ip, or cosine` algorithm that is specified when you create an index.
	//
	// example:
	//
	// 10
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The vector data.
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

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) GetMetadata() map[string]interface{} {
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

func (s *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches) SetMetadata(v map[string]interface{}) *ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultMatches {
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

type ChatWithKnowledgeBaseStreamResponseBodyMultiCollectionRecallResultUsage struct {
	// The number of tokens that are used during vectorization.
	//
	// >  A token is the minimum unit for splitting text. A token can be a word, phrase, punctuation, or character.
	//
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
