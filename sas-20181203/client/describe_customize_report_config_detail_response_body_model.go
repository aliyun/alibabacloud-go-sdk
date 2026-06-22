// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeReportConfigDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChartIds(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetChartIds() *string
	SetGroupType(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetGroupType() *string
	SetIsDefault(v int32) *DescribeCustomizeReportConfigDetailResponseBody
	GetIsDefault() *int32
	SetMemberAccountSyncFlag(v bool) *DescribeCustomizeReportConfigDetailResponseBody
	GetMemberAccountSyncFlag() *bool
	SetPinnedTime(v int64) *DescribeCustomizeReportConfigDetailResponseBody
	GetPinnedTime() *int64
	SetRecipients(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetRecipients() *string
	SetReportDays(v int32) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportDays() *int32
	SetReportEndDate(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportEndDate() *string
	SetReportId(v int64) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportId() *int64
	SetReportLang(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportLang() *string
	SetReportSendType(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportSendType() *string
	SetReportStartDate(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportStartDate() *string
	SetReportStatus(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportStatus() *string
	SetReportType(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetReportType() *string
	SetRequestId(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetRequestId() *string
	SetSendEndTime(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetSendEndTime() *string
	SetSendPeriodDays(v int32) *DescribeCustomizeReportConfigDetailResponseBody
	GetSendPeriodDays() *int32
	SetSendPeriodType(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetSendPeriodType() *string
	SetSendStartTime(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetSendStartTime() *string
	SetSendTime(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetSendTime() *string
	SetTargetGroups(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetTargetGroups() *string
	SetTargetUids(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetTargetUids() *string
	SetTitle(v string) *DescribeCustomizeReportConfigDetailResponseBody
	GetTitle() *string
}

type DescribeCustomizeReportConfigDetailResponseBody struct {
	// The report chart configuration IDs, separated by commas.
	//
	// example:
	//
	// BIZ_STAT_QUERY_KEY_ATTACK,CUSTOM_VUL_CVE_LIST,CUSTOM_VUL_SYS_LIST,CUSTOM_VUL_WEBCMS_LIST,CUSTOM_AUTO_BREAKING_PIE,CUSTOM_AK_LEAK_LIST,KEY_HP_TAMPERPROOF,KEY_HP_DEFENCE
	ChartIds *string `json:"ChartIds,omitempty" xml:"ChartIds,omitempty"`
	// The group type. Valid values:
	//
	// - **ALIYUN_RG**: ALIYUN_RG.
	//
	// - **SAS_GROUP**: SAS_GROUP.
	//
	// example:
	//
	// SAS_GROUP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// Indicates whether the report is a default report. Valid values:
	//
	// - **0**: Not a default report.
	//
	// - **1**: A default report.
	//
	// example:
	//
	// 1
	IsDefault *int32 `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Specifies whether newly added accounts are included by default. Valid values:
	//
	// - **true**: Included.
	//
	// - **false**: Not included.
	//
	// > Only version 2.0.0 supports this parameter.
	//
	// example:
	//
	// true
	MemberAccountSyncFlag *bool `json:"MemberAccountSyncFlag,omitempty" xml:"MemberAccountSyncFlag,omitempty"`
	// The pinned time.
	//
	// example:
	//
	// 1717430400000
	PinnedTime *int64 `json:"PinnedTime,omitempty" xml:"PinnedTime,omitempty"`
	// The recipient email addresses, separated by commas.
	//
	// example:
	//
	// PengZheng@eaton.com,ZhongJi@Eaton.com
	Recipients *string `json:"Recipients,omitempty" xml:"Recipients,omitempty"`
	// The number of recent days covered by the report statistics.
	//
	// example:
	//
	// 30
	ReportDays *int32 `json:"ReportDays,omitempty" xml:"ReportDays,omitempty"`
	// The end date for report delivery.
	//
	// example:
	//
	// 1720022399999
	ReportEndDate *string `json:"ReportEndDate,omitempty" xml:"ReportEndDate,omitempty"`
	// The report ID.
	//
	// example:
	//
	// 663434
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The language type. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	ReportLang *string `json:"ReportLang,omitempty" xml:"ReportLang,omitempty"`
	// The report delivery time range. Valid values:
	//
	// - **1**: 0:00 to 6:00.
	//
	// - **2**: 6:00 to 12:00.
	//
	// - **3**: 12:00 to 18:00.
	//
	// - **4**: 18:00 to 24:00.
	//
	// example:
	//
	// 2
	ReportSendType *string `json:"ReportSendType,omitempty" xml:"ReportSendType,omitempty"`
	// The start date for report delivery.
	//
	// example:
	//
	// 1717430400000
	ReportStartDate *string `json:"ReportStartDate,omitempty" xml:"ReportStartDate,omitempty"`
	// The report status. Valid values:
	//
	//  - **0**: Disabled.
	//
	//  - **1**: Enabled.
	//
	// example:
	//
	// 1
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The report type. Valid values:
	//
	// - **0**: Daily report.
	//
	// - **1**: Weekly report.
	//
	// - **2**: Monthly report.
	//
	// - **3**: Custom period.
	//
	// example:
	//
	// 3
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 379a9b8f-107b-4630-9e95-2299a1ea****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The delivery end time, in the format of HH:mm:ss.
	//
	// example:
	//
	// 10:00:00
	SendEndTime *string `json:"SendEndTime,omitempty" xml:"SendEndTime,omitempty"`
	// The specific execution dates within the delivery period.
	//
	// example:
	//
	// 12
	SendPeriodDays *int32 `json:"SendPeriodDays,omitempty" xml:"SendPeriodDays,omitempty"`
	// The delivery period type. Valid values:
	//
	// - **DAY**: day.
	//
	// - **WEEK**: week.
	//
	// - **MONTH**: month.
	//
	// example:
	//
	// MONTH
	SendPeriodType *string `json:"SendPeriodType,omitempty" xml:"SendPeriodType,omitempty"`
	// The delivery start time, in the format of HH:mm:ss.
	//
	// example:
	//
	// 09:00:00
	SendStartTime *string `json:"SendStartTime,omitempty" xml:"SendStartTime,omitempty"`
	// The delivery time, in the format of HH:mm:ss.
	//
	// example:
	//
	// 09:00:00
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// The targets within the group.
	//
	// example:
	//
	// 12125884,12140191
	TargetGroups *string `json:"TargetGroups,omitempty" xml:"TargetGroups,omitempty"`
	// The list of target UIDs, separated by commas.
	//
	// example:
	//
	// 1457515594445744,1600011353839072,1766185894104675,1674080148055995,1627510829033157
	TargetUids *string `json:"TargetUids,omitempty" xml:"TargetUids,omitempty"`
	// The title.
	//
	// example:
	//
	// marketing report
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeCustomizeReportConfigDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeReportConfigDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetChartIds() *string {
	return s.ChartIds
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetIsDefault() *int32 {
	return s.IsDefault
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetMemberAccountSyncFlag() *bool {
	return s.MemberAccountSyncFlag
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetPinnedTime() *int64 {
	return s.PinnedTime
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetRecipients() *string {
	return s.Recipients
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportDays() *int32 {
	return s.ReportDays
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportEndDate() *string {
	return s.ReportEndDate
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportLang() *string {
	return s.ReportLang
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportSendType() *string {
	return s.ReportSendType
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportStartDate() *string {
	return s.ReportStartDate
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetReportType() *string {
	return s.ReportType
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetSendEndTime() *string {
	return s.SendEndTime
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetSendPeriodDays() *int32 {
	return s.SendPeriodDays
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetSendPeriodType() *string {
	return s.SendPeriodType
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetSendStartTime() *string {
	return s.SendStartTime
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetSendTime() *string {
	return s.SendTime
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetTargetGroups() *string {
	return s.TargetGroups
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetTargetUids() *string {
	return s.TargetUids
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) GetTitle() *string {
	return s.Title
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetChartIds(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ChartIds = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetGroupType(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.GroupType = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetIsDefault(v int32) *DescribeCustomizeReportConfigDetailResponseBody {
	s.IsDefault = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetMemberAccountSyncFlag(v bool) *DescribeCustomizeReportConfigDetailResponseBody {
	s.MemberAccountSyncFlag = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetPinnedTime(v int64) *DescribeCustomizeReportConfigDetailResponseBody {
	s.PinnedTime = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetRecipients(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.Recipients = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportDays(v int32) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportDays = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportEndDate(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportEndDate = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportId(v int64) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportId = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportLang(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportLang = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportSendType(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportSendType = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportStartDate(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportStartDate = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportStatus(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportStatus = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetReportType(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.ReportType = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetRequestId(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetSendEndTime(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.SendEndTime = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetSendPeriodDays(v int32) *DescribeCustomizeReportConfigDetailResponseBody {
	s.SendPeriodDays = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetSendPeriodType(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.SendPeriodType = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetSendStartTime(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.SendStartTime = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetSendTime(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.SendTime = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetTargetGroups(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.TargetGroups = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetTargetUids(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.TargetUids = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) SetTitle(v string) *DescribeCustomizeReportConfigDetailResponseBody {
	s.Title = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
