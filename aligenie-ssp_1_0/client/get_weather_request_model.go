// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWeatherRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetWeatherRequestDeviceInfo) *GetWeatherRequest
	GetDeviceInfo() *GetWeatherRequestDeviceInfo
	SetPayload(v *GetWeatherRequestPayload) *GetWeatherRequest
	GetPayload() *GetWeatherRequestPayload
	SetUserInfo(v *GetWeatherRequestUserInfo) *GetWeatherRequest
	GetUserInfo() *GetWeatherRequestUserInfo
}

type GetWeatherRequest struct {
	// Device ID information
	//
	// This parameter is required.
	DeviceInfo *GetWeatherRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request
	//
	// if can be null:
	// false
	Payload *GetWeatherRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *GetWeatherRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetWeatherRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequest) GoString() string {
	return s.String()
}

func (s *GetWeatherRequest) GetDeviceInfo() *GetWeatherRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetWeatherRequest) GetPayload() *GetWeatherRequestPayload {
	return s.Payload
}

func (s *GetWeatherRequest) GetUserInfo() *GetWeatherRequestUserInfo {
	return s.UserInfo
}

func (s *GetWeatherRequest) SetDeviceInfo(v *GetWeatherRequestDeviceInfo) *GetWeatherRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetWeatherRequest) SetPayload(v *GetWeatherRequestPayload) *GetWeatherRequest {
	s.Payload = v
	return s
}

func (s *GetWeatherRequest) SetUserInfo(v *GetWeatherRequestUserInfo) *GetWeatherRequest {
	s.UserInfo = v
	return s
}

func (s *GetWeatherRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
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

type GetWeatherRequestDeviceInfo struct {
	// The value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID of Maojing, and each way corresponds to a different encoding type: PACKAGE_NAME refers to the APK package name, which is the encoding type for the Android application customer link; SKILL_ID refers to the skill ID, which is the encoding type for the cloud link.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID (deviceOpenId or deviceUnionId)
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the device ID: OPEN_ID is the default device ID identity; UNION_ID is the organization-dimension device ID identity, which is available only after an organization has been requested on the Maojing Skill Application Open Platform.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This field is required when IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetWeatherRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetWeatherRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetWeatherRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetWeatherRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetWeatherRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetWeatherRequestDeviceInfo) SetEncodeKey(v string) *GetWeatherRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetEncodeType(v string) *GetWeatherRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetId(v string) *GetWeatherRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetIdType(v string) *GetWeatherRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) SetOrganizationId(v string) *GetWeatherRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetWeatherRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetWeatherRequestPayload struct {
}

func (s GetWeatherRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequestPayload) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetWeatherRequestUserInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME: APK package name, used for Android application customer links; SKILL_ID: skill ID, used for cloud-based links.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier (userOpenId or userUnionId)
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of user ID: OPEN_ID: default user ID identifier; UNION_ID: organization-dimension user ID identifier, available only after an organization has been requested on the Maojing Skill Application Open Platform.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetWeatherRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetWeatherRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetWeatherRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetWeatherRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetWeatherRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetWeatherRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetWeatherRequestUserInfo) SetEncodeKey(v string) *GetWeatherRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetEncodeType(v string) *GetWeatherRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetId(v string) *GetWeatherRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetIdType(v string) *GetWeatherRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetWeatherRequestUserInfo) SetOrganizationId(v string) *GetWeatherRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetWeatherRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
