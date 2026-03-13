// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateEvaluationReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *GenerateEvaluationReportResponseBody
	GetAccountId() *int64
	SetEvaluationScore(v float64) *GenerateEvaluationReportResponseBody
	GetEvaluationScore() *float64
	SetEvaluationTime(v string) *GenerateEvaluationReportResponseBody
	GetEvaluationTime() *string
	SetFinished(v string) *GenerateEvaluationReportResponseBody
	GetFinished() *string
	SetReportType(v string) *GenerateEvaluationReportResponseBody
	GetReportType() *string
	SetReportUrl(v string) *GenerateEvaluationReportResponseBody
	GetReportUrl() *string
	SetRequestId(v string) *GenerateEvaluationReportResponseBody
	GetRequestId() *string
}

type GenerateEvaluationReportResponseBody struct {
	// example:
	//
	// 103144549568****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 0.7684
	EvaluationScore *float64 `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty"`
	// example:
	//
	// 2026-01-12T07:25:33Z
	EvaluationTime *string `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// example:
	//
	// true
	Finished *string `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// EvaluationAccountExcelReport
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// example:
	//
	// https://governance-prod-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/reports-html/*****
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7DCF863F-CBBB-57C4-8AF2-5D4EE35D1EB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateEvaluationReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateEvaluationReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateEvaluationReportResponseBody) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GenerateEvaluationReportResponseBody) GetEvaluationScore() *float64 {
	return s.EvaluationScore
}

func (s *GenerateEvaluationReportResponseBody) GetEvaluationTime() *string {
	return s.EvaluationTime
}

func (s *GenerateEvaluationReportResponseBody) GetFinished() *string {
	return s.Finished
}

func (s *GenerateEvaluationReportResponseBody) GetReportType() *string {
	return s.ReportType
}

func (s *GenerateEvaluationReportResponseBody) GetReportUrl() *string {
	return s.ReportUrl
}

func (s *GenerateEvaluationReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateEvaluationReportResponseBody) SetAccountId(v int64) *GenerateEvaluationReportResponseBody {
	s.AccountId = &v
	return s
}

func (s *GenerateEvaluationReportResponseBody) SetEvaluationScore(v float64) *GenerateEvaluationReportResponseBody {
	s.EvaluationScore = &v
	return s
}

func (s *GenerateEvaluationReportResponseBody) SetEvaluationTime(v string) *GenerateEvaluationReportResponseBody {
	s.EvaluationTime = &v
	return s
}

func (s *GenerateEvaluationReportResponseBody) SetFinished(v string) *GenerateEvaluationReportResponseBody {
	s.Finished = &v
	return s
}

func (s *GenerateEvaluationReportResponseBody) SetReportType(v string) *GenerateEvaluationReportResponseBody {
	s.ReportType = &v
	return s
}

func (s *GenerateEvaluationReportResponseBody) SetReportUrl(v string) *GenerateEvaluationReportResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *GenerateEvaluationReportResponseBody) SetRequestId(v string) *GenerateEvaluationReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateEvaluationReportResponseBody) Validate() error {
	return dara.Validate(s)
}
