// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceDirectoryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEnableMode(v string) *EnableResourceDirectoryRequest
  GetEnableMode() *string 
  SetMAName(v string) *EnableResourceDirectoryRequest
  GetMAName() *string 
  SetMASecureMobilePhone(v string) *EnableResourceDirectoryRequest
  GetMASecureMobilePhone() *string 
  SetVerificationCode(v string) *EnableResourceDirectoryRequest
  GetVerificationCode() *string 
}

type EnableResourceDirectoryRequest struct {
  // The method used to enable the resource directory. Valid values:
  // 
  // - CurrentAccount: enables the resource directory by using the current logon account.
  // 
  // - NewManagementAccount: enables the resource directory by creating a new management account. If you select this method, you must configure `MAName`, `MASecureMobilePhone`, and `VerificationCode`.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // CurrentAccount
  EnableMode *string `json:"EnableMode,omitempty" xml:"EnableMode,omitempty"`
  // The name of the new management account.
  // 
  // Format: `<Prefix>@rdadmin.aliyunid.com`. The prefix can contain letters, digits, and the special characters `_.-`. The prefix must start and end with a letter or digit, and cannot contain consecutive special characters. The prefix must be 2 to 50 characters in length.
  // 
  // example:
  // 
  // user01@rdadmin.aliyunid.com
  MAName *string `json:"MAName,omitempty" xml:"MAName,omitempty"`
  // The secure phone number of the new management account.
  // 
  // If you leave this parameter empty, the secure phone number of the current logon account is used. The phone number must be the same as the one used to authenticate and obtain the verification code in [SendVerificationCodeForEnableRD](https://help.aliyun.com/document_detail/364248.html).
  // 
  // Format: `<Country code>-<Phone number>`.
  // 
  // <props="intl">
  // 
  // > Phone numbers from the Chinese mainland in the format `86-<Phone number>` are not supported.
  // 
  // 
  // <props="china">
  // 
  // > Only phone numbers from the Chinese mainland in the format `86-<Phone number>` are supported.
  // 
  // example:
  // 
  // xx-13900001234
  MASecureMobilePhone *string `json:"MASecureMobilePhone,omitempty" xml:"MASecureMobilePhone,omitempty"`
  // The verification code.
  // 
  // You can call [SendVerificationCodeForEnableRD](https://help.aliyun.com/document_detail/364248.html) to obtain the verification code.
  // 
  // example:
  // 
  // 123456
  VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
}

func (s EnableResourceDirectoryRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceDirectoryRequest) GoString() string {
  return s.String()
}

func (s *EnableResourceDirectoryRequest) GetEnableMode() *string  {
  return s.EnableMode
}

func (s *EnableResourceDirectoryRequest) GetMAName() *string  {
  return s.MAName
}

func (s *EnableResourceDirectoryRequest) GetMASecureMobilePhone() *string  {
  return s.MASecureMobilePhone
}

func (s *EnableResourceDirectoryRequest) GetVerificationCode() *string  {
  return s.VerificationCode
}

func (s *EnableResourceDirectoryRequest) SetEnableMode(v string) *EnableResourceDirectoryRequest {
  s.EnableMode = &v
  return s
}

func (s *EnableResourceDirectoryRequest) SetMAName(v string) *EnableResourceDirectoryRequest {
  s.MAName = &v
  return s
}

func (s *EnableResourceDirectoryRequest) SetMASecureMobilePhone(v string) *EnableResourceDirectoryRequest {
  s.MASecureMobilePhone = &v
  return s
}

func (s *EnableResourceDirectoryRequest) SetVerificationCode(v string) *EnableResourceDirectoryRequest {
  s.VerificationCode = &v
  return s
}

func (s *EnableResourceDirectoryRequest) Validate() error {
  return dara.Validate(s)
}

