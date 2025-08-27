// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserAddRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBirthday(v string) *ExternalUserAddRequest
  GetBirthday() *string 
  SetCertRequestList(v []*ExternalUserAddRequestCertRequestList) *ExternalUserAddRequest
  GetCertRequestList() []*ExternalUserAddRequestCertRequestList 
  SetEmail(v string) *ExternalUserAddRequest
  GetEmail() *string 
  SetExternalUserId(v string) *ExternalUserAddRequest
  GetExternalUserId() *string 
  SetPhone(v string) *ExternalUserAddRequest
  GetPhone() *string 
  SetRealName(v string) *ExternalUserAddRequest
  GetRealName() *string 
  SetRealNameEn(v string) *ExternalUserAddRequest
  GetRealNameEn() *string 
  SetUserType(v int32) *ExternalUserAddRequest
  GetUserType() *int32 
}

type ExternalUserAddRequest struct {
  // example:
  // 
  // 2000-01-02
  Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
  CertRequestList []*ExternalUserAddRequestCertRequestList `json:"cert_request_list,omitempty" xml:"cert_request_list,omitempty" type:"Repeated"`
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

func (s ExternalUserAddRequest) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserAddRequest) GoString() string {
  return s.String()
}

func (s *ExternalUserAddRequest) GetBirthday() *string  {
  return s.Birthday
}

func (s *ExternalUserAddRequest) GetCertRequestList() []*ExternalUserAddRequestCertRequestList  {
  return s.CertRequestList
}

func (s *ExternalUserAddRequest) GetEmail() *string  {
  return s.Email
}

func (s *ExternalUserAddRequest) GetExternalUserId() *string  {
  return s.ExternalUserId
}

func (s *ExternalUserAddRequest) GetPhone() *string  {
  return s.Phone
}

func (s *ExternalUserAddRequest) GetRealName() *string  {
  return s.RealName
}

func (s *ExternalUserAddRequest) GetRealNameEn() *string  {
  return s.RealNameEn
}

func (s *ExternalUserAddRequest) GetUserType() *int32  {
  return s.UserType
}

func (s *ExternalUserAddRequest) SetBirthday(v string) *ExternalUserAddRequest {
  s.Birthday = &v
  return s
}

func (s *ExternalUserAddRequest) SetCertRequestList(v []*ExternalUserAddRequestCertRequestList) *ExternalUserAddRequest {
  s.CertRequestList = v
  return s
}

func (s *ExternalUserAddRequest) SetEmail(v string) *ExternalUserAddRequest {
  s.Email = &v
  return s
}

func (s *ExternalUserAddRequest) SetExternalUserId(v string) *ExternalUserAddRequest {
  s.ExternalUserId = &v
  return s
}

func (s *ExternalUserAddRequest) SetPhone(v string) *ExternalUserAddRequest {
  s.Phone = &v
  return s
}

func (s *ExternalUserAddRequest) SetRealName(v string) *ExternalUserAddRequest {
  s.RealName = &v
  return s
}

func (s *ExternalUserAddRequest) SetRealNameEn(v string) *ExternalUserAddRequest {
  s.RealNameEn = &v
  return s
}

func (s *ExternalUserAddRequest) SetUserType(v int32) *ExternalUserAddRequest {
  s.UserType = &v
  return s
}

func (s *ExternalUserAddRequest) Validate() error {
  return dara.Validate(s)
}

type ExternalUserAddRequestCertRequestList struct {
  // example:
  // 
  // 2034-10-01
  CertExpiredTime *string `json:"cert_expired_time,omitempty" xml:"cert_expired_time,omitempty"`
  // example:
  // 
  // CN
  CertNation *string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 330101199010010213
  CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 0
  CertType *int32 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
  // example:
  // 
  // CN
  Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
}

func (s ExternalUserAddRequestCertRequestList) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserAddRequestCertRequestList) GoString() string {
  return s.String()
}

func (s *ExternalUserAddRequestCertRequestList) GetCertExpiredTime() *string  {
  return s.CertExpiredTime
}

func (s *ExternalUserAddRequestCertRequestList) GetCertNation() *string  {
  return s.CertNation
}

func (s *ExternalUserAddRequestCertRequestList) GetCertNo() *string  {
  return s.CertNo
}

func (s *ExternalUserAddRequestCertRequestList) GetCertType() *int32  {
  return s.CertType
}

func (s *ExternalUserAddRequestCertRequestList) GetNationality() *string  {
  return s.Nationality
}

func (s *ExternalUserAddRequestCertRequestList) SetCertExpiredTime(v string) *ExternalUserAddRequestCertRequestList {
  s.CertExpiredTime = &v
  return s
}

func (s *ExternalUserAddRequestCertRequestList) SetCertNation(v string) *ExternalUserAddRequestCertRequestList {
  s.CertNation = &v
  return s
}

func (s *ExternalUserAddRequestCertRequestList) SetCertNo(v string) *ExternalUserAddRequestCertRequestList {
  s.CertNo = &v
  return s
}

func (s *ExternalUserAddRequestCertRequestList) SetCertType(v int32) *ExternalUserAddRequestCertRequestList {
  s.CertType = &v
  return s
}

func (s *ExternalUserAddRequestCertRequestList) SetNationality(v string) *ExternalUserAddRequestCertRequestList {
  s.Nationality = &v
  return s
}

func (s *ExternalUserAddRequestCertRequestList) Validate() error {
  return dara.Validate(s)
}

