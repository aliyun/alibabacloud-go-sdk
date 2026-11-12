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
	// Device identification information
	//
	// This parameter is required.
	DeviceInfo *GetPhoneNumberRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// User identifier information
	//
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
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhoneNumberRequestDeviceInfo struct {
	// Value corresponding to the encoding type. If the encoding type is `SKILL_ID`, the value is the application\\"s Skill ID. If the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the Maojing device ID, and each method corresponds to a different encoding type:
	//
	// - PACKAGE_NAME: APK package name, used in the Android application customer flow
	//
	// - SKILL_ID: Skill ID, used in the cloud-based flow
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID (deviceOpenId or deviceUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of device ID:
	//
	// - OPEN_ID: Default device ID
	//
	// - UNION_ID: Organization-dimension device ID, available only after an organization has been requested on the Maojing Skill Application Open Platform
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is `UNION_ID`.
	//
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
	// Value corresponding to the encoding type. When the encoding type is `SKILL_ID`, the value is the application\\"s Skill ID. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Maojing, and each method corresponds to a different encoding type:
	//
	// - PACKAGE_NAME: APK package name, used for Android application customer journeys
	//
	// - SKILL_ID: Skill ID, used for cloud-based journeys
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier (userOpenId or userUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID:
	//
	// - OPEN_ID: Default User ID identifier
	//
	// - UNION_ID: Organization-dimension User ID identifier, available only after an organization has been requested on the Maojing Skill Application Open Platform
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is `UNION_ID`.
	//
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
