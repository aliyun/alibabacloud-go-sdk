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
	SetRegionId(v string) *GenerateEvaluationReportRequest
	GetRegionId() *string
	SetReportType(v string) *GenerateEvaluationReportRequest
	GetReportType() *string
}

type GenerateEvaluationReportRequest struct {
	// example:
	//
	// 103144549568****
	AccountId  *int64   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
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
