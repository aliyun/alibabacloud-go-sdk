// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemRunSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetWorkItemRunSummariesRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *GetWorkItemRunSummariesRequest
	GetCollaborationGroupId() *string
	SetNodeKey(v string) *GetWorkItemRunSummariesRequest
	GetNodeKey() *string
	SetOpsUserCode(v string) *GetWorkItemRunSummariesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetWorkItemRunSummariesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetWorkItemRunSummariesRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemRunSummariesRequest
	GetWorkItemId() *string
}

type GetWorkItemRunSummariesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The node key.
	//
	// example:
	//
	// query.amount
	NodeKey *string `json:"nodeKey,omitempty" xml:"nodeKey,omitempty"`
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
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the work item.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_0123456789ab4cde
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetWorkItemRunSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRunSummariesRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemRunSummariesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetWorkItemRunSummariesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemRunSummariesRequest) GetNodeKey() *string {
	return s.NodeKey
}

func (s *GetWorkItemRunSummariesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetWorkItemRunSummariesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetWorkItemRunSummariesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemRunSummariesRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemRunSummariesRequest) SetAppToken(v string) *GetWorkItemRunSummariesRequest {
	s.AppToken = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetCollaborationGroupId(v string) *GetWorkItemRunSummariesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetNodeKey(v string) *GetWorkItemRunSummariesRequest {
	s.NodeKey = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetOpsUserCode(v string) *GetWorkItemRunSummariesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetOpsUserId(v string) *GetWorkItemRunSummariesRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetTenantId(v string) *GetWorkItemRunSummariesRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetWorkItemId(v string) *GetWorkItemRunSummariesRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) Validate() error {
	return dara.Validate(s)
}
