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
	// The logon configurations of the RAM user.
	LoginProfile *CreateLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
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
	if s.LoginProfile != nil {
		if err := s.LoginProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLoginProfileResponseBodyLoginProfile struct {
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether an MFA device must be bound to the RAM user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user must change the password upon logon.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateLoginProfileResponseBodyLoginProfile) String() string {
	return dara.Prettify(s)
}

func (s CreateLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetMFABindRequired() *bool {
	return s.MFABindRequired
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetPasswordResetRequired() *bool {
	return s.PasswordResetRequired
}

func (s *CreateLoginProfileResponseBodyLoginProfile) GetUserName() *string {
	return s.UserName
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetCreateDate(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.CreateDate = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetUserName(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.UserName = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) Validate() error {
	return dara.Validate(s)
}
