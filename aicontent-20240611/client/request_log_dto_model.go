// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestLogDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *RequestLogDTO
	GetApiKeyId() *int64
	SetClientId(v int64) *RequestLogDTO
	GetClientId() *int64
	SetCompletionTokens(v int32) *RequestLogDTO
	GetCompletionTokens() *int32
	SetDeleteTag(v int32) *RequestLogDTO
	GetDeleteTag() *int32
	SetGmtCreate(v string) *RequestLogDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *RequestLogDTO
	GetGmtModified() *string
	SetId(v int64) *RequestLogDTO
	GetId() *int64
	SetModelId(v int64) *RequestLogDTO
	GetModelId() *int64
	SetModelName(v string) *RequestLogDTO
	GetModelName() *string
	SetPromptTokens(v int32) *RequestLogDTO
	GetPromptTokens() *int32
	SetRequestBody(v string) *RequestLogDTO
	GetRequestBody() *string
	SetRequestId(v string) *RequestLogDTO
	GetRequestId() *string
	SetRequestTime(v string) *RequestLogDTO
	GetRequestTime() *string
	SetResponseBody(v string) *RequestLogDTO
	GetResponseBody() *string
	SetResponseTimeMs(v int32) *RequestLogDTO
	GetResponseTimeMs() *int32
	SetStatusCode(v int32) *RequestLogDTO
	GetStatusCode() *int32
	SetTotalTokens(v int32) *RequestLogDTO
	GetTotalTokens() *int32
}

type RequestLogDTO struct {
	// example:
	//
	// 1
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// 50
	CompletionTokens *int32 `json:"completionTokens,omitempty" xml:"completionTokens,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 通义千问
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// 100
	PromptTokens *int32 `json:"promptTokens,omitempty" xml:"promptTokens,omitempty"`
	// example:
	//
	// {}
	RequestBody *string `json:"requestBody,omitempty" xml:"requestBody,omitempty"`
	// example:
	//
	// req-xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	RequestTime *string `json:"requestTime,omitempty" xml:"requestTime,omitempty"`
	// example:
	//
	// {}
	ResponseBody *string `json:"responseBody,omitempty" xml:"responseBody,omitempty"`
	// example:
	//
	// 100
	ResponseTimeMs *int32 `json:"responseTimeMs,omitempty" xml:"responseTimeMs,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	// example:
	//
	// 150
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RequestLogDTO) String() string {
	return dara.Prettify(s)
}

func (s RequestLogDTO) GoString() string {
	return s.String()
}

func (s *RequestLogDTO) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *RequestLogDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *RequestLogDTO) GetCompletionTokens() *int32 {
	return s.CompletionTokens
}

func (s *RequestLogDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *RequestLogDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *RequestLogDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *RequestLogDTO) GetId() *int64 {
	return s.Id
}

func (s *RequestLogDTO) GetModelId() *int64 {
	return s.ModelId
}

func (s *RequestLogDTO) GetModelName() *string {
	return s.ModelName
}

func (s *RequestLogDTO) GetPromptTokens() *int32 {
	return s.PromptTokens
}

func (s *RequestLogDTO) GetRequestBody() *string {
	return s.RequestBody
}

func (s *RequestLogDTO) GetRequestId() *string {
	return s.RequestId
}

func (s *RequestLogDTO) GetRequestTime() *string {
	return s.RequestTime
}

func (s *RequestLogDTO) GetResponseBody() *string {
	return s.ResponseBody
}

func (s *RequestLogDTO) GetResponseTimeMs() *int32 {
	return s.ResponseTimeMs
}

func (s *RequestLogDTO) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RequestLogDTO) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *RequestLogDTO) SetApiKeyId(v int64) *RequestLogDTO {
	s.ApiKeyId = &v
	return s
}

func (s *RequestLogDTO) SetClientId(v int64) *RequestLogDTO {
	s.ClientId = &v
	return s
}

func (s *RequestLogDTO) SetCompletionTokens(v int32) *RequestLogDTO {
	s.CompletionTokens = &v
	return s
}

func (s *RequestLogDTO) SetDeleteTag(v int32) *RequestLogDTO {
	s.DeleteTag = &v
	return s
}

func (s *RequestLogDTO) SetGmtCreate(v string) *RequestLogDTO {
	s.GmtCreate = &v
	return s
}

func (s *RequestLogDTO) SetGmtModified(v string) *RequestLogDTO {
	s.GmtModified = &v
	return s
}

func (s *RequestLogDTO) SetId(v int64) *RequestLogDTO {
	s.Id = &v
	return s
}

func (s *RequestLogDTO) SetModelId(v int64) *RequestLogDTO {
	s.ModelId = &v
	return s
}

func (s *RequestLogDTO) SetModelName(v string) *RequestLogDTO {
	s.ModelName = &v
	return s
}

func (s *RequestLogDTO) SetPromptTokens(v int32) *RequestLogDTO {
	s.PromptTokens = &v
	return s
}

func (s *RequestLogDTO) SetRequestBody(v string) *RequestLogDTO {
	s.RequestBody = &v
	return s
}

func (s *RequestLogDTO) SetRequestId(v string) *RequestLogDTO {
	s.RequestId = &v
	return s
}

func (s *RequestLogDTO) SetRequestTime(v string) *RequestLogDTO {
	s.RequestTime = &v
	return s
}

func (s *RequestLogDTO) SetResponseBody(v string) *RequestLogDTO {
	s.ResponseBody = &v
	return s
}

func (s *RequestLogDTO) SetResponseTimeMs(v int32) *RequestLogDTO {
	s.ResponseTimeMs = &v
	return s
}

func (s *RequestLogDTO) SetStatusCode(v int32) *RequestLogDTO {
	s.StatusCode = &v
	return s
}

func (s *RequestLogDTO) SetTotalTokens(v int32) *RequestLogDTO {
	s.TotalTokens = &v
	return s
}

func (s *RequestLogDTO) Validate() error {
	return dara.Validate(s)
}
