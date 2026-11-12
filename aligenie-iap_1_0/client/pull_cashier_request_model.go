// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullCashierRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *PullCashierRequestDeviceInfo) *PullCashierRequest
	GetDeviceInfo() *PullCashierRequestDeviceInfo
	SetPayload(v *PullCashierRequestPayload) *PullCashierRequest
	GetPayload() *PullCashierRequestPayload
	SetUserInfo(v *PullCashierRequestUserInfo) *PullCashierRequest
	GetUserInfo() *PullCashierRequestUserInfo
}

type PullCashierRequest struct {
	// device ID information
	DeviceInfo *PullCashierRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	Payload *PullCashierRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *PullCashierRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PullCashierRequest) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequest) GoString() string {
	return s.String()
}

func (s *PullCashierRequest) GetDeviceInfo() *PullCashierRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *PullCashierRequest) GetPayload() *PullCashierRequestPayload {
	return s.Payload
}

func (s *PullCashierRequest) GetUserInfo() *PullCashierRequestUserInfo {
	return s.UserInfo
}

func (s *PullCashierRequest) SetDeviceInfo(v *PullCashierRequestDeviceInfo) *PullCashierRequest {
	s.DeviceInfo = v
	return s
}

func (s *PullCashierRequest) SetPayload(v *PullCashierRequestPayload) *PullCashierRequest {
	s.Payload = v
	return s
}

func (s *PullCashierRequest) SetUserInfo(v *PullCashierRequestUserInfo) *PullCashierRequest {
	s.UserInfo = v
	return s
}

func (s *PullCashierRequest) Validate() error {
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

type PullCashierRequestDeviceInfo struct {
	// The value corresponding to the encoding type.
	//
	// - When the encoding type is SKILL_ID, the value is the Skill ID of the application.
	//
	// - When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID of Tmall Genie, and each method corresponds to a different encoding type.
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
	// Type of device ID
	//
	// - OPEN_ID: The default device ID.
	//
	// - UNION_ID: Organization-level device ID. This value is available only after an organization has been requested on the Tmall Genie skill application Open Platform.
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

func (s PullCashierRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PullCashierRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PullCashierRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PullCashierRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *PullCashierRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *PullCashierRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PullCashierRequestDeviceInfo) SetEncodeKey(v string) *PullCashierRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetEncodeType(v string) *PullCashierRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetId(v string) *PullCashierRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetIdType(v string) *PullCashierRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetOrganizationId(v string) *PullCashierRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type PullCashierRequestPayload struct {
	// Plaintext UUID
	//
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s PullCashierRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequestPayload) GoString() string {
	return s.String()
}

func (s *PullCashierRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *PullCashierRequestPayload) SetOriginUuid(v string) *PullCashierRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *PullCashierRequestPayload) Validate() error {
	return dara.Validate(s)
}

type PullCashierRequestUserInfo struct {
	// The value corresponding to the encoding type.
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
	// Encoding type. There are multiple ways to obtain the user identifier from Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used for the Android application customer ingest endpoint.
	//
	// - SKILL_ID: Skill ID, used for the cloud-based ingest endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User identifier, set to either userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the User ID.
	//
	// - OPEN_ID: The default User ID identity.
	//
	// - UNION_ID: Organization-dimension User ID identity. This value is available only after an organization has been requested on the Tmall Genie Skill Application Platform.
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

func (s PullCashierRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s PullCashierRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PullCashierRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PullCashierRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PullCashierRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *PullCashierRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *PullCashierRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PullCashierRequestUserInfo) SetEncodeKey(v string) *PullCashierRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetEncodeType(v string) *PullCashierRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetId(v string) *PullCashierRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetIdType(v string) *PullCashierRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetOrganizationId(v string) *PullCashierRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *PullCashierRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
