// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListKnowledgeBaseDirectoriesRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ListKnowledgeBaseDirectoriesRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *ListKnowledgeBaseDirectoriesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListKnowledgeBaseDirectoriesRequest
	GetOpsUserId() *string
	SetSortField(v string) *ListKnowledgeBaseDirectoriesRequest
	GetSortField() *string
	SetSortOrder(v string) *ListKnowledgeBaseDirectoriesRequest
	GetSortOrder() *string
	SetTenantId(v string) *ListKnowledgeBaseDirectoriesRequest
	GetTenantId() *string
}

type ListKnowledgeBaseDirectoriesRequest struct {
	// The platform application token issued by the console. If this parameter is passed, the request is executed as the user specified by opsUserId or opsUserCode.
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
	// The field by which the returned results are sorted. Valid values:
	//
	// - event_time: the event creation time
	//
	// - event_execute_start_time: the event execution time
	//
	// - event_execute_finish_time: the event completion time
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order. This parameter takes effect only when sortBy is specified. Valid values: ASC and DESC. The values are case-insensitive.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The tenant ID. This is a common parameter. For winnexo-cli, this parameter is explicitly passed by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListKnowledgeBaseDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetAppToken(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.AppToken = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetDirectoryId(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetOpsUserCode(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetOpsUserId(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetSortField(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.SortField = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetSortOrder(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetTenantId(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.TenantId = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
