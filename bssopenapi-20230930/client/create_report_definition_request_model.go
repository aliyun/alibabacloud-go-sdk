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
	SetMcProject(v string) *CreateReportDefinitionRequest
	GetMcProject() *string
	SetMcTableName(v string) *CreateReportDefinitionRequest
	GetMcTableName() *string
	SetNbid(v string) *CreateReportDefinitionRequest
	GetNbid() *string
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
}

type CreateReportDefinitionRequest struct {
	// example:
	//
	// 2025-05
	BeginBillingCycle *string `json:"BeginBillingCycle,omitempty" xml:"BeginBillingCycle,omitempty"`
	// example:
	//
	// project
	McProject *string `json:"McProject,omitempty" xml:"McProject,omitempty"`
	// example:
	//
	// table
	McTableName *string `json:"McTableName,omitempty" xml:"McTableName,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// sh-bill
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// 1234567812345678
	OssBucketOwnerAccountId *int64 `json:"OssBucketOwnerAccountId,omitempty" xml:"OssBucketOwnerAccountId,omitempty"`
	// example:
	//
	// bill/
	OssBucketPath *string `json:"OssBucketPath,omitempty" xml:"OssBucketPath,omitempty"`
	// example:
	//
	// OSS
	ReportSourceType *string `json:"ReportSourceType,omitempty" xml:"ReportSourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BillingItemDetailForBillingPeriod
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
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

func (s *CreateReportDefinitionRequest) GetMcProject() *string {
	return s.McProject
}

func (s *CreateReportDefinitionRequest) GetMcTableName() *string {
	return s.McTableName
}

func (s *CreateReportDefinitionRequest) GetNbid() *string {
	return s.Nbid
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

func (s *CreateReportDefinitionRequest) SetBeginBillingCycle(v string) *CreateReportDefinitionRequest {
	s.BeginBillingCycle = &v
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

func (s *CreateReportDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
