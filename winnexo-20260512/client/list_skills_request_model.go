// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindStatus(v string) *ListSkillsRequest
	GetBindStatus() *string
	SetFilterType(v string) *ListSkillsRequest
	GetFilterType() *string
	SetKeyword(v string) *ListSkillsRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListSkillsRequest
	GetOperatingObjectName() *string
	SetPage(v int32) *ListSkillsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSkillsRequest
	GetPageSize() *int32
	SetTags(v []*string) *ListSkillsRequest
	GetTags() []*string
	SetTenantId(v string) *ListSkillsRequest
	GetTenantId() *string
}

type ListSkillsRequest struct {
	// 绑定状态：BOUND(已绑定) / UNBOUND(未绑定的全局技能)；必须与 operatingObjectName 同时传入
	//
	// example:
	//
	// BOUND
	BindStatus *string `json:"bindStatus,omitempty" xml:"bindStatus,omitempty"`
	// 技能筛选维度：ALL/BUILTIN/CUSTOM/DRAFT/ALL_WITH_DRAFTS
	//
	// example:
	//
	// ALL
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// 按技能名称或描述模糊匹配
	//
	// example:
	//
	// 示例关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 数字员工名称；必须与 bindStatus 同时传入
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 页码，从 1 开始
	//
	// example:
	//
	// string_value
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量，范围 1-100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 按标签过滤，数组任一命中即匹配
	//
	// example:
	//
	// string_value
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetBindStatus() *string {
	return s.BindStatus
}

func (s *ListSkillsRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *ListSkillsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSkillsRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListSkillsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSkillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsRequest) GetTags() []*string {
	return s.Tags
}

func (s *ListSkillsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListSkillsRequest) SetBindStatus(v string) *ListSkillsRequest {
	s.BindStatus = &v
	return s
}

func (s *ListSkillsRequest) SetFilterType(v string) *ListSkillsRequest {
	s.FilterType = &v
	return s
}

func (s *ListSkillsRequest) SetKeyword(v string) *ListSkillsRequest {
	s.Keyword = &v
	return s
}

func (s *ListSkillsRequest) SetOperatingObjectName(v string) *ListSkillsRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListSkillsRequest) SetPage(v int32) *ListSkillsRequest {
	s.Page = &v
	return s
}

func (s *ListSkillsRequest) SetPageSize(v int32) *ListSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsRequest) SetTags(v []*string) *ListSkillsRequest {
	s.Tags = v
	return s
}

func (s *ListSkillsRequest) SetTenantId(v string) *ListSkillsRequest {
	s.TenantId = &v
	return s
}

func (s *ListSkillsRequest) Validate() error {
	return dara.Validate(s)
}
