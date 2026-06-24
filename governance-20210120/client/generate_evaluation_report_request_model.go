// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateEvaluationReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *GenerateEvaluationReportRequest
	GetAccountId() *int64
	SetAccountIds(v []*int64) *GenerateEvaluationReportRequest
	GetAccountIds() []*int64
	SetEvaluationDomain(v string) *GenerateEvaluationReportRequest
	GetEvaluationDomain() *string
	SetRegionId(v string) *GenerateEvaluationReportRequest
	GetRegionId() *string
	SetReportType(v string) *GenerateEvaluationReportRequest
	GetReportType() *string
}

type GenerateEvaluationReportRequest struct {
	// The account ID. If this parameter is not specified, the report is generated for the current account by default. A management account (MA) can pass in a member account ID to generate a report for the member account.
	//
	// example:
	//
	// 103144549568****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The list of member account IDs for which to generate reports.
	//
	// Note: This parameter is required only when you generate a multi-account report and want to specify the scope of accounts.
	AccountIds       []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	EvaluationDomain *string  `json:"EvaluationDomain,omitempty" xml:"EvaluationDomain,omitempty"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The report type. Valid values:
	//
	// - EvaluationAccountHtmlReport: single-account HTML report.
	//
	// - EvaluationAccountExcelReport: single-account Excel report.
	//
	// - EvaluationMultiAccountExcelReport: multi-account Excel report.
	//
	// example:
	//
	// EvaluationAccountExcelReport
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GenerateEvaluationReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateEvaluationReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateEvaluationReportRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GenerateEvaluationReportRequest) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *GenerateEvaluationReportRequest) GetEvaluationDomain() *string {
	return s.EvaluationDomain
}

func (s *GenerateEvaluationReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateEvaluationReportRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GenerateEvaluationReportRequest) SetAccountId(v int64) *GenerateEvaluationReportRequest {
	s.AccountId = &v
	return s
}

func (s *GenerateEvaluationReportRequest) SetAccountIds(v []*int64) *GenerateEvaluationReportRequest {
	s.AccountIds = v
	return s
}

func (s *GenerateEvaluationReportRequest) SetEvaluationDomain(v string) *GenerateEvaluationReportRequest {
	s.EvaluationDomain = &v
	return s
}

func (s *GenerateEvaluationReportRequest) SetRegionId(v string) *GenerateEvaluationReportRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateEvaluationReportRequest) SetReportType(v string) *GenerateEvaluationReportRequest {
	s.ReportType = &v
	return s
}

func (s *GenerateEvaluationReportRequest) Validate() error {
	return dara.Validate(s)
}
