// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTaskDialoguesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSearchTaskDialoguesResponseBody
	GetCode() *string
	SetData(v []*ListSearchTaskDialoguesResponseBodyData) *ListSearchTaskDialoguesResponseBody
	GetData() []*ListSearchTaskDialoguesResponseBodyData
	SetHttpStatusCode(v int32) *ListSearchTaskDialoguesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSearchTaskDialoguesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListSearchTaskDialoguesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSearchTaskDialoguesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSearchTaskDialoguesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSearchTaskDialoguesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSearchTaskDialoguesResponseBody
	GetTotalCount() *int32
}

type ListSearchTaskDialoguesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data []*ListSearchTaskDialoguesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. \\`true\\` indicates success. \\`false\\` indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSearchTaskDialoguesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialoguesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialoguesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSearchTaskDialoguesResponseBody) GetData() []*ListSearchTaskDialoguesResponseBodyData {
	return s.Data
}

func (s *ListSearchTaskDialoguesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSearchTaskDialoguesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSearchTaskDialoguesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSearchTaskDialoguesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchTaskDialoguesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchTaskDialoguesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSearchTaskDialoguesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSearchTaskDialoguesResponseBody) SetCode(v string) *ListSearchTaskDialoguesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetData(v []*ListSearchTaskDialoguesResponseBodyData) *ListSearchTaskDialoguesResponseBody {
	s.Data = v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetHttpStatusCode(v int32) *ListSearchTaskDialoguesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetMessage(v string) *ListSearchTaskDialoguesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetPageNumber(v int32) *ListSearchTaskDialoguesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetPageSize(v int32) *ListSearchTaskDialoguesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetRequestId(v string) *ListSearchTaskDialoguesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetSuccess(v bool) *ListSearchTaskDialoguesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) SetTotalCount(v int32) *ListSearchTaskDialoguesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSearchTaskDialoguesResponseBodyData struct {
	// The configuration.
	ChatConfig *ListSearchTaskDialoguesResponseBodyDataChatConfig `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty" type:"Struct"`
	// The time when the task was created.
	//
	// example:
	//
	// 2024-11-25 13:33:01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of generation.
	//
	// example:
	//
	// 24
	DialogueType *int32 `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	// The generated content that the customer rated as good.
	//
	// example:
	//
	// xxx
	GoodText *string `json:"GoodText,omitempty" xml:"GoodText,omitempty"`
	// The original session ID.
	//
	// example:
	//
	// xxxx
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// The input prompt.
	//
	// example:
	//
	// xxx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The rating of the content, which can be Like or Dislike.
	//
	// example:
	//
	// thumbsUp
	Rating *string `json:"Rating,omitempty" xml:"Rating,omitempty"`
	// The response body as a string.
	//
	// example:
	//
	// {}
	ResponseBodyStr *string `json:"ResponseBodyStr,omitempty" xml:"ResponseBodyStr,omitempty"`
	// The unique ID of a single-turn conversation.
	//
	// example:
	//
	// xxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The tags.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The unique ID of the task.
	//
	// example:
	//
	// xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The customer feedback.
	//
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListSearchTaskDialoguesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialoguesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetChatConfig() *ListSearchTaskDialoguesResponseBodyDataChatConfig {
	return s.ChatConfig
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetDialogueType() *int32 {
	return s.DialogueType
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetGoodText() *string {
	return s.GoodText
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetPrompt() *string {
	return s.Prompt
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetRating() *string {
	return s.Rating
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetResponseBodyStr() *string {
	return s.ResponseBodyStr
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetTags() []*string {
	return s.Tags
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListSearchTaskDialoguesResponseBodyData) GetText() *string {
	return s.Text
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetChatConfig(v *ListSearchTaskDialoguesResponseBodyDataChatConfig) *ListSearchTaskDialoguesResponseBodyData {
	s.ChatConfig = v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetCreateTime(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetDialogueType(v int32) *ListSearchTaskDialoguesResponseBodyData {
	s.DialogueType = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetGoodText(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.GoodText = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetOriginSessionId(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.OriginSessionId = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetPrompt(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetRating(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.Rating = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetResponseBodyStr(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.ResponseBodyStr = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetSessionId(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetTags(v []*string) *ListSearchTaskDialoguesResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetTaskId(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) SetText(v string) *ListSearchTaskDialoguesResponseBodyData {
	s.Text = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyData) Validate() error {
	if s.ChatConfig != nil {
		if err := s.ChatConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSearchTaskDialoguesResponseBodyDataChatConfig struct {
	// The request type.
	//
	// example:
	//
	// 24
	DialogueType *int32 `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	// Indicates whether the process is end-to-end.
	EndToEnd *bool `json:"EndToEnd,omitempty" xml:"EndToEnd,omitempty"`
	// The generation level.
	//
	// example:
	//
	// concise
	GenerateLevel *string `json:"GenerateLevel,omitempty" xml:"GenerateLevel,omitempty"`
	// The type of content generation.
	//
	// example:
	//
	// copilotReference
	GenerateTechnology *string `json:"GenerateTechnology,omitempty" xml:"GenerateTechnology,omitempty"`
	// The search type.
	SearchModels []*string                                                     `json:"SearchModels,omitempty" xml:"SearchModels,omitempty" type:"Repeated"`
	SearchParam  *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam `json:"SearchParam,omitempty" xml:"SearchParam,omitempty" type:"Struct"`
}

func (s ListSearchTaskDialoguesResponseBodyDataChatConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialoguesResponseBodyDataChatConfig) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) GetDialogueType() *int32 {
	return s.DialogueType
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) GetEndToEnd() *bool {
	return s.EndToEnd
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) GetGenerateLevel() *string {
	return s.GenerateLevel
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) GetGenerateTechnology() *string {
	return s.GenerateTechnology
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) GetSearchModels() []*string {
	return s.SearchModels
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) GetSearchParam() *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam {
	return s.SearchParam
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) SetDialogueType(v int32) *ListSearchTaskDialoguesResponseBodyDataChatConfig {
	s.DialogueType = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) SetEndToEnd(v bool) *ListSearchTaskDialoguesResponseBodyDataChatConfig {
	s.EndToEnd = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) SetGenerateLevel(v string) *ListSearchTaskDialoguesResponseBodyDataChatConfig {
	s.GenerateLevel = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) SetGenerateTechnology(v string) *ListSearchTaskDialoguesResponseBodyDataChatConfig {
	s.GenerateTechnology = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) SetSearchModels(v []*string) *ListSearchTaskDialoguesResponseBodyDataChatConfig {
	s.SearchModels = v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) SetSearchParam(v *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) *ListSearchTaskDialoguesResponseBodyDataChatConfig {
	s.SearchParam = v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfig) Validate() error {
	if s.SearchParam != nil {
		if err := s.SearchParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam struct {
	EndTime               *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MultimodalSearchTypes []*string `json:"MultimodalSearchTypes,omitempty" xml:"MultimodalSearchTypes,omitempty" type:"Repeated"`
	// The search sources.
	SearchSources []*ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	StartTime     *string                                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) GetMultimodalSearchTypes() []*string {
	return s.MultimodalSearchTypes
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) GetSearchSources() []*ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources {
	return s.SearchSources
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) SetEndTime(v string) *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam {
	s.EndTime = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) SetMultimodalSearchTypes(v []*string) *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam {
	s.MultimodalSearchTypes = v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) SetSearchSources(v []*ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam {
	s.SearchSources = v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) SetStartTime(v string) *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam {
	s.StartTime = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParam) Validate() error {
	if s.SearchSources != nil {
		for _, item := range s.SearchSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources struct {
	// The type.
	//
	// example:
	//
	// x\\"x
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique ID.
	//
	// example:
	//
	// xx
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The description.
	//
	// example:
	//
	// x
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) GetCode() *string {
	return s.Code
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) GetName() *string {
	return s.Name
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) SetCode(v string) *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources {
	s.Code = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) SetDatasetName(v string) *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources {
	s.DatasetName = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) SetName(v string) *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources {
	s.Name = &v
	return s
}

func (s *ListSearchTaskDialoguesResponseBodyDataChatConfigSearchParamSearchSources) Validate() error {
	return dara.Validate(s)
}
