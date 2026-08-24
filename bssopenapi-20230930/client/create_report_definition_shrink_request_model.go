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
	// The start billing cycle for push. After successful subscription, the system automatically pushes data from the start billing cycle to the current time. This parameter is invalid for monthly bill PDF subscriptions and does not re-push historical data. Data within the last year can be pushed.
	//
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// Email subscription configuration. Specifies whether to include enterprise multi-account members in the bill.
	//
	// example:
	//
	// true
	IncludeMembers *string `json:"IncludeMembers,omitempty" xml:"IncludeMembers,omitempty"`
	// The MaxCompute project name.
	//
	// example:
	//
	// project
	McProject *string `json:"McProject,omitempty" xml:"McProject,omitempty"`
	// The MaxCompute subscription table name.
	//
	// example:
	//
	// table
	McTableName *string `json:"McTableName,omitempty" xml:"McTableName,omitempty"`
	// The primary sales site ID. If left empty, the system uses the site ID of the current user by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// Email subscription configuration. Specifies whether to skip sending emails when no bills are available.
	//
	// example:
	//
	// true
	NotSendOnNoData *string `json:"NotSendOnNoData,omitempty" xml:"NotSendOnNoData,omitempty"`
	// The name of the OSS bucket for file storage.
	//
	// example:
	//
	// sh-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The UID of the OSS owner that stores the files. If this is a Bid/Reseller subscription and you need to push to a sub-account\\"s OSS, specify this parameter. The account must be a sub-account of the calling account, and the AliyunConsumeDump2OSSRole permission must be granted to this account. Regular users do not need to specify this parameter. The default value is the calling account.
	//
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// The OSS bucket storage path.
	//
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// The subscription source. Valid values: OSS, MC, or MSC_EMAIL.
	//
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// The subscription type. Valid values:
	//
	// - consumeDetailBillV2: consumption details (supported only for OSS/MC subscriptions).
	//
	// - splitDetailBillV2: split details (supported only for OSS/MC subscriptions).
	//
	// - costDetailBillV2: cost details (supported only for OSS/MC subscriptions).
	//
	// - monthBillOverview: monthly bill summary (supported only for OSS/MSC_EMAIL subscriptions).
	//
	// - focus: FOCUS bill (supported only for OSS/MC subscriptions).
	//
	// This parameter is required.
	//
	// example:
	//
	// consumeDetailBillV2
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The user-specified subscription fields.
	SelectedFieldsShrink *string `json:"SelectedFields,omitempty" xml:"SelectedFields,omitempty"`
	// Email subscription configuration. Specifies whether to include bill attachments in emails.
	//
	// example:
	//
	// true
	SendWithAttach *string `json:"SendWithAttach,omitempty" xml:"SendWithAttach,omitempty"`
	// Email subscription configuration. Specifies whether to split attachments by user ID.
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
