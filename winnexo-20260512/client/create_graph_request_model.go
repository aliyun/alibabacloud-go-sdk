// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGraphRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessProfile(v string) *CreateGraphRequest
	GetBusinessProfile() *string
	SetDataSourceId(v int64) *CreateGraphRequest
	GetDataSourceId() *int64
	SetDisplayName(v string) *CreateGraphRequest
	GetDisplayName() *string
	SetGraphName(v string) *CreateGraphRequest
	GetGraphName() *string
	SetTenantId(v string) *CreateGraphRequest
	GetTenantId() *string
}

type CreateGraphRequest struct {
	// 业务说明（可选）
	//
	// example:
	//
	// 客户域语义图谱
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// 绑定的数据源 ID（控制台已创建的 RDB 类数据源）
	//
	// This parameter is required.
	//
	// example:
	//
	// 198001
	DataSourceId *int64 `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// 图谱展示名（可选，租户内大小写不敏感唯一，最多200字）
	//
	// example:
	//
	// CRM 图谱
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 图谱名称，字母开头+字母/数字/下划线，长度不超过64，租户内唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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

func (s *CreateGraphRequest) GetTenantId() *string {
	return s.TenantId
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

func (s *CreateGraphRequest) SetTenantId(v string) *CreateGraphRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGraphRequest) Validate() error {
	return dara.Validate(s)
}
