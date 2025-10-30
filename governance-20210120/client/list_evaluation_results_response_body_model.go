// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *ListEvaluationResultsResponseBody
	GetAccountId() *int64
	SetRequestId(v string) *ListEvaluationResultsResponseBody
	GetRequestId() *string
	SetResults(v *ListEvaluationResultsResponseBodyResults) *ListEvaluationResultsResponseBody
	GetResults() *ListEvaluationResultsResponseBodyResults
}

type ListEvaluationResultsResponseBody struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD57329E-131A-59F4-8746-E1CD8D7B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The check results, including the status of the overall check and the results of check items.
	Results *ListEvaluationResultsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
}

func (s ListEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBody) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluationResultsResponseBody) GetResults() *ListEvaluationResultsResponseBodyResults {
	return s.Results
}

func (s *ListEvaluationResultsResponseBody) SetAccountId(v int64) *ListEvaluationResultsResponseBody {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationResultsResponseBody) SetRequestId(v string) *ListEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationResultsResponseBody) SetResults(v *ListEvaluationResultsResponseBodyResults) *ListEvaluationResultsResponseBody {
	s.Results = v
	return s
}

func (s *ListEvaluationResultsResponseBody) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEvaluationResultsResponseBodyResults struct {
	// The end time of the overall check. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-13T03:35:00Z
	EvaluationTime *string `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// The check result.
	MetricResults []*ListEvaluationResultsResponseBodyResultsMetricResults `json:"MetricResults,omitempty" xml:"MetricResults,omitempty" type:"Repeated"`
	// The status of the overall check. Valid values:
	//
	// 	- Running: The check is in progress.
	//
	// 	- Finished: The check is complete.
	//
	// 	- failed: The check fails.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The overall score.
	//
	// example:
	//
	// 0.6453
	TotalScore *float64 `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResults) GetEvaluationTime() *string {
	return s.EvaluationTime
}

func (s *ListEvaluationResultsResponseBodyResults) GetMetricResults() []*ListEvaluationResultsResponseBodyResultsMetricResults {
	return s.MetricResults
}

func (s *ListEvaluationResultsResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *ListEvaluationResultsResponseBodyResults) GetTotalScore() *float64 {
	return s.TotalScore
}

func (s *ListEvaluationResultsResponseBodyResults) SetEvaluationTime(v string) *ListEvaluationResultsResponseBodyResults {
	s.EvaluationTime = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResults) SetMetricResults(v []*ListEvaluationResultsResponseBodyResultsMetricResults) *ListEvaluationResultsResponseBodyResults {
	s.MetricResults = v
	return s
}

func (s *ListEvaluationResultsResponseBodyResults) SetStatus(v string) *ListEvaluationResultsResponseBodyResults {
	s.Status = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResults) SetTotalScore(v float64) *ListEvaluationResultsResponseBodyResults {
	s.TotalScore = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResults) Validate() error {
	if s.MetricResults != nil {
		for _, item := range s.MetricResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationResultsResponseBodyResultsMetricResults struct {
	AccountSummary *ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary `json:"AccountSummary,omitempty" xml:"AccountSummary,omitempty" type:"Struct"`
	// The error information.
	//
	// >  This parameter is returned only if the value of `Status` is `Failed`.
	ErrorInfo *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" type:"Struct"`
	// The end time of the check item. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-13T03:34:02Z
	EvaluationTime *string `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// r7xdcu****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0.2
	PotentialScoreIncrease *float64 `json:"PotentialScoreIncrease,omitempty" xml:"PotentialScoreIncrease,omitempty"`
	// The checked resources.
	ResourcesSummary *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary `json:"ResourcesSummary,omitempty" xml:"ResourcesSummary,omitempty" type:"Struct"`
	// The rate of the non-compliant resources.
	//
	// example:
	//
	// 0.67
	Result *float64 `json:"Result,omitempty" xml:"Result,omitempty"`
	// The risk level. Valid values:
	//
	// 	- Error: high risk
	//
	// 	- Warning: medium risk
	//
	// 	- None: no risk
	//
	// example:
	//
	// Error
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- Running: The check is in progress.
	//
	// 	- Finished: The check is complete.
	//
	// 	- failed: The check fails.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResultsMetricResults) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResultsMetricResults) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetAccountSummary() *ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary {
	return s.AccountSummary
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetErrorInfo() *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo {
	return s.ErrorInfo
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetEvaluationTime() *string {
	return s.EvaluationTime
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetId() *string {
	return s.Id
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetPotentialScoreIncrease() *float64 {
	return s.PotentialScoreIncrease
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetResourcesSummary() *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary {
	return s.ResourcesSummary
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetResult() *float64 {
	return s.Result
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetRisk() *string {
	return s.Risk
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) GetStatus() *string {
	return s.Status
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetAccountSummary(v *ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.AccountSummary = v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetErrorInfo(v *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.ErrorInfo = v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetEvaluationTime(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.EvaluationTime = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetId(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Id = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetPotentialScoreIncrease(v float64) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.PotentialScoreIncrease = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetResourcesSummary(v *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.ResourcesSummary = v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetResult(v float64) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Result = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetRisk(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Risk = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetStatus(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Status = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) Validate() error {
	if s.AccountSummary != nil {
		if err := s.AccountSummary.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorInfo != nil {
		if err := s.ErrorInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResourcesSummary != nil {
		if err := s.ResourcesSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary struct {
	// example:
	//
	// 1
	NonCompliant *int32 `json:"NonCompliant,omitempty" xml:"NonCompliant,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary) GetNonCompliant() *int32 {
	return s.NonCompliant
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary) SetNonCompliant(v int32) *ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary {
	s.NonCompliant = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsAccountSummary) Validate() error {
	return dara.Validate(s)
}

type ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo struct {
	// The error code.
	//
	// example:
	//
	// EcsInsightEnableFailed
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unable to enable ECS Insight due to a server error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) GetCode() *string {
	return s.Code
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) GetMessage() *string {
	return s.Message
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) SetCode(v string) *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo {
	s.Code = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) SetMessage(v string) *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo {
	s.Message = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) Validate() error {
	return dara.Validate(s)
}

type ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary struct {
	// The number of non-compliant resources.
	//
	// example:
	//
	// 2
	NonCompliant *int32 `json:"NonCompliant,omitempty" xml:"NonCompliant,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) GetNonCompliant() *int32 {
	return s.NonCompliant
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) SetNonCompliant(v int32) *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary {
	s.NonCompliant = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) Validate() error {
	return dara.Validate(s)
}
