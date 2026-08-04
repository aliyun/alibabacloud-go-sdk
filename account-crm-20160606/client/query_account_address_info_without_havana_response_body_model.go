// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountAddressInfoWithoutHavanaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountAddressInfoWithoutHavanaResponseBody
	GetCode() *string
	SetMessage(v string) *QueryAccountAddressInfoWithoutHavanaResponseBody
	GetMessage() *string
	SetProfileInfo(v *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) *QueryAccountAddressInfoWithoutHavanaResponseBody
	GetProfileInfo() *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo
	SetRequestId(v string) *QueryAccountAddressInfoWithoutHavanaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountAddressInfoWithoutHavanaResponseBody
	GetSuccess() *bool
}

type QueryAccountAddressInfoWithoutHavanaResponseBody struct {
	Code        *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	ProfileInfo *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo `json:"ProfileInfo,omitempty" xml:"ProfileInfo,omitempty" type:"Struct"`
	RequestId   *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) GetProfileInfo() *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	return s.ProfileInfo
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) SetCode(v string) *QueryAccountAddressInfoWithoutHavanaResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) SetMessage(v string) *QueryAccountAddressInfoWithoutHavanaResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) SetProfileInfo(v *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) *QueryAccountAddressInfoWithoutHavanaResponseBody {
	s.ProfileInfo = v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) SetRequestId(v string) *QueryAccountAddressInfoWithoutHavanaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) SetSuccess(v bool) *QueryAccountAddressInfoWithoutHavanaResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBody) Validate() error {
	if s.ProfileInfo != nil {
		if err := s.ProfileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo struct {
	AccountAttr *string                                                              `json:"AccountAttr,omitempty" xml:"AccountAttr,omitempty"`
	Address     *string                                                              `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2    *string                                                              `json:"Address2,omitempty" xml:"Address2,omitempty"`
	Address3    *string                                                              `json:"Address3,omitempty" xml:"Address3,omitempty"`
	Address4    *string                                                              `json:"Address4,omitempty" xml:"Address4,omitempty"`
	Address5    *string                                                              `json:"Address5,omitempty" xml:"Address5,omitempty"`
	Address6    *string                                                              `json:"Address6,omitempty" xml:"Address6,omitempty"`
	City        *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity     `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	HavanaId    *string                                                              `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	PostCode    *string                                                              `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	Province    *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	TrueName    *string                                                              `json:"TrueName,omitempty" xml:"TrueName,omitempty"`
	Version     *string                                                              `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetAccountAttr() *string {
	return s.AccountAttr
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetAddress() *string {
	return s.Address
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetAddress2() *string {
	return s.Address2
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetAddress3() *string {
	return s.Address3
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetAddress4() *string {
	return s.Address4
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetAddress5() *string {
	return s.Address5
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetAddress6() *string {
	return s.Address6
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetCity() *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity {
	return s.City
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetPostCode() *string {
	return s.PostCode
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetProvince() *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince {
	return s.Province
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetTrueName() *string {
	return s.TrueName
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) GetVersion() *string {
	return s.Version
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetAccountAttr(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.AccountAttr = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetAddress(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Address = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetAddress2(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Address2 = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetAddress3(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Address3 = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetAddress4(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Address4 = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetAddress5(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Address5 = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetAddress6(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Address6 = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetCity(v *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.City = v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetHavanaId(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.HavanaId = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetPostCode(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.PostCode = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetProvince(v *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Province = v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetTrueName(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.TrueName = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) SetVersion(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo {
	s.Version = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfo) Validate() error {
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

type QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) GetId() *string {
	return s.Id
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) GetName() *string {
	return s.Name
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) SetId(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity {
	s.Id = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) SetName(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity {
	s.Name = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoCity) Validate() error {
	return dara.Validate(s)
}

type QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) GetId() *string {
	return s.Id
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) GetName() *string {
	return s.Name
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) SetId(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince {
	s.Id = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) SetName(v string) *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince {
	s.Name = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponseBodyProfileInfoProvince) Validate() error {
	return dara.Validate(s)
}
