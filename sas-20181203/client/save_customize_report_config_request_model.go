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
	// The grouping type. Valid values:
	//
	// 	- **ALIYUN_RG**
	//
	// 	- **SAS_GROUP**
	//
	// >  This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// SAS_GROUP
	GroupType             *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	MemberAccountSyncFlag *bool   `json:"MemberAccountSyncFlag,omitempty" xml:"MemberAccountSyncFlag,omitempty"`
	// The time when the report is pinned. Unit: milliseconds.
	//
	// >  This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 1717430400000
	PinnedTime *int64 `json:"PinnedTime,omitempty" xml:"PinnedTime,omitempty"`
	// The email address of the recipient. Separate multiple email addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx@163.com
	Recipients *string `json:"Recipients,omitempty" xml:"Recipients,omitempty"`
	// The most recent days for report statistics.
	//
	// >  This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 7
	ReportDays *int32 `json:"ReportDays,omitempty" xml:"ReportDays,omitempty"`
	// The end date on which the report is sent. The value is in the yyyy-MM-dd format.
	//
	// >  This parameter is required if the ReportType parameter is set to 3.
	//
	// example:
	//
	// 2024-01-15
	ReportEndDate *string `json:"ReportEndDate,omitempty" xml:"ReportEndDate,omitempty"`
	// The ID of the report.
	//
	// >  You can call the [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) operation to query the ID.
	//
	// example:
	//
	// 123
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The language of the report. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	ReportLang *string `json:"ReportLang,omitempty" xml:"ReportLang,omitempty"`
	// The time range in which the report is sent. Valid values:
	//
	// 	- **1**: 00:00 to 06:00.
	//
	// 	- **2**: 06:00 to 12:00.
	//
	// 	- **3**: 12:00 to 18:00.
	//
	// 	- **4**: 18:00 to 24:00.
	//
	// example:
	//
	// 2
	ReportSendType *int32 `json:"ReportSendType,omitempty" xml:"ReportSendType,omitempty"`
	// The start date on which the report is sent. The value is in the yyyy-MM-dd format.
	//
	// >  This parameter is required if the ReportType parameter is set to 3.
	//
	// example:
	//
	// 2024-01-01
	ReportStartDate *string `json:"ReportStartDate,omitempty" xml:"ReportStartDate,omitempty"`
	// The status of the report. Valid values:
	//
	// 	- **0**: disabled.
	//
	// 	- **1**: enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReportStatus *int32 `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The type of the report. Valid values:
	//
	// 	- **0**: daily report.
	//
	// 	- **1**: weekly report.
	//
	// 	- **2**: monthly report.
	//
	// 	- **3**: report whose statistics are collected within a custom time range.
	//
	// 	- **4**: report of the most recent time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	ReportType *int32 `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The version of the report. Valid values:
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// example:
	//
	// 2.0.0
	ReportVersion *string `json:"ReportVersion,omitempty" xml:"ReportVersion,omitempty"`
	// The end time at which the report is sent. The value is in the HH:mm:ss format.
	//
	// >  This parameter is required if the ReportType parameter is set to 0, 1, 2, or 4.
	//
	// example:
	//
	// 10:00:00
	SendEndTime *string `json:"SendEndTime,omitempty" xml:"SendEndTime,omitempty"`
	// The exact day within the sending period.
	//
	// >  This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 12
	SendPeriodDays *int32 `json:"SendPeriodDays,omitempty" xml:"SendPeriodDays,omitempty"`
	// The interval at which the report is sent. Valid values:
	//
	// 	- **DAY**
	//
	// 	- **WEEK**
	//
	// 	- **MONTH**
	//
	// >  This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// MONTH
	SendPeriodType *string `json:"SendPeriodType,omitempty" xml:"SendPeriodType,omitempty"`
	// The start time at which the report is sent. The value is in the HH:mm:ss format.
	//
	// >  This parameter is required if the ReportType parameter is set to 0, 1, 2, or 4.
	//
	// example:
	//
	// 09:00:00
	SendStartTime *string `json:"SendStartTime,omitempty" xml:"SendStartTime,omitempty"`
	// The groups.
	//
	// >  This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 12,123
	TargetGroups *string `json:"TargetGroups,omitempty" xml:"TargetGroups,omitempty"`
	// The ID of the Alibaba Cloud account. Separate multiple IDs with commas (,).
	//
	// >  This parameter is supported only in version 2.0.0.
	//
	// example:
	//
	// 12,123
	TargetUids *string `json:"TargetUids,omitempty" xml:"TargetUids,omitempty"`
	// The title of the report.
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
