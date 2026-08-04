// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ListRecommendContentRequestDeviceInfo) *ListRecommendContentRequest
	GetDeviceInfo() *ListRecommendContentRequestDeviceInfo
	SetRequest(v *ListRecommendContentRequestRequest) *ListRecommendContentRequest
	GetRequest() *ListRecommendContentRequestRequest
	SetUserInfo(v *ListRecommendContentRequestUserInfo) *ListRecommendContentRequest
	GetUserInfo() *ListRecommendContentRequestUserInfo
}

type ListRecommendContentRequest struct {
	// Device identification information
	//
	// This parameter is required.
	DeviceInfo *ListRecommendContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Request Parameters
	//
	// This parameter is required.
	Request *ListRecommendContentRequestRequest `json:"Request,omitempty" xml:"Request,omitempty" type:"Struct"`
	// User identification information
	//
	// This parameter is required.
	UserInfo *ListRecommendContentRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListRecommendContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequest) GetDeviceInfo() *ListRecommendContentRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ListRecommendContentRequest) GetRequest() *ListRecommendContentRequestRequest {
	return s.Request
}

func (s *ListRecommendContentRequest) GetUserInfo() *ListRecommendContentRequestUserInfo {
	return s.UserInfo
}

func (s *ListRecommendContentRequest) SetDeviceInfo(v *ListRecommendContentRequestDeviceInfo) *ListRecommendContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListRecommendContentRequest) SetRequest(v *ListRecommendContentRequestRequest) *ListRecommendContentRequest {
	s.Request = v
	return s
}

func (s *ListRecommendContentRequest) SetUserInfo(v *ListRecommendContentRequestUserInfo) *ListRecommendContentRequest {
	s.UserInfo = v
	return s
}

func (s *ListRecommendContentRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Request != nil {
		if err := s.Request.Validate(); err != nil {
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

type ListRecommendContentRequestDeviceInfo struct {
	// Value corresponding to the encoding type
	//
	// When the encoding type is SKILL_ID, the value is the application\\"s Skill ID.
	//
	// When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// PACKAGE_NAME: APK package name, used for the Android application client path.
	//
	// SKILL_ID: Skill ID, used for the cloud-based path.
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
	// OPEN_ID: Default device ID identity.
	//
	// UNION_ID: Organization-dimension device ID identity. This value is available only after an organization has been registered on the Tmall Genie Skill Application Open Platform.
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

func (s ListRecommendContentRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRecommendContentRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRecommendContentRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ListRecommendContentRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRecommendContentRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRecommendContentRequestDeviceInfo) SetEncodeKey(v string) *ListRecommendContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetEncodeType(v string) *ListRecommendContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetId(v string) *ListRecommendContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetIdType(v string) *ListRecommendContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) SetOrganizationId(v string) *ListRecommendContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRecommendContentRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentRequestRequest struct {
	// Quantity of recommendations
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Default value: song (currently, the extension field supports only song)
	//
	// example:
	//
	// song
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecommendContentRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequestRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestRequest) GetCount() *int32 {
	return s.Count
}

func (s *ListRecommendContentRequestRequest) GetType() *string {
	return s.Type
}

func (s *ListRecommendContentRequestRequest) SetCount(v int32) *ListRecommendContentRequestRequest {
	s.Count = &v
	return s
}

func (s *ListRecommendContentRequestRequest) SetType(v string) *ListRecommendContentRequestRequest {
	s.Type = &v
	return s
}

func (s *ListRecommendContentRequestRequest) Validate() error {
	return dara.Validate(s)
}

type ListRecommendContentRequestUserInfo struct {
	// Value corresponding to the encoding type.
	//
	// When the encoding type is SKILL_ID, the value is the Skill ID of the application.
	//
	// When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// PACKAGE_NAME: APK package name, used for the Android application client path.
	//
	// SKILL_ID: Skill ID, used for the cloud-based path.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier, set to userOpenId or userUnionId.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of User ID.
	//
	// OPEN_ID: The default User ID identifier.
	//
	// UNION_ID: The organization-dimension User ID identifier. This value is available only after an organization has been requested on the Tmall Genie Skills Application Open Platform.
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

func (s ListRecommendContentRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListRecommendContentRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ListRecommendContentRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ListRecommendContentRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ListRecommendContentRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ListRecommendContentRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRecommendContentRequestUserInfo) SetEncodeKey(v string) *ListRecommendContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetEncodeType(v string) *ListRecommendContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetId(v string) *ListRecommendContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetIdType(v string) *ListRecommendContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) SetOrganizationId(v string) *ListRecommendContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ListRecommendContentRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
