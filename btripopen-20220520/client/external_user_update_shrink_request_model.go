// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserUpdateShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBirthday(v string) *ExternalUserUpdateShrinkRequest
  GetBirthday() *string 
  SetCertRequestListShrink(v string) *ExternalUserUpdateShrinkRequest
  GetCertRequestListShrink() *string 
  SetEmail(v string) *ExternalUserUpdateShrinkRequest
  GetEmail() *string 
  SetPhone(v string) *ExternalUserUpdateShrinkRequest
  GetPhone() *string 
  SetRealName(v string) *ExternalUserUpdateShrinkRequest
  GetRealName() *string 
  SetRealNameEn(v string) *ExternalUserUpdateShrinkRequest
  GetRealNameEn() *string 
}

type ExternalUserUpdateShrinkRequest struct {
  // example:
  // 
  // 2000-01-02
  Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
  CertRequestListShrink *string `json:"cert_request_list,omitempty" xml:"cert_request_list,omitempty"`
  // example:
  // 
  // zhangsan@alibaba-inc.com
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // example:
  // 
  // 13438009765
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  RealName *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
  // example:
  // 
  // zhang/san
  RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
}

func (s ExternalUserUpdateShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserUpdateShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExternalUserUpdateShrinkRequest) GetBirthday() *string  {
  return s.Birthday
}

func (s *ExternalUserUpdateShrinkRequest) GetCertRequestListShrink() *string  {
  return s.CertRequestListShrink
}

func (s *ExternalUserUpdateShrinkRequest) GetEmail() *string  {
  return s.Email
}

func (s *ExternalUserUpdateShrinkRequest) GetPhone() *string  {
  return s.Phone
}

func (s *ExternalUserUpdateShrinkRequest) GetRealName() *string  {
  return s.RealName
}

func (s *ExternalUserUpdateShrinkRequest) GetRealNameEn() *string  {
  return s.RealNameEn
}

func (s *ExternalUserUpdateShrinkRequest) SetBirthday(v string) *ExternalUserUpdateShrinkRequest {
  s.Birthday = &v
  return s
}

func (s *ExternalUserUpdateShrinkRequest) SetCertRequestListShrink(v string) *ExternalUserUpdateShrinkRequest {
  s.CertRequestListShrink = &v
  return s
}

func (s *ExternalUserUpdateShrinkRequest) SetEmail(v string) *ExternalUserUpdateShrinkRequest {
  s.Email = &v
  return s
}

func (s *ExternalUserUpdateShrinkRequest) SetPhone(v string) *ExternalUserUpdateShrinkRequest {
  s.Phone = &v
  return s
}

func (s *ExternalUserUpdateShrinkRequest) SetRealName(v string) *ExternalUserUpdateShrinkRequest {
  s.RealName = &v
  return s
}

func (s *ExternalUserUpdateShrinkRequest) SetRealNameEn(v string) *ExternalUserUpdateShrinkRequest {
  s.RealNameEn = &v
  return s
}

func (s *ExternalUserUpdateShrinkRequest) Validate() error {
  return dara.Validate(s)
}

