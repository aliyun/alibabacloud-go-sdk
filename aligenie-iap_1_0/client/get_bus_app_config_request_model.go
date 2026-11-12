// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetBusAppConfigRequestDeviceInfo) *GetBusAppConfigRequest
	GetDeviceInfo() *GetBusAppConfigRequestDeviceInfo
	SetPayload(v *GetBusAppConfigRequestPayload) *GetBusAppConfigRequest
	GetPayload() *GetBusAppConfigRequestPayload
	SetUserInfo(v *GetBusAppConfigRequestUserInfo) *GetBusAppConfigRequest
	GetUserInfo() *GetBusAppConfigRequestUserInfo
}

type GetBusAppConfigRequest struct {
	// Device ID information
	DeviceInfo *GetBusAppConfigRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	Payload *GetBusAppConfigRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identity information
	//
	// This parameter is required.
	UserInfo *GetBusAppConfigRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetBusAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequest) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequest) GetDeviceInfo() *GetBusAppConfigRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetBusAppConfigRequest) GetPayload() *GetBusAppConfigRequestPayload {
	return s.Payload
}

func (s *GetBusAppConfigRequest) GetUserInfo() *GetBusAppConfigRequestUserInfo {
	return s.UserInfo
}

func (s *GetBusAppConfigRequest) SetDeviceInfo(v *GetBusAppConfigRequestDeviceInfo) *GetBusAppConfigRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetBusAppConfigRequest) SetPayload(v *GetBusAppConfigRequestPayload) *GetBusAppConfigRequest {
	s.Payload = v
	return s
}

func (s *GetBusAppConfigRequest) SetUserInfo(v *GetBusAppConfigRequestUserInfo) *GetBusAppConfigRequest {
	s.UserInfo = v
	return s
}

func (s *GetBusAppConfigRequest) Validate() error {
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

type GetBusAppConfigRequestDeviceInfo struct {
	// Value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the application\\"s Skill ID. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME: APK package name, used for Android application customer journeys; SKILL_ID: skill ID, used for cloud-based journeys.
	//
	// example:
	//
	// SKILL_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID (deviceOpenId or deviceUnionId)
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of device ID: OPEN_ID: default device ID; UNION_ID: organization-dimension device ID, available only after an organization has been requested on the Maojing Skill Application Open Platform.
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

func (s GetBusAppConfigRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetBusAppConfigRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetBusAppConfigRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetBusAppConfigRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetBusAppConfigRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetBusAppConfigRequestDeviceInfo) SetEncodeKey(v string) *GetBusAppConfigRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetEncodeType(v string) *GetBusAppConfigRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetId(v string) *GetBusAppConfigRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetIdType(v string) *GetBusAppConfigRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetOrganizationId(v string) *GetBusAppConfigRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetBusAppConfigRequestPayload struct {
	// Plaintext UUID
	//
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
	// Phone number
	//
	// example:
	//
	// 136****1111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s GetBusAppConfigRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequestPayload) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *GetBusAppConfigRequestPayload) GetPhone() *string {
	return s.Phone
}

func (s *GetBusAppConfigRequestPayload) SetOriginUuid(v string) *GetBusAppConfigRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *GetBusAppConfigRequestPayload) SetPhone(v string) *GetBusAppConfigRequestPayload {
	s.Phone = &v
	return s
}

func (s *GetBusAppConfigRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetBusAppConfigRequestUserInfo struct {
	// Value corresponding to the encoding type. If the encoding type is SKILL_ID, the value is the application\\"s Skill ID. If the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Maojing, and each method corresponds to a different encoding type: PACKAGE_NAME: APK package name, used for Android application customer ingest endpoints; SKILL_ID: skill ID, used for cloud-based ingest endpoints.
	//
	// This parameter is required.
	//
	// example:
	//
	// SKILL_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier (UserOpenId or UserUnionId)
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

func (s GetBusAppConfigRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetBusAppConfigRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetBusAppConfigRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetBusAppConfigRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetBusAppConfigRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetBusAppConfigRequestUserInfo) SetEncodeKey(v string) *GetBusAppConfigRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetEncodeType(v string) *GetBusAppConfigRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetId(v string) *GetBusAppConfigRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetIdType(v string) *GetBusAppConfigRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetOrganizationId(v string) *GetBusAppConfigRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
