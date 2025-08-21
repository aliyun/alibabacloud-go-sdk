// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoginProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginProfile(v *CreateLoginProfileResponseBodyLoginProfile) *CreateLoginProfileResponseBody
	GetLoginProfile() *CreateLoginProfileResponseBodyLoginProfile
	SetRequestId(v string) *CreateLoginProfileResponseBody
	GetRequestId() *string
}

type CreateLoginProfileResponseBody struct {
	// The logon information.
	LoginProfile *CreateLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 29CB303C-1F05-43A6-A6BC-EBC5A797F8DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoginProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseBody) GetLoginProfile() *CreateLoginProfileResponseBodyLoginProfile {
	return s.LoginProfile
}

func (s *CreateLoginProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoginProfileResponseBody) SetLoginProfile(v *CreateLoginProfileResponseBodyLoginProfile) *CreateLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *CreateLoginProfileResponseBody) SetRequestId(v string) *CreateLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoginProfileResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateLoginProfileResponseBodyLoginProfile struct {
	// Indicates whether to forcefully enable MFA for the RAM user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user is required to reset the password upon the next logon.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Indicates whether to enable password-based logons to the console.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-14T03:47:51Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateLoginProfileResponseBodyLoginProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetMFABindRequired() *bool {
	return s.MFABindRequired
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetPasswordResetRequired() *bool {
	return s.PasswordResetRequired
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetStatus() *string {
	return s.Status
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetStatus(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.Status = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetUpdateDate(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetUserPrincipalName(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) Validate() error {
	return dara.Validate(s)
}
