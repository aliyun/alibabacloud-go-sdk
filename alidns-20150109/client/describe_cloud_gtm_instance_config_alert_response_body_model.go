// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmInstanceConfigAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig) *DescribeCloudGtmInstanceConfigAlertResponseBody
	GetAlertConfig() *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig
	SetAlertGroup(v *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup) *DescribeCloudGtmInstanceConfigAlertResponseBody
	GetAlertGroup() *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup
	SetAlertMode(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody
	GetAlertMode() *string
	SetConfigId(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody
	GetConfigId() *string
	SetInstanceId(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody
	GetRequestId() *string
}

type DescribeCloudGtmInstanceConfigAlertResponseBody struct {
	// The alert configurations.
	AlertConfig *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	// The alert contact groups.
	AlertGroup *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Struct"`
	// The alert configuration mode of the instance. Valid values:
	//
	// 	- global: global alert configuration
	//
	// 	- instance_config: custom alert configuration
	//
	// example:
	//
	// global
	AlertMode *string `json:"AlertMode,omitempty" xml:"AlertMode,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) GetAlertConfig() *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig {
	return s.AlertConfig
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) GetAlertGroup() *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup {
	return s.AlertGroup
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) GetAlertMode() *string {
	return s.AlertMode
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) SetAlertConfig(v *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig) *DescribeCloudGtmInstanceConfigAlertResponseBody {
	s.AlertConfig = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) SetAlertGroup(v *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup) *DescribeCloudGtmInstanceConfigAlertResponseBody {
	s.AlertGroup = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) SetAlertMode(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody {
	s.AlertMode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) SetConfigId(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody {
	s.ConfigId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) SetInstanceId(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) SetRequestId(v string) *DescribeCloudGtmInstanceConfigAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig struct {
	AlertConfig []*DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig) GetAlertConfig() []*DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig) SetAlertConfig(v []*DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig struct {
	// Indicates whether DingTalk notifications are configured. Valid values:
	//
	// 	- true: DingTalk notifications are configured. DingTalk notifications are sent after alerts are triggered.
	//
	// 	- false: DingTalk notifications are not configured.
	//
	// example:
	//
	// true
	DingtalkNotice *bool `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// Indicates whether email notifications are configured. Valid values:
	//
	// 	- true: Email notifications are configured. Emails are sent after alerts are triggered.
	//
	// 	- false: Email notifications are not configured.
	//
	// example:
	//
	// true
	EmailNotice *bool `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// The type of the alert event. Valid values:
	//
	// 	- addr_alert: The address is unavailable.
	//
	// 	- addr_resume: The address becomes available.
	//
	// 	- addr_pool_unavailable: The address pool is unavailable.
	//
	// 	- addr_pool_available: The address pool becomes available.
	//
	// example:
	//
	// addr_alert
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// Indicates whether text message notifications are configured. Valid values:
	//
	// 	- true: Text message notifications are configured. Text messages are sent after alerts are triggered.
	//
	// 	- false: Text message notifications are not configured.
	//
	// Only the China site (aliyun.com) supports text message notifications.
	//
	// example:
	//
	// true
	SmsNotice *bool `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) GetDingtalkNotice() *bool {
	return s.DingtalkNotice
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) GetEmailNotice() *bool {
	return s.EmailNotice
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) SetDingtalkNotice(v bool) *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) SetEmailNotice(v bool) *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) SetNoticeType(v string) *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) SetSmsNotice(v bool) *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertConfigAlertConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup struct {
	AlertGroup []*string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup) GetAlertGroup() []*string {
	return s.AlertGroup
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup) SetAlertGroup(v []*string) *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup {
	s.AlertGroup = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponseBodyAlertGroup) Validate() error {
	return dara.Validate(s)
}
