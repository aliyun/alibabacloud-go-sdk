// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMusicTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *QueryMusicTypeRequestDeviceInfo) *QueryMusicTypeRequest
	GetDeviceInfo() *QueryMusicTypeRequestDeviceInfo
	SetPayload(v *QueryMusicTypeRequestPayload) *QueryMusicTypeRequest
	GetPayload() *QueryMusicTypeRequestPayload
	SetUserInfo(v *QueryMusicTypeRequestUserInfo) *QueryMusicTypeRequest
	GetUserInfo() *QueryMusicTypeRequestUserInfo
}

type QueryMusicTypeRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *QueryMusicTypeRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request
	Payload *QueryMusicTypeRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *QueryMusicTypeRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s QueryMusicTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequest) GetDeviceInfo() *QueryMusicTypeRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *QueryMusicTypeRequest) GetPayload() *QueryMusicTypeRequestPayload {
	return s.Payload
}

func (s *QueryMusicTypeRequest) GetUserInfo() *QueryMusicTypeRequestUserInfo {
	return s.UserInfo
}

func (s *QueryMusicTypeRequest) SetDeviceInfo(v *QueryMusicTypeRequestDeviceInfo) *QueryMusicTypeRequest {
	s.DeviceInfo = v
	return s
}

func (s *QueryMusicTypeRequest) SetPayload(v *QueryMusicTypeRequestPayload) *QueryMusicTypeRequest {
	s.Payload = v
	return s
}

func (s *QueryMusicTypeRequest) SetUserInfo(v *QueryMusicTypeRequestUserInfo) *QueryMusicTypeRequest {
	s.UserInfo = v
	return s
}

func (s *QueryMusicTypeRequest) Validate() error {
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

type QueryMusicTypeRequestDeviceInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device identity for Maojing, and each way corresponds to a different encoding type: PACKAGE_NAME: APK package name, used for the Android application customer link; SKILL_ID: skill ID, used for the cloud link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// device ID (deviceOpenId or deviceUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the device ID: OPEN_ID: default device ID; UNION_ID: organization-level device ID, available only after requesting an organization in the Maojing Skill Application Open Platform.
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

func (s QueryMusicTypeRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *QueryMusicTypeRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *QueryMusicTypeRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *QueryMusicTypeRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *QueryMusicTypeRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryMusicTypeRequestDeviceInfo) SetEncodeKey(v string) *QueryMusicTypeRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetEncodeType(v string) *QueryMusicTypeRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetId(v string) *QueryMusicTypeRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetIdType(v string) *QueryMusicTypeRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) SetOrganizationId(v string) *QueryMusicTypeRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *QueryMusicTypeRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type QueryMusicTypeRequestPayload struct {
}

func (s QueryMusicTypeRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequestPayload) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestPayload) Validate() error {
	return dara.Validate(s)
}

type QueryMusicTypeRequestUserInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s SkillID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the Maojing user identifier, and each way corresponds to a different encoding type: PACKAGE_NAME: APK package name, used for Android application customer journeys; SKILL_ID: skill ID, used for cloud-based journeys.
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
	// Type of User ID: OPEN_ID: default User ID identity; UNION_ID: organization-dimension User ID identity, available only after an organization has been requested on the Maojing Skill Application Open Platform.
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

func (s QueryMusicTypeRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeRequestUserInfo) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *QueryMusicTypeRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *QueryMusicTypeRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *QueryMusicTypeRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *QueryMusicTypeRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryMusicTypeRequestUserInfo) SetEncodeKey(v string) *QueryMusicTypeRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetEncodeType(v string) *QueryMusicTypeRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetId(v string) *QueryMusicTypeRequestUserInfo {
	s.Id = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetIdType(v string) *QueryMusicTypeRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) SetOrganizationId(v string) *QueryMusicTypeRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *QueryMusicTypeRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
