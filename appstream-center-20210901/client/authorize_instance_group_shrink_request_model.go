// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeInstanceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupShrinkRequest
	GetAppInstanceGroupId() *string
	SetAppInstancePersistentId(v string) *AuthorizeInstanceGroupShrinkRequest
	GetAppInstancePersistentId() *string
	SetAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupShrinkRequest
	GetAuthorizeUserGroupIds() []*string
	SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupShrinkRequest
	GetAuthorizeUserIds() []*string
	SetAvatarId(v string) *AuthorizeInstanceGroupShrinkRequest
	GetAvatarId() *string
	SetProductType(v string) *AuthorizeInstanceGroupShrinkRequest
	GetProductType() *string
	SetUnAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupShrinkRequest
	GetUnAuthorizeUserGroupIds() []*string
	SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupShrinkRequest
	GetUnAuthorizeUserIds() []*string
	SetUserMetaShrink(v string) *AuthorizeInstanceGroupShrinkRequest
	GetUserMetaShrink() *string
}

type AuthorizeInstanceGroupShrinkRequest struct {
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
	UserMetaShrink *string `json:"UserMeta,omitempty" xml:"UserMeta,omitempty"`
}

func (s AuthorizeInstanceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetAppInstancePersistentId() *string {
	return s.AppInstancePersistentId
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetAuthorizeUserGroupIds() []*string {
	return s.AuthorizeUserGroupIds
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetAuthorizeUserIds() []*string {
	return s.AuthorizeUserIds
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetAvatarId() *string {
	return s.AvatarId
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetUnAuthorizeUserGroupIds() []*string {
	return s.UnAuthorizeUserGroupIds
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetUnAuthorizeUserIds() []*string {
	return s.UnAuthorizeUserIds
}

func (s *AuthorizeInstanceGroupShrinkRequest) GetUserMetaShrink() *string {
	return s.UserMetaShrink
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAppInstancePersistentId(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupShrinkRequest {
	s.AuthorizeUserGroupIds = v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupShrinkRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAvatarId(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.AvatarId = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetProductType(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetUnAuthorizeUserGroupIds(v []*string) *AuthorizeInstanceGroupShrinkRequest {
	s.UnAuthorizeUserGroupIds = v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupShrinkRequest {
	s.UnAuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetUserMetaShrink(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.UserMetaShrink = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
