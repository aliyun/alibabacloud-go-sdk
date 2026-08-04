// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexControlPlayingListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *IndexControlPlayingListRequestDeviceInfo) *IndexControlPlayingListRequest
	GetDeviceInfo() *IndexControlPlayingListRequestDeviceInfo
	SetOpenIndexControlRequest(v *IndexControlPlayingListRequestOpenIndexControlRequest) *IndexControlPlayingListRequest
	GetOpenIndexControlRequest() *IndexControlPlayingListRequestOpenIndexControlRequest
	SetUserInfo(v *IndexControlPlayingListRequestUserInfo) *IndexControlPlayingListRequest
	GetUserInfo() *IndexControlPlayingListRequestUserInfo
}

type IndexControlPlayingListRequest struct {
	// This parameter is required.
	DeviceInfo *IndexControlPlayingListRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	//
	// This parameter is required.
	OpenIndexControlRequest *IndexControlPlayingListRequestOpenIndexControlRequest `json:"OpenIndexControlRequest,omitempty" xml:"OpenIndexControlRequest,omitempty" type:"Struct"`
	// User Identifier information
	//
	// This parameter is required.
	UserInfo *IndexControlPlayingListRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s IndexControlPlayingListRequest) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequest) GetDeviceInfo() *IndexControlPlayingListRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *IndexControlPlayingListRequest) GetOpenIndexControlRequest() *IndexControlPlayingListRequestOpenIndexControlRequest {
	return s.OpenIndexControlRequest
}

func (s *IndexControlPlayingListRequest) GetUserInfo() *IndexControlPlayingListRequestUserInfo {
	return s.UserInfo
}

func (s *IndexControlPlayingListRequest) SetDeviceInfo(v *IndexControlPlayingListRequestDeviceInfo) *IndexControlPlayingListRequest {
	s.DeviceInfo = v
	return s
}

func (s *IndexControlPlayingListRequest) SetOpenIndexControlRequest(v *IndexControlPlayingListRequestOpenIndexControlRequest) *IndexControlPlayingListRequest {
	s.OpenIndexControlRequest = v
	return s
}

func (s *IndexControlPlayingListRequest) SetUserInfo(v *IndexControlPlayingListRequestUserInfo) *IndexControlPlayingListRequest {
	s.UserInfo = v
	return s
}

func (s *IndexControlPlayingListRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OpenIndexControlRequest != nil {
		if err := s.OpenIndexControlRequest.Validate(); err != nil {
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

type IndexControlPlayingListRequestDeviceInfo struct {
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
	// `PACKAGE_NAME`: APK package name, used for the Android application customer link. `SKILL_ID`: Skill ID, used for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID, set to either deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of device ID
	//
	// `OPEN_ID`: The default device ID identity. `UNION_ID`: Device ID identity at the organization dimension. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
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

func (s IndexControlPlayingListRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *IndexControlPlayingListRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetEncodeKey(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetEncodeType(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetId(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetIdType(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) SetOrganizationId(v string) *IndexControlPlayingListRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *IndexControlPlayingListRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type IndexControlPlayingListRequestOpenIndexControlRequest struct {
	// Extension information
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The index to be played back
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Whether content playback should continue. Default is false.
	//
	// example:
	//
	// false
	NeedContentContinued *bool `json:"NeedContentContinued,omitempty" xml:"NeedContentContinued,omitempty"`
}

func (s IndexControlPlayingListRequestOpenIndexControlRequest) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequestOpenIndexControlRequest) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) GetIndex() *int32 {
	return s.Index
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) GetNeedContentContinued() *bool {
	return s.NeedContentContinued
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetExtendInfo(v map[string]interface{}) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.ExtendInfo = v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetIndex(v int32) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.Index = &v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) SetNeedContentContinued(v bool) *IndexControlPlayingListRequestOpenIndexControlRequest {
	s.NeedContentContinued = &v
	return s
}

func (s *IndexControlPlayingListRequestOpenIndexControlRequest) Validate() error {
	return dara.Validate(s)
}

type IndexControlPlayingListRequestUserInfo struct {
	// Value corresponding to the encoding type.
	//
	// When the encoding type is `SKILL_ID`, the value is the Skill ID of the application. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the User Identifier for Tmall Genie, and each method corresponds to a different encoding type.
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
	// Type of User ID
	//
	// `OPEN_ID`: The default User ID identity. `UNION_ID`: The organization-dimension User ID identity, which is available only after an organization has been requested on the Tmall Genie Skills Application Open Platform.
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

func (s IndexControlPlayingListRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListRequestUserInfo) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *IndexControlPlayingListRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *IndexControlPlayingListRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *IndexControlPlayingListRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *IndexControlPlayingListRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *IndexControlPlayingListRequestUserInfo) SetEncodeKey(v string) *IndexControlPlayingListRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetEncodeType(v string) *IndexControlPlayingListRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetId(v string) *IndexControlPlayingListRequestUserInfo {
	s.Id = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetIdType(v string) *IndexControlPlayingListRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) SetOrganizationId(v string) *IndexControlPlayingListRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *IndexControlPlayingListRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
