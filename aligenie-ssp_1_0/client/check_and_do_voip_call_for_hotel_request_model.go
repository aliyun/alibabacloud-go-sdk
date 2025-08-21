// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAndDoVoipCallForHotelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizData(v string) *CheckAndDoVoipCallForHotelRequest
	GetBizData() *string
	SetCalleeNick(v string) *CheckAndDoVoipCallForHotelRequest
	GetCalleeNick() *string
	SetCalleePhoneNum(v string) *CheckAndDoVoipCallForHotelRequest
	GetCalleePhoneNum() *string
	SetDeviceInfo(v *CheckAndDoVoipCallForHotelRequestDeviceInfo) *CheckAndDoVoipCallForHotelRequest
	GetDeviceInfo() *CheckAndDoVoipCallForHotelRequestDeviceInfo
	SetUserInfo(v *CheckAndDoVoipCallForHotelRequestUserInfo) *CheckAndDoVoipCallForHotelRequest
	GetUserInfo() *CheckAndDoVoipCallForHotelRequestUserInfo
}

type CheckAndDoVoipCallForHotelRequest struct {
	BizData        *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	CalleeNick     *string `json:"CalleeNick,omitempty" xml:"CalleeNick,omitempty"`
	CalleePhoneNum *string `json:"CalleePhoneNum,omitempty" xml:"CalleePhoneNum,omitempty"`
	// This parameter is required.
	DeviceInfo *CheckAndDoVoipCallForHotelRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *CheckAndDoVoipCallForHotelRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CheckAndDoVoipCallForHotelRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelRequest) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelRequest) GetBizData() *string {
	return s.BizData
}

func (s *CheckAndDoVoipCallForHotelRequest) GetCalleeNick() *string {
	return s.CalleeNick
}

func (s *CheckAndDoVoipCallForHotelRequest) GetCalleePhoneNum() *string {
	return s.CalleePhoneNum
}

func (s *CheckAndDoVoipCallForHotelRequest) GetDeviceInfo() *CheckAndDoVoipCallForHotelRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *CheckAndDoVoipCallForHotelRequest) GetUserInfo() *CheckAndDoVoipCallForHotelRequestUserInfo {
	return s.UserInfo
}

func (s *CheckAndDoVoipCallForHotelRequest) SetBizData(v string) *CheckAndDoVoipCallForHotelRequest {
	s.BizData = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequest) SetCalleeNick(v string) *CheckAndDoVoipCallForHotelRequest {
	s.CalleeNick = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequest) SetCalleePhoneNum(v string) *CheckAndDoVoipCallForHotelRequest {
	s.CalleePhoneNum = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequest) SetDeviceInfo(v *CheckAndDoVoipCallForHotelRequestDeviceInfo) *CheckAndDoVoipCallForHotelRequest {
	s.DeviceInfo = v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequest) SetUserInfo(v *CheckAndDoVoipCallForHotelRequestUserInfo) *CheckAndDoVoipCallForHotelRequest {
	s.UserInfo = v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequest) Validate() error {
	return dara.Validate(s)
}

type CheckAndDoVoipCallForHotelRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CheckAndDoVoipCallForHotelRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) SetEncodeKey(v string) *CheckAndDoVoipCallForHotelRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) SetEncodeType(v string) *CheckAndDoVoipCallForHotelRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) SetId(v string) *CheckAndDoVoipCallForHotelRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) SetIdType(v string) *CheckAndDoVoipCallForHotelRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) SetOrganizationId(v string) *CheckAndDoVoipCallForHotelRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type CheckAndDoVoipCallForHotelRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CheckAndDoVoipCallForHotelRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) SetEncodeKey(v string) *CheckAndDoVoipCallForHotelRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) SetEncodeType(v string) *CheckAndDoVoipCallForHotelRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) SetId(v string) *CheckAndDoVoipCallForHotelRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) SetIdType(v string) *CheckAndDoVoipCallForHotelRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) SetOrganizationId(v string) *CheckAndDoVoipCallForHotelRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
