// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppUseTimeReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *AppUseTimeReportRequestDeviceInfo) *AppUseTimeReportRequest
	GetDeviceInfo() *AppUseTimeReportRequestDeviceInfo
	SetPayload(v *AppUseTimeReportRequestPayload) *AppUseTimeReportRequest
	GetPayload() *AppUseTimeReportRequestPayload
	SetUserInfo(v *AppUseTimeReportRequestUserInfo) *AppUseTimeReportRequest
	GetUserInfo() *AppUseTimeReportRequestUserInfo
}

type AppUseTimeReportRequest struct {
	// Device identity information
	DeviceInfo *AppUseTimeReportRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Business parameters
	Payload *AppUseTimeReportRequestPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// User identifier information
	//
	// This parameter is required.
	UserInfo *AppUseTimeReportRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s AppUseTimeReportRequest) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportRequest) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequest) GetDeviceInfo() *AppUseTimeReportRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *AppUseTimeReportRequest) GetPayload() *AppUseTimeReportRequestPayload {
	return s.Payload
}

func (s *AppUseTimeReportRequest) GetUserInfo() *AppUseTimeReportRequestUserInfo {
	return s.UserInfo
}

func (s *AppUseTimeReportRequest) SetDeviceInfo(v *AppUseTimeReportRequestDeviceInfo) *AppUseTimeReportRequest {
	s.DeviceInfo = v
	return s
}

func (s *AppUseTimeReportRequest) SetPayload(v *AppUseTimeReportRequestPayload) *AppUseTimeReportRequest {
	s.Payload = v
	return s
}

func (s *AppUseTimeReportRequest) SetUserInfo(v *AppUseTimeReportRequestUserInfo) *AppUseTimeReportRequest {
	s.UserInfo = v
	return s
}

func (s *AppUseTimeReportRequest) Validate() error {
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

type AppUseTimeReportRequestDeviceInfo struct {
	// The value corresponding to the encoding type.
	//
	// - When the encoding type is SKILL_ID, the value is the application\\"s Skill ID.
	//
	// - When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client application.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device identity for Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used for Android application customer-side integration.
	//
	// - SKILL_ID: Skill ID, used for cloud-side integration.
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
	// - OPEN_ID: Default device ID.
	//
	// - UNION_ID: Organization-dimension device ID, available only after an organization has been requested in the Tmall Genie Skill Application Open Platform.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID, required if IdType is UNION_ID.
	//
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AppUseTimeReportRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *AppUseTimeReportRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *AppUseTimeReportRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *AppUseTimeReportRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *AppUseTimeReportRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AppUseTimeReportRequestDeviceInfo) SetEncodeKey(v string) *AppUseTimeReportRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetEncodeType(v string) *AppUseTimeReportRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetId(v string) *AppUseTimeReportRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetIdType(v string) *AppUseTimeReportRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetOrganizationId(v string) *AppUseTimeReportRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type AppUseTimeReportRequestPayload struct {
	// Action item
	//
	// - start indicates Start
	//
	// - end indicates End
	//
	// This parameter is required.
	//
	// example:
	//
	// start
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Indicates whether the benefit is exclusive to members:
	//
	// - 1 indicates Yes
	//
	// - 0 indicates No
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsPrivilege *int32 `json:"IsPrivilege,omitempty" xml:"IsPrivilege,omitempty"`
	// Resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource type:
	//
	// - 0 indicates game
	//
	// - 1 indicates audio
	//
	// - 2 indicates video
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Step code. The start and end actions must be uniquely matched.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	StepCode *string `json:"StepCode,omitempty" xml:"StepCode,omitempty"`
	// Membership type:
	//
	// - 0 indicates Cat Spirit Membership
	//
	// - 1 indicates third-party proprietary membership
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	VipType *int32 `json:"VipType,omitempty" xml:"VipType,omitempty"`
	// Plaintext UUID
	//
	// example:
	//
	// 731D5F********DC3B
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s AppUseTimeReportRequestPayload) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportRequestPayload) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequestPayload) GetAction() *string {
	return s.Action
}

func (s *AppUseTimeReportRequestPayload) GetIsPrivilege() *int32 {
	return s.IsPrivilege
}

func (s *AppUseTimeReportRequestPayload) GetResourceId() *string {
	return s.ResourceId
}

func (s *AppUseTimeReportRequestPayload) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *AppUseTimeReportRequestPayload) GetStepCode() *string {
	return s.StepCode
}

func (s *AppUseTimeReportRequestPayload) GetVipType() *int32 {
	return s.VipType
}

func (s *AppUseTimeReportRequestPayload) GetOriginUuid() *string {
	return s.OriginUuid
}

func (s *AppUseTimeReportRequestPayload) SetAction(v string) *AppUseTimeReportRequestPayload {
	s.Action = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetIsPrivilege(v int32) *AppUseTimeReportRequestPayload {
	s.IsPrivilege = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetResourceId(v string) *AppUseTimeReportRequestPayload {
	s.ResourceId = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetResourceType(v int32) *AppUseTimeReportRequestPayload {
	s.ResourceType = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetStepCode(v string) *AppUseTimeReportRequestPayload {
	s.StepCode = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetVipType(v int32) *AppUseTimeReportRequestPayload {
	s.VipType = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetOriginUuid(v string) *AppUseTimeReportRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) Validate() error {
	return dara.Validate(s)
}

type AppUseTimeReportRequestUserInfo struct {
	// Value corresponding to the encoding type
	//
	// - When the encoding type is SKILL_ID, the value is the application\\"s Skill ID.
	//
	// - When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the user identifier from Tmall Genie, and each method corresponds to a different encoding type.
	//
	// - PACKAGE_NAME: APK package name, used for the Android application customer flow.
	//
	// - SKILL_ID: Skill ID, used for the cloud-based flow.
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
	// Type of User ID
	//
	// - OPEN_ID: Default User ID identifier.
	//
	// - UNION_ID: Organization-dimension User ID identifier. This value is available only after an organization has been requested on the Tmall Genie Skill Application Open Platform.
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

func (s AppUseTimeReportRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s AppUseTimeReportRequestUserInfo) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *AppUseTimeReportRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *AppUseTimeReportRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *AppUseTimeReportRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *AppUseTimeReportRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AppUseTimeReportRequestUserInfo) SetEncodeKey(v string) *AppUseTimeReportRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetEncodeType(v string) *AppUseTimeReportRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetId(v string) *AppUseTimeReportRequestUserInfo {
	s.Id = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetIdType(v string) *AppUseTimeReportRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetOrganizationId(v string) *AppUseTimeReportRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
