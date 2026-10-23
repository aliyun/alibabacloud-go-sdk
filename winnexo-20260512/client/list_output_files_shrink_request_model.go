// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListOutputFilesShrinkRequest
	GetBizId() *string
	SetCollaborationGroupId(v string) *ListOutputFilesShrinkRequest
	GetCollaborationGroupId() *string
	SetItemType(v string) *ListOutputFilesShrinkRequest
	GetItemType() *string
	SetKeyword(v string) *ListOutputFilesShrinkRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListOutputFilesShrinkRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListOutputFilesShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListOutputFilesShrinkRequest
	GetPageSize() *int64
	SetSessionId(v string) *ListOutputFilesShrinkRequest
	GetSessionId() *string
	SetSharedOnly(v bool) *ListOutputFilesShrinkRequest
	GetSharedOnly() *bool
	SetTagsShrink(v string) *ListOutputFilesShrinkRequest
	GetTagsShrink() *string
	SetTenantId(v string) *ListOutputFilesShrinkRequest
	GetTenantId() *string
}

type ListOutputFilesShrinkRequest struct {
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
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used. The identity and tenant permissions are determined by the gateway and backend authentication.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOutputFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOutputFilesShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListOutputFilesShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListOutputFilesShrinkRequest) GetItemType() *string {
	return s.ItemType
}

func (s *ListOutputFilesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOutputFilesShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOutputFilesShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListOutputFilesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOutputFilesShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListOutputFilesShrinkRequest) GetSharedOnly() *bool {
	return s.SharedOnly
}

func (s *ListOutputFilesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListOutputFilesShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOutputFilesShrinkRequest) SetBizId(v string) *ListOutputFilesShrinkRequest {
	s.BizId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetCollaborationGroupId(v string) *ListOutputFilesShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetItemType(v string) *ListOutputFilesShrinkRequest {
	s.ItemType = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetKeyword(v string) *ListOutputFilesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetOperatingObjectName(v string) *ListOutputFilesShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetPage(v int64) *ListOutputFilesShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetPageSize(v int64) *ListOutputFilesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetSessionId(v string) *ListOutputFilesShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetSharedOnly(v bool) *ListOutputFilesShrinkRequest {
	s.SharedOnly = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetTagsShrink(v string) *ListOutputFilesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetTenantId(v string) *ListOutputFilesShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
