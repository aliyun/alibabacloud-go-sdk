// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeUsersForAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AuthorizeUsersForAppShrinkRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *AuthorizeUsersForAppShrinkRequest
	GetAppInstanceGroupId() *string
	SetAuthorizeUserIds(v []*string) *AuthorizeUsersForAppShrinkRequest
	GetAuthorizeUserIds() []*string
	SetProductType(v string) *AuthorizeUsersForAppShrinkRequest
	GetProductType() *string
	SetUnAuthorizeUserIds(v []*string) *AuthorizeUsersForAppShrinkRequest
	GetUnAuthorizeUserIds() []*string
	SetUserMetaShrink(v string) *AuthorizeUsersForAppShrinkRequest
	GetUserMetaShrink() *string
}

type AuthorizeUsersForAppShrinkRequest struct {
	// The ID of the application. The application must be deployed in the image used by the app instance group. You can obtain the application ID from the Apps list returned by the [GetAppInstanceGroup](https://help.aliyun.com/document_detail/600836.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the app instance group. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to query the ID.
	//
	// The application specified by AppId must be deployed in the image used by the app instance group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The list of usernames of the users to be authorized for the application. You can specify up to 100 usernames in a single call.
	//
	// You must specify at least one of AuthorizeUserIds and UnAuthorizeUserIds, or specify both. Adding authorization is subject to the authorized user quota limit of the application.
	AuthorizeUserIds []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	// The product type. Application-level authorization applies to WUYING Cloud App app instance groups.
	//
	// Valid values:
	//
	// - CloudApp: WUYING Cloud App.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of usernames of the users whose authorization for the application is to be removed. You can specify up to 100 usernames in a single call.
	//
	// You must specify at least one of AuthorizeUserIds and UnAuthorizeUserIds, or specify both. Removing authorization is not subject to quota limits.
	UnAuthorizeUserIds []*string `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
	// The account information of the authorized user, which is used to specify the account type corresponding to the username.
	//
	// - **Required*	- if the workspace to which the app instance group belongs is an AD workspace: Set Type to ad, and set AdDomain to the AD domain bound to the workspace.
	//
	// - If this parameter is not specified, the default account type is WUYING convenience account (simple).
	UserMetaShrink *string `json:"UserMeta,omitempty" xml:"UserMeta,omitempty"`
}

func (s AuthorizeUsersForAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeUsersForAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeUsersForAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *AuthorizeUsersForAppShrinkRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *AuthorizeUsersForAppShrinkRequest) GetAuthorizeUserIds() []*string {
	return s.AuthorizeUserIds
}

func (s *AuthorizeUsersForAppShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *AuthorizeUsersForAppShrinkRequest) GetUnAuthorizeUserIds() []*string {
	return s.UnAuthorizeUserIds
}

func (s *AuthorizeUsersForAppShrinkRequest) GetUserMetaShrink() *string {
	return s.UserMetaShrink
}

func (s *AuthorizeUsersForAppShrinkRequest) SetAppId(v string) *AuthorizeUsersForAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *AuthorizeUsersForAppShrinkRequest) SetAppInstanceGroupId(v string) *AuthorizeUsersForAppShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeUsersForAppShrinkRequest) SetAuthorizeUserIds(v []*string) *AuthorizeUsersForAppShrinkRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeUsersForAppShrinkRequest) SetProductType(v string) *AuthorizeUsersForAppShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeUsersForAppShrinkRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeUsersForAppShrinkRequest {
	s.UnAuthorizeUserIds = v
	return s
}

func (s *AuthorizeUsersForAppShrinkRequest) SetUserMetaShrink(v string) *AuthorizeUsersForAppShrinkRequest {
	s.UserMetaShrink = &v
	return s
}

func (s *AuthorizeUsersForAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
