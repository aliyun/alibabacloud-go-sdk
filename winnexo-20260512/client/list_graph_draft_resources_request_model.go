// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphDraftResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListGraphDraftResourcesRequest
	GetAppToken() *string
	SetGraphName(v string) *ListGraphDraftResourcesRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *ListGraphDraftResourcesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListGraphDraftResourcesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ListGraphDraftResourcesRequest
	GetTenantId() *string
}

type ListGraphDraftResourcesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The graph name.
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
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphDraftResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphDraftResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListGraphDraftResourcesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListGraphDraftResourcesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ListGraphDraftResourcesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListGraphDraftResourcesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListGraphDraftResourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphDraftResourcesRequest) SetAppToken(v string) *ListGraphDraftResourcesRequest {
	s.AppToken = &v
	return s
}

func (s *ListGraphDraftResourcesRequest) SetGraphName(v string) *ListGraphDraftResourcesRequest {
	s.GraphName = &v
	return s
}

func (s *ListGraphDraftResourcesRequest) SetOpsUserCode(v string) *ListGraphDraftResourcesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListGraphDraftResourcesRequest) SetOpsUserId(v string) *ListGraphDraftResourcesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListGraphDraftResourcesRequest) SetTenantId(v string) *ListGraphDraftResourcesRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphDraftResourcesRequest) Validate() error {
	return dara.Validate(s)
}
