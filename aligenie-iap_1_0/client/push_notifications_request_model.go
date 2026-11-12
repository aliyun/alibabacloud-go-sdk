// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNotificationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationUnicastRequest(v *PushNotificationsRequestNotificationUnicastRequest) *PushNotificationsRequest
	GetNotificationUnicastRequest() *PushNotificationsRequestNotificationUnicastRequest
	SetTenantInfo(v *PushNotificationsRequestTenantInfo) *PushNotificationsRequest
	GetTenantInfo() *PushNotificationsRequestTenantInfo
}

type PushNotificationsRequest struct {
	// Input parameter object for message push.
	//
	// This parameter is required.
	NotificationUnicastRequest *PushNotificationsRequestNotificationUnicastRequest `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty" type:"Struct"`
	// Identity information.
	TenantInfo *PushNotificationsRequestTenantInfo `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
}

func (s PushNotificationsRequest) String() string {
	return dara.Prettify(s)
}

func (s PushNotificationsRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequest) GetNotificationUnicastRequest() *PushNotificationsRequestNotificationUnicastRequest {
	return s.NotificationUnicastRequest
}

func (s *PushNotificationsRequest) GetTenantInfo() *PushNotificationsRequestTenantInfo {
	return s.TenantInfo
}

func (s *PushNotificationsRequest) SetNotificationUnicastRequest(v *PushNotificationsRequestNotificationUnicastRequest) *PushNotificationsRequest {
	s.NotificationUnicastRequest = v
	return s
}

func (s *PushNotificationsRequest) SetTenantInfo(v *PushNotificationsRequestTenantInfo) *PushNotificationsRequest {
	s.TenantInfo = v
	return s
}

func (s *PushNotificationsRequest) Validate() error {
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
	return nil
}

type PushNotificationsRequestNotificationUnicastRequest struct {
	// The value corresponding to the encoding type.
	//
	// - When the encoding type is SKILL_ID, the value is the skill ID obtained from the webhook service.
	//
	// - When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// apk包名
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the Tmall Genie device ID, and each method corresponds to a different encoding type:
	//
	// - PACKAGE_NAME: APK package name
	//
	// - SKILL_ID: Skill ID
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Debug flag. When the skill application is in the debug phase, setting this flag determines whether the retrieved skill application information is from the editing state or the online state.
	//
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// Message template ID obtained when applying for a message template on the Tmall Genie Application Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2iU81*****G9elJ
	MessageTemplateId *string `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	// Organization Identifier. This parameter is required only when the push type is `XX_UNION_XX`. If you obtain CatGenie device or user identifiers through multiple channels and need interoperability among them, you must request an organization from the platform. After approval, the platform will assign an organization identifier.
	//
	// example:
	//
	// 2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// Placeholder information. For example, for the template ["Hello, {nick}!"], this could be {"nick":"Sweetie"}. Additionally, within the placeholder parameters, you can append parameters to the application launch link to support hyperlinking to a specific address, such as: `{"nick":"Sweetie","_URL_SUFFIX_":"xxx"}`. The hyperlink address in the message pop-up will then be: application launch link + `${_URL_SUFFIX_}`. This feature only works for applications that have configured a launch link.
	//
	// example:
	//
	// {"nick":"张三"}
	PlaceHolder map[string]*string `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	// The target information for message push.
	//
	// This parameter is required.
	SendTarget *PushNotificationsRequestNotificationUnicastRequestSendTarget `json:"SendTarget,omitempty" xml:"SendTarget,omitempty" type:"Struct"`
}

func (s PushNotificationsRequestNotificationUnicastRequest) String() string {
	return dara.Prettify(s)
}

func (s PushNotificationsRequestNotificationUnicastRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestNotificationUnicastRequest) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PushNotificationsRequestNotificationUnicastRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PushNotificationsRequestNotificationUnicastRequest) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *PushNotificationsRequestNotificationUnicastRequest) GetMessageTemplateId() *string {
	return s.MessageTemplateId
}

func (s *PushNotificationsRequestNotificationUnicastRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PushNotificationsRequestNotificationUnicastRequest) GetPlaceHolder() map[string]*string {
	return s.PlaceHolder
}

func (s *PushNotificationsRequestNotificationUnicastRequest) GetSendTarget() *PushNotificationsRequestNotificationUnicastRequestSendTarget {
	return s.SendTarget
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetEncodeKey(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.EncodeKey = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetEncodeType(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.EncodeType = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetIsDebug(v bool) *PushNotificationsRequestNotificationUnicastRequest {
	s.IsDebug = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetMessageTemplateId(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.MessageTemplateId = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetOrganizationId(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.OrganizationId = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetPlaceHolder(v map[string]*string) *PushNotificationsRequestNotificationUnicastRequest {
	s.PlaceHolder = v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetSendTarget(v *PushNotificationsRequestNotificationUnicastRequestSendTarget) *PushNotificationsRequestNotificationUnicastRequest {
	s.SendTarget = v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) Validate() error {
	if s.SendTarget != nil {
		if err := s.SendTarget.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushNotificationsRequestNotificationUnicastRequestSendTarget struct {
	// The identity value corresponding to the Push Target Type.
	//
	// example:
	//
	// 2VpiDQ6aMjxz******Eo7r6e08oIVZ3fKrm5TyEfY=
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
	// The type of the Push Target, indicating the type when obtaining the corresponding device or User Identifier:
	//
	// - `DEVICE_UNION_ID`: device unionId
	//
	// - `DEVICE_OPEN_ID`: device openId
	//
	// - `USER_UNION_ID`: user unionId
	//
	// - `USER_OPEN_ID`: user openId
	//
	// example:
	//
	// DEVICE_OPEN_ID
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s PushNotificationsRequestNotificationUnicastRequestSendTarget) String() string {
	return dara.Prettify(s)
}

func (s PushNotificationsRequestNotificationUnicastRequestSendTarget) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) GetTargetIdentity() *string {
	return s.TargetIdentity
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) SetTargetIdentity(v string) *PushNotificationsRequestNotificationUnicastRequestSendTarget {
	s.TargetIdentity = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) SetTargetType(v string) *PushNotificationsRequestNotificationUnicastRequestSendTarget {
	s.TargetType = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) Validate() error {
	return dara.Validate(s)
}

type PushNotificationsRequestTenantInfo struct {
}

func (s PushNotificationsRequestTenantInfo) String() string {
	return dara.Prettify(s)
}

func (s PushNotificationsRequestTenantInfo) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestTenantInfo) Validate() error {
	return dara.Validate(s)
}
