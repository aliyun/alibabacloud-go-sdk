// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListSkillsShrinkRequest
	GetAppToken() *string
	SetBindStatus(v string) *ListSkillsShrinkRequest
	GetBindStatus() *string
	SetFilterType(v string) *ListSkillsShrinkRequest
	GetFilterType() *string
	SetKeyword(v string) *ListSkillsShrinkRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListSkillsShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *ListSkillsShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListSkillsShrinkRequest
	GetOpsUserId() *string
	SetPage(v int32) *ListSkillsShrinkRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSkillsShrinkRequest
	GetPageSize() *int32
	SetTagsShrink(v string) *ListSkillsShrinkRequest
	GetTagsShrink() *string
	SetTenantId(v string) *ListSkillsShrinkRequest
	GetTenantId() *string
}

type ListSkillsShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The binding status. Valid values: BOUND (bound) and UNBOUND (unbound global skills). Must be specified together with operatingObjectName.
	//
	// example:
	//
	// BOUND
	BindStatus *string `json:"bindStatus,omitempty" xml:"bindStatus,omitempty"`
	// The filter expression type. Valid values:
	//
	// - SQL: SQL filtering.
	//
	// - TAG: Tag filtering.
	//
	// example:
	//
	// ALL
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// The search keyword. Supports fuzzy search by API name or exact search by API ID.
	//
	// example:
	//
	// Sample keyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The name of the digital employee. Used to calculate the CodeAgent allowedSkills whitelist based on binding relationships.
	//
	// example:
	//
	// 11111
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The proxy user code, which is unique within the tenant. Used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. Used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tag filtering parameter.
	//
	// example:
	//
	// ["[{\\\"value\\\":\\\"dataworks#0#1#34167#432313#T3_23030136354#86969667366\\\",\\\"key\\\":\\\"dataworksTag\\\"}]"]
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListSkillsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListSkillsShrinkRequest) GetBindStatus() *string {
	return s.BindStatus
}

func (s *ListSkillsShrinkRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *ListSkillsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSkillsShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListSkillsShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListSkillsShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListSkillsShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSkillsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListSkillsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListSkillsShrinkRequest) SetAppToken(v string) *ListSkillsShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetBindStatus(v string) *ListSkillsShrinkRequest {
	s.BindStatus = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetFilterType(v string) *ListSkillsShrinkRequest {
	s.FilterType = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetKeyword(v string) *ListSkillsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetOperatingObjectName(v string) *ListSkillsShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetOpsUserCode(v string) *ListSkillsShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetOpsUserId(v string) *ListSkillsShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetPage(v int32) *ListSkillsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetPageSize(v int32) *ListSkillsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetTagsShrink(v string) *ListSkillsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetTenantId(v string) *ListSkillsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListSkillsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
