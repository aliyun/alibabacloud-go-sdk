// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnterpriseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEnterpriseInfoResponseBody
	GetCode() *string
	SetMessage(v string) *QueryEnterpriseInfoResponseBody
	GetMessage() *string
	SetProfileInfo(v *QueryEnterpriseInfoResponseBodyProfileInfo) *QueryEnterpriseInfoResponseBody
	GetProfileInfo() *QueryEnterpriseInfoResponseBodyProfileInfo
	SetRequestId(v string) *QueryEnterpriseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEnterpriseInfoResponseBody
	GetSuccess() *bool
}

type QueryEnterpriseInfoResponseBody struct {
	Code        *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	ProfileInfo *QueryEnterpriseInfoResponseBodyProfileInfo `json:"ProfileInfo,omitempty" xml:"ProfileInfo,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEnterpriseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEnterpriseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEnterpriseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEnterpriseInfoResponseBody) GetProfileInfo() *QueryEnterpriseInfoResponseBodyProfileInfo {
	return s.ProfileInfo
}

func (s *QueryEnterpriseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEnterpriseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEnterpriseInfoResponseBody) SetCode(v string) *QueryEnterpriseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBody) SetMessage(v string) *QueryEnterpriseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBody) SetProfileInfo(v *QueryEnterpriseInfoResponseBodyProfileInfo) *QueryEnterpriseInfoResponseBody {
	s.ProfileInfo = v
	return s
}

func (s *QueryEnterpriseInfoResponseBody) SetRequestId(v string) *QueryEnterpriseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBody) SetSuccess(v bool) *QueryEnterpriseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBody) Validate() error {
	if s.ProfileInfo != nil {
		if err := s.ProfileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEnterpriseInfoResponseBodyProfileInfo struct {
	Alias                 *string                                             `json:"Alias,omitempty" xml:"Alias,omitempty"`
	AliyunPK              *string                                             `json:"AliyunPK,omitempty" xml:"AliyunPK,omitempty"`
	AuditStatus           *string                                             `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	BusinessLicenseImgSrc *string                                             `json:"BusinessLicenseImgSrc,omitempty" xml:"BusinessLicenseImgSrc,omitempty"`
	BusinessLicenseNumber *string                                             `json:"BusinessLicenseNumber,omitempty" xml:"BusinessLicenseNumber,omitempty"`
	BusinessLicenseType   *string                                             `json:"BusinessLicenseType,omitempty" xml:"BusinessLicenseType,omitempty"`
	CertifiedFrom         *string                                             `json:"CertifiedFrom,omitempty" xml:"CertifiedFrom,omitempty"`
	CertifiedTime         *string                                             `json:"CertifiedTime,omitempty" xml:"CertifiedTime,omitempty"`
	City                  *QueryEnterpriseInfoResponseBodyProfileInfoCity     `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	CreateTime            *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DetailAddress         *string                                             `json:"DetailAddress,omitempty" xml:"DetailAddress,omitempty"`
	EInterpriseSize       *string                                             `json:"EInterpriseSize,omitempty" xml:"EInterpriseSize,omitempty"`
	EnterpriseEntity      *string                                             `json:"EnterpriseEntity,omitempty" xml:"EnterpriseEntity,omitempty"`
	EntityIDNumber        *string                                             `json:"EntityIDNumber,omitempty" xml:"EntityIDNumber,omitempty"`
	Extend                *string                                             `json:"Extend,omitempty" xml:"Extend,omitempty"`
	Fax                   *string                                             `json:"Fax,omitempty" xml:"Fax,omitempty"`
	Name                  *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone                 *string                                             `json:"Phone,omitempty" xml:"Phone,omitempty"`
	PostalCode            *string                                             `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Profile               *string                                             `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Province              *QueryEnterpriseInfoResponseBodyProfileInfoProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	UpdateTime            *string                                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Years                 *string                                             `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s QueryEnterpriseInfoResponseBodyProfileInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryEnterpriseInfoResponseBodyProfileInfo) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetAlias() *string {
	return s.Alias
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetAliyunPK() *string {
	return s.AliyunPK
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetBusinessLicenseImgSrc() *string {
	return s.BusinessLicenseImgSrc
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetBusinessLicenseNumber() *string {
	return s.BusinessLicenseNumber
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetBusinessLicenseType() *string {
	return s.BusinessLicenseType
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetCertifiedFrom() *string {
	return s.CertifiedFrom
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetCertifiedTime() *string {
	return s.CertifiedTime
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetCity() *QueryEnterpriseInfoResponseBodyProfileInfoCity {
	return s.City
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetDetailAddress() *string {
	return s.DetailAddress
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetEInterpriseSize() *string {
	return s.EInterpriseSize
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetEnterpriseEntity() *string {
	return s.EnterpriseEntity
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetEntityIDNumber() *string {
	return s.EntityIDNumber
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetExtend() *string {
	return s.Extend
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetFax() *string {
	return s.Fax
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetName() *string {
	return s.Name
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetPhone() *string {
	return s.Phone
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetPostalCode() *string {
	return s.PostalCode
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetProfile() *string {
	return s.Profile
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetProvince() *QueryEnterpriseInfoResponseBodyProfileInfoProvince {
	return s.Province
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) GetYears() *string {
	return s.Years
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetAlias(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Alias = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetAliyunPK(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.AliyunPK = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetAuditStatus(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.AuditStatus = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetBusinessLicenseImgSrc(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.BusinessLicenseImgSrc = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetBusinessLicenseNumber(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.BusinessLicenseNumber = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetBusinessLicenseType(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.BusinessLicenseType = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetCertifiedFrom(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.CertifiedFrom = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetCertifiedTime(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.CertifiedTime = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetCity(v *QueryEnterpriseInfoResponseBodyProfileInfoCity) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.City = v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetCreateTime(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetDetailAddress(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.DetailAddress = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetEInterpriseSize(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.EInterpriseSize = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetEnterpriseEntity(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.EnterpriseEntity = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetEntityIDNumber(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.EntityIDNumber = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetExtend(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Extend = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetFax(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Fax = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetName(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Name = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetPhone(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Phone = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetPostalCode(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.PostalCode = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetProfile(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Profile = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetProvince(v *QueryEnterpriseInfoResponseBodyProfileInfoProvince) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Province = v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetUpdateTime(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.UpdateTime = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) SetYears(v string) *QueryEnterpriseInfoResponseBodyProfileInfo {
	s.Years = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfo) Validate() error {
	if s.City != nil {
		if err := s.City.Validate(); err != nil {
			return err
		}
	}
	if s.Province != nil {
		if err := s.Province.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEnterpriseInfoResponseBodyProfileInfoCity struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryEnterpriseInfoResponseBodyProfileInfoCity) String() string {
	return dara.Prettify(s)
}

func (s QueryEnterpriseInfoResponseBodyProfileInfoCity) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoCity) GetId() *string {
	return s.Id
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoCity) GetName() *string {
	return s.Name
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoCity) SetId(v string) *QueryEnterpriseInfoResponseBodyProfileInfoCity {
	s.Id = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoCity) SetName(v string) *QueryEnterpriseInfoResponseBodyProfileInfoCity {
	s.Name = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoCity) Validate() error {
	return dara.Validate(s)
}

type QueryEnterpriseInfoResponseBodyProfileInfoProvince struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryEnterpriseInfoResponseBodyProfileInfoProvince) String() string {
	return dara.Prettify(s)
}

func (s QueryEnterpriseInfoResponseBodyProfileInfoProvince) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoProvince) GetId() *string {
	return s.Id
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoProvince) GetName() *string {
	return s.Name
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoProvince) SetId(v string) *QueryEnterpriseInfoResponseBodyProfileInfoProvince {
	s.Id = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoProvince) SetName(v string) *QueryEnterpriseInfoResponseBodyProfileInfoProvince {
	s.Name = &v
	return s
}

func (s *QueryEnterpriseInfoResponseBodyProfileInfoProvince) Validate() error {
	return dara.Validate(s)
}
