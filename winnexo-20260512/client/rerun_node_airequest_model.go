// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunNodeAIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RerunNodeAIRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *RerunNodeAIRequest
	GetCollaborationGroupId() *string
	SetInstanceId(v string) *RerunNodeAIRequest
	GetInstanceId() *string
	SetNodeId(v string) *RerunNodeAIRequest
	GetNodeId() *string
	SetOpsUserCode(v string) *RerunNodeAIRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RerunNodeAIRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RerunNodeAIRequest
	GetTenantId() *string
}

type RerunNodeAIRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The physical workflow instance ID (planId).
	//
	// This parameter is required.
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The key of the failed node (nodeKey).
	//
	// This parameter is required.
	//
	// example:
	//
	// query.amount
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
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RerunNodeAIRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunNodeAIRequest) GoString() string {
	return s.String()
}

func (s *RerunNodeAIRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RerunNodeAIRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *RerunNodeAIRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RerunNodeAIRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *RerunNodeAIRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RerunNodeAIRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RerunNodeAIRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RerunNodeAIRequest) SetAppToken(v string) *RerunNodeAIRequest {
	s.AppToken = &v
	return s
}

func (s *RerunNodeAIRequest) SetCollaborationGroupId(v string) *RerunNodeAIRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *RerunNodeAIRequest) SetInstanceId(v string) *RerunNodeAIRequest {
	s.InstanceId = &v
	return s
}

func (s *RerunNodeAIRequest) SetNodeId(v string) *RerunNodeAIRequest {
	s.NodeId = &v
	return s
}

func (s *RerunNodeAIRequest) SetOpsUserCode(v string) *RerunNodeAIRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RerunNodeAIRequest) SetOpsUserId(v string) *RerunNodeAIRequest {
	s.OpsUserId = &v
	return s
}

func (s *RerunNodeAIRequest) SetTenantId(v string) *RerunNodeAIRequest {
	s.TenantId = &v
	return s
}

func (s *RerunNodeAIRequest) Validate() error {
	return dara.Validate(s)
}
