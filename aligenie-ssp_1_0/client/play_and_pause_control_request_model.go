// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayAndPauseControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *PlayAndPauseControlRequestDeviceInfo) *PlayAndPauseControlRequest
	GetDeviceInfo() *PlayAndPauseControlRequestDeviceInfo
	SetOpenPlayAndPauseControlParam(v *PlayAndPauseControlRequestOpenPlayAndPauseControlParam) *PlayAndPauseControlRequest
	GetOpenPlayAndPauseControlParam() *PlayAndPauseControlRequestOpenPlayAndPauseControlParam
	SetUserInfo(v *PlayAndPauseControlRequestUserInfo) *PlayAndPauseControlRequest
	GetUserInfo() *PlayAndPauseControlRequestUserInfo
}

type PlayAndPauseControlRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *PlayAndPauseControlRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	//
	// This parameter is required.
	OpenPlayAndPauseControlParam *PlayAndPauseControlRequestOpenPlayAndPauseControlParam `json:"OpenPlayAndPauseControlParam,omitempty" xml:"OpenPlayAndPauseControlParam,omitempty" type:"Struct"`
	// User identity information
	//
	// This parameter is required.
	UserInfo *PlayAndPauseControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PlayAndPauseControlRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlRequest) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequest) GetDeviceInfo() *PlayAndPauseControlRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *PlayAndPauseControlRequest) GetOpenPlayAndPauseControlParam() *PlayAndPauseControlRequestOpenPlayAndPauseControlParam {
	return s.OpenPlayAndPauseControlParam
}

func (s *PlayAndPauseControlRequest) GetUserInfo() *PlayAndPauseControlRequestUserInfo {
	return s.UserInfo
}

func (s *PlayAndPauseControlRequest) SetDeviceInfo(v *PlayAndPauseControlRequestDeviceInfo) *PlayAndPauseControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *PlayAndPauseControlRequest) SetOpenPlayAndPauseControlParam(v *PlayAndPauseControlRequestOpenPlayAndPauseControlParam) *PlayAndPauseControlRequest {
	s.OpenPlayAndPauseControlParam = v
	return s
}

func (s *PlayAndPauseControlRequest) SetUserInfo(v *PlayAndPauseControlRequestUserInfo) *PlayAndPauseControlRequest {
	s.UserInfo = v
	return s
}

func (s *PlayAndPauseControlRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OpenPlayAndPauseControlParam != nil {
		if err := s.OpenPlayAndPauseControlParam.Validate(); err != nil {
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

type PlayAndPauseControlRequestDeviceInfo struct {
	// Value corresponding to the encoding type.
	//
	// If the encoding type is `SKILL_ID`, the value is the application\\"s Skill ID. If the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device identity for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application customer link. `SKILL_ID`: Skill ID, used for the cloud-based link.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID, which can be set to deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of Device ID.
	//
	// `OPEN_ID`: The default device ID identity. `UNION_ID`: The organization-dimension device ID identity. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. This parameter is required if IdType is UNION_ID.
	//
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PlayAndPauseControlRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PlayAndPauseControlRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PlayAndPauseControlRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *PlayAndPauseControlRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *PlayAndPauseControlRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetEncodeKey(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetEncodeType(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetId(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetIdType(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) SetOrganizationId(v string) *PlayAndPauseControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *PlayAndPauseControlRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type PlayAndPauseControlRequestOpenPlayAndPauseControlParam struct {
	// Playback: Play; Pause: Pause.
	//
	// This parameter is required.
	//
	// example:
	//
	// Play
	OpenPlayAndPauseCommand *string `json:"OpenPlayAndPauseCommand,omitempty" xml:"OpenPlayAndPauseCommand,omitempty"`
}

func (s PlayAndPauseControlRequestOpenPlayAndPauseControlParam) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlRequestOpenPlayAndPauseControlParam) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequestOpenPlayAndPauseControlParam) GetOpenPlayAndPauseCommand() *string {
	return s.OpenPlayAndPauseCommand
}

func (s *PlayAndPauseControlRequestOpenPlayAndPauseControlParam) SetOpenPlayAndPauseCommand(v string) *PlayAndPauseControlRequestOpenPlayAndPauseControlParam {
	s.OpenPlayAndPauseCommand = &v
	return s
}

func (s *PlayAndPauseControlRequestOpenPlayAndPauseControlParam) Validate() error {
	return dara.Validate(s)
}

type PlayAndPauseControlRequestUserInfo struct {
	// The value corresponding to the encoding type.
	//
	// When the encoding type is `SKILL_ID`, the value is the application\\"s Skill ID. When the encoding type is `PACKAGE_NAME`, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier from Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application customer flow. `SKILL_ID`: Skill ID, used for the cloud-based flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User identifier, set to either userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of User ID.
	//
	// `OPEN_ID`: The default user ID identity. `UNION_ID`: The organization-dimension user ID identity. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required when IdType is `UNION_ID`.
	//
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PlayAndPauseControlRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PlayAndPauseControlRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PlayAndPauseControlRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *PlayAndPauseControlRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *PlayAndPauseControlRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PlayAndPauseControlRequestUserInfo) SetEncodeKey(v string) *PlayAndPauseControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetEncodeType(v string) *PlayAndPauseControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetId(v string) *PlayAndPauseControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetIdType(v string) *PlayAndPauseControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) SetOrganizationId(v string) *PlayAndPauseControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *PlayAndPauseControlRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
