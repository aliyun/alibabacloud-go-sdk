// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertGraphDraftResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RevertGraphDraftResourceRequest
	GetAppToken() *string
	SetDraftChangeId(v int64) *RevertGraphDraftResourceRequest
	GetDraftChangeId() *int64
	SetGraphName(v string) *RevertGraphDraftResourceRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *RevertGraphDraftResourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RevertGraphDraftResourceRequest
	GetOpsUserId() *string
	SetTenantId(v string) *RevertGraphDraftResourceRequest
	GetTenantId() *string
}

type RevertGraphDraftResourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The draft change ID (the draftChangeId returned by listGraphDraftResources).
	//
	// This parameter is required.
	//
	// example:
	//
	// 401001
	DraftChangeId *int64 `json:"draftChangeId,omitempty" xml:"draftChangeId,omitempty"`
	// The knowledge graph name.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
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
	// The tenant ID. This is a common parameter. Pass it explicitly by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RevertGraphDraftResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RevertGraphDraftResourceRequest) GoString() string {
	return s.String()
}

func (s *RevertGraphDraftResourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RevertGraphDraftResourceRequest) GetDraftChangeId() *int64 {
	return s.DraftChangeId
}

func (s *RevertGraphDraftResourceRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *RevertGraphDraftResourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RevertGraphDraftResourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RevertGraphDraftResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevertGraphDraftResourceRequest) SetAppToken(v string) *RevertGraphDraftResourceRequest {
	s.AppToken = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) SetDraftChangeId(v int64) *RevertGraphDraftResourceRequest {
	s.DraftChangeId = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) SetGraphName(v string) *RevertGraphDraftResourceRequest {
	s.GraphName = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) SetOpsUserCode(v string) *RevertGraphDraftResourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) SetOpsUserId(v string) *RevertGraphDraftResourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) SetTenantId(v string) *RevertGraphDraftResourceRequest {
	s.TenantId = &v
	return s
}

func (s *RevertGraphDraftResourceRequest) Validate() error {
	return dara.Validate(s)
}
