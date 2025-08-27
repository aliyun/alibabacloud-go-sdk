// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserAddShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBirthday(v string) *ExternalUserAddShrinkRequest
  GetBirthday() *string 
  SetCertRequestListShrink(v string) *ExternalUserAddShrinkRequest
  GetCertRequestListShrink() *string 
  SetEmail(v string) *ExternalUserAddShrinkRequest
  GetEmail() *string 
  SetExternalUserId(v string) *ExternalUserAddShrinkRequest
  GetExternalUserId() *string 
  SetPhone(v string) *ExternalUserAddShrinkRequest
  GetPhone() *string 
  SetRealName(v string) *ExternalUserAddShrinkRequest
  GetRealName() *string 
  SetRealNameEn(v string) *ExternalUserAddShrinkRequest
  GetRealNameEn() *string 
  SetUserType(v int32) *ExternalUserAddShrinkRequest
  GetUserType() *int32 
}

type ExternalUserAddShrinkRequest struct {
  // example:
  // 
  // 2000-01-02
  Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
  CertRequestListShrink *string `json:"cert_request_list,omitempty" xml:"cert_request_list,omitempty"`
  // example:
  // 
  // zhangsan@alibaba-inc.com
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 0012
  ExternalUserId *string `json:"external_user_id,omitempty" xml:"external_user_id,omitempty"`
  // example:
  // 
  // 13438009765
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  RealName *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
  // example:
  // 
  // zhang/san
  RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 2
  UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s ExternalUserAddShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserAddShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExternalUserAddShrinkRequest) GetBirthday() *string  {
  return s.Birthday
}

func (s *ExternalUserAddShrinkRequest) GetCertRequestListShrink() *string  {
  return s.CertRequestListShrink
}

func (s *ExternalUserAddShrinkRequest) GetEmail() *string  {
  return s.Email
}

func (s *ExternalUserAddShrinkRequest) GetExternalUserId() *string  {
  return s.ExternalUserId
}

func (s *ExternalUserAddShrinkRequest) GetPhone() *string  {
  return s.Phone
}

func (s *ExternalUserAddShrinkRequest) GetRealName() *string  {
  return s.RealName
}

func (s *ExternalUserAddShrinkRequest) GetRealNameEn() *string  {
  return s.RealNameEn
}

func (s *ExternalUserAddShrinkRequest) GetUserType() *int32  {
  return s.UserType
}

func (s *ExternalUserAddShrinkRequest) SetBirthday(v string) *ExternalUserAddShrinkRequest {
  s.Birthday = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) SetCertRequestListShrink(v string) *ExternalUserAddShrinkRequest {
  s.CertRequestListShrink = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) SetEmail(v string) *ExternalUserAddShrinkRequest {
  s.Email = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) SetExternalUserId(v string) *ExternalUserAddShrinkRequest {
  s.ExternalUserId = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) SetPhone(v string) *ExternalUserAddShrinkRequest {
  s.Phone = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) SetRealName(v string) *ExternalUserAddShrinkRequest {
  s.RealName = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) SetRealNameEn(v string) *ExternalUserAddShrinkRequest {
  s.RealNameEn = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) SetUserType(v int32) *ExternalUserAddShrinkRequest {
  s.UserType = &v
  return s
}

func (s *ExternalUserAddShrinkRequest) Validate() error {
  return dara.Validate(s)
}

