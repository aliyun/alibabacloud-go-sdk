// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountTodosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CountTodosRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *CountTodosRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *CountTodosRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CountTodosRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CountTodosRequest
	GetTenantId() *string
}

type CountTodosRequest struct {
	// The platform application token issued by the console. If specified, the request is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID used for filtering. If not specified, all workspaces are included.
	//
	// example:
	//
	// cg_123
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
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CountTodosRequest) String() string {
	return dara.Prettify(s)
}

func (s CountTodosRequest) GoString() string {
	return s.String()
}

func (s *CountTodosRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CountTodosRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CountTodosRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CountTodosRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CountTodosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CountTodosRequest) SetAppToken(v string) *CountTodosRequest {
	s.AppToken = &v
	return s
}

func (s *CountTodosRequest) SetCollaborationGroupId(v string) *CountTodosRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CountTodosRequest) SetOpsUserCode(v string) *CountTodosRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CountTodosRequest) SetOpsUserId(v string) *CountTodosRequest {
	s.OpsUserId = &v
	return s
}

func (s *CountTodosRequest) SetTenantId(v string) *CountTodosRequest {
	s.TenantId = &v
	return s
}

func (s *CountTodosRequest) Validate() error {
	return dara.Validate(s)
}
