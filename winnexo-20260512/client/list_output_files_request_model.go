// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListOutputFilesRequest
	GetBizId() *string
	SetCollaborationGroupId(v string) *ListOutputFilesRequest
	GetCollaborationGroupId() *string
	SetItemType(v string) *ListOutputFilesRequest
	GetItemType() *string
	SetKeyword(v string) *ListOutputFilesRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListOutputFilesRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListOutputFilesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListOutputFilesRequest
	GetPageSize() *int64
	SetSessionId(v string) *ListOutputFilesRequest
	GetSessionId() *string
	SetSharedOnly(v bool) *ListOutputFilesRequest
	GetSharedOnly() *bool
	SetTags(v []*string) *ListOutputFilesRequest
	GetTags() []*string
	SetTenantId(v string) *ListOutputFilesRequest
	GetTenantId() *string
}

type ListOutputFilesRequest struct {
	// The business ID used for exact filtering. The length is 1 to 128 characters. This value is derived from the businessId of the file registration or associated workflow instance.
	//
	// example:
	//
	// REQ-20261012-001
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// The ID of the collaboration workspace. The length is 1 to 64 characters. This parameter is optional. If not specified, the personal outputs of the current user are queried. If specified, the real-time permissions of the workspace members and source resources are verified.
	//
	// example:
	//
	// cg_123e4567-e89b-42d3-a456-426614174001
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The output item type. Valid values: ppt, html, document, picture, slides, video, audio, email, and others.
	//
	// example:
	//
	// document
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The keyword used to match the output title or detail name.
	//
	// example:
	//
	// Quarterly operations
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The name of the digital employee (operating object). Filters outputs by name.
	//
	// example:
	//
	// Business Analysis Assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of items per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The session ID used for exact filtering. The length is 1 to 128 characters. This parameter applies to regular sessions, skill sessions, and the executionId of task executions.
	//
	// example:
	//
	// session_123e4567-e89b-42d3-a456-426614174008
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to display only shared outputs and output details.
	//
	// example:
	//
	// False
	SharedOnly *bool `json:"sharedOnly,omitempty" xml:"sharedOnly,omitempty"`
	// The tags used for exact filtering. All specified tags must be matched simultaneously. You can specify up to 20 tags. Each tag must be a non-empty string of 1 to 128 characters. Tags are derived from file registrations or the variables.tags of associated instances.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used. The identity and tenant permissions are determined by the gateway and backend authentication.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOutputFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesRequest) GoString() string {
	return s.String()
}

func (s *ListOutputFilesRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListOutputFilesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListOutputFilesRequest) GetItemType() *string {
	return s.ItemType
}

func (s *ListOutputFilesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOutputFilesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOutputFilesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListOutputFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOutputFilesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListOutputFilesRequest) GetSharedOnly() *bool {
	return s.SharedOnly
}

func (s *ListOutputFilesRequest) GetTags() []*string {
	return s.Tags
}

func (s *ListOutputFilesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOutputFilesRequest) SetBizId(v string) *ListOutputFilesRequest {
	s.BizId = &v
	return s
}

func (s *ListOutputFilesRequest) SetCollaborationGroupId(v string) *ListOutputFilesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListOutputFilesRequest) SetItemType(v string) *ListOutputFilesRequest {
	s.ItemType = &v
	return s
}

func (s *ListOutputFilesRequest) SetKeyword(v string) *ListOutputFilesRequest {
	s.Keyword = &v
	return s
}

func (s *ListOutputFilesRequest) SetOperatingObjectName(v string) *ListOutputFilesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOutputFilesRequest) SetPage(v int64) *ListOutputFilesRequest {
	s.Page = &v
	return s
}

func (s *ListOutputFilesRequest) SetPageSize(v int64) *ListOutputFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutputFilesRequest) SetSessionId(v string) *ListOutputFilesRequest {
	s.SessionId = &v
	return s
}

func (s *ListOutputFilesRequest) SetSharedOnly(v bool) *ListOutputFilesRequest {
	s.SharedOnly = &v
	return s
}

func (s *ListOutputFilesRequest) SetTags(v []*string) *ListOutputFilesRequest {
	s.Tags = v
	return s
}

func (s *ListOutputFilesRequest) SetTenantId(v string) *ListOutputFilesRequest {
	s.TenantId = &v
	return s
}

func (s *ListOutputFilesRequest) Validate() error {
	return dara.Validate(s)
}
