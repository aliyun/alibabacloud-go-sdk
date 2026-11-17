// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputItemId(v string) *UpdateShareRequest
	GetOutputItemId() *string
	SetOutputItemName(v string) *UpdateShareRequest
	GetOutputItemName() *string
	SetRefreshShareLink(v bool) *UpdateShareRequest
	GetRefreshShareLink() *bool
	SetSessionId(v string) *UpdateShareRequest
	GetSessionId() *string
	SetShareAccessType(v string) *UpdateShareRequest
	GetShareAccessType() *string
	SetShareEnabled(v bool) *UpdateShareRequest
	GetShareEnabled() *bool
	SetShareUserGroupIds(v []*string) *UpdateShareRequest
	GetShareUserGroupIds() []*string
	SetShareUserIds(v []*int64) *UpdateShareRequest
	GetShareUserIds() []*int64
	SetTargetType(v string) *UpdateShareRequest
	GetTargetType() *string
	SetTenantId(v string) *UpdateShareRequest
	GetTenantId() *string
}

type UpdateShareRequest struct {
	// The output item ID.
	//
	// example:
	//
	// exampleOutputItemId
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// The name of the output item. This parameter must be used together with sessionId.
	//
	// example:
	//
	// Quarterly Review.pptx
	OutputItemName *string `json:"outputItemName,omitempty" xml:"outputItemName,omitempty"`
	// Specifies whether to refresh the sharing link. This is supported only for output items that have sharing enabled or are being enabled in the current request.
	//
	// example:
	//
	// false
	RefreshShareLink *bool `json:"refreshShareLink,omitempty" xml:"refreshShareLink,omitempty"`
	// The session ID.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The sharing scope. Valid values:
	//
	// - PUBLIC: Accessible to all users.
	//
	// - SPECIFIED: Accessible to specified users or user groups.
	//
	// example:
	//
	// SPECIFIED
	ShareAccessType *string `json:"shareAccessType,omitempty" xml:"shareAccessType,omitempty"`
	// Specifies whether sharing is enabled.
	//
	// example:
	//
	// true
	ShareEnabled *bool `json:"shareEnabled,omitempty" xml:"shareEnabled,omitempty"`
	// The list of user group IDs within the specified scope. Pass an empty array to clear the list.
	//
	// example:
	//
	// ["group-1"]
	ShareUserGroupIds []*string `json:"shareUserGroupIds,omitempty" xml:"shareUserGroupIds,omitempty" type:"Repeated"`
	// The list of platform user IDs within the specified scope. Pass an empty array to clear the list.
	//
	// example:
	//
	// [10001]
	ShareUserIds []*int64 `json:"shareUserIds,omitempty" xml:"shareUserIds,omitempty" type:"Repeated"`
	// The type of the current permission. Valid values:
	//
	// - period: Period type.
	//
	// - objective: Objective type.
	//
	// - kr: Key result type.
	//
	// This parameter is required.
	//
	// example:
	//
	// outputItem
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateShareRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateShareRequest) GoString() string {
	return s.String()
}

func (s *UpdateShareRequest) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *UpdateShareRequest) GetOutputItemName() *string {
	return s.OutputItemName
}

func (s *UpdateShareRequest) GetRefreshShareLink() *bool {
	return s.RefreshShareLink
}

func (s *UpdateShareRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateShareRequest) GetShareAccessType() *string {
	return s.ShareAccessType
}

func (s *UpdateShareRequest) GetShareEnabled() *bool {
	return s.ShareEnabled
}

func (s *UpdateShareRequest) GetShareUserGroupIds() []*string {
	return s.ShareUserGroupIds
}

func (s *UpdateShareRequest) GetShareUserIds() []*int64 {
	return s.ShareUserIds
}

func (s *UpdateShareRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateShareRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateShareRequest) SetOutputItemId(v string) *UpdateShareRequest {
	s.OutputItemId = &v
	return s
}

func (s *UpdateShareRequest) SetOutputItemName(v string) *UpdateShareRequest {
	s.OutputItemName = &v
	return s
}

func (s *UpdateShareRequest) SetRefreshShareLink(v bool) *UpdateShareRequest {
	s.RefreshShareLink = &v
	return s
}

func (s *UpdateShareRequest) SetSessionId(v string) *UpdateShareRequest {
	s.SessionId = &v
	return s
}

func (s *UpdateShareRequest) SetShareAccessType(v string) *UpdateShareRequest {
	s.ShareAccessType = &v
	return s
}

func (s *UpdateShareRequest) SetShareEnabled(v bool) *UpdateShareRequest {
	s.ShareEnabled = &v
	return s
}

func (s *UpdateShareRequest) SetShareUserGroupIds(v []*string) *UpdateShareRequest {
	s.ShareUserGroupIds = v
	return s
}

func (s *UpdateShareRequest) SetShareUserIds(v []*int64) *UpdateShareRequest {
	s.ShareUserIds = v
	return s
}

func (s *UpdateShareRequest) SetTargetType(v string) *UpdateShareRequest {
	s.TargetType = &v
	return s
}

func (s *UpdateShareRequest) SetTenantId(v string) *UpdateShareRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateShareRequest) Validate() error {
	return dara.Validate(s)
}
