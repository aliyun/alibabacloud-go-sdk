// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAndRemoveFavoriteContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *AddAndRemoveFavoriteContentRequestDeviceInfo) *AddAndRemoveFavoriteContentRequest
	GetDeviceInfo() *AddAndRemoveFavoriteContentRequestDeviceInfo
	SetOpenAddAndRemoveFavoriteContentRequest(v *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) *AddAndRemoveFavoriteContentRequest
	GetOpenAddAndRemoveFavoriteContentRequest() *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest
	SetUserInfo(v *AddAndRemoveFavoriteContentRequestUserInfo) *AddAndRemoveFavoriteContentRequest
	GetUserInfo() *AddAndRemoveFavoriteContentRequestUserInfo
}

type AddAndRemoveFavoriteContentRequest struct {
	// Device identification information
	//
	// This parameter is required.
	DeviceInfo *AddAndRemoveFavoriteContentRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	//
	// This parameter is required.
	OpenAddAndRemoveFavoriteContentRequest *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest `json:"OpenAddAndRemoveFavoriteContentRequest,omitempty" xml:"OpenAddAndRemoveFavoriteContentRequest,omitempty" type:"Struct"`
	// User identification information
	//
	// This parameter is required.
	UserInfo *AddAndRemoveFavoriteContentRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s AddAndRemoveFavoriteContentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequest) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequest) GetDeviceInfo() *AddAndRemoveFavoriteContentRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *AddAndRemoveFavoriteContentRequest) GetOpenAddAndRemoveFavoriteContentRequest() *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest {
	return s.OpenAddAndRemoveFavoriteContentRequest
}

func (s *AddAndRemoveFavoriteContentRequest) GetUserInfo() *AddAndRemoveFavoriteContentRequestUserInfo {
	return s.UserInfo
}

func (s *AddAndRemoveFavoriteContentRequest) SetDeviceInfo(v *AddAndRemoveFavoriteContentRequestDeviceInfo) *AddAndRemoveFavoriteContentRequest {
	s.DeviceInfo = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequest) SetOpenAddAndRemoveFavoriteContentRequest(v *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) *AddAndRemoveFavoriteContentRequest {
	s.OpenAddAndRemoveFavoriteContentRequest = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequest) SetUserInfo(v *AddAndRemoveFavoriteContentRequestUserInfo) *AddAndRemoveFavoriteContentRequest {
	s.UserInfo = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OpenAddAndRemoveFavoriteContentRequest != nil {
		if err := s.OpenAddAndRemoveFavoriteContentRequest.Validate(); err != nil {
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

type AddAndRemoveFavoriteContentRequestDeviceInfo struct {
	// Value corresponding to the encoding type
	//
	// When the encoding type is `SKILL_ID`, the value is the Skill ID of the application. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application customer link. `SKILL_ID`: Skill ID, used for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID, set to deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of device ID
	//
	// `OPEN_ID`: Default device ID identifier. `UNION_ID`: Device ID identifier at the organization dimension. This value is available only after an organization has been registered on the Tmall Genie Skill Application Open Platform.
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
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AddAndRemoveFavoriteContentRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetEncodeKey(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetEncodeType(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetId(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetIdType(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) SetOrganizationId(v string) *AddAndRemoveFavoriteContentRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest struct {
	// Operation Type
	//
	// ADD for collect; REMOVE for remove from favorites
	//
	// This parameter is required.
	//
	// example:
	//
	// ADD
	FavoriteCmd *string `json:"FavoriteCmd,omitempty" xml:"FavoriteCmd,omitempty"`
	// Object to collect or remove from favorites
	//
	// This parameter is required.
	OpenSourceRawIdPair *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair `json:"OpenSourceRawIdPair,omitempty" xml:"OpenSourceRawIdPair,omitempty" type:"Struct"`
	// Content type
	//
	// Content: CONTENT; Album: ALBUM; Playlist: COLLECT.
	//
	// This parameter is required.
	//
	// example:
	//
	// CONTENT
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) GetFavoriteCmd() *string {
	return s.FavoriteCmd
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) GetOpenSourceRawIdPair() *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair {
	return s.OpenSourceRawIdPair
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) SetFavoriteCmd(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest {
	s.FavoriteCmd = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) SetOpenSourceRawIdPair(v *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest {
	s.OpenSourceRawIdPair = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) SetPackageType(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest {
	s.PackageType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequest) Validate() error {
	if s.OpenSourceRawIdPair != nil {
		if err := s.OpenSourceRawIdPair.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair struct {
	// Extension information
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// Third-party ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2105818057
	RawId *string `json:"RawId,omitempty" xml:"RawId,omitempty"`
	// Source
	//
	// This parameter is required.
	//
	// example:
	//
	// xiami
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) GetRawId() *string {
	return s.RawId
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) GetSource() *string {
	return s.Source
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) SetExtendInfo(v map[string]interface{}) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair {
	s.ExtendInfo = v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) SetRawId(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair {
	s.RawId = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) SetSource(v string) *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair {
	s.Source = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestOpenAddAndRemoveFavoriteContentRequestOpenSourceRawIdPair) Validate() error {
	return dara.Validate(s)
}

type AddAndRemoveFavoriteContentRequestUserInfo struct {
	// The value corresponding to the encoding type.
	//
	// When the encoding type is `SKILL_ID`, the value is the Skill ID of the application. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application customer link. `SKILL_ID`: Skill ID, used for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier, set to userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of User ID.
	//
	// `OPEN_ID`: The default User ID identity. `UNION_ID`: The organization-dimension User ID identity. This value is available only after an organization has been requested on the Tmall Genie Skills Open Platform.
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

func (s AddAndRemoveFavoriteContentRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentRequestUserInfo) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetEncodeKey(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetEncodeType(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetId(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.Id = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetIdType(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) SetOrganizationId(v string) *AddAndRemoveFavoriteContentRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *AddAndRemoveFavoriteContentRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
