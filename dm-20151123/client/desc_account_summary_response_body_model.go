// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescAccountSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDailyQuota(v int32) *DescAccountSummaryResponseBody
	GetDailyQuota() *int32
	SetDailyRemainFreeQuota(v int32) *DescAccountSummaryResponseBody
	GetDailyRemainFreeQuota() *int32
	SetDayuStatus(v int32) *DescAccountSummaryResponseBody
	GetDayuStatus() *int32
	SetDomains(v int32) *DescAccountSummaryResponseBody
	GetDomains() *int32
	SetEnableTimes(v int32) *DescAccountSummaryResponseBody
	GetEnableTimes() *int32
	SetIpChannelType(v string) *DescAccountSummaryResponseBody
	GetIpChannelType() *string
	SetMailAddresses(v int32) *DescAccountSummaryResponseBody
	GetMailAddresses() *int32
	SetMaxQuotaLevel(v int32) *DescAccountSummaryResponseBody
	GetMaxQuotaLevel() *int32
	SetMonthQuota(v int32) *DescAccountSummaryResponseBody
	GetMonthQuota() *int32
	SetQuotaLevel(v int32) *DescAccountSummaryResponseBody
	GetQuotaLevel() *int32
	SetReceivers(v int32) *DescAccountSummaryResponseBody
	GetReceivers() *int32
	SetRemainFreeQuota(v int32) *DescAccountSummaryResponseBody
	GetRemainFreeQuota() *int32
	SetRequestId(v string) *DescAccountSummaryResponseBody
	GetRequestId() *string
	SetSmsRecord(v int32) *DescAccountSummaryResponseBody
	GetSmsRecord() *int32
	SetSmsSign(v int32) *DescAccountSummaryResponseBody
	GetSmsSign() *int32
	SetSmsTemplates(v int32) *DescAccountSummaryResponseBody
	GetSmsTemplates() *int32
	SetTags(v int32) *DescAccountSummaryResponseBody
	GetTags() *int32
	SetTemplates(v int32) *DescAccountSummaryResponseBody
	GetTemplates() *int32
	SetUserStatus(v int32) *DescAccountSummaryResponseBody
	GetUserStatus() *int32
}

type DescAccountSummaryResponseBody struct {
	// Daily quota
	//
	// example:
	//
	// 2000
	DailyQuota *int32 `json:"DailyQuota,omitempty" xml:"DailyQuota,omitempty"`
	// remaining amount of daily free quota
	//
	// example:
	//
	// 100
	DailyRemainFreeQuota *int32 `json:"DailyRemainFreeQuota,omitempty" xml:"DailyRemainFreeQuota,omitempty"`
	// Dayu status (deprecated, retained for compatibility reasons.)
	//
	// example:
	//
	// 0
	DayuStatus *int32 `json:"DayuStatus,omitempty" xml:"DayuStatus,omitempty"`
	// Number of domains
	//
	// example:
	//
	// 1
	Domains *int32 `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// Effective time
	//
	// example:
	//
	// 0
	EnableTimes *int32 `json:"EnableTimes,omitempty" xml:"EnableTimes,omitempty"`
	// if can be null:
	// true
	IpChannelType *string `json:"IpChannelType,omitempty" xml:"IpChannelType,omitempty"`
	// Number of sending addresses
	//
	// example:
	//
	// 0
	MailAddresses *int32 `json:"MailAddresses,omitempty" xml:"MailAddresses,omitempty"`
	// Maximum level
	//
	// example:
	//
	// 10
	MaxQuotaLevel *int32 `json:"MaxQuotaLevel,omitempty" xml:"MaxQuotaLevel,omitempty"`
	// Monthly quota
	//
	// example:
	//
	// 60000
	MonthQuota *int32 `json:"MonthQuota,omitempty" xml:"MonthQuota,omitempty"`
	// Credit level
	//
	// example:
	//
	// 2
	QuotaLevel *int32 `json:"QuotaLevel,omitempty" xml:"QuotaLevel,omitempty"`
	// Number of recipients
	//
	// example:
	//
	// 0
	Receivers *int32 `json:"Receivers,omitempty" xml:"Receivers,omitempty"`
	// Remaining amount of total free quota
	//
	// example:
	//
	// 1910
	RemainFreeQuota *int32 `json:"RemainFreeQuota,omitempty" xml:"RemainFreeQuota,omitempty"`
	// Request ID
	//
	// example:
	//
	// 82B295BB-7E69-491F-9896-ECEAFF09E1A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Deprecated, retained for compatibility reasons.
	//
	// example:
	//
	// 0
	SmsRecord *int32 `json:"SmsRecord,omitempty" xml:"SmsRecord,omitempty"`
	// Deprecated, retained for compatibility reasons.
	//
	// example:
	//
	// 0
	SmsSign *int32 `json:"SmsSign,omitempty" xml:"SmsSign,omitempty"`
	// Deprecated, retained for compatibility reasons.
	//
	// example:
	//
	// 0
	SmsTemplates *int32 `json:"SmsTemplates,omitempty" xml:"SmsTemplates,omitempty"`
	// Number of tags
	//
	// example:
	//
	// 0
	Tags *int32 `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Number of templates
	//
	// example:
	//
	// 1
	Templates *int32 `json:"Templates,omitempty" xml:"Templates,omitempty"`
	// User status:
	//
	// 1 Frozen
	//
	// 2 In arrears
	//
	// 4 Restricted from sending
	//
	// 8 Logically deleted
	//
	// example:
	//
	// 0
	UserStatus *int32 `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescAccountSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescAccountSummaryResponseBody) GetDailyQuota() *int32 {
	return s.DailyQuota
}

func (s *DescAccountSummaryResponseBody) GetDailyRemainFreeQuota() *int32 {
	return s.DailyRemainFreeQuota
}

func (s *DescAccountSummaryResponseBody) GetDayuStatus() *int32 {
	return s.DayuStatus
}

func (s *DescAccountSummaryResponseBody) GetDomains() *int32 {
	return s.Domains
}

func (s *DescAccountSummaryResponseBody) GetEnableTimes() *int32 {
	return s.EnableTimes
}

func (s *DescAccountSummaryResponseBody) GetIpChannelType() *string {
	return s.IpChannelType
}

func (s *DescAccountSummaryResponseBody) GetMailAddresses() *int32 {
	return s.MailAddresses
}

func (s *DescAccountSummaryResponseBody) GetMaxQuotaLevel() *int32 {
	return s.MaxQuotaLevel
}

func (s *DescAccountSummaryResponseBody) GetMonthQuota() *int32 {
	return s.MonthQuota
}

func (s *DescAccountSummaryResponseBody) GetQuotaLevel() *int32 {
	return s.QuotaLevel
}

func (s *DescAccountSummaryResponseBody) GetReceivers() *int32 {
	return s.Receivers
}

func (s *DescAccountSummaryResponseBody) GetRemainFreeQuota() *int32 {
	return s.RemainFreeQuota
}

func (s *DescAccountSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescAccountSummaryResponseBody) GetSmsRecord() *int32 {
	return s.SmsRecord
}

func (s *DescAccountSummaryResponseBody) GetSmsSign() *int32 {
	return s.SmsSign
}

func (s *DescAccountSummaryResponseBody) GetSmsTemplates() *int32 {
	return s.SmsTemplates
}

func (s *DescAccountSummaryResponseBody) GetTags() *int32 {
	return s.Tags
}

func (s *DescAccountSummaryResponseBody) GetTemplates() *int32 {
	return s.Templates
}

func (s *DescAccountSummaryResponseBody) GetUserStatus() *int32 {
	return s.UserStatus
}

func (s *DescAccountSummaryResponseBody) SetDailyQuota(v int32) *DescAccountSummaryResponseBody {
	s.DailyQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDailyRemainFreeQuota(v int32) *DescAccountSummaryResponseBody {
	s.DailyRemainFreeQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDayuStatus(v int32) *DescAccountSummaryResponseBody {
	s.DayuStatus = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetDomains(v int32) *DescAccountSummaryResponseBody {
	s.Domains = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetEnableTimes(v int32) *DescAccountSummaryResponseBody {
	s.EnableTimes = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetIpChannelType(v string) *DescAccountSummaryResponseBody {
	s.IpChannelType = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMailAddresses(v int32) *DescAccountSummaryResponseBody {
	s.MailAddresses = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMaxQuotaLevel(v int32) *DescAccountSummaryResponseBody {
	s.MaxQuotaLevel = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetMonthQuota(v int32) *DescAccountSummaryResponseBody {
	s.MonthQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetQuotaLevel(v int32) *DescAccountSummaryResponseBody {
	s.QuotaLevel = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetReceivers(v int32) *DescAccountSummaryResponseBody {
	s.Receivers = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetRemainFreeQuota(v int32) *DescAccountSummaryResponseBody {
	s.RemainFreeQuota = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetRequestId(v string) *DescAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsRecord(v int32) *DescAccountSummaryResponseBody {
	s.SmsRecord = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsSign(v int32) *DescAccountSummaryResponseBody {
	s.SmsSign = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetSmsTemplates(v int32) *DescAccountSummaryResponseBody {
	s.SmsTemplates = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetTags(v int32) *DescAccountSummaryResponseBody {
	s.Tags = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetTemplates(v int32) *DescAccountSummaryResponseBody {
	s.Templates = &v
	return s
}

func (s *DescAccountSummaryResponseBody) SetUserStatus(v int32) *DescAccountSummaryResponseBody {
	s.UserStatus = &v
	return s
}

func (s *DescAccountSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
