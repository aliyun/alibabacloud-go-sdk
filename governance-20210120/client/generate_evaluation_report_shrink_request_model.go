// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateEvaluationReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *GenerateEvaluationReportShrinkRequest
	GetAccountId() *int64
	SetAccountIdsShrink(v string) *GenerateEvaluationReportShrinkRequest
	GetAccountIdsShrink() *string
	SetEvaluationDomain(v string) *GenerateEvaluationReportShrinkRequest
	GetEvaluationDomain() *string
	SetRegionId(v string) *GenerateEvaluationReportShrinkRequest
	GetRegionId() *string
	SetReportType(v string) *GenerateEvaluationReportShrinkRequest
	GetReportType() *string
}

type GenerateEvaluationReportShrinkRequest struct {
	// The account ID. If this parameter is not specified, the report is generated for the current account by default. A management account (MA) can pass in a member account ID to generate a report for the member account.
	//
	// example:
	//
	// 103144549568****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The list of member account IDs for which to generate reports.
	//
	// Note: This parameter is required only when you generate a multi-account report and want to specify the scope of accounts.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	EvaluationDomain *string `json:"EvaluationDomain,omitempty" xml:"EvaluationDomain,omitempty"`
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

func (s GenerateEvaluationReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateEvaluationReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateEvaluationReportShrinkRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GenerateEvaluationReportShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *GenerateEvaluationReportShrinkRequest) GetEvaluationDomain() *string {
	return s.EvaluationDomain
}

func (s *GenerateEvaluationReportShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateEvaluationReportShrinkRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GenerateEvaluationReportShrinkRequest) SetAccountId(v int64) *GenerateEvaluationReportShrinkRequest {
	s.AccountId = &v
	return s
}

func (s *GenerateEvaluationReportShrinkRequest) SetAccountIdsShrink(v string) *GenerateEvaluationReportShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *GenerateEvaluationReportShrinkRequest) SetEvaluationDomain(v string) *GenerateEvaluationReportShrinkRequest {
	s.EvaluationDomain = &v
	return s
}

func (s *GenerateEvaluationReportShrinkRequest) SetRegionId(v string) *GenerateEvaluationReportShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateEvaluationReportShrinkRequest) SetReportType(v string) *GenerateEvaluationReportShrinkRequest {
	s.ReportType = &v
	return s
}

func (s *GenerateEvaluationReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
