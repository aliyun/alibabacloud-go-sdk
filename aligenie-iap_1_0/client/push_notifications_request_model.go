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
	// This parameter is required.
	NotificationUnicastRequest *PushNotificationsRequestNotificationUnicastRequest `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty" type:"Struct"`
	TenantInfo                 *PushNotificationsRequestTenantInfo                 `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type PushNotificationsRequestNotificationUnicastRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// apk包名
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iU81*****G9elJ
	MessageTemplateId *string `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	// example:
	//
	// 2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// {"nick":"张三"}
	PlaceHolder map[string]*string `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
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
	return dara.Validate(s)
}

type PushNotificationsRequestNotificationUnicastRequestSendTarget struct {
	// example:
	//
	// 2VpiDQ6aMjxz******Eo7r6e08oIVZ3fKrm5TyEfY=
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
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
