// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveGroupResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *MoveGroupResourceRequest
	GetAppToken() *string
	SetGroupId(v string) *MoveGroupResourceRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *MoveGroupResourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *MoveGroupResourceRequest
	GetOpsUserId() *string
	SetSourceDirectoryId(v string) *MoveGroupResourceRequest
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveGroupResourceRequest
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveGroupResourceRequest
	GetTargetDirectoryId() *string
	SetTenantId(v string) *MoveGroupResourceRequest
	GetTenantId() *string
}

type MoveGroupResourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration space ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The real ID of the physical directory in the space where the resource currently resides. The root sentinel is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// The physical GROUP resource ID to be moved. Referenced resources are read-only.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The real ID of the target physical directory in the same space. This value must be different from the source directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MoveGroupResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveGroupResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveGroupResourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *MoveGroupResourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *MoveGroupResourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *MoveGroupResourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *MoveGroupResourceRequest) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveGroupResourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveGroupResourceRequest) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveGroupResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MoveGroupResourceRequest) SetAppToken(v string) *MoveGroupResourceRequest {
	s.AppToken = &v
	return s
}

func (s *MoveGroupResourceRequest) SetGroupId(v string) *MoveGroupResourceRequest {
	s.GroupId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetOpsUserCode(v string) *MoveGroupResourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *MoveGroupResourceRequest) SetOpsUserId(v string) *MoveGroupResourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetSourceDirectoryId(v string) *MoveGroupResourceRequest {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetSourceId(v string) *MoveGroupResourceRequest {
	s.SourceId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetTargetDirectoryId(v string) *MoveGroupResourceRequest {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveGroupResourceRequest) SetTenantId(v string) *MoveGroupResourceRequest {
	s.TenantId = &v
	return s
}

func (s *MoveGroupResourceRequest) Validate() error {
	return dara.Validate(s)
}
