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
	// The alert configurations.
	AlertConfig []*UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// The alert contact group. The value is a JSON-formatted \\`List\\<string>\\`.
	//
	// example:
	//
	// ["test1","test2"]
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The type of the CNAME record. Valid value:
	//
	// - PUBLIC: The CNAME record is used for Internet access.
	//
	// example:
	//
	// PUBLIC
	CnameType *string `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	// Specifies whether to forcefully update the instance. Valid values:
	//
	// - true: Forcefully updates the instance without checking for conflicts.
	//
	// - false or null: Does not forcefully update the instance. The system checks for conflicts before the update.
	//
	// example:
	//
	// true
	ForceUpdate *bool `json:"ForceUpdate,omitempty" xml:"ForceUpdate,omitempty"`
	// The ID of the GTM instance. To obtain the instance ID, call the [DescribeDnsGtmInstances](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance. This parameter is required when you update the instance for the first time. It is optional for subsequent updates.
	//
	// example:
	//
	// test-1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language of the response. Valid values: en, zh, and ja. The default value is en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The method used to access the instance over the Internet. Valid values:
	//
	// - SYSTEM_ASSIGN: The system assigns a canonical name (CNAME) record. This option is disabled.
	//
	// - CUSTOM: You specify a CNAME record.
	//
	// example:
	//
	// CUSTOM
	PublicCnameMode *string `json:"PublicCnameMode,omitempty" xml:"PublicCnameMode,omitempty"`
	// The hostname of the CNAME record that is used for Internet access.
	//
	// example:
	//
	// test.rr
	PublicRr *string `json:"PublicRr,omitempty" xml:"PublicRr,omitempty"`
	// The service domain name that is accessed over the Internet.
	//
	// example:
	//
	// example.com
	PublicUserDomainName *string `json:"PublicUserDomainName,omitempty" xml:"PublicUserDomainName,omitempty"`
	// The primary domain name that is used to access the instance over the Internet using a CNAME record. This parameter is required if you set PublicCnameMode to CUSTOM.
	//
	// > Enter the primary domain name. Do not include the hostname specified by the PublicRr parameter.
	//
	// example:
	//
	// www.example.com
	PublicZoneName *string `json:"PublicZoneName,omitempty" xml:"PublicZoneName,omitempty"`
	// The global time to live (TTL).
	//
	// example:
	//
	// 60
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
	if s.AlertConfig != nil {
		for _, item := range s.AlertConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDnsGtmInstanceGlobalConfigRequestAlertConfig struct {
	// Specifies whether to send alerts through DingTalk. Valid values:
	//
	// - true: yes
	//
	// - false: no
	//
	// example:
	//
	// true
	DingtalkNotice *bool `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// Specifies whether to send alerts by email. Valid values:
	//
	// - true: yes
	//
	// - false or null: no
	//
	// example:
	//
	// true
	EmailNotice *bool `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// The type of the alert event. Valid values:
	//
	// - ADDR_ALERT: An address becomes unavailable.
	//
	// - ADDR_RESUME: An address becomes available.
	//
	// - ADDR_POOL_GROUP_UNAVAILABLE: An address pool group becomes unavailable.
	//
	// - ADDR_POOL_GROUP_AVAILABLE: An address pool group becomes available.
	//
	// - ACCESS_STRATEGY_POOL_GROUP_SWITCH: A switchover occurs between the primary and secondary address pools.
	//
	// - MONITOR_NODE_IP_CHANGE: The IP address of a monitoring node changes.
	//
	// example:
	//
	// ADDR_ALERT
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// Specifies whether to send alerts through text messages. Valid values:
	//
	// - true: yes
	//
	// - false or null: no
	//
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
