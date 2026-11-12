// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *SendNotificationsRequestDeviceInfo) *SendNotificationsRequest
	GetDeviceInfo() *SendNotificationsRequestDeviceInfo
	SetNotificationUnicastRequest(v *SendNotificationsRequestNotificationUnicastRequest) *SendNotificationsRequest
	GetNotificationUnicastRequest() *SendNotificationsRequestNotificationUnicastRequest
	SetTenantInfo(v *SendNotificationsRequestTenantInfo) *SendNotificationsRequest
	GetTenantInfo() *SendNotificationsRequestTenantInfo
	SetUserInfo(v *SendNotificationsRequestUserInfo) *SendNotificationsRequest
	GetUserInfo() *SendNotificationsRequestUserInfo
}

type SendNotificationsRequest struct {
	// Device identity information
	//
	// This parameter is required.
	DeviceInfo *SendNotificationsRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Input parameter object for message push.
	//
	// This parameter is required.
	NotificationUnicastRequest *SendNotificationsRequestNotificationUnicastRequest `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty" type:"Struct"`
	// Identity information.
	TenantInfo *SendNotificationsRequestTenantInfo `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
	// User Identifier information
	//
	// This parameter is required.
	UserInfo *SendNotificationsRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SendNotificationsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequest) GetDeviceInfo() *SendNotificationsRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *SendNotificationsRequest) GetNotificationUnicastRequest() *SendNotificationsRequestNotificationUnicastRequest {
	return s.NotificationUnicastRequest
}

func (s *SendNotificationsRequest) GetTenantInfo() *SendNotificationsRequestTenantInfo {
	return s.TenantInfo
}

func (s *SendNotificationsRequest) GetUserInfo() *SendNotificationsRequestUserInfo {
	return s.UserInfo
}

func (s *SendNotificationsRequest) SetDeviceInfo(v *SendNotificationsRequestDeviceInfo) *SendNotificationsRequest {
	s.DeviceInfo = v
	return s
}

func (s *SendNotificationsRequest) SetNotificationUnicastRequest(v *SendNotificationsRequestNotificationUnicastRequest) *SendNotificationsRequest {
	s.NotificationUnicastRequest = v
	return s
}

func (s *SendNotificationsRequest) SetTenantInfo(v *SendNotificationsRequestTenantInfo) *SendNotificationsRequest {
	s.TenantInfo = v
	return s
}

func (s *SendNotificationsRequest) SetUserInfo(v *SendNotificationsRequestUserInfo) *SendNotificationsRequest {
	s.UserInfo = v
	return s
}

func (s *SendNotificationsRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.NotificationUnicastRequest != nil {
		if err := s.NotificationUnicastRequest.Validate(); err != nil {
			return err
		}
	}
	if s.TenantInfo != nil {
		if err := s.TenantInfo.Validate(); err != nil {
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

type SendNotificationsRequestDeviceInfo struct {
	// Value corresponding to the encoding type. If the encoding type is SKILL_ID, the value is the application\\"s Skill ID. If the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the device ID for Maojing, and each method corresponds to a different encoding type:
	//
	//  - PACKAGE_NAME: APK package name, used for the Android application customer link
	//
	//  - SKILL_ID: Skill ID, used for the cloud link
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID (deviceOpenId or deviceUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of device ID:
	//
	//  - OPEN_ID: Default device ID
	//
	//  - UNION_ID: Organization-dimension device ID, available only after an organization has been requested on the Maojing Skill Application Open Platform
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

func (s SendNotificationsRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *SendNotificationsRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *SendNotificationsRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *SendNotificationsRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *SendNotificationsRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *SendNotificationsRequestDeviceInfo) SetEncodeKey(v string) *SendNotificationsRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetEncodeType(v string) *SendNotificationsRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetId(v string) *SendNotificationsRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetIdType(v string) *SendNotificationsRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetOrganizationId(v string) *SendNotificationsRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type SendNotificationsRequestNotificationUnicastRequest struct {
	// Debug flag. When the skill application is in the debugging phase, setting this flag determines whether the retrieved skill application information is from the editing state or the online state.
	//
	// example:
	//
	// false
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// Message template ID obtained when requesting a message template in the Tmall Genie Application Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2iU81*****G9elJ
	MessageTemplateId *string `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	// Placeholder information. For example, for the template ["Hello, {nick}!"], this can be: {"nick":"Xiao Tian Tian"}. Additionally, within the placeholder parameters, you can append parameters to the application launch link to support hyperlinking to a specific address, such as: `{"nick":"Xiao Tian Tian","_URL_SUFFIX_":"xxx"}`. The hyperlink address in the message pop-up will then be: application launch link + `${_URL_SUFFIX_}`. This feature is only effective for applications that have configured a launch link.
	//
	// example:
	//
	// {"nick":"张三"}
	PlaceHolder map[string]*string `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	// Target information for message push.
	//
	// This parameter is required.
	SendTarget *SendNotificationsRequestNotificationUnicastRequestSendTarget `json:"SendTarget,omitempty" xml:"SendTarget,omitempty" type:"Struct"`
}

func (s SendNotificationsRequestNotificationUnicastRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsRequestNotificationUnicastRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestNotificationUnicastRequest) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *SendNotificationsRequestNotificationUnicastRequest) GetMessageTemplateId() *string {
	return s.MessageTemplateId
}

func (s *SendNotificationsRequestNotificationUnicastRequest) GetPlaceHolder() map[string]*string {
	return s.PlaceHolder
}

func (s *SendNotificationsRequestNotificationUnicastRequest) GetSendTarget() *SendNotificationsRequestNotificationUnicastRequestSendTarget {
	return s.SendTarget
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetIsDebug(v bool) *SendNotificationsRequestNotificationUnicastRequest {
	s.IsDebug = &v
	return s
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetMessageTemplateId(v string) *SendNotificationsRequestNotificationUnicastRequest {
	s.MessageTemplateId = &v
	return s
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetPlaceHolder(v map[string]*string) *SendNotificationsRequestNotificationUnicastRequest {
	s.PlaceHolder = v
	return s
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetSendTarget(v *SendNotificationsRequestNotificationUnicastRequestSendTarget) *SendNotificationsRequestNotificationUnicastRequest {
	s.SendTarget = v
	return s
}

func (s *SendNotificationsRequestNotificationUnicastRequest) Validate() error {
	if s.SendTarget != nil {
		if err := s.SendTarget.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNotificationsRequestNotificationUnicastRequestSendTarget struct {
}

func (s SendNotificationsRequestNotificationUnicastRequestSendTarget) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsRequestNotificationUnicastRequestSendTarget) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestNotificationUnicastRequestSendTarget) Validate() error {
	return dara.Validate(s)
}

type SendNotificationsRequestTenantInfo struct {
}

func (s SendNotificationsRequestTenantInfo) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsRequestTenantInfo) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestTenantInfo) Validate() error {
	return dara.Validate(s)
}

type SendNotificationsRequestUserInfo struct {
	// Value corresponding to the encoding type. If the encoding type is SKILL_ID, the value is the application\\"s Skill ID. If the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12**45
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the Maojing User Identifier, and each method corresponds to a different encoding type:
	//
	// - PACKAGE_NAME: APK package name, used for the Android application customer link
	//
	// - SKILL_ID: Skill ID, used for the cloud-based link
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// User Identifier (userOpenId or userUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// HOFF****my7Iw=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Type of the User ID:
	//
	// - OPEN_ID: Default User ID identifier
	//
	// - UNION_ID: Organization-dimension User ID identifier, available only after an organization has been requested on the Maojing Skill Application Open Platform
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

func (s SendNotificationsRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *SendNotificationsRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *SendNotificationsRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *SendNotificationsRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *SendNotificationsRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *SendNotificationsRequestUserInfo) SetEncodeKey(v string) *SendNotificationsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetEncodeType(v string) *SendNotificationsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetId(v string) *SendNotificationsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetIdType(v string) *SendNotificationsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetOrganizationId(v string) *SendNotificationsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
