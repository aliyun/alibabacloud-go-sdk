// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportDefinitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginBillingCycle(v string) *CreateReportDefinitionShrinkRequest
	GetBeginBillingCycle() *string
	SetIncludeMembers(v string) *CreateReportDefinitionShrinkRequest
	GetIncludeMembers() *string
	SetMcProject(v string) *CreateReportDefinitionShrinkRequest
	GetMcProject() *string
	SetMcTableName(v string) *CreateReportDefinitionShrinkRequest
	GetMcTableName() *string
	SetNbid(v string) *CreateReportDefinitionShrinkRequest
	GetNbid() *string
	SetNotSendOnNoData(v string) *CreateReportDefinitionShrinkRequest
	GetNotSendOnNoData() *string
	SetOssBucketName(v string) *CreateReportDefinitionShrinkRequest
	GetOssBucketName() *string
	SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionShrinkRequest
	GetOssBucketOwnerAccountId() *int64
	SetOssBucketPath(v string) *CreateReportDefinitionShrinkRequest
	GetOssBucketPath() *string
	SetReportSourceType(v string) *CreateReportDefinitionShrinkRequest
	GetReportSourceType() *string
	SetReportType(v string) *CreateReportDefinitionShrinkRequest
	GetReportType() *string
	SetSelectedFieldsShrink(v string) *CreateReportDefinitionShrinkRequest
	GetSelectedFieldsShrink() *string
	SetSendWithAttach(v string) *CreateReportDefinitionShrinkRequest
	GetSendWithAttach() *string
	SetSplitFileOnUserId(v string) *CreateReportDefinitionShrinkRequest
	GetSplitFileOnUserId() *string
}

type CreateReportDefinitionShrinkRequest struct {
	// The start billing cycle for data push. After the subscription is created, the system automatically pushes data from the start billing cycle to the current time. This parameter is invalid for monthly billing PDF subscriptions, and historical data is not pushed again. Data within the most recent year can be pushed.
	//
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// Specifies whether to include enterprise multi-account members in bills for email subscriptions.
	//
	// example:
	//
	// true
	IncludeMembers *string `json:"IncludeMembers,omitempty" xml:"IncludeMembers,omitempty"`
	// The name of the MC project.
	//
	// example:
	//
	// project
	McProject *string `json:"McProject,omitempty" xml:"McProject,omitempty"`
	// The name of the MC subscription table.
	//
	// example:
	//
	// table
	McTableName *string `json:"McTableName,omitempty" xml:"McTableName,omitempty"`
	// The ID of the level-1 site. If this parameter is left empty, the ID of the site to which the current user belongs is used by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// Specifies whether to skip sending emails when no bills are generated for email subscriptions.
	//
	// example:
	//
	// true
	NotSendOnNoData *string `json:"NotSendOnNoData,omitempty" xml:"NotSendOnNoData,omitempty"`
	// The name of the OSS bucket that stores the files.
	//
	// example:
	//
	// sh-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The UID of the OSS owner that stores the files. Specify this parameter if the subscription is made by a bidder or reseller and the files need to be pushed to the OSS bucket of a RAM user. The RAM user must be a subordinate account of the calling account and must be granted the AliyunConsumeDump2OSSRole permission. Regular users do not need to specify this parameter. The default value is the calling account.
	//
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// The storage path in the OSS bucket.
	//
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// The subscription source. Valid values: OSS, MC, and MSC_EMAIL.
	//
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// The subscription type. Valid values:
	//
	// - consumeDetailBillV2: consumption details. Supported only for OSS and MC subscriptions.
	//
	// - splitDetailBillV2: split details. Supported only for OSS and MC subscriptions.
	//
	// - costDetailBillV2: cost details. Supported only for OSS and MC subscriptions.
	//
	// - monthBillOverview: monthly bill overview. Supported only for OSS and MSC_EMAIL subscriptions.
	//
	// - focus: focus bills. Supported only for OSS and MC subscriptions.
	//
	// This parameter is required.
	//
	// example:
	//
	// consumeDetailBillV2
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The fields to which the user subscribes.
	SelectedFieldsShrink *string `json:"SelectedFields,omitempty" xml:"SelectedFields,omitempty"`
	// Specifies whether to include bill attachments in emails for email subscriptions.
	//
	// example:
	//
	// true
	SendWithAttach *string `json:"SendWithAttach,omitempty" xml:"SendWithAttach,omitempty"`
	// Specifies whether to split attachments by user ID for email subscriptions.
	//
	// example:
	//
	// true
	SplitFileOnUserId *string `json:"SplitFileOnUserId,omitempty" xml:"SplitFileOnUserId,omitempty"`
}

func (s CreateReportDefinitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReportDefinitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionShrinkRequest) GetBeginBillingCycle() *string {
	return s.BeginBillingCycle
}

func (s *CreateReportDefinitionShrinkRequest) GetIncludeMembers() *string {
	return s.IncludeMembers
}

func (s *CreateReportDefinitionShrinkRequest) GetMcProject() *string {
	return s.McProject
}

func (s *CreateReportDefinitionShrinkRequest) GetMcTableName() *string {
	return s.McTableName
}

func (s *CreateReportDefinitionShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateReportDefinitionShrinkRequest) GetNotSendOnNoData() *string {
	return s.NotSendOnNoData
}

func (s *CreateReportDefinitionShrinkRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateReportDefinitionShrinkRequest) GetOssBucketOwnerAccountId() *int64 {
	return s.OssBucketOwnerAccountId
}

func (s *CreateReportDefinitionShrinkRequest) GetOssBucketPath() *string {
	return s.OssBucketPath
}

func (s *CreateReportDefinitionShrinkRequest) GetReportSourceType() *string {
	return s.ReportSourceType
}

func (s *CreateReportDefinitionShrinkRequest) GetReportType() *string {
	return s.ReportType
}

func (s *CreateReportDefinitionShrinkRequest) GetSelectedFieldsShrink() *string {
	return s.SelectedFieldsShrink
}

func (s *CreateReportDefinitionShrinkRequest) GetSendWithAttach() *string {
	return s.SendWithAttach
}

func (s *CreateReportDefinitionShrinkRequest) GetSplitFileOnUserId() *string {
	return s.SplitFileOnUserId
}

func (s *CreateReportDefinitionShrinkRequest) SetBeginBillingCycle(v string) *CreateReportDefinitionShrinkRequest {
	s.BeginBillingCycle = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetIncludeMembers(v string) *CreateReportDefinitionShrinkRequest {
	s.IncludeMembers = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetMcProject(v string) *CreateReportDefinitionShrinkRequest {
	s.McProject = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetMcTableName(v string) *CreateReportDefinitionShrinkRequest {
	s.McTableName = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetNbid(v string) *CreateReportDefinitionShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetNotSendOnNoData(v string) *CreateReportDefinitionShrinkRequest {
	s.NotSendOnNoData = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetOssBucketName(v string) *CreateReportDefinitionShrinkRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionShrinkRequest {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetOssBucketPath(v string) *CreateReportDefinitionShrinkRequest {
	s.OssBucketPath = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetReportSourceType(v string) *CreateReportDefinitionShrinkRequest {
	s.ReportSourceType = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetReportType(v string) *CreateReportDefinitionShrinkRequest {
	s.ReportType = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetSelectedFieldsShrink(v string) *CreateReportDefinitionShrinkRequest {
	s.SelectedFieldsShrink = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetSendWithAttach(v string) *CreateReportDefinitionShrinkRequest {
	s.SendWithAttach = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) SetSplitFileOnUserId(v string) *CreateReportDefinitionShrinkRequest {
	s.SplitFileOnUserId = &v
	return s
}

func (s *CreateReportDefinitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
