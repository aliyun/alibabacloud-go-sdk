// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentVersionResponseBody
	GetCode() *string
	SetData(v *GetAgentVersionResponseBodyData) *GetAgentVersionResponseBody
	GetData() *GetAgentVersionResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentVersionResponseBody
	GetSuccess() *bool
}

type GetAgentVersionResponseBody struct {
	// 业务状态码，请求成功时为 SUCCESS，失败时为错误码。
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体固定版本详情，configuration 为版本保存时的 Agent 配置。
	Data *GetAgentVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// HTTP 状态码。
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// 响应消息，请求成功时为 success，失败时为错误信息。
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求 ID，用于定位和排查请求。
	//
	// example:
	//
	// 5E8F3C2A-1B4D-4F6E-9A7C-0D2B8E6F4A1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 请求是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAgentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentVersionResponseBody) GetData() *GetAgentVersionResponseBodyData {
	return s.Data
}

func (s *GetAgentVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentVersionResponseBody) SetCode(v string) *GetAgentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetData(v *GetAgentVersionResponseBodyData) *GetAgentVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentVersionResponseBody) SetHttpStatusCode(v int32) *GetAgentVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetMessage(v string) *GetAgentVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetRequestId(v string) *GetAgentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetSuccess(v bool) *GetAgentVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentVersionResponseBodyData struct {
	// 版本所属的托管智能体或自定义智能体 ID。
	//
	// example:
	//
	// agent-0123456789abcdef0123456789abcdef
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// 系统按 Agent 递增生成的版本号，例如 v1、v2。查询、删除版本及服务端点引用版本时使用。
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// 仅表示版本模板与制品归档是否完整，不代表部署或运行状态。READY 为完整，INVALID 为归档缺失或不一致。
	//
	// example:
	//
	// READY
	ArtifactStatus *string `json:"artifactStatus,omitempty" xml:"artifactStatus,omitempty"`
	// 托管智能体与 GetManagedAgent 的 data 结构一致，自定义智能体与 GetCustomAgent 的 data 结构一致，脱敏规则相同。
	//
	// example:
	//
	// {"agentId":"agent-0123456789abcdef0123456789abcdef","workspaceId":"ws-0123456789abcdef01234","name":"customer-service","description":"客服智能体","status":"Running","latestSpecVersion":2,"instruction":"你是一名客服助手。"}
	Configuration map[string]interface{} `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// 版本保存的时间，采用 ISO 8601 UTC 格式。
	//
	// example:
	//
	// 2026-08-12T01:02:03Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 保存该版本的调用者阿里云账号 UID；RAM 用户调用时为其所属阿里云主账号 UID。
	//
	// example:
	//
	// 1234567890123456
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// 保存版本时填写的版本描述，最多 1024 个字符；保存时未填写则为空。
	//
	// example:
	//
	// 生产环境稳定版本
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 保存版本时 Agent 已生效的发布流水号，即当时的 effectiveSpecVersion。版本保存后该值不再变化。
	//
	// example:
	//
	// 2
	TargetSpecVersion *int64 `json:"targetSpecVersion,omitempty" xml:"targetSpecVersion,omitempty"`
	// 保存版本时用户指定的展示名称。早期保存且未记录名称的版本返回 agentVersion。
	//
	// example:
	//
	// 生产稳定版 V1.0
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty"`
	// 版本所属智能体所在的 AgentCore 工作空间 ID。
	//
	// example:
	//
	// ws-0123456789abcdef01234
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetAgentVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentVersionResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentVersionResponseBodyData) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *GetAgentVersionResponseBodyData) GetArtifactStatus() *string {
	return s.ArtifactStatus
}

func (s *GetAgentVersionResponseBodyData) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *GetAgentVersionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAgentVersionResponseBodyData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetAgentVersionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentVersionResponseBodyData) GetTargetSpecVersion() *int64 {
	return s.TargetSpecVersion
}

func (s *GetAgentVersionResponseBodyData) GetVersionName() *string {
	return s.VersionName
}

func (s *GetAgentVersionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAgentVersionResponseBodyData) SetAgentId(v string) *GetAgentVersionResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetAgentVersion(v string) *GetAgentVersionResponseBodyData {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetArtifactStatus(v string) *GetAgentVersionResponseBodyData {
	s.ArtifactStatus = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetConfiguration(v map[string]interface{}) *GetAgentVersionResponseBodyData {
	s.Configuration = v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetCreatedAt(v string) *GetAgentVersionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetCreatedBy(v string) *GetAgentVersionResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetDescription(v string) *GetAgentVersionResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetTargetSpecVersion(v int64) *GetAgentVersionResponseBodyData {
	s.TargetSpecVersion = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetVersionName(v string) *GetAgentVersionResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetWorkspaceId(v string) *GetAgentVersionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
