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
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// 持久会话ID。
	//
	// example:
	//
	// p-0cc7s3mw2fg4j****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// if can be null:
	// true
	AuthorizeUserGroupIds []*string `json:"AuthorizeUserGroupIds,omitempty" xml:"AuthorizeUserGroupIds,omitempty" type:"Repeated"`
	// The IDs of the users that you want to add to the authorization list of the delivery group. You can specify 1 to 100 user IDs.
	AuthorizeUserIds []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// default
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// if can be null:
	// true
	UnAuthorizeUserGroupIds []*string `json:"UnAuthorizeUserGroupIds,omitempty" xml:"UnAuthorizeUserGroupIds,omitempty" type:"Repeated"`
	// The IDs of the users that you want to remove from the authorization list of the delivery group. You can specify 1 to 100 user IDs.
	UnAuthorizeUserIds []*string `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
	// The user information.
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
	return dara.Validate(s)
}

type AuthorizeInstanceGroupRequestUserMeta struct {
	// The AD domain name.
	//
	// example:
	//
	// example.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The user type.
	//
	// Valid values:
	//
	// 	- ad: Active Directory (AD) account
	//
	// 	- simple: convenience account
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
