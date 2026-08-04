// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *DeleteScheduleTaskRequestDeviceInfo) *DeleteScheduleTaskRequest
	GetDeviceInfo() *DeleteScheduleTaskRequestDeviceInfo
	SetPayload(v *DeleteScheduleTaskRequestPayload) *DeleteScheduleTaskRequest
	GetPayload() *DeleteScheduleTaskRequestPayload
	SetUserInfo(v *DeleteScheduleTaskRequestUserInfo) *DeleteScheduleTaskRequest
	GetUserInfo() *DeleteScheduleTaskRequestUserInfo
}

type DeleteScheduleTaskRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *DeleteScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request
	//
	// This parameter is required.
	Payload *DeleteScheduleTaskRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *DeleteScheduleTaskRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeleteScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequest) GetDeviceInfo() *DeleteScheduleTaskRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *DeleteScheduleTaskRequest) GetPayload() *DeleteScheduleTaskRequestPayload {
	return s.Payload
}

func (s *DeleteScheduleTaskRequest) GetUserInfo() *DeleteScheduleTaskRequestUserInfo {
	return s.UserInfo
}

func (s *DeleteScheduleTaskRequest) SetDeviceInfo(v *DeleteScheduleTaskRequestDeviceInfo) *DeleteScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeleteScheduleTaskRequest) SetPayload(v *DeleteScheduleTaskRequestPayload) *DeleteScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *DeleteScheduleTaskRequest) SetUserInfo(v *DeleteScheduleTaskRequestUserInfo) *DeleteScheduleTaskRequest {
	s.UserInfo = v
	return s
}

func (s *DeleteScheduleTaskRequest) Validate() error {
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

type DeleteScheduleTaskRequestDeviceInfo struct {
	// The value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME: APK package name, used in the Android application customer link; SKILL_ID: Skill ID, used in the cloud link.
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
	// Type of the device ID: OPEN_ID: default device ID; UNION_ID: organization-level device ID, available only after an organization has been registered on the Maojing skill application Open Platform.
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

func (s DeleteScheduleTaskRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteScheduleTaskRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetId(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetIdType(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *DeleteScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteScheduleTaskRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type DeleteScheduleTaskRequestPayload struct {
	// ID of the job to delete
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScheduleTaskRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestPayload) GetId() *int64 {
	return s.Id
}

func (s *DeleteScheduleTaskRequestPayload) SetId(v int64) *DeleteScheduleTaskRequestPayload {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestPayload) Validate() error {
	return dara.Validate(s)
}

type DeleteScheduleTaskRequestUserInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the Maojing user identifier, and each way corresponds to a different encoding type: PACKAGE_NAME refers to the APK package name, which is the encoding type for the Android application customer link; SKILL_ID refers to the skill ID, which is the encoding type for the cloud link.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User identifier (userOpenId or userUnionId)
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID: OPEN_ID: default User ID identifier; UNION_ID: organization-dimension User ID identifier, available only after an organization has been requested on the Maojing Skill Application Open Platform.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteScheduleTaskRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *DeleteScheduleTaskRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *DeleteScheduleTaskRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *DeleteScheduleTaskRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *DeleteScheduleTaskRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteScheduleTaskRequestUserInfo) SetEncodeKey(v string) *DeleteScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetEncodeType(v string) *DeleteScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetId(v string) *DeleteScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetIdType(v string) *DeleteScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) SetOrganizationId(v string) *DeleteScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *DeleteScheduleTaskRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
