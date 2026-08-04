// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayModeControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *PlayModeControlRequestDeviceInfo) *PlayModeControlRequest
	GetDeviceInfo() *PlayModeControlRequestDeviceInfo
	SetOpenPlayModeControlRequest(v *PlayModeControlRequestOpenPlayModeControlRequest) *PlayModeControlRequest
	GetOpenPlayModeControlRequest() *PlayModeControlRequestOpenPlayModeControlRequest
	SetUserInfo(v *PlayModeControlRequestUserInfo) *PlayModeControlRequest
	GetUserInfo() *PlayModeControlRequestUserInfo
}

type PlayModeControlRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *PlayModeControlRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	//
	// This parameter is required.
	OpenPlayModeControlRequest *PlayModeControlRequestOpenPlayModeControlRequest `json:"OpenPlayModeControlRequest,omitempty" xml:"OpenPlayModeControlRequest,omitempty" type:"Struct"`
	// User Identifier information
	//
	// This parameter is required.
	UserInfo *PlayModeControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PlayModeControlRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequest) GetDeviceInfo() *PlayModeControlRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *PlayModeControlRequest) GetOpenPlayModeControlRequest() *PlayModeControlRequestOpenPlayModeControlRequest {
	return s.OpenPlayModeControlRequest
}

func (s *PlayModeControlRequest) GetUserInfo() *PlayModeControlRequestUserInfo {
	return s.UserInfo
}

func (s *PlayModeControlRequest) SetDeviceInfo(v *PlayModeControlRequestDeviceInfo) *PlayModeControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *PlayModeControlRequest) SetOpenPlayModeControlRequest(v *PlayModeControlRequestOpenPlayModeControlRequest) *PlayModeControlRequest {
	s.OpenPlayModeControlRequest = v
	return s
}

func (s *PlayModeControlRequest) SetUserInfo(v *PlayModeControlRequestUserInfo) *PlayModeControlRequest {
	s.UserInfo = v
	return s
}

func (s *PlayModeControlRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OpenPlayModeControlRequest != nil {
		if err := s.OpenPlayModeControlRequest.Validate(); err != nil {
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

type PlayModeControlRequestDeviceInfo struct {
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
	// `PACKAGE_NAME`: APK package name, used for Android application customer journeys. `SKILL_ID`: Skill ID, used for cloud-based journeys.
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
	// `OPEN_ID`: Default device ID identity. `UNION_ID`: Organization-dimension device ID identity, available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
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

func (s PlayModeControlRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PlayModeControlRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PlayModeControlRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *PlayModeControlRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *PlayModeControlRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PlayModeControlRequestDeviceInfo) SetEncodeKey(v string) *PlayModeControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetEncodeType(v string) *PlayModeControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetId(v string) *PlayModeControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetIdType(v string) *PlayModeControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) SetOrganizationId(v string) *PlayModeControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *PlayModeControlRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type PlayModeControlRequestOpenPlayModeControlRequest struct {
	// Playback mode
	//
	// List loop: Repeat; Shuffle: Shuffle; Single track loop: RepeatOne; NAT mode: Normal;
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	OpenPlayMode *string `json:"OpenPlayMode,omitempty" xml:"OpenPlayMode,omitempty"`
}

func (s PlayModeControlRequestOpenPlayModeControlRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequestOpenPlayModeControlRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestOpenPlayModeControlRequest) GetOpenPlayMode() *string {
	return s.OpenPlayMode
}

func (s *PlayModeControlRequestOpenPlayModeControlRequest) SetOpenPlayMode(v string) *PlayModeControlRequestOpenPlayModeControlRequest {
	s.OpenPlayMode = &v
	return s
}

func (s *PlayModeControlRequestOpenPlayModeControlRequest) Validate() error {
	return dara.Validate(s)
}

type PlayModeControlRequestUserInfo struct {
	// Value corresponding to the encoding type.
	//
	// When the encoding type is `SKILL_ID`, the value is the application\\"s Skill ID. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the User Identifier for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application customer ingest endpoint. `SKILL_ID`: Skill ID, used for the cloud-side ingest endpoint.
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
	// `OPEN_ID`: The default User ID identity. `UNION_ID`: Organization-dimension User ID identity. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
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

func (s PlayModeControlRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PlayModeControlRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PlayModeControlRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PlayModeControlRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *PlayModeControlRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *PlayModeControlRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PlayModeControlRequestUserInfo) SetEncodeKey(v string) *PlayModeControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetEncodeType(v string) *PlayModeControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetId(v string) *PlayModeControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetIdType(v string) *PlayModeControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) SetOrganizationId(v string) *PlayModeControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *PlayModeControlRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
