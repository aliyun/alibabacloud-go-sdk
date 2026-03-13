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
	SetRegionId(v string) *GenerateEvaluationReportShrinkRequest
	GetRegionId() *string
	SetReportType(v string) *GenerateEvaluationReportShrinkRequest
	GetReportType() *string
}

type GenerateEvaluationReportShrinkRequest struct {
	// example:
	//
	// 103144549568****
	AccountId        *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
