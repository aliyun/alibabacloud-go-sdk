// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupSourceContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateGroupSourceContentRequest
	GetAppToken() *string
	SetContent(v string) *UpdateGroupSourceContentRequest
	GetContent() *string
	SetForceSync(v bool) *UpdateGroupSourceContentRequest
	GetForceSync() *bool
	SetGroupId(v string) *UpdateGroupSourceContentRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *UpdateGroupSourceContentRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateGroupSourceContentRequest
	GetOpsUserId() *string
	SetSourceId(v string) *UpdateGroupSourceContentRequest
	GetSourceId() *string
	SetTenantId(v string) *UpdateGroupSourceContentRequest
	GetTenantId() *string
}

type UpdateGroupSourceContentRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Updated content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Specifies whether to force synchronization.
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
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
	// The source project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateGroupSourceContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupSourceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupSourceContentRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateGroupSourceContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateGroupSourceContentRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *UpdateGroupSourceContentRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupSourceContentRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateGroupSourceContentRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateGroupSourceContentRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateGroupSourceContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGroupSourceContentRequest) SetAppToken(v string) *UpdateGroupSourceContentRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetContent(v string) *UpdateGroupSourceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetForceSync(v bool) *UpdateGroupSourceContentRequest {
	s.ForceSync = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetGroupId(v string) *UpdateGroupSourceContentRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetOpsUserCode(v string) *UpdateGroupSourceContentRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetOpsUserId(v string) *UpdateGroupSourceContentRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetSourceId(v string) *UpdateGroupSourceContentRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetTenantId(v string) *UpdateGroupSourceContentRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) Validate() error {
	return dara.Validate(s)
}
