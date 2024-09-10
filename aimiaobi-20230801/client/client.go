// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskRequest) SetAgentKey(v string) *CancelAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelAsyncTaskRequest) SetTaskId(v string) *CancelAsyncTaskRequest {
	s.TaskId = &v
	return s
}

type CancelAsyncTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FB698445-61DA-5361-BF73-1C5F1157E888
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskResponseBody) SetCode(v string) *CancelAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetData(v bool) *CancelAsyncTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetHttpStatusCode(v int32) *CancelAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetMessage(v string) *CancelAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetRequestId(v string) *CancelAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetSuccess(v bool) *CancelAsyncTaskResponseBody {
	s.Success = &v
	return s
}

type CancelAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskResponse) SetHeaders(v map[string]*string) *CancelAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelAsyncTaskResponse) SetStatusCode(v int32) *CancelAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAsyncTaskResponse) SetBody(v *CancelAsyncTaskResponseBody) *CancelAsyncTaskResponse {
	s.Body = v
	return s
}

type ClearIntervenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ClearIntervenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ClearIntervenesRequest) SetAgentKey(v string) *ClearIntervenesRequest {
	s.AgentKey = &v
	return s
}

type ClearIntervenesResponseBody struct {
	// example:
	//
	// 0
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClearIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClearIntervenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponseBody) SetCode(v string) *ClearIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetData(v *ClearIntervenesResponseBodyData) *ClearIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ClearIntervenesResponseBody) SetHttpStatusCode(v int32) *ClearIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetMessage(v string) *ClearIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetRequestId(v string) *ClearIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetSuccess(v bool) *ClearIntervenesResponseBody {
	s.Success = &v
	return s
}

type ClearIntervenesResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ClearIntervenesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponseBodyData) SetFailIdList(v []*string) *ClearIntervenesResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ClearIntervenesResponseBodyData) SetTaskId(v string) *ClearIntervenesResponseBodyData {
	s.TaskId = &v
	return s
}

type ClearIntervenesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearIntervenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponse) SetHeaders(v map[string]*string) *ClearIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ClearIntervenesResponse) SetStatusCode(v int32) *ClearIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearIntervenesResponse) SetBody(v *ClearIntervenesResponseBody) *ClearIntervenesResponse {
	s.Body = v
	return s
}

type CreateGeneratedContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// government
	ContentDomain *string   `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText   *string   `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	Keywords      []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Prompt        *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentRequest) SetAgentKey(v string) *CreateGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContent(v string) *CreateGeneratedContentRequest {
	s.Content = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContentDomain(v string) *CreateGeneratedContentRequest {
	s.ContentDomain = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContentText(v string) *CreateGeneratedContentRequest {
	s.ContentText = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetKeywords(v []*string) *CreateGeneratedContentRequest {
	s.Keywords = v
	return s
}

func (s *CreateGeneratedContentRequest) SetPrompt(v string) *CreateGeneratedContentRequest {
	s.Prompt = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetTaskId(v string) *CreateGeneratedContentRequest {
	s.TaskId = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetTitle(v string) *CreateGeneratedContentRequest {
	s.Title = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetUuid(v string) *CreateGeneratedContentRequest {
	s.Uuid = &v
	return s
}

type CreateGeneratedContentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// government
	ContentDomain  *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText    *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// xxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateGeneratedContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentShrinkRequest) SetAgentKey(v string) *CreateGeneratedContentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContent(v string) *CreateGeneratedContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContentDomain(v string) *CreateGeneratedContentShrinkRequest {
	s.ContentDomain = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContentText(v string) *CreateGeneratedContentShrinkRequest {
	s.ContentText = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetKeywordsShrink(v string) *CreateGeneratedContentShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetPrompt(v string) *CreateGeneratedContentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetTaskId(v string) *CreateGeneratedContentShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetTitle(v string) *CreateGeneratedContentShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetUuid(v string) *CreateGeneratedContentShrinkRequest {
	s.Uuid = &v
	return s
}

type CreateGeneratedContentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 42
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentResponseBody) SetCode(v string) *CreateGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetData(v int64) *CreateGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetHttpStatusCode(v int32) *CreateGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetMessage(v string) *CreateGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetRequestId(v string) *CreateGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetSuccess(v bool) *CreateGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type CreateGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentResponse) SetHeaders(v map[string]*string) *CreateGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *CreateGeneratedContentResponse) SetStatusCode(v int32) *CreateGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGeneratedContentResponse) SetBody(v *CreateGeneratedContentResponseBody) *CreateGeneratedContentResponse {
	s.Body = v
	return s
}

type CreateTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daaa2e0c209xb26acb97009ea77bd4b_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) SetAgentKey(v string) *CreateTokenRequest {
	s.AgentKey = &v
	return s
}

type CreateTokenResponseBody struct {
	// example:
	//
	// NoData
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) SetCode(v string) *CreateTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTokenResponseBody) SetData(v *CreateTokenResponseBodyData) *CreateTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateTokenResponseBody) SetHttpStatusCode(v int32) *CreateTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTokenResponseBody) SetMessage(v string) *CreateTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTokenResponseBody) SetRequestId(v string) *CreateTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResponseBody) SetSuccess(v bool) *CreateTokenResponseBody {
	s.Success = &v
	return s
}

type CreateTokenResponseBodyData struct {
	// example:
	//
	// 1705388704855
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBodyData) SetExpiredTime(v int64) *CreateTokenResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *CreateTokenResponseBodyData) SetToken(v string) *CreateTokenResponseBodyData {
	s.Token = &v
	return s
}

type CreateTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) SetHeaders(v map[string]*string) *CreateTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenResponse) SetStatusCode(v int32) *CreateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenResponse) SetBody(v *CreateTokenResponseBody) *CreateTokenResponse {
	s.Body = v
	return s
}

type DeleteCustomTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 85
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCustomTextRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTextRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTextRequest) SetAgentKey(v string) *DeleteCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCustomTextRequest) SetCommodityCode(v string) *DeleteCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *DeleteCustomTextRequest) SetId(v int64) *DeleteCustomTextRequest {
	s.Id = &v
	return s
}

type DeleteCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTextResponseBody) SetCode(v string) *DeleteCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetData(v bool) *DeleteCustomTextResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetHttpStatusCode(v int32) *DeleteCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetMessage(v string) *DeleteCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetRequestId(v string) *DeleteCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTextResponseBody) SetSuccess(v bool) *DeleteCustomTextResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTextResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTextResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTextResponse) SetHeaders(v map[string]*string) *DeleteCustomTextResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTextResponse) SetStatusCode(v int32) *DeleteCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTextResponse) SetBody(v *DeleteCustomTextResponseBody) *DeleteCustomTextResponse {
	s.Body = v
	return s
}

type DeleteCustomTopicByTopicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 话题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteCustomTopicByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTopicByTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicByTopicRequest) SetAgentKey(v string) *DeleteCustomTopicByTopicRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCustomTopicByTopicRequest) SetTopic(v string) *DeleteCustomTopicByTopicRequest {
	s.Topic = &v
	return s
}

type DeleteCustomTopicByTopicResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 33
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomTopicByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTopicByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicByTopicResponseBody) SetCode(v string) *DeleteCustomTopicByTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetData(v int32) *DeleteCustomTopicByTopicResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetHttpStatusCode(v int32) *DeleteCustomTopicByTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetMessage(v string) *DeleteCustomTopicByTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetRequestId(v string) *DeleteCustomTopicByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponseBody) SetSuccess(v bool) *DeleteCustomTopicByTopicResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomTopicByTopicResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTopicByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTopicByTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTopicByTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicByTopicResponse) SetHeaders(v map[string]*string) *DeleteCustomTopicByTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTopicByTopicResponse) SetStatusCode(v int32) *DeleteCustomTopicByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponse) SetBody(v *DeleteCustomTopicByTopicResponseBody) *DeleteCustomTopicByTopicResponse {
	s.Body = v
	return s
}

type DeleteCustomTopicViewPointByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dfd73894e6a94fd79fe7ffbe865796fb
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
}

func (s DeleteCustomTopicViewPointByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTopicViewPointByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicViewPointByIdRequest) SetAgentKey(v string) *DeleteCustomTopicViewPointByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdRequest) SetCustomViewPointId(v string) *DeleteCustomTopicViewPointByIdRequest {
	s.CustomViewPointId = &v
	return s
}

type DeleteCustomTopicViewPointByIdResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 7
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomTopicViewPointByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTopicViewPointByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetCode(v string) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetData(v int32) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetHttpStatusCode(v int32) *DeleteCustomTopicViewPointByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetMessage(v string) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetRequestId(v string) *DeleteCustomTopicViewPointByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponseBody) SetSuccess(v bool) *DeleteCustomTopicViewPointByIdResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomTopicViewPointByIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTopicViewPointByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTopicViewPointByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomTopicViewPointByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicViewPointByIdResponse) SetHeaders(v map[string]*string) *DeleteCustomTopicViewPointByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponse) SetStatusCode(v int32) *DeleteCustomTopicViewPointByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponse) SetBody(v *DeleteCustomTopicViewPointByIdResponseBody) *DeleteCustomTopicViewPointByIdResponse {
	s.Body = v
	return s
}

type DeleteGeneratedContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 99
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentRequest) SetAgentKey(v string) *DeleteGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteGeneratedContentRequest) SetId(v int64) *DeleteGeneratedContentRequest {
	s.Id = &v
	return s
}

func (s *DeleteGeneratedContentRequest) SetRegionId(v string) *DeleteGeneratedContentRequest {
	s.RegionId = &v
	return s
}

type DeleteGeneratedContentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentResponseBody) SetCode(v string) *DeleteGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetData(v bool) *DeleteGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetHttpStatusCode(v int32) *DeleteGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetMessage(v string) *DeleteGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetRequestId(v string) *DeleteGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetSuccess(v bool) *DeleteGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type DeleteGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentResponse) SetHeaders(v map[string]*string) *DeleteGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *DeleteGeneratedContentResponse) SetStatusCode(v int32) *DeleteGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGeneratedContentResponse) SetBody(v *DeleteGeneratedContentResponseBody) *DeleteGeneratedContentResponse {
	s.Body = v
	return s
}

type DeleteInterveneRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 12345
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteInterveneRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleRequest) SetAgentKey(v string) *DeleteInterveneRuleRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteInterveneRuleRequest) SetRuleId(v int64) *DeleteInterveneRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteInterveneRuleResponseBody struct {
	// example:
	//
	// 0
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteInterveneRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInterveneRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponseBody) SetCode(v string) *DeleteInterveneRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetData(v *DeleteInterveneRuleResponseBodyData) *DeleteInterveneRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetHttpStatusCode(v int32) *DeleteInterveneRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetMessage(v string) *DeleteInterveneRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetRequestId(v string) *DeleteInterveneRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetSuccess(v bool) *DeleteInterveneRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteInterveneRuleResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// dt-s50ntwtywb4y
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteInterveneRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponseBodyData) SetFailIdList(v []*string) *DeleteInterveneRuleResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *DeleteInterveneRuleResponseBodyData) SetTaskId(v string) *DeleteInterveneRuleResponseBodyData {
	s.TaskId = &v
	return s
}

type DeleteInterveneRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInterveneRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInterveneRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponse) SetHeaders(v map[string]*string) *DeleteInterveneRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteInterveneRuleResponse) SetStatusCode(v int32) *DeleteInterveneRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInterveneRuleResponse) SetBody(v *DeleteInterveneRuleResponseBody) *DeleteInterveneRuleResponse {
	s.Body = v
	return s
}

type DeleteMaterialByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteMaterialByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdRequest) SetAgentKey(v string) *DeleteMaterialByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteMaterialByIdRequest) SetId(v int64) *DeleteMaterialByIdRequest {
	s.Id = &v
	return s
}

type DeleteMaterialByIdResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMaterialByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdResponseBody) SetCode(v string) *DeleteMaterialByIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetData(v bool) *DeleteMaterialByIdResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetHttpStatusCode(v int32) *DeleteMaterialByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetMessage(v string) *DeleteMaterialByIdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetRequestId(v string) *DeleteMaterialByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetSuccess(v bool) *DeleteMaterialByIdResponseBody {
	s.Success = &v
	return s
}

type DeleteMaterialByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaterialByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaterialByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdResponse) SetHeaders(v map[string]*string) *DeleteMaterialByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaterialByIdResponse) SetStatusCode(v int32) *DeleteMaterialByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaterialByIdResponse) SetBody(v *DeleteMaterialByIdResponseBody) *DeleteMaterialByIdResponse {
	s.Body = v
	return s
}

type DocumentExtractionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Urls []*string `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Repeated"`
}

func (s DocumentExtractionRequest) String() string {
	return tea.Prettify(s)
}

func (s DocumentExtractionRequest) GoString() string {
	return s.String()
}

func (s *DocumentExtractionRequest) SetAgentKey(v string) *DocumentExtractionRequest {
	s.AgentKey = &v
	return s
}

func (s *DocumentExtractionRequest) SetUrls(v []*string) *DocumentExtractionRequest {
	s.Urls = v
	return s
}

type DocumentExtractionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	UrlsShrink *string `json:"Urls,omitempty" xml:"Urls,omitempty"`
}

func (s DocumentExtractionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DocumentExtractionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DocumentExtractionShrinkRequest) SetAgentKey(v string) *DocumentExtractionShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *DocumentExtractionShrinkRequest) SetUrlsShrink(v string) *DocumentExtractionShrinkRequest {
	s.UrlsShrink = &v
	return s
}

type DocumentExtractionResponseBody struct {
	// example:
	//
	// NoData
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DocumentExtractionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DocumentExtractionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DocumentExtractionResponseBody) GoString() string {
	return s.String()
}

func (s *DocumentExtractionResponseBody) SetCode(v string) *DocumentExtractionResponseBody {
	s.Code = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetData(v []*DocumentExtractionResponseBodyData) *DocumentExtractionResponseBody {
	s.Data = v
	return s
}

func (s *DocumentExtractionResponseBody) SetHttpStatusCode(v int32) *DocumentExtractionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetMessage(v string) *DocumentExtractionResponseBody {
	s.Message = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetRequestId(v string) *DocumentExtractionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DocumentExtractionResponseBody) SetSuccess(v bool) *DocumentExtractionResponseBody {
	s.Success = &v
	return s
}

type DocumentExtractionResponseBodyData struct {
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
	// 8df2d69d63a247b6b52ff455b2d426b6
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-05-14 08:54:33
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

func (s DocumentExtractionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DocumentExtractionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DocumentExtractionResponseBodyData) SetAuthor(v string) *DocumentExtractionResponseBodyData {
	s.Author = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetContent(v string) *DocumentExtractionResponseBodyData {
	s.Content = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetDocId(v string) *DocumentExtractionResponseBodyData {
	s.DocId = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetDocUuid(v string) *DocumentExtractionResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetPubTime(v string) *DocumentExtractionResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetSource(v string) *DocumentExtractionResponseBodyData {
	s.Source = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetSummary(v string) *DocumentExtractionResponseBodyData {
	s.Summary = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetTag(v string) *DocumentExtractionResponseBodyData {
	s.Tag = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetTitle(v string) *DocumentExtractionResponseBodyData {
	s.Title = &v
	return s
}

func (s *DocumentExtractionResponseBodyData) SetUrl(v string) *DocumentExtractionResponseBodyData {
	s.Url = &v
	return s
}

type DocumentExtractionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocumentExtractionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocumentExtractionResponse) String() string {
	return tea.Prettify(s)
}

func (s DocumentExtractionResponse) GoString() string {
	return s.String()
}

func (s *DocumentExtractionResponse) SetHeaders(v map[string]*string) *DocumentExtractionResponse {
	s.Headers = v
	return s
}

func (s *DocumentExtractionResponse) SetStatusCode(v int32) *DocumentExtractionResponse {
	s.StatusCode = &v
	return s
}

func (s *DocumentExtractionResponse) SetBody(v *DocumentExtractionResponseBody) *DocumentExtractionResponse {
	s.Body = v
	return s
}

type ExportGeneratedContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ExportGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *ExportGeneratedContentRequest) SetAgentKey(v string) *ExportGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *ExportGeneratedContentRequest) SetId(v int64) *ExportGeneratedContentRequest {
	s.Id = &v
	return s
}

type ExportGeneratedContentResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *ExportGeneratedContentResponseBody) SetCode(v string) *ExportGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetData(v string) *ExportGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetHttpStatusCode(v int32) *ExportGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetMessage(v string) *ExportGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetRequestId(v string) *ExportGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetSuccess(v bool) *ExportGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type ExportGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *ExportGeneratedContentResponse) SetHeaders(v map[string]*string) *ExportGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *ExportGeneratedContentResponse) SetStatusCode(v int32) *ExportGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportGeneratedContentResponse) SetBody(v *ExportGeneratedContentResponseBody) *ExportGeneratedContentResponse {
	s.Body = v
	return s
}

type ExportHotTopicPlanningProposalsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 025c6cee437741368098b790c90166f8
	CustomViewPointIds []*string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty" type:"Repeated"`
	// example:
	//
	// 导出文档类型，word:导出为word,xmind:导处为xmind
	ExportType *string   `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	Titles     []*string `json:"Titles,omitempty" xml:"Titles,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 选题策划类型：CustomViewPoints:自定义视角，HotViewPoints:热门视角、TimedViewPoints:时效性视角、WebReviewPoints:网友视角、FreshViewPoints:新颖视角
	ViewPointType *string `json:"ViewPointType,omitempty" xml:"ViewPointType,omitempty"`
}

func (s ExportHotTopicPlanningProposalsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsRequest) GoString() string {
	return s.String()
}

func (s *ExportHotTopicPlanningProposalsRequest) SetAgentKey(v string) *ExportHotTopicPlanningProposalsRequest {
	s.AgentKey = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetCustomViewPointIds(v []*string) *ExportHotTopicPlanningProposalsRequest {
	s.CustomViewPointIds = v
	return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetExportType(v string) *ExportHotTopicPlanningProposalsRequest {
	s.ExportType = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetTitles(v []*string) *ExportHotTopicPlanningProposalsRequest {
	s.Titles = v
	return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetTopic(v string) *ExportHotTopicPlanningProposalsRequest {
	s.Topic = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetTopicSource(v string) *ExportHotTopicPlanningProposalsRequest {
	s.TopicSource = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsRequest) SetViewPointType(v string) *ExportHotTopicPlanningProposalsRequest {
	s.ViewPointType = &v
	return s
}

type ExportHotTopicPlanningProposalsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 025c6cee437741368098b790c90166f8
	CustomViewPointIdsShrink *string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty"`
	// example:
	//
	// 导出文档类型，word:导出为word,xmind:导处为xmind
	ExportType   *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	TitlesShrink *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 选题策划类型：CustomViewPoints:自定义视角，HotViewPoints:热门视角、TimedViewPoints:时效性视角、WebReviewPoints:网友视角、FreshViewPoints:新颖视角
	ViewPointType *string `json:"ViewPointType,omitempty" xml:"ViewPointType,omitempty"`
}

func (s ExportHotTopicPlanningProposalsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetAgentKey(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetCustomViewPointIdsShrink(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
	s.CustomViewPointIdsShrink = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetExportType(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
	s.ExportType = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetTitlesShrink(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
	s.TitlesShrink = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetTopic(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
	s.Topic = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetTopicSource(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
	s.TopicSource = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsShrinkRequest) SetViewPointType(v string) *ExportHotTopicPlanningProposalsShrinkRequest {
	s.ViewPointType = &v
	return s
}

type ExportHotTopicPlanningProposalsResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 业务数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportHotTopicPlanningProposalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetCode(v string) *ExportHotTopicPlanningProposalsResponseBody {
	s.Code = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetData(v string) *ExportHotTopicPlanningProposalsResponseBody {
	s.Data = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetHttpStatusCode(v int32) *ExportHotTopicPlanningProposalsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetMessage(v string) *ExportHotTopicPlanningProposalsResponseBody {
	s.Message = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetRequestId(v string) *ExportHotTopicPlanningProposalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsResponseBody) SetSuccess(v bool) *ExportHotTopicPlanningProposalsResponseBody {
	s.Success = &v
	return s
}

type ExportHotTopicPlanningProposalsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportHotTopicPlanningProposalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportHotTopicPlanningProposalsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportHotTopicPlanningProposalsResponse) GoString() string {
	return s.String()
}

func (s *ExportHotTopicPlanningProposalsResponse) SetHeaders(v map[string]*string) *ExportHotTopicPlanningProposalsResponse {
	s.Headers = v
	return s
}

func (s *ExportHotTopicPlanningProposalsResponse) SetStatusCode(v int32) *ExportHotTopicPlanningProposalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportHotTopicPlanningProposalsResponse) SetBody(v *ExportHotTopicPlanningProposalsResponseBody) *ExportHotTopicPlanningProposalsResponse {
	s.Body = v
	return s
}

type ExportIntervenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fed6555ec9e24b92aeecc34be484b887_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ExportIntervenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ExportIntervenesRequest) SetAgentKey(v string) *ExportIntervenesRequest {
	s.AgentKey = &v
	return s
}

type ExportIntervenesResponseBody struct {
	// example:
	//
	// 0
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ExportIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportIntervenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ExportIntervenesResponseBody) SetCode(v string) *ExportIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetData(v *ExportIntervenesResponseBodyData) *ExportIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ExportIntervenesResponseBody) SetHttpStatusCode(v int32) *ExportIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetMessage(v string) *ExportIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetRequestId(v string) *ExportIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetSuccess(v bool) *ExportIntervenesResponseBody {
	s.Success = &v
	return s
}

type ExportIntervenesResponseBodyData struct {
	// example:
	//
	// http://xxx/xxx.xls
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ExportIntervenesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExportIntervenesResponseBodyData) SetFileUrl(v string) *ExportIntervenesResponseBodyData {
	s.FileUrl = &v
	return s
}

type ExportIntervenesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportIntervenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ExportIntervenesResponse) SetHeaders(v map[string]*string) *ExportIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ExportIntervenesResponse) SetStatusCode(v int32) *ExportIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportIntervenesResponse) SetBody(v *ExportIntervenesResponseBody) *ExportIntervenesResponse {
	s.Body = v
	return s
}

type FeedbackDialogueRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// test
	CustomerResponse *string `json:"CustomerResponse,omitempty" xml:"CustomerResponse,omitempty"`
	// example:
	//
	// test
	GoodText *string `json:"GoodText,omitempty" xml:"GoodText,omitempty"`
	// example:
	//
	// test
	ModifiedResponse *string `json:"ModifiedResponse,omitempty" xml:"ModifiedResponse,omitempty"`
	// example:
	//
	// thumbsDown
	Rating     *string   `json:"Rating,omitempty" xml:"Rating,omitempty"`
	RatingTags []*string `json:"RatingTags,omitempty" xml:"RatingTags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 75bf82fa-b71b-45d7-ae40-0b00e496cd9e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FeedbackDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueRequest) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueRequest) SetAgentKey(v string) *FeedbackDialogueRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackDialogueRequest) SetCustomerResponse(v string) *FeedbackDialogueRequest {
	s.CustomerResponse = &v
	return s
}

func (s *FeedbackDialogueRequest) SetGoodText(v string) *FeedbackDialogueRequest {
	s.GoodText = &v
	return s
}

func (s *FeedbackDialogueRequest) SetModifiedResponse(v string) *FeedbackDialogueRequest {
	s.ModifiedResponse = &v
	return s
}

func (s *FeedbackDialogueRequest) SetRating(v string) *FeedbackDialogueRequest {
	s.Rating = &v
	return s
}

func (s *FeedbackDialogueRequest) SetRatingTags(v []*string) *FeedbackDialogueRequest {
	s.RatingTags = v
	return s
}

func (s *FeedbackDialogueRequest) SetSessionId(v string) *FeedbackDialogueRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackDialogueRequest) SetTaskId(v string) *FeedbackDialogueRequest {
	s.TaskId = &v
	return s
}

type FeedbackDialogueShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// test
	CustomerResponse *string `json:"CustomerResponse,omitempty" xml:"CustomerResponse,omitempty"`
	// example:
	//
	// test
	GoodText *string `json:"GoodText,omitempty" xml:"GoodText,omitempty"`
	// example:
	//
	// test
	ModifiedResponse *string `json:"ModifiedResponse,omitempty" xml:"ModifiedResponse,omitempty"`
	// example:
	//
	// thumbsDown
	Rating           *string `json:"Rating,omitempty" xml:"Rating,omitempty"`
	RatingTagsShrink *string `json:"RatingTags,omitempty" xml:"RatingTags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 75bf82fa-b71b-45d7-ae40-0b00e496cd9e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FeedbackDialogueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueShrinkRequest) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueShrinkRequest) SetAgentKey(v string) *FeedbackDialogueShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetCustomerResponse(v string) *FeedbackDialogueShrinkRequest {
	s.CustomerResponse = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetGoodText(v string) *FeedbackDialogueShrinkRequest {
	s.GoodText = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetModifiedResponse(v string) *FeedbackDialogueShrinkRequest {
	s.ModifiedResponse = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetRating(v string) *FeedbackDialogueShrinkRequest {
	s.Rating = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetRatingTagsShrink(v string) *FeedbackDialogueShrinkRequest {
	s.RatingTagsShrink = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetSessionId(v string) *FeedbackDialogueShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetTaskId(v string) *FeedbackDialogueShrinkRequest {
	s.TaskId = &v
	return s
}

type FeedbackDialogueResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FeedbackDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueResponseBody) SetCode(v string) *FeedbackDialogueResponseBody {
	s.Code = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetHttpStatusCode(v int32) *FeedbackDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetMessage(v string) *FeedbackDialogueResponseBody {
	s.Message = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetRequestId(v string) *FeedbackDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetSuccess(v bool) *FeedbackDialogueResponseBody {
	s.Success = &v
	return s
}

type FeedbackDialogueResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FeedbackDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueResponse) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueResponse) SetHeaders(v map[string]*string) *FeedbackDialogueResponse {
	s.Headers = v
	return s
}

func (s *FeedbackDialogueResponse) SetStatusCode(v int32) *FeedbackDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackDialogueResponse) SetBody(v *FeedbackDialogueResponseBody) *FeedbackDialogueResponse {
	s.Body = v
	return s
}

type FetchImageTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	ArticleTaskId *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["9d8c9185-3f75-4a20-aca1-c5bb53dd97b3"]
	TaskIdList []*string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty" type:"Repeated"`
}

func (s FetchImageTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskRequest) GoString() string {
	return s.String()
}

func (s *FetchImageTaskRequest) SetAgentKey(v string) *FetchImageTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *FetchImageTaskRequest) SetArticleTaskId(v string) *FetchImageTaskRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *FetchImageTaskRequest) SetTaskIdList(v []*string) *FetchImageTaskRequest {
	s.TaskIdList = v
	return s
}

type FetchImageTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	ArticleTaskId *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["9d8c9185-3f75-4a20-aca1-c5bb53dd97b3"]
	TaskIdListShrink *string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty"`
}

func (s FetchImageTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *FetchImageTaskShrinkRequest) SetAgentKey(v string) *FetchImageTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *FetchImageTaskShrinkRequest) SetArticleTaskId(v string) *FetchImageTaskShrinkRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *FetchImageTaskShrinkRequest) SetTaskIdListShrink(v string) *FetchImageTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

type FetchImageTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FetchImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DD656AF9-0839-521A-A3D2-F320009F9C87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchImageTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBody) SetCode(v string) *FetchImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetData(v *FetchImageTaskResponseBodyData) *FetchImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *FetchImageTaskResponseBody) SetHttpStatusCode(v int32) *FetchImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetMessage(v string) *FetchImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetRequestId(v string) *FetchImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetSuccess(v bool) *FetchImageTaskResponseBody {
	s.Success = &v
	return s
}

type FetchImageTaskResponseBodyData struct {
	TaskInfoList []*FetchImageTaskResponseBodyDataTaskInfoList `json:"TaskInfoList,omitempty" xml:"TaskInfoList,omitempty" type:"Repeated"`
}

func (s FetchImageTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyData) SetTaskInfoList(v []*FetchImageTaskResponseBodyDataTaskInfoList) *FetchImageTaskResponseBodyData {
	s.TaskInfoList = v
	return s
}

type FetchImageTaskResponseBodyDataTaskInfoList struct {
	// example:
	//
	// 1
	Id        *int64                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageList []*FetchImageTaskResponseBodyDataTaskInfoListImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// example:
	//
	// net-7eb32699000d4193a3c59fc64ae1e55f
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s FetchImageTaskResponseBodyDataTaskInfoList) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBodyDataTaskInfoList) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetId(v int64) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.Id = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetImageList(v []*FetchImageTaskResponseBodyDataTaskInfoListImageList) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.ImageList = v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetTaskId(v string) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.TaskId = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetTaskStatus(v string) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.TaskStatus = &v
	return s
}

type FetchImageTaskResponseBodyDataTaskInfoListImageList struct {
	// example:
	//
	// NoData
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// https://a-hbr-temp-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/r-000aham3nsx9gc7a8r5l.csv?Expires=1678260131&OSSAccessKeyId=LTAIjGotF8wXIEjy&Signature=WpMfqBnjeR0w5UL1xFAd1J556Pw%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s FetchImageTaskResponseBodyDataTaskInfoListImageList) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBodyDataTaskInfoListImageList) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetCode(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Code = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetMessage(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Message = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetUrl(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Url = &v
	return s
}

type FetchImageTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchImageTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponse) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponse) SetHeaders(v map[string]*string) *FetchImageTaskResponse {
	s.Headers = v
	return s
}

func (s *FetchImageTaskResponse) SetStatusCode(v int32) *FetchImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchImageTaskResponse) SetBody(v *FetchImageTaskResponseBody) *FetchImageTaskResponse {
	s.Body = v
	return s
}

type GenerateFileUrlByKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GenerateFileUrlByKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateFileUrlByKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyRequest) SetAgentKey(v string) *GenerateFileUrlByKeyRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateFileUrlByKeyRequest) SetFileKey(v string) *GenerateFileUrlByKeyRequest {
	s.FileKey = &v
	return s
}

func (s *GenerateFileUrlByKeyRequest) SetFileName(v string) *GenerateFileUrlByKeyRequest {
	s.FileName = &v
	return s
}

type GenerateFileUrlByKeyResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://www.example.com/a.txt
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateFileUrlByKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateFileUrlByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyResponseBody) SetCode(v string) *GenerateFileUrlByKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetData(v string) *GenerateFileUrlByKeyResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetHttpStatusCode(v int32) *GenerateFileUrlByKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetMessage(v string) *GenerateFileUrlByKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetRequestId(v string) *GenerateFileUrlByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetSuccess(v bool) *GenerateFileUrlByKeyResponseBody {
	s.Success = &v
	return s
}

type GenerateFileUrlByKeyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateFileUrlByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateFileUrlByKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateFileUrlByKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyResponse) SetHeaders(v map[string]*string) *GenerateFileUrlByKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateFileUrlByKeyResponse) SetStatusCode(v int32) *GenerateFileUrlByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateFileUrlByKeyResponse) SetBody(v *GenerateFileUrlByKeyResponseBody) *GenerateFileUrlByKeyResponse {
	s.Body = v
	return s
}

type GenerateImageTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	ArticleTaskId *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	// This parameter is required.
	ParagraphList []*GenerateImageTaskRequestParagraphList `json:"ParagraphList,omitempty" xml:"ParagraphList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1024*1024
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <auto>
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s GenerateImageTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskRequest) SetAgentKey(v string) *GenerateImageTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateImageTaskRequest) SetArticleTaskId(v string) *GenerateImageTaskRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *GenerateImageTaskRequest) SetParagraphList(v []*GenerateImageTaskRequestParagraphList) *GenerateImageTaskRequest {
	s.ParagraphList = v
	return s
}

func (s *GenerateImageTaskRequest) SetSize(v string) *GenerateImageTaskRequest {
	s.Size = &v
	return s
}

func (s *GenerateImageTaskRequest) SetStyle(v string) *GenerateImageTaskRequest {
	s.Style = &v
	return s
}

type GenerateImageTaskRequestParagraphList struct {
	// This parameter is required.
	//
	// example:
	//
	// 一直忧伤的猫
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GenerateImageTaskRequestParagraphList) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskRequestParagraphList) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskRequestParagraphList) SetContent(v string) *GenerateImageTaskRequestParagraphList {
	s.Content = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetId(v int64) *GenerateImageTaskRequestParagraphList {
	s.Id = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetTaskId(v string) *GenerateImageTaskRequestParagraphList {
	s.TaskId = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetTaskStatus(v string) *GenerateImageTaskRequestParagraphList {
	s.TaskStatus = &v
	return s
}

type GenerateImageTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	ArticleTaskId *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	// This parameter is required.
	ParagraphListShrink *string `json:"ParagraphList,omitempty" xml:"ParagraphList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1024*1024
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <auto>
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s GenerateImageTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskShrinkRequest) SetAgentKey(v string) *GenerateImageTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetArticleTaskId(v string) *GenerateImageTaskShrinkRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetParagraphListShrink(v string) *GenerateImageTaskShrinkRequest {
	s.ParagraphListShrink = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetSize(v string) *GenerateImageTaskShrinkRequest {
	s.Size = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetStyle(v string) *GenerateImageTaskShrinkRequest {
	s.Style = &v
	return s
}

type GenerateImageTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GenerateImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateImageTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBody) SetCode(v string) *GenerateImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetData(v *GenerateImageTaskResponseBodyData) *GenerateImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *GenerateImageTaskResponseBody) SetHttpStatusCode(v int32) *GenerateImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetMessage(v string) *GenerateImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetRequestId(v string) *GenerateImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetSuccess(v bool) *GenerateImageTaskResponseBody {
	s.Success = &v
	return s
}

type GenerateImageTaskResponseBodyData struct {
	TaskList []*GenerateImageTaskResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s GenerateImageTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBodyData) SetTaskList(v []*GenerateImageTaskResponseBodyDataTaskList) *GenerateImageTaskResponseBodyData {
	s.TaskList = v
	return s
}

type GenerateImageTaskResponseBodyDataTaskList struct {
	// example:
	//
	// 一直忧伤的猫
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GenerateImageTaskResponseBodyDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetContent(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.Content = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetId(v int64) *GenerateImageTaskResponseBodyDataTaskList {
	s.Id = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetTaskId(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.TaskId = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetTaskStatus(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.TaskStatus = &v
	return s
}

type GenerateImageTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateImageTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponse) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponse) SetHeaders(v map[string]*string) *GenerateImageTaskResponse {
	s.Headers = v
	return s
}

func (s *GenerateImageTaskResponse) SetStatusCode(v int32) *GenerateImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateImageTaskResponse) SetBody(v *GenerateImageTaskResponseBody) *GenerateImageTaskResponse {
	s.Body = v
	return s
}

type GenerateUploadConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// test.docx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// datasetUpload
	ParentDir *string `json:"ParentDir,omitempty" xml:"ParentDir,omitempty"`
}

func (s GenerateUploadConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigRequest) SetAgentKey(v string) *GenerateUploadConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateUploadConfigRequest) SetFileName(v string) *GenerateUploadConfigRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadConfigRequest) SetParentDir(v string) *GenerateUploadConfigRequest {
	s.ParentDir = &v
	return s
}

type GenerateUploadConfigResponseBody struct {
	// example:
	//
	// NoData
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GenerateUploadConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponseBody) SetCode(v string) *GenerateUploadConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetData(v *GenerateUploadConfigResponseBodyData) *GenerateUploadConfigResponseBody {
	s.Data = v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetHttpStatusCode(v int32) *GenerateUploadConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetMessage(v string) *GenerateUploadConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetRequestId(v string) *GenerateUploadConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetSuccess(v bool) *GenerateUploadConfigResponseBody {
	s.Success = &v
	return s
}

type GenerateUploadConfigResponseBodyData struct {
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	FileKey   *string                `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FormDatas map[string]interface{} `json:"FormDatas,omitempty" xml:"FormDatas,omitempty"`
	// example:
	//
	// https://bucket-name.oss-cn-hangzhou.aliyuncs.com
	PostUrl *string `json:"PostUrl,omitempty" xml:"PostUrl,omitempty"`
}

func (s GenerateUploadConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponseBodyData) SetFileKey(v string) *GenerateUploadConfigResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *GenerateUploadConfigResponseBodyData) SetFormDatas(v map[string]interface{}) *GenerateUploadConfigResponseBodyData {
	s.FormDatas = v
	return s
}

func (s *GenerateUploadConfigResponseBodyData) SetPostUrl(v string) *GenerateUploadConfigResponseBodyData {
	s.PostUrl = &v
	return s
}

type GenerateUploadConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponse) SetHeaders(v map[string]*string) *GenerateUploadConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadConfigResponse) SetStatusCode(v int32) *GenerateUploadConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadConfigResponse) SetBody(v *GenerateUploadConfigResponseBody) *GenerateUploadConfigResponse {
	s.Body = v
	return s
}

type GenerateViewPointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey      *string                                `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReferenceData *GenerateViewPointRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
}

func (s GenerateViewPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointRequest) GoString() string {
	return s.String()
}

func (s *GenerateViewPointRequest) SetAgentKey(v string) *GenerateViewPointRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateViewPointRequest) SetReferenceData(v *GenerateViewPointRequestReferenceData) *GenerateViewPointRequest {
	s.ReferenceData = v
	return s
}

type GenerateViewPointRequestReferenceData struct {
	MiniDoc []*string `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
}

func (s GenerateViewPointRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointRequestReferenceData) GoString() string {
	return s.String()
}

func (s *GenerateViewPointRequestReferenceData) SetMiniDoc(v []*string) *GenerateViewPointRequestReferenceData {
	s.MiniDoc = v
	return s
}

type GenerateViewPointShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey            *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
}

func (s GenerateViewPointShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateViewPointShrinkRequest) SetAgentKey(v string) *GenerateViewPointShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateViewPointShrinkRequest) SetReferenceDataShrink(v string) *GenerateViewPointShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

type GenerateViewPointResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GenerateViewPointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateViewPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponseBody) SetCode(v string) *GenerateViewPointResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetData(v []*GenerateViewPointResponseBodyData) *GenerateViewPointResponseBody {
	s.Data = v
	return s
}

func (s *GenerateViewPointResponseBody) SetHttpStatusCode(v int32) *GenerateViewPointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetMessage(v string) *GenerateViewPointResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetRequestId(v string) *GenerateViewPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetSuccess(v bool) *GenerateViewPointResponseBody {
	s.Success = &v
	return s
}

type GenerateViewPointResponseBodyData struct {
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
}

func (s GenerateViewPointResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponseBodyData) SetPoint(v string) *GenerateViewPointResponseBodyData {
	s.Point = &v
	return s
}

type GenerateViewPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateViewPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateViewPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointResponse) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponse) SetHeaders(v map[string]*string) *GenerateViewPointResponse {
	s.Headers = v
	return s
}

func (s *GenerateViewPointResponse) SetStatusCode(v int32) *GenerateViewPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateViewPointResponse) SetBody(v *GenerateViewPointResponseBody) *GenerateViewPointResponse {
	s.Body = v
	return s
}

type GetCustomTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 商品code
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 63
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCustomTextRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTextRequest) GoString() string {
	return s.String()
}

func (s *GetCustomTextRequest) SetAgentKey(v string) *GetCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *GetCustomTextRequest) SetCommodityCode(v string) *GetCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetCustomTextRequest) SetId(v int64) *GetCustomTextRequest {
	s.Id = &v
	return s
}

type GetCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCustomTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomTextResponseBody) SetCode(v string) *GetCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomTextResponseBody) SetData(v *GetCustomTextResponseBodyData) *GetCustomTextResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomTextResponseBody) SetHttpStatusCode(v int32) *GetCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomTextResponseBody) SetMessage(v string) *GetCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomTextResponseBody) SetRequestId(v string) *GetCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomTextResponseBody) SetSuccess(v bool) *GetCustomTextResponseBody {
	s.Success = &v
	return s
}

type GetCustomTextResponseBodyData struct {
	// example:
	//
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 创建用户
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 34
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 修改用户
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s GetCustomTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomTextResponseBodyData) SetContent(v string) *GetCustomTextResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetCreateTime(v string) *GetCustomTextResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetCreateUser(v string) *GetCustomTextResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetId(v int64) *GetCustomTextResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetTitle(v string) *GetCustomTextResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetUpdateTime(v string) *GetCustomTextResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetCustomTextResponseBodyData) SetUpdateUser(v string) *GetCustomTextResponseBodyData {
	s.UpdateUser = &v
	return s
}

type GetCustomTextResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomTextResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTextResponse) GoString() string {
	return s.String()
}

func (s *GetCustomTextResponse) SetHeaders(v map[string]*string) *GetCustomTextResponse {
	s.Headers = v
	return s
}

func (s *GetCustomTextResponse) SetStatusCode(v int32) *GetCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomTextResponse) SetBody(v *GetCustomTextResponseBody) *GetCustomTextResponse {
	s.Body = v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0dbf1055f8a2475d99904c3b76a0ffba
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetTaskId(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	CustomViewPointsResult *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult `json:"CustomViewPointsResult,omitempty" xml:"CustomViewPointsResult,omitempty" type:"Struct"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetCustomViewPointsResult(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.CustomViewPointsResult = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetStatus(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult struct {
	Attitudes []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
	// example:
	//
	// 热点主题事件
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) SetAttitudes(v []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult {
	s.Attitudes = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult) SetTopic(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResult {
	s.Topic = &v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                                                                                    `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetAttitude(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetAttitudeType(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetRatio(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes) SetViewPoints(v []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyDataCustomViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *GetCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

type GetDataSourceOrderConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d9a1f6146a37446495d9985c2e7b267e_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miaobi
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetDataSourceOrderConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigRequest) SetAgentKey(v string) *GetDataSourceOrderConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *GetDataSourceOrderConfigRequest) SetProductCode(v string) *GetDataSourceOrderConfigRequest {
	s.ProductCode = &v
	return s
}

type GetDataSourceOrderConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDataSourceOrderConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataSourceOrderConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBody) SetCode(v string) *GetDataSourceOrderConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetData(v *GetDataSourceOrderConfigResponseBodyData) *GetDataSourceOrderConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetHttpStatusCode(v int32) *GetDataSourceOrderConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetMessage(v string) *GetDataSourceOrderConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetRequestId(v string) *GetDataSourceOrderConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetSuccess(v bool) *GetDataSourceOrderConfigResponseBody {
	s.Success = &v
	return s
}

type GetDataSourceOrderConfigResponseBodyData struct {
	UserConfigDataSourceList []*GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty" type:"Repeated"`
}

func (s GetDataSourceOrderConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBodyData) SetUserConfigDataSourceList(v []*GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) *GetDataSourceOrderConfigResponseBodyData {
	s.UserConfigDataSourceList = v
	return s
}

type GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList struct {
	// example:
	//
	// QuarkCommonNews
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 20
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// SystemSearch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetCode(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Code = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetName(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Name = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetNumber(v int32) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Number = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetType(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Type = &v
	return s
}

type GetDataSourceOrderConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceOrderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceOrderConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponse) SetHeaders(v map[string]*string) *GetDataSourceOrderConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceOrderConfigResponse) SetStatusCode(v int32) *GetDataSourceOrderConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceOrderConfigResponse) SetBody(v *GetDataSourceOrderConfigResponseBody) *GetDataSourceOrderConfigResponse {
	s.Body = v
	return s
}

type GetDocClusterTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 93771c8e1142467fb1aedf1763feba1e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDocClusterTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskRequest) SetAgentKey(v string) *GetDocClusterTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GetDocClusterTaskRequest) SetTaskId(v string) *GetDocClusterTaskRequest {
	s.TaskId = &v
	return s
}

type GetDocClusterTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDocClusterTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocClusterTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponseBody) SetCode(v string) *GetDocClusterTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetData(v *GetDocClusterTaskResponseBodyData) *GetDocClusterTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetHttpStatusCode(v int32) *GetDocClusterTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetMessage(v string) *GetDocClusterTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetRequestId(v string) *GetDocClusterTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocClusterTaskResponseBody) SetSuccess(v bool) *GetDocClusterTaskResponseBody {
	s.Success = &v
	return s
}

type GetDocClusterTaskResponseBodyData struct {
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// PENDING
	Status *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Topics []*GetDocClusterTaskResponseBodyDataTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s GetDocClusterTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocClusterTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponseBodyData) SetErrorMessage(v string) *GetDocClusterTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetDocClusterTaskResponseBodyData) SetStatus(v string) *GetDocClusterTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDocClusterTaskResponseBodyData) SetTopics(v []*GetDocClusterTaskResponseBodyDataTopics) *GetDocClusterTaskResponseBodyData {
	s.Topics = v
	return s
}

type GetDocClusterTaskResponseBodyDataTopics struct {
	DocIds []*string `json:"DocIds,omitempty" xml:"DocIds,omitempty" type:"Repeated"`
	// example:
	//
	// 聚类主题摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 聚类主题名
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetDocClusterTaskResponseBodyDataTopics) String() string {
	return tea.Prettify(s)
}

func (s GetDocClusterTaskResponseBodyDataTopics) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponseBodyDataTopics) SetDocIds(v []*string) *GetDocClusterTaskResponseBodyDataTopics {
	s.DocIds = v
	return s
}

func (s *GetDocClusterTaskResponseBodyDataTopics) SetSummary(v string) *GetDocClusterTaskResponseBodyDataTopics {
	s.Summary = &v
	return s
}

func (s *GetDocClusterTaskResponseBodyDataTopics) SetTitle(v string) *GetDocClusterTaskResponseBodyDataTopics {
	s.Title = &v
	return s
}

type GetDocClusterTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocClusterTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskResponse) SetHeaders(v map[string]*string) *GetDocClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDocClusterTaskResponse) SetStatusCode(v int32) *GetDocClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocClusterTaskResponse) SetBody(v *GetDocClusterTaskResponseBody) *GetDocClusterTaskResponse {
	s.Body = v
	return s
}

type GetGeneratedContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentRequest) SetAgentKey(v string) *GetGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *GetGeneratedContentRequest) SetId(v int64) *GetGeneratedContentRequest {
	s.Id = &v
	return s
}

type GetGeneratedContentResponseBody struct {
	// example:
	//
	// NoData
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetGeneratedContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponseBody) SetCode(v string) *GetGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetData(v *GetGeneratedContentResponseBodyData) *GetGeneratedContentResponseBody {
	s.Data = v
	return s
}

func (s *GetGeneratedContentResponseBody) SetHttpStatusCode(v int32) *GetGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetMessage(v string) *GetGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetRequestId(v string) *GetGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetSuccess(v bool) *GetGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type GetGeneratedContentResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// media
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText   *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// xxx
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 86
	Id          *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordList []*string `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	Keywords    *string   `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Prompt      *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// example:
	//
	// 0961a514-2e26-4aa6-b22b-f592d145fe47
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetGeneratedContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponseBodyData) SetContent(v string) *GetGeneratedContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetContentDomain(v string) *GetGeneratedContentResponseBodyData {
	s.ContentDomain = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetContentText(v string) *GetGeneratedContentResponseBodyData {
	s.ContentText = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetCreateTime(v string) *GetGeneratedContentResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetCreateUser(v string) *GetGeneratedContentResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetDeviceId(v string) *GetGeneratedContentResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetId(v int64) *GetGeneratedContentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetKeywordList(v []*string) *GetGeneratedContentResponseBodyData {
	s.KeywordList = v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetKeywords(v string) *GetGeneratedContentResponseBodyData {
	s.Keywords = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetPrompt(v string) *GetGeneratedContentResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetTaskId(v string) *GetGeneratedContentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetTitle(v string) *GetGeneratedContentResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUpdateTime(v string) *GetGeneratedContentResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUpdateUser(v string) *GetGeneratedContentResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUuid(v string) *GetGeneratedContentResponseBodyData {
	s.Uuid = &v
	return s
}

type GetGeneratedContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponse) SetHeaders(v map[string]*string) *GetGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *GetGeneratedContentResponse) SetStatusCode(v int32) *GetGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGeneratedContentResponse) SetBody(v *GetGeneratedContentResponseBody) *GetGeneratedContentResponse {
	s.Body = v
	return s
}

type GetInterveneGlobalReplyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetInterveneGlobalReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyRequest) SetAgentKey(v string) *GetInterveneGlobalReplyRequest {
	s.AgentKey = &v
	return s
}

type GetInterveneGlobalReplyResponseBody struct {
	// example:
	//
	// 0
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneGlobalReplyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneGlobalReplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBody) SetCode(v string) *GetInterveneGlobalReplyResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetData(v *GetInterveneGlobalReplyResponseBodyData) *GetInterveneGlobalReplyResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetHttpStatusCode(v int32) *GetInterveneGlobalReplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetMessage(v string) *GetInterveneGlobalReplyResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetRequestId(v string) *GetInterveneGlobalReplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetSuccess(v bool) *GetInterveneGlobalReplyResponseBody {
	s.Success = &v
	return s
}

type GetInterveneGlobalReplyResponseBodyData struct {
	ReplyMessagList []*GetInterveneGlobalReplyResponseBodyDataReplyMessagList `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty" type:"Repeated"`
}

func (s GetInterveneGlobalReplyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBodyData) SetReplyMessagList(v []*GetInterveneGlobalReplyResponseBodyDataReplyMessagList) *GetInterveneGlobalReplyResponseBodyData {
	s.ReplyMessagList = v
	return s
}

type GetInterveneGlobalReplyResponseBodyDataReplyMessagList struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// namespace_qa_query
	ReplyType *string `json:"ReplyType,omitempty" xml:"ReplyType,omitempty"`
}

func (s GetInterveneGlobalReplyResponseBodyDataReplyMessagList) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBodyDataReplyMessagList) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) SetMessage(v string) *GetInterveneGlobalReplyResponseBodyDataReplyMessagList {
	s.Message = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) SetReplyType(v string) *GetInterveneGlobalReplyResponseBodyDataReplyMessagList {
	s.ReplyType = &v
	return s
}

type GetInterveneGlobalReplyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneGlobalReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneGlobalReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponse) SetHeaders(v map[string]*string) *GetInterveneGlobalReplyResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneGlobalReplyResponse) SetStatusCode(v int32) *GetInterveneGlobalReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneGlobalReplyResponse) SetBody(v *GetInterveneGlobalReplyResponseBody) *GetInterveneGlobalReplyResponse {
	s.Body = v
	return s
}

type GetInterveneImportTaskInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 19162157
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetInterveneImportTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoRequest) SetAgentKey(v string) *GetInterveneImportTaskInfoRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneImportTaskInfoRequest) SetTaskId(v string) *GetInterveneImportTaskInfoRequest {
	s.TaskId = &v
	return s
}

type GetInterveneImportTaskInfoResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneImportTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneImportTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBody) SetCode(v string) *GetInterveneImportTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetData(v *GetInterveneImportTaskInfoResponseBodyData) *GetInterveneImportTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetHttpStatusCode(v int32) *GetInterveneImportTaskInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetMessage(v string) *GetInterveneImportTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetRequestId(v string) *GetInterveneImportTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetSuccess(v bool) *GetInterveneImportTaskInfoResponseBody {
	s.Success = &v
	return s
}

type GetInterveneImportTaskInfoResponseBodyData struct {
	Status *GetInterveneImportTaskInfoResponseBodyDataStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s GetInterveneImportTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBodyData) SetStatus(v *GetInterveneImportTaskInfoResponseBodyDataStatus) *GetInterveneImportTaskInfoResponseBodyData {
	s.Status = v
	return s
}

type GetInterveneImportTaskInfoResponseBodyDataStatus struct {
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 80
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 41405255
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ft-task-20190101m8rnK
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetInterveneImportTaskInfoResponseBodyDataStatus) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetMsg(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Msg = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetPercentage(v int32) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Percentage = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetStatus(v int32) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Status = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetTaskId(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.TaskId = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetTaskName(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.TaskName = &v
	return s
}

type GetInterveneImportTaskInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneImportTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneImportTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponse) SetHeaders(v map[string]*string) *GetInterveneImportTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneImportTaskInfoResponse) SetStatusCode(v int32) *GetInterveneImportTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponse) SetBody(v *GetInterveneImportTaskInfoResponseBody) *GetInterveneImportTaskInfoResponse {
	s.Body = v
	return s
}

type GetInterveneRuleDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daaa2e0c209xb26acb97009ea77bd4b_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 12345
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetInterveneRuleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailRequest) SetAgentKey(v string) *GetInterveneRuleDetailRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneRuleDetailRequest) SetRuleId(v int64) *GetInterveneRuleDetailRequest {
	s.RuleId = &v
	return s
}

type GetInterveneRuleDetailResponseBody struct {
	// example:
	//
	// 0
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneRuleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBody) SetCode(v string) *GetInterveneRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetData(v *GetInterveneRuleDetailResponseBodyData) *GetInterveneRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetHttpStatusCode(v int32) *GetInterveneRuleDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetMessage(v string) *GetInterveneRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetRequestId(v string) *GetInterveneRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetSuccess(v bool) *GetInterveneRuleDetailResponseBody {
	s.Success = &v
	return s
}

type GetInterveneRuleDetailResponseBodyData struct {
	InterveneRuleDetail *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail `json:"InterveneRuleDetail,omitempty" xml:"InterveneRuleDetail,omitempty" type:"Struct"`
}

func (s GetInterveneRuleDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyData) SetInterveneRuleDetail(v *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) *GetInterveneRuleDetailResponseBodyData {
	s.InterveneRuleDetail = v
	return s
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail struct {
	AnswerConfig []*GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	EffectConfig *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig   `json:"EffectConfig,omitempty" xml:"EffectConfig,omitempty" type:"Struct"`
	// example:
	//
	// 0
	InterveneType *int32    `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// example:
	//
	// 100418
	RuleId   *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetAnswerConfig(v []*GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.AnswerConfig = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetEffectConfig(v *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.EffectConfig = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetInterveneType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.InterveneType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetNamespaceList(v []*string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.NamespaceList = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetRuleId(v int64) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.RuleId = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetRuleName(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.RuleName = &v
	return s
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig struct {
	// example:
	//
	// 0
	AnswerType *int32  `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// namespace_qa_query
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetAnswerType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetMessage(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.Message = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetNamespace(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.Namespace = &v
	return s
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig struct {
	// example:
	//
	// 0
	EffectType *int32 `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	// example:
	//
	// 2023-11-25 14:21:15
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2023-11-25 14:21:15
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetEffectType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.EffectType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetEndTime(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.EndTime = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetStartTime(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.StartTime = &v
	return s
}

type GetInterveneRuleDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneRuleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponse) SetHeaders(v map[string]*string) *GetInterveneRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneRuleDetailResponse) SetStatusCode(v int32) *GetInterveneRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneRuleDetailResponse) SetBody(v *GetInterveneRuleDetailResponseBody) *GetInterveneRuleDetailResponse {
	s.Body = v
	return s
}

type GetInterveneTemplateFileUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetInterveneTemplateFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlRequest) SetAgentKey(v string) *GetInterveneTemplateFileUrlRequest {
	s.AgentKey = &v
	return s
}

type GetInterveneTemplateFileUrlResponseBody struct {
	// example:
	//
	// 0
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneTemplateFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetCode(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetData(v *GetInterveneTemplateFileUrlResponseBodyData) *GetInterveneTemplateFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetHttpStatusCode(v int32) *GetInterveneTemplateFileUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetMessage(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetRequestId(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetSuccess(v bool) *GetInterveneTemplateFileUrlResponseBody {
	s.Success = &v
	return s
}

type GetInterveneTemplateFileUrlResponseBodyData struct {
	// example:
	//
	// http://xxx/xxx.xls
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponseBodyData) SetFileUrl(v string) *GetInterveneTemplateFileUrlResponseBodyData {
	s.FileUrl = &v
	return s
}

type GetInterveneTemplateFileUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneTemplateFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponse) SetHeaders(v map[string]*string) *GetInterveneTemplateFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneTemplateFileUrlResponse) SetStatusCode(v int32) *GetInterveneTemplateFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponse) SetBody(v *GetInterveneTemplateFileUrlResponseBody) *GetInterveneTemplateFileUrlResponse {
	s.Body = v
	return s
}

type GetMaterialByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMaterialByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdRequest) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdRequest) SetAgentKey(v string) *GetMaterialByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *GetMaterialByIdRequest) SetId(v int64) *GetMaterialByIdRequest {
	s.Id = &v
	return s
}

type GetMaterialByIdResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetMaterialByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMaterialByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponseBody) SetCode(v string) *GetMaterialByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetData(v *GetMaterialByIdResponseBodyData) *GetMaterialByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetMaterialByIdResponseBody) SetHttpStatusCode(v int32) *GetMaterialByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetMessage(v string) *GetMaterialByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetRequestId(v string) *GetMaterialByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetSuccess(v bool) *GetMaterialByIdResponseBody {
	s.Success = &v
	return s
}

type GetMaterialByIdResponseBodyData struct {
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 2023-03-21 11:34:19
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser  *string   `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DocKeywords []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// https://www.example.com
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// example:
	//
	// 32
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// https://www.example.com
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// user_upload
	SrcFrom           *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary           *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent       *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	ThumbnailInBase64 *string `json:"ThumbnailInBase64,omitempty" xml:"ThumbnailInBase64,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2022-04-08 19:33:01
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMaterialByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponseBodyData) SetAuthor(v string) *GetMaterialByIdResponseBodyData {
	s.Author = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetCreateTime(v string) *GetMaterialByIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetCreateUser(v string) *GetMaterialByIdResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetDocKeywords(v []*string) *GetMaterialByIdResponseBodyData {
	s.DocKeywords = v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetDocType(v string) *GetMaterialByIdResponseBodyData {
	s.DocType = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetExternalUrl(v string) *GetMaterialByIdResponseBodyData {
	s.ExternalUrl = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetHtmlContent(v string) *GetMaterialByIdResponseBodyData {
	s.HtmlContent = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetId(v int64) *GetMaterialByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetPubTime(v string) *GetMaterialByIdResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetPublicUrl(v string) *GetMaterialByIdResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetShareAttr(v int32) *GetMaterialByIdResponseBodyData {
	s.ShareAttr = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetSrcFrom(v string) *GetMaterialByIdResponseBodyData {
	s.SrcFrom = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetSummary(v string) *GetMaterialByIdResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetTextContent(v string) *GetMaterialByIdResponseBodyData {
	s.TextContent = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetThumbnailInBase64(v string) *GetMaterialByIdResponseBodyData {
	s.ThumbnailInBase64 = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetTitle(v string) *GetMaterialByIdResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUpdateTime(v string) *GetMaterialByIdResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUpdateUser(v string) *GetMaterialByIdResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUrl(v string) *GetMaterialByIdResponseBodyData {
	s.Url = &v
	return s
}

type GetMaterialByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMaterialByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMaterialByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdResponse) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponse) SetHeaders(v map[string]*string) *GetMaterialByIdResponse {
	s.Headers = v
	return s
}

func (s *GetMaterialByIdResponse) SetStatusCode(v int32) *GetMaterialByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaterialByIdResponse) SetBody(v *GetMaterialByIdResponseBody) *GetMaterialByIdResponse {
	s.Body = v
	return s
}

type GetPropertiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetPropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetPropertiesRequest) SetAgentKey(v string) *GetPropertiesRequest {
	s.AgentKey = &v
	return s
}

type GetPropertiesResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPropertiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBody) SetCode(v string) *GetPropertiesResponseBody {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBody) SetData(v *GetPropertiesResponseBodyData) *GetPropertiesResponseBody {
	s.Data = v
	return s
}

func (s *GetPropertiesResponseBody) SetHttpStatusCode(v int32) *GetPropertiesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPropertiesResponseBody) SetMessage(v string) *GetPropertiesResponseBody {
	s.Message = &v
	return s
}

func (s *GetPropertiesResponseBody) SetRequestId(v string) *GetPropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPropertiesResponseBody) SetSuccess(v bool) *GetPropertiesResponseBody {
	s.Success = &v
	return s
}

type GetPropertiesResponseBodyData struct {
	ChatConfig              map[string]interface{}                                `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	ConsoleConfig           *GetPropertiesResponseBodyDataConsoleConfig           `json:"ConsoleConfig,omitempty" xml:"ConsoleConfig,omitempty" type:"Struct"`
	GeneralConfigMap        map[string]interface{}                                `json:"GeneralConfigMap,omitempty" xml:"GeneralConfigMap,omitempty"`
	IntelligentSearchConfig *GetPropertiesResponseBodyDataIntelligentSearchConfig `json:"IntelligentSearchConfig,omitempty" xml:"IntelligentSearchConfig,omitempty" type:"Struct"`
	SearchSources           []*GetPropertiesResponseBodyDataSearchSources         `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SlrAuthorized            *bool                                                    `json:"SlrAuthorized,omitempty" xml:"SlrAuthorized,omitempty"`
	UserInfo                 *GetPropertiesResponseBodyDataUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	WanxiangImageSizeConfig  []*GetPropertiesResponseBodyDataWanxiangImageSizeConfig  `json:"WanxiangImageSizeConfig,omitempty" xml:"WanxiangImageSizeConfig,omitempty" type:"Repeated"`
	WanxiangImageStyleConfig []*GetPropertiesResponseBodyDataWanxiangImageStyleConfig `json:"WanxiangImageStyleConfig,omitempty" xml:"WanxiangImageStyleConfig,omitempty" type:"Repeated"`
}

func (s GetPropertiesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyData) SetChatConfig(v map[string]interface{}) *GetPropertiesResponseBodyData {
	s.ChatConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetConsoleConfig(v *GetPropertiesResponseBodyDataConsoleConfig) *GetPropertiesResponseBodyData {
	s.ConsoleConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetGeneralConfigMap(v map[string]interface{}) *GetPropertiesResponseBodyData {
	s.GeneralConfigMap = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetIntelligentSearchConfig(v *GetPropertiesResponseBodyDataIntelligentSearchConfig) *GetPropertiesResponseBodyData {
	s.IntelligentSearchConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetSearchSources(v []*GetPropertiesResponseBodyDataSearchSources) *GetPropertiesResponseBodyData {
	s.SearchSources = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetSlrAuthorized(v bool) *GetPropertiesResponseBodyData {
	s.SlrAuthorized = &v
	return s
}

func (s *GetPropertiesResponseBodyData) SetUserInfo(v *GetPropertiesResponseBodyDataUserInfo) *GetPropertiesResponseBodyData {
	s.UserInfo = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetWanxiangImageSizeConfig(v []*GetPropertiesResponseBodyDataWanxiangImageSizeConfig) *GetPropertiesResponseBodyData {
	s.WanxiangImageSizeConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetWanxiangImageStyleConfig(v []*GetPropertiesResponseBodyDataWanxiangImageStyleConfig) *GetPropertiesResponseBodyData {
	s.WanxiangImageStyleConfig = v
	return s
}

type GetPropertiesResponseBodyDataConsoleConfig struct {
	// example:
	//
	// xx
	TipContent *string `json:"TipContent,omitempty" xml:"TipContent,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetPropertiesResponseBodyDataConsoleConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataConsoleConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) SetTipContent(v string) *GetPropertiesResponseBodyDataConsoleConfig {
	s.TipContent = &v
	return s
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) SetTitle(v string) *GetPropertiesResponseBodyDataConsoleConfig {
	s.Title = &v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfig struct {
	// example:
	//
	// xxx
	ProductDescription *string                                                              `json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	SearchSamples      []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples `json:"SearchSamples,omitempty" xml:"SearchSamples,omitempty" type:"Repeated"`
	SearchSources      []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetProductDescription(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.ProductDescription = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetSearchSamples(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.SearchSamples = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetSearchSources(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.SearchSources = v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples struct {
	Articles []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetArticles(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Articles = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetPrompt(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Prompt = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetText(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Text = &v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles struct {
	// example:
	//
	// true
	Select *bool `json:"Select,omitempty" xml:"Select,omitempty"`
	// example:
	//
	// false
	Stared *bool `json:"Stared,omitempty" xml:"Stared,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxx.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetSelect(v bool) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Select = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetStared(v bool) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Stared = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetTitle(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Title = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetUrl(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Url = &v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources struct {
	// example:
	//
	// xx
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xx
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetCode(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetDatasetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.DatasetName = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.Name = &v
	return s
}

type GetPropertiesResponseBodyDataSearchSources struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// SystemSearch
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataSearchSources) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataSearchSources) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataSearchSources) SetLabel(v string) *GetPropertiesResponseBodyDataSearchSources {
	s.Label = &v
	return s
}

func (s *GetPropertiesResponseBodyDataSearchSources) SetValue(v string) *GetPropertiesResponseBodyDataSearchSources {
	s.Value = &v
	return s
}

type GetPropertiesResponseBodyDataUserInfo struct {
	// example:
	//
	// 1
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetPropertiesResponseBodyDataUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataUserInfo) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetAgentId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.AgentId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetTenantId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.TenantId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetUserId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.UserId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetUsername(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.Username = &v
	return s
}

type GetPropertiesResponseBodyDataWanxiangImageSizeConfig struct {
	// example:
	//
	// 1:1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1024*1024
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataWanxiangImageSizeConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataWanxiangImageSizeConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) SetName(v string) *GetPropertiesResponseBodyDataWanxiangImageSizeConfig {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) SetValue(v string) *GetPropertiesResponseBodyDataWanxiangImageSizeConfig {
	s.Value = &v
	return s
}

type GetPropertiesResponseBodyDataWanxiangImageStyleConfig struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01RzKicz1W0YWzYkWcK_!!6000000002726-2-tps-132-104.png
	Pic *string `json:"Pic,omitempty" xml:"Pic,omitempty"`
	// example:
	//
	// <auto>
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataWanxiangImageStyleConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataWanxiangImageStyleConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetName(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetPic(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Pic = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetValue(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Value = &v
	return s
}

type GetPropertiesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponse) SetHeaders(v map[string]*string) *GetPropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetPropertiesResponse) SetStatusCode(v int32) *GetPropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPropertiesResponse) SetBody(v *GetPropertiesResponseBody) *GetPropertiesResponse {
	s.Body = v
	return s
}

type GetTopicByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 数据ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTopicByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicByIdRequest) GoString() string {
	return s.String()
}

func (s *GetTopicByIdRequest) SetAgentKey(v string) *GetTopicByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *GetTopicByIdRequest) SetId(v string) *GetTopicByIdRequest {
	s.Id = &v
	return s
}

type GetTopicByIdResponseBody struct {
	// example:
	//
	// NoData
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTopicByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTopicByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBody) SetCode(v string) *GetTopicByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetData(v *GetTopicByIdResponseBodyData) *GetTopicByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicByIdResponseBody) SetHttpStatusCode(v int32) *GetTopicByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetMessage(v string) *GetTopicByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetRequestId(v string) *GetTopicByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicByIdResponseBody) SetSuccess(v bool) *GetTopicByIdResponseBody {
	s.Success = &v
	return s
}

type GetTopicByIdResponseBodyData struct {
	// example:
	//
	// 异步任务ID（自定义主题场景下使用）
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 创建用户ID（自定义主题场景下使用）
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 43
	HotValue *int64 `json:"HotValue,omitempty" xml:"HotValue,omitempty"`
	// example:
	//
	// 热榜ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// PENDING
	Status           *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureSummary []*GetTopicByIdResponseBodyDataStructureSummary `json:"StructureSummary,omitempty" xml:"StructureSummary,omitempty" type:"Repeated"`
	// example:
	//
	// 热榜摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 异步任务失败错误信息
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 14
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 主题唯一名称
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 热榜源，目前支持的热榜源: Toutiao：头条、Quark：夸克、Baidu：百度、Sina：新浪。Custom：自定义、Aggregation：热点话题榜
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 数据版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetTopicByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTopicByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBodyData) SetAsyncTaskId(v string) *GetTopicByIdResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetCreateUser(v string) *GetTopicByIdResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetHotValue(v int64) *GetTopicByIdResponseBodyData {
	s.HotValue = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetId(v string) *GetTopicByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetStatus(v string) *GetTopicByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetStructureSummary(v []*GetTopicByIdResponseBodyDataStructureSummary) *GetTopicByIdResponseBodyData {
	s.StructureSummary = v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetSummary(v string) *GetTopicByIdResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTaskErrorMessage(v string) *GetTopicByIdResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTaskStatus(v int32) *GetTopicByIdResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTopic(v string) *GetTopicByIdResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetTopicSource(v string) *GetTopicByIdResponseBodyData {
	s.TopicSource = &v
	return s
}

func (s *GetTopicByIdResponseBodyData) SetVersion(v string) *GetTopicByIdResponseBodyData {
	s.Version = &v
	return s
}

type GetTopicByIdResponseBodyDataStructureSummary struct {
	DocList []*GetTopicByIdResponseBodyDataStructureSummaryDocList `json:"DocList,omitempty" xml:"DocList,omitempty" type:"Repeated"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTopicByIdResponseBodyDataStructureSummary) String() string {
	return tea.Prettify(s)
}

func (s GetTopicByIdResponseBodyDataStructureSummary) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) SetDocList(v []*GetTopicByIdResponseBodyDataStructureSummaryDocList) *GetTopicByIdResponseBodyDataStructureSummary {
	s.DocList = v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) SetSummary(v string) *GetTopicByIdResponseBodyDataStructureSummary {
	s.Summary = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummary) SetTitle(v string) *GetTopicByIdResponseBodyDataStructureSummary {
	s.Title = &v
	return s
}

type GetTopicByIdResponseBodyDataStructureSummaryDocList struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTopicByIdResponseBodyDataStructureSummaryDocList) String() string {
	return tea.Prettify(s)
}

func (s GetTopicByIdResponseBodyDataStructureSummaryDocList) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) SetSource(v string) *GetTopicByIdResponseBodyDataStructureSummaryDocList {
	s.Source = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) SetTitle(v string) *GetTopicByIdResponseBodyDataStructureSummaryDocList {
	s.Title = &v
	return s
}

func (s *GetTopicByIdResponseBodyDataStructureSummaryDocList) SetUrl(v string) *GetTopicByIdResponseBodyDataStructureSummaryDocList {
	s.Url = &v
	return s
}

type GetTopicByIdResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicByIdResponse) GoString() string {
	return s.String()
}

func (s *GetTopicByIdResponse) SetHeaders(v map[string]*string) *GetTopicByIdResponse {
	s.Headers = v
	return s
}

func (s *GetTopicByIdResponse) SetStatusCode(v int32) *GetTopicByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicByIdResponse) SetBody(v *GetTopicByIdResponseBody) *GetTopicByIdResponse {
	s.Body = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// c9f226b02cca4f42a84c5e955c39dfd2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *GetTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskRequest) SetTaskId(v string) *GetTopicSelectionPerspectiveAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *GetTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 错误信息
	ErrorMessage          *string                                                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FreshViewPointsResult *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult `json:"FreshViewPointsResult,omitempty" xml:"FreshViewPointsResult,omitempty" type:"Struct"`
	HotViewPointsResult   *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult   `json:"HotViewPointsResult,omitempty" xml:"HotViewPointsResult,omitempty" type:"Struct"`
	// example:
	//
	// SUSPENDED
	Status                *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	TimedViewPointsResult *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult `json:"TimedViewPointsResult,omitempty" xml:"TimedViewPointsResult,omitempty" type:"Struct"`
	// example:
	//
	// 热点主题事件
	Topic                 *string                                                                        `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TopicSummaryResult    *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult    `json:"TopicSummaryResult,omitempty" xml:"TopicSummaryResult,omitempty" type:"Struct"`
	WebReviewPointsResult *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult `json:"WebReviewPointsResult,omitempty" xml:"WebReviewPointsResult,omitempty" type:"Struct"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetFreshViewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.FreshViewPointsResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetHotViewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.HotViewPointsResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetStatus(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTimedViewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TimedViewPointsResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTopic(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.Topic = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTopicSummaryResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TopicSummaryResult = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetWebReviewPointsResult(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.WebReviewPointsResult = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResult {
	s.Attitudes = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                                                                             `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataFreshViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResult {
	s.Attitudes = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                                                                     `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	News         []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                                                                           `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetNews(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.News = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews struct {
	// example:
	//
	// 9957175DEDCF49C5ACF7A956B4FD67B2
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 123456
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.png
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// ["标签1","标签2"]
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 文章主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetDocId(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.DocId = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetDocUuid(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.DocUuid = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetImageUrls(v []*string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.ImageUrls = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetTags(v []*string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Tags = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews) SetTopic(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesNews {
	s.Topic = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataHotViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResult {
	s.Attitudes = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	// example:
	//
	// 2024-01-22 10:29
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 当前观点占比
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com/news/1.html
	Url        *string                                                                                             `json:"Url,omitempty" xml:"Url,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetPubTime(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.PubTime = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetSource(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Source = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetUrl(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.Url = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTimedViewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult struct {
	Summaries []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries `json:"Summaries,omitempty" xml:"Summaries,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult) SetSummaries(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResult {
	s.Summaries = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries struct {
	DocList []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList `json:"DocList,omitempty" xml:"DocList,omitempty" type:"Repeated"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) SetDocList(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries {
	s.DocList = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries {
	s.Summary = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummaries {
	s.Title = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) SetSource(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList {
	s.Source = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList) SetUrl(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataTopicSummaryResultSummariesDocList {
	s.Url = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult struct {
	Attitudes []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult) SetAttitudes(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResult {
	s.Attitudes = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                                                                           `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	Comments     []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                                                                             `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetAttitude(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.Attitude = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetAttitudeType(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.AttitudeType = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetComments(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.Comments = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetRatio(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.Ratio = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes) SetViewPoints(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudes {
	s.ViewPoints = v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments struct {
	// example:
	//
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 评论内容
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 当前评论所属的URL
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 评论用户名
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetSource(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Source = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetText(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Text = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetTitle(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Title = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetUrl(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Url = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments) SetUsername(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesComments {
	s.Username = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints struct {
	Outlines []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) SetOutlines(v []*GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints {
	s.Outlines = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) SetPoint(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints {
	s.Point = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPoints {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) SetOutline(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines) SetSummary(v string) *GetTopicSelectionPerspectiveAnalysisTaskResponseBodyDataWebReviewPointsResultAttitudesViewPointsOutlines {
	s.Summary = &v
	return s
}

type GetTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *GetTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *GetTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *GetTopicSelectionPerspectiveAnalysisTaskResponseBody) *GetTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

type ImportInterveneFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// import.xls
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// import.xsl
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// http://xxx/xxx.xls
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ImportInterveneFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileRequest) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileRequest) SetAgentKey(v string) *ImportInterveneFileRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportInterveneFileRequest) SetDocName(v string) *ImportInterveneFileRequest {
	s.DocName = &v
	return s
}

func (s *ImportInterveneFileRequest) SetFileKey(v string) *ImportInterveneFileRequest {
	s.FileKey = &v
	return s
}

func (s *ImportInterveneFileRequest) SetFileUrl(v string) *ImportInterveneFileRequest {
	s.FileUrl = &v
	return s
}

type ImportInterveneFileResponseBody struct {
	// example:
	//
	// 0
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ImportInterveneFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportInterveneFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileResponseBody) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponseBody) SetCode(v string) *ImportInterveneFileResponseBody {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetData(v *ImportInterveneFileResponseBodyData) *ImportInterveneFileResponseBody {
	s.Data = v
	return s
}

func (s *ImportInterveneFileResponseBody) SetHttpStatusCode(v int32) *ImportInterveneFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetMessage(v string) *ImportInterveneFileResponseBody {
	s.Message = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetRequestId(v string) *ImportInterveneFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetSuccess(v bool) *ImportInterveneFileResponseBody {
	s.Success = &v
	return s
}

type ImportInterveneFileResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportInterveneFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponseBodyData) SetFailIdList(v []*string) *ImportInterveneFileResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ImportInterveneFileResponseBodyData) SetTaskId(v string) *ImportInterveneFileResponseBodyData {
	s.TaskId = &v
	return s
}

type ImportInterveneFileResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportInterveneFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportInterveneFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileResponse) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponse) SetHeaders(v map[string]*string) *ImportInterveneFileResponse {
	s.Headers = v
	return s
}

func (s *ImportInterveneFileResponse) SetStatusCode(v int32) *ImportInterveneFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportInterveneFileResponse) SetBody(v *ImportInterveneFileResponseBody) *ImportInterveneFileResponse {
	s.Body = v
	return s
}

type ImportInterveneFileAsyncRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// import.xls
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// import.xls
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// https://xxx/import.xls
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ImportInterveneFileAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncRequest) SetAgentKey(v string) *ImportInterveneFileAsyncRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetDocName(v string) *ImportInterveneFileAsyncRequest {
	s.DocName = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetFileKey(v string) *ImportInterveneFileAsyncRequest {
	s.FileKey = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetFileUrl(v string) *ImportInterveneFileAsyncRequest {
	s.FileUrl = &v
	return s
}

type ImportInterveneFileAsyncResponseBody struct {
	// example:
	//
	// 0
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ImportInterveneFileAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportInterveneFileAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponseBody) SetCode(v string) *ImportInterveneFileAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetData(v *ImportInterveneFileAsyncResponseBodyData) *ImportInterveneFileAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetHttpStatusCode(v int32) *ImportInterveneFileAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetMessage(v string) *ImportInterveneFileAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetRequestId(v string) *ImportInterveneFileAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetSuccess(v bool) *ImportInterveneFileAsyncResponseBody {
	s.Success = &v
	return s
}

type ImportInterveneFileAsyncResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportInterveneFileAsyncResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponseBodyData) SetFailIdList(v []*string) *ImportInterveneFileAsyncResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ImportInterveneFileAsyncResponseBodyData) SetTaskId(v string) *ImportInterveneFileAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

type ImportInterveneFileAsyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportInterveneFileAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportInterveneFileAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponse) SetHeaders(v map[string]*string) *ImportInterveneFileAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImportInterveneFileAsyncResponse) SetStatusCode(v int32) *ImportInterveneFileAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportInterveneFileAsyncResponse) SetBody(v *ImportInterveneFileAsyncResponseBody) *ImportInterveneFileAsyncResponse {
	s.Body = v
	return s
}

type InsertInterveneGlobalReplyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey        *string                                             `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReplyMessagList []*InsertInterveneGlobalReplyRequestReplyMessagList `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty" type:"Repeated"`
}

func (s InsertInterveneGlobalReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyRequest) SetAgentKey(v string) *InsertInterveneGlobalReplyRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneGlobalReplyRequest) SetReplyMessagList(v []*InsertInterveneGlobalReplyRequestReplyMessagList) *InsertInterveneGlobalReplyRequest {
	s.ReplyMessagList = v
	return s
}

type InsertInterveneGlobalReplyRequestReplyMessagList struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReplyType *string `json:"ReplyType,omitempty" xml:"ReplyType,omitempty"`
}

func (s InsertInterveneGlobalReplyRequestReplyMessagList) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyRequestReplyMessagList) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) SetMessage(v string) *InsertInterveneGlobalReplyRequestReplyMessagList {
	s.Message = &v
	return s
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) SetReplyType(v string) *InsertInterveneGlobalReplyRequestReplyMessagList {
	s.ReplyType = &v
	return s
}

type InsertInterveneGlobalReplyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey              *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReplyMessagListShrink *string `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty"`
}

func (s InsertInterveneGlobalReplyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyShrinkRequest) SetAgentKey(v string) *InsertInterveneGlobalReplyShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneGlobalReplyShrinkRequest) SetReplyMessagListShrink(v string) *InsertInterveneGlobalReplyShrinkRequest {
	s.ReplyMessagListShrink = &v
	return s
}

type InsertInterveneGlobalReplyResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InsertInterveneGlobalReplyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertInterveneGlobalReplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponseBody) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponseBody) SetCode(v string) *InsertInterveneGlobalReplyResponseBody {
	s.Code = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetData(v *InsertInterveneGlobalReplyResponseBodyData) *InsertInterveneGlobalReplyResponseBody {
	s.Data = v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetHttpStatusCode(v int32) *InsertInterveneGlobalReplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetMessage(v string) *InsertInterveneGlobalReplyResponseBody {
	s.Message = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetRequestId(v string) *InsertInterveneGlobalReplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetSuccess(v bool) *InsertInterveneGlobalReplyResponseBody {
	s.Success = &v
	return s
}

type InsertInterveneGlobalReplyResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 4829
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InsertInterveneGlobalReplyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponseBodyData) SetFailIdList(v []*string) *InsertInterveneGlobalReplyResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBodyData) SetTaskId(v string) *InsertInterveneGlobalReplyResponseBodyData {
	s.TaskId = &v
	return s
}

type InsertInterveneGlobalReplyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertInterveneGlobalReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertInterveneGlobalReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponse) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponse) SetHeaders(v map[string]*string) *InsertInterveneGlobalReplyResponse {
	s.Headers = v
	return s
}

func (s *InsertInterveneGlobalReplyResponse) SetStatusCode(v int32) *InsertInterveneGlobalReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponse) SetBody(v *InsertInterveneGlobalReplyResponseBody) *InsertInterveneGlobalReplyResponse {
	s.Body = v
	return s
}

type InsertInterveneRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey            *string                                        `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneRuleConfig *InsertInterveneRuleRequestInterveneRuleConfig `json:"InterveneRuleConfig,omitempty" xml:"InterveneRuleConfig,omitempty" type:"Struct"`
}

func (s InsertInterveneRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequest) SetAgentKey(v string) *InsertInterveneRuleRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneRuleRequest) SetInterveneRuleConfig(v *InsertInterveneRuleRequestInterveneRuleConfig) *InsertInterveneRuleRequest {
	s.InterveneRuleConfig = v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfig struct {
	AnswerConfig        []*InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig        `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	EffectConfig        *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig          `json:"EffectConfig,omitempty" xml:"EffectConfig,omitempty" type:"Struct"`
	InterveneConfigList []*InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList `json:"InterveneConfigList,omitempty" xml:"InterveneConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	InterveneType *int32    `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// tf-test-rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetAnswerConfig(v []*InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.AnswerConfig = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetEffectConfig(v *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.EffectConfig = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetInterveneConfigList(v []*InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.InterveneConfigList = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetInterveneType(v int32) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.InterveneType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetNamespaceList(v []*string) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.NamespaceList = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetRuleId(v int64) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.RuleId = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetRuleName(v string) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.RuleName = &v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig struct {
	// example:
	//
	// 0
	AnswerType *int32  `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// namespace_qa_query
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetAnswerType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetMessage(v string) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.Message = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetNamespace(v string) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.Namespace = &v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfigEffectConfig struct {
	// example:
	//
	// 0
	EffectType *int32 `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	// example:
	//
	// 2023-03-28 06:04:29
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2023-03-28 06:04:29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetEffectType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.EffectType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetEndTime(v string) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.EndTime = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetStartTime(v string) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.StartTime = &v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList struct {
	// id
	//
	// example:
	//
	// 37249
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	OperationType *int32  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Query         *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetId(v string) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.Id = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetOperationType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.OperationType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetQuery(v string) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.Query = &v
	return s
}

type InsertInterveneRuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey                  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneRuleConfigShrink *string `json:"InterveneRuleConfig,omitempty" xml:"InterveneRuleConfig,omitempty"`
}

func (s InsertInterveneRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleShrinkRequest) SetAgentKey(v string) *InsertInterveneRuleShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneRuleShrinkRequest) SetInterveneRuleConfigShrink(v string) *InsertInterveneRuleShrinkRequest {
	s.InterveneRuleConfigShrink = &v
	return s
}

type InsertInterveneRuleResponseBody struct {
	// example:
	//
	// 0
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *InsertInterveneRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DD656AF9-0839-521A-A3D2-F320009F9C87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertInterveneRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponseBody) SetCode(v string) *InsertInterveneRuleResponseBody {
	s.Code = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetData(v *InsertInterveneRuleResponseBodyData) *InsertInterveneRuleResponseBody {
	s.Data = v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetHttpStatusCode(v int32) *InsertInterveneRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetMessage(v string) *InsertInterveneRuleResponseBody {
	s.Message = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetRequestId(v string) *InsertInterveneRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetSuccess(v bool) *InsertInterveneRuleResponseBody {
	s.Success = &v
	return s
}

type InsertInterveneRuleResponseBodyData struct {
	// example:
	//
	// 12345
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s InsertInterveneRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponseBodyData) SetRuleId(v int64) *InsertInterveneRuleResponseBodyData {
	s.RuleId = &v
	return s
}

type InsertInterveneRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertInterveneRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertInterveneRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleResponse) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponse) SetHeaders(v map[string]*string) *InsertInterveneRuleResponse {
	s.Headers = v
	return s
}

func (s *InsertInterveneRuleResponse) SetStatusCode(v int32) *InsertInterveneRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertInterveneRuleResponse) SetBody(v *InsertInterveneRuleResponseBody) *InsertInterveneRuleResponse {
	s.Body = v
	return s
}

type ListAsyncTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 1
	TaskStatus     *int32    `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusList []*int32  `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty" type:"Repeated"`
	TaskType       *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeList   []*string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty" type:"Repeated"`
}

func (s ListAsyncTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksRequest) SetAgentKey(v string) *ListAsyncTasksRequest {
	s.AgentKey = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCreateTimeEnd(v string) *ListAsyncTasksRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCreateTimeStart(v string) *ListAsyncTasksRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCurrent(v int32) *ListAsyncTasksRequest {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksRequest) SetSize(v int32) *ListAsyncTasksRequest {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskCode(v string) *ListAsyncTasksRequest {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskName(v string) *ListAsyncTasksRequest {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskStatus(v int32) *ListAsyncTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskStatusList(v []*int32) *ListAsyncTasksRequest {
	s.TaskStatusList = v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskType(v string) *ListAsyncTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskTypeList(v []*string) *ListAsyncTasksRequest {
	s.TaskTypeList = v
	return s
}

type ListAsyncTasksShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 1
	TaskStatus           *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusListShrink *string `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeListShrink   *string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty"`
}

func (s ListAsyncTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksShrinkRequest) SetAgentKey(v string) *ListAsyncTasksShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCreateTimeEnd(v string) *ListAsyncTasksShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCreateTimeStart(v string) *ListAsyncTasksShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCurrent(v int32) *ListAsyncTasksShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetSize(v int32) *ListAsyncTasksShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskCode(v string) *ListAsyncTasksShrinkRequest {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskName(v string) *ListAsyncTasksShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskStatus(v int32) *ListAsyncTasksShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskStatusListShrink(v string) *ListAsyncTasksShrinkRequest {
	s.TaskStatusListShrink = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskType(v string) *ListAsyncTasksShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskTypeListShrink(v string) *ListAsyncTasksShrinkRequest {
	s.TaskTypeListShrink = &v
	return s
}

type ListAsyncTasksResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                            `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListAsyncTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAsyncTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBody) SetCode(v string) *ListAsyncTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetCurrent(v int32) *ListAsyncTasksResponseBody {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetData(v []*ListAsyncTasksResponseBodyData) *ListAsyncTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListAsyncTasksResponseBody) SetHttpStatusCode(v int32) *ListAsyncTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetMessage(v string) *ListAsyncTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetRequestId(v string) *ListAsyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetSize(v int32) *ListAsyncTasksResponseBody {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetSuccess(v bool) *ListAsyncTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetTotal(v int32) *ListAsyncTasksResponseBody {
	s.Total = &v
	return s
}

type ListAsyncTasksResponseBodyData struct {
	// example:
	//
	// 2020-12-23 15:41:58
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1111
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// example:
	//
	// {}
	TaskDefinition *string `json:"TaskDefinition,omitempty" xml:"TaskDefinition,omitempty"`
	// example:
	//
	// 2023-03-09 00:00:00
	TaskEndTime      *string `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 2023-10-14 14:30:00
	TaskExecuteTime *string `json:"TaskExecuteTime,omitempty" xml:"TaskExecuteTime,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId                *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskInnerErrorMessage *string `json:"TaskInnerErrorMessage,omitempty" xml:"TaskInnerErrorMessage,omitempty"`
	// example:
	//
	// {}
	TaskIntermediateResult *string `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	TaskName               *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// {}
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// example:
	//
	// {}
	TaskProgressMessage *string `json:"TaskProgressMessage,omitempty" xml:"TaskProgressMessage,omitempty"`
	// example:
	//
	// {}
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// example:
	//
	// 1
	TaskRetryCount *string `json:"TaskRetryCount,omitempty" xml:"TaskRetryCount,omitempty"`
	// example:
	//
	// 2023-03-20 10:53:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 2023-02-16 10:29:16
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 111
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListAsyncTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBodyData) SetCreateTime(v string) *ListAsyncTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetCreateUser(v string) *ListAsyncTasksResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetId(v int64) *ListAsyncTasksResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskCode(v string) *ListAsyncTasksResponseBodyData {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskDefinition(v string) *ListAsyncTasksResponseBodyData {
	s.TaskDefinition = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskEndTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskEndTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskErrorMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskExecuteTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskExecuteTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskId(v string) *ListAsyncTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskInnerErrorMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskInnerErrorMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskIntermediateResult(v string) *ListAsyncTasksResponseBodyData {
	s.TaskIntermediateResult = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskName(v string) *ListAsyncTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskParam(v string) *ListAsyncTasksResponseBodyData {
	s.TaskParam = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskProgressMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskProgressMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskResult(v string) *ListAsyncTasksResponseBodyData {
	s.TaskResult = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskRetryCount(v string) *ListAsyncTasksResponseBodyData {
	s.TaskRetryCount = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskStartTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskStartTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskStatus(v int32) *ListAsyncTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskType(v string) *ListAsyncTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetUpdateTime(v string) *ListAsyncTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetUpdateUser(v string) *ListAsyncTasksResponseBodyData {
	s.UpdateUser = &v
	return s
}

type ListAsyncTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsyncTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponse) SetHeaders(v map[string]*string) *ListAsyncTasksResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncTasksResponse) SetStatusCode(v int32) *ListAsyncTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncTasksResponse) SetBody(v *ListAsyncTasksResponseBody) *ListAsyncTasksResponse {
	s.Body = v
	return s
}

type ListBuildConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBuildConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsRequest) SetAgentKey(v string) *ListBuildConfigsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBuildConfigsRequest) SetRegionId(v string) *ListBuildConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBuildConfigsRequest) SetType(v string) *ListBuildConfigsRequest {
	s.Type = &v
	return s
}

type ListBuildConfigsResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListBuildConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBuildConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBody) SetCode(v string) *ListBuildConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetData(v []*ListBuildConfigsResponseBodyData) *ListBuildConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListBuildConfigsResponseBody) SetHttpStatusCode(v int32) *ListBuildConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetMessage(v string) *ListBuildConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetRequestId(v string) *ListBuildConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetSuccess(v bool) *ListBuildConfigsResponseBody {
	s.Success = &v
	return s
}

type ListBuildConfigsResponseBodyData struct {
	// example:
	//
	// true
	BuildIn *bool `json:"BuildIn,omitempty" xml:"BuildIn,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *string                                     `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Id         *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords   []*ListBuildConfigsResponseBodyDataKeywords `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// example:
	//
	// writingStyle
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// example:
	//
	// media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListBuildConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBodyData) SetBuildIn(v bool) *ListBuildConfigsResponseBodyData {
	s.BuildIn = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetCreateTime(v string) *ListBuildConfigsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetCreateUser(v string) *ListBuildConfigsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetId(v int64) *ListBuildConfigsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetKeywords(v []*ListBuildConfigsResponseBodyDataKeywords) *ListBuildConfigsResponseBodyData {
	s.Keywords = v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetTag(v string) *ListBuildConfigsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetTagDescription(v string) *ListBuildConfigsResponseBodyData {
	s.TagDescription = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetType(v string) *ListBuildConfigsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetUpdateTime(v string) *ListBuildConfigsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetUpdateUser(v string) *ListBuildConfigsResponseBodyData {
	s.UpdateUser = &v
	return s
}

type ListBuildConfigsResponseBodyDataKeywords struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListBuildConfigsResponseBodyDataKeywords) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponseBodyDataKeywords) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBodyDataKeywords) SetDescription(v string) *ListBuildConfigsResponseBodyDataKeywords {
	s.Description = &v
	return s
}

func (s *ListBuildConfigsResponseBodyDataKeywords) SetKey(v string) *ListBuildConfigsResponseBodyDataKeywords {
	s.Key = &v
	return s
}

type ListBuildConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBuildConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBuildConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponse) SetHeaders(v map[string]*string) *ListBuildConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListBuildConfigsResponse) SetStatusCode(v int32) *ListBuildConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBuildConfigsResponse) SetBody(v *ListBuildConfigsResponseBody) *ListBuildConfigsResponse {
	s.Body = v
	return s
}

type ListCustomTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 商品code
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s ListCustomTextRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomTextRequest) GoString() string {
	return s.String()
}

func (s *ListCustomTextRequest) SetAgentKey(v string) *ListCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCustomTextRequest) SetCommodityCode(v string) *ListCustomTextRequest {
	s.CommodityCode = &v
	return s
}

type ListCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListCustomTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomTextResponseBody) SetCode(v string) *ListCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomTextResponseBody) SetData(v []*ListCustomTextResponseBodyData) *ListCustomTextResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomTextResponseBody) SetHttpStatusCode(v int32) *ListCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCustomTextResponseBody) SetMessage(v string) *ListCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomTextResponseBody) SetRequestId(v string) *ListCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomTextResponseBody) SetSuccess(v bool) *ListCustomTextResponseBody {
	s.Success = &v
	return s
}

type ListCustomTextResponseBodyData struct {
	// example:
	//
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 创建用户
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 40
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 修改用户
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListCustomTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCustomTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomTextResponseBodyData) SetContent(v string) *ListCustomTextResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetCreateTime(v string) *ListCustomTextResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetCreateUser(v string) *ListCustomTextResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetId(v int64) *ListCustomTextResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetTitle(v string) *ListCustomTextResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetUpdateTime(v string) *ListCustomTextResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListCustomTextResponseBodyData) SetUpdateUser(v string) *ListCustomTextResponseBodyData {
	s.UpdateUser = &v
	return s
}

type ListCustomTextResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomTextResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomTextResponse) GoString() string {
	return s.String()
}

func (s *ListCustomTextResponse) SetHeaders(v map[string]*string) *ListCustomTextResponse {
	s.Headers = v
	return s
}

func (s *ListCustomTextResponse) SetStatusCode(v int32) *ListCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomTextResponse) SetBody(v *ListCustomTextResponseBody) *ListCustomTextResponse {
	s.Body = v
	return s
}

type ListCustomViewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点
	//
	//      *
	Attitudes []*string `json:"Attitudes,omitempty" xml:"Attitudes,omitempty" type:"Repeated"`
	// example:
	//
	// 461591f4880747f890702c1b90494d1a
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 7ece3d1212e04c9ca716ae2486228f3f
	CustomViewPointIds []*string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty" type:"Repeated"`
	// example:
	//
	// 52
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 1d20ed14db0840efb1c7eaaf4d46352b
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ListCustomViewPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomViewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsRequest) SetAgentKey(v string) *ListCustomViewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetAttitude(v string) *ListCustomViewPointsRequest {
	s.Attitude = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetAttitudes(v []*string) *ListCustomViewPointsRequest {
	s.Attitudes = v
	return s
}

func (s *ListCustomViewPointsRequest) SetCustomViewPointId(v string) *ListCustomViewPointsRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetCustomViewPointIds(v []*string) *ListCustomViewPointsRequest {
	s.CustomViewPointIds = v
	return s
}

func (s *ListCustomViewPointsRequest) SetMaxResults(v int32) *ListCustomViewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetNextToken(v string) *ListCustomViewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetTopic(v string) *ListCustomViewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListCustomViewPointsRequest) SetTopicId(v string) *ListCustomViewPointsRequest {
	s.TopicId = &v
	return s
}

type ListCustomViewPointsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点
	//
	//      *
	AttitudesShrink *string `json:"Attitudes,omitempty" xml:"Attitudes,omitempty"`
	// example:
	//
	// 461591f4880747f890702c1b90494d1a
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 7ece3d1212e04c9ca716ae2486228f3f
	CustomViewPointIdsShrink *string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty"`
	// example:
	//
	// 52
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 1d20ed14db0840efb1c7eaaf4d46352b
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ListCustomViewPointsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomViewPointsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsShrinkRequest) SetAgentKey(v string) *ListCustomViewPointsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetAttitude(v string) *ListCustomViewPointsShrinkRequest {
	s.Attitude = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetAttitudesShrink(v string) *ListCustomViewPointsShrinkRequest {
	s.AttitudesShrink = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetCustomViewPointId(v string) *ListCustomViewPointsShrinkRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetCustomViewPointIdsShrink(v string) *ListCustomViewPointsShrinkRequest {
	s.CustomViewPointIdsShrink = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetMaxResults(v int32) *ListCustomViewPointsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetNextToken(v string) *ListCustomViewPointsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetTopic(v string) *ListCustomViewPointsShrinkRequest {
	s.Topic = &v
	return s
}

func (s *ListCustomViewPointsShrinkRequest) SetTopicId(v string) *ListCustomViewPointsShrinkRequest {
	s.TopicId = &v
	return s
}

type ListCustomViewPointsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListCustomViewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 60
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 73
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomViewPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomViewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBody) SetCode(v string) *ListCustomViewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetData(v []*ListCustomViewPointsResponseBodyData) *ListCustomViewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetHttpStatusCode(v int32) *ListCustomViewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetMaxResults(v int32) *ListCustomViewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetMessage(v string) *ListCustomViewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetNextToken(v string) *ListCustomViewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetRequestId(v string) *ListCustomViewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetSuccess(v bool) *ListCustomViewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomViewPointsResponseBody) SetTotalCount(v int32) *ListCustomViewPointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomViewPointsResponseBodyData struct {
	// example:
	//
	// 2323ac73e174428a98c91097a59c67e0
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 2024-08-15 16:18:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 709806dd051042d5ab9de8bdbb3a64ca
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 参数校验失败
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32                                            `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	ViewPoints []*ListCustomViewPointsResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListCustomViewPointsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCustomViewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBodyData) SetAsyncTaskId(v string) *ListCustomViewPointsResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetAttitude(v string) *ListCustomViewPointsResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetCreateTime(v string) *ListCustomViewPointsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetCreateUser(v string) *ListCustomViewPointsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetId(v string) *ListCustomViewPointsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetStatus(v string) *ListCustomViewPointsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetTaskErrorMessage(v string) *ListCustomViewPointsResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetTaskStatus(v int32) *ListCustomViewPointsResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyData) SetViewPoints(v []*ListCustomViewPointsResponseBodyDataViewPoints) *ListCustomViewPointsResponseBodyData {
	s.ViewPoints = v
	return s
}

type ListCustomViewPointsResponseBodyDataViewPoints struct {
	Outlines []*ListCustomViewPointsResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListCustomViewPointsResponseBodyDataViewPoints) String() string {
	return tea.Prettify(s)
}

func (s ListCustomViewPointsResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) SetOutlines(v []*ListCustomViewPointsResponseBodyDataViewPointsOutlines) *ListCustomViewPointsResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) SetPoint(v string) *ListCustomViewPointsResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPoints) SetSummary(v string) *ListCustomViewPointsResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

type ListCustomViewPointsResponseBodyDataViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListCustomViewPointsResponseBodyDataViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s ListCustomViewPointsResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListCustomViewPointsResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListCustomViewPointsResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListCustomViewPointsResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

type ListCustomViewPointsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomViewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomViewPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomViewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomViewPointsResponse) SetHeaders(v map[string]*string) *ListCustomViewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomViewPointsResponse) SetStatusCode(v int32) *ListCustomViewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomViewPointsResponse) SetBody(v *ListCustomViewPointsResponseBody) *ListCustomViewPointsResponse {
	s.Body = v
	return s
}

type ListDialoguesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 2
	DialogueType *int32 `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListDialoguesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesRequest) GoString() string {
	return s.String()
}

func (s *ListDialoguesRequest) SetAgentKey(v string) *ListDialoguesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDialoguesRequest) SetCurrent(v int32) *ListDialoguesRequest {
	s.Current = &v
	return s
}

func (s *ListDialoguesRequest) SetDialogueType(v int32) *ListDialoguesRequest {
	s.DialogueType = &v
	return s
}

func (s *ListDialoguesRequest) SetEndTime(v string) *ListDialoguesRequest {
	s.EndTime = &v
	return s
}

func (s *ListDialoguesRequest) SetSize(v int32) *ListDialoguesRequest {
	s.Size = &v
	return s
}

func (s *ListDialoguesRequest) SetStartTime(v string) *ListDialoguesRequest {
	s.StartTime = &v
	return s
}

func (s *ListDialoguesRequest) SetTaskId(v string) *ListDialoguesRequest {
	s.TaskId = &v
	return s
}

type ListDialoguesResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                           `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListDialoguesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDialoguesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponseBody) SetCode(v string) *ListDialoguesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialoguesResponseBody) SetCurrent(v int32) *ListDialoguesResponseBody {
	s.Current = &v
	return s
}

func (s *ListDialoguesResponseBody) SetData(v []*ListDialoguesResponseBodyData) *ListDialoguesResponseBody {
	s.Data = v
	return s
}

func (s *ListDialoguesResponseBody) SetHttpStatusCode(v int32) *ListDialoguesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDialoguesResponseBody) SetMessage(v string) *ListDialoguesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialoguesResponseBody) SetRequestId(v string) *ListDialoguesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialoguesResponseBody) SetSize(v int32) *ListDialoguesResponseBody {
	s.Size = &v
	return s
}

func (s *ListDialoguesResponseBody) SetSuccess(v bool) *ListDialoguesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDialoguesResponseBody) SetTotal(v int32) *ListDialoguesResponseBody {
	s.Total = &v
	return s
}

type ListDialoguesResponseBodyData struct {
	// example:
	//
	// xx
	Bot *string `json:"Bot,omitempty" xml:"Bot,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 2
	DialogueType *int32 `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	// example:
	//
	// xx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// x
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListDialoguesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponseBodyData) SetBot(v string) *ListDialoguesResponseBodyData {
	s.Bot = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetCreateTime(v string) *ListDialoguesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetCreateUser(v string) *ListDialoguesResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetDialogueType(v int32) *ListDialoguesResponseBodyData {
	s.DialogueType = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetTaskId(v string) *ListDialoguesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetUser(v string) *ListDialoguesResponseBodyData {
	s.User = &v
	return s
}

type ListDialoguesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDialoguesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDialoguesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesResponse) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponse) SetHeaders(v map[string]*string) *ListDialoguesResponse {
	s.Headers = v
	return s
}

func (s *ListDialoguesResponse) SetStatusCode(v int32) *ListDialoguesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDialoguesResponse) SetBody(v *ListDialoguesResponseBody) *ListDialoguesResponse {
	s.Body = v
	return s
}

type ListFreshViewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 6
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
}

func (s ListFreshViewPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFreshViewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsRequest) SetAgentKey(v string) *ListFreshViewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetMaxResults(v int32) *ListFreshViewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetNextToken(v string) *ListFreshViewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetTopic(v string) *ListFreshViewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetTopicSource(v string) *ListFreshViewPointsRequest {
	s.TopicSource = &v
	return s
}

type ListFreshViewPointsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListFreshViewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 94
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 26
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFreshViewPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFreshViewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponseBody) SetCode(v string) *ListFreshViewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetData(v []*ListFreshViewPointsResponseBodyData) *ListFreshViewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetHttpStatusCode(v int32) *ListFreshViewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetMaxResults(v int32) *ListFreshViewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetMessage(v string) *ListFreshViewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetNextToken(v string) *ListFreshViewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetRequestId(v string) *ListFreshViewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetSuccess(v bool) *ListFreshViewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListFreshViewPointsResponseBody) SetTotalCount(v int32) *ListFreshViewPointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFreshViewPointsResponseBodyData struct {
	Outlines []*ListFreshViewPointsResponseBodyDataOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListFreshViewPointsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFreshViewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponseBodyData) SetOutlines(v []*ListFreshViewPointsResponseBodyDataOutlines) *ListFreshViewPointsResponseBodyData {
	s.Outlines = v
	return s
}

func (s *ListFreshViewPointsResponseBodyData) SetPoint(v string) *ListFreshViewPointsResponseBodyData {
	s.Point = &v
	return s
}

func (s *ListFreshViewPointsResponseBodyData) SetSummary(v string) *ListFreshViewPointsResponseBodyData {
	s.Summary = &v
	return s
}

type ListFreshViewPointsResponseBodyDataOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListFreshViewPointsResponseBodyDataOutlines) String() string {
	return tea.Prettify(s)
}

func (s ListFreshViewPointsResponseBodyDataOutlines) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponseBodyDataOutlines) SetOutline(v string) *ListFreshViewPointsResponseBodyDataOutlines {
	s.Outline = &v
	return s
}

func (s *ListFreshViewPointsResponseBodyDataOutlines) SetSummary(v string) *ListFreshViewPointsResponseBodyDataOutlines {
	s.Summary = &v
	return s
}

type ListFreshViewPointsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFreshViewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFreshViewPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFreshViewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsResponse) SetHeaders(v map[string]*string) *ListFreshViewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListFreshViewPointsResponse) SetStatusCode(v int32) *ListFreshViewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFreshViewPointsResponse) SetBody(v *ListFreshViewPointsResponseBody) *ListFreshViewPointsResponse {
	s.Body = v
	return s
}

type ListGeneratedContentsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// media
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Query   *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListGeneratedContentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsRequest) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsRequest) SetAgentKey(v string) *ListGeneratedContentsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetContentDomain(v string) *ListGeneratedContentsRequest {
	s.ContentDomain = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetCurrent(v int32) *ListGeneratedContentsRequest {
	s.Current = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetEndTime(v string) *ListGeneratedContentsRequest {
	s.EndTime = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetQuery(v string) *ListGeneratedContentsRequest {
	s.Query = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetSize(v int32) *ListGeneratedContentsRequest {
	s.Size = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetStartTime(v string) *ListGeneratedContentsRequest {
	s.StartTime = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetTaskId(v string) *ListGeneratedContentsRequest {
	s.TaskId = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetTitle(v string) *ListGeneratedContentsRequest {
	s.Title = &v
	return s
}

type ListGeneratedContentsResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                                   `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListGeneratedContentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGeneratedContentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponseBody) SetCode(v string) *ListGeneratedContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetCurrent(v int32) *ListGeneratedContentsResponseBody {
	s.Current = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetData(v []*ListGeneratedContentsResponseBodyData) *ListGeneratedContentsResponseBody {
	s.Data = v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetHttpStatusCode(v int32) *ListGeneratedContentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetMessage(v string) *ListGeneratedContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetRequestId(v string) *ListGeneratedContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetSize(v int32) *ListGeneratedContentsResponseBody {
	s.Size = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetSuccess(v bool) *ListGeneratedContentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetTotal(v int32) *ListGeneratedContentsResponseBody {
	s.Total = &v
	return s
}

type ListGeneratedContentsResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// media
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText   *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// xxx
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 10
	Id          *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordList []*string `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	Keywords    *string   `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Prompt      *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1111
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// example:
	//
	// xxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListGeneratedContentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponseBodyData) SetContent(v string) *ListGeneratedContentsResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetContentDomain(v string) *ListGeneratedContentsResponseBodyData {
	s.ContentDomain = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetContentText(v string) *ListGeneratedContentsResponseBodyData {
	s.ContentText = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetCreateTime(v string) *ListGeneratedContentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetCreateUser(v string) *ListGeneratedContentsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetDeviceId(v string) *ListGeneratedContentsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetId(v int64) *ListGeneratedContentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetKeywordList(v []*string) *ListGeneratedContentsResponseBodyData {
	s.KeywordList = v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetKeywords(v string) *ListGeneratedContentsResponseBodyData {
	s.Keywords = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetPrompt(v string) *ListGeneratedContentsResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetTaskId(v string) *ListGeneratedContentsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetTitle(v string) *ListGeneratedContentsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUpdateTime(v string) *ListGeneratedContentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUpdateUser(v string) *ListGeneratedContentsResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUuid(v string) *ListGeneratedContentsResponseBodyData {
	s.Uuid = &v
	return s
}

type ListGeneratedContentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGeneratedContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGeneratedContentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsResponse) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponse) SetHeaders(v map[string]*string) *ListGeneratedContentsResponse {
	s.Headers = v
	return s
}

func (s *ListGeneratedContentsResponse) SetStatusCode(v int32) *ListGeneratedContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGeneratedContentsResponse) SetBody(v *ListGeneratedContentsResponseBody) *ListGeneratedContentsResponse {
	s.Body = v
	return s
}

type ListHotNewsWithTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// society
	NewsType  *string   `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypes []*string `json:"NewsTypes,omitempty" xml:"NewsTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListHotNewsWithTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeRequest) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeRequest) SetAgentKey(v string) *ListHotNewsWithTypeRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetCurrent(v int32) *ListHotNewsWithTypeRequest {
	s.Current = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetNewsType(v string) *ListHotNewsWithTypeRequest {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetNewsTypes(v []*string) *ListHotNewsWithTypeRequest {
	s.NewsTypes = v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetSize(v int32) *ListHotNewsWithTypeRequest {
	s.Size = &v
	return s
}

type ListHotNewsWithTypeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// society
	NewsType        *string `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypesShrink *string `json:"NewsTypes,omitempty" xml:"NewsTypes,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListHotNewsWithTypeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeShrinkRequest) SetAgentKey(v string) *ListHotNewsWithTypeShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetCurrent(v int32) *ListHotNewsWithTypeShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetNewsType(v string) *ListHotNewsWithTypeShrinkRequest {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetNewsTypesShrink(v string) *ListHotNewsWithTypeShrinkRequest {
	s.NewsTypesShrink = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetSize(v int32) *ListHotNewsWithTypeShrinkRequest {
	s.Size = &v
	return s
}

type ListHotNewsWithTypeResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListHotNewsWithTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotNewsWithTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBody) SetCode(v string) *ListHotNewsWithTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetData(v []*ListHotNewsWithTypeResponseBodyData) *ListHotNewsWithTypeResponseBody {
	s.Data = v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetHttpStatusCode(v int32) *ListHotNewsWithTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetMessage(v string) *ListHotNewsWithTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetRequestId(v string) *ListHotNewsWithTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetSuccess(v bool) *ListHotNewsWithTypeResponseBody {
	s.Success = &v
	return s
}

type ListHotNewsWithTypeResponseBodyData struct {
	News []*ListHotNewsWithTypeResponseBodyDataNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// example:
	//
	// society
	NewsType     *string `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypeName *string `json:"NewsTypeName,omitempty" xml:"NewsTypeName,omitempty"`
	// example:
	//
	// 77
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListHotNewsWithTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNews(v []*ListHotNewsWithTypeResponseBodyDataNews) *ListHotNewsWithTypeResponseBodyData {
	s.News = v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNewsType(v string) *ListHotNewsWithTypeResponseBodyData {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNewsTypeName(v string) *ListHotNewsWithTypeResponseBodyData {
	s.NewsTypeName = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetTotalPages(v int32) *ListHotNewsWithTypeResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListHotNewsWithTypeResponseBodyDataNews struct {
	Author  *string `json:"Author,omitempty" xml:"Author,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	DocUuid   *string   `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime          *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SearchSource     *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	Source           *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary          *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tag              *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2023-10-14 14:30:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotNewsWithTypeResponseBodyDataNews) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetAuthor(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Author = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetContent(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetDocUuid(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.DocUuid = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetImageUrls(v []*string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.ImageUrls = v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetPubTime(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSearchSource(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.SearchSource = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSearchSourceName(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.SearchSourceName = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSource(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Source = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSummary(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Summary = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetTag(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Tag = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetTitle(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetUpdateTime(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.UpdateTime = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetUrl(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Url = &v
	return s
}

type ListHotNewsWithTypeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotNewsWithTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotNewsWithTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponse) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponse) SetHeaders(v map[string]*string) *ListHotNewsWithTypeResponse {
	s.Headers = v
	return s
}

func (s *ListHotNewsWithTypeResponse) SetStatusCode(v int32) *ListHotNewsWithTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotNewsWithTypeResponse) SetBody(v *ListHotNewsWithTypeResponseBody) *ListHotNewsWithTypeResponse {
	s.Body = v
	return s
}

type ListHotSourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 66
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListHotSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListHotSourcesRequest) SetAgentKey(v string) *ListHotSourcesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotSourcesRequest) SetMaxResults(v int32) *ListHotSourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotSourcesRequest) SetNextToken(v string) *ListHotSourcesRequest {
	s.NextToken = &v
	return s
}

type ListHotSourcesResponseBody struct {
	// example:
	//
	// NoData
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListHotSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHotSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotSourcesResponseBody) SetCode(v string) *ListHotSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetData(v []*ListHotSourcesResponseBodyData) *ListHotSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListHotSourcesResponseBody) SetHttpStatusCode(v int32) *ListHotSourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetMaxResults(v int32) *ListHotSourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetMessage(v string) *ListHotSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetNextToken(v string) *ListHotSourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetRequestId(v string) *ListHotSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetSuccess(v bool) *ListHotSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotSourcesResponseBody) SetTotalCount(v int32) *ListHotSourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListHotSourcesResponseBodyData struct {
	// example:
	//
	// 热榜源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Show *bool `json:"Show,omitempty" xml:"Show,omitempty"`
	// example:
	//
	// 86
	Sort *int32 `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 热榜源标识
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListHotSourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotSourcesResponseBodyData) SetDescription(v string) *ListHotSourcesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListHotSourcesResponseBodyData) SetShow(v bool) *ListHotSourcesResponseBodyData {
	s.Show = &v
	return s
}

func (s *ListHotSourcesResponseBodyData) SetSort(v int32) *ListHotSourcesResponseBodyData {
	s.Sort = &v
	return s
}

func (s *ListHotSourcesResponseBodyData) SetSource(v string) *ListHotSourcesResponseBodyData {
	s.Source = &v
	return s
}

type ListHotSourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListHotSourcesResponse) SetHeaders(v map[string]*string) *ListHotSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListHotSourcesResponse) SetStatusCode(v int32) *ListHotSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotSourcesResponse) SetBody(v *ListHotSourcesResponseBody) *ListHotSourcesResponse {
	s.Body = v
	return s
}

type ListHotTopicsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TopicIds  []*string `json:"TopicIds,omitempty" xml:"TopicIds,omitempty" type:"Repeated"`
	// example:
	//
	// 根据热榜主题全文检索
	TopicQuery *string `json:"TopicQuery,omitempty" xml:"TopicQuery,omitempty"`
	// example:
	//
	// 热榜源筛选，支持的热榜源。热榜源详见API：ListHotSources
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 数据版本筛选
	TopicVersion *string   `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	Topics       []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// example:
	//
	// true
	WithNews *bool `json:"WithNews,omitempty" xml:"WithNews,omitempty"`
}

func (s ListHotTopicsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListHotTopicsRequest) SetAgentKey(v string) *ListHotTopicsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotTopicsRequest) SetMaxResults(v int32) *ListHotTopicsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicsRequest) SetNextToken(v string) *ListHotTopicsRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopicIds(v []*string) *ListHotTopicsRequest {
	s.TopicIds = v
	return s
}

func (s *ListHotTopicsRequest) SetTopicQuery(v string) *ListHotTopicsRequest {
	s.TopicQuery = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopicSource(v string) *ListHotTopicsRequest {
	s.TopicSource = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopicVersion(v string) *ListHotTopicsRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopics(v []*string) *ListHotTopicsRequest {
	s.Topics = v
	return s
}

func (s *ListHotTopicsRequest) SetWithNews(v bool) *ListHotTopicsRequest {
	s.WithNews = &v
	return s
}

type ListHotTopicsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TopicIdsShrink *string `json:"TopicIds,omitempty" xml:"TopicIds,omitempty"`
	// example:
	//
	// 根据热榜主题全文检索
	TopicQuery *string `json:"TopicQuery,omitempty" xml:"TopicQuery,omitempty"`
	// example:
	//
	// 热榜源筛选，支持的热榜源。热榜源详见API：ListHotSources
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 数据版本筛选
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// example:
	//
	// true
	WithNews *bool `json:"WithNews,omitempty" xml:"WithNews,omitempty"`
}

func (s ListHotTopicsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotTopicsShrinkRequest) SetAgentKey(v string) *ListHotTopicsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetMaxResults(v int32) *ListHotTopicsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetNextToken(v string) *ListHotTopicsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicIdsShrink(v string) *ListHotTopicsShrinkRequest {
	s.TopicIdsShrink = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicQuery(v string) *ListHotTopicsShrinkRequest {
	s.TopicQuery = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicSource(v string) *ListHotTopicsShrinkRequest {
	s.TopicSource = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicVersion(v string) *ListHotTopicsShrinkRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicsShrink(v string) *ListHotTopicsShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetWithNews(v bool) *ListHotTopicsShrinkRequest {
	s.WithNews = &v
	return s
}

type ListHotTopicsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListHotTopicsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 94
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHotTopicsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBody) SetCode(v string) *ListHotTopicsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetData(v []*ListHotTopicsResponseBodyData) *ListHotTopicsResponseBody {
	s.Data = v
	return s
}

func (s *ListHotTopicsResponseBody) SetHttpStatusCode(v int32) *ListHotTopicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetMaxResults(v int32) *ListHotTopicsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetMessage(v string) *ListHotTopicsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetNextToken(v string) *ListHotTopicsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetRequestId(v string) *ListHotTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetSuccess(v bool) *ListHotTopicsResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotTopicsResponseBody) SetTotalCount(v int32) *ListHotTopicsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHotTopicsResponseBodyData struct {
	// example:
	//
	// 异步任务ID（自定义主题场景下使用）
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 创建用户ID（自定义主题场景下使用）
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 61
	HotValue *int64 `json:"HotValue,omitempty" xml:"HotValue,omitempty"`
	// example:
	//
	// 热榜ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// FAILED
	Status           *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureSummary []*ListHotTopicsResponseBodyDataStructureSummary `json:"StructureSummary,omitempty" xml:"StructureSummary,omitempty" type:"Repeated"`
	// example:
	//
	// 热榜摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 异步任务失败错误信息
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 26
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 主题唯一名称
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 热榜源，目前支持的热榜源: Toutiao：头条、Quark：夸克、Baidu：百度、Sina：新浪。Custom：自定义、Aggregation：热点话题榜
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 数据版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListHotTopicsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBodyData) SetAsyncTaskId(v string) *ListHotTopicsResponseBodyData {
	s.AsyncTaskId = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetCreateUser(v string) *ListHotTopicsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetHotValue(v int64) *ListHotTopicsResponseBodyData {
	s.HotValue = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetId(v string) *ListHotTopicsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetStatus(v string) *ListHotTopicsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetStructureSummary(v []*ListHotTopicsResponseBodyDataStructureSummary) *ListHotTopicsResponseBodyData {
	s.StructureSummary = v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetSummary(v string) *ListHotTopicsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTaskErrorMessage(v string) *ListHotTopicsResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTaskStatus(v int32) *ListHotTopicsResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTopic(v string) *ListHotTopicsResponseBodyData {
	s.Topic = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetTopicSource(v string) *ListHotTopicsResponseBodyData {
	s.TopicSource = &v
	return s
}

func (s *ListHotTopicsResponseBodyData) SetVersion(v string) *ListHotTopicsResponseBodyData {
	s.Version = &v
	return s
}

type ListHotTopicsResponseBodyDataStructureSummary struct {
	DocList []*ListHotTopicsResponseBodyDataStructureSummaryDocList `json:"DocList,omitempty" xml:"DocList,omitempty" type:"Repeated"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListHotTopicsResponseBodyDataStructureSummary) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicsResponseBodyDataStructureSummary) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) SetDocList(v []*ListHotTopicsResponseBodyDataStructureSummaryDocList) *ListHotTopicsResponseBodyDataStructureSummary {
	s.DocList = v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) SetSummary(v string) *ListHotTopicsResponseBodyDataStructureSummary {
	s.Summary = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummary) SetTitle(v string) *ListHotTopicsResponseBodyDataStructureSummary {
	s.Title = &v
	return s
}

type ListHotTopicsResponseBodyDataStructureSummaryDocList struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// xxxxx
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotTopicsResponseBodyDataStructureSummaryDocList) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicsResponseBodyDataStructureSummaryDocList) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) SetSource(v string) *ListHotTopicsResponseBodyDataStructureSummaryDocList {
	s.Source = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) SetTitle(v string) *ListHotTopicsResponseBodyDataStructureSummaryDocList {
	s.Title = &v
	return s
}

func (s *ListHotTopicsResponseBodyDataStructureSummaryDocList) SetUrl(v string) *ListHotTopicsResponseBodyDataStructureSummaryDocList {
	s.Url = &v
	return s
}

type ListHotTopicsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotTopicsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponse) SetHeaders(v map[string]*string) *ListHotTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListHotTopicsResponse) SetStatusCode(v int32) *ListHotTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotTopicsResponse) SetBody(v *ListHotTopicsResponseBody) *ListHotTopicsResponse {
	s.Body = v
	return s
}

type ListHotViewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 56
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
}

func (s ListHotViewPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotViewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsRequest) SetAgentKey(v string) *ListHotViewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotViewPointsRequest) SetMaxResults(v int32) *ListHotViewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotViewPointsRequest) SetNextToken(v string) *ListHotViewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotViewPointsRequest) SetTopic(v string) *ListHotViewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListHotViewPointsRequest) SetTopicSource(v string) *ListHotViewPointsRequest {
	s.TopicSource = &v
	return s
}

type ListHotViewPointsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListHotViewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 67
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 70
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHotViewPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotViewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBody) SetCode(v string) *ListHotViewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetData(v []*ListHotViewPointsResponseBodyData) *ListHotViewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListHotViewPointsResponseBody) SetHttpStatusCode(v int32) *ListHotViewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetMaxResults(v int32) *ListHotViewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetMessage(v string) *ListHotViewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetNextToken(v string) *ListHotViewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetRequestId(v string) *ListHotViewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetSuccess(v bool) *ListHotViewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotViewPointsResponseBody) SetTotalCount(v int32) *ListHotViewPointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHotViewPointsResponseBodyData struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                  `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	News         []*ListHotViewPointsResponseBodyDataNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                        `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*ListHotViewPointsResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListHotViewPointsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotViewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyData) SetAttitude(v string) *ListHotViewPointsResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetAttitudeType(v string) *ListHotViewPointsResponseBodyData {
	s.AttitudeType = &v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetNews(v []*ListHotViewPointsResponseBodyDataNews) *ListHotViewPointsResponseBodyData {
	s.News = v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetRatio(v string) *ListHotViewPointsResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *ListHotViewPointsResponseBodyData) SetViewPoints(v []*ListHotViewPointsResponseBodyDataViewPoints) *ListHotViewPointsResponseBodyData {
	s.ViewPoints = v
	return s
}

type ListHotViewPointsResponseBodyDataNews struct {
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
	// xxxxx
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 123456
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.png
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-01-22 10:29:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// ["标签1","标签2"]
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 文章主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotViewPointsResponseBodyDataNews) String() string {
	return tea.Prettify(s)
}

func (s ListHotViewPointsResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyDataNews) SetAuthor(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Author = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetContent(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetDocId(v string) *ListHotViewPointsResponseBodyDataNews {
	s.DocId = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetDocUuid(v string) *ListHotViewPointsResponseBodyDataNews {
	s.DocUuid = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetImageUrls(v []*string) *ListHotViewPointsResponseBodyDataNews {
	s.ImageUrls = v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetPubTime(v string) *ListHotViewPointsResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetSource(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Source = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetSummary(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Summary = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetTags(v []*string) *ListHotViewPointsResponseBodyDataNews {
	s.Tags = v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetTitle(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetTopic(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Topic = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataNews) SetUrl(v string) *ListHotViewPointsResponseBodyDataNews {
	s.Url = &v
	return s
}

type ListHotViewPointsResponseBodyDataViewPoints struct {
	Outlines []*ListHotViewPointsResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListHotViewPointsResponseBodyDataViewPoints) String() string {
	return tea.Prettify(s)
}

func (s ListHotViewPointsResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) SetOutlines(v []*ListHotViewPointsResponseBodyDataViewPointsOutlines) *ListHotViewPointsResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) SetPoint(v string) *ListHotViewPointsResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPoints) SetSummary(v string) *ListHotViewPointsResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

type ListHotViewPointsResponseBodyDataViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListHotViewPointsResponseBodyDataViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s ListHotViewPointsResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListHotViewPointsResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListHotViewPointsResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListHotViewPointsResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

type ListHotViewPointsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotViewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotViewPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotViewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsResponse) SetHeaders(v map[string]*string) *ListHotViewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListHotViewPointsResponse) SetStatusCode(v int32) *ListHotViewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotViewPointsResponse) SetBody(v *ListHotViewPointsResponseBody) *ListHotViewPointsResponse {
	s.Body = v
	return s
}

type ListInterveneCntRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneCntRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneCntRequest) SetAgentKey(v string) *ListInterveneCntRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneCntRequest) SetPageIndex(v int32) *ListInterveneCntRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneCntRequest) SetPageSize(v int32) *ListInterveneCntRequest {
	s.PageSize = &v
	return s
}

type ListInterveneCntResponseBody struct {
	// example:
	//
	// 0
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInterveneCntResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneCntResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponseBody) SetCode(v string) *ListInterveneCntResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetData(v *ListInterveneCntResponseBodyData) *ListInterveneCntResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneCntResponseBody) SetHttpStatusCode(v int32) *ListInterveneCntResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetMessage(v string) *ListInterveneCntResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetRequestId(v string) *ListInterveneCntResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetSuccess(v bool) *ListInterveneCntResponseBody {
	s.Success = &v
	return s
}

type ListInterveneCntResponseBodyData struct {
	CntList []interface{} `json:"CntList,omitempty" xml:"CntList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageCnt *int32 `json:"PageCnt,omitempty" xml:"PageCnt,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneCntResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponseBodyData) SetCntList(v []interface{}) *ListInterveneCntResponseBodyData {
	s.CntList = v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageCnt(v int32) *ListInterveneCntResponseBodyData {
	s.PageCnt = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageIndex(v int32) *ListInterveneCntResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageSize(v int32) *ListInterveneCntResponseBodyData {
	s.PageSize = &v
	return s
}

type ListInterveneCntResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneCntResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneCntResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponse) SetHeaders(v map[string]*string) *ListInterveneCntResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneCntResponse) SetStatusCode(v int32) *ListInterveneCntResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneCntResponse) SetBody(v *ListInterveneCntResponseBody) *ListInterveneCntResponse {
	s.Body = v
	return s
}

type ListInterveneImportTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneImportTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksRequest) SetAgentKey(v string) *ListInterveneImportTasksRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneImportTasksRequest) SetPageIndex(v int32) *ListInterveneImportTasksRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneImportTasksRequest) SetPageSize(v int32) *ListInterveneImportTasksRequest {
	s.PageSize = &v
	return s
}

type ListInterveneImportTasksResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInterveneImportTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneImportTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBody) SetCode(v string) *ListInterveneImportTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetData(v *ListInterveneImportTasksResponseBodyData) *ListInterveneImportTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetHttpStatusCode(v int32) *ListInterveneImportTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetMessage(v string) *ListInterveneImportTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetRequestId(v string) *ListInterveneImportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetSuccess(v bool) *ListInterveneImportTasksResponseBody {
	s.Success = &v
	return s
}

type ListInterveneImportTasksResponseBodyData struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList []*ListInterveneImportTasksResponseBodyDataStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListInterveneImportTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBodyData) SetPageIndex(v int32) *ListInterveneImportTasksResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetPageSize(v int32) *ListInterveneImportTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetStatusList(v []*ListInterveneImportTasksResponseBodyDataStatusList) *ListInterveneImportTasksResponseBodyData {
	s.StatusList = v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetTotalSize(v int32) *ListInterveneImportTasksResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListInterveneImportTasksResponseBodyDataStatusList struct {
	// example:
	//
	// Success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 5
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// example:
	//
	// Success
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 4854
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 12344454
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListInterveneImportTasksResponseBodyDataStatusList) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponseBodyDataStatusList) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetMsg(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Msg = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetPercentage(v int32) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Percentage = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetStatus(v int32) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Status = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetTaskId(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.TaskId = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetTaskName(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.TaskName = &v
	return s
}

type ListInterveneImportTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneImportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneImportTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponse) SetHeaders(v map[string]*string) *ListInterveneImportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneImportTasksResponse) SetStatusCode(v int32) *ListInterveneImportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneImportTasksResponse) SetBody(v *ListInterveneImportTasksResponseBody) *ListInterveneImportTasksResponse {
	s.Body = v
	return s
}

type ListInterveneRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesRequest) SetAgentKey(v string) *ListInterveneRulesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneRulesRequest) SetPageIndex(v int32) *ListInterveneRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneRulesRequest) SetPageSize(v int32) *ListInterveneRulesRequest {
	s.PageSize = &v
	return s
}

type ListInterveneRulesResponseBody struct {
	// example:
	//
	// 0
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListInterveneRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBody) SetCode(v string) *ListInterveneRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetData(v *ListInterveneRulesResponseBodyData) *ListInterveneRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneRulesResponseBody) SetHttpStatusCode(v int32) *ListInterveneRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetMessage(v string) *ListInterveneRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetRequestId(v string) *ListInterveneRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetSuccess(v bool) *ListInterveneRulesResponseBody {
	s.Success = &v
	return s
}

type ListInterveneRulesResponseBodyData struct {
	// example:
	//
	// 1
	Count             *int64                                                 `json:"Count,omitempty" xml:"Count,omitempty"`
	InterveneRuleList []*ListInterveneRulesResponseBodyDataInterveneRuleList `json:"InterveneRuleList,omitempty" xml:"InterveneRuleList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyData) SetCount(v int64) *ListInterveneRulesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetInterveneRuleList(v []*ListInterveneRulesResponseBodyDataInterveneRuleList) *ListInterveneRulesResponseBodyData {
	s.InterveneRuleList = v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetPageIndex(v int32) *ListInterveneRulesResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetPageSize(v int32) *ListInterveneRulesResponseBodyData {
	s.PageSize = &v
	return s
}

type ListInterveneRulesResponseBodyDataInterveneRuleList struct {
	AnswerConfig []*ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-06-05 15:17:01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2023-04-03 02:42:01
	EffectTime *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// example:
	//
	// 0
	InterveneType *int32    `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// example:
	//
	// mr-iuo9pi9w555phfbb
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// ruletest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleList) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetAnswerConfig(v []*ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.AnswerConfig = v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetCreateTime(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.CreateTime = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetEffectTime(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.EffectTime = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetInterveneType(v int32) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.InterveneType = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetNamespaceList(v []*string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.NamespaceList = v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetRuleId(v int64) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.RuleId = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetRuleName(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.RuleName = &v
	return s
}

type ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig struct {
	// example:
	//
	// 0
	AnswerType *int32  `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// namespace_qa_query
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetAnswerType(v int32) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetMessage(v string) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.Message = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetNamespace(v string) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.Namespace = &v
	return s
}

type ListInterveneRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponse) SetHeaders(v map[string]*string) *ListInterveneRulesResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneRulesResponse) SetStatusCode(v int32) *ListInterveneRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneRulesResponse) SetBody(v *ListInterveneRulesResponseBody) *ListInterveneRulesResponse {
	s.Body = v
	return s
}

type ListIntervenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneType *int32  `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query    *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// mqtt_outRule_1679019634514
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListIntervenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ListIntervenesRequest) SetAgentKey(v string) *ListIntervenesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListIntervenesRequest) SetInterveneType(v int32) *ListIntervenesRequest {
	s.InterveneType = &v
	return s
}

func (s *ListIntervenesRequest) SetPageIndex(v int32) *ListIntervenesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListIntervenesRequest) SetPageSize(v int32) *ListIntervenesRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntervenesRequest) SetQuery(v string) *ListIntervenesRequest {
	s.Query = &v
	return s
}

func (s *ListIntervenesRequest) SetRuleId(v int64) *ListIntervenesRequest {
	s.RuleId = &v
	return s
}

type ListIntervenesResponseBody struct {
	// example:
	//
	// 0
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIntervenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBody) SetCode(v string) *ListIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervenesResponseBody) SetData(v *ListIntervenesResponseBodyData) *ListIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervenesResponseBody) SetHttpStatusCode(v int32) *ListIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervenesResponseBody) SetMessage(v string) *ListIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervenesResponseBody) SetRequestId(v string) *ListIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervenesResponseBody) SetSuccess(v bool) *ListIntervenesResponseBody {
	s.Success = &v
	return s
}

type ListIntervenesResponseBodyData struct {
	InterveneList []*ListIntervenesResponseBodyDataInterveneList `json:"InterveneList,omitempty" xml:"InterveneList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListIntervenesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBodyData) SetInterveneList(v []*ListIntervenesResponseBodyDataInterveneList) *ListIntervenesResponseBodyData {
	s.InterveneList = v
	return s
}

func (s *ListIntervenesResponseBodyData) SetPageIndex(v int32) *ListIntervenesResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListIntervenesResponseBodyData) SetPageSize(v int32) *ListIntervenesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIntervenesResponseBodyData) SetTotalSize(v int64) *ListIntervenesResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListIntervenesResponseBodyDataInterveneList struct {
	// id
	//
	// example:
	//
	// 36559
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListIntervenesResponseBodyDataInterveneList) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponseBodyDataInterveneList) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBodyDataInterveneList) SetId(v string) *ListIntervenesResponseBodyDataInterveneList {
	s.Id = &v
	return s
}

func (s *ListIntervenesResponseBodyDataInterveneList) SetQuery(v string) *ListIntervenesResponseBodyDataInterveneList {
	s.Query = &v
	return s
}

type ListIntervenesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntervenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponse) SetHeaders(v map[string]*string) *ListIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ListIntervenesResponse) SetStatusCode(v int32) *ListIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervenesResponse) SetBody(v *ListIntervenesResponseBody) *ListIntervenesResponse {
	s.Body = v
	return s
}

type ListMaterialDocumentsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// jsonLine
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// excel
	DocTypeList []*string `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	GeneratePublicUrl *bool `json:"GeneratePublicUrl,omitempty" xml:"GeneratePublicUrl,omitempty"`
	// example:
	//
	// 69
	Id       *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Query    *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// 10
	Size  *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2023-03-18 03:00:00
	UpdateTimeEnd *string `json:"UpdateTimeEnd,omitempty" xml:"UpdateTimeEnd,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	UpdateTimeStart *string `json:"UpdateTimeStart,omitempty" xml:"UpdateTimeStart,omitempty"`
}

func (s ListMaterialDocumentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsRequest) SetAgentKey(v string) *ListMaterialDocumentsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetContent(v string) *ListMaterialDocumentsRequest {
	s.Content = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCreateTimeEnd(v string) *ListMaterialDocumentsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCreateTimeStart(v string) *ListMaterialDocumentsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCurrent(v int32) *ListMaterialDocumentsRequest {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetDocType(v string) *ListMaterialDocumentsRequest {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetDocTypeList(v []*string) *ListMaterialDocumentsRequest {
	s.DocTypeList = v
	return s
}

func (s *ListMaterialDocumentsRequest) SetGeneratePublicUrl(v bool) *ListMaterialDocumentsRequest {
	s.GeneratePublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetId(v int64) *ListMaterialDocumentsRequest {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetKeywords(v []*string) *ListMaterialDocumentsRequest {
	s.Keywords = v
	return s
}

func (s *ListMaterialDocumentsRequest) SetQuery(v string) *ListMaterialDocumentsRequest {
	s.Query = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetShareAttr(v int32) *ListMaterialDocumentsRequest {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetSize(v int32) *ListMaterialDocumentsRequest {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetTitle(v string) *ListMaterialDocumentsRequest {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetUpdateTimeEnd(v string) *ListMaterialDocumentsRequest {
	s.UpdateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetUpdateTimeStart(v string) *ListMaterialDocumentsRequest {
	s.UpdateTimeStart = &v
	return s
}

type ListMaterialDocumentsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// jsonLine
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// excel
	DocTypeListShrink *string `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty"`
	// example:
	//
	// true
	GeneratePublicUrl *bool `json:"GeneratePublicUrl,omitempty" xml:"GeneratePublicUrl,omitempty"`
	// example:
	//
	// 69
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// 10
	Size  *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2023-03-18 03:00:00
	UpdateTimeEnd *string `json:"UpdateTimeEnd,omitempty" xml:"UpdateTimeEnd,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	UpdateTimeStart *string `json:"UpdateTimeStart,omitempty" xml:"UpdateTimeStart,omitempty"`
}

func (s ListMaterialDocumentsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsShrinkRequest) SetAgentKey(v string) *ListMaterialDocumentsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetContent(v string) *ListMaterialDocumentsShrinkRequest {
	s.Content = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCreateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCreateTimeStart(v string) *ListMaterialDocumentsShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCurrent(v int32) *ListMaterialDocumentsShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetDocType(v string) *ListMaterialDocumentsShrinkRequest {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetDocTypeListShrink(v string) *ListMaterialDocumentsShrinkRequest {
	s.DocTypeListShrink = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetGeneratePublicUrl(v bool) *ListMaterialDocumentsShrinkRequest {
	s.GeneratePublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetId(v int64) *ListMaterialDocumentsShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetKeywordsShrink(v string) *ListMaterialDocumentsShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetQuery(v string) *ListMaterialDocumentsShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetShareAttr(v int32) *ListMaterialDocumentsShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetSize(v int32) *ListMaterialDocumentsShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetTitle(v string) *ListMaterialDocumentsShrinkRequest {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetUpdateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest {
	s.UpdateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetUpdateTimeStart(v string) *ListMaterialDocumentsShrinkRequest {
	s.UpdateTimeStart = &v
	return s
}

type ListMaterialDocumentsResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                                   `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListMaterialDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMaterialDocumentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponseBody) SetCode(v string) *ListMaterialDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetCurrent(v int32) *ListMaterialDocumentsResponseBody {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetData(v []*ListMaterialDocumentsResponseBodyData) *ListMaterialDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetHttpStatusCode(v int32) *ListMaterialDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetMessage(v string) *ListMaterialDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetRequestId(v string) *ListMaterialDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetSize(v int32) *ListMaterialDocumentsResponseBody {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetSuccess(v bool) *ListMaterialDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetTotal(v int32) *ListMaterialDocumentsResponseBody {
	s.Total = &v
	return s
}

type ListMaterialDocumentsResponseBodyData struct {
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser     *string   `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	CreateUserName *string   `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DocKeywords    []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// https://www.example.com
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// example:
	//
	// 35
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// https://www.example.com
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// user_upload
	SrcFrom           *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary           *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent       *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	ThumbnailInBase64 *string `json:"ThumbnailInBase64,omitempty" xml:"ThumbnailInBase64,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser     *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListMaterialDocumentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponseBodyData) SetAuthor(v string) *ListMaterialDocumentsResponseBodyData {
	s.Author = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateUser(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateUserName(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetDocKeywords(v []*string) *ListMaterialDocumentsResponseBodyData {
	s.DocKeywords = v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetDocType(v string) *ListMaterialDocumentsResponseBodyData {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetExternalUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.ExternalUrl = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetHtmlContent(v string) *ListMaterialDocumentsResponseBodyData {
	s.HtmlContent = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetId(v int64) *ListMaterialDocumentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetPubTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetPublicUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetShareAttr(v int32) *ListMaterialDocumentsResponseBodyData {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetSrcFrom(v string) *ListMaterialDocumentsResponseBodyData {
	s.SrcFrom = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetSummary(v string) *ListMaterialDocumentsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetTextContent(v string) *ListMaterialDocumentsResponseBodyData {
	s.TextContent = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetThumbnailInBase64(v string) *ListMaterialDocumentsResponseBodyData {
	s.ThumbnailInBase64 = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetTitle(v string) *ListMaterialDocumentsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateUser(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateUserName(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.Url = &v
	return s
}

type ListMaterialDocumentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMaterialDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMaterialDocumentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponse) SetHeaders(v map[string]*string) *ListMaterialDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListMaterialDocumentsResponse) SetStatusCode(v int32) *ListMaterialDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMaterialDocumentsResponse) SetBody(v *ListMaterialDocumentsResponseBody) *ListMaterialDocumentsResponse {
	s.Body = v
	return s
}

type ListPlanningProposalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// e7b26a9e1211444db8f0a984361a5e0f
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 27971fc8f3ce4ed58c7e7fc4b503e432
	CustomViewPointIds []*string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty" type:"Repeated"`
	// example:
	//
	// 73
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 标题
	//
	//      *
	Titles []*string `json:"Titles,omitempty" xml:"Titles,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 2024-09-10_08
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CustomViewPoints
	ViewPointType *string `json:"ViewPointType,omitempty" xml:"ViewPointType,omitempty"`
}

func (s ListPlanningProposalRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlanningProposalRequest) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalRequest) SetAgentKey(v string) *ListPlanningProposalRequest {
	s.AgentKey = &v
	return s
}

func (s *ListPlanningProposalRequest) SetCustomViewPointId(v string) *ListPlanningProposalRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListPlanningProposalRequest) SetCustomViewPointIds(v []*string) *ListPlanningProposalRequest {
	s.CustomViewPointIds = v
	return s
}

func (s *ListPlanningProposalRequest) SetMaxResults(v int32) *ListPlanningProposalRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPlanningProposalRequest) SetNextToken(v string) *ListPlanningProposalRequest {
	s.NextToken = &v
	return s
}

func (s *ListPlanningProposalRequest) SetTitles(v []*string) *ListPlanningProposalRequest {
	s.Titles = v
	return s
}

func (s *ListPlanningProposalRequest) SetTopic(v string) *ListPlanningProposalRequest {
	s.Topic = &v
	return s
}

func (s *ListPlanningProposalRequest) SetTopicSource(v string) *ListPlanningProposalRequest {
	s.TopicSource = &v
	return s
}

func (s *ListPlanningProposalRequest) SetTopicVersion(v string) *ListPlanningProposalRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListPlanningProposalRequest) SetViewPointType(v string) *ListPlanningProposalRequest {
	s.ViewPointType = &v
	return s
}

type ListPlanningProposalShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// e7b26a9e1211444db8f0a984361a5e0f
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 27971fc8f3ce4ed58c7e7fc4b503e432
	CustomViewPointIdsShrink *string `json:"CustomViewPointIds,omitempty" xml:"CustomViewPointIds,omitempty"`
	// example:
	//
	// 73
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 标题
	//
	//      *
	TitlesShrink *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 2024-09-10_08
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CustomViewPoints
	ViewPointType *string `json:"ViewPointType,omitempty" xml:"ViewPointType,omitempty"`
}

func (s ListPlanningProposalShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlanningProposalShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalShrinkRequest) SetAgentKey(v string) *ListPlanningProposalShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetCustomViewPointId(v string) *ListPlanningProposalShrinkRequest {
	s.CustomViewPointId = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetCustomViewPointIdsShrink(v string) *ListPlanningProposalShrinkRequest {
	s.CustomViewPointIdsShrink = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetMaxResults(v int32) *ListPlanningProposalShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetNextToken(v string) *ListPlanningProposalShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTitlesShrink(v string) *ListPlanningProposalShrinkRequest {
	s.TitlesShrink = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTopic(v string) *ListPlanningProposalShrinkRequest {
	s.Topic = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTopicSource(v string) *ListPlanningProposalShrinkRequest {
	s.TopicSource = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetTopicVersion(v string) *ListPlanningProposalShrinkRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListPlanningProposalShrinkRequest) SetViewPointType(v string) *ListPlanningProposalShrinkRequest {
	s.ViewPointType = &v
	return s
}

type ListPlanningProposalResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListPlanningProposalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 77
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 80
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPlanningProposalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPlanningProposalResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponseBody) SetCode(v string) *ListPlanningProposalResponseBody {
	s.Code = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetData(v []*ListPlanningProposalResponseBodyData) *ListPlanningProposalResponseBody {
	s.Data = v
	return s
}

func (s *ListPlanningProposalResponseBody) SetHttpStatusCode(v int32) *ListPlanningProposalResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetMaxResults(v int32) *ListPlanningProposalResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetMessage(v string) *ListPlanningProposalResponseBody {
	s.Message = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetNextToken(v string) *ListPlanningProposalResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetRequestId(v string) *ListPlanningProposalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetSuccess(v bool) *ListPlanningProposalResponseBody {
	s.Success = &v
	return s
}

func (s *ListPlanningProposalResponseBody) SetTotalCount(v int32) *ListPlanningProposalResponseBody {
	s.TotalCount = &v
	return s
}

type ListPlanningProposalResponseBodyData struct {
	Outlines []*ListPlanningProposalResponseBodyDataOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	Summary  *string                                         `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title    *string                                         `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListPlanningProposalResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPlanningProposalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponseBodyData) SetOutlines(v []*ListPlanningProposalResponseBodyDataOutlines) *ListPlanningProposalResponseBodyData {
	s.Outlines = v
	return s
}

func (s *ListPlanningProposalResponseBodyData) SetSummary(v string) *ListPlanningProposalResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListPlanningProposalResponseBodyData) SetTitle(v string) *ListPlanningProposalResponseBodyData {
	s.Title = &v
	return s
}

type ListPlanningProposalResponseBodyDataOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListPlanningProposalResponseBodyDataOutlines) String() string {
	return tea.Prettify(s)
}

func (s ListPlanningProposalResponseBodyDataOutlines) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponseBodyDataOutlines) SetOutline(v string) *ListPlanningProposalResponseBodyDataOutlines {
	s.Outline = &v
	return s
}

func (s *ListPlanningProposalResponseBodyDataOutlines) SetSummary(v string) *ListPlanningProposalResponseBodyDataOutlines {
	s.Summary = &v
	return s
}

type ListPlanningProposalResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlanningProposalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlanningProposalResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPlanningProposalResponse) GoString() string {
	return s.String()
}

func (s *ListPlanningProposalResponse) SetHeaders(v map[string]*string) *ListPlanningProposalResponse {
	s.Headers = v
	return s
}

func (s *ListPlanningProposalResponse) SetStatusCode(v int32) *ListPlanningProposalResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlanningProposalResponse) SetBody(v *ListPlanningProposalResponseBody) *ListPlanningProposalResponse {
	s.Body = v
	return s
}

type ListTimedViewAttitudeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 53
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
}

func (s ListTimedViewAttitudeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTimedViewAttitudeRequest) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeRequest) SetAgentKey(v string) *ListTimedViewAttitudeRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetMaxResults(v int32) *ListTimedViewAttitudeRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetNextToken(v string) *ListTimedViewAttitudeRequest {
	s.NextToken = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetTopic(v string) *ListTimedViewAttitudeRequest {
	s.Topic = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetTopicSource(v string) *ListTimedViewAttitudeRequest {
	s.TopicSource = &v
	return s
}

type ListTimedViewAttitudeResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListTimedViewAttitudeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 15
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 58
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTimedViewAttitudeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBody) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBody) SetCode(v string) *ListTimedViewAttitudeResponseBody {
	s.Code = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetData(v []*ListTimedViewAttitudeResponseBodyData) *ListTimedViewAttitudeResponseBody {
	s.Data = v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetHttpStatusCode(v int32) *ListTimedViewAttitudeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetMaxResults(v int32) *ListTimedViewAttitudeResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetMessage(v string) *ListTimedViewAttitudeResponseBody {
	s.Message = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetNextToken(v string) *ListTimedViewAttitudeResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetRequestId(v string) *ListTimedViewAttitudeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetSuccess(v bool) *ListTimedViewAttitudeResponseBody {
	s.Success = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBody) SetTotalCount(v int32) *ListTimedViewAttitudeResponseBody {
	s.TotalCount = &v
	return s
}

type ListTimedViewAttitudeResponseBodyData struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	// example:
	//
	// 2024-01-22 10:29
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 当前观点占比
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com/news/1.html
	Url        *string                                            `json:"Url,omitempty" xml:"Url,omitempty"`
	ViewPoints []*ListTimedViewAttitudeResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListTimedViewAttitudeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBodyData) SetAttitude(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetAttitudeType(v string) *ListTimedViewAttitudeResponseBodyData {
	s.AttitudeType = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetPubTime(v string) *ListTimedViewAttitudeResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetRatio(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetSource(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetTitle(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetUrl(v string) *ListTimedViewAttitudeResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyData) SetViewPoints(v []*ListTimedViewAttitudeResponseBodyDataViewPoints) *ListTimedViewAttitudeResponseBodyData {
	s.ViewPoints = v
	return s
}

type ListTimedViewAttitudeResponseBodyDataViewPoints struct {
	Outlines []*ListTimedViewAttitudeResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListTimedViewAttitudeResponseBodyDataViewPoints) String() string {
	return tea.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) SetOutlines(v []*ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) *ListTimedViewAttitudeResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) SetPoint(v string) *ListTimedViewAttitudeResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPoints) SetSummary(v string) *ListTimedViewAttitudeResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

type ListTimedViewAttitudeResponseBodyDataViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListTimedViewAttitudeResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

type ListTimedViewAttitudeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTimedViewAttitudeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTimedViewAttitudeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTimedViewAttitudeResponse) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponse) SetHeaders(v map[string]*string) *ListTimedViewAttitudeResponse {
	s.Headers = v
	return s
}

func (s *ListTimedViewAttitudeResponse) SetStatusCode(v int32) *ListTimedViewAttitudeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTimedViewAttitudeResponse) SetBody(v *ListTimedViewAttitudeResponseBody) *ListTimedViewAttitudeResponse {
	s.Body = v
	return s
}

type ListTopicRecommendEventListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 72
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTopicRecommendEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTopicRecommendEventListRequest) GoString() string {
	return s.String()
}

func (s *ListTopicRecommendEventListRequest) SetAgentKey(v string) *ListTopicRecommendEventListRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTopicRecommendEventListRequest) SetMaxResults(v int32) *ListTopicRecommendEventListRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTopicRecommendEventListRequest) SetNextToken(v string) *ListTopicRecommendEventListRequest {
	s.NextToken = &v
	return s
}

type ListTopicRecommendEventListResponseBody struct {
	// example:
	//
	// NoData
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 71
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// x\\"x\\"x
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 60
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTopicRecommendEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTopicRecommendEventListResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicRecommendEventListResponseBody) SetCode(v string) *ListTopicRecommendEventListResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetData(v []*string) *ListTopicRecommendEventListResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetHttpStatusCode(v int32) *ListTopicRecommendEventListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetMaxResults(v int32) *ListTopicRecommendEventListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetMessage(v string) *ListTopicRecommendEventListResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetNextToken(v string) *ListTopicRecommendEventListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetRequestId(v string) *ListTopicRecommendEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetSuccess(v bool) *ListTopicRecommendEventListResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicRecommendEventListResponseBody) SetTotalCount(v int32) *ListTopicRecommendEventListResponseBody {
	s.TotalCount = &v
	return s
}

type ListTopicRecommendEventListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicRecommendEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicRecommendEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTopicRecommendEventListResponse) GoString() string {
	return s.String()
}

func (s *ListTopicRecommendEventListResponse) SetHeaders(v map[string]*string) *ListTopicRecommendEventListResponse {
	s.Headers = v
	return s
}

func (s *ListTopicRecommendEventListResponse) SetStatusCode(v int32) *ListTopicRecommendEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicRecommendEventListResponse) SetBody(v *ListTopicRecommendEventListResponseBody) *ListTopicRecommendEventListResponse {
	s.Body = v
	return s
}

type ListTopicViewPointRecommendEventListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 66
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTopicViewPointRecommendEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTopicViewPointRecommendEventListRequest) GoString() string {
	return s.String()
}

func (s *ListTopicViewPointRecommendEventListRequest) SetAgentKey(v string) *ListTopicViewPointRecommendEventListRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListRequest) SetId(v string) *ListTopicViewPointRecommendEventListRequest {
	s.Id = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListRequest) SetMaxResults(v int32) *ListTopicViewPointRecommendEventListRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListRequest) SetNextToken(v string) *ListTopicViewPointRecommendEventListRequest {
	s.NextToken = &v
	return s
}

type ListTopicViewPointRecommendEventListResponseBody struct {
	// example:
	//
	// NoData
	Code *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 8
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTopicViewPointRecommendEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTopicViewPointRecommendEventListResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetCode(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetData(v []*string) *ListTopicViewPointRecommendEventListResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetHttpStatusCode(v int32) *ListTopicViewPointRecommendEventListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetMaxResults(v int32) *ListTopicViewPointRecommendEventListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetMessage(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetNextToken(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetRequestId(v string) *ListTopicViewPointRecommendEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetSuccess(v bool) *ListTopicViewPointRecommendEventListResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponseBody) SetTotalCount(v int32) *ListTopicViewPointRecommendEventListResponseBody {
	s.TotalCount = &v
	return s
}

type ListTopicViewPointRecommendEventListResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicViewPointRecommendEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicViewPointRecommendEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTopicViewPointRecommendEventListResponse) GoString() string {
	return s.String()
}

func (s *ListTopicViewPointRecommendEventListResponse) SetHeaders(v map[string]*string) *ListTopicViewPointRecommendEventListResponse {
	s.Headers = v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponse) SetStatusCode(v int32) *ListTopicViewPointRecommendEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponse) SetBody(v *ListTopicViewPointRecommendEventListResponseBody) *ListTopicViewPointRecommendEventListResponse {
	s.Body = v
	return s
}

type ListVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ListVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListVersionsRequest) SetAgentKey(v string) *ListVersionsRequest {
	s.AgentKey = &v
	return s
}

type ListVersionsResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBody) SetCode(v string) *ListVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVersionsResponseBody) SetData(v []*ListVersionsResponseBodyData) *ListVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListVersionsResponseBody) SetHttpStatusCode(v int32) *ListVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVersionsResponseBody) SetMessage(v string) *ListVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVersionsResponseBody) SetRequestId(v string) *ListVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVersionsResponseBody) SetSuccess(v bool) *ListVersionsResponseBody {
	s.Success = &v
	return s
}

type ListVersionsResponseBodyData struct {
	// example:
	//
	// 43
	ConcurrentCount *int32 `json:"ConcurrentCount,omitempty" xml:"ConcurrentCount,omitempty"`
	// example:
	//
	// 2023-04-23 02:00:34
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 55
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// example:
	//
	// ga-bp12pismsw4v3tzhf62p1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 7
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// CUSTOMIZE
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 13
	Quota *int32 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// 2023-05-27 04:11:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 65
	UseQuota      *int32  `json:"UseQuota,omitempty" xml:"UseQuota,omitempty"`
	VersionDetail *string `json:"VersionDetail,omitempty" xml:"VersionDetail,omitempty"`
	// example:
	//
	// 试用版
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// example:
	//
	// 87
	VersionStatus *int32 `json:"VersionStatus,omitempty" xml:"VersionStatus,omitempty"`
}

func (s ListVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBodyData) SetConcurrentCount(v int32) *ListVersionsResponseBodyData {
	s.ConcurrentCount = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetEndTime(v string) *ListVersionsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetInstanceCount(v int32) *ListVersionsResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetInstanceId(v string) *ListVersionsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetOrderId(v int64) *ListVersionsResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetProductType(v string) *ListVersionsResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetQuota(v int32) *ListVersionsResponseBodyData {
	s.Quota = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetStartTime(v string) *ListVersionsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetUseQuota(v int32) *ListVersionsResponseBodyData {
	s.UseQuota = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionDetail(v string) *ListVersionsResponseBodyData {
	s.VersionDetail = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionName(v string) *ListVersionsResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionStatus(v int32) *ListVersionsResponseBodyData {
	s.VersionStatus = &v
	return s
}

type ListVersionsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListVersionsResponse) SetHeaders(v map[string]*string) *ListVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListVersionsResponse) SetStatusCode(v int32) *ListVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVersionsResponse) SetBody(v *ListVersionsResponseBody) *ListVersionsResponse {
	s.Body = v
	return s
}

type ListWebReviewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 81
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
}

func (s ListWebReviewPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWebReviewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsRequest) SetAgentKey(v string) *ListWebReviewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetMaxResults(v int32) *ListWebReviewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetNextToken(v string) *ListWebReviewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetTopic(v string) *ListWebReviewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetTopicSource(v string) *ListWebReviewPointsRequest {
	s.TopicSource = &v
	return s
}

type ListWebReviewPointsResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListWebReviewPointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 79
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWebReviewPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWebReviewPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBody) SetCode(v string) *ListWebReviewPointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetData(v []*ListWebReviewPointsResponseBodyData) *ListWebReviewPointsResponseBody {
	s.Data = v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetHttpStatusCode(v int32) *ListWebReviewPointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetMaxResults(v int32) *ListWebReviewPointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetMessage(v string) *ListWebReviewPointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetNextToken(v string) *ListWebReviewPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetRequestId(v string) *ListWebReviewPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetSuccess(v bool) *ListWebReviewPointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWebReviewPointsResponseBody) SetTotalCount(v int32) *ListWebReviewPointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListWebReviewPointsResponseBodyData struct {
	// example:
	//
	// 当前观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 观点类型
	AttitudeType *string                                        `json:"AttitudeType,omitempty" xml:"AttitudeType,omitempty"`
	Comments     []*ListWebReviewPointsResponseBodyDataComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	// example:
	//
	// 当前观点占比
	Ratio      *string                                          `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	ViewPoints []*ListWebReviewPointsResponseBodyDataViewPoints `json:"ViewPoints,omitempty" xml:"ViewPoints,omitempty" type:"Repeated"`
}

func (s ListWebReviewPointsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyData) SetAttitude(v string) *ListWebReviewPointsResponseBodyData {
	s.Attitude = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetAttitudeType(v string) *ListWebReviewPointsResponseBodyData {
	s.AttitudeType = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetComments(v []*ListWebReviewPointsResponseBodyDataComments) *ListWebReviewPointsResponseBodyData {
	s.Comments = v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetRatio(v string) *ListWebReviewPointsResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyData) SetViewPoints(v []*ListWebReviewPointsResponseBodyDataViewPoints) *ListWebReviewPointsResponseBodyData {
	s.ViewPoints = v
	return s
}

type ListWebReviewPointsResponseBodyDataComments struct {
	// example:
	//
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 评论内容
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 当前评论所属的URL
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 评论用户名
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListWebReviewPointsResponseBodyDataComments) String() string {
	return tea.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyDataComments) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetSource(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Source = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetText(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Text = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetTitle(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Title = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetUrl(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Url = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataComments) SetUsername(v string) *ListWebReviewPointsResponseBodyDataComments {
	s.Username = &v
	return s
}

type ListWebReviewPointsResponseBodyDataViewPoints struct {
	Outlines []*ListWebReviewPointsResponseBodyDataViewPointsOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 视角
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListWebReviewPointsResponseBodyDataViewPoints) String() string {
	return tea.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyDataViewPoints) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) SetOutlines(v []*ListWebReviewPointsResponseBodyDataViewPointsOutlines) *ListWebReviewPointsResponseBodyDataViewPoints {
	s.Outlines = v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) SetPoint(v string) *ListWebReviewPointsResponseBodyDataViewPoints {
	s.Point = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPoints) SetSummary(v string) *ListWebReviewPointsResponseBodyDataViewPoints {
	s.Summary = &v
	return s
}

type ListWebReviewPointsResponseBodyDataViewPointsOutlines struct {
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// example:
	//
	// 大纲摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListWebReviewPointsResponseBodyDataViewPointsOutlines) String() string {
	return tea.Prettify(s)
}

func (s ListWebReviewPointsResponseBodyDataViewPointsOutlines) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponseBodyDataViewPointsOutlines) SetOutline(v string) *ListWebReviewPointsResponseBodyDataViewPointsOutlines {
	s.Outline = &v
	return s
}

func (s *ListWebReviewPointsResponseBodyDataViewPointsOutlines) SetSummary(v string) *ListWebReviewPointsResponseBodyDataViewPointsOutlines {
	s.Summary = &v
	return s
}

type ListWebReviewPointsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebReviewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebReviewPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWebReviewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponse) SetHeaders(v map[string]*string) *ListWebReviewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListWebReviewPointsResponse) SetStatusCode(v int32) *ListWebReviewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebReviewPointsResponse) SetBody(v *ListWebReviewPointsResponseBody) *ListWebReviewPointsResponse {
	s.Body = v
	return s
}

type QueryAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskRequest) SetAgentKey(v string) *QueryAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryAsyncTaskRequest) SetTaskId(v string) *QueryAsyncTaskRequest {
	s.TaskId = &v
	return s
}

type QueryAsyncTaskResponseBody struct {
	// example:
	//
	// Success
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryAsyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 867C4ABE-4381-5BC2-9810-5A5F334F71CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponseBody) SetCode(v string) *QueryAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetData(v *QueryAsyncTaskResponseBodyData) *QueryAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetHttpStatusCode(v int32) *QueryAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetMessage(v string) *QueryAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetRequestId(v string) *QueryAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetSuccess(v bool) *QueryAsyncTaskResponseBody {
	s.Success = &v
	return s
}

type QueryAsyncTaskResponseBodyData struct {
	// example:
	//
	// 2021-07-25 14:34:33
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 12121
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// example:
	//
	// error
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// {}
	TaskIntermediateResult *string `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	TaskName               *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// {"fileKey":"oss://default/xxxx/xxxx/xxx","fileName":"xxxxx.doc"}
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// example:
	//
	// {}
	TaskProgressMessage *string `json:"TaskProgressMessage,omitempty" xml:"TaskProgressMessage,omitempty"`
	// example:
	//
	// {}
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// example:
	//
	// 3
	TaskRetryCount *string `json:"TaskRetryCount,omitempty" xml:"TaskRetryCount,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 2023-04-27 18:07:43
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 12121
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s QueryAsyncTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponseBodyData) SetCreateTime(v string) *QueryAsyncTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetCreateUser(v string) *QueryAsyncTaskResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskCode(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskCode = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskErrorMessage(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskId(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskIntermediateResult(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskIntermediateResult = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskName(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskParam(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskParam = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskProgressMessage(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskProgressMessage = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskResult(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskResult = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskRetryCount(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskRetryCount = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskStatus(v int32) *QueryAsyncTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetUpdateTime(v string) *QueryAsyncTaskResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetUpdateUser(v string) *QueryAsyncTaskResponseBodyData {
	s.UpdateUser = &v
	return s
}

type QueryAsyncTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponse) SetHeaders(v map[string]*string) *QueryAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryAsyncTaskResponse) SetStatusCode(v int32) *QueryAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAsyncTaskResponse) SetBody(v *QueryAsyncTaskResponseBody) *QueryAsyncTaskResponse {
	s.Body = v
	return s
}

type RunAbbreviationContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 创新政务社交媒体功能。鼓励各地区、各部门结合实际，开发政务社交媒体的特色功能，如在线咨询服务、政策解读、互动问答等，增强政务社交媒体的互动性和实用性，提升公众参与度。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunAbbreviationContentRequest) String() string {
	return tea.Prettify(s)
}

func (s RunAbbreviationContentRequest) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentRequest) SetContent(v string) *RunAbbreviationContentRequest {
	s.Content = &v
	return s
}

func (s *RunAbbreviationContentRequest) SetWorkspaceId(v string) *RunAbbreviationContentRequest {
	s.WorkspaceId = &v
	return s
}

type RunAbbreviationContentResponseBody struct {
	End     *bool                                      `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunAbbreviationContentResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunAbbreviationContentResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunAbbreviationContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunAbbreviationContentResponseBody) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBody) SetEnd(v bool) *RunAbbreviationContentResponseBody {
	s.End = &v
	return s
}

func (s *RunAbbreviationContentResponseBody) SetHeader(v *RunAbbreviationContentResponseBodyHeader) *RunAbbreviationContentResponseBody {
	s.Header = v
	return s
}

func (s *RunAbbreviationContentResponseBody) SetPayload(v *RunAbbreviationContentResponseBodyPayload) *RunAbbreviationContentResponseBody {
	s.Payload = v
	return s
}

func (s *RunAbbreviationContentResponseBody) SetRequestId(v string) *RunAbbreviationContentResponseBody {
	s.RequestId = &v
	return s
}

type RunAbbreviationContentResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunAbbreviationContentResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyHeader) SetErrorCode(v string) *RunAbbreviationContentResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetErrorMessage(v string) *RunAbbreviationContentResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetEvent(v string) *RunAbbreviationContentResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetEventInfo(v string) *RunAbbreviationContentResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetSessionId(v string) *RunAbbreviationContentResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetTaskId(v string) *RunAbbreviationContentResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetTraceId(v string) *RunAbbreviationContentResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunAbbreviationContentResponseBodyPayload struct {
	Output *RunAbbreviationContentResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunAbbreviationContentResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunAbbreviationContentResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyPayload) SetOutput(v *RunAbbreviationContentResponseBodyPayloadOutput) *RunAbbreviationContentResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayload) SetUsage(v *RunAbbreviationContentResponseBodyPayloadUsage) *RunAbbreviationContentResponseBodyPayload {
	s.Usage = v
	return s
}

type RunAbbreviationContentResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunAbbreviationContentResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyPayloadOutput) SetText(v string) *RunAbbreviationContentResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunAbbreviationContentResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunAbbreviationContentResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) SetInputTokens(v int64) *RunAbbreviationContentResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunAbbreviationContentResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunAbbreviationContentResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunAbbreviationContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunAbbreviationContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunAbbreviationContentResponse) String() string {
	return tea.Prettify(s)
}

func (s RunAbbreviationContentResponse) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponse) SetHeaders(v map[string]*string) *RunAbbreviationContentResponse {
	s.Headers = v
	return s
}

func (s *RunAbbreviationContentResponse) SetStatusCode(v int32) *RunAbbreviationContentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunAbbreviationContentResponse) SetBody(v *RunAbbreviationContentResponseBody) *RunAbbreviationContentResponse {
	s.Body = v
	return s
}

type RunContinueContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 创新政务社交媒体功能。鼓励各地区、各部门结合实际，开发政务社交媒体的特色功能，如在线咨询服务、政策解读、互动问答等，增强政务社交媒体的互动性和实用性，提升公众参与度。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunContinueContentRequest) String() string {
	return tea.Prettify(s)
}

func (s RunContinueContentRequest) GoString() string {
	return s.String()
}

func (s *RunContinueContentRequest) SetContent(v string) *RunContinueContentRequest {
	s.Content = &v
	return s
}

func (s *RunContinueContentRequest) SetWorkspaceId(v string) *RunContinueContentRequest {
	s.WorkspaceId = &v
	return s
}

type RunContinueContentResponseBody struct {
	End     *bool                                  `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunContinueContentResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunContinueContentResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunContinueContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunContinueContentResponseBody) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBody) SetEnd(v bool) *RunContinueContentResponseBody {
	s.End = &v
	return s
}

func (s *RunContinueContentResponseBody) SetHeader(v *RunContinueContentResponseBodyHeader) *RunContinueContentResponseBody {
	s.Header = v
	return s
}

func (s *RunContinueContentResponseBody) SetPayload(v *RunContinueContentResponseBodyPayload) *RunContinueContentResponseBody {
	s.Payload = v
	return s
}

func (s *RunContinueContentResponseBody) SetRequestId(v string) *RunContinueContentResponseBody {
	s.RequestId = &v
	return s
}

type RunContinueContentResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunContinueContentResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunContinueContentResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyHeader) SetErrorCode(v string) *RunContinueContentResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetErrorMessage(v string) *RunContinueContentResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetEvent(v string) *RunContinueContentResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetEventInfo(v string) *RunContinueContentResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetSessionId(v string) *RunContinueContentResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetTaskId(v string) *RunContinueContentResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetTraceId(v string) *RunContinueContentResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunContinueContentResponseBodyPayload struct {
	Output *RunContinueContentResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunContinueContentResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunContinueContentResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunContinueContentResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyPayload) SetOutput(v *RunContinueContentResponseBodyPayloadOutput) *RunContinueContentResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunContinueContentResponseBodyPayload) SetUsage(v *RunContinueContentResponseBodyPayloadUsage) *RunContinueContentResponseBodyPayload {
	s.Usage = v
	return s
}

type RunContinueContentResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunContinueContentResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunContinueContentResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyPayloadOutput) SetText(v string) *RunContinueContentResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunContinueContentResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunContinueContentResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunContinueContentResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyPayloadUsage) SetInputTokens(v int64) *RunContinueContentResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunContinueContentResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunContinueContentResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunContinueContentResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunContinueContentResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunContinueContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunContinueContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContinueContentResponse) String() string {
	return tea.Prettify(s)
}

func (s RunContinueContentResponse) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponse) SetHeaders(v map[string]*string) *RunContinueContentResponse {
	s.Headers = v
	return s
}

func (s *RunContinueContentResponse) SetStatusCode(v int32) *RunContinueContentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContinueContentResponse) SetBody(v *RunContinueContentResponseBody) *RunContinueContentResponse {
	s.Body = v
	return s
}

type RunCustomHotTopicAnalysisRequest struct {
	// example:
	//
	// 模型反问
	AskUser *string `json:"AskUser,omitempty" xml:"AskUser,omitempty"`
	// example:
	//
	// false
	ForceAnalysisExistsTopic *bool `json:"ForceAnalysisExistsTopic,omitempty" xml:"ForceAnalysisExistsTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 用户输入Prompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
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
	// 用户针对模型反问的输入
	UserBack *string `json:"UserBack,omitempty" xml:"UserBack,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunCustomHotTopicAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisRequest) SetAskUser(v string) *RunCustomHotTopicAnalysisRequest {
	s.AskUser = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetForceAnalysisExistsTopic(v bool) *RunCustomHotTopicAnalysisRequest {
	s.ForceAnalysisExistsTopic = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetPrompt(v string) *RunCustomHotTopicAnalysisRequest {
	s.Prompt = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetSessionId(v string) *RunCustomHotTopicAnalysisRequest {
	s.SessionId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetTaskId(v string) *RunCustomHotTopicAnalysisRequest {
	s.TaskId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetUserBack(v string) *RunCustomHotTopicAnalysisRequest {
	s.UserBack = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetWorkspaceId(v string) *RunCustomHotTopicAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

type RunCustomHotTopicAnalysisResponseBody struct {
	Header  *RunCustomHotTopicAnalysisResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunCustomHotTopicAnalysisResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBody) SetHeader(v *RunCustomHotTopicAnalysisResponseBodyHeader) *RunCustomHotTopicAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBody) SetPayload(v *RunCustomHotTopicAnalysisResponseBodyPayload) *RunCustomHotTopicAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBody) SetRequestId(v string) *RunCustomHotTopicAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type RunCustomHotTopicAnalysisResponseBodyHeader struct {
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

func (s RunCustomHotTopicAnalysisResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetErrorCode(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetEvent(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetOriginSessionId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetSessionId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetTaskId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetTraceId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunCustomHotTopicAnalysisResponseBodyPayload struct {
	Output *RunCustomHotTopicAnalysisResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunCustomHotTopicAnalysisResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunCustomHotTopicAnalysisResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayload) SetOutput(v *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) *RunCustomHotTopicAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayload) SetUsage(v *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) *RunCustomHotTopicAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

type RunCustomHotTopicAnalysisResponseBodyPayloadOutput struct {
	Articles []*RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	AskUser  []*string                                                     `json:"AskUser,omitempty" xml:"AskUser,omitempty" type:"Repeated"`
	// example:
	//
	// 异步任务ID
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 自定义选题视角
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 大模型改变世界
	SearchQuery *string `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 话题ID
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetArticles(v []*RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetAskUser(v []*string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.AskUser = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetAsyncTaskId(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.AsyncTaskId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetAttitude(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.Attitude = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetSearchQuery(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetText(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetTopicId(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.TopicId = &v
	return s
}

type RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles struct {
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
	// a2103fcfbd5441f1991c72f8834833e3
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-08-27 14:50:47
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

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetAuthor(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Author = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetContent(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetDocId(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.DocId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetDocUuid(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.DocUuid = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetSource(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetSummary(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetTag(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Tag = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetTitle(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetUrl(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

type RunCustomHotTopicAnalysisResponseBodyPayloadUsage struct {
	// example:
	//
	// 60
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 13
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 73
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunCustomHotTopicAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunCustomHotTopicAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunCustomHotTopicAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunCustomHotTopicAnalysisResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCustomHotTopicAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponse) SetHeaders(v map[string]*string) *RunCustomHotTopicAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponse) SetStatusCode(v int32) *RunCustomHotTopicAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponse) SetBody(v *RunCustomHotTopicAnalysisResponseBody) *RunCustomHotTopicAnalysisResponse {
	s.Body = v
	return s
}

type RunCustomHotTopicViewPointAnalysisRequest struct {
	// example:
	//
	// 模型反问
	AskUser *string `json:"AskUser,omitempty" xml:"AskUser,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 自定义选题视角的Prompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 改写后的Query
	SearchQuery *string `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	// example:
	//
	// true
	SkipAskUser *bool `json:"SkipAskUser,omitempty" xml:"SkipAskUser,omitempty"`
	// example:
	//
	// 热点主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 热点主题ID
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// example:
	//
	// 热点主题来源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 热点主题版本
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	// example:
	//
	// 用户反馈
	UserBack *string `json:"UserBack,omitempty" xml:"UserBack,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetAskUser(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.AskUser = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetPrompt(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.Prompt = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetSearchQuery(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.SearchQuery = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetSkipAskUser(v bool) *RunCustomHotTopicViewPointAnalysisRequest {
	s.SkipAskUser = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopic(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.Topic = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopicId(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.TopicId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopicSource(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.TopicSource = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetTopicVersion(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.TopicVersion = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetUserBack(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.UserBack = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisRequest) SetWorkspaceId(v string) *RunCustomHotTopicViewPointAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

type RunCustomHotTopicViewPointAnalysisResponseBody struct {
	Header  *RunCustomHotTopicViewPointAnalysisResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunCustomHotTopicViewPointAnalysisResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) SetHeader(v *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) *RunCustomHotTopicViewPointAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) SetPayload(v *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) *RunCustomHotTopicViewPointAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) SetRequestId(v string) *RunCustomHotTopicViewPointAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type RunCustomHotTopicViewPointAnalysisResponseBodyHeader struct {
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

func (s RunCustomHotTopicViewPointAnalysisResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetErrorCode(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetEvent(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetOriginSessionId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetSessionId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetTaskId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetTraceId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunCustomHotTopicViewPointAnalysisResponseBodyPayload struct {
	Output *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) SetOutput(v *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) *RunCustomHotTopicViewPointAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) SetUsage(v *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) *RunCustomHotTopicViewPointAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

type RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput struct {
	AskUser []*string `json:"AskUser,omitempty" xml:"AskUser,omitempty" type:"Repeated"`
	// example:
	//
	// 异步任务ID
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 模型生成的自定义选题视角的观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// xxxxxx
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 话题ID
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetAskUser(v []*string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.AskUser = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetAsyncTaskId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.AsyncTaskId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetAttitude(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.Attitude = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetCustomViewPointId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.CustomViewPointId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetText(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetTopicId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.TopicId = &v
	return s
}

type RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage struct {
	// example:
	//
	// 51
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 79
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 130
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunCustomHotTopicViewPointAnalysisResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCustomHotTopicViewPointAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) SetHeaders(v map[string]*string) *RunCustomHotTopicViewPointAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) SetStatusCode(v int32) *RunCustomHotTopicViewPointAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponse) SetBody(v *RunCustomHotTopicViewPointAnalysisResponseBody) *RunCustomHotTopicViewPointAnalysisResponse {
	s.Body = v
	return s
}

type RunExpandContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 创新政务社交媒体功能。鼓励各地区、各部门结合实际，开发政务社交媒体的特色功能，如在线咨询服务、政策解读、互动问答等，增强政务社交媒体的互动性和实用性，提升公众参与度。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunExpandContentRequest) String() string {
	return tea.Prettify(s)
}

func (s RunExpandContentRequest) GoString() string {
	return s.String()
}

func (s *RunExpandContentRequest) SetContent(v string) *RunExpandContentRequest {
	s.Content = &v
	return s
}

func (s *RunExpandContentRequest) SetWorkspaceId(v string) *RunExpandContentRequest {
	s.WorkspaceId = &v
	return s
}

type RunExpandContentResponseBody struct {
	End     *bool                                `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunExpandContentResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunExpandContentResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunExpandContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunExpandContentResponseBody) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBody) SetEnd(v bool) *RunExpandContentResponseBody {
	s.End = &v
	return s
}

func (s *RunExpandContentResponseBody) SetHeader(v *RunExpandContentResponseBodyHeader) *RunExpandContentResponseBody {
	s.Header = v
	return s
}

func (s *RunExpandContentResponseBody) SetPayload(v *RunExpandContentResponseBodyPayload) *RunExpandContentResponseBody {
	s.Payload = v
	return s
}

func (s *RunExpandContentResponseBody) SetRequestId(v string) *RunExpandContentResponseBody {
	s.RequestId = &v
	return s
}

type RunExpandContentResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunExpandContentResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunExpandContentResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyHeader) SetErrorCode(v string) *RunExpandContentResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetErrorMessage(v string) *RunExpandContentResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetEvent(v string) *RunExpandContentResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetEventInfo(v string) *RunExpandContentResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetSessionId(v string) *RunExpandContentResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetTaskId(v string) *RunExpandContentResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetTraceId(v string) *RunExpandContentResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunExpandContentResponseBodyPayload struct {
	Output *RunExpandContentResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunExpandContentResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunExpandContentResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunExpandContentResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyPayload) SetOutput(v *RunExpandContentResponseBodyPayloadOutput) *RunExpandContentResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunExpandContentResponseBodyPayload) SetUsage(v *RunExpandContentResponseBodyPayloadUsage) *RunExpandContentResponseBodyPayload {
	s.Usage = v
	return s
}

type RunExpandContentResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunExpandContentResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunExpandContentResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyPayloadOutput) SetText(v string) *RunExpandContentResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunExpandContentResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunExpandContentResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunExpandContentResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyPayloadUsage) SetInputTokens(v int64) *RunExpandContentResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunExpandContentResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunExpandContentResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunExpandContentResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunExpandContentResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunExpandContentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunExpandContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunExpandContentResponse) String() string {
	return tea.Prettify(s)
}

func (s RunExpandContentResponse) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponse) SetHeaders(v map[string]*string) *RunExpandContentResponse {
	s.Headers = v
	return s
}

func (s *RunExpandContentResponse) SetStatusCode(v int32) *RunExpandContentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunExpandContentResponse) SetBody(v *RunExpandContentResponseBody) *RunExpandContentResponse {
	s.Body = v
	return s
}

type RunKeywordsExtractionGenerationRequest struct {
	// This parameter is required.
	ReferenceData *RunKeywordsExtractionGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunKeywordsExtractionGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationRequest) SetReferenceData(v *RunKeywordsExtractionGenerationRequestReferenceData) *RunKeywordsExtractionGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunKeywordsExtractionGenerationRequest) SetTaskId(v string) *RunKeywordsExtractionGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationRequest) SetWorkspaceId(v string) *RunKeywordsExtractionGenerationRequest {
	s.WorkspaceId = &v
	return s
}

type RunKeywordsExtractionGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunKeywordsExtractionGenerationRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationRequestReferenceData) SetContents(v []*string) *RunKeywordsExtractionGenerationRequestReferenceData {
	s.Contents = v
	return s
}

type RunKeywordsExtractionGenerationShrinkRequest struct {
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunKeywordsExtractionGenerationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunKeywordsExtractionGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) SetTaskId(v string) *RunKeywordsExtractionGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) SetWorkspaceId(v string) *RunKeywordsExtractionGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type RunKeywordsExtractionGenerationResponseBody struct {
	Header  *RunKeywordsExtractionGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunKeywordsExtractionGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 419F3FBE-5C8D-5949-AC29-E9615235D15A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBody) SetHeader(v *RunKeywordsExtractionGenerationResponseBodyHeader) *RunKeywordsExtractionGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBody) SetPayload(v *RunKeywordsExtractionGenerationResponseBodyPayload) *RunKeywordsExtractionGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBody) SetRequestId(v string) *RunKeywordsExtractionGenerationResponseBody {
	s.RequestId = &v
	return s
}

type RunKeywordsExtractionGenerationResponseBodyHeader struct {
	// example:
	//
	// AccessForbid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1a3d7c9f-3a6d-4e49-b176-2d8721a27397
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 8d55b429d7c6d321fcff54823e8d317b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 210bc4e817219607963985396de8bd
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetErrorCode(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetErrorMessage(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetEvent(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetSessionId(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetTaskId(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetTraceId(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunKeywordsExtractionGenerationResponseBodyPayload struct {
	Output *RunKeywordsExtractionGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunKeywordsExtractionGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunKeywordsExtractionGenerationResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayload) SetOutput(v *RunKeywordsExtractionGenerationResponseBodyPayloadOutput) *RunKeywordsExtractionGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayload) SetUsage(v *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) *RunKeywordsExtractionGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

type RunKeywordsExtractionGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadOutput) SetText(v string) *RunKeywordsExtractionGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunKeywordsExtractionGenerationResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunKeywordsExtractionGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunKeywordsExtractionGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunKeywordsExtractionGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunKeywordsExtractionGenerationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunKeywordsExtractionGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponse) SetHeaders(v map[string]*string) *RunKeywordsExtractionGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponse) SetStatusCode(v int32) *RunKeywordsExtractionGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponse) SetBody(v *RunKeywordsExtractionGenerationResponseBody) *RunKeywordsExtractionGenerationResponse {
	s.Body = v
	return s
}

type RunStepByStepWritingRequest struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt        *string                                   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	ReferenceData *RunStepByStepWritingRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId   *string                                   `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingConfig *RunStepByStepWritingRequestWritingConfig `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty" type:"Struct"`
}

func (s RunStepByStepWritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequest) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequest) SetOriginSessionId(v string) *RunStepByStepWritingRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetPrompt(v string) *RunStepByStepWritingRequest {
	s.Prompt = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetReferenceData(v *RunStepByStepWritingRequestReferenceData) *RunStepByStepWritingRequest {
	s.ReferenceData = v
	return s
}

func (s *RunStepByStepWritingRequest) SetSessionId(v string) *RunStepByStepWritingRequest {
	s.SessionId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetTaskId(v string) *RunStepByStepWritingRequest {
	s.TaskId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetWorkspaceId(v string) *RunStepByStepWritingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunStepByStepWritingRequest) SetWritingConfig(v *RunStepByStepWritingRequestWritingConfig) *RunStepByStepWritingRequest {
	s.WritingConfig = v
	return s
}

type RunStepByStepWritingRequestReferenceData struct {
	Articles      []*RunStepByStepWritingRequestReferenceDataArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	MiniDoc       []*string                                           `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
	Outlines      []*RunStepByStepWritingRequestReferenceDataOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	Summarization []*string                                           `json:"Summarization,omitempty" xml:"Summarization,omitempty" type:"Repeated"`
}

func (s RunStepByStepWritingRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceData) SetArticles(v []*RunStepByStepWritingRequestReferenceDataArticles) *RunStepByStepWritingRequestReferenceData {
	s.Articles = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceData) SetMiniDoc(v []*string) *RunStepByStepWritingRequestReferenceData {
	s.MiniDoc = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceData) SetOutlines(v []*RunStepByStepWritingRequestReferenceDataOutlines) *RunStepByStepWritingRequestReferenceData {
	s.Outlines = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceData) SetSummarization(v []*string) *RunStepByStepWritingRequestReferenceData {
	s.Summarization = v
	return s
}

type RunStepByStepWritingRequestReferenceDataArticles struct {
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
	// 8a20e007a6174522af4d6a2657d5526f
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// http://www.example.com
	MediaUrl *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	// example:
	//
	// 2024-09-10 14:17:54
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

func (s RunStepByStepWritingRequestReferenceDataArticles) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceDataArticles) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetAuthor(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Author = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetContent(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Content = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetDocId(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.DocId = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetDocUuid(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.DocUuid = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetMediaUrl(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.MediaUrl = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetPubTime(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.PubTime = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetSource(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Source = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetSummary(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Summary = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetTag(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Tag = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetTitle(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Title = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataArticles) SetUrl(v string) *RunStepByStepWritingRequestReferenceDataArticles {
	s.Url = &v
	return s
}

type RunStepByStepWritingRequestReferenceDataOutlines struct {
	Articles []*RunStepByStepWritingRequestReferenceDataOutlinesArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// example:
	//
	// 大纲
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
}

func (s RunStepByStepWritingRequestReferenceDataOutlines) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceDataOutlines) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceDataOutlines) SetArticles(v []*RunStepByStepWritingRequestReferenceDataOutlinesArticles) *RunStepByStepWritingRequestReferenceDataOutlines {
	s.Articles = v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlines) SetOutline(v string) *RunStepByStepWritingRequestReferenceDataOutlines {
	s.Outline = &v
	return s
}

type RunStepByStepWritingRequestReferenceDataOutlinesArticles struct {
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 文章链接
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunStepByStepWritingRequestReferenceDataOutlinesArticles) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequestReferenceDataOutlinesArticles) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) SetContent(v string) *RunStepByStepWritingRequestReferenceDataOutlinesArticles {
	s.Content = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) SetTitle(v string) *RunStepByStepWritingRequestReferenceDataOutlinesArticles {
	s.Title = &v
	return s
}

func (s *RunStepByStepWritingRequestReferenceDataOutlinesArticles) SetUrl(v string) *RunStepByStepWritingRequestReferenceDataOutlinesArticles {
	s.Url = &v
	return s
}

type RunStepByStepWritingRequestWritingConfig struct {
	// example:
	//
	// media
	Domain    *string                                            `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Keywords  []*string                                          `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	PromptTag *RunStepByStepWritingRequestWritingConfigPromptTag `json:"PromptTag,omitempty" xml:"PromptTag,omitempty" type:"Struct"`
	// example:
	//
	// 分步骤写作场景，传媒写作支持的写作场景:新闻写作(默认),新闻评论,通用文体，公文写作支持的写作场景:通知(默认),通告,通报,请示,决定,函,通用文体
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// Writing
	Step *string                                         `json:"Step,omitempty" xml:"Step,omitempty"`
	Tags []*RunStepByStepWritingRequestWritingConfigTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
}

func (s RunStepByStepWritingRequestWritingConfig) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequestWritingConfig) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestWritingConfig) SetDomain(v string) *RunStepByStepWritingRequestWritingConfig {
	s.Domain = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetKeywords(v []*string) *RunStepByStepWritingRequestWritingConfig {
	s.Keywords = v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetPromptTag(v *RunStepByStepWritingRequestWritingConfigPromptTag) *RunStepByStepWritingRequestWritingConfig {
	s.PromptTag = v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetScene(v string) *RunStepByStepWritingRequestWritingConfig {
	s.Scene = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetStep(v string) *RunStepByStepWritingRequestWritingConfig {
	s.Step = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetTags(v []*RunStepByStepWritingRequestWritingConfigTags) *RunStepByStepWritingRequestWritingConfig {
	s.Tags = v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfig) SetUseSearch(v bool) *RunStepByStepWritingRequestWritingConfig {
	s.UseSearch = &v
	return s
}

type RunStepByStepWritingRequestWritingConfigPromptTag struct {
	// example:
	//
	// 必要提示
	NecessaryTips *string `json:"NecessaryTips,omitempty" xml:"NecessaryTips,omitempty"`
	// example:
	//
	// 立场
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// 反向词
	ReverseWords *string `json:"ReverseWords,omitempty" xml:"ReverseWords,omitempty"`
	// example:
	//
	// 主题
	Theme *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
}

func (s RunStepByStepWritingRequestWritingConfigPromptTag) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequestWritingConfigPromptTag) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetNecessaryTips(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.NecessaryTips = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetPosition(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.Position = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetReverseWords(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.ReverseWords = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigPromptTag) SetTheme(v string) *RunStepByStepWritingRequestWritingConfigPromptTag {
	s.Theme = &v
	return s
}

type RunStepByStepWritingRequestWritingConfigTags struct {
	// example:
	//
	// 10
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// gcNumberSizeTag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s RunStepByStepWritingRequestWritingConfigTags) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingRequestWritingConfigTags) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingRequestWritingConfigTags) SetKeyword(v string) *RunStepByStepWritingRequestWritingConfigTags {
	s.Keyword = &v
	return s
}

func (s *RunStepByStepWritingRequestWritingConfigTags) SetTag(v string) *RunStepByStepWritingRequestWritingConfigTags {
	s.Tag = &v
	return s
}

type RunStepByStepWritingShrinkRequest struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt              *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingConfigShrink *string `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty"`
}

func (s RunStepByStepWritingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingShrinkRequest) SetOriginSessionId(v string) *RunStepByStepWritingShrinkRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetPrompt(v string) *RunStepByStepWritingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetReferenceDataShrink(v string) *RunStepByStepWritingShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetSessionId(v string) *RunStepByStepWritingShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetTaskId(v string) *RunStepByStepWritingShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetWorkspaceId(v string) *RunStepByStepWritingShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetWritingConfigShrink(v string) *RunStepByStepWritingShrinkRequest {
	s.WritingConfigShrink = &v
	return s
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
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyHeader) GoString() string {
	return s.String()
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

type RunStepByStepWritingResponseBodyPayload struct {
	Output *RunStepByStepWritingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunStepByStepWritingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunStepByStepWritingResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyPayload) SetOutput(v *RunStepByStepWritingResponseBodyPayloadOutput) *RunStepByStepWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunStepByStepWritingResponseBodyPayload) SetUsage(v *RunStepByStepWritingResponseBodyPayloadUsage) *RunStepByStepWritingResponseBodyPayload {
	s.Usage = v
	return s
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
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
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

type RunStepByStepWritingResponseBodyPayloadOutputExtraOutput struct {
	Summarization []*string `json:"summarization,omitempty" xml:"summarization,omitempty" type:"Repeated"`
}

func (s RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput) SetSummarization(v []*string) *RunStepByStepWritingResponseBodyPayloadOutputExtraOutput {
	s.Summarization = v
	return s
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
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
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

type RunStepByStepWritingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunStepByStepWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunStepByStepWritingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunStepByStepWritingResponse) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingResponse) SetHeaders(v map[string]*string) *RunStepByStepWritingResponse {
	s.Headers = v
	return s
}

func (s *RunStepByStepWritingResponse) SetStatusCode(v int32) *RunStepByStepWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunStepByStepWritingResponse) SetBody(v *RunStepByStepWritingResponseBody) *RunStepByStepWritingResponse {
	s.Body = v
	return s
}

type RunStyleFeatureAnalysisRequest struct {
	Contents    []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	MaterialIds []*int64  `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunStyleFeatureAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisRequest) SetContents(v []*string) *RunStyleFeatureAnalysisRequest {
	s.Contents = v
	return s
}

func (s *RunStyleFeatureAnalysisRequest) SetMaterialIds(v []*int64) *RunStyleFeatureAnalysisRequest {
	s.MaterialIds = v
	return s
}

func (s *RunStyleFeatureAnalysisRequest) SetWorkspaceId(v string) *RunStyleFeatureAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

type RunStyleFeatureAnalysisShrinkRequest struct {
	ContentsShrink    *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	MaterialIdsShrink *string `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunStyleFeatureAnalysisShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisShrinkRequest) SetContentsShrink(v string) *RunStyleFeatureAnalysisShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *RunStyleFeatureAnalysisShrinkRequest) SetMaterialIdsShrink(v string) *RunStyleFeatureAnalysisShrinkRequest {
	s.MaterialIdsShrink = &v
	return s
}

func (s *RunStyleFeatureAnalysisShrinkRequest) SetWorkspaceId(v string) *RunStyleFeatureAnalysisShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type RunStyleFeatureAnalysisResponseBody struct {
	End     *bool                                       `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunStyleFeatureAnalysisResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunStyleFeatureAnalysisResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunStyleFeatureAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBody) SetEnd(v bool) *RunStyleFeatureAnalysisResponseBody {
	s.End = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBody) SetHeader(v *RunStyleFeatureAnalysisResponseBodyHeader) *RunStyleFeatureAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBody) SetPayload(v *RunStyleFeatureAnalysisResponseBodyPayload) *RunStyleFeatureAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBody) SetRequestId(v string) *RunStyleFeatureAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type RunStyleFeatureAnalysisResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunStyleFeatureAnalysisResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetErrorCode(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetEvent(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetEventInfo(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetSessionId(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetTaskId(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetTraceId(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunStyleFeatureAnalysisResponseBodyPayload struct {
	Output *RunStyleFeatureAnalysisResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunStyleFeatureAnalysisResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunStyleFeatureAnalysisResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyPayload) SetOutput(v *RunStyleFeatureAnalysisResponseBodyPayloadOutput) *RunStyleFeatureAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayload) SetUsage(v *RunStyleFeatureAnalysisResponseBodyPayloadUsage) *RunStyleFeatureAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

type RunStyleFeatureAnalysisResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunStyleFeatureAnalysisResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadOutput) SetText(v string) *RunStyleFeatureAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunStyleFeatureAnalysisResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunStyleFeatureAnalysisResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunStyleFeatureAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunStyleFeatureAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunStyleFeatureAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunStyleFeatureAnalysisResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunStyleFeatureAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunStyleFeatureAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponse) SetHeaders(v map[string]*string) *RunStyleFeatureAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetStatusCode(v int32) *RunStyleFeatureAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponse) SetBody(v *RunStyleFeatureAnalysisResponseBody) *RunStyleFeatureAnalysisResponse {
	s.Body = v
	return s
}

type RunSummaryGenerateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 创新政务社交媒体功能。鼓励各地区、各部门结合实际，开发政务社交媒体的特色功能，如在线咨询服务、政策解读、互动问答等，增强政务社交媒体的互动性和实用性，提升公众参与度。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 请为上述内容生成一段摘要，字数在100~200字以内。
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunSummaryGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSummaryGenerateRequest) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateRequest) SetContent(v string) *RunSummaryGenerateRequest {
	s.Content = &v
	return s
}

func (s *RunSummaryGenerateRequest) SetPrompt(v string) *RunSummaryGenerateRequest {
	s.Prompt = &v
	return s
}

func (s *RunSummaryGenerateRequest) SetWorkspaceId(v string) *RunSummaryGenerateRequest {
	s.WorkspaceId = &v
	return s
}

type RunSummaryGenerateResponseBody struct {
	End     *bool                                  `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunSummaryGenerateResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunSummaryGenerateResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunSummaryGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunSummaryGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBody) SetEnd(v bool) *RunSummaryGenerateResponseBody {
	s.End = &v
	return s
}

func (s *RunSummaryGenerateResponseBody) SetHeader(v *RunSummaryGenerateResponseBodyHeader) *RunSummaryGenerateResponseBody {
	s.Header = v
	return s
}

func (s *RunSummaryGenerateResponseBody) SetPayload(v *RunSummaryGenerateResponseBodyPayload) *RunSummaryGenerateResponseBody {
	s.Payload = v
	return s
}

func (s *RunSummaryGenerateResponseBody) SetRequestId(v string) *RunSummaryGenerateResponseBody {
	s.RequestId = &v
	return s
}

type RunSummaryGenerateResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunSummaryGenerateResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyHeader) SetErrorCode(v string) *RunSummaryGenerateResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetErrorMessage(v string) *RunSummaryGenerateResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetEvent(v string) *RunSummaryGenerateResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetEventInfo(v string) *RunSummaryGenerateResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetSessionId(v string) *RunSummaryGenerateResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetTaskId(v string) *RunSummaryGenerateResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetTraceId(v string) *RunSummaryGenerateResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunSummaryGenerateResponseBodyPayload struct {
	Output *RunSummaryGenerateResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunSummaryGenerateResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunSummaryGenerateResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyPayload) SetOutput(v *RunSummaryGenerateResponseBodyPayloadOutput) *RunSummaryGenerateResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayload) SetUsage(v *RunSummaryGenerateResponseBodyPayloadUsage) *RunSummaryGenerateResponseBodyPayload {
	s.Usage = v
	return s
}

type RunSummaryGenerateResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunSummaryGenerateResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyPayloadOutput) SetText(v string) *RunSummaryGenerateResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunSummaryGenerateResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunSummaryGenerateResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) SetInputTokens(v int64) *RunSummaryGenerateResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunSummaryGenerateResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunSummaryGenerateResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunSummaryGenerateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSummaryGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSummaryGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s RunSummaryGenerateResponse) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponse) SetHeaders(v map[string]*string) *RunSummaryGenerateResponse {
	s.Headers = v
	return s
}

func (s *RunSummaryGenerateResponse) SetStatusCode(v int32) *RunSummaryGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSummaryGenerateResponse) SetBody(v *RunSummaryGenerateResponseBody) *RunSummaryGenerateResponse {
	s.Body = v
	return s
}

type RunTextPolishingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 文本内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTextPolishingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunTextPolishingRequest) GoString() string {
	return s.String()
}

func (s *RunTextPolishingRequest) SetContent(v string) *RunTextPolishingRequest {
	s.Content = &v
	return s
}

func (s *RunTextPolishingRequest) SetWorkspaceId(v string) *RunTextPolishingRequest {
	s.WorkspaceId = &v
	return s
}

type RunTextPolishingResponseBody struct {
	Header  *RunTextPolishingResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunTextPolishingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunTextPolishingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunTextPolishingResponseBody) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBody) SetHeader(v *RunTextPolishingResponseBodyHeader) *RunTextPolishingResponseBody {
	s.Header = v
	return s
}

func (s *RunTextPolishingResponseBody) SetPayload(v *RunTextPolishingResponseBodyPayload) *RunTextPolishingResponseBody {
	s.Payload = v
	return s
}

func (s *RunTextPolishingResponseBody) SetRequestId(v string) *RunTextPolishingResponseBody {
	s.RequestId = &v
	return s
}

type RunTextPolishingResponseBodyHeader struct {
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

func (s RunTextPolishingResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunTextPolishingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyHeader) SetErrorCode(v string) *RunTextPolishingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetErrorMessage(v string) *RunTextPolishingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetEvent(v string) *RunTextPolishingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetOriginSessionId(v string) *RunTextPolishingResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetSessionId(v string) *RunTextPolishingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetTaskId(v string) *RunTextPolishingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetTraceId(v string) *RunTextPolishingResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunTextPolishingResponseBodyPayload struct {
	Output *RunTextPolishingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunTextPolishingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunTextPolishingResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunTextPolishingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyPayload) SetOutput(v *RunTextPolishingResponseBodyPayloadOutput) *RunTextPolishingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTextPolishingResponseBodyPayload) SetUsage(v *RunTextPolishingResponseBodyPayloadUsage) *RunTextPolishingResponseBodyPayload {
	s.Usage = v
	return s
}

type RunTextPolishingResponseBodyPayloadOutput struct {
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunTextPolishingResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunTextPolishingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyPayloadOutput) SetText(v string) *RunTextPolishingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunTextPolishingResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunTextPolishingResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunTextPolishingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTextPolishingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTextPolishingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTextPolishingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTextPolishingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTextPolishingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunTextPolishingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTextPolishingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTextPolishingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunTextPolishingResponse) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponse) SetHeaders(v map[string]*string) *RunTextPolishingResponse {
	s.Headers = v
	return s
}

func (s *RunTextPolishingResponse) SetStatusCode(v int32) *RunTextPolishingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTextPolishingResponse) SetBody(v *RunTextPolishingResponseBody) *RunTextPolishingResponse {
	s.Body = v
	return s
}

type RunTitleGenerationRequest struct {
	// This parameter is required.
	ReferenceData *RunTitleGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTitleGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationRequest) SetReferenceData(v *RunTitleGenerationRequestReferenceData) *RunTitleGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunTitleGenerationRequest) SetTaskId(v string) *RunTitleGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunTitleGenerationRequest) SetWorkspaceId(v string) *RunTitleGenerationRequest {
	s.WorkspaceId = &v
	return s
}

type RunTitleGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunTitleGenerationRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationRequestReferenceData) SetContents(v []*string) *RunTitleGenerationRequestReferenceData {
	s.Contents = v
	return s
}

type RunTitleGenerationShrinkRequest struct {
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTitleGenerationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunTitleGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunTitleGenerationShrinkRequest) SetTaskId(v string) *RunTitleGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunTitleGenerationShrinkRequest) SetWorkspaceId(v string) *RunTitleGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type RunTitleGenerationResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Header         *RunTitleGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	HttpStatusCode *string                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Payload        *RunTitleGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunTitleGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBody) SetCode(v string) *RunTitleGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetHeader(v *RunTitleGenerationResponseBodyHeader) *RunTitleGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunTitleGenerationResponseBody) SetHttpStatusCode(v string) *RunTitleGenerationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetMessage(v string) *RunTitleGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetPayload(v *RunTitleGenerationResponseBodyPayload) *RunTitleGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunTitleGenerationResponseBody) SetRequestId(v string) *RunTitleGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetSuccess(v bool) *RunTitleGenerationResponseBody {
	s.Success = &v
	return s
}

type RunTitleGenerationResponseBodyHeader struct {
	// example:
	//
	// AccessForbid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// 50a1cc8e-717e-4a2b-a76b-dc9734a8564b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0a3d448f17000139741898287e0eb3
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunTitleGenerationResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyHeader) SetErrorCode(v string) *RunTitleGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetErrorMessage(v string) *RunTitleGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetEvent(v string) *RunTitleGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetSessionId(v string) *RunTitleGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetStatusCode(v int32) *RunTitleGenerationResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetTaskId(v string) *RunTitleGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetTraceId(v string) *RunTitleGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunTitleGenerationResponseBodyPayload struct {
	Output *RunTitleGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunTitleGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunTitleGenerationResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyPayload) SetOutput(v *RunTitleGenerationResponseBodyPayloadOutput) *RunTitleGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTitleGenerationResponseBodyPayload) SetUsage(v *RunTitleGenerationResponseBodyPayloadUsage) *RunTitleGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

type RunTitleGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunTitleGenerationResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyPayloadOutput) SetText(v string) *RunTitleGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunTitleGenerationResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunTitleGenerationResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTitleGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTitleGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTitleGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunTitleGenerationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTitleGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTitleGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunTitleGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponse) SetHeaders(v map[string]*string) *RunTitleGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunTitleGenerationResponse) SetStatusCode(v int32) *RunTitleGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTitleGenerationResponse) SetBody(v *RunTitleGenerationResponseBody) *RunTitleGenerationResponse {
	s.Body = v
	return s
}

type RunTranslateGenerationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// toEnglish
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceData *RunTranslateGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTranslateGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationRequest) SetPrompt(v string) *RunTranslateGenerationRequest {
	s.Prompt = &v
	return s
}

func (s *RunTranslateGenerationRequest) SetReferenceData(v *RunTranslateGenerationRequestReferenceData) *RunTranslateGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunTranslateGenerationRequest) SetTaskId(v string) *RunTranslateGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunTranslateGenerationRequest) SetWorkspaceId(v string) *RunTranslateGenerationRequest {
	s.WorkspaceId = &v
	return s
}

type RunTranslateGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunTranslateGenerationRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationRequestReferenceData) SetContents(v []*string) *RunTranslateGenerationRequestReferenceData {
	s.Contents = v
	return s
}

type RunTranslateGenerationShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// toEnglish
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTranslateGenerationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationShrinkRequest) SetPrompt(v string) *RunTranslateGenerationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunTranslateGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunTranslateGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunTranslateGenerationShrinkRequest) SetTaskId(v string) *RunTranslateGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunTranslateGenerationShrinkRequest) SetWorkspaceId(v string) *RunTranslateGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type RunTranslateGenerationResponseBody struct {
	Header  *RunTranslateGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunTranslateGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunTranslateGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBody) SetHeader(v *RunTranslateGenerationResponseBodyHeader) *RunTranslateGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunTranslateGenerationResponseBody) SetPayload(v *RunTranslateGenerationResponseBodyPayload) *RunTranslateGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunTranslateGenerationResponseBody) SetRequestId(v string) *RunTranslateGenerationResponseBody {
	s.RequestId = &v
	return s
}

type RunTranslateGenerationResponseBodyHeader struct {
	// example:
	//
	// AccessForbid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 91C2B2B8-7D12-4A8D-A724-1E576D30C096
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0abb781d17146157564845243e20b5
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunTranslateGenerationResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyHeader) SetErrorCode(v string) *RunTranslateGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetErrorMessage(v string) *RunTranslateGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetEvent(v string) *RunTranslateGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetSessionId(v string) *RunTranslateGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetTaskId(v string) *RunTranslateGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetTraceId(v string) *RunTranslateGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunTranslateGenerationResponseBodyPayload struct {
	Output *RunTranslateGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunTranslateGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunTranslateGenerationResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyPayload) SetOutput(v *RunTranslateGenerationResponseBodyPayloadOutput) *RunTranslateGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayload) SetUsage(v *RunTranslateGenerationResponseBodyPayloadUsage) *RunTranslateGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

type RunTranslateGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunTranslateGenerationResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyPayloadOutput) SetText(v string) *RunTranslateGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunTranslateGenerationResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunTranslateGenerationResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTranslateGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTranslateGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTranslateGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunTranslateGenerationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTranslateGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTranslateGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunTranslateGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponse) SetHeaders(v map[string]*string) *RunTranslateGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunTranslateGenerationResponse) SetStatusCode(v int32) *RunTranslateGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTranslateGenerationResponse) SetBody(v *RunTranslateGenerationResponseBody) *RunTranslateGenerationResponse {
	s.Body = v
	return s
}

type RunWriteToneGenerationRequest struct {
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceData *RunWriteToneGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunWriteToneGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationRequest) SetPrompt(v string) *RunWriteToneGenerationRequest {
	s.Prompt = &v
	return s
}

func (s *RunWriteToneGenerationRequest) SetReferenceData(v *RunWriteToneGenerationRequestReferenceData) *RunWriteToneGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunWriteToneGenerationRequest) SetTaskId(v string) *RunWriteToneGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunWriteToneGenerationRequest) SetWorkspaceId(v string) *RunWriteToneGenerationRequest {
	s.WorkspaceId = &v
	return s
}

type RunWriteToneGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunWriteToneGenerationRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationRequestReferenceData) SetContents(v []*string) *RunWriteToneGenerationRequestReferenceData {
	s.Contents = v
	return s
}

type RunWriteToneGenerationShrinkRequest struct {
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunWriteToneGenerationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationShrinkRequest) SetPrompt(v string) *RunWriteToneGenerationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunWriteToneGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunWriteToneGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunWriteToneGenerationShrinkRequest) SetTaskId(v string) *RunWriteToneGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunWriteToneGenerationShrinkRequest) SetWorkspaceId(v string) *RunWriteToneGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type RunWriteToneGenerationResponseBody struct {
	Header  *RunWriteToneGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunWriteToneGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// FB698445-61DA-5361-BF73-1C5F1157E888
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunWriteToneGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBody) SetHeader(v *RunWriteToneGenerationResponseBodyHeader) *RunWriteToneGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunWriteToneGenerationResponseBody) SetPayload(v *RunWriteToneGenerationResponseBodyPayload) *RunWriteToneGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunWriteToneGenerationResponseBody) SetRequestId(v string) *RunWriteToneGenerationResponseBody {
	s.RequestId = &v
	return s
}

type RunWriteToneGenerationResponseBodyHeader struct {
	// example:
	//
	// AccessForbid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// F1953EE6-157C-40DC-BBF1-87C98AC27C51
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// F1953EE6-157C-40DC-BBF1-87C98AC27C51
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// F1953EE6-157C-40DC-BBF1-87C98AC27C51
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunWriteToneGenerationResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetErrorCode(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetErrorMessage(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetEvent(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetSessionId(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetTaskId(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetTraceId(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunWriteToneGenerationResponseBodyPayload struct {
	Output *RunWriteToneGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunWriteToneGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunWriteToneGenerationResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyPayload) SetOutput(v *RunWriteToneGenerationResponseBodyPayloadOutput) *RunWriteToneGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayload) SetUsage(v *RunWriteToneGenerationResponseBodyPayloadUsage) *RunWriteToneGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

type RunWriteToneGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunWriteToneGenerationResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyPayloadOutput) SetText(v string) *RunWriteToneGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunWriteToneGenerationResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunWriteToneGenerationResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunWriteToneGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunWriteToneGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunWriteToneGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunWriteToneGenerationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunWriteToneGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunWriteToneGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunWriteToneGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponse) SetHeaders(v map[string]*string) *RunWriteToneGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunWriteToneGenerationResponse) SetStatusCode(v int32) *RunWriteToneGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunWriteToneGenerationResponse) SetBody(v *RunWriteToneGenerationResponseBody) *RunWriteToneGenerationResponse {
	s.Body = v
	return s
}

type RunWritingRequest struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt        *string                         `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	ReferenceData *RunWritingRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId   *string                         `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingConfig *RunWritingRequestWritingConfig `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty" type:"Struct"`
}

func (s RunWritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunWritingRequest) GoString() string {
	return s.String()
}

func (s *RunWritingRequest) SetOriginSessionId(v string) *RunWritingRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunWritingRequest) SetPrompt(v string) *RunWritingRequest {
	s.Prompt = &v
	return s
}

func (s *RunWritingRequest) SetReferenceData(v *RunWritingRequestReferenceData) *RunWritingRequest {
	s.ReferenceData = v
	return s
}

func (s *RunWritingRequest) SetSessionId(v string) *RunWritingRequest {
	s.SessionId = &v
	return s
}

func (s *RunWritingRequest) SetTaskId(v string) *RunWritingRequest {
	s.TaskId = &v
	return s
}

func (s *RunWritingRequest) SetWorkspaceId(v string) *RunWritingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunWritingRequest) SetWritingConfig(v *RunWritingRequestWritingConfig) *RunWritingRequest {
	s.WritingConfig = v
	return s
}

type RunWritingRequestReferenceData struct {
	Articles []*RunWritingRequestReferenceDataArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
}

func (s RunWritingRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s RunWritingRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunWritingRequestReferenceData) SetArticles(v []*RunWritingRequestReferenceDataArticles) *RunWritingRequestReferenceData {
	s.Articles = v
	return s
}

type RunWritingRequestReferenceDataArticles struct {
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
	// 2124ca4d48a542d788aa86151e1a8c8b
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-08-28 11:38:28
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

func (s RunWritingRequestReferenceDataArticles) String() string {
	return tea.Prettify(s)
}

func (s RunWritingRequestReferenceDataArticles) GoString() string {
	return s.String()
}

func (s *RunWritingRequestReferenceDataArticles) SetAuthor(v string) *RunWritingRequestReferenceDataArticles {
	s.Author = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetContent(v string) *RunWritingRequestReferenceDataArticles {
	s.Content = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetDocId(v string) *RunWritingRequestReferenceDataArticles {
	s.DocId = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetDocUuid(v string) *RunWritingRequestReferenceDataArticles {
	s.DocUuid = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetPubTime(v string) *RunWritingRequestReferenceDataArticles {
	s.PubTime = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetSource(v string) *RunWritingRequestReferenceDataArticles {
	s.Source = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetSummary(v string) *RunWritingRequestReferenceDataArticles {
	s.Summary = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetTag(v string) *RunWritingRequestReferenceDataArticles {
	s.Tag = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetTitle(v string) *RunWritingRequestReferenceDataArticles {
	s.Title = &v
	return s
}

func (s *RunWritingRequestReferenceDataArticles) SetUrl(v string) *RunWritingRequestReferenceDataArticles {
	s.Url = &v
	return s
}

type RunWritingRequestWritingConfig struct {
	// example:
	//
	// 写作领域，media:传媒,government:政务,market:营销
	Domain    *string                                  `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PromptTag *RunWritingRequestWritingConfigPromptTag `json:"PromptTag,omitempty" xml:"PromptTag,omitempty" type:"Struct"`
	Tags      []*RunWritingRequestWritingConfigTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
}

func (s RunWritingRequestWritingConfig) String() string {
	return tea.Prettify(s)
}

func (s RunWritingRequestWritingConfig) GoString() string {
	return s.String()
}

func (s *RunWritingRequestWritingConfig) SetDomain(v string) *RunWritingRequestWritingConfig {
	s.Domain = &v
	return s
}

func (s *RunWritingRequestWritingConfig) SetPromptTag(v *RunWritingRequestWritingConfigPromptTag) *RunWritingRequestWritingConfig {
	s.PromptTag = v
	return s
}

func (s *RunWritingRequestWritingConfig) SetTags(v []*RunWritingRequestWritingConfigTags) *RunWritingRequestWritingConfig {
	s.Tags = v
	return s
}

func (s *RunWritingRequestWritingConfig) SetUseSearch(v bool) *RunWritingRequestWritingConfig {
	s.UseSearch = &v
	return s
}

type RunWritingRequestWritingConfigPromptTag struct {
	// example:
	//
	// 必要提示
	NecessaryTips *string `json:"NecessaryTips,omitempty" xml:"NecessaryTips,omitempty"`
	// example:
	//
	// 立场
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// 反向词
	ReverseWords *string `json:"ReverseWords,omitempty" xml:"ReverseWords,omitempty"`
	// example:
	//
	// 主题
	Theme *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
}

func (s RunWritingRequestWritingConfigPromptTag) String() string {
	return tea.Prettify(s)
}

func (s RunWritingRequestWritingConfigPromptTag) GoString() string {
	return s.String()
}

func (s *RunWritingRequestWritingConfigPromptTag) SetNecessaryTips(v string) *RunWritingRequestWritingConfigPromptTag {
	s.NecessaryTips = &v
	return s
}

func (s *RunWritingRequestWritingConfigPromptTag) SetPosition(v string) *RunWritingRequestWritingConfigPromptTag {
	s.Position = &v
	return s
}

func (s *RunWritingRequestWritingConfigPromptTag) SetReverseWords(v string) *RunWritingRequestWritingConfigPromptTag {
	s.ReverseWords = &v
	return s
}

func (s *RunWritingRequestWritingConfigPromptTag) SetTheme(v string) *RunWritingRequestWritingConfigPromptTag {
	s.Theme = &v
	return s
}

type RunWritingRequestWritingConfigTags struct {
	// example:
	//
	// 10
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// gcNumberSizeTag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s RunWritingRequestWritingConfigTags) String() string {
	return tea.Prettify(s)
}

func (s RunWritingRequestWritingConfigTags) GoString() string {
	return s.String()
}

func (s *RunWritingRequestWritingConfigTags) SetKeyword(v string) *RunWritingRequestWritingConfigTags {
	s.Keyword = &v
	return s
}

func (s *RunWritingRequestWritingConfigTags) SetTag(v string) *RunWritingRequestWritingConfigTags {
	s.Tag = &v
	return s
}

type RunWritingShrinkRequest struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt              *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingConfigShrink *string `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty"`
}

func (s RunWritingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunWritingShrinkRequest) SetOriginSessionId(v string) *RunWritingShrinkRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetPrompt(v string) *RunWritingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunWritingShrinkRequest) SetReferenceDataShrink(v string) *RunWritingShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunWritingShrinkRequest) SetSessionId(v string) *RunWritingShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetTaskId(v string) *RunWritingShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetWorkspaceId(v string) *RunWritingShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetWritingConfigShrink(v string) *RunWritingShrinkRequest {
	s.WritingConfigShrink = &v
	return s
}

type RunWritingResponseBody struct {
	Header  *RunWritingResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunWritingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunWritingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBody) SetHeader(v *RunWritingResponseBodyHeader) *RunWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunWritingResponseBody) SetPayload(v *RunWritingResponseBodyPayload) *RunWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunWritingResponseBody) SetRequestId(v string) *RunWritingResponseBody {
	s.RequestId = &v
	return s
}

type RunWritingResponseBodyHeader struct {
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
	// 400
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 全链路ID
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunWritingResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s RunWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyHeader) SetErrorCode(v string) *RunWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetErrorMessage(v string) *RunWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetEvent(v string) *RunWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetOriginSessionId(v string) *RunWritingResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetSessionId(v string) *RunWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetStatusCode(v int32) *RunWritingResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetTaskId(v string) *RunWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetTraceId(v string) *RunWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

type RunWritingResponseBodyPayload struct {
	Output *RunWritingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunWritingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunWritingResponseBodyPayload) String() string {
	return tea.Prettify(s)
}

func (s RunWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayload) SetOutput(v *RunWritingResponseBodyPayloadOutput) *RunWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunWritingResponseBodyPayload) SetUsage(v *RunWritingResponseBodyPayloadUsage) *RunWritingResponseBodyPayload {
	s.Usage = v
	return s
}

type RunWritingResponseBodyPayloadOutput struct {
	Articles []*RunWritingResponseBodyPayloadOutputArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
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

func (s RunWritingResponseBodyPayloadOutput) String() string {
	return tea.Prettify(s)
}

func (s RunWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayloadOutput) SetArticles(v []*RunWritingResponseBodyPayloadOutputArticles) *RunWritingResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunWritingResponseBodyPayloadOutput) SetMiniDoc(v []*string) *RunWritingResponseBodyPayloadOutput {
	s.MiniDoc = v
	return s
}

func (s *RunWritingResponseBodyPayloadOutput) SetSearchQuery(v string) *RunWritingResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutput) SetText(v string) *RunWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

type RunWritingResponseBodyPayloadOutputArticles struct {
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
	// 98229f6001cf4deeb1668191d4eccc75
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-08-28 11:38:28
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

func (s RunWritingResponseBodyPayloadOutputArticles) String() string {
	return tea.Prettify(s)
}

func (s RunWritingResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetAuthor(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Author = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetContent(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetDocId(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.DocId = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetDocUuid(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.DocUuid = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetSource(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetSummary(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetTag(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Tag = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetTitle(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetUrl(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

type RunWritingResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64            `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	TokenMap     map[string]*int64 `json:"TokenMap,omitempty" xml:"TokenMap,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunWritingResponseBodyPayloadUsage) String() string {
	return tea.Prettify(s)
}

func (s RunWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunWritingResponseBodyPayloadUsage) SetTokenMap(v map[string]*int64) *RunWritingResponseBodyPayloadUsage {
	s.TokenMap = v
	return s
}

func (s *RunWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

type RunWritingResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunWritingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunWritingResponse) GoString() string {
	return s.String()
}

func (s *RunWritingResponse) SetHeaders(v map[string]*string) *RunWritingResponse {
	s.Headers = v
	return s
}

func (s *RunWritingResponse) SetStatusCode(v int32) *RunWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunWritingResponse) SetBody(v *RunWritingResponseBody) *RunWritingResponse {
	s.Body = v
	return s
}

type SaveCustomTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 商品code
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SaveCustomTextRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomTextRequest) GoString() string {
	return s.String()
}

func (s *SaveCustomTextRequest) SetAgentKey(v string) *SaveCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveCustomTextRequest) SetCommodityCode(v string) *SaveCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *SaveCustomTextRequest) SetContent(v string) *SaveCustomTextRequest {
	s.Content = &v
	return s
}

func (s *SaveCustomTextRequest) SetTitle(v string) *SaveCustomTextRequest {
	s.Title = &v
	return s
}

type SaveCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 5
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveCustomTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCustomTextResponseBody) SetCode(v string) *SaveCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetData(v int64) *SaveCustomTextResponseBody {
	s.Data = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetHttpStatusCode(v int32) *SaveCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetMessage(v string) *SaveCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetRequestId(v string) *SaveCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveCustomTextResponseBody) SetSuccess(v bool) *SaveCustomTextResponseBody {
	s.Success = &v
	return s
}

type SaveCustomTextResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveCustomTextResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveCustomTextResponse) GoString() string {
	return s.String()
}

func (s *SaveCustomTextResponse) SetHeaders(v map[string]*string) *SaveCustomTextResponse {
	s.Headers = v
	return s
}

func (s *SaveCustomTextResponse) SetStatusCode(v int32) *SaveCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveCustomTextResponse) SetBody(v *SaveCustomTextResponseBody) *SaveCustomTextResponse {
	s.Body = v
	return s
}

type SaveDataSourceOrderConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miaobi
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	UserConfigDataSourceList []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty" type:"Repeated"`
}

func (s SaveDataSourceOrderConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigRequest) SetAgentKey(v string) *SaveDataSourceOrderConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) SetProductCode(v string) *SaveDataSourceOrderConfigRequest {
	s.ProductCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) SetUserConfigDataSourceList(v []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList) *SaveDataSourceOrderConfigRequest {
	s.UserConfigDataSourceList = v
	return s
}

type SaveDataSourceOrderConfigRequestUserConfigDataSourceList struct {
	// This parameter is required.
	//
	// example:
	//
	// QuarkCommonNews
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Number *int32 `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SystemSearch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SaveDataSourceOrderConfigRequestUserConfigDataSourceList) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetCode(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Code = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetName(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Name = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetNumber(v int32) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Number = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetType(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Type = &v
	return s
}

type SaveDataSourceOrderConfigShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// miaobi
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	UserConfigDataSourceListShrink *string `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty"`
}

func (s SaveDataSourceOrderConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetAgentKey(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetProductCode(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetUserConfigDataSourceListShrink(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.UserConfigDataSourceListShrink = &v
	return s
}

type SaveDataSourceOrderConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveDataSourceOrderConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigResponseBody) SetCode(v string) *SaveDataSourceOrderConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetData(v bool) *SaveDataSourceOrderConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetHttpStatusCode(v int32) *SaveDataSourceOrderConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetMessage(v string) *SaveDataSourceOrderConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetRequestId(v string) *SaveDataSourceOrderConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetSuccess(v bool) *SaveDataSourceOrderConfigResponseBody {
	s.Success = &v
	return s
}

type SaveDataSourceOrderConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveDataSourceOrderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveDataSourceOrderConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigResponse) SetHeaders(v map[string]*string) *SaveDataSourceOrderConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveDataSourceOrderConfigResponse) SetStatusCode(v int32) *SaveDataSourceOrderConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponse) SetBody(v *SaveDataSourceOrderConfigResponseBody) *SaveDataSourceOrderConfigResponse {
	s.Body = v
	return s
}

type SaveMaterialDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author   *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// false
	BothSavePrivateAndShare *bool     `json:"BothSavePrivateAndShare,omitempty" xml:"BothSavePrivateAndShare,omitempty"`
	DocKeywords             []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// excel
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// IntellijSearch
	SrcFrom     *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SaveMaterialDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentRequest) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentRequest) SetAgentKey(v string) *SaveMaterialDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetAuthor(v string) *SaveMaterialDocumentRequest {
	s.Author = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentRequest {
	s.BothSavePrivateAndShare = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetDocKeywords(v []*string) *SaveMaterialDocumentRequest {
	s.DocKeywords = v
	return s
}

func (s *SaveMaterialDocumentRequest) SetDocType(v string) *SaveMaterialDocumentRequest {
	s.DocType = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetExternalUrl(v string) *SaveMaterialDocumentRequest {
	s.ExternalUrl = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetHtmlContent(v string) *SaveMaterialDocumentRequest {
	s.HtmlContent = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetPubTime(v string) *SaveMaterialDocumentRequest {
	s.PubTime = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetShareAttr(v int32) *SaveMaterialDocumentRequest {
	s.ShareAttr = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetSrcFrom(v string) *SaveMaterialDocumentRequest {
	s.SrcFrom = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetSummary(v string) *SaveMaterialDocumentRequest {
	s.Summary = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetTextContent(v string) *SaveMaterialDocumentRequest {
	s.TextContent = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetTitle(v string) *SaveMaterialDocumentRequest {
	s.Title = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetUrl(v string) *SaveMaterialDocumentRequest {
	s.Url = &v
	return s
}

type SaveMaterialDocumentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author   *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// false
	BothSavePrivateAndShare *bool   `json:"BothSavePrivateAndShare,omitempty" xml:"BothSavePrivateAndShare,omitempty"`
	DocKeywordsShrink       *string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// excel
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// IntellijSearch
	SrcFrom     *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SaveMaterialDocumentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentShrinkRequest) SetAgentKey(v string) *SaveMaterialDocumentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetAuthor(v string) *SaveMaterialDocumentShrinkRequest {
	s.Author = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentShrinkRequest {
	s.BothSavePrivateAndShare = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetDocKeywordsShrink(v string) *SaveMaterialDocumentShrinkRequest {
	s.DocKeywordsShrink = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetDocType(v string) *SaveMaterialDocumentShrinkRequest {
	s.DocType = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetExternalUrl(v string) *SaveMaterialDocumentShrinkRequest {
	s.ExternalUrl = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetHtmlContent(v string) *SaveMaterialDocumentShrinkRequest {
	s.HtmlContent = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetPubTime(v string) *SaveMaterialDocumentShrinkRequest {
	s.PubTime = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetShareAttr(v int32) *SaveMaterialDocumentShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetSrcFrom(v string) *SaveMaterialDocumentShrinkRequest {
	s.SrcFrom = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetSummary(v string) *SaveMaterialDocumentShrinkRequest {
	s.Summary = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetTextContent(v string) *SaveMaterialDocumentShrinkRequest {
	s.TextContent = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetTitle(v string) *SaveMaterialDocumentShrinkRequest {
	s.Title = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetUrl(v string) *SaveMaterialDocumentShrinkRequest {
	s.Url = &v
	return s
}

type SaveMaterialDocumentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 12
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveMaterialDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentResponseBody) SetCode(v string) *SaveMaterialDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetData(v int64) *SaveMaterialDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetHttpStatusCode(v int32) *SaveMaterialDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetMessage(v string) *SaveMaterialDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetRequestId(v string) *SaveMaterialDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetSuccess(v bool) *SaveMaterialDocumentResponseBody {
	s.Success = &v
	return s
}

type SaveMaterialDocumentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveMaterialDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveMaterialDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentResponse) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentResponse) SetHeaders(v map[string]*string) *SaveMaterialDocumentResponse {
	s.Headers = v
	return s
}

func (s *SaveMaterialDocumentResponse) SetStatusCode(v int32) *SaveMaterialDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveMaterialDocumentResponse) SetBody(v *SaveMaterialDocumentResponseBody) *SaveMaterialDocumentResponse {
	s.Body = v
	return s
}

type SearchNewsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// false
	FilterNotNull *bool `json:"FilterNotNull,omitempty" xml:"FilterNotNull,omitempty"`
	// example:
	//
	// false
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// example:
	//
	// 81
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 35
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 检索Query
	Query         *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	SearchSources []*string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
}

func (s SearchNewsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsRequest) GoString() string {
	return s.String()
}

func (s *SearchNewsRequest) SetAgentKey(v string) *SearchNewsRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchNewsRequest) SetFilterNotNull(v bool) *SearchNewsRequest {
	s.FilterNotNull = &v
	return s
}

func (s *SearchNewsRequest) SetIncludeContent(v bool) *SearchNewsRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchNewsRequest) SetPage(v int32) *SearchNewsRequest {
	s.Page = &v
	return s
}

func (s *SearchNewsRequest) SetPageSize(v int32) *SearchNewsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchNewsRequest) SetQuery(v string) *SearchNewsRequest {
	s.Query = &v
	return s
}

func (s *SearchNewsRequest) SetSearchSources(v []*string) *SearchNewsRequest {
	s.SearchSources = v
	return s
}

type SearchNewsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// false
	FilterNotNull *bool `json:"FilterNotNull,omitempty" xml:"FilterNotNull,omitempty"`
	// example:
	//
	// false
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// example:
	//
	// 81
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 35
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 检索Query
	Query               *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SearchSourcesShrink *string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty"`
}

func (s SearchNewsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchNewsShrinkRequest) SetAgentKey(v string) *SearchNewsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetFilterNotNull(v bool) *SearchNewsShrinkRequest {
	s.FilterNotNull = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetIncludeContent(v bool) *SearchNewsShrinkRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetPage(v int32) *SearchNewsShrinkRequest {
	s.Page = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetPageSize(v int32) *SearchNewsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetQuery(v string) *SearchNewsShrinkRequest {
	s.Query = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetSearchSourcesShrink(v string) *SearchNewsShrinkRequest {
	s.SearchSourcesShrink = &v
	return s
}

type SearchNewsResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                        `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*SearchNewsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchNewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchNewsResponseBody) SetCode(v string) *SearchNewsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchNewsResponseBody) SetCurrent(v int32) *SearchNewsResponseBody {
	s.Current = &v
	return s
}

func (s *SearchNewsResponseBody) SetData(v []*SearchNewsResponseBodyData) *SearchNewsResponseBody {
	s.Data = v
	return s
}

func (s *SearchNewsResponseBody) SetHttpStatusCode(v int32) *SearchNewsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchNewsResponseBody) SetMessage(v string) *SearchNewsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchNewsResponseBody) SetRequestId(v string) *SearchNewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchNewsResponseBody) SetSize(v int32) *SearchNewsResponseBody {
	s.Size = &v
	return s
}

func (s *SearchNewsResponseBody) SetSuccess(v bool) *SearchNewsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchNewsResponseBody) SetTotal(v int32) *SearchNewsResponseBody {
	s.Total = &v
	return s
}

type SearchNewsResponseBodyData struct {
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
	// 9a598b44c6444da5907b8ea68a5f82c4
	DocUuid   *string   `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-01-18 06:46:22
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// QuarkCommonNews
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// example:
	//
	// 夸克检索
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
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
	// 2024-01-18 06:46:22
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 文章URL
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchNewsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchNewsResponseBodyData) SetAuthor(v string) *SearchNewsResponseBodyData {
	s.Author = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetContent(v string) *SearchNewsResponseBodyData {
	s.Content = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetDocUuid(v string) *SearchNewsResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetImageUrls(v []*string) *SearchNewsResponseBodyData {
	s.ImageUrls = v
	return s
}

func (s *SearchNewsResponseBodyData) SetPubTime(v string) *SearchNewsResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSearchSource(v string) *SearchNewsResponseBodyData {
	s.SearchSource = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSearchSourceName(v string) *SearchNewsResponseBodyData {
	s.SearchSourceName = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSource(v string) *SearchNewsResponseBodyData {
	s.Source = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSummary(v string) *SearchNewsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetTag(v string) *SearchNewsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetTitle(v string) *SearchNewsResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetUpdateTime(v string) *SearchNewsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetUrl(v string) *SearchNewsResponseBodyData {
	s.Url = &v
	return s
}

type SearchNewsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchNewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchNewsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsResponse) GoString() string {
	return s.String()
}

func (s *SearchNewsResponse) SetHeaders(v map[string]*string) *SearchNewsResponse {
	s.Headers = v
	return s
}

func (s *SearchNewsResponse) SetStatusCode(v int32) *SearchNewsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchNewsResponse) SetBody(v *SearchNewsResponseBody) *SearchNewsResponse {
	s.Body = v
	return s
}

type SubmitAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daaa2e0c209xb26acb97009ea77bd4b_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// example:
	//
	// 2023-10-14 14:30:00
	TaskExecuteTime *string `json:"TaskExecuteTime,omitempty" xml:"TaskExecuteTime,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 任务提交参数
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
}

func (s SubmitAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskRequest) SetAgentKey(v string) *SubmitAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskCode(v string) *SubmitAsyncTaskRequest {
	s.TaskCode = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskExecuteTime(v string) *SubmitAsyncTaskRequest {
	s.TaskExecuteTime = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskName(v string) *SubmitAsyncTaskRequest {
	s.TaskName = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskParam(v string) *SubmitAsyncTaskRequest {
	s.TaskParam = &v
	return s
}

type SubmitAsyncTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitAsyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponseBody) SetCode(v string) *SubmitAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetData(v *SubmitAsyncTaskResponseBodyData) *SubmitAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetHttpStatusCode(v int32) *SubmitAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetMessage(v string) *SubmitAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetRequestId(v string) *SubmitAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetSuccess(v bool) *SubmitAsyncTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitAsyncTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// {}
	TaskIntermediateResult interface{} `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitAsyncTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskId(v string) *SubmitAsyncTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskIntermediateResult(v interface{}) *SubmitAsyncTaskResponseBodyData {
	s.TaskIntermediateResult = v
	return s
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskName(v string) *SubmitAsyncTaskResponseBodyData {
	s.TaskName = &v
	return s
}

type SubmitAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponse) SetHeaders(v map[string]*string) *SubmitAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAsyncTaskResponse) SetStatusCode(v int32) *SubmitAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAsyncTaskResponse) SetBody(v *SubmitAsyncTaskResponseBody) *SubmitAsyncTaskResponse {
	s.Body = v
	return s
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey  *string                                                              `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Documents []*SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 自定义观点的输入Prompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 待分析的主题名（documents与topic二者至少传一个）
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetDocuments(v []*SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Documents = v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetPrompt(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetTopic(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Topic = &v
	return s
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2024-01-22 10:29:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) String() string {
	return tea.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetAuthor(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Author = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetContent(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Content = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetPubTime(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.PubTime = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSource(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Source = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSummary(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Summary = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetTitle(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Title = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetUrl(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Url = &v
	return s
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 自定义观点的输入Prompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 待分析的主题名（documents与topic二者至少传一个）
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetAgentKey(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetDocumentsShrink(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetPrompt(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetTopic(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.Topic = &v
	return s
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponseBody) *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

type SubmitDocClusterTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Documents []*SubmitDocClusterTaskRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// 49
	SummaryLength *int32 `json:"SummaryLength,omitempty" xml:"SummaryLength,omitempty"`
	// example:
	//
	// 69
	TitleLength *int32 `json:"TitleLength,omitempty" xml:"TitleLength,omitempty"`
	// example:
	//
	// 15
	TopicCount *int32 `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s SubmitDocClusterTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskRequest) SetAgentKey(v string) *SubmitDocClusterTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetDocuments(v []*SubmitDocClusterTaskRequestDocuments) *SubmitDocClusterTaskRequest {
	s.Documents = v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetSummaryLength(v int32) *SubmitDocClusterTaskRequest {
	s.SummaryLength = &v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetTitleLength(v int32) *SubmitDocClusterTaskRequest {
	s.TitleLength = &v
	return s
}

func (s *SubmitDocClusterTaskRequest) SetTopicCount(v int32) *SubmitDocClusterTaskRequest {
	s.TopicCount = &v
	return s
}

type SubmitDocClusterTaskRequestDocuments struct {
	// This parameter is required.
	//
	// example:
	//
	// 文档内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文档ID。用于在返回聚类文章时标识文章。如果文章列表中都不传则使用数组索引作为ID。如果部分传则会报错
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 文档标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SubmitDocClusterTaskRequestDocuments) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocClusterTaskRequestDocuments) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskRequestDocuments) SetContent(v string) *SubmitDocClusterTaskRequestDocuments {
	s.Content = &v
	return s
}

func (s *SubmitDocClusterTaskRequestDocuments) SetDocId(v string) *SubmitDocClusterTaskRequestDocuments {
	s.DocId = &v
	return s
}

func (s *SubmitDocClusterTaskRequestDocuments) SetTitle(v string) *SubmitDocClusterTaskRequestDocuments {
	s.Title = &v
	return s
}

type SubmitDocClusterTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// example:
	//
	// 49
	SummaryLength *int32 `json:"SummaryLength,omitempty" xml:"SummaryLength,omitempty"`
	// example:
	//
	// 69
	TitleLength *int32 `json:"TitleLength,omitempty" xml:"TitleLength,omitempty"`
	// example:
	//
	// 15
	TopicCount *int32 `json:"TopicCount,omitempty" xml:"TopicCount,omitempty"`
}

func (s SubmitDocClusterTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocClusterTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskShrinkRequest) SetAgentKey(v string) *SubmitDocClusterTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetDocumentsShrink(v string) *SubmitDocClusterTaskShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetSummaryLength(v int32) *SubmitDocClusterTaskShrinkRequest {
	s.SummaryLength = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetTitleLength(v int32) *SubmitDocClusterTaskShrinkRequest {
	s.TitleLength = &v
	return s
}

func (s *SubmitDocClusterTaskShrinkRequest) SetTopicCount(v int32) *SubmitDocClusterTaskShrinkRequest {
	s.TopicCount = &v
	return s
}

type SubmitDocClusterTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocClusterTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitDocClusterTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskResponseBody) SetCode(v string) *SubmitDocClusterTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetData(v *SubmitDocClusterTaskResponseBodyData) *SubmitDocClusterTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetHttpStatusCode(v int32) *SubmitDocClusterTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetMessage(v string) *SubmitDocClusterTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetRequestId(v string) *SubmitDocClusterTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocClusterTaskResponseBody) SetSuccess(v bool) *SubmitDocClusterTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitDocClusterTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitDocClusterTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocClusterTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskResponseBodyData) SetTaskId(v string) *SubmitDocClusterTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type SubmitDocClusterTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocClusterTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocClusterTaskResponse) SetHeaders(v map[string]*string) *SubmitDocClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocClusterTaskResponse) SetStatusCode(v int32) *SubmitDocClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocClusterTaskResponse) SetBody(v *SubmitDocClusterTaskResponseBody) *SubmitDocClusterTaskResponse {
	s.Body = v
	return s
}

type SubmitTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey  *string                                                        `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Documents []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// TimedViewPoints
	PerspectiveTypes []*string `json:"PerspectiveTypes,omitempty" xml:"PerspectiveTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 待分析的主题名（documents与topic二者至少传一个）
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetDocuments(v []*SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Documents = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetPerspectiveTypes(v []*string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.PerspectiveTypes = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) SetTopic(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequest {
	s.Topic = &v
	return s
}

type SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2024-01-22 10:29:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 新浪
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) String() string {
	return tea.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetAuthor(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Author = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetContent(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Content = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetPubTime(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.PubTime = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSource(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Source = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetSummary(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Summary = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetTitle(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Title = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments) SetUrl(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskRequestDocuments {
	s.Url = &v
	return s
}

type SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// example:
	//
	// TimedViewPoints
	PerspectiveTypesShrink *string `json:"PerspectiveTypes,omitempty" xml:"PerspectiveTypes,omitempty"`
	// example:
	//
	// 待分析的主题名（documents与topic二者至少传一个）
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetAgentKey(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetDocumentsShrink(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetPerspectiveTypesShrink(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.PerspectiveTypesShrink = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest) SetTopic(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest {
	s.Topic = &v
	return s
}

type SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetCode(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetData(v *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetMessage(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetRequestId(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) SetSuccess(v bool) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTaskId(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData) SetTaskName(v string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBodyData {
	s.TaskName = &v
	return s
}

type SubmitTopicSelectionPerspectiveAnalysisTaskResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTopicSelectionPerspectiveAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) SetHeaders(v map[string]*string) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) SetStatusCode(v int32) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTopicSelectionPerspectiveAnalysisTaskResponse) SetBody(v *SubmitTopicSelectionPerspectiveAnalysisTaskResponseBody) *SubmitTopicSelectionPerspectiveAnalysisTaskResponse {
	s.Body = v
	return s
}

type UpdateCustomTextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 商品code
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 96
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateCustomTextRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomTextRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomTextRequest) SetAgentKey(v string) *UpdateCustomTextRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateCustomTextRequest) SetCommodityCode(v string) *UpdateCustomTextRequest {
	s.CommodityCode = &v
	return s
}

func (s *UpdateCustomTextRequest) SetContent(v string) *UpdateCustomTextRequest {
	s.Content = &v
	return s
}

func (s *UpdateCustomTextRequest) SetId(v int64) *UpdateCustomTextRequest {
	s.Id = &v
	return s
}

func (s *UpdateCustomTextRequest) SetTitle(v string) *UpdateCustomTextRequest {
	s.Title = &v
	return s
}

type UpdateCustomTextResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 48
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomTextResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomTextResponseBody) SetCode(v string) *UpdateCustomTextResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetData(v int64) *UpdateCustomTextResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetHttpStatusCode(v int32) *UpdateCustomTextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetMessage(v string) *UpdateCustomTextResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetRequestId(v string) *UpdateCustomTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomTextResponseBody) SetSuccess(v bool) *UpdateCustomTextResponseBody {
	s.Success = &v
	return s
}

type UpdateCustomTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomTextResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomTextResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomTextResponse) SetHeaders(v map[string]*string) *UpdateCustomTextResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomTextResponse) SetStatusCode(v int32) *UpdateCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomTextResponse) SetBody(v *UpdateCustomTextResponseBody) *UpdateCustomTextResponse {
	s.Body = v
	return s
}

type UpdateGeneratedContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 正文
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 正文
	ContentText *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36
	Id       *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// example:
	//
	// 创作xx文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 评论类文章
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentRequest) SetAgentKey(v string) *UpdateGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetContent(v string) *UpdateGeneratedContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetContentText(v string) *UpdateGeneratedContentRequest {
	s.ContentText = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetId(v int64) *UpdateGeneratedContentRequest {
	s.Id = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetKeywords(v []*string) *UpdateGeneratedContentRequest {
	s.Keywords = v
	return s
}

func (s *UpdateGeneratedContentRequest) SetPrompt(v string) *UpdateGeneratedContentRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetTitle(v string) *UpdateGeneratedContentRequest {
	s.Title = &v
	return s
}

type UpdateGeneratedContentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 正文
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 正文
	ContentText *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// 创作xx文章
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 评论类文章
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateGeneratedContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentShrinkRequest) SetAgentKey(v string) *UpdateGeneratedContentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetContent(v string) *UpdateGeneratedContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetContentText(v string) *UpdateGeneratedContentShrinkRequest {
	s.ContentText = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetId(v int64) *UpdateGeneratedContentShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetKeywordsShrink(v string) *UpdateGeneratedContentShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetPrompt(v string) *UpdateGeneratedContentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetTitle(v string) *UpdateGeneratedContentShrinkRequest {
	s.Title = &v
	return s
}

type UpdateGeneratedContentResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentResponseBody) SetCode(v string) *UpdateGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetData(v bool) *UpdateGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetHttpStatusCode(v int32) *UpdateGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetMessage(v string) *UpdateGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetRequestId(v string) *UpdateGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetSuccess(v bool) *UpdateGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type UpdateGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentResponse) SetHeaders(v map[string]*string) *UpdateGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateGeneratedContentResponse) SetStatusCode(v int32) *UpdateGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGeneratedContentResponse) SetBody(v *UpdateGeneratedContentResponseBody) *UpdateGeneratedContentResponse {
	s.Body = v
	return s
}

type UpdateMaterialDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey    *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author      *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	DocKeywords []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// image
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 44
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime  *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// UserUpload
	SrcFrom     *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateMaterialDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentRequest) SetAgentKey(v string) *UpdateMaterialDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetAuthor(v string) *UpdateMaterialDocumentRequest {
	s.Author = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetDocKeywords(v []*string) *UpdateMaterialDocumentRequest {
	s.DocKeywords = v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetDocType(v string) *UpdateMaterialDocumentRequest {
	s.DocType = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetExternalUrl(v string) *UpdateMaterialDocumentRequest {
	s.ExternalUrl = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetHtmlContent(v string) *UpdateMaterialDocumentRequest {
	s.HtmlContent = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetId(v int64) *UpdateMaterialDocumentRequest {
	s.Id = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetPubTime(v string) *UpdateMaterialDocumentRequest {
	s.PubTime = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetRegionId(v string) *UpdateMaterialDocumentRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetShareAttr(v int32) *UpdateMaterialDocumentRequest {
	s.ShareAttr = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetSrcFrom(v string) *UpdateMaterialDocumentRequest {
	s.SrcFrom = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetSummary(v string) *UpdateMaterialDocumentRequest {
	s.Summary = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetTextContent(v string) *UpdateMaterialDocumentRequest {
	s.TextContent = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetTitle(v string) *UpdateMaterialDocumentRequest {
	s.Title = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetUrl(v string) *UpdateMaterialDocumentRequest {
	s.Url = &v
	return s
}

type UpdateMaterialDocumentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey          *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author            *string `json:"Author,omitempty" xml:"Author,omitempty"`
	DocKeywordsShrink *string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// image
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 44
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime  *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// UserUpload
	SrcFrom     *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateMaterialDocumentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentShrinkRequest) SetAgentKey(v string) *UpdateMaterialDocumentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetAuthor(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Author = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetDocKeywordsShrink(v string) *UpdateMaterialDocumentShrinkRequest {
	s.DocKeywordsShrink = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetDocType(v string) *UpdateMaterialDocumentShrinkRequest {
	s.DocType = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetExternalUrl(v string) *UpdateMaterialDocumentShrinkRequest {
	s.ExternalUrl = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetHtmlContent(v string) *UpdateMaterialDocumentShrinkRequest {
	s.HtmlContent = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetId(v int64) *UpdateMaterialDocumentShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetPubTime(v string) *UpdateMaterialDocumentShrinkRequest {
	s.PubTime = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetRegionId(v string) *UpdateMaterialDocumentShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetShareAttr(v int32) *UpdateMaterialDocumentShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetSrcFrom(v string) *UpdateMaterialDocumentShrinkRequest {
	s.SrcFrom = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetSummary(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Summary = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetTextContent(v string) *UpdateMaterialDocumentShrinkRequest {
	s.TextContent = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetTitle(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetUrl(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Url = &v
	return s
}

type UpdateMaterialDocumentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 82
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMaterialDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentResponseBody) SetCode(v string) *UpdateMaterialDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetData(v int64) *UpdateMaterialDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetHttpStatusCode(v int32) *UpdateMaterialDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetMessage(v string) *UpdateMaterialDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetRequestId(v string) *UpdateMaterialDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetSuccess(v bool) *UpdateMaterialDocumentResponseBody {
	s.Success = &v
	return s
}

type UpdateMaterialDocumentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMaterialDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMaterialDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentResponse) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentResponse) SetHeaders(v map[string]*string) *UpdateMaterialDocumentResponse {
	s.Headers = v
	return s
}

func (s *UpdateMaterialDocumentResponse) SetStatusCode(v int32) *UpdateMaterialDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMaterialDocumentResponse) SetBody(v *UpdateMaterialDocumentResponseBody) *UpdateMaterialDocumentResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aimiaobi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消异步任务
//
// @param request - CancelAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAsyncTaskResponse
func (client *Client) CancelAsyncTaskWithOptions(request *CancelAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *CancelAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelAsyncTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消异步任务
//
// @param request - CancelAsyncTaskRequest
//
// @return CancelAsyncTaskResponse
func (client *Client) CancelAsyncTask(request *CancelAsyncTaskRequest) (_result *CancelAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CancelAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 清除所有干预内容
//
// @param request - ClearIntervenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearIntervenesResponse
func (client *Client) ClearIntervenesWithOptions(request *ClearIntervenesRequest, runtime *util.RuntimeOptions) (_result *ClearIntervenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearIntervenes"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearIntervenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清除所有干预内容
//
// @param request - ClearIntervenesRequest
//
// @return ClearIntervenesResponse
func (client *Client) ClearIntervenes(request *ClearIntervenesRequest) (_result *ClearIntervenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearIntervenesResponse{}
	_body, _err := client.ClearIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档管理-创建
//
// @param tmpReq - CreateGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGeneratedContentResponse
func (client *Client) CreateGeneratedContentWithOptions(tmpReq *CreateGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *CreateGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateGeneratedContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Keywords)) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, tea.String("Keywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentDomain)) {
		body["ContentDomain"] = request.ContentDomain
	}

	if !tea.BoolValue(util.IsUnset(request.ContentText)) {
		body["ContentText"] = request.ContentText
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsShrink)) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-创建
//
// @param request - CreateGeneratedContentRequest
//
// @return CreateGeneratedContentResponse
func (client *Client) CreateGeneratedContent(request *CreateGeneratedContentRequest) (_result *CreateGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGeneratedContentResponse{}
	_body, _err := client.CreateGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取授权token
//
// @param request - CreateTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTokenResponse
func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, runtime *util.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateToken"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取授权token
//
// @param request - CreateTokenRequest
//
// @return CreateTokenResponse
func (client *Client) CreateToken(request *CreateTokenRequest) (_result *CreateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTokenResponse{}
	_body, _err := client.CreateTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义文本
//
// @param request - DeleteCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTextResponse
func (client *Client) DeleteCustomTextWithOptions(request *DeleteCustomTextRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomText"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义文本
//
// @param request - DeleteCustomTextRequest
//
// @return DeleteCustomTextResponse
func (client *Client) DeleteCustomText(request *DeleteCustomTextRequest) (_result *DeleteCustomTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomTextResponse{}
	_body, _err := client.DeleteCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据主题删除自定义主题事件
//
// @param request - DeleteCustomTopicByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTopicByTopicResponse
func (client *Client) DeleteCustomTopicByTopicWithOptions(request *DeleteCustomTopicByTopicRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomTopicByTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomTopicByTopic"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomTopicByTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据主题删除自定义主题事件
//
// @param request - DeleteCustomTopicByTopicRequest
//
// @return DeleteCustomTopicByTopicResponse
func (client *Client) DeleteCustomTopicByTopic(request *DeleteCustomTopicByTopicRequest) (_result *DeleteCustomTopicByTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomTopicByTopicResponse{}
	_body, _err := client.DeleteCustomTopicByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据自定义观点ID删除自定义观点
//
// @param request - DeleteCustomTopicViewPointByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTopicViewPointByIdResponse
func (client *Client) DeleteCustomTopicViewPointByIdWithOptions(request *DeleteCustomTopicViewPointByIdRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomTopicViewPointByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomViewPointId)) {
		body["CustomViewPointId"] = request.CustomViewPointId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomTopicViewPointById"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomTopicViewPointByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据自定义观点ID删除自定义观点
//
// @param request - DeleteCustomTopicViewPointByIdRequest
//
// @return DeleteCustomTopicViewPointByIdResponse
func (client *Client) DeleteCustomTopicViewPointById(request *DeleteCustomTopicViewPointByIdRequest) (_result *DeleteCustomTopicViewPointByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomTopicViewPointByIdResponse{}
	_body, _err := client.DeleteCustomTopicViewPointByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档管理-删除。
//
// @param request - DeleteGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGeneratedContentResponse
func (client *Client) DeleteGeneratedContentWithOptions(request *DeleteGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *DeleteGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-删除。
//
// @param request - DeleteGeneratedContentRequest
//
// @return DeleteGeneratedContentResponse
func (client *Client) DeleteGeneratedContent(request *DeleteGeneratedContentRequest) (_result *DeleteGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGeneratedContentResponse{}
	_body, _err := client.DeleteGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除干预规则
//
// @param request - DeleteInterveneRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInterveneRuleResponse
func (client *Client) DeleteInterveneRuleWithOptions(request *DeleteInterveneRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteInterveneRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInterveneRule"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInterveneRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除干预规则
//
// @param request - DeleteInterveneRuleRequest
//
// @return DeleteInterveneRuleResponse
func (client *Client) DeleteInterveneRule(request *DeleteInterveneRuleRequest) (_result *DeleteInterveneRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInterveneRuleResponse{}
	_body, _err := client.DeleteInterveneRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据ID删除素材
//
// @param request - DeleteMaterialByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMaterialByIdResponse
func (client *Client) DeleteMaterialByIdWithOptions(request *DeleteMaterialByIdRequest, runtime *util.RuntimeOptions) (_result *DeleteMaterialByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMaterialById"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMaterialByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID删除素材
//
// @param request - DeleteMaterialByIdRequest
//
// @return DeleteMaterialByIdResponse
func (client *Client) DeleteMaterialById(request *DeleteMaterialByIdRequest) (_result *DeleteMaterialByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMaterialByIdResponse{}
	_body, _err := client.DeleteMaterialByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从链接中提取文档内容
//
// @param tmpReq - DocumentExtractionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocumentExtractionResponse
func (client *Client) DocumentExtractionWithOptions(tmpReq *DocumentExtractionRequest, runtime *util.RuntimeOptions) (_result *DocumentExtractionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DocumentExtractionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Urls)) {
		request.UrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Urls, tea.String("Urls"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UrlsShrink)) {
		body["Urls"] = request.UrlsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DocumentExtraction"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DocumentExtractionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从链接中提取文档内容
//
// @param request - DocumentExtractionRequest
//
// @return DocumentExtractionResponse
func (client *Client) DocumentExtraction(request *DocumentExtractionRequest) (_result *DocumentExtractionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DocumentExtractionResponse{}
	_body, _err := client.DocumentExtractionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档管理-导出。
//
// @param request - ExportGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportGeneratedContentResponse
func (client *Client) ExportGeneratedContentWithOptions(request *ExportGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *ExportGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-导出。
//
// @param request - ExportGeneratedContentRequest
//
// @return ExportGeneratedContentResponse
func (client *Client) ExportGeneratedContent(request *ExportGeneratedContentRequest) (_result *ExportGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportGeneratedContentResponse{}
	_body, _err := client.ExportGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出选题策划文档，响应为一个可公开访问的URL。一小时后失效
//
// @param tmpReq - ExportHotTopicPlanningProposalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportHotTopicPlanningProposalsResponse
func (client *Client) ExportHotTopicPlanningProposalsWithOptions(tmpReq *ExportHotTopicPlanningProposalsRequest, runtime *util.RuntimeOptions) (_result *ExportHotTopicPlanningProposalsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExportHotTopicPlanningProposalsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomViewPointIds)) {
		request.CustomViewPointIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomViewPointIds, tea.String("CustomViewPointIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Titles)) {
		request.TitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Titles, tea.String("Titles"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomViewPointIdsShrink)) {
		body["CustomViewPointIds"] = request.CustomViewPointIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExportType)) {
		body["ExportType"] = request.ExportType
	}

	if !tea.BoolValue(util.IsUnset(request.TitlesShrink)) {
		body["Titles"] = request.TitlesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	if !tea.BoolValue(util.IsUnset(request.ViewPointType)) {
		body["ViewPointType"] = request.ViewPointType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportHotTopicPlanningProposals"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportHotTopicPlanningProposalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出选题策划文档，响应为一个可公开访问的URL。一小时后失效
//
// @param request - ExportHotTopicPlanningProposalsRequest
//
// @return ExportHotTopicPlanningProposalsResponse
func (client *Client) ExportHotTopicPlanningProposals(request *ExportHotTopicPlanningProposalsRequest) (_result *ExportHotTopicPlanningProposalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportHotTopicPlanningProposalsResponse{}
	_body, _err := client.ExportHotTopicPlanningProposalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出所有干预内容
//
// @param request - ExportIntervenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportIntervenesResponse
func (client *Client) ExportIntervenesWithOptions(request *ExportIntervenesRequest, runtime *util.RuntimeOptions) (_result *ExportIntervenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportIntervenes"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportIntervenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出所有干预内容
//
// @param request - ExportIntervenesRequest
//
// @return ExportIntervenesResponse
func (client *Client) ExportIntervenes(request *ExportIntervenesRequest) (_result *ExportIntervenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportIntervenesResponse{}
	_body, _err := client.ExportIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 反馈某次生成的结果
//
// @param tmpReq - FeedbackDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FeedbackDialogueResponse
func (client *Client) FeedbackDialogueWithOptions(tmpReq *FeedbackDialogueRequest, runtime *util.RuntimeOptions) (_result *FeedbackDialogueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FeedbackDialogueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RatingTags)) {
		request.RatingTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RatingTags, tea.String("RatingTags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerResponse)) {
		body["CustomerResponse"] = request.CustomerResponse
	}

	if !tea.BoolValue(util.IsUnset(request.GoodText)) {
		body["GoodText"] = request.GoodText
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedResponse)) {
		body["ModifiedResponse"] = request.ModifiedResponse
	}

	if !tea.BoolValue(util.IsUnset(request.Rating)) {
		body["Rating"] = request.Rating
	}

	if !tea.BoolValue(util.IsUnset(request.RatingTagsShrink)) {
		body["RatingTags"] = request.RatingTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FeedbackDialogue"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FeedbackDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 反馈某次生成的结果
//
// @param request - FeedbackDialogueRequest
//
// @return FeedbackDialogueResponse
func (client *Client) FeedbackDialogue(request *FeedbackDialogueRequest) (_result *FeedbackDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FeedbackDialogueResponse{}
	_body, _err := client.FeedbackDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片任务执行结果
//
// @param tmpReq - FetchImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchImageTaskResponse
func (client *Client) FetchImageTaskWithOptions(tmpReq *FetchImageTaskRequest, runtime *util.RuntimeOptions) (_result *FetchImageTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FetchImageTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskIdList)) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, tea.String("TaskIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArticleTaskId)) {
		body["ArticleTaskId"] = request.ArticleTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdListShrink)) {
		body["TaskIdList"] = request.TaskIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FetchImageTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchImageTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片任务执行结果
//
// @param request - FetchImageTaskRequest
//
// @return FetchImageTaskResponse
func (client *Client) FetchImageTask(request *FetchImageTaskRequest) (_result *FetchImageTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchImageTaskResponse{}
	_body, _err := client.FetchImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成临时可访问的公开url
//
// @param request - GenerateFileUrlByKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateFileUrlByKeyResponse
func (client *Client) GenerateFileUrlByKeyWithOptions(request *GenerateFileUrlByKeyRequest, runtime *util.RuntimeOptions) (_result *GenerateFileUrlByKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateFileUrlByKey"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateFileUrlByKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成临时可访问的公开url
//
// @param request - GenerateFileUrlByKeyRequest
//
// @return GenerateFileUrlByKeyResponse
func (client *Client) GenerateFileUrlByKey(request *GenerateFileUrlByKeyRequest) (_result *GenerateFileUrlByKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateFileUrlByKeyResponse{}
	_body, _err := client.GenerateFileUrlByKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能配图，图片生成任务
//
// @param tmpReq - GenerateImageTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateImageTaskResponse
func (client *Client) GenerateImageTaskWithOptions(tmpReq *GenerateImageTaskRequest, runtime *util.RuntimeOptions) (_result *GenerateImageTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateImageTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ParagraphList)) {
		request.ParagraphListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParagraphList, tea.String("ParagraphList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArticleTaskId)) {
		body["ArticleTaskId"] = request.ArticleTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ParagraphListShrink)) {
		body["ParagraphList"] = request.ParagraphListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Style)) {
		body["Style"] = request.Style
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateImageTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateImageTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能配图，图片生成任务
//
// @param request - GenerateImageTaskRequest
//
// @return GenerateImageTaskResponse
func (client *Client) GenerateImageTask(request *GenerateImageTaskRequest) (_result *GenerateImageTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateImageTaskResponse{}
	_body, _err := client.GenerateImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成上传配置
//
// @param request - GenerateUploadConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadConfigResponse
func (client *Client) GenerateUploadConfigWithOptions(request *GenerateUploadConfigRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDir)) {
		body["ParentDir"] = request.ParentDir
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUploadConfig"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUploadConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成上传配置
//
// @param request - GenerateUploadConfigRequest
//
// @return GenerateUploadConfigResponse
func (client *Client) GenerateUploadConfig(request *GenerateUploadConfigRequest) (_result *GenerateUploadConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadConfigResponse{}
	_body, _err := client.GenerateUploadConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视角生成
//
// @param tmpReq - GenerateViewPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateViewPointResponse
func (client *Client) GenerateViewPointWithOptions(tmpReq *GenerateViewPointRequest, runtime *util.RuntimeOptions) (_result *GenerateViewPointResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateViewPointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateViewPoint"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateViewPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视角生成
//
// @param request - GenerateViewPointRequest
//
// @return GenerateViewPointResponse
func (client *Client) GenerateViewPoint(request *GenerateViewPointRequest) (_result *GenerateViewPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateViewPointResponse{}
	_body, _err := client.GenerateViewPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自定义文本
//
// @param request - GetCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomTextResponse
func (client *Client) GetCustomTextWithOptions(request *GetCustomTextRequest, runtime *util.RuntimeOptions) (_result *GetCustomTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomText"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义文本
//
// @param request - GetCustomTextRequest
//
// @return GetCustomTextResponse
func (client *Client) GetCustomText(request *GetCustomTextRequest) (_result *GetCustomTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomTextResponse{}
	_body, _err := client.GetCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自定义选题视角分析任务结果
//
// @param request - GetCustomTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(request *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest, runtime *util.RuntimeOptions) (_result *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomTopicSelectionPerspectiveAnalysisTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义选题视角分析任务结果
//
// @param request - GetCustomTopicSelectionPerspectiveAnalysisTaskRequest
//
// @return GetCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetCustomTopicSelectionPerspectiveAnalysisTask(request *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) (_result *GetCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.GetCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取系统数据源配置和个人配置
//
// @param request - GetDataSourceOrderConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceOrderConfigResponse
func (client *Client) GetDataSourceOrderConfigWithOptions(request *GetDataSourceOrderConfigRequest, runtime *util.RuntimeOptions) (_result *GetDataSourceOrderConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSourceOrderConfig"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceOrderConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统数据源配置和个人配置
//
// @param request - GetDataSourceOrderConfigRequest
//
// @return GetDataSourceOrderConfigResponse
func (client *Client) GetDataSourceOrderConfig(request *GetDataSourceOrderConfigRequest) (_result *GetDataSourceOrderConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataSourceOrderConfigResponse{}
	_body, _err := client.GetDataSourceOrderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档聚合任务结果
//
// @param request - GetDocClusterTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocClusterTaskResponse
func (client *Client) GetDocClusterTaskWithOptions(request *GetDocClusterTaskRequest, runtime *util.RuntimeOptions) (_result *GetDocClusterTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocClusterTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocClusterTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档聚合任务结果
//
// @param request - GetDocClusterTaskRequest
//
// @return GetDocClusterTaskResponse
func (client *Client) GetDocClusterTask(request *GetDocClusterTaskRequest) (_result *GetDocClusterTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocClusterTaskResponse{}
	_body, _err := client.GetDocClusterTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档管理-查询详情。
//
// @param request - GetGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGeneratedContentResponse
func (client *Client) GetGeneratedContentWithOptions(request *GetGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *GetGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-查询详情。
//
// @param request - GetGeneratedContentRequest
//
// @return GetGeneratedContentResponse
func (client *Client) GetGeneratedContent(request *GetGeneratedContentRequest) (_result *GetGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGeneratedContentResponse{}
	_body, _err := client.GetGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得干预全局回复
//
// @param request - GetInterveneGlobalReplyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneGlobalReplyResponse
func (client *Client) GetInterveneGlobalReplyWithOptions(request *GetInterveneGlobalReplyRequest, runtime *util.RuntimeOptions) (_result *GetInterveneGlobalReplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneGlobalReply"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneGlobalReplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预全局回复
//
// @param request - GetInterveneGlobalReplyRequest
//
// @return GetInterveneGlobalReplyResponse
func (client *Client) GetInterveneGlobalReply(request *GetInterveneGlobalReplyRequest) (_result *GetInterveneGlobalReplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneGlobalReplyResponse{}
	_body, _err := client.GetInterveneGlobalReplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得导入任务信息
//
// @param request - GetInterveneImportTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneImportTaskInfoResponse
func (client *Client) GetInterveneImportTaskInfoWithOptions(request *GetInterveneImportTaskInfoRequest, runtime *util.RuntimeOptions) (_result *GetInterveneImportTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneImportTaskInfo"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneImportTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得导入任务信息
//
// @param request - GetInterveneImportTaskInfoRequest
//
// @return GetInterveneImportTaskInfoResponse
func (client *Client) GetInterveneImportTaskInfo(request *GetInterveneImportTaskInfoRequest) (_result *GetInterveneImportTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneImportTaskInfoResponse{}
	_body, _err := client.GetInterveneImportTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得干预项规则详情
//
// @param request - GetInterveneRuleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneRuleDetailResponse
func (client *Client) GetInterveneRuleDetailWithOptions(request *GetInterveneRuleDetailRequest, runtime *util.RuntimeOptions) (_result *GetInterveneRuleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneRuleDetail"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneRuleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预项规则详情
//
// @param request - GetInterveneRuleDetailRequest
//
// @return GetInterveneRuleDetailResponse
func (client *Client) GetInterveneRuleDetail(request *GetInterveneRuleDetailRequest) (_result *GetInterveneRuleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneRuleDetailResponse{}
	_body, _err := client.GetInterveneRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得干预导入模版文件下载地址
//
// @param request - GetInterveneTemplateFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterveneTemplateFileUrlResponse
func (client *Client) GetInterveneTemplateFileUrlWithOptions(request *GetInterveneTemplateFileUrlRequest, runtime *util.RuntimeOptions) (_result *GetInterveneTemplateFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneTemplateFileUrl"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneTemplateFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预导入模版文件下载地址
//
// @param request - GetInterveneTemplateFileUrlRequest
//
// @return GetInterveneTemplateFileUrlResponse
func (client *Client) GetInterveneTemplateFileUrl(request *GetInterveneTemplateFileUrlRequest) (_result *GetInterveneTemplateFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneTemplateFileUrlResponse{}
	_body, _err := client.GetInterveneTemplateFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据ID获取素材内容
//
// @param request - GetMaterialByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMaterialByIdResponse
func (client *Client) GetMaterialByIdWithOptions(request *GetMaterialByIdRequest, runtime *util.RuntimeOptions) (_result *GetMaterialByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMaterialById"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMaterialByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID获取素材内容
//
// @param request - GetMaterialByIdRequest
//
// @return GetMaterialByIdResponse
func (client *Client) GetMaterialById(request *GetMaterialByIdRequest) (_result *GetMaterialByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMaterialByIdResponse{}
	_body, _err := client.GetMaterialByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前用户的配置
//
// @param request - GetPropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPropertiesResponse
func (client *Client) GetPropertiesWithOptions(request *GetPropertiesRequest, runtime *util.RuntimeOptions) (_result *GetPropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProperties"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前用户的配置
//
// @param request - GetPropertiesRequest
//
// @return GetPropertiesResponse
func (client *Client) GetProperties(request *GetPropertiesRequest) (_result *GetPropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPropertiesResponse{}
	_body, _err := client.GetPropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据ID获取热点事件信息
//
// @param request - GetTopicByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicByIdResponse
func (client *Client) GetTopicByIdWithOptions(request *GetTopicByIdRequest, runtime *util.RuntimeOptions) (_result *GetTopicByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopicById"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID获取热点事件信息
//
// @param request - GetTopicByIdRequest
//
// @return GetTopicByIdResponse
func (client *Client) GetTopicById(request *GetTopicByIdRequest) (_result *GetTopicByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicByIdResponse{}
	_body, _err := client.GetTopicByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取选题视角分析任务结果
//
// @param request - GetTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetTopicSelectionPerspectiveAnalysisTaskWithOptions(request *GetTopicSelectionPerspectiveAnalysisTaskRequest, runtime *util.RuntimeOptions) (_result *GetTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopicSelectionPerspectiveAnalysisTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取选题视角分析任务结果
//
// @param request - GetTopicSelectionPerspectiveAnalysisTaskRequest
//
// @return GetTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) GetTopicSelectionPerspectiveAnalysisTask(request *GetTopicSelectionPerspectiveAnalysisTaskRequest) (_result *GetTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.GetTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入干预文件
//
// @param request - ImportInterveneFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportInterveneFileResponse
func (client *Client) ImportInterveneFileWithOptions(request *ImportInterveneFileRequest, runtime *util.RuntimeOptions) (_result *ImportInterveneFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		body["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportInterveneFile"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportInterveneFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入干预文件
//
// @param request - ImportInterveneFileRequest
//
// @return ImportInterveneFileResponse
func (client *Client) ImportInterveneFile(request *ImportInterveneFileRequest) (_result *ImportInterveneFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportInterveneFileResponse{}
	_body, _err := client.ImportInterveneFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 异步导入干预文件
//
// @param request - ImportInterveneFileAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportInterveneFileAsyncResponse
func (client *Client) ImportInterveneFileAsyncWithOptions(request *ImportInterveneFileAsyncRequest, runtime *util.RuntimeOptions) (_result *ImportInterveneFileAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		body["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportInterveneFileAsync"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportInterveneFileAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 异步导入干预文件
//
// @param request - ImportInterveneFileAsyncRequest
//
// @return ImportInterveneFileAsyncResponse
func (client *Client) ImportInterveneFileAsync(request *ImportInterveneFileAsyncRequest) (_result *ImportInterveneFileAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportInterveneFileAsyncResponse{}
	_body, _err := client.ImportInterveneFileAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置干预全局回复
//
// @param tmpReq - InsertInterveneGlobalReplyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertInterveneGlobalReplyResponse
func (client *Client) InsertInterveneGlobalReplyWithOptions(tmpReq *InsertInterveneGlobalReplyRequest, runtime *util.RuntimeOptions) (_result *InsertInterveneGlobalReplyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertInterveneGlobalReplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReplyMessagList)) {
		request.ReplyMessagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReplyMessagList, tea.String("ReplyMessagList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReplyMessagListShrink)) {
		body["ReplyMessagList"] = request.ReplyMessagListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertInterveneGlobalReply"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertInterveneGlobalReplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置干预全局回复
//
// @param request - InsertInterveneGlobalReplyRequest
//
// @return InsertInterveneGlobalReplyResponse
func (client *Client) InsertInterveneGlobalReply(request *InsertInterveneGlobalReplyRequest) (_result *InsertInterveneGlobalReplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertInterveneGlobalReplyResponse{}
	_body, _err := client.InsertInterveneGlobalReplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 插入干预规则
//
// @param tmpReq - InsertInterveneRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertInterveneRuleResponse
func (client *Client) InsertInterveneRuleWithOptions(tmpReq *InsertInterveneRuleRequest, runtime *util.RuntimeOptions) (_result *InsertInterveneRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertInterveneRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InterveneRuleConfig)) {
		request.InterveneRuleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterveneRuleConfig, tea.String("InterveneRuleConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InterveneRuleConfigShrink)) {
		body["InterveneRuleConfig"] = request.InterveneRuleConfigShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertInterveneRule"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertInterveneRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插入干预规则
//
// @param request - InsertInterveneRuleRequest
//
// @return InsertInterveneRuleResponse
func (client *Client) InsertInterveneRule(request *InsertInterveneRuleRequest) (_result *InsertInterveneRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertInterveneRuleResponse{}
	_body, _err := client.InsertInterveneRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param tmpReq - ListAsyncTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasksWithOptions(tmpReq *ListAsyncTasksRequest, runtime *util.RuntimeOptions) (_result *ListAsyncTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAsyncTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskStatusList)) {
		request.TaskStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskStatusList, tea.String("TaskStatusList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TaskTypeList)) {
		request.TaskTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskTypeList, tea.String("TaskTypeList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.TaskCode)) {
		body["TaskCode"] = request.TaskCode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatusListShrink)) {
		body["TaskStatusList"] = request.TaskStatusListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTypeListShrink)) {
		body["TaskTypeList"] = request.TaskTypeListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsyncTasks"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param request - ListAsyncTasksRequest
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasks(request *ListAsyncTasksRequest) (_result *ListAsyncTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.ListAsyncTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取系统自定义预设
//
// @param request - ListBuildConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBuildConfigsResponse
func (client *Client) ListBuildConfigsWithOptions(request *ListBuildConfigsRequest, runtime *util.RuntimeOptions) (_result *ListBuildConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBuildConfigs"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBuildConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统自定义预设
//
// @param request - ListBuildConfigsRequest
//
// @return ListBuildConfigsResponse
func (client *Client) ListBuildConfigs(request *ListBuildConfigsRequest) (_result *ListBuildConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBuildConfigsResponse{}
	_body, _err := client.ListBuildConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自定义文本列表
//
// @param request - ListCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomTextResponse
func (client *Client) ListCustomTextWithOptions(request *ListCustomTextRequest, runtime *util.RuntimeOptions) (_result *ListCustomTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["CommodityCode"] = request.CommodityCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomText"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义文本列表
//
// @param request - ListCustomTextRequest
//
// @return ListCustomTextResponse
func (client *Client) ListCustomText(request *ListCustomTextRequest) (_result *ListCustomTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomTextResponse{}
	_body, _err := client.ListCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自定义视角列表
//
// @param tmpReq - ListCustomViewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomViewPointsResponse
func (client *Client) ListCustomViewPointsWithOptions(tmpReq *ListCustomViewPointsRequest, runtime *util.RuntimeOptions) (_result *ListCustomViewPointsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCustomViewPointsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Attitudes)) {
		request.AttitudesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attitudes, tea.String("Attitudes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CustomViewPointIds)) {
		request.CustomViewPointIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomViewPointIds, tea.String("CustomViewPointIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attitude)) {
		body["Attitude"] = request.Attitude
	}

	if !tea.BoolValue(util.IsUnset(request.AttitudesShrink)) {
		body["Attitudes"] = request.AttitudesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomViewPointId)) {
		body["CustomViewPointId"] = request.CustomViewPointId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomViewPointIdsShrink)) {
		body["CustomViewPointIds"] = request.CustomViewPointIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicId)) {
		body["TopicId"] = request.TopicId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomViewPoints"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomViewPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义视角列表
//
// @param request - ListCustomViewPointsRequest
//
// @return ListCustomViewPointsResponse
func (client *Client) ListCustomViewPoints(request *ListCustomViewPointsRequest) (_result *ListCustomViewPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomViewPointsResponse{}
	_body, _err := client.ListCustomViewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成历史列表
//
// @param request - ListDialoguesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDialoguesResponse
func (client *Client) ListDialoguesWithOptions(request *ListDialoguesRequest, runtime *util.RuntimeOptions) (_result *ListDialoguesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueType)) {
		body["DialogueType"] = request.DialogueType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDialogues"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDialoguesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成历史列表
//
// @param request - ListDialoguesRequest
//
// @return ListDialoguesResponse
func (client *Client) ListDialogues(request *ListDialoguesRequest) (_result *ListDialoguesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDialoguesResponse{}
	_body, _err := client.ListDialoguesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新颖视角列表
//
// @param request - ListFreshViewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFreshViewPointsResponse
func (client *Client) ListFreshViewPointsWithOptions(request *ListFreshViewPointsRequest, runtime *util.RuntimeOptions) (_result *ListFreshViewPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFreshViewPoints"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFreshViewPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新颖视角列表
//
// @param request - ListFreshViewPointsRequest
//
// @return ListFreshViewPointsResponse
func (client *Client) ListFreshViewPoints(request *ListFreshViewPointsRequest) (_result *ListFreshViewPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFreshViewPointsResponse{}
	_body, _err := client.ListFreshViewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档管理-列表。
//
// @param request - ListGeneratedContentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGeneratedContentsResponse
func (client *Client) ListGeneratedContentsWithOptions(request *ListGeneratedContentsRequest, runtime *util.RuntimeOptions) (_result *ListGeneratedContentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentDomain)) {
		body["ContentDomain"] = request.ContentDomain
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGeneratedContents"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGeneratedContentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-列表。
//
// @param request - ListGeneratedContentsRequest
//
// @return ListGeneratedContentsResponse
func (client *Client) ListGeneratedContents(request *ListGeneratedContentsRequest) (_result *ListGeneratedContentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGeneratedContentsResponse{}
	_body, _err := client.ListGeneratedContentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取分类的热点新闻
//
// @param tmpReq - ListHotNewsWithTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotNewsWithTypeResponse
func (client *Client) ListHotNewsWithTypeWithOptions(tmpReq *ListHotNewsWithTypeRequest, runtime *util.RuntimeOptions) (_result *ListHotNewsWithTypeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotNewsWithTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NewsTypes)) {
		request.NewsTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NewsTypes, tea.String("NewsTypes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.NewsType)) {
		body["NewsType"] = request.NewsType
	}

	if !tea.BoolValue(util.IsUnset(request.NewsTypesShrink)) {
		body["NewsTypes"] = request.NewsTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotNewsWithType"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotNewsWithTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取分类的热点新闻
//
// @param request - ListHotNewsWithTypeRequest
//
// @return ListHotNewsWithTypeResponse
func (client *Client) ListHotNewsWithType(request *ListHotNewsWithTypeRequest) (_result *ListHotNewsWithTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotNewsWithTypeResponse{}
	_body, _err := client.ListHotNewsWithTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取所有平台热榜源列表
//
// @param request - ListHotSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotSourcesResponse
func (client *Client) ListHotSourcesWithOptions(request *ListHotSourcesRequest, runtime *util.RuntimeOptions) (_result *ListHotSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotSources"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有平台热榜源列表
//
// @param request - ListHotSourcesRequest
//
// @return ListHotSourcesResponse
func (client *Client) ListHotSources(request *ListHotSourcesRequest) (_result *ListHotSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotSourcesResponse{}
	_body, _err := client.ListHotSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取热点事件列表
//
// @param tmpReq - ListHotTopicsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotTopicsResponse
func (client *Client) ListHotTopicsWithOptions(tmpReq *ListHotTopicsRequest, runtime *util.RuntimeOptions) (_result *ListHotTopicsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotTopicsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TopicIds)) {
		request.TopicIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TopicIds, tea.String("TopicIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Topics)) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, tea.String("Topics"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TopicIdsShrink)) {
		body["TopicIds"] = request.TopicIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TopicQuery)) {
		body["TopicQuery"] = request.TopicQuery
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	if !tea.BoolValue(util.IsUnset(request.TopicVersion)) {
		body["TopicVersion"] = request.TopicVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TopicsShrink)) {
		body["Topics"] = request.TopicsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WithNews)) {
		body["WithNews"] = request.WithNews
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotTopics"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotTopicsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热点事件列表
//
// @param request - ListHotTopicsRequest
//
// @return ListHotTopicsResponse
func (client *Client) ListHotTopics(request *ListHotTopicsRequest) (_result *ListHotTopicsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotTopicsResponse{}
	_body, _err := client.ListHotTopicsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 热门视角列表
//
// @param request - ListHotViewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotViewPointsResponse
func (client *Client) ListHotViewPointsWithOptions(request *ListHotViewPointsRequest, runtime *util.RuntimeOptions) (_result *ListHotViewPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotViewPoints"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotViewPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 热门视角列表
//
// @param request - ListHotViewPointsRequest
//
// @return ListHotViewPointsResponse
func (client *Client) ListHotViewPoints(request *ListHotViewPointsRequest) (_result *ListHotViewPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotViewPointsResponse{}
	_body, _err := client.ListHotViewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得干预项目数量列表
//
// @param request - ListInterveneCntRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterveneCntResponse
func (client *Client) ListInterveneCntWithOptions(request *ListInterveneCntRequest, runtime *util.RuntimeOptions) (_result *ListInterveneCntResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterveneCnt"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterveneCntResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预项目数量列表
//
// @param request - ListInterveneCntRequest
//
// @return ListInterveneCntResponse
func (client *Client) ListInterveneCnt(request *ListInterveneCntRequest) (_result *ListInterveneCntResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInterveneCntResponse{}
	_body, _err := client.ListInterveneCntWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得导入任务列表
//
// @param request - ListInterveneImportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterveneImportTasksResponse
func (client *Client) ListInterveneImportTasksWithOptions(request *ListInterveneImportTasksRequest, runtime *util.RuntimeOptions) (_result *ListInterveneImportTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterveneImportTasks"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterveneImportTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得导入任务列表
//
// @param request - ListInterveneImportTasksRequest
//
// @return ListInterveneImportTasksResponse
func (client *Client) ListInterveneImportTasks(request *ListInterveneImportTasksRequest) (_result *ListInterveneImportTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInterveneImportTasksResponse{}
	_body, _err := client.ListInterveneImportTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得干预规则列表
//
// @param request - ListInterveneRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterveneRulesResponse
func (client *Client) ListInterveneRulesWithOptions(request *ListInterveneRulesRequest, runtime *util.RuntimeOptions) (_result *ListInterveneRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterveneRules"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterveneRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预规则列表
//
// @param request - ListInterveneRulesRequest
//
// @return ListInterveneRulesResponse
func (client *Client) ListInterveneRules(request *ListInterveneRulesRequest) (_result *ListInterveneRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInterveneRulesResponse{}
	_body, _err := client.ListInterveneRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得干预项列表
//
// @param request - ListIntervenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntervenesResponse
func (client *Client) ListIntervenesWithOptions(request *ListIntervenesRequest, runtime *util.RuntimeOptions) (_result *ListIntervenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InterveneType)) {
		body["InterveneType"] = request.InterveneType
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntervenes"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntervenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得干预项列表
//
// @param request - ListIntervenesRequest
//
// @return ListIntervenesResponse
func (client *Client) ListIntervenes(request *ListIntervenesRequest) (_result *ListIntervenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntervenesResponse{}
	_body, _err := client.ListIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询素材列表
//
// @param tmpReq - ListMaterialDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMaterialDocumentsResponse
func (client *Client) ListMaterialDocumentsWithOptions(tmpReq *ListMaterialDocumentsRequest, runtime *util.RuntimeOptions) (_result *ListMaterialDocumentsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMaterialDocumentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocTypeList)) {
		request.DocTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocTypeList, tea.String("DocTypeList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Keywords)) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, tea.String("Keywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.DocTypeListShrink)) {
		body["DocTypeList"] = request.DocTypeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GeneratePublicUrl)) {
		body["GeneratePublicUrl"] = request.GeneratePublicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsShrink)) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ShareAttr)) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTimeEnd)) {
		body["UpdateTimeEnd"] = request.UpdateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTimeStart)) {
		body["UpdateTimeStart"] = request.UpdateTimeStart
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMaterialDocuments"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMaterialDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询素材列表
//
// @param request - ListMaterialDocumentsRequest
//
// @return ListMaterialDocumentsResponse
func (client *Client) ListMaterialDocuments(request *ListMaterialDocumentsRequest) (_result *ListMaterialDocumentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMaterialDocumentsResponse{}
	_body, _err := client.ListMaterialDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取选题策划列表
//
// @param tmpReq - ListPlanningProposalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPlanningProposalResponse
func (client *Client) ListPlanningProposalWithOptions(tmpReq *ListPlanningProposalRequest, runtime *util.RuntimeOptions) (_result *ListPlanningProposalResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPlanningProposalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomViewPointIds)) {
		request.CustomViewPointIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomViewPointIds, tea.String("CustomViewPointIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Titles)) {
		request.TitlesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Titles, tea.String("Titles"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomViewPointId)) {
		body["CustomViewPointId"] = request.CustomViewPointId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomViewPointIdsShrink)) {
		body["CustomViewPointIds"] = request.CustomViewPointIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TitlesShrink)) {
		body["Titles"] = request.TitlesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	if !tea.BoolValue(util.IsUnset(request.TopicVersion)) {
		body["TopicVersion"] = request.TopicVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ViewPointType)) {
		body["ViewPointType"] = request.ViewPointType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPlanningProposal"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPlanningProposalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取选题策划列表
//
// @param request - ListPlanningProposalRequest
//
// @return ListPlanningProposalResponse
func (client *Client) ListPlanningProposal(request *ListPlanningProposalRequest) (_result *ListPlanningProposalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPlanningProposalResponse{}
	_body, _err := client.ListPlanningProposalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 时效性视角列表
//
// @param request - ListTimedViewAttitudeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTimedViewAttitudeResponse
func (client *Client) ListTimedViewAttitudeWithOptions(request *ListTimedViewAttitudeRequest, runtime *util.RuntimeOptions) (_result *ListTimedViewAttitudeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTimedViewAttitude"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTimedViewAttitudeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 时效性视角列表
//
// @param request - ListTimedViewAttitudeRequest
//
// @return ListTimedViewAttitudeResponse
func (client *Client) ListTimedViewAttitude(request *ListTimedViewAttitudeRequest) (_result *ListTimedViewAttitudeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTimedViewAttitudeResponse{}
	_body, _err := client.ListTimedViewAttitudeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取热点推荐事件
//
// @param request - ListTopicRecommendEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicRecommendEventListResponse
func (client *Client) ListTopicRecommendEventListWithOptions(request *ListTopicRecommendEventListRequest, runtime *util.RuntimeOptions) (_result *ListTopicRecommendEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTopicRecommendEventList"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTopicRecommendEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热点推荐事件
//
// @param request - ListTopicRecommendEventListRequest
//
// @return ListTopicRecommendEventListResponse
func (client *Client) ListTopicRecommendEventList(request *ListTopicRecommendEventListRequest) (_result *ListTopicRecommendEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTopicRecommendEventListResponse{}
	_body, _err := client.ListTopicRecommendEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取主题事件推荐观点列表
//
// @param request - ListTopicViewPointRecommendEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicViewPointRecommendEventListResponse
func (client *Client) ListTopicViewPointRecommendEventListWithOptions(request *ListTopicViewPointRecommendEventListRequest, runtime *util.RuntimeOptions) (_result *ListTopicViewPointRecommendEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTopicViewPointRecommendEventList"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTopicViewPointRecommendEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主题事件推荐观点列表
//
// @param request - ListTopicViewPointRecommendEventListRequest
//
// @return ListTopicViewPointRecommendEventListResponse
func (client *Client) ListTopicViewPointRecommendEventList(request *ListTopicViewPointRecommendEventListRequest) (_result *ListTopicViewPointRecommendEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTopicViewPointRecommendEventListResponse{}
	_body, _err := client.ListTopicViewPointRecommendEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取系统所有实例信息
//
// @param request - ListVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVersionsResponse
func (client *Client) ListVersionsWithOptions(request *ListVersionsRequest, runtime *util.RuntimeOptions) (_result *ListVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVersions"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取系统所有实例信息
//
// @param request - ListVersionsRequest
//
// @return ListVersionsResponse
func (client *Client) ListVersions(request *ListVersionsRequest) (_result *ListVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVersionsResponse{}
	_body, _err := client.ListVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 网友视角列表
//
// @param request - ListWebReviewPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWebReviewPointsResponse
func (client *Client) ListWebReviewPointsWithOptions(request *ListWebReviewPointsRequest, runtime *util.RuntimeOptions) (_result *ListWebReviewPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWebReviewPoints"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWebReviewPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 网友视角列表
//
// @param request - ListWebReviewPointsRequest
//
// @return ListWebReviewPointsResponse
func (client *Client) ListWebReviewPoints(request *ListWebReviewPointsRequest) (_result *ListWebReviewPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWebReviewPointsResponse{}
	_body, _err := client.ListWebReviewPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据taskId查询异步任务状态
//
// @param request - QueryAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAsyncTaskResponse
func (client *Client) QueryAsyncTaskWithOptions(request *QueryAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *QueryAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAsyncTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据taskId查询异步任务状态
//
// @param request - QueryAsyncTaskRequest
//
// @return QueryAsyncTaskResponse
func (client *Client) QueryAsyncTask(request *QueryAsyncTaskRequest) (_result *QueryAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAsyncTaskResponse{}
	_body, _err := client.QueryAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 内容缩写
//
// @param request - RunAbbreviationContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAbbreviationContentResponse
func (client *Client) RunAbbreviationContentWithOptions(request *RunAbbreviationContentRequest, runtime *util.RuntimeOptions) (_result *RunAbbreviationContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunAbbreviationContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunAbbreviationContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容缩写
//
// @param request - RunAbbreviationContentRequest
//
// @return RunAbbreviationContentResponse
func (client *Client) RunAbbreviationContent(request *RunAbbreviationContentRequest) (_result *RunAbbreviationContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunAbbreviationContentResponse{}
	_body, _err := client.RunAbbreviationContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 内容续写
//
// @param request - RunContinueContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunContinueContentResponse
func (client *Client) RunContinueContentWithOptions(request *RunContinueContentRequest, runtime *util.RuntimeOptions) (_result *RunContinueContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunContinueContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunContinueContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容续写
//
// @param request - RunContinueContentRequest
//
// @return RunContinueContentResponse
func (client *Client) RunContinueContent(request *RunContinueContentRequest) (_result *RunContinueContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunContinueContentResponse{}
	_body, _err := client.RunContinueContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自定义热点话题分析
//
// @param request - RunCustomHotTopicAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCustomHotTopicAnalysisResponse
func (client *Client) RunCustomHotTopicAnalysisWithOptions(request *RunCustomHotTopicAnalysisRequest, runtime *util.RuntimeOptions) (_result *RunCustomHotTopicAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AskUser)) {
		body["AskUser"] = request.AskUser
	}

	if !tea.BoolValue(util.IsUnset(request.ForceAnalysisExistsTopic)) {
		body["ForceAnalysisExistsTopic"] = request.ForceAnalysisExistsTopic
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBack)) {
		body["UserBack"] = request.UserBack
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCustomHotTopicAnalysis"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCustomHotTopicAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义热点话题分析
//
// @param request - RunCustomHotTopicAnalysisRequest
//
// @return RunCustomHotTopicAnalysisResponse
func (client *Client) RunCustomHotTopicAnalysis(request *RunCustomHotTopicAnalysisRequest) (_result *RunCustomHotTopicAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCustomHotTopicAnalysisResponse{}
	_body, _err := client.RunCustomHotTopicAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自定义选题视角分析
//
// @param request - RunCustomHotTopicViewPointAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCustomHotTopicViewPointAnalysisResponse
func (client *Client) RunCustomHotTopicViewPointAnalysisWithOptions(request *RunCustomHotTopicViewPointAnalysisRequest, runtime *util.RuntimeOptions) (_result *RunCustomHotTopicViewPointAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AskUser)) {
		body["AskUser"] = request.AskUser
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.SearchQuery)) {
		body["SearchQuery"] = request.SearchQuery
	}

	if !tea.BoolValue(util.IsUnset(request.SkipAskUser)) {
		body["SkipAskUser"] = request.SkipAskUser
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.TopicId)) {
		body["TopicId"] = request.TopicId
	}

	if !tea.BoolValue(util.IsUnset(request.TopicSource)) {
		body["TopicSource"] = request.TopicSource
	}

	if !tea.BoolValue(util.IsUnset(request.TopicVersion)) {
		body["TopicVersion"] = request.TopicVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserBack)) {
		body["UserBack"] = request.UserBack
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCustomHotTopicViewPointAnalysis"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCustomHotTopicViewPointAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自定义选题视角分析
//
// @param request - RunCustomHotTopicViewPointAnalysisRequest
//
// @return RunCustomHotTopicViewPointAnalysisResponse
func (client *Client) RunCustomHotTopicViewPointAnalysis(request *RunCustomHotTopicViewPointAnalysisRequest) (_result *RunCustomHotTopicViewPointAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCustomHotTopicViewPointAnalysisResponse{}
	_body, _err := client.RunCustomHotTopicViewPointAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 内容扩写
//
// @param request - RunExpandContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunExpandContentResponse
func (client *Client) RunExpandContentWithOptions(request *RunExpandContentRequest, runtime *util.RuntimeOptions) (_result *RunExpandContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunExpandContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunExpandContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容扩写
//
// @param request - RunExpandContentRequest
//
// @return RunExpandContentResponse
func (client *Client) RunExpandContent(request *RunExpandContentRequest) (_result *RunExpandContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunExpandContentResponse{}
	_body, _err := client.RunExpandContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// AI妙笔-创作-抽取关键词
//
// @param tmpReq - RunKeywordsExtractionGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunKeywordsExtractionGenerationResponse
func (client *Client) RunKeywordsExtractionGenerationWithOptions(tmpReq *RunKeywordsExtractionGenerationRequest, runtime *util.RuntimeOptions) (_result *RunKeywordsExtractionGenerationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunKeywordsExtractionGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunKeywordsExtractionGeneration"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunKeywordsExtractionGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// AI妙笔-创作-抽取关键词
//
// @param request - RunKeywordsExtractionGenerationRequest
//
// @return RunKeywordsExtractionGenerationResponse
func (client *Client) RunKeywordsExtractionGeneration(request *RunKeywordsExtractionGenerationRequest) (_result *RunKeywordsExtractionGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunKeywordsExtractionGenerationResponse{}
	_body, _err := client.RunKeywordsExtractionGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创作-分步骤写作
//
// @param tmpReq - RunStepByStepWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStepByStepWritingResponse
func (client *Client) RunStepByStepWritingWithOptions(tmpReq *RunStepByStepWritingRequest, runtime *util.RuntimeOptions) (_result *RunStepByStepWritingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunStepByStepWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.WritingConfig)) {
		request.WritingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingConfig, tea.String("WritingConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OriginSessionId)) {
		body["OriginSessionId"] = request.OriginSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WritingConfigShrink)) {
		body["WritingConfig"] = request.WritingConfigShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunStepByStepWriting"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunStepByStepWritingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创作-分步骤写作
//
// @param request - RunStepByStepWritingRequest
//
// @return RunStepByStepWritingResponse
func (client *Client) RunStepByStepWriting(request *RunStepByStepWritingRequest) (_result *RunStepByStepWritingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunStepByStepWritingResponse{}
	_body, _err := client.RunStepByStepWritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 内容特点分析
//
// @param tmpReq - RunStyleFeatureAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunStyleFeatureAnalysisResponse
func (client *Client) RunStyleFeatureAnalysisWithOptions(tmpReq *RunStyleFeatureAnalysisRequest, runtime *util.RuntimeOptions) (_result *RunStyleFeatureAnalysisResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunStyleFeatureAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Contents)) {
		request.ContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contents, tea.String("Contents"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.MaterialIds)) {
		request.MaterialIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MaterialIds, tea.String("MaterialIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentsShrink)) {
		body["Contents"] = request.ContentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialIdsShrink)) {
		body["MaterialIds"] = request.MaterialIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunStyleFeatureAnalysis"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunStyleFeatureAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容特点分析
//
// @param request - RunStyleFeatureAnalysisRequest
//
// @return RunStyleFeatureAnalysisResponse
func (client *Client) RunStyleFeatureAnalysis(request *RunStyleFeatureAnalysisRequest) (_result *RunStyleFeatureAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunStyleFeatureAnalysisResponse{}
	_body, _err := client.RunStyleFeatureAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 内容摘要生成
//
// @param request - RunSummaryGenerateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSummaryGenerateResponse
func (client *Client) RunSummaryGenerateWithOptions(request *RunSummaryGenerateRequest, runtime *util.RuntimeOptions) (_result *RunSummaryGenerateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunSummaryGenerate"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunSummaryGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内容摘要生成
//
// @param request - RunSummaryGenerateRequest
//
// @return RunSummaryGenerateResponse
func (client *Client) RunSummaryGenerate(request *RunSummaryGenerateRequest) (_result *RunSummaryGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunSummaryGenerateResponse{}
	_body, _err := client.RunSummaryGenerateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创作-文本润色
//
// @param request - RunTextPolishingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTextPolishingResponse
func (client *Client) RunTextPolishingWithOptions(request *RunTextPolishingRequest, runtime *util.RuntimeOptions) (_result *RunTextPolishingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunTextPolishing"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunTextPolishingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创作-文本润色
//
// @param request - RunTextPolishingRequest
//
// @return RunTextPolishingResponse
func (client *Client) RunTextPolishing(request *RunTextPolishingRequest) (_result *RunTextPolishingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunTextPolishingResponse{}
	_body, _err := client.RunTextPolishingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 妙笔：标题生成
//
// @param tmpReq - RunTitleGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTitleGenerationResponse
func (client *Client) RunTitleGenerationWithOptions(tmpReq *RunTitleGenerationRequest, runtime *util.RuntimeOptions) (_result *RunTitleGenerationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunTitleGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunTitleGeneration"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunTitleGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 妙笔：标题生成
//
// @param request - RunTitleGenerationRequest
//
// @return RunTitleGenerationResponse
func (client *Client) RunTitleGeneration(request *RunTitleGenerationRequest) (_result *RunTitleGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunTitleGenerationResponse{}
	_body, _err := client.RunTitleGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// AI妙笔-创作-中英文翻译
//
// @param tmpReq - RunTranslateGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTranslateGenerationResponse
func (client *Client) RunTranslateGenerationWithOptions(tmpReq *RunTranslateGenerationRequest, runtime *util.RuntimeOptions) (_result *RunTranslateGenerationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunTranslateGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunTranslateGeneration"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunTranslateGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// AI妙笔-创作-中英文翻译
//
// @param request - RunTranslateGenerationRequest
//
// @return RunTranslateGenerationResponse
func (client *Client) RunTranslateGeneration(request *RunTranslateGenerationRequest) (_result *RunTranslateGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunTranslateGenerationResponse{}
	_body, _err := client.RunTranslateGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// AI妙笔-创作-文风改写
//
// @param tmpReq - RunWriteToneGenerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWriteToneGenerationResponse
func (client *Client) RunWriteToneGenerationWithOptions(tmpReq *RunWriteToneGenerationRequest, runtime *util.RuntimeOptions) (_result *RunWriteToneGenerationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunWriteToneGenerationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunWriteToneGeneration"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunWriteToneGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// AI妙笔-创作-文风改写
//
// @param request - RunWriteToneGenerationRequest
//
// @return RunWriteToneGenerationResponse
func (client *Client) RunWriteToneGeneration(request *RunWriteToneGenerationRequest) (_result *RunWriteToneGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunWriteToneGenerationResponse{}
	_body, _err := client.RunWriteToneGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 直接写作
//
// @param tmpReq - RunWritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunWritingResponse
func (client *Client) RunWritingWithOptions(tmpReq *RunWritingRequest, runtime *util.RuntimeOptions) (_result *RunWritingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunWritingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.WritingConfig)) {
		request.WritingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WritingConfig, tea.String("WritingConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OriginSessionId)) {
		body["OriginSessionId"] = request.OriginSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WritingConfigShrink)) {
		body["WritingConfig"] = request.WritingConfigShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunWriting"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunWritingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直接写作
//
// @param request - RunWritingRequest
//
// @return RunWritingResponse
func (client *Client) RunWriting(request *RunWritingRequest) (_result *RunWritingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunWritingResponse{}
	_body, _err := client.RunWritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存自定义文本
//
// @param request - SaveCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveCustomTextResponse
func (client *Client) SaveCustomTextWithOptions(request *SaveCustomTextRequest, runtime *util.RuntimeOptions) (_result *SaveCustomTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveCustomText"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveCustomTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存自定义文本
//
// @param request - SaveCustomTextRequest
//
// @return SaveCustomTextResponse
func (client *Client) SaveCustomText(request *SaveCustomTextRequest) (_result *SaveCustomTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveCustomTextResponse{}
	_body, _err := client.SaveCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存用户的信源配置
//
// @param tmpReq - SaveDataSourceOrderConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveDataSourceOrderConfigResponse
func (client *Client) SaveDataSourceOrderConfigWithOptions(tmpReq *SaveDataSourceOrderConfigRequest, runtime *util.RuntimeOptions) (_result *SaveDataSourceOrderConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveDataSourceOrderConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserConfigDataSourceList)) {
		request.UserConfigDataSourceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserConfigDataSourceList, tea.String("UserConfigDataSourceList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserConfigDataSourceListShrink)) {
		body["UserConfigDataSourceList"] = request.UserConfigDataSourceListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveDataSourceOrderConfig"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveDataSourceOrderConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存用户的信源配置
//
// @param request - SaveDataSourceOrderConfigRequest
//
// @return SaveDataSourceOrderConfigResponse
func (client *Client) SaveDataSourceOrderConfig(request *SaveDataSourceOrderConfigRequest) (_result *SaveDataSourceOrderConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveDataSourceOrderConfigResponse{}
	_body, _err := client.SaveDataSourceOrderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存素材
//
// @param tmpReq - SaveMaterialDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveMaterialDocumentResponse
func (client *Client) SaveMaterialDocumentWithOptions(tmpReq *SaveMaterialDocumentRequest, runtime *util.RuntimeOptions) (_result *SaveMaterialDocumentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveMaterialDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocKeywords)) {
		request.DocKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocKeywords, tea.String("DocKeywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Author)) {
		body["Author"] = request.Author
	}

	if !tea.BoolValue(util.IsUnset(request.BothSavePrivateAndShare)) {
		body["BothSavePrivateAndShare"] = request.BothSavePrivateAndShare
	}

	if !tea.BoolValue(util.IsUnset(request.DocKeywordsShrink)) {
		body["DocKeywords"] = request.DocKeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUrl)) {
		body["ExternalUrl"] = request.ExternalUrl
	}

	if !tea.BoolValue(util.IsUnset(request.HtmlContent)) {
		body["HtmlContent"] = request.HtmlContent
	}

	if !tea.BoolValue(util.IsUnset(request.PubTime)) {
		body["PubTime"] = request.PubTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShareAttr)) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !tea.BoolValue(util.IsUnset(request.SrcFrom)) {
		body["SrcFrom"] = request.SrcFrom
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["Summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["TextContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveMaterialDocument"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveMaterialDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存素材
//
// @param request - SaveMaterialDocumentRequest
//
// @return SaveMaterialDocumentResponse
func (client *Client) SaveMaterialDocument(request *SaveMaterialDocumentRequest) (_result *SaveMaterialDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveMaterialDocumentResponse{}
	_body, _err := client.SaveMaterialDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新闻检索
//
// @param tmpReq - SearchNewsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchNewsResponse
func (client *Client) SearchNewsWithOptions(tmpReq *SearchNewsRequest, runtime *util.RuntimeOptions) (_result *SearchNewsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchNewsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SearchSources)) {
		request.SearchSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchSources, tea.String("SearchSources"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterNotNull)) {
		body["FilterNotNull"] = request.FilterNotNull
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeContent)) {
		body["IncludeContent"] = request.IncludeContent
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SearchSourcesShrink)) {
		body["SearchSources"] = request.SearchSourcesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchNews"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchNewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新闻检索
//
// @param request - SearchNewsRequest
//
// @return SearchNewsResponse
func (client *Client) SearchNews(request *SearchNewsRequest) (_result *SearchNewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchNewsResponse{}
	_body, _err := client.SearchNewsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交异步任务
//
// @param request - SubmitAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAsyncTaskResponse
func (client *Client) SubmitAsyncTaskWithOptions(request *SubmitAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskCode)) {
		body["TaskCode"] = request.TaskCode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecuteTime)) {
		body["TaskExecuteTime"] = request.TaskExecuteTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskParam)) {
		body["TaskParam"] = request.TaskParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAsyncTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交异步任务
//
// @param request - SubmitAsyncTaskRequest
//
// @return SubmitAsyncTaskResponse
func (client *Client) SubmitAsyncTask(request *SubmitAsyncTaskRequest) (_result *SubmitAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAsyncTaskResponse{}
	_body, _err := client.SubmitAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交自定义热点选题视角分析任务
//
// @param tmpReq - SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(tmpReq *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitCustomTopicSelectionPerspectiveAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Documents)) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, tea.String("Documents"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentsShrink)) {
		body["Documents"] = request.DocumentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitCustomTopicSelectionPerspectiveAnalysisTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交自定义热点选题视角分析任务
//
// @param request - SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest
//
// @return SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitCustomTopicSelectionPerspectiveAnalysisTask(request *SubmitCustomTopicSelectionPerspectiveAnalysisTaskRequest) (_result *SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitCustomTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.SubmitCustomTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交文档聚合任务
//
// @param tmpReq - SubmitDocClusterTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocClusterTaskResponse
func (client *Client) SubmitDocClusterTaskWithOptions(tmpReq *SubmitDocClusterTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitDocClusterTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitDocClusterTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Documents)) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, tea.String("Documents"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentsShrink)) {
		body["Documents"] = request.DocumentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryLength)) {
		body["SummaryLength"] = request.SummaryLength
	}

	if !tea.BoolValue(util.IsUnset(request.TitleLength)) {
		body["TitleLength"] = request.TitleLength
	}

	if !tea.BoolValue(util.IsUnset(request.TopicCount)) {
		body["TopicCount"] = request.TopicCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocClusterTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocClusterTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交文档聚合任务
//
// @param request - SubmitDocClusterTaskRequest
//
// @return SubmitDocClusterTaskResponse
func (client *Client) SubmitDocClusterTask(request *SubmitDocClusterTaskRequest) (_result *SubmitDocClusterTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDocClusterTaskResponse{}
	_body, _err := client.SubmitDocClusterTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交选题热点分析任务
//
// @param tmpReq - SubmitTopicSelectionPerspectiveAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitTopicSelectionPerspectiveAnalysisTaskWithOptions(tmpReq *SubmitTopicSelectionPerspectiveAnalysisTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTopicSelectionPerspectiveAnalysisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Documents)) {
		request.DocumentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Documents, tea.String("Documents"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PerspectiveTypes)) {
		request.PerspectiveTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PerspectiveTypes, tea.String("PerspectiveTypes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentsShrink)) {
		body["Documents"] = request.DocumentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PerspectiveTypesShrink)) {
		body["PerspectiveTypes"] = request.PerspectiveTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTopicSelectionPerspectiveAnalysisTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交选题热点分析任务
//
// @param request - SubmitTopicSelectionPerspectiveAnalysisTaskRequest
//
// @return SubmitTopicSelectionPerspectiveAnalysisTaskResponse
func (client *Client) SubmitTopicSelectionPerspectiveAnalysisTask(request *SubmitTopicSelectionPerspectiveAnalysisTaskRequest) (_result *SubmitTopicSelectionPerspectiveAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitTopicSelectionPerspectiveAnalysisTaskResponse{}
	_body, _err := client.SubmitTopicSelectionPerspectiveAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新自定义文本
//
// @param request - UpdateCustomTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomTextResponse
func (client *Client) UpdateCustomTextWithOptions(request *UpdateCustomTextRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomText"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新自定义文本
//
// @param request - UpdateCustomTextRequest
//
// @return UpdateCustomTextResponse
func (client *Client) UpdateCustomText(request *UpdateCustomTextRequest) (_result *UpdateCustomTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomTextResponse{}
	_body, _err := client.UpdateCustomTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档管理-更新。
//
// @param tmpReq - UpdateGeneratedContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGeneratedContentResponse
func (client *Client) UpdateGeneratedContentWithOptions(tmpReq *UpdateGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *UpdateGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateGeneratedContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Keywords)) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, tea.String("Keywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentText)) {
		body["ContentText"] = request.ContentText
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsShrink)) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档管理-更新。
//
// @param request - UpdateGeneratedContentRequest
//
// @return UpdateGeneratedContentResponse
func (client *Client) UpdateGeneratedContent(request *UpdateGeneratedContentRequest) (_result *UpdateGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGeneratedContentResponse{}
	_body, _err := client.UpdateGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据ID更新素材
//
// @param tmpReq - UpdateMaterialDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMaterialDocumentResponse
func (client *Client) UpdateMaterialDocumentWithOptions(tmpReq *UpdateMaterialDocumentRequest, runtime *util.RuntimeOptions) (_result *UpdateMaterialDocumentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateMaterialDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocKeywords)) {
		request.DocKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocKeywords, tea.String("DocKeywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Author)) {
		body["Author"] = request.Author
	}

	if !tea.BoolValue(util.IsUnset(request.DocKeywordsShrink)) {
		body["DocKeywords"] = request.DocKeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUrl)) {
		body["ExternalUrl"] = request.ExternalUrl
	}

	if !tea.BoolValue(util.IsUnset(request.HtmlContent)) {
		body["HtmlContent"] = request.HtmlContent
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PubTime)) {
		body["PubTime"] = request.PubTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShareAttr)) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !tea.BoolValue(util.IsUnset(request.SrcFrom)) {
		body["SrcFrom"] = request.SrcFrom
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["Summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["TextContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMaterialDocument"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMaterialDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID更新素材
//
// @param request - UpdateMaterialDocumentRequest
//
// @return UpdateMaterialDocumentResponse
func (client *Client) UpdateMaterialDocument(request *UpdateMaterialDocumentRequest) (_result *UpdateMaterialDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMaterialDocumentResponse{}
	_body, _err := client.UpdateMaterialDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
