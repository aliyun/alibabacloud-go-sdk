// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentVersionResponseBody
	GetCode() *string
	SetData(v *CreateAgentVersionResponseBodyData) *CreateAgentVersionResponseBody
	GetData() *CreateAgentVersionResponseBodyData
	SetHttpStatusCode(v int32) *CreateAgentVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAgentVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentVersionResponseBody
	GetSuccess() *bool
}

type CreateAgentVersionResponseBody struct {
	// 业务状态码，请求成功时为 SUCCESS，失败时为错误码。
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 已保存的智能体固定版本详情；重复提交命中已有版本时返回该已有版本。
	Data *CreateAgentVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateAgentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentVersionResponseBody) GetData() *CreateAgentVersionResponseBodyData {
	return s.Data
}

func (s *CreateAgentVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAgentVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentVersionResponseBody) SetCode(v string) *CreateAgentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentVersionResponseBody) SetData(v *CreateAgentVersionResponseBodyData) *CreateAgentVersionResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgentVersionResponseBody) SetHttpStatusCode(v int32) *CreateAgentVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAgentVersionResponseBody) SetMessage(v string) *CreateAgentVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentVersionResponseBody) SetRequestId(v string) *CreateAgentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentVersionResponseBody) SetSuccess(v bool) *CreateAgentVersionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentVersionResponseBodyData struct {
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

func (s CreateAgentVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentVersionResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAgentVersionResponseBodyData) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *CreateAgentVersionResponseBodyData) GetArtifactStatus() *string {
	return s.ArtifactStatus
}

func (s *CreateAgentVersionResponseBodyData) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *CreateAgentVersionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateAgentVersionResponseBodyData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateAgentVersionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentVersionResponseBodyData) GetTargetSpecVersion() *int64 {
	return s.TargetSpecVersion
}

func (s *CreateAgentVersionResponseBodyData) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateAgentVersionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAgentVersionResponseBodyData) SetAgentId(v string) *CreateAgentVersionResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetAgentVersion(v string) *CreateAgentVersionResponseBodyData {
	s.AgentVersion = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetArtifactStatus(v string) *CreateAgentVersionResponseBodyData {
	s.ArtifactStatus = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetConfiguration(v map[string]interface{}) *CreateAgentVersionResponseBodyData {
	s.Configuration = v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetCreatedAt(v string) *CreateAgentVersionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetCreatedBy(v string) *CreateAgentVersionResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetDescription(v string) *CreateAgentVersionResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetTargetSpecVersion(v int64) *CreateAgentVersionResponseBodyData {
	s.TargetSpecVersion = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetVersionName(v string) *CreateAgentVersionResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) SetWorkspaceId(v string) *CreateAgentVersionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAgentVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
