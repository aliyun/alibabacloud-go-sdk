// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetScheduleTaskRequestDeviceInfo) *GetScheduleTaskRequest
	GetDeviceInfo() *GetScheduleTaskRequestDeviceInfo
	SetPayload(v *GetScheduleTaskRequestPayload) *GetScheduleTaskRequest
	GetPayload() *GetScheduleTaskRequestPayload
	SetUserInfo(v *GetScheduleTaskRequestUserInfo) *GetScheduleTaskRequest
	GetUserInfo() *GetScheduleTaskRequestUserInfo
}

type GetScheduleTaskRequest struct {
	// Device ID information
	//
	// This parameter is required.
	DeviceInfo *GetScheduleTaskRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request
	//
	// This parameter is required.
	Payload *GetScheduleTaskRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *GetScheduleTaskRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequest) GetDeviceInfo() *GetScheduleTaskRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetScheduleTaskRequest) GetPayload() *GetScheduleTaskRequestPayload {
	return s.Payload
}

func (s *GetScheduleTaskRequest) GetUserInfo() *GetScheduleTaskRequestUserInfo {
	return s.UserInfo
}

func (s *GetScheduleTaskRequest) SetDeviceInfo(v *GetScheduleTaskRequestDeviceInfo) *GetScheduleTaskRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetScheduleTaskRequest) SetPayload(v *GetScheduleTaskRequestPayload) *GetScheduleTaskRequest {
	s.Payload = v
	return s
}

func (s *GetScheduleTaskRequest) SetUserInfo(v *GetScheduleTaskRequestUserInfo) *GetScheduleTaskRequest {
	s.UserInfo = v
	return s
}

func (s *GetScheduleTaskRequest) Validate() error {
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

type GetScheduleTaskRequestDeviceInfo struct {
	// The value corresponding to the encoding type. If the encoding type is SKILL_ID, the value is the application\\"s SkillID. If the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME: APK package name, used in the Android application customer link; SKILL_ID: skill ID, used in the cloud link.
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
	// Device ID type: OPEN_ID: default device ID; UNION_ID: organization-level device ID, available only after an organization has been requested on the Maojing Skill Application Open Platform.
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

func (s GetScheduleTaskRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetScheduleTaskRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetScheduleTaskRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetScheduleTaskRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetScheduleTaskRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetScheduleTaskRequestDeviceInfo) SetEncodeKey(v string) *GetScheduleTaskRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetEncodeType(v string) *GetScheduleTaskRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetId(v string) *GetScheduleTaskRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetIdType(v string) *GetScheduleTaskRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) SetOrganizationId(v string) *GetScheduleTaskRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetScheduleTaskRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetScheduleTaskRequestPayload struct {
	// ID of the job to query
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetScheduleTaskRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequestPayload) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestPayload) GetId() *int64 {
	return s.Id
}

func (s *GetScheduleTaskRequestPayload) SetId(v int64) *GetScheduleTaskRequestPayload {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetScheduleTaskRequestUserInfo struct {
	// The value corresponding to the encoding type. When the encoding type is SKILLID, the value is the application\\"s SkillID. When the encoding type is PACKAGENAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME indicates the APK package name, used in the Android application customer flow; SKILL_ID indicates the skill ID, used in the cloud-based flow.
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
	// Type of the User ID: OPENID—the default User ID identity; UNIONID—the organization-dimension User ID identity, available only after an organization has been requested on the Maojing Skill Application Open Platform.
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

func (s GetScheduleTaskRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetScheduleTaskRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetScheduleTaskRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetScheduleTaskRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetScheduleTaskRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetScheduleTaskRequestUserInfo) SetEncodeKey(v string) *GetScheduleTaskRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetEncodeType(v string) *GetScheduleTaskRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetId(v string) *GetScheduleTaskRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetIdType(v string) *GetScheduleTaskRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) SetOrganizationId(v string) *GetScheduleTaskRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetScheduleTaskRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
