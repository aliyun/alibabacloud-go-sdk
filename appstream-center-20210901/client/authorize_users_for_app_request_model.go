// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeUsersForAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AuthorizeUsersForAppRequest
	GetAppId() *string
	SetAppInstanceGroupId(v string) *AuthorizeUsersForAppRequest
	GetAppInstanceGroupId() *string
	SetAuthorizeUserIds(v []*string) *AuthorizeUsersForAppRequest
	GetAuthorizeUserIds() []*string
	SetProductType(v string) *AuthorizeUsersForAppRequest
	GetProductType() *string
	SetUnAuthorizeUserIds(v []*string) *AuthorizeUsersForAppRequest
	GetUnAuthorizeUserIds() []*string
	SetUserMeta(v *AuthorizeUsersForAppRequestUserMeta) *AuthorizeUsersForAppRequest
	GetUserMeta() *AuthorizeUsersForAppRequestUserMeta
}

type AuthorizeUsersForAppRequest struct {
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
	UserMeta *AuthorizeUsersForAppRequestUserMeta `json:"UserMeta,omitempty" xml:"UserMeta,omitempty" type:"Struct"`
}

func (s AuthorizeUsersForAppRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeUsersForAppRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeUsersForAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *AuthorizeUsersForAppRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *AuthorizeUsersForAppRequest) GetAuthorizeUserIds() []*string {
	return s.AuthorizeUserIds
}

func (s *AuthorizeUsersForAppRequest) GetProductType() *string {
	return s.ProductType
}

func (s *AuthorizeUsersForAppRequest) GetUnAuthorizeUserIds() []*string {
	return s.UnAuthorizeUserIds
}

func (s *AuthorizeUsersForAppRequest) GetUserMeta() *AuthorizeUsersForAppRequestUserMeta {
	return s.UserMeta
}

func (s *AuthorizeUsersForAppRequest) SetAppId(v string) *AuthorizeUsersForAppRequest {
	s.AppId = &v
	return s
}

func (s *AuthorizeUsersForAppRequest) SetAppInstanceGroupId(v string) *AuthorizeUsersForAppRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeUsersForAppRequest) SetAuthorizeUserIds(v []*string) *AuthorizeUsersForAppRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeUsersForAppRequest) SetProductType(v string) *AuthorizeUsersForAppRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeUsersForAppRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeUsersForAppRequest {
	s.UnAuthorizeUserIds = v
	return s
}

func (s *AuthorizeUsersForAppRequest) SetUserMeta(v *AuthorizeUsersForAppRequestUserMeta) *AuthorizeUsersForAppRequest {
	s.UserMeta = v
	return s
}

func (s *AuthorizeUsersForAppRequest) Validate() error {
	if s.UserMeta != nil {
		if err := s.UserMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuthorizeUsersForAppRequestUserMeta struct {
	// The AD domain name. Specify this parameter when Type is set to ad. The value must match the AD domain bound to the workspace of the delivery group.
	//
	// example:
	//
	// example.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The account type. Default value: simple.
	//
	// Valid values:
	//
	// - ad: AD account.
	//
	// - simple: WUYING convenience account.
	//
	// example:
	//
	// simple
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AuthorizeUsersForAppRequestUserMeta) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeUsersForAppRequestUserMeta) GoString() string {
	return s.String()
}

func (s *AuthorizeUsersForAppRequestUserMeta) GetAdDomain() *string {
	return s.AdDomain
}

func (s *AuthorizeUsersForAppRequestUserMeta) GetType() *string {
	return s.Type
}

func (s *AuthorizeUsersForAppRequestUserMeta) SetAdDomain(v string) *AuthorizeUsersForAppRequestUserMeta {
	s.AdDomain = &v
	return s
}

func (s *AuthorizeUsersForAppRequestUserMeta) SetType(v string) *AuthorizeUsersForAppRequestUserMeta {
	s.Type = &v
	return s
}

func (s *AuthorizeUsersForAppRequestUserMeta) Validate() error {
	return dara.Validate(s)
}
