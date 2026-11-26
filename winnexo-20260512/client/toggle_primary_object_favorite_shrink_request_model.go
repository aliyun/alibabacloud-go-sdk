// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePrimaryObjectFavoriteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetAction() *string
	SetAppToken(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetAppToken() *string
	SetObjectIdsShrink(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetObjectIdsShrink() *string
	SetObjectType(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *TogglePrimaryObjectFavoriteShrinkRequest
	GetTenantId() *string
}

type TogglePrimaryObjectFavoriteShrinkRequest struct {
	// The operation type. Valid values:
	//
	// - **1**: Add to whitelist.
	//
	// - **2**: Remove from whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The list of primary object business IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ObjectIdsShrink *string `json:"objectIds,omitempty" xml:"objectIds,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TogglePrimaryObjectFavoriteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TogglePrimaryObjectFavoriteShrinkRequest) GoString() string {
	return s.String()
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetAction() *string {
	return s.Action
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetAction(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.Action = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetAppToken(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetObjectIdsShrink(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetObjectType(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetOperatingObjectName(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetOpsUserCode(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetOpsUserId(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) SetTenantId(v string) *TogglePrimaryObjectFavoriteShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
