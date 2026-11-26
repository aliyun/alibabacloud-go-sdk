// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdminKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListAdminKnowledgeBasesRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ListAdminKnowledgeBasesRequest
	GetDirectoryId() *string
	SetKeyword(v string) *ListAdminKnowledgeBasesRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListAdminKnowledgeBasesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListAdminKnowledgeBasesRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListAdminKnowledgeBasesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAdminKnowledgeBasesRequest
	GetPageSize() *int64
	SetSortField(v string) *ListAdminKnowledgeBasesRequest
	GetSortField() *string
	SetSortOrder(v string) *ListAdminKnowledgeBasesRequest
	GetSortOrder() *string
	SetSourceTypes(v []*string) *ListAdminKnowledgeBasesRequest
	GetSourceTypes() []*string
	SetTenantId(v string) *ListAdminKnowledgeBasesRequest
	GetTenantId() *string
}

type ListAdminKnowledgeBasesRequest struct {
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
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAdminKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListAdminKnowledgeBasesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAdminKnowledgeBasesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAdminKnowledgeBasesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListAdminKnowledgeBasesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListAdminKnowledgeBasesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAdminKnowledgeBasesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdminKnowledgeBasesRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListAdminKnowledgeBasesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAdminKnowledgeBasesRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListAdminKnowledgeBasesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAdminKnowledgeBasesRequest) SetAppToken(v string) *ListAdminKnowledgeBasesRequest {
	s.AppToken = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetDirectoryId(v string) *ListAdminKnowledgeBasesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetKeyword(v string) *ListAdminKnowledgeBasesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetOpsUserCode(v string) *ListAdminKnowledgeBasesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetOpsUserId(v string) *ListAdminKnowledgeBasesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetPage(v int64) *ListAdminKnowledgeBasesRequest {
	s.Page = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetPageSize(v int64) *ListAdminKnowledgeBasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetSortField(v string) *ListAdminKnowledgeBasesRequest {
	s.SortField = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetSortOrder(v string) *ListAdminKnowledgeBasesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetSourceTypes(v []*string) *ListAdminKnowledgeBasesRequest {
	s.SourceTypes = v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) SetTenantId(v string) *ListAdminKnowledgeBasesRequest {
	s.TenantId = &v
	return s
}

func (s *ListAdminKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
