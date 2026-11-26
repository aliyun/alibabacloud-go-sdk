// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdminKnowledgeBasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetDirectoryId() *string
	SetKeyword(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListAdminKnowledgeBasesShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAdminKnowledgeBasesShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSortOrder() *string
	SetSourceTypesShrink(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetTenantId() *string
}

type ListAdminKnowledgeBasesShrinkRequest struct {
	// The platform application token issued by the console. If this parameter is passed, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The keyword for fuzzy search of component data in the form.
	//
	// example:
	//
	// Sample keyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The field by which to sort the returned results. Valid values:
	//
	// - event_time: the event creation time
	//
	// - event_execute_start_time: the event execution start time
	//
	// - event_execute_finish_time: the event execution finish time
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order. This parameter takes effect only when the sort field is specified. Valid values: ASC and DESC. The values are case-insensitive.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The list of service source types.
	//
	// example:
	//
	// ["K8S,SAE_K8S_SERVICE,FC3,VIP,DNS,AI,AGENT,CloudFlow,MSE_NACOS"]
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAdminKnowledgeBasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetAppToken(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetDirectoryId(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetKeyword(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetOpsUserCode(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetOpsUserId(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetPage(v int64) *ListAdminKnowledgeBasesShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetPageSize(v int64) *ListAdminKnowledgeBasesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSortField(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSortOrder(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSourceTypesShrink(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetTenantId(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
