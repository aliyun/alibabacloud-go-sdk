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
  // The mode in which you enable a resource directory. Valid values:
  // 
  // 	- CurrentAccount: indicates that the current account is used to enable a resource directory.
  // 
  // 	- NewManagementAccount: indicates that a newly created account is used to enable a resource directory. If you select this mode, you must configure the `MAName`, `MASecureMobilePhone`, and `VerificationCode` parameters.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // CurrentAccount
  EnableMode *string `json:"EnableMode,omitempty" xml:"EnableMode,omitempty"`
  // The name of the newly created account.
  // 
  // Specify the name in the `<Prefix>@rdadmin.aliyunid.com` format. The prefix can contain letters, digits, and special characters but cannot contain consecutive special characters. The prefix must start with a letter or digit and end with a letter or digit. Valid special characters include underscores (_), periods (.), and hyphens (-). The prefix must be 2 to 50 characters in length.
  // 
  // example:
  // 
  // user01@rdadmin.aliyunid.com
  MAName *string `json:"MAName,omitempty" xml:"MAName,omitempty"`
  // The mobile phone number that is bound to the newly created account.
  // 
  // If you leave this parameter empty, the mobile phone number that is bound to the current account is used. The mobile phone number you specify must be the same as the mobile phone number that you specify when you call the [SendVerificationCodeForEnableRD](https://help.aliyun.com/document_detail/364248.html) operation to obtain a verification code.
  // 
  // Specify the mobile phone number in the `<Country code>-<Mobile phone number>` format.
  // 
  // >  Mobile phone numbers in the `86-<Mobile phone number>` format in the Chinese mainland are not supported.
  // 
  // example:
  // 
  // xx-13900001234
  MASecureMobilePhone *string `json:"MASecureMobilePhone,omitempty" xml:"MASecureMobilePhone,omitempty"`
  // The verification code.
  // 
  // You can call the [SendVerificationCodeForEnableRD](https://help.aliyun.com/document_detail/364248.html) operation to obtain the verification code.
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

