// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListSkillsRequest
	GetAppToken() *string
	SetBindStatus(v string) *ListSkillsRequest
	GetBindStatus() *string
	SetFilterType(v string) *ListSkillsRequest
	GetFilterType() *string
	SetKeyword(v string) *ListSkillsRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListSkillsRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *ListSkillsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListSkillsRequest
	GetOpsUserId() *string
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
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
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

func (s *ListSkillsRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *ListSkillsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListSkillsRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *ListSkillsRequest) SetAppToken(v string) *ListSkillsRequest {
	s.AppToken = &v
	return s
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

func (s *ListSkillsRequest) SetOpsUserCode(v string) *ListSkillsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListSkillsRequest) SetOpsUserId(v string) *ListSkillsRequest {
	s.OpsUserId = &v
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
