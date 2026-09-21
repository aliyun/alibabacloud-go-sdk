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
	// The application ID. The application must be deployed in the image used by the delivery group. You can obtain the ID from the Apps list returned by the [GetAppInstanceGroup](https://help.aliyun.com/document_detail/600836.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// The application specified by AppId must be deployed in the image used by this delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The list of usernames to add authorization for the application. A maximum of 100 usernames can be specified in a single request.
	//
	// At least one of AuthorizeUserIds and UnAuthorizeUserIds must be specified. You can also specify both. Adding authorization is subject to the authorized user quota for the application.
	AuthorizeUserIds []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	// The product type. Application-level authorization applies to WUYING Cloud Application delivery groups.
	//
	// Valid values:
	//
	// - CloudApp: WUYING Cloud Application.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of usernames to be unauthorized for the application. A maximum of 100 usernames can be specified in a single request.
	//
	// At least one of AuthorizeUserIds and UnAuthorizeUserIds must be specified. You can also specify both. Removing authorizations is not subject to quota limits.
	UnAuthorizeUserIds []*string `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
	// The account information of the authorized user, which specifies the account type corresponding to the username.
	//
	// - If the workspace to which the delivery group belongs is an AD workspace, **this parameter is required**: set Type to ad and set AdDomain to the AD domain bound to the workspace.
	//
	// - If this parameter is not specified, the WUYING convenience account (simple) is used by default.
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
