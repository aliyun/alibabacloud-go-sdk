// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginBillingCycle(v string) *CreateReportDefinitionRequest
	GetBeginBillingCycle() *string
	SetIncludeMembers(v string) *CreateReportDefinitionRequest
	GetIncludeMembers() *string
	SetMcProject(v string) *CreateReportDefinitionRequest
	GetMcProject() *string
	SetMcTableName(v string) *CreateReportDefinitionRequest
	GetMcTableName() *string
	SetNbid(v string) *CreateReportDefinitionRequest
	GetNbid() *string
	SetNotSendOnNoData(v string) *CreateReportDefinitionRequest
	GetNotSendOnNoData() *string
	SetOssBucketName(v string) *CreateReportDefinitionRequest
	GetOssBucketName() *string
	SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionRequest
	GetOssBucketOwnerAccountId() *int64
	SetOssBucketPath(v string) *CreateReportDefinitionRequest
	GetOssBucketPath() *string
	SetReportSourceType(v string) *CreateReportDefinitionRequest
	GetReportSourceType() *string
	SetReportType(v string) *CreateReportDefinitionRequest
	GetReportType() *string
	SetSendWithAttach(v string) *CreateReportDefinitionRequest
	GetSendWithAttach() *string
	SetSplitFileOnUserId(v string) *CreateReportDefinitionRequest
	GetSplitFileOnUserId() *string
}

type CreateReportDefinitionRequest struct {
	// The start billing cycle for push. After the subscription succeeds, the system automatically pushes data from the start billing cycle to the current time. This parameter is invalid for monthly bill PDF subscriptions, and historical data will not be re-pushed. You can push data within the last year.
	//
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// The email subscription configuration that specifies whether to include multi-account members of the organization in bills.
	//
	// example:
	//
	// true
	IncludeMembers *string `json:"IncludeMembers,omitempty" xml:"IncludeMembers,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// project
	McProject *string `json:"McProject,omitempty" xml:"McProject,omitempty"`
	// The name of the MaxCompute subscription table.
	//
	// example:
	//
	// table
	McTableName *string `json:"McTableName,omitempty" xml:"McTableName,omitempty"`
	// The first-level site ID. If this parameter is left empty, the site ID of the current user is used by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// The email subscription configuration that specifies whether to skip sending emails when no bills are available.
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
	// The UID of the OSS bucket owner that stores the files. If this is a Bid/Reseller subscription and you need to push data to a sub-account\\"s OSS, specify this parameter. The account must be a sub-account of the calling account, and the AliyunConsumeDump2OSSRole permission must be granted to this account. Regular users do not need to specify this parameter. The default value is the calling account.
	//
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// The storage path of the OSS bucket.
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
	// - consumeDetailBillV2: consumption details. This value is supported only by OSS/MC subscriptions.
	//
	// - splitDetailBillV2: split details. This value is supported only by OSS/MC subscriptions.
	//
	// - costDetailBillV2: cost details. This value is supported only by OSS/MC subscriptions.
	//
	// - monthBillOverview: monthly bill overview. This value is supported only by OSS/MSC_EMAIL subscriptions.
	//
	// - focus: FOCUS bill. This value is supported only by OSS/MC subscriptions.
	//
	// This parameter is required.
	//
	// example:
	//
	// consumeDetailBillV2
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The email subscription configuration that specifies whether to send emails with bill attachments.
	//
	// example:
	//
	// true
	SendWithAttach *string `json:"SendWithAttach,omitempty" xml:"SendWithAttach,omitempty"`
	// The email subscription configuration that specifies whether to split attachments by user ID.
	//
	// example:
	//
	// true
	SplitFileOnUserId *string `json:"SplitFileOnUserId,omitempty" xml:"SplitFileOnUserId,omitempty"`
}

func (s CreateReportDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReportDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateReportDefinitionRequest) GetBeginBillingCycle() *string {
	return s.BeginBillingCycle
}

func (s *CreateReportDefinitionRequest) GetIncludeMembers() *string {
	return s.IncludeMembers
}

func (s *CreateReportDefinitionRequest) GetMcProject() *string {
	return s.McProject
}

func (s *CreateReportDefinitionRequest) GetMcTableName() *string {
	return s.McTableName
}

func (s *CreateReportDefinitionRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateReportDefinitionRequest) GetNotSendOnNoData() *string {
	return s.NotSendOnNoData
}

func (s *CreateReportDefinitionRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateReportDefinitionRequest) GetOssBucketOwnerAccountId() *int64 {
	return s.OssBucketOwnerAccountId
}

func (s *CreateReportDefinitionRequest) GetOssBucketPath() *string {
	return s.OssBucketPath
}

func (s *CreateReportDefinitionRequest) GetReportSourceType() *string {
	return s.ReportSourceType
}

func (s *CreateReportDefinitionRequest) GetReportType() *string {
	return s.ReportType
}

func (s *CreateReportDefinitionRequest) GetSendWithAttach() *string {
	return s.SendWithAttach
}

func (s *CreateReportDefinitionRequest) GetSplitFileOnUserId() *string {
	return s.SplitFileOnUserId
}

func (s *CreateReportDefinitionRequest) SetBeginBillingCycle(v string) *CreateReportDefinitionRequest {
	s.BeginBillingCycle = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetIncludeMembers(v string) *CreateReportDefinitionRequest {
	s.IncludeMembers = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetMcProject(v string) *CreateReportDefinitionRequest {
	s.McProject = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetMcTableName(v string) *CreateReportDefinitionRequest {
	s.McTableName = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetNbid(v string) *CreateReportDefinitionRequest {
	s.Nbid = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetNotSendOnNoData(v string) *CreateReportDefinitionRequest {
	s.NotSendOnNoData = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetOssBucketName(v string) *CreateReportDefinitionRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetOssBucketOwnerAccountId(v int64) *CreateReportDefinitionRequest {
	s.OssBucketOwnerAccountId = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetOssBucketPath(v string) *CreateReportDefinitionRequest {
	s.OssBucketPath = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetReportSourceType(v string) *CreateReportDefinitionRequest {
	s.ReportSourceType = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetReportType(v string) *CreateReportDefinitionRequest {
	s.ReportType = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetSendWithAttach(v string) *CreateReportDefinitionRequest {
	s.SendWithAttach = &v
	return s
}

func (s *CreateReportDefinitionRequest) SetSplitFileOnUserId(v string) *CreateReportDefinitionRequest {
	s.SplitFileOnUserId = &v
	return s
}

func (s *CreateReportDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
