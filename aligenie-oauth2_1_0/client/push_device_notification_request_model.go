// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDeviceNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantInfo(v *PushDeviceNotificationRequestTenantInfo) *PushDeviceNotificationRequest
	GetTenantInfo() *PushDeviceNotificationRequestTenantInfo
	SetBody(v *PushDeviceNotificationRequestBody) *PushDeviceNotificationRequest
	GetBody() *PushDeviceNotificationRequestBody
}

type PushDeviceNotificationRequest struct {
	TenantInfo *PushDeviceNotificationRequestTenantInfo `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
	Body       *PushDeviceNotificationRequestBody       `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s PushDeviceNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceNotificationRequest) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequest) GetTenantInfo() *PushDeviceNotificationRequestTenantInfo {
	return s.TenantInfo
}

func (s *PushDeviceNotificationRequest) GetBody() *PushDeviceNotificationRequestBody {
	return s.Body
}

func (s *PushDeviceNotificationRequest) SetTenantInfo(v *PushDeviceNotificationRequestTenantInfo) *PushDeviceNotificationRequest {
	s.TenantInfo = v
	return s
}

func (s *PushDeviceNotificationRequest) SetBody(v *PushDeviceNotificationRequestBody) *PushDeviceNotificationRequest {
	s.Body = v
	return s
}

func (s *PushDeviceNotificationRequest) Validate() error {
	return dara.Validate(s)
}

type PushDeviceNotificationRequestTenantInfo struct {
	// example:
	//
	// 12797******304102
	SubjectId *string `json:"SubjectId,omitempty" xml:"SubjectId,omitempty"`
}

func (s PushDeviceNotificationRequestTenantInfo) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceNotificationRequestTenantInfo) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequestTenantInfo) GetSubjectId() *string {
	return s.SubjectId
}

func (s *PushDeviceNotificationRequestTenantInfo) SetSubjectId(v string) *PushDeviceNotificationRequestTenantInfo {
	s.SubjectId = &v
	return s
}

func (s *PushDeviceNotificationRequestTenantInfo) Validate() error {
	return dara.Validate(s)
}

type PushDeviceNotificationRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1923792******8R7392
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// false
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2iU81*****G9elJ
	MessageTemplateId *string `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	// example:
	//
	// 29837******2938
	OrganizationId *string            `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	PlaceHolder    map[string]*string `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	// This parameter is required.
	SendTarget *PushDeviceNotificationRequestBodySendTarget `json:"SendTarget,omitempty" xml:"SendTarget,omitempty" type:"Struct"`
}

func (s PushDeviceNotificationRequestBody) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceNotificationRequestBody) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequestBody) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *PushDeviceNotificationRequestBody) GetEncodeType() *string {
	return s.EncodeType
}

func (s *PushDeviceNotificationRequestBody) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *PushDeviceNotificationRequestBody) GetMessageTemplateId() *string {
	return s.MessageTemplateId
}

func (s *PushDeviceNotificationRequestBody) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *PushDeviceNotificationRequestBody) GetPlaceHolder() map[string]*string {
	return s.PlaceHolder
}

func (s *PushDeviceNotificationRequestBody) GetSendTarget() *PushDeviceNotificationRequestBodySendTarget {
	return s.SendTarget
}

func (s *PushDeviceNotificationRequestBody) SetEncodeKey(v string) *PushDeviceNotificationRequestBody {
	s.EncodeKey = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetEncodeType(v string) *PushDeviceNotificationRequestBody {
	s.EncodeType = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetIsDebug(v bool) *PushDeviceNotificationRequestBody {
	s.IsDebug = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetMessageTemplateId(v string) *PushDeviceNotificationRequestBody {
	s.MessageTemplateId = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetOrganizationId(v string) *PushDeviceNotificationRequestBody {
	s.OrganizationId = &v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetPlaceHolder(v map[string]*string) *PushDeviceNotificationRequestBody {
	s.PlaceHolder = v
	return s
}

func (s *PushDeviceNotificationRequestBody) SetSendTarget(v *PushDeviceNotificationRequestBodySendTarget) *PushDeviceNotificationRequestBody {
	s.SendTarget = v
	return s
}

func (s *PushDeviceNotificationRequestBody) Validate() error {
	return dara.Validate(s)
}

type PushDeviceNotificationRequestBodySendTarget struct {
	// example:
	//
	// 2VpiDQ6aMjxz******Eo7r6e08oIVZ3fKrm5TyEfY=
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
	// example:
	//
	// DEVICE_OPEN_ID
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s PushDeviceNotificationRequestBodySendTarget) String() string {
	return dara.Prettify(s)
}

func (s PushDeviceNotificationRequestBodySendTarget) GoString() string {
	return s.String()
}

func (s *PushDeviceNotificationRequestBodySendTarget) GetTargetIdentity() *string {
	return s.TargetIdentity
}

func (s *PushDeviceNotificationRequestBodySendTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *PushDeviceNotificationRequestBodySendTarget) SetTargetIdentity(v string) *PushDeviceNotificationRequestBodySendTarget {
	s.TargetIdentity = &v
	return s
}

func (s *PushDeviceNotificationRequestBodySendTarget) SetTargetType(v string) *PushDeviceNotificationRequestBodySendTarget {
	s.TargetType = &v
	return s
}

func (s *PushDeviceNotificationRequestBodySendTarget) Validate() error {
	return dara.Validate(s)
}
