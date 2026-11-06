// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentVersionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAgentVersionsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListAgentVersionsResponseBodyItems) *ListAgentVersionsResponseBody
	GetItems() []*ListAgentVersionsResponseBodyItems
	SetMaxResults(v int32) *ListAgentVersionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAgentVersionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAgentVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgentVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentVersionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgentVersionsResponseBody
	GetTotalCount() *int64
}

type ListAgentVersionsResponseBody struct {
	// 业务状态码，请求成功时为 SUCCESS，失败时为错误码。
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP 状态码。
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// 当前页的固定版本列表，按创建时间倒序排列，不包含 configuration。
	Items []*ListAgentVersionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 本次查询使用的单页最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 响应消息，请求成功时为 success，失败时为错误信息。
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 下一页的分页令牌，查询下一页时作为请求参数 nextToken 传入；为空表示没有更多数据。
	//
	// example:
	//
	// YWdlbnQtdmVyc2lvbnM6MjA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求 ID，用于定位和排查请求。
	//
	// example:
	//
	// 5E8F3C2A-1B4D-4F6E-9A7C-0D2B8E6F4A1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 请求是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 该智能体未删除的固定版本总数，不包含系统维护的 latest 版本。
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAgentVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentVersionsResponseBody) GetItems() []*ListAgentVersionsResponseBodyItems {
	return s.Items
}

func (s *ListAgentVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgentVersionsResponseBody) SetCode(v string) *ListAgentVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetHttpStatusCode(v int32) *ListAgentVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetItems(v []*ListAgentVersionsResponseBodyItems) *ListAgentVersionsResponseBody {
	s.Items = v
	return s
}

func (s *ListAgentVersionsResponseBody) SetMaxResults(v int32) *ListAgentVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetMessage(v string) *ListAgentVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetNextToken(v string) *ListAgentVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetRequestId(v string) *ListAgentVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetSuccess(v bool) *ListAgentVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetTotalCount(v int64) *ListAgentVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentVersionsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentVersionsResponseBodyItems struct {
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

func (s ListAgentVersionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAgentVersionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAgentVersionsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentVersionsResponseBodyItems) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *ListAgentVersionsResponseBodyItems) GetArtifactStatus() *string {
	return s.ArtifactStatus
}

func (s *ListAgentVersionsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAgentVersionsResponseBodyItems) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListAgentVersionsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListAgentVersionsResponseBodyItems) GetTargetSpecVersion() *int64 {
	return s.TargetSpecVersion
}

func (s *ListAgentVersionsResponseBodyItems) GetVersionName() *string {
	return s.VersionName
}

func (s *ListAgentVersionsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAgentVersionsResponseBodyItems) SetAgentId(v string) *ListAgentVersionsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetAgentVersion(v string) *ListAgentVersionsResponseBodyItems {
	s.AgentVersion = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetArtifactStatus(v string) *ListAgentVersionsResponseBodyItems {
	s.ArtifactStatus = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetCreatedAt(v string) *ListAgentVersionsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetCreatedBy(v string) *ListAgentVersionsResponseBodyItems {
	s.CreatedBy = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetDescription(v string) *ListAgentVersionsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetTargetSpecVersion(v int64) *ListAgentVersionsResponseBodyItems {
	s.TargetSpecVersion = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetVersionName(v string) *ListAgentVersionsResponseBodyItems {
	s.VersionName = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetWorkspaceId(v string) *ListAgentVersionsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
