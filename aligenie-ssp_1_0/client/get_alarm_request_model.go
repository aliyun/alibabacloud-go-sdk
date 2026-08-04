// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetAlarmRequestDeviceInfo) *GetAlarmRequest
	GetDeviceInfo() *GetAlarmRequestDeviceInfo
	SetPayload(v *GetAlarmRequestPayload) *GetAlarmRequest
	GetPayload() *GetAlarmRequestPayload
	SetUserInfo(v *GetAlarmRequestUserInfo) *GetAlarmRequest
	GetUserInfo() *GetAlarmRequestUserInfo
}

type GetAlarmRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *GetAlarmRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request
	//
	// This parameter is required.
	Payload *GetAlarmRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *GetAlarmRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmRequest) GetDeviceInfo() *GetAlarmRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetAlarmRequest) GetPayload() *GetAlarmRequestPayload {
	return s.Payload
}

func (s *GetAlarmRequest) GetUserInfo() *GetAlarmRequestUserInfo {
	return s.UserInfo
}

func (s *GetAlarmRequest) SetDeviceInfo(v *GetAlarmRequestDeviceInfo) *GetAlarmRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetAlarmRequest) SetPayload(v *GetAlarmRequestPayload) *GetAlarmRequest {
	s.Payload = v
	return s
}

func (s *GetAlarmRequest) SetUserInfo(v *GetAlarmRequestUserInfo) *GetAlarmRequest {
	s.UserInfo = v
	return s
}

func (s *GetAlarmRequest) Validate() error {
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

type GetAlarmRequestDeviceInfo struct {
	// Value corresponding to the encoding type: when the encoding type is SKILL_ID, the value is the application\\"s SkillID; when the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding Type. There are multiple ways to obtain the device ID for Maojing, and each method corresponds to a different encoding Type: PACKAGE_NAME: APK package name, used for the Android application Customer link; SKILL_ID: Skill ID, used for the cloud link.
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
	// Type of Device ID:
	//
	// - OPEN_ID: default device ID identity
	//
	// - UNION_ID: organization-dimension device ID identity, available only after an organization has been requested on the Maojing Skill Application Open Platform
	//
	// This parameter is required.
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

func (s GetAlarmRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAlarmRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAlarmRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetAlarmRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAlarmRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAlarmRequestDeviceInfo) SetEncodeKey(v string) *GetAlarmRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetEncodeType(v string) *GetAlarmRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetId(v string) *GetAlarmRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetIdType(v string) *GetAlarmRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) SetOrganizationId(v string) *GetAlarmRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAlarmRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetAlarmRequestPayload struct {
	// Alarm ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
}

func (s GetAlarmRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequestPayload) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestPayload) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *GetAlarmRequestPayload) SetAlarmId(v int64) *GetAlarmRequestPayload {
	s.AlarmId = &v
	return s
}

func (s *GetAlarmRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetAlarmRequestUserInfo struct {
	// Value corresponding to the encoding type. If the encoding type is SKILL_ID, the value is the application\\"s SkillID. If the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the Maojing User Identifier, and each corresponds to a different encoding type:
	//
	// - PACKAGE_NAME: APK package name, used for Android application customer-side flows
	//
	// - SKILL_ID: Skill ID, used for cloud-side flows
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
	// Type of User ID: OPEN_ID: default User ID identifier; UNION_ID: organization-dimension User ID identifier, available only after an organization has been requested on the Maojing Skill Application Open Platform.
	//
	// This parameter is required.
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

func (s GetAlarmRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetAlarmRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAlarmRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAlarmRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetAlarmRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAlarmRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAlarmRequestUserInfo) SetEncodeKey(v string) *GetAlarmRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetEncodeType(v string) *GetAlarmRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetId(v string) *GetAlarmRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetIdType(v string) *GetAlarmRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetAlarmRequestUserInfo) SetOrganizationId(v string) *GetAlarmRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAlarmRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
