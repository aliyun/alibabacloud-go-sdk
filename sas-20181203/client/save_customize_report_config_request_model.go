// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCustomizeReportConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupType(v string) *SaveCustomizeReportConfigRequest
	GetGroupType() *string
	SetMemberAccountSyncFlag(v bool) *SaveCustomizeReportConfigRequest
	GetMemberAccountSyncFlag() *bool
	SetPinnedTime(v int64) *SaveCustomizeReportConfigRequest
	GetPinnedTime() *int64
	SetRecipients(v string) *SaveCustomizeReportConfigRequest
	GetRecipients() *string
	SetReportDays(v int32) *SaveCustomizeReportConfigRequest
	GetReportDays() *int32
	SetReportEndDate(v string) *SaveCustomizeReportConfigRequest
	GetReportEndDate() *string
	SetReportId(v int64) *SaveCustomizeReportConfigRequest
	GetReportId() *int64
	SetReportLang(v string) *SaveCustomizeReportConfigRequest
	GetReportLang() *string
	SetReportSendType(v int32) *SaveCustomizeReportConfigRequest
	GetReportSendType() *int32
	SetReportStartDate(v string) *SaveCustomizeReportConfigRequest
	GetReportStartDate() *string
	SetReportStatus(v int32) *SaveCustomizeReportConfigRequest
	GetReportStatus() *int32
	SetReportType(v int32) *SaveCustomizeReportConfigRequest
	GetReportType() *int32
	SetReportVersion(v string) *SaveCustomizeReportConfigRequest
	GetReportVersion() *string
	SetResourceDirectoryAccountId(v int64) *SaveCustomizeReportConfigRequest
	GetResourceDirectoryAccountId() *int64
	SetSendEndTime(v string) *SaveCustomizeReportConfigRequest
	GetSendEndTime() *string
	SetSendPeriodDays(v int32) *SaveCustomizeReportConfigRequest
	GetSendPeriodDays() *int32
	SetSendPeriodType(v string) *SaveCustomizeReportConfigRequest
	GetSendPeriodType() *string
	SetSendStartTime(v string) *SaveCustomizeReportConfigRequest
	GetSendStartTime() *string
	SetTargetGroups(v string) *SaveCustomizeReportConfigRequest
	GetTargetGroups() *string
	SetTargetUids(v string) *SaveCustomizeReportConfigRequest
	GetTargetUids() *string
	SetTitle(v string) *SaveCustomizeReportConfigRequest
	GetTitle() *string
}

type SaveCustomizeReportConfigRequest struct {
	// The group type. Valid values:
	//
	// - **ALIYUN_RG**: Alibaba Cloud resource group.
	//
	// - **SAS_GROUP**: Security Center group.
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// SAS_GROUP
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// Specifies whether newly added accounts are included by default. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// true
	MemberAccountSyncFlag *bool `json:"MemberAccountSyncFlag,omitempty" xml:"MemberAccountSyncFlag,omitempty"`
	// The pinned time. Unit: milliseconds.
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 1717430400000
	PinnedTime *int64 `json:"PinnedTime,omitempty" xml:"PinnedTime,omitempty"`
	// The email addresses of contacts. Separate multiple email addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx@163.com
	Recipients *string `json:"Recipients,omitempty" xml:"Recipients,omitempty"`
	// The number of recent days for report statistics.
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 7
	ReportDays *int32 `json:"ReportDays,omitempty" xml:"ReportDays,omitempty"`
	// The end date for report statistics. Format: yyyy-MM-dd.
	//
	// > This parameter is required when ReportType is set to 3.
	//
	// example:
	//
	// 2024-01-15
	ReportEndDate *string `json:"ReportEndDate,omitempty" xml:"ReportEndDate,omitempty"`
	// The report ID.
	//
	// >Call [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) to obtain this parameter.
	//
	// example:
	//
	// 123
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The language of the report. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	ReportLang *string `json:"ReportLang,omitempty" xml:"ReportLang,omitempty"`
	// The report sending type. Valid values:
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
	ReportSendType *int32 `json:"ReportSendType,omitempty" xml:"ReportSendType,omitempty"`
	// The start date for report statistics. Format: yyyy-MM-dd.
	//
	// > This parameter is required when ReportType is set to 3.
	//
	// example:
	//
	// 2024-01-01
	ReportStartDate *string `json:"ReportStartDate,omitempty" xml:"ReportStartDate,omitempty"`
	// The report status. Valid values:
	//
	//  - **0**: disabled.
	//
	//  - **1**: enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReportStatus *int32 `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The report type. Valid values:
	//
	// - **0**: daily report.
	//
	// - **1**: weekly report.
	//
	// - **2**: monthly report.
	//
	// - **3**: custom period.
	//
	// - **4**: latest period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	ReportType *int32 `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The security report version. Valid values:
	//
	// - **1.0.0**
	//
	// - **2.0.0**
	//
	// example:
	//
	// 2.0.0
	ReportVersion *string `json:"ReportVersion,omitempty" xml:"ReportVersion,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource folder.
	//
	// >Invoke [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The send end time. Format: HH:mm:ss.
	//
	// > This parameter is required when ReportType is set to 0, 1, 2, or 4.
	//
	// example:
	//
	// 10:00:00
	SendEndTime *string `json:"SendEndTime,omitempty" xml:"SendEndTime,omitempty"`
	// The specific execution dates within the send period.
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 12
	SendPeriodDays *int32 `json:"SendPeriodDays,omitempty" xml:"SendPeriodDays,omitempty"`
	// The send period type. Valid values:
	//
	// - **DAY**: day.
	//
	// - **WEEK**: week.
	//
	// - **MONTH**: month.
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// MONTH
	SendPeriodType *string `json:"SendPeriodType,omitempty" xml:"SendPeriodType,omitempty"`
	// The send start time. Format: HH:mm:ss.
	//
	// > This parameter is required when ReportType is set to 0, 1, 2, or 4.
	//
	// example:
	//
	// 09:00:00
	SendStartTime *string `json:"SendStartTime,omitempty" xml:"SendStartTime,omitempty"`
	// The targets within the group.
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 12,123
	TargetGroups *string `json:"TargetGroups,omitempty" xml:"TargetGroups,omitempty"`
	// The list of target users. Separate multiple values with commas (,).
	//
	// > This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 12,123
	TargetUids *string `json:"TargetUids,omitempty" xml:"TargetUids,omitempty"`
	// The report name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Daily Report
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SaveCustomizeReportConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveCustomizeReportConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveCustomizeReportConfigRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *SaveCustomizeReportConfigRequest) GetMemberAccountSyncFlag() *bool {
	return s.MemberAccountSyncFlag
}

func (s *SaveCustomizeReportConfigRequest) GetPinnedTime() *int64 {
	return s.PinnedTime
}

func (s *SaveCustomizeReportConfigRequest) GetRecipients() *string {
	return s.Recipients
}

func (s *SaveCustomizeReportConfigRequest) GetReportDays() *int32 {
	return s.ReportDays
}

func (s *SaveCustomizeReportConfigRequest) GetReportEndDate() *string {
	return s.ReportEndDate
}

func (s *SaveCustomizeReportConfigRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *SaveCustomizeReportConfigRequest) GetReportLang() *string {
	return s.ReportLang
}

func (s *SaveCustomizeReportConfigRequest) GetReportSendType() *int32 {
	return s.ReportSendType
}

func (s *SaveCustomizeReportConfigRequest) GetReportStartDate() *string {
	return s.ReportStartDate
}

func (s *SaveCustomizeReportConfigRequest) GetReportStatus() *int32 {
	return s.ReportStatus
}

func (s *SaveCustomizeReportConfigRequest) GetReportType() *int32 {
	return s.ReportType
}

func (s *SaveCustomizeReportConfigRequest) GetReportVersion() *string {
	return s.ReportVersion
}

func (s *SaveCustomizeReportConfigRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *SaveCustomizeReportConfigRequest) GetSendEndTime() *string {
	return s.SendEndTime
}

func (s *SaveCustomizeReportConfigRequest) GetSendPeriodDays() *int32 {
	return s.SendPeriodDays
}

func (s *SaveCustomizeReportConfigRequest) GetSendPeriodType() *string {
	return s.SendPeriodType
}

func (s *SaveCustomizeReportConfigRequest) GetSendStartTime() *string {
	return s.SendStartTime
}

func (s *SaveCustomizeReportConfigRequest) GetTargetGroups() *string {
	return s.TargetGroups
}

func (s *SaveCustomizeReportConfigRequest) GetTargetUids() *string {
	return s.TargetUids
}

func (s *SaveCustomizeReportConfigRequest) GetTitle() *string {
	return s.Title
}

func (s *SaveCustomizeReportConfigRequest) SetGroupType(v string) *SaveCustomizeReportConfigRequest {
	s.GroupType = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetMemberAccountSyncFlag(v bool) *SaveCustomizeReportConfigRequest {
	s.MemberAccountSyncFlag = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetPinnedTime(v int64) *SaveCustomizeReportConfigRequest {
	s.PinnedTime = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetRecipients(v string) *SaveCustomizeReportConfigRequest {
	s.Recipients = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportDays(v int32) *SaveCustomizeReportConfigRequest {
	s.ReportDays = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportEndDate(v string) *SaveCustomizeReportConfigRequest {
	s.ReportEndDate = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportId(v int64) *SaveCustomizeReportConfigRequest {
	s.ReportId = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportLang(v string) *SaveCustomizeReportConfigRequest {
	s.ReportLang = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportSendType(v int32) *SaveCustomizeReportConfigRequest {
	s.ReportSendType = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportStartDate(v string) *SaveCustomizeReportConfigRequest {
	s.ReportStartDate = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportStatus(v int32) *SaveCustomizeReportConfigRequest {
	s.ReportStatus = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportType(v int32) *SaveCustomizeReportConfigRequest {
	s.ReportType = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetReportVersion(v string) *SaveCustomizeReportConfigRequest {
	s.ReportVersion = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetResourceDirectoryAccountId(v int64) *SaveCustomizeReportConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetSendEndTime(v string) *SaveCustomizeReportConfigRequest {
	s.SendEndTime = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetSendPeriodDays(v int32) *SaveCustomizeReportConfigRequest {
	s.SendPeriodDays = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetSendPeriodType(v string) *SaveCustomizeReportConfigRequest {
	s.SendPeriodType = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetSendStartTime(v string) *SaveCustomizeReportConfigRequest {
	s.SendStartTime = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetTargetGroups(v string) *SaveCustomizeReportConfigRequest {
	s.TargetGroups = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetTargetUids(v string) *SaveCustomizeReportConfigRequest {
	s.TargetUids = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) SetTitle(v string) *SaveCustomizeReportConfigRequest {
	s.Title = &v
	return s
}

func (s *SaveCustomizeReportConfigRequest) Validate() error {
	return dara.Validate(s)
}
