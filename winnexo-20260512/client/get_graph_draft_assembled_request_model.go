// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphDraftAssembledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetGraphDraftAssembledRequest
	GetAppToken() *string
	SetGraphName(v string) *GetGraphDraftAssembledRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *GetGraphDraftAssembledRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetGraphDraftAssembledRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetGraphDraftAssembledRequest
	GetTenantId() *string
}

type GetGraphDraftAssembledRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The knowledge graph name.
	//
	// This parameter is required.
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
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGraphDraftAssembledRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGraphDraftAssembledRequest) GoString() string {
	return s.String()
}

func (s *GetGraphDraftAssembledRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetGraphDraftAssembledRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphDraftAssembledRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetGraphDraftAssembledRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetGraphDraftAssembledRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGraphDraftAssembledRequest) SetAppToken(v string) *GetGraphDraftAssembledRequest {
	s.AppToken = &v
	return s
}

func (s *GetGraphDraftAssembledRequest) SetGraphName(v string) *GetGraphDraftAssembledRequest {
	s.GraphName = &v
	return s
}

func (s *GetGraphDraftAssembledRequest) SetOpsUserCode(v string) *GetGraphDraftAssembledRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetGraphDraftAssembledRequest) SetOpsUserId(v string) *GetGraphDraftAssembledRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetGraphDraftAssembledRequest) SetTenantId(v string) *GetGraphDraftAssembledRequest {
	s.TenantId = &v
	return s
}

func (s *GetGraphDraftAssembledRequest) Validate() error {
	return dara.Validate(s)
}
