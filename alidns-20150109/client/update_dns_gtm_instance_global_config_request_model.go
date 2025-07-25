// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmInstanceGlobalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetAlertConfig() []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig
	SetAlertGroup(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetAlertGroup() *string
	SetCnameType(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetCnameType() *string
	SetForceUpdate(v bool) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetForceUpdate() *bool
	SetInstanceId(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetInstanceName() *string
	SetLang(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetLang() *string
	SetPublicCnameMode(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetPublicCnameMode() *string
	SetPublicRr(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetPublicRr() *string
	SetPublicUserDomainName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetPublicUserDomainName() *string
	SetPublicZoneName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetPublicZoneName() *string
	SetTtl(v int32) *UpdateDnsGtmInstanceGlobalConfigRequest
	GetTtl() *int32
}

type UpdateDnsGtmInstanceGlobalConfigRequest struct {
	AlertConfig []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// The name of the alert group in the JSON format.
	//
	// example:
	//
	// alertGroup1
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The type of the canonical name (CNAME).
	//
	// 	- Set the value to PUBLIC.
	//
	// example:
	//
	// public
	CnameType *string `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	// Specifies whether to enable force updates. Valid values:
	//
	// 	- true: enables force update without a conflict alert.
	//
	// 	- false: disables force update. If a conflict occurs, the system displays an alert. null: This valid value of ForceUpdate provides the same information as the false value.
	//
	// example:
	//
	// true
	ForceUpdate *bool `json:"ForceUpdate,omitempty" xml:"ForceUpdate,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance. This parameter is required only for the first update.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to use a custom CNAME domain name or a CNAME domain name assigned by the system to access the instance over the Internet. Valid values:
	//
	// 	- SYSTEM_ASSIGN: a CNAME domain name assigned by the system
	//
	// 	- CUSTOM: a custom CNAME domain name
	//
	// example:
	//
	// custom
	PublicCnameMode *string `json:"PublicCnameMode,omitempty" xml:"PublicCnameMode,omitempty"`
	// The hostname corresponding to the CNAME domain name that is used to access the instance over the Internet.
	//
	// example:
	//
	// test.rr
	PublicRr *string `json:"PublicRr,omitempty" xml:"PublicRr,omitempty"`
	// The service domain name that is used over the Internet.
	//
	// example:
	//
	// example.com
	PublicUserDomainName *string `json:"PublicUserDomainName,omitempty" xml:"PublicUserDomainName,omitempty"`
	// The CNAME domain name that is used to access the instance over the Internet, which is the primary domain name. This parameter is required when the PublicCnameMode parameter is set to CUSTOM.
	//
	// >  You must use the primary domain name. Do not include the hostname specified by the PublicRr parameter.
	//
	// example:
	//
	// gtm-003.com
	PublicZoneName *string `json:"PublicZoneName,omitempty" xml:"PublicZoneName,omitempty"`
	// The global time to live (TTL).
	//
	// example:
	//
	// 1
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateDnsGtmInstanceGlobalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetAlertConfig() []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	return s.AlertConfig
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetCnameType() *string {
	return s.CnameType
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetForceUpdate() *bool {
	return s.ForceUpdate
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetPublicCnameMode() *string {
	return s.PublicCnameMode
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetPublicRr() *string {
	return s.PublicRr
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetPublicUserDomainName() *string {
	return s.PublicUserDomainName
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetPublicZoneName() *string {
	return s.PublicZoneName
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetAlertConfig(v []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.AlertConfig = v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetAlertGroup(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.AlertGroup = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetCnameType(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.CnameType = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetForceUpdate(v bool) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.ForceUpdate = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetInstanceId(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetInstanceName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetLang(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetPublicCnameMode(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.PublicCnameMode = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetPublicRr(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.PublicRr = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetPublicUserDomainName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.PublicUserDomainName = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetPublicZoneName(v string) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.PublicZoneName = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) SetTtl(v int32) *UpdateDnsGtmInstanceGlobalConfigRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig struct {
	// example:
	//
	// true
	DingtalkNotice *bool `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// example:
	//
	// true
	EmailNotice *bool `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// example:
	//
	// ADDR_ALERT
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// example:
	//
	// true
	SmsNotice *bool `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
}

func (s UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) GetDingtalkNotice() *bool {
	return s.DingtalkNotice
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) GetEmailNotice() *bool {
	return s.EmailNotice
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) SetDingtalkNotice(v bool) *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) SetEmailNotice(v bool) *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) SetNoticeType(v string) *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) SetSmsNotice(v bool) *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}
