// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetAppInstancePersistentId(v string) *AuthorizeInstanceGroupRequest
	GetAppInstancePersistentId() *string
	SetAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupRequest
	GetAuthorizeUserGroupIds() []*string
	SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest
	GetAuthorizeUserIds() []*string
	SetAvatarId(v string) *AuthorizeInstanceGroupRequest
	GetAvatarId() *string
	SetProductType(v string) *AuthorizeInstanceGroupRequest
	GetProductType() *string
	SetUnAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupRequest
	GetUnAuthorizeUserGroupIds() []*string
	SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest
	GetUnAuthorizeUserIds() []*string
	SetUserMeta(v *AuthorizeInstanceGroupRequestUserMeta) *AuthorizeInstanceGroupRequest
	GetUserMeta() *AuthorizeInstanceGroupRequestUserMeta
}

type AuthorizeInstanceGroupRequest struct {
	// 交付群組 ID。可呼叫 [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) 介面取得。
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// 持續性工作階段 ID。
	//
	// example:
	//
	// p-0cc7s3mw2fg4j****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// 授權使用者群組 ID 清單。
	//
	// if can be null:
	// true
	AuthorizeUserGroupIds []*string `json:"AuthorizeUserGroupIds,omitempty" xml:"AuthorizeUserGroupIds,omitempty" type:"Repeated"`
	// 要新增交付群組授權的使用者名稱清單。可設定 1\\~100 個。
	AuthorizeUserIds []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	// 使用者分身 ID。
	//
	// > 此參數未開放使用。
	//
	// example:
	//
	// default
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// 產品類型。
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// 取消授權使用者群組 ID 清單。
	//
	// if can be null:
	// true
	UnAuthorizeUserGroupIds []*string `json:"UnAuthorizeUserGroupIds,omitempty" xml:"UnAuthorizeUserGroupIds,omitempty" type:"Repeated"`
	// 要移除交付群組授權的使用者名稱清單。可設定 1\\~100 個。
	UnAuthorizeUserIds []*string `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
	// 使用者資訊。
	UserMeta *AuthorizeInstanceGroupRequestUserMeta `json:"UserMeta,omitempty" xml:"UserMeta,omitempty" type:"Struct"`
}

func (s AuthorizeInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *AuthorizeInstanceGroupRequest) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *AuthorizeInstanceGroupRequest) GetAuthorizeUserGroupIds() []*string {
	return s.AuthorizeUserGroupIds
}

func (s *AuthorizeInstanceGroupRequest) GetAuthorizeUserIds() []*string {
	return s.AuthorizeUserIds
}

func (s *AuthorizeInstanceGroupRequest) GetAvatarId() *string {
	return s.AvatarId
}

func (s *AuthorizeInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *AuthorizeInstanceGroupRequest) GetUnAuthorizeUserGroupIds() []*string {
	return s.UnAuthorizeUserGroupIds
}

func (s *AuthorizeInstanceGroupRequest) GetUnAuthorizeUserIds() []*string {
	return s.UnAuthorizeUserIds
}

func (s *AuthorizeInstanceGroupRequest) GetUserMeta() *AuthorizeInstanceGroupRequestUserMeta {
	return s.UserMeta
}

func (s *AuthorizeInstanceGroupRequest) SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAppInstancePersistentId(v string) *AuthorizeInstanceGroupRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.AuthorizeUserGroupIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAvatarId(v string) *AuthorizeInstanceGroupRequest {
	s.AvatarId = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetProductType(v string) *AuthorizeInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetUnAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.UnAuthorizeUserGroupIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.UnAuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetUserMeta(v *AuthorizeInstanceGroupRequestUserMeta) *AuthorizeInstanceGroupRequest {
	s.UserMeta = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) Validate() error {
	if s.UserMeta != nil {
		if err := s.UserMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthorizeInstanceGroupRequestUserMeta struct {
	// AD 網域名稱。
	//
	// example:
	//
	// example.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// 使用者類型。
	//
	// example:
	//
	// simple
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AuthorizeInstanceGroupRequestUserMeta) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeInstanceGroupRequestUserMeta) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupRequestUserMeta) GetAdDomain() *string {
	return s.AdDomain
}

func (s *AuthorizeInstanceGroupRequestUserMeta) GetType() *string {
	return s.Type
}

func (s *AuthorizeInstanceGroupRequestUserMeta) SetAdDomain(v string) *AuthorizeInstanceGroupRequestUserMeta {
	s.AdDomain = &v
	return s
}

func (s *AuthorizeInstanceGroupRequestUserMeta) SetType(v string) *AuthorizeInstanceGroupRequestUserMeta {
	s.Type = &v
	return s
}

func (s *AuthorizeInstanceGroupRequestUserMeta) Validate() error {
	return dara.Validate(s)
}
