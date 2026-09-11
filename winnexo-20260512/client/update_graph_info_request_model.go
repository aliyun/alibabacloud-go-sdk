// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGraphInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessProfile(v string) *UpdateGraphInfoRequest
	GetBusinessProfile() *string
	SetDisplayName(v string) *UpdateGraphInfoRequest
	GetDisplayName() *string
	SetGraphName(v string) *UpdateGraphInfoRequest
	GetGraphName() *string
	SetTenantId(v string) *UpdateGraphInfoRequest
	GetTenantId() *string
}

type UpdateGraphInfoRequest struct {
	// 业务说明（可选；传空串表示清空；与 displayName 至少传其一）
	//
	// example:
	//
	// 客户域语义图谱
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// 图谱展示名（可选，最多200字；传空串或纯空白会被拒绝；与 businessProfile 至少传其一）
	//
	// example:
	//
	// CRM 图谱
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 图谱名称
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

func (s UpdateGraphInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGraphInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateGraphInfoRequest) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *UpdateGraphInfoRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateGraphInfoRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *UpdateGraphInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGraphInfoRequest) SetBusinessProfile(v string) *UpdateGraphInfoRequest {
	s.BusinessProfile = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetDisplayName(v string) *UpdateGraphInfoRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetGraphName(v string) *UpdateGraphInfoRequest {
	s.GraphName = &v
	return s
}

func (s *UpdateGraphInfoRequest) SetTenantId(v string) *UpdateGraphInfoRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGraphInfoRequest) Validate() error {
	return dara.Validate(s)
}
