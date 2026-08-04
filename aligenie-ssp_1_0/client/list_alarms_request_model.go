// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListAlarmsRequestDeviceInfo) *ListAlarmsRequest
	GetDeviceInfo() *ListAlarmsRequestDeviceInfo
	SetPayload(v *ListAlarmsRequestPayload) *ListAlarmsRequest
	GetPayload() *ListAlarmsRequestPayload
	SetUserInfo(v *ListAlarmsRequestUserInfo) *ListAlarmsRequest
	GetUserInfo() *ListAlarmsRequestUserInfo
}

type ListAlarmsRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *ListAlarmsRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request
	//
	// This parameter is required.
	Payload *ListAlarmsRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *ListAlarmsRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListAlarmsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequest) GetDeviceInfo() *ListAlarmsRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListAlarmsRequest) GetPayload() *ListAlarmsRequestPayload {
	return s.Payload
}

func (s *ListAlarmsRequest) GetUserInfo() *ListAlarmsRequestUserInfo {
	return s.UserInfo
}

func (s *ListAlarmsRequest) SetDeviceInfo(v *ListAlarmsRequestDeviceInfo) *ListAlarmsRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListAlarmsRequest) SetPayload(v *ListAlarmsRequestPayload) *ListAlarmsRequest {
	s.Payload = v
	return s
}

func (s *ListAlarmsRequest) SetUserInfo(v *ListAlarmsRequestUserInfo) *ListAlarmsRequest {
	s.UserInfo = v
	return s
}

func (s *ListAlarmsRequest) Validate() error {
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

type ListAlarmsRequestDeviceInfo struct {
	// The value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME: APK package name, used in the Android application customer link; SKILL_ID: skill ID, used in the cloud link.
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
	// Device ID type: OPEN_ID: default device ID; UNION_ID: organization-level device ID, available only after an organization has been requested on the Maojing Skill Application Open Platform.
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

func (s ListAlarmsRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListAlarmsRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListAlarmsRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListAlarmsRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListAlarmsRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAlarmsRequestDeviceInfo) SetEncodeKey(v string) *ListAlarmsRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetEncodeType(v string) *ListAlarmsRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetId(v string) *ListAlarmsRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetIdType(v string) *ListAlarmsRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) SetOrganizationId(v string) *ListAlarmsRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListAlarmsRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsRequestPayload struct {
	// Current page
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of entries per page: maximum value is 100
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAlarmsRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequestPayload) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestPayload) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAlarmsRequestPayload) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlarmsRequestPayload) SetCurrentPage(v int32) *ListAlarmsRequestPayload {
	s.CurrentPage = &v
	return s
}

func (s *ListAlarmsRequestPayload) SetPageSize(v int32) *ListAlarmsRequestPayload {
	s.PageSize = &v
	return s
}

func (s *ListAlarmsRequestPayload) Validate() error {
	return dara.Validate(s)
}

type ListAlarmsRequestUserInfo struct {
	// The value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME refers to the APK package name, which is the encoding type for the Android application customer link; SKILL_ID refers to the skill ID, which is the encoding type for the cloud link.
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
	// Type of the User ID:
	//
	// - OPEN_ID: The default User ID identity.
	//
	// - UNION_ID: The User ID identity at the organization dimension. This is available only after an organization has been requested on the Maojing Skill Application Open Platform.
	//
	// This parameter is required.
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

func (s ListAlarmsRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListAlarmsRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListAlarmsRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListAlarmsRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListAlarmsRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListAlarmsRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListAlarmsRequestUserInfo) SetEncodeKey(v string) *ListAlarmsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetEncodeType(v string) *ListAlarmsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetId(v string) *ListAlarmsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetIdType(v string) *ListAlarmsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) SetOrganizationId(v string) *ListAlarmsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListAlarmsRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
