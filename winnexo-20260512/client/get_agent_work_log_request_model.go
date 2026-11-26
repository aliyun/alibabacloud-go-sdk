// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentWorkLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetAgentWorkLogRequest
	GetAppToken() *string
	SetNodeId(v string) *GetAgentWorkLogRequest
	GetNodeId() *string
	SetOpsUserCode(v string) *GetAgentWorkLogRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetAgentWorkLogRequest
	GetOpsUserId() *string
	SetPage(v int64) *GetAgentWorkLogRequest
	GetPage() *int64
	SetPageSize(v int64) *GetAgentWorkLogRequest
	GetPageSize() *int64
	SetTenantId(v string) *GetAgentWorkLogRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetAgentWorkLogRequest
	GetWorkItemId() *string
}

type GetAgentWorkLogRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The optional node key.
	//
	// example:
	//
	// step_1
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The work item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetAgentWorkLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentWorkLogRequest) GoString() string {
	return s.String()
}

func (s *GetAgentWorkLogRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetAgentWorkLogRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAgentWorkLogRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetAgentWorkLogRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetAgentWorkLogRequest) GetPage() *int64 {
	return s.Page
}

func (s *GetAgentWorkLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAgentWorkLogRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAgentWorkLogRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetAgentWorkLogRequest) SetAppToken(v string) *GetAgentWorkLogRequest {
	s.AppToken = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetNodeId(v string) *GetAgentWorkLogRequest {
	s.NodeId = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetOpsUserCode(v string) *GetAgentWorkLogRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetOpsUserId(v string) *GetAgentWorkLogRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetPage(v int64) *GetAgentWorkLogRequest {
	s.Page = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetPageSize(v int64) *GetAgentWorkLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetTenantId(v string) *GetAgentWorkLogRequest {
	s.TenantId = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetWorkItemId(v string) *GetAgentWorkLogRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetAgentWorkLogRequest) Validate() error {
	return dara.Validate(s)
}
