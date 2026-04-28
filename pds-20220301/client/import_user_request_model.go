// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationDisplayName(v string) *ImportUserRequest
	GetAuthenticationDisplayName() *string
	SetAuthenticationType(v string) *ImportUserRequest
	GetAuthenticationType() *string
	SetAutoCreateDrive(v bool) *ImportUserRequest
	GetAutoCreateDrive() *bool
	SetDriveTotalSize(v int64) *ImportUserRequest
	GetDriveTotalSize() *int64
	SetExtra(v string) *ImportUserRequest
	GetExtra() *string
	SetIdentity(v string) *ImportUserRequest
	GetIdentity() *string
	SetNickName(v string) *ImportUserRequest
	GetNickName() *string
	SetParentGroupId(v string) *ImportUserRequest
	GetParentGroupId() *string
}

type ImportUserRequest struct {
	// The display name of the authentication type.
	//
	// example:
	//
	// 10000
	AuthenticationDisplayName *string `json:"authentication_display_name,omitempty" xml:"authentication_display_name,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- mobile: mobile number.
	//
	// 	- email: email address.
	//
	// 	- ding: DingTalk account.
	//
	// 	- ram: Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: WeCom account.
	//
	// 	- ldap: Lightweight Directory Access Protocol (LDAP) account.
	//
	// 	- custom: custom account.
	//
	// This parameter is required.
	//
	// example:
	//
	// mobile
	AuthenticationType *string `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	// Specifies whether to automatically create a drive.
	//
	// example:
	//
	// false
	AutoCreateDrive *bool `json:"auto_create_drive,omitempty" xml:"auto_create_drive,omitempty"`
	// The size of the drive. The value cannot be smaller than -1. A value of -1 specifies that the size is unlimited.
	//
	// example:
	//
	// 10240
	DriveTotalSize *int64 `json:"drive_total_size,omitempty" xml:"drive_total_size,omitempty"`
	// The additional information.
	//
	// If authentication_type is set to mobile, set this parameter to a country code. If you do not specify this parameter, 86 is used by default.
	//
	// example:
	//
	// 1
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The unique identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130****
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// The nickname of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The ID of the group to which the user is added.
	//
	// example:
	//
	// g12
	ParentGroupId *string `json:"parent_group_id,omitempty" xml:"parent_group_id,omitempty"`
}

func (s ImportUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportUserRequest) GoString() string {
	return s.String()
}

func (s *ImportUserRequest) GetAuthenticationDisplayName() *string {
	return s.AuthenticationDisplayName
}

func (s *ImportUserRequest) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *ImportUserRequest) GetAutoCreateDrive() *bool {
	return s.AutoCreateDrive
}

func (s *ImportUserRequest) GetDriveTotalSize() *int64 {
	return s.DriveTotalSize
}

func (s *ImportUserRequest) GetExtra() *string {
	return s.Extra
}

func (s *ImportUserRequest) GetIdentity() *string {
	return s.Identity
}

func (s *ImportUserRequest) GetNickName() *string {
	return s.NickName
}

func (s *ImportUserRequest) GetParentGroupId() *string {
	return s.ParentGroupId
}

func (s *ImportUserRequest) SetAuthenticationDisplayName(v string) *ImportUserRequest {
	s.AuthenticationDisplayName = &v
	return s
}

func (s *ImportUserRequest) SetAuthenticationType(v string) *ImportUserRequest {
	s.AuthenticationType = &v
	return s
}

func (s *ImportUserRequest) SetAutoCreateDrive(v bool) *ImportUserRequest {
	s.AutoCreateDrive = &v
	return s
}

func (s *ImportUserRequest) SetDriveTotalSize(v int64) *ImportUserRequest {
	s.DriveTotalSize = &v
	return s
}

func (s *ImportUserRequest) SetExtra(v string) *ImportUserRequest {
	s.Extra = &v
	return s
}

func (s *ImportUserRequest) SetIdentity(v string) *ImportUserRequest {
	s.Identity = &v
	return s
}

func (s *ImportUserRequest) SetNickName(v string) *ImportUserRequest {
	s.NickName = &v
	return s
}

func (s *ImportUserRequest) SetParentGroupId(v string) *ImportUserRequest {
	s.ParentGroupId = &v
	return s
}

func (s *ImportUserRequest) Validate() error {
	return dara.Validate(s)
}
