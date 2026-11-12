// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListRemindersRequestDeviceInfo) *ListRemindersRequest
	GetDeviceInfo() *ListRemindersRequestDeviceInfo
	SetPayload(v *ListRemindersRequestPayload) *ListRemindersRequest
	GetPayload() *ListRemindersRequestPayload
	SetUserInfo(v *ListRemindersRequestUserInfo) *ListRemindersRequest
	GetUserInfo() *ListRemindersRequestUserInfo
}

type ListRemindersRequest struct {
	// List of device identity information.
	//
	// This parameter is required.
	DeviceInfo *ListRemindersRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameters for the service request.
	//
	// This parameter is required.
	Payload *ListRemindersRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// List of user identity information.
	//
	// This parameter is required.
	UserInfo *ListRemindersRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListRemindersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequest) GoString() string {
	return s.String()
}

func (s *ListRemindersRequest) GetDeviceInfo() *ListRemindersRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListRemindersRequest) GetPayload() *ListRemindersRequestPayload {
	return s.Payload
}

func (s *ListRemindersRequest) GetUserInfo() *ListRemindersRequestUserInfo {
	return s.UserInfo
}

func (s *ListRemindersRequest) SetDeviceInfo(v *ListRemindersRequestDeviceInfo) *ListRemindersRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListRemindersRequest) SetPayload(v *ListRemindersRequestPayload) *ListRemindersRequest {
	s.Payload = v
	return s
}

func (s *ListRemindersRequest) SetUserInfo(v *ListRemindersRequestUserInfo) *ListRemindersRequest {
	s.UserInfo = v
	return s
}

func (s *ListRemindersRequest) Validate() error {
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

type ListRemindersRequestDeviceInfo struct {
	// The value corresponding to the encoding type.
	//
	// - When the encoding type is SKILLID, the value is the application\\"s Skill ID.
	//
	// - When the encoding type is PACKAGENAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used in the Android application customer flow.
	//
	// - SKILL_ID: Skill ID, used in the cloud-based flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID, either deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the device ID.
	//
	// - OPEN_ID: The default device ID.
	//
	// - UNION_ID: Organization-level device ID. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when **IdType*	- is set to **UNION_ID**.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRemindersRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRemindersRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRemindersRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListRemindersRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRemindersRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRemindersRequestDeviceInfo) SetEncodeKey(v string) *ListRemindersRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetEncodeType(v string) *ListRemindersRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetId(v string) *ListRemindersRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetIdType(v string) *ListRemindersRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetOrganizationId(v string) *ListRemindersRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListRemindersRequestPayload struct {
	// Indicates whether this invocation is for debug purposes.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s ListRemindersRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequestPayload) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestPayload) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *ListRemindersRequestPayload) SetIsDebug(v bool) *ListRemindersRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *ListRemindersRequestPayload) Validate() error {
	return dara.Validate(s)
}

type ListRemindersRequestUserInfo struct {
	// The value corresponding to the encoding type:
	//
	//  - When the encoding type is SKILLID, the value is the application\\"s Skill ID.
	//
	//  - When the encoding type is PACKAGENAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used for the Android application customer ingest endpoint.
	//
	// - SKILL_ID: Skill ID, used for the cloud-side ingest endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier, either userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the User ID:
	//
	//  - OPEN_ID: The default User ID identity.
	//
	//  - UNION_ID: The User ID identity at the organization dimension. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This parameter is required when **IdType*	- is configured as **UNION_ID**.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRemindersRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRemindersRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRemindersRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListRemindersRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRemindersRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRemindersRequestUserInfo) SetEncodeKey(v string) *ListRemindersRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetEncodeType(v string) *ListRemindersRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetId(v string) *ListRemindersRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetIdType(v string) *ListRemindersRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetOrganizationId(v string) *ListRemindersRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRemindersRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
