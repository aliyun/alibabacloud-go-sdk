// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginProfile(v *GetLoginProfileResponseBodyLoginProfile) *GetLoginProfileResponseBody
	GetLoginProfile() *GetLoginProfileResponseBodyLoginProfile
	SetRequestId(v string) *GetLoginProfileResponseBody
	GetRequestId() *string
}

type GetLoginProfileResponseBody struct {
	// The logon configurations of the RAM user.
	LoginProfile *GetLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBody) GetLoginProfile() *GetLoginProfileResponseBodyLoginProfile {
	return s.LoginProfile
}

func (s *GetLoginProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoginProfileResponseBody) SetLoginProfile(v *GetLoginProfileResponseBodyLoginProfile) *GetLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *GetLoginProfileResponseBody) SetRequestId(v string) *GetLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginProfileResponseBody) Validate() error {
	if s.LoginProfile != nil {
		if err := s.LoginProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoginProfileResponseBodyLoginProfile struct {
	// The creation time.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// Indicates whether a multi-factor authentication (MFA) device must be bound to the RAM user.
	//
	// example:
	//
	// true
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user must change the password upon logon.
	//
	// example:
	//
	// true
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetLoginProfileResponseBodyLoginProfile) String() string {
	return dara.Prettify(s)
}

func (s GetLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetMFABindRequired() *bool {
	return s.MFABindRequired
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetPasswordResetRequired() *bool {
	return s.PasswordResetRequired
}

func (s *GetLoginProfileResponseBodyLoginProfile) GetUserName() *string {
	return s.UserName
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetCreateDate(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.CreateDate = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetUserName(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.UserName = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) Validate() error {
	return dara.Validate(s)
}
