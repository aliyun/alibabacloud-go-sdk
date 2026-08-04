// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *GetCurrentPlayingItemRequestDeviceInfo) *GetCurrentPlayingItemRequest
	GetDeviceInfo() *GetCurrentPlayingItemRequestDeviceInfo
	SetUserInfo(v *GetCurrentPlayingItemRequestUserInfo) *GetCurrentPlayingItemRequest
	GetUserInfo() *GetCurrentPlayingItemRequestUserInfo
}

type GetCurrentPlayingItemRequest struct {
	// Device identification information
	//
	// This parameter is required.
	DeviceInfo *GetCurrentPlayingItemRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// User identification information
	//
	// This parameter is required.
	UserInfo *GetCurrentPlayingItemRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetCurrentPlayingItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemRequest) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequest) GetDeviceInfo() *GetCurrentPlayingItemRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *GetCurrentPlayingItemRequest) GetUserInfo() *GetCurrentPlayingItemRequestUserInfo {
	return s.UserInfo
}

func (s *GetCurrentPlayingItemRequest) SetDeviceInfo(v *GetCurrentPlayingItemRequestDeviceInfo) *GetCurrentPlayingItemRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetCurrentPlayingItemRequest) SetUserInfo(v *GetCurrentPlayingItemRequestUserInfo) *GetCurrentPlayingItemRequest {
	s.UserInfo = v
	return s
}

func (s *GetCurrentPlayingItemRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
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

type GetCurrentPlayingItemRequestDeviceInfo struct {
	// Value corresponding to the encoding type
	//
	// When the encoding type is `SKILL_ID`, the value is the application\\"s Skill ID. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application client link. `SKILL_ID`: Skill ID, used for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device identifier, set to either deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of device ID
	//
	// `OPEN_ID`: The default device ID identifier. `UNION_ID`: Device ID identifier at the organization dimension. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
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
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingItemRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetEncodeKey(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetEncodeType(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetId(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetIdType(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) SetOrganizationId(v string) *GetCurrentPlayingItemRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetCurrentPlayingItemRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type GetCurrentPlayingItemRequestUserInfo struct {
	// The value corresponding to the encoding type.
	//
	// When the encoding type is `SKILL_ID`, the value is the Skill ID of the application. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// The encoding type. There are multiple ways to obtain the User Identifier for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: The APK package name, which is the encoding type for the Android application customer link. `SKILL_ID`: The skill ID, which is the encoding type for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// The User Identifier, which can be set to userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of User ID.
	//
	// `OPEN_ID`: The default User ID identity. `UNION_ID`: The organization-dimension User ID identity. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// The organization ID. This field is required if IdType is set to UNION_ID.
	//
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetCurrentPlayingItemRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *GetCurrentPlayingItemRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetEncodeKey(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetEncodeType(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetId(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetIdType(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) SetOrganizationId(v string) *GetCurrentPlayingItemRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *GetCurrentPlayingItemRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
