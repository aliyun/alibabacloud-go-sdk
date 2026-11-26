// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGraphRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateGraphRequest
	GetAppToken() *string
	SetBusinessProfile(v string) *CreateGraphRequest
	GetBusinessProfile() *string
	SetDataSourceId(v int64) *CreateGraphRequest
	GetDataSourceId() *int64
	SetDisplayName(v string) *CreateGraphRequest
	GetDisplayName() *string
	SetGraphName(v string) *CreateGraphRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *CreateGraphRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateGraphRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CreateGraphRequest
	GetTenantId() *string
}

type CreateGraphRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The business description of the graph. If not configured, it is an empty string.
	//
	// example:
	//
	// Customer domain semantic graph
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 198001
	DataSourceId *int64 `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// CRM Graph
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
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
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGraphRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGraphRequest) GoString() string {
	return s.String()
}

func (s *CreateGraphRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateGraphRequest) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *CreateGraphRequest) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateGraphRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateGraphRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *CreateGraphRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateGraphRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateGraphRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGraphRequest) SetAppToken(v string) *CreateGraphRequest {
	s.AppToken = &v
	return s
}

func (s *CreateGraphRequest) SetBusinessProfile(v string) *CreateGraphRequest {
	s.BusinessProfile = &v
	return s
}

func (s *CreateGraphRequest) SetDataSourceId(v int64) *CreateGraphRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateGraphRequest) SetDisplayName(v string) *CreateGraphRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateGraphRequest) SetGraphName(v string) *CreateGraphRequest {
	s.GraphName = &v
	return s
}

func (s *CreateGraphRequest) SetOpsUserCode(v string) *CreateGraphRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateGraphRequest) SetOpsUserId(v string) *CreateGraphRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateGraphRequest) SetTenantId(v string) *CreateGraphRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGraphRequest) Validate() error {
	return dara.Validate(s)
}
