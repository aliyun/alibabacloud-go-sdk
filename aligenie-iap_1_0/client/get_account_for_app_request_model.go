// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountForAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetAccountForAppRequestDeviceInfo) *GetAccountForAppRequest
	GetDeviceInfo() *GetAccountForAppRequestDeviceInfo
	SetPayload(v *GetAccountForAppRequestPayload) *GetAccountForAppRequest
	GetPayload() *GetAccountForAppRequestPayload
	SetUserInfo(v *GetAccountForAppRequestUserInfo) *GetAccountForAppRequest
	GetUserInfo() *GetAccountForAppRequestUserInfo
}

type GetAccountForAppRequest struct {
	// Device ID information
	DeviceInfo *GetAccountForAppRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	Payload *GetAccountForAppRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *GetAccountForAppRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetAccountForAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequest) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequest) GetDeviceInfo() *GetAccountForAppRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetAccountForAppRequest) GetPayload() *GetAccountForAppRequestPayload {
	return s.Payload
}

func (s *GetAccountForAppRequest) GetUserInfo() *GetAccountForAppRequestUserInfo {
	return s.UserInfo
}

func (s *GetAccountForAppRequest) SetDeviceInfo(v *GetAccountForAppRequestDeviceInfo) *GetAccountForAppRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetAccountForAppRequest) SetPayload(v *GetAccountForAppRequestPayload) *GetAccountForAppRequest {
	s.Payload = v
	return s
}

func (s *GetAccountForAppRequest) SetUserInfo(v *GetAccountForAppRequestUserInfo) *GetAccountForAppRequest {
	s.UserInfo = v
	return s
}

func (s *GetAccountForAppRequest) Validate() error {
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

type GetAccountForAppRequestDeviceInfo struct {
	// Value corresponding to the encoding type:
	//
	// - When the encoding type is SKILL_ID, the value is the application\\"s Skill ID.
	//
	// - When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
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
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID, set to deviceOpenId or deviceUnionId.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the device ID:
	//
	// - OPEN_ID: Default device ID identifier
	//
	// - UNION_ID: Device ID identifier at the organization dimension. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
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

func (s GetAccountForAppRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAccountForAppRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAccountForAppRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetAccountForAppRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAccountForAppRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAccountForAppRequestDeviceInfo) SetEncodeKey(v string) *GetAccountForAppRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetEncodeType(v string) *GetAccountForAppRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetId(v string) *GetAccountForAppRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetIdType(v string) *GetAccountForAppRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetOrganizationId(v string) *GetAccountForAppRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetAccountForAppRequestPayload struct {
	// Phone number of the authorized membership user.
	//
	// example:
	//
	// 188*777
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// Plaintext UUID
	//
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s GetAccountForAppRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequestPayload) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestPayload) GetPhone() *string {
	return s.Phone
}

func (s *GetAccountForAppRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *GetAccountForAppRequestPayload) SetPhone(v string) *GetAccountForAppRequestPayload {
	s.Phone = &v
	return s
}

func (s *GetAccountForAppRequestPayload) SetOriginUuid(v string) *GetAccountForAppRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *GetAccountForAppRequestPayload) Validate() error {
	return dara.Validate(s)
}

type GetAccountForAppRequestUserInfo struct {
	// Value corresponding to the encoding type:
	//
	// - When the encoding type is SKILL_ID, the value is the application\\"s Skill ID.
	//
	// - When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Tmall Genie, and each way corresponds to a different encoding type:
	//
	// - PACKAGE_NAME: APK package name, used for the Android application customer link.
	//
	// - SKILL_ID: Skill ID, used for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User identifier, which can be set to userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID:
	//
	// - OPEN_ID: The default user ID identity.
	//
	// - UNION_ID: Organization-dimension user ID identity. This value is available only after an organization has been requested on the Tmall Genie Skill Application Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This field is required if IdType is set to UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetAccountForAppRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetAccountForAppRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetAccountForAppRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetAccountForAppRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetAccountForAppRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetAccountForAppRequestUserInfo) SetEncodeKey(v string) *GetAccountForAppRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetEncodeType(v string) *GetAccountForAppRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetId(v string) *GetAccountForAppRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetIdType(v string) *GetAccountForAppRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetOrganizationId(v string) *GetAccountForAppRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
