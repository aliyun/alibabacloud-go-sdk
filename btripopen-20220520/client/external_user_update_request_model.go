// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserUpdateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBirthday(v string) *ExternalUserUpdateRequest
  GetBirthday() *string 
  SetCertRequestList(v []*ExternalUserUpdateRequestCertRequestList) *ExternalUserUpdateRequest
  GetCertRequestList() []*ExternalUserUpdateRequestCertRequestList 
  SetEmail(v string) *ExternalUserUpdateRequest
  GetEmail() *string 
  SetPhone(v string) *ExternalUserUpdateRequest
  GetPhone() *string 
  SetRealName(v string) *ExternalUserUpdateRequest
  GetRealName() *string 
  SetRealNameEn(v string) *ExternalUserUpdateRequest
  GetRealNameEn() *string 
}

type ExternalUserUpdateRequest struct {
  // example:
  // 
  // 2000-01-02
  Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
  CertRequestList []*ExternalUserUpdateRequestCertRequestList `json:"cert_request_list,omitempty" xml:"cert_request_list,omitempty" type:"Repeated"`
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

func (s ExternalUserUpdateRequest) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserUpdateRequest) GoString() string {
  return s.String()
}

func (s *ExternalUserUpdateRequest) GetBirthday() *string  {
  return s.Birthday
}

func (s *ExternalUserUpdateRequest) GetCertRequestList() []*ExternalUserUpdateRequestCertRequestList  {
  return s.CertRequestList
}

func (s *ExternalUserUpdateRequest) GetEmail() *string  {
  return s.Email
}

func (s *ExternalUserUpdateRequest) GetPhone() *string  {
  return s.Phone
}

func (s *ExternalUserUpdateRequest) GetRealName() *string  {
  return s.RealName
}

func (s *ExternalUserUpdateRequest) GetRealNameEn() *string  {
  return s.RealNameEn
}

func (s *ExternalUserUpdateRequest) SetBirthday(v string) *ExternalUserUpdateRequest {
  s.Birthday = &v
  return s
}

func (s *ExternalUserUpdateRequest) SetCertRequestList(v []*ExternalUserUpdateRequestCertRequestList) *ExternalUserUpdateRequest {
  s.CertRequestList = v
  return s
}

func (s *ExternalUserUpdateRequest) SetEmail(v string) *ExternalUserUpdateRequest {
  s.Email = &v
  return s
}

func (s *ExternalUserUpdateRequest) SetPhone(v string) *ExternalUserUpdateRequest {
  s.Phone = &v
  return s
}

func (s *ExternalUserUpdateRequest) SetRealName(v string) *ExternalUserUpdateRequest {
  s.RealName = &v
  return s
}

func (s *ExternalUserUpdateRequest) SetRealNameEn(v string) *ExternalUserUpdateRequest {
  s.RealNameEn = &v
  return s
}

func (s *ExternalUserUpdateRequest) Validate() error {
  return dara.Validate(s)
}

type ExternalUserUpdateRequestCertRequestList struct {
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

func (s ExternalUserUpdateRequestCertRequestList) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserUpdateRequestCertRequestList) GoString() string {
  return s.String()
}

func (s *ExternalUserUpdateRequestCertRequestList) GetCertExpiredTime() *string  {
  return s.CertExpiredTime
}

func (s *ExternalUserUpdateRequestCertRequestList) GetCertNation() *string  {
  return s.CertNation
}

func (s *ExternalUserUpdateRequestCertRequestList) GetCertNo() *string  {
  return s.CertNo
}

func (s *ExternalUserUpdateRequestCertRequestList) GetCertType() *int32  {
  return s.CertType
}

func (s *ExternalUserUpdateRequestCertRequestList) GetNationality() *string  {
  return s.Nationality
}

func (s *ExternalUserUpdateRequestCertRequestList) SetCertExpiredTime(v string) *ExternalUserUpdateRequestCertRequestList {
  s.CertExpiredTime = &v
  return s
}

func (s *ExternalUserUpdateRequestCertRequestList) SetCertNation(v string) *ExternalUserUpdateRequestCertRequestList {
  s.CertNation = &v
  return s
}

func (s *ExternalUserUpdateRequestCertRequestList) SetCertNo(v string) *ExternalUserUpdateRequestCertRequestList {
  s.CertNo = &v
  return s
}

func (s *ExternalUserUpdateRequestCertRequestList) SetCertType(v int32) *ExternalUserUpdateRequestCertRequestList {
  s.CertType = &v
  return s
}

func (s *ExternalUserUpdateRequestCertRequestList) SetNationality(v string) *ExternalUserUpdateRequestCertRequestList {
  s.Nationality = &v
  return s
}

func (s *ExternalUserUpdateRequestCertRequestList) Validate() error {
  return dara.Validate(s)
}

