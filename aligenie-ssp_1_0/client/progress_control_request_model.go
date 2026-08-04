// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProgressControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *ProgressControlRequestDeviceInfo) *ProgressControlRequest
	GetDeviceInfo() *ProgressControlRequestDeviceInfo
	SetOpenProgressControlRequest(v *ProgressControlRequestOpenProgressControlRequest) *ProgressControlRequest
	GetOpenProgressControlRequest() *ProgressControlRequestOpenProgressControlRequest
	SetUserInfo(v *ProgressControlRequestUserInfo) *ProgressControlRequest
	GetUserInfo() *ProgressControlRequestUserInfo
}

type ProgressControlRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *ProgressControlRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	//
	// This parameter is required.
	OpenProgressControlRequest *ProgressControlRequestOpenProgressControlRequest `json:"OpenProgressControlRequest,omitempty" xml:"OpenProgressControlRequest,omitempty" type:"Struct"`
	// User identity information
	//
	// This parameter is required.
	UserInfo *ProgressControlRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ProgressControlRequest) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlRequest) GoString() string {
	return s.String()
}

func (s *ProgressControlRequest) GetDeviceInfo() *ProgressControlRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *ProgressControlRequest) GetOpenProgressControlRequest() *ProgressControlRequestOpenProgressControlRequest {
	return s.OpenProgressControlRequest
}

func (s *ProgressControlRequest) GetUserInfo() *ProgressControlRequestUserInfo {
	return s.UserInfo
}

func (s *ProgressControlRequest) SetDeviceInfo(v *ProgressControlRequestDeviceInfo) *ProgressControlRequest {
	s.DeviceInfo = v
	return s
}

func (s *ProgressControlRequest) SetOpenProgressControlRequest(v *ProgressControlRequestOpenProgressControlRequest) *ProgressControlRequest {
	s.OpenProgressControlRequest = v
	return s
}

func (s *ProgressControlRequest) SetUserInfo(v *ProgressControlRequestUserInfo) *ProgressControlRequest {
	s.UserInfo = v
	return s
}

func (s *ProgressControlRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OpenProgressControlRequest != nil {
		if err := s.OpenProgressControlRequest.Validate(); err != nil {
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

type ProgressControlRequestDeviceInfo struct {
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
	// Encoding type. There are multiple ways to obtain the device ID for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application customer flow. `SKILL_ID`: Skill ID, used for the cloud-based flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID. Set to either deviceOpenId or deviceUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of Device ID.
	//
	// `OPEN_ID`: The default device identity. `UNION_ID`: The organization-dimension device identity, which is available only after an organization has been requested on the Tmall Genie Skills Application Open Platform.
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

func (s ProgressControlRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ProgressControlRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ProgressControlRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ProgressControlRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *ProgressControlRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *ProgressControlRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ProgressControlRequestDeviceInfo) SetEncodeKey(v string) *ProgressControlRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetEncodeType(v string) *ProgressControlRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetId(v string) *ProgressControlRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetIdType(v string) *ProgressControlRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) SetOrganizationId(v string) *ProgressControlRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *ProgressControlRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type ProgressControlRequestOpenProgressControlRequest struct {
	// Extension information
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// Song progress, in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s ProgressControlRequestOpenProgressControlRequest) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlRequestOpenProgressControlRequest) GoString() string {
	return s.String()
}

func (s *ProgressControlRequestOpenProgressControlRequest) GetExtendInfo() map[string]interface{} {
	return s.ExtendInfo
}

func (s *ProgressControlRequestOpenProgressControlRequest) GetProgress() *int32 {
	return s.Progress
}

func (s *ProgressControlRequestOpenProgressControlRequest) SetExtendInfo(v map[string]interface{}) *ProgressControlRequestOpenProgressControlRequest {
	s.ExtendInfo = v
	return s
}

func (s *ProgressControlRequestOpenProgressControlRequest) SetProgress(v int32) *ProgressControlRequestOpenProgressControlRequest {
	s.Progress = &v
	return s
}

func (s *ProgressControlRequestOpenProgressControlRequest) Validate() error {
	return dara.Validate(s)
}

type ProgressControlRequestUserInfo struct {
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
	// Encoding type. There are multiple ways to obtain the Tmall Genie user identity, and each way corresponds to a different encoding type.
	//
	// `PACKAGE_NAME`: APK package name, used for the Android application client path. `SKILL_ID`: Skill ID, used for the cloud-based path.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User identifier, set to userOpenId or userUnionId.
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ouRjKOX95tn1a0kwb2+Ao6Q1CAxASJUZDWtlk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of User ID.
	//
	// `OPEN_ID`: The default User ID identity. `UNION_ID`: The organization-dimension User ID identity, which is available only after an organization has been requested on the Tmall Genie Skills Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is `UNION_ID`.
	//
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ProgressControlRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ProgressControlRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ProgressControlRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ProgressControlRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ProgressControlRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ProgressControlRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ProgressControlRequestUserInfo) SetEncodeKey(v string) *ProgressControlRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetEncodeType(v string) *ProgressControlRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetId(v string) *ProgressControlRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetIdType(v string) *ProgressControlRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ProgressControlRequestUserInfo) SetOrganizationId(v string) *ProgressControlRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ProgressControlRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
