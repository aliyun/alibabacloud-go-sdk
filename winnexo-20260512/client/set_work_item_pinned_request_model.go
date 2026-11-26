// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkItemPinnedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SetWorkItemPinnedRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *SetWorkItemPinnedRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *SetWorkItemPinnedRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SetWorkItemPinnedRequest
	GetOpsUserId() *string
	SetPinned(v bool) *SetWorkItemPinnedRequest
	GetPinned() *bool
	SetTenantId(v string) *SetWorkItemPinnedRequest
	GetTenantId() *string
	SetWorkItemId(v string) *SetWorkItemPinnedRequest
	GetWorkItemId() *string
}

type SetWorkItemPinnedRequest struct {
	// The platform application token issued by the console. If specified, the operation is performed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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
	// The pin status.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the work item.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s SetWorkItemPinnedRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWorkItemPinnedRequest) GoString() string {
	return s.String()
}

func (s *SetWorkItemPinnedRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SetWorkItemPinnedRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *SetWorkItemPinnedRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SetWorkItemPinnedRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SetWorkItemPinnedRequest) GetPinned() *bool {
	return s.Pinned
}

func (s *SetWorkItemPinnedRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SetWorkItemPinnedRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *SetWorkItemPinnedRequest) SetAppToken(v string) *SetWorkItemPinnedRequest {
	s.AppToken = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetCollaborationGroupId(v string) *SetWorkItemPinnedRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetOpsUserCode(v string) *SetWorkItemPinnedRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetOpsUserId(v string) *SetWorkItemPinnedRequest {
	s.OpsUserId = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetPinned(v bool) *SetWorkItemPinnedRequest {
	s.Pinned = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetTenantId(v string) *SetWorkItemPinnedRequest {
	s.TenantId = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetWorkItemId(v string) *SetWorkItemPinnedRequest {
	s.WorkItemId = &v
	return s
}

func (s *SetWorkItemPinnedRequest) Validate() error {
	return dara.Validate(s)
}
