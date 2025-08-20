// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetPhoneNumberRequestDeviceInfo) *GetPhoneNumberRequest
	GetDeviceInfo() *GetPhoneNumberRequestDeviceInfo
	SetUserInfo(v *GetPhoneNumberRequestUserInfo) *GetPhoneNumberRequest
	GetUserInfo() *GetPhoneNumberRequestUserInfo
}

type GetPhoneNumberRequest struct {
	// This parameter is required.
	DeviceInfo *GetPhoneNumberRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *GetPhoneNumberRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberRequest) GetDeviceInfo() *GetPhoneNumberRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetPhoneNumberRequest) GetUserInfo() *GetPhoneNumberRequestUserInfo {
	return s.UserInfo
}

func (s *GetPhoneNumberRequest) SetDeviceInfo(v *GetPhoneNumberRequestDeviceInfo) *GetPhoneNumberRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetPhoneNumberRequest) SetUserInfo(v *GetPhoneNumberRequestUserInfo) *GetPhoneNumberRequest {
	s.UserInfo = v
	return s
}

func (s *GetPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}

type GetPhoneNumberRequestDeviceInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetPhoneNumberRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetPhoneNumberRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetPhoneNumberRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetPhoneNumberRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetPhoneNumberRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetPhoneNumberRequestDeviceInfo) SetEncodeKey(v string) *GetPhoneNumberRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetEncodeType(v string) *GetPhoneNumberRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetId(v string) *GetPhoneNumberRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetIdType(v string) *GetPhoneNumberRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetOrganizationId(v string) *GetPhoneNumberRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetPhoneNumberRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetPhoneNumberRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetPhoneNumberRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetPhoneNumberRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetPhoneNumberRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetPhoneNumberRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetPhoneNumberRequestUserInfo) SetEncodeKey(v string) *GetPhoneNumberRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetEncodeType(v string) *GetPhoneNumberRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetId(v string) *GetPhoneNumberRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetIdType(v string) *GetPhoneNumberRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetOrganizationId(v string) *GetPhoneNumberRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
