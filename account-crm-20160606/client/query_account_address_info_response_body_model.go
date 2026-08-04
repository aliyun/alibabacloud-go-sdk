// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountAddressInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountAddressInfoResponseBody
	GetCode() *string
	SetMessage(v string) *QueryAccountAddressInfoResponseBody
	GetMessage() *string
	SetProfileInfo(v *QueryAccountAddressInfoResponseBodyProfileInfo) *QueryAccountAddressInfoResponseBody
	GetProfileInfo() *QueryAccountAddressInfoResponseBodyProfileInfo
	SetRequestId(v string) *QueryAccountAddressInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountAddressInfoResponseBody
	GetSuccess() *bool
}

type QueryAccountAddressInfoResponseBody struct {
	Code        *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	ProfileInfo *QueryAccountAddressInfoResponseBodyProfileInfo `json:"ProfileInfo,omitempty" xml:"ProfileInfo,omitempty" type:"Struct"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountAddressInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountAddressInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountAddressInfoResponseBody) GetProfileInfo() *QueryAccountAddressInfoResponseBodyProfileInfo {
	return s.ProfileInfo
}

func (s *QueryAccountAddressInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountAddressInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountAddressInfoResponseBody) SetCode(v string) *QueryAccountAddressInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBody) SetMessage(v string) *QueryAccountAddressInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBody) SetProfileInfo(v *QueryAccountAddressInfoResponseBodyProfileInfo) *QueryAccountAddressInfoResponseBody {
	s.ProfileInfo = v
	return s
}

func (s *QueryAccountAddressInfoResponseBody) SetRequestId(v string) *QueryAccountAddressInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBody) SetSuccess(v bool) *QueryAccountAddressInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBody) Validate() error {
	if s.ProfileInfo != nil {
		if err := s.ProfileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountAddressInfoResponseBodyProfileInfo struct {
	AccountAttr                    *string                                                 `json:"AccountAttr,omitempty" xml:"AccountAttr,omitempty"`
	Address                        *string                                                 `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2                       *string                                                 `json:"Address2,omitempty" xml:"Address2,omitempty"`
	Address3                       *string                                                 `json:"Address3,omitempty" xml:"Address3,omitempty"`
	Address4                       *string                                                 `json:"Address4,omitempty" xml:"Address4,omitempty"`
	Address5                       *string                                                 `json:"Address5,omitempty" xml:"Address5,omitempty"`
	Address6                       *string                                                 `json:"Address6,omitempty" xml:"Address6,omitempty"`
	City                           *QueryAccountAddressInfoResponseBodyProfileInfoCity     `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	District                       *QueryAccountAddressInfoResponseBodyProfileInfoDistrict `json:"District,omitempty" xml:"District,omitempty" type:"Struct"`
	Email                          *string                                                 `json:"Email,omitempty" xml:"Email,omitempty"`
	HavanaId                       *string                                                 `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	NationalityCode                *string                                                 `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	PostCode                       *string                                                 `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	Province                       *QueryAccountAddressInfoResponseBodyProfileInfoProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	SelfServicingBusinessRegNum    *string                                                 `json:"SelfServicingBusinessRegNum,omitempty" xml:"SelfServicingBusinessRegNum,omitempty"`
	SelfServicingIdentificationNum *string                                                 `json:"SelfServicingIdentificationNum,omitempty" xml:"SelfServicingIdentificationNum,omitempty"`
	TrueName                       *string                                                 `json:"TrueName,omitempty" xml:"TrueName,omitempty"`
	Version                        *string                                                 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryAccountAddressInfoResponseBodyProfileInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoResponseBodyProfileInfo) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetAccountAttr() *string {
	return s.AccountAttr
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetAddress() *string {
	return s.Address
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetAddress2() *string {
	return s.Address2
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetAddress3() *string {
	return s.Address3
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetAddress4() *string {
	return s.Address4
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetAddress5() *string {
	return s.Address5
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetAddress6() *string {
	return s.Address6
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetCity() *QueryAccountAddressInfoResponseBodyProfileInfoCity {
	return s.City
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetDistrict() *QueryAccountAddressInfoResponseBodyProfileInfoDistrict {
	return s.District
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetEmail() *string {
	return s.Email
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetPostCode() *string {
	return s.PostCode
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetProvince() *QueryAccountAddressInfoResponseBodyProfileInfoProvince {
	return s.Province
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetSelfServicingBusinessRegNum() *string {
	return s.SelfServicingBusinessRegNum
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetSelfServicingIdentificationNum() *string {
	return s.SelfServicingIdentificationNum
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetTrueName() *string {
	return s.TrueName
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) GetVersion() *string {
	return s.Version
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetAccountAttr(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.AccountAttr = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetAddress(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Address = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetAddress2(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Address2 = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetAddress3(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Address3 = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetAddress4(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Address4 = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetAddress5(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Address5 = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetAddress6(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Address6 = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetCity(v *QueryAccountAddressInfoResponseBodyProfileInfoCity) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.City = v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetDistrict(v *QueryAccountAddressInfoResponseBodyProfileInfoDistrict) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.District = v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetEmail(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Email = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetHavanaId(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.HavanaId = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetNationalityCode(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.NationalityCode = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetPostCode(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.PostCode = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetProvince(v *QueryAccountAddressInfoResponseBodyProfileInfoProvince) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Province = v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetSelfServicingBusinessRegNum(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.SelfServicingBusinessRegNum = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetSelfServicingIdentificationNum(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.SelfServicingIdentificationNum = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetTrueName(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.TrueName = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) SetVersion(v string) *QueryAccountAddressInfoResponseBodyProfileInfo {
	s.Version = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfo) Validate() error {
	if s.City != nil {
		if err := s.City.Validate(); err != nil {
			return err
		}
	}
	if s.District != nil {
		if err := s.District.Validate(); err != nil {
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

type QueryAccountAddressInfoResponseBodyProfileInfoCity struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountAddressInfoResponseBodyProfileInfoCity) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoResponseBodyProfileInfoCity) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoCity) GetId() *string {
	return s.Id
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoCity) GetName() *string {
	return s.Name
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoCity) SetId(v string) *QueryAccountAddressInfoResponseBodyProfileInfoCity {
	s.Id = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoCity) SetName(v string) *QueryAccountAddressInfoResponseBodyProfileInfoCity {
	s.Name = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoCity) Validate() error {
	return dara.Validate(s)
}

type QueryAccountAddressInfoResponseBodyProfileInfoDistrict struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountAddressInfoResponseBodyProfileInfoDistrict) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoResponseBodyProfileInfoDistrict) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoDistrict) GetId() *string {
	return s.Id
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoDistrict) GetName() *string {
	return s.Name
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoDistrict) SetId(v string) *QueryAccountAddressInfoResponseBodyProfileInfoDistrict {
	s.Id = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoDistrict) SetName(v string) *QueryAccountAddressInfoResponseBodyProfileInfoDistrict {
	s.Name = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoDistrict) Validate() error {
	return dara.Validate(s)
}

type QueryAccountAddressInfoResponseBodyProfileInfoProvince struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountAddressInfoResponseBodyProfileInfoProvince) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoResponseBodyProfileInfoProvince) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoProvince) GetId() *string {
	return s.Id
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoProvince) GetName() *string {
	return s.Name
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoProvince) SetId(v string) *QueryAccountAddressInfoResponseBodyProfileInfoProvince {
	s.Id = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoProvince) SetName(v string) *QueryAccountAddressInfoResponseBodyProfileInfoProvince {
	s.Name = &v
	return s
}

func (s *QueryAccountAddressInfoResponseBodyProfileInfoProvince) Validate() error {
	return dara.Validate(s)
}
