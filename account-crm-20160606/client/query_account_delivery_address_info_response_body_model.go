// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountDeliveryAddressInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountDeliveryAddressInfoResponseBody
	GetCode() *string
	SetData(v []*QueryAccountDeliveryAddressInfoResponseBodyData) *QueryAccountDeliveryAddressInfoResponseBody
	GetData() []*QueryAccountDeliveryAddressInfoResponseBodyData
	SetMessage(v string) *QueryAccountDeliveryAddressInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAccountDeliveryAddressInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountDeliveryAddressInfoResponseBody
	GetSuccess() *bool
}

type QueryAccountDeliveryAddressInfoResponseBody struct {
	Code      *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryAccountDeliveryAddressInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) GetData() []*QueryAccountDeliveryAddressInfoResponseBodyData {
	return s.Data
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) SetCode(v string) *QueryAccountDeliveryAddressInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) SetData(v []*QueryAccountDeliveryAddressInfoResponseBodyData) *QueryAccountDeliveryAddressInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) SetMessage(v string) *QueryAccountDeliveryAddressInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) SetRequestId(v string) *QueryAccountDeliveryAddressInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) SetSuccess(v bool) *QueryAccountDeliveryAddressInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAccountDeliveryAddressInfoResponseBodyData struct {
	Address         *string                                                         `json:"Address,omitempty" xml:"Address,omitempty"`
	AreaDivision    *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision    `json:"AreaDivision,omitempty" xml:"AreaDivision,omitempty" type:"Struct"`
	AreaId          *string                                                         `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	CityDivision    *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision    `json:"CityDivision,omitempty" xml:"CityDivision,omitempty" type:"Struct"`
	CityId          *string                                                         `json:"CityId,omitempty" xml:"CityId,omitempty"`
	Contacts        *string                                                         `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	DefaultAddress  *bool                                                           `json:"DefaultAddress,omitempty" xml:"DefaultAddress,omitempty"`
	Email           *string                                                         `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile          *string                                                         `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Phone           *string                                                         `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Pk              *string                                                         `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Postalcode      *string                                                         `json:"Postalcode,omitempty" xml:"Postalcode,omitempty"`
	ProviceDivision *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision `json:"ProviceDivision,omitempty" xml:"ProviceDivision,omitempty" type:"Struct"`
	ProviceId       *string                                                         `json:"ProviceId,omitempty" xml:"ProviceId,omitempty"`
	TownDivision    *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision    `json:"TownDivision,omitempty" xml:"TownDivision,omitempty" type:"Struct"`
	TownId          *string                                                         `json:"TownId,omitempty" xml:"TownId,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetAreaDivision() *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	return s.AreaDivision
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetAreaId() *string {
	return s.AreaId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetCityDivision() *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	return s.CityDivision
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetCityId() *string {
	return s.CityId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetContacts() *string {
	return s.Contacts
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetDefaultAddress() *bool {
	return s.DefaultAddress
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetPhone() *string {
	return s.Phone
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetPk() *string {
	return s.Pk
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetPostalcode() *string {
	return s.Postalcode
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetProviceDivision() *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	return s.ProviceDivision
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetProviceId() *string {
	return s.ProviceId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetTownDivision() *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	return s.TownDivision
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) GetTownId() *string {
	return s.TownId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetAddress(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.Address = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetAreaDivision(v *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.AreaDivision = v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetAreaId(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.AreaId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetCityDivision(v *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.CityDivision = v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetCityId(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.CityId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetContacts(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.Contacts = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetDefaultAddress(v bool) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.DefaultAddress = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetEmail(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetMobile(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetPhone(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.Phone = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetPk(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.Pk = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetPostalcode(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.Postalcode = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetProviceDivision(v *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.ProviceDivision = v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetProviceId(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.ProviceId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetTownDivision(v *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.TownDivision = v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) SetTownId(v string) *QueryAccountDeliveryAddressInfoResponseBodyData {
	s.TownId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyData) Validate() error {
	if s.AreaDivision != nil {
		if err := s.AreaDivision.Validate(); err != nil {
			return err
		}
	}
	if s.CityDivision != nil {
		if err := s.CityDivision.Validate(); err != nil {
			return err
		}
	}
	if s.ProviceDivision != nil {
		if err := s.ProviceDivision.Validate(); err != nil {
			return err
		}
	}
	if s.TownDivision != nil {
		if err := s.TownDivision.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision struct {
	DivisionAbbName *string `json:"DivisionAbbName,omitempty" xml:"DivisionAbbName,omitempty"`
	DivisionId      *int64  `json:"DivisionId,omitempty" xml:"DivisionId,omitempty"`
	DivisionLevel   *int64  `json:"DivisionLevel,omitempty" xml:"DivisionLevel,omitempty"`
	DivisionName    *string `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
	DivisionTname   *string `json:"DivisionTname,omitempty" xml:"DivisionTname,omitempty"`
	NewDivisionId   *int64  `json:"NewDivisionId,omitempty" xml:"NewDivisionId,omitempty"`
	ParentId        *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Pinyin          *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetDivisionAbbName() *string {
	return s.DivisionAbbName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetDivisionId() *int64 {
	return s.DivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetDivisionLevel() *int64 {
	return s.DivisionLevel
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetDivisionName() *string {
	return s.DivisionName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetDivisionTname() *string {
	return s.DivisionTname
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetNewDivisionId() *int64 {
	return s.NewDivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetPinyin() *string {
	return s.Pinyin
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) GetRemark() *string {
	return s.Remark
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetDivisionAbbName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.DivisionAbbName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.DivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetDivisionLevel(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.DivisionLevel = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetDivisionName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.DivisionName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetDivisionTname(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.DivisionTname = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetNewDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.NewDivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetParentId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.ParentId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetPinyin(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.Pinyin = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) SetRemark(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision {
	s.Remark = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataAreaDivision) Validate() error {
	return dara.Validate(s)
}

type QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision struct {
	DivisionAbbName *string `json:"DivisionAbbName,omitempty" xml:"DivisionAbbName,omitempty"`
	DivisionId      *int64  `json:"DivisionId,omitempty" xml:"DivisionId,omitempty"`
	DivisionLevel   *int64  `json:"DivisionLevel,omitempty" xml:"DivisionLevel,omitempty"`
	DivisionName    *string `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
	DivisionTname   *string `json:"DivisionTname,omitempty" xml:"DivisionTname,omitempty"`
	NewDivisionId   *int64  `json:"NewDivisionId,omitempty" xml:"NewDivisionId,omitempty"`
	ParentId        *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Pinyin          *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetDivisionAbbName() *string {
	return s.DivisionAbbName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetDivisionId() *int64 {
	return s.DivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetDivisionLevel() *int64 {
	return s.DivisionLevel
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetDivisionName() *string {
	return s.DivisionName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetDivisionTname() *string {
	return s.DivisionTname
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetNewDivisionId() *int64 {
	return s.NewDivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetPinyin() *string {
	return s.Pinyin
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) GetRemark() *string {
	return s.Remark
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetDivisionAbbName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.DivisionAbbName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.DivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetDivisionLevel(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.DivisionLevel = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetDivisionName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.DivisionName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetDivisionTname(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.DivisionTname = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetNewDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.NewDivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetParentId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.ParentId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetPinyin(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.Pinyin = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) SetRemark(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision {
	s.Remark = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataCityDivision) Validate() error {
	return dara.Validate(s)
}

type QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision struct {
	DivisionAbbName *string `json:"DivisionAbbName,omitempty" xml:"DivisionAbbName,omitempty"`
	DivisionId      *int64  `json:"DivisionId,omitempty" xml:"DivisionId,omitempty"`
	DivisionLevel   *int64  `json:"DivisionLevel,omitempty" xml:"DivisionLevel,omitempty"`
	DivisionName    *string `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
	DivisionTname   *string `json:"DivisionTname,omitempty" xml:"DivisionTname,omitempty"`
	NewDivisionId   *int64  `json:"NewDivisionId,omitempty" xml:"NewDivisionId,omitempty"`
	ParentId        *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Pinyin          *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetDivisionAbbName() *string {
	return s.DivisionAbbName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetDivisionId() *int64 {
	return s.DivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetDivisionLevel() *int64 {
	return s.DivisionLevel
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetDivisionName() *string {
	return s.DivisionName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetDivisionTname() *string {
	return s.DivisionTname
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetNewDivisionId() *int64 {
	return s.NewDivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetPinyin() *string {
	return s.Pinyin
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) GetRemark() *string {
	return s.Remark
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetDivisionAbbName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.DivisionAbbName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.DivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetDivisionLevel(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.DivisionLevel = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetDivisionName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.DivisionName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetDivisionTname(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.DivisionTname = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetNewDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.NewDivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetParentId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.ParentId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetPinyin(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.Pinyin = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) SetRemark(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision {
	s.Remark = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataProviceDivision) Validate() error {
	return dara.Validate(s)
}

type QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision struct {
	DivisionAbbName *string `json:"DivisionAbbName,omitempty" xml:"DivisionAbbName,omitempty"`
	DivisionId      *int64  `json:"DivisionId,omitempty" xml:"DivisionId,omitempty"`
	DivisionLevel   *int64  `json:"DivisionLevel,omitempty" xml:"DivisionLevel,omitempty"`
	DivisionName    *string `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
	DivisionTname   *string `json:"DivisionTname,omitempty" xml:"DivisionTname,omitempty"`
	NewDivisionId   *int64  `json:"NewDivisionId,omitempty" xml:"NewDivisionId,omitempty"`
	ParentId        *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Pinyin          *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetDivisionAbbName() *string {
	return s.DivisionAbbName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetDivisionId() *int64 {
	return s.DivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetDivisionLevel() *int64 {
	return s.DivisionLevel
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetDivisionName() *string {
	return s.DivisionName
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetDivisionTname() *string {
	return s.DivisionTname
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetNewDivisionId() *int64 {
	return s.NewDivisionId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetParentId() *int64 {
	return s.ParentId
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetPinyin() *string {
	return s.Pinyin
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) GetRemark() *string {
	return s.Remark
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetDivisionAbbName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.DivisionAbbName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.DivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetDivisionLevel(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.DivisionLevel = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetDivisionName(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.DivisionName = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetDivisionTname(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.DivisionTname = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetNewDivisionId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.NewDivisionId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetParentId(v int64) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.ParentId = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetPinyin(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.Pinyin = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) SetRemark(v string) *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision {
	s.Remark = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponseBodyDataTownDivision) Validate() error {
	return dara.Validate(s)
}
