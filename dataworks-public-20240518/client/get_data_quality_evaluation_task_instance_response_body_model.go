// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityEvaluationTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTaskInstance(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) *GetDataQualityEvaluationTaskInstanceResponseBody
	GetDataQualityEvaluationTaskInstance() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance
	SetRequestId(v string) *GetDataQualityEvaluationTaskInstanceResponseBody
	GetRequestId() *string
}

type GetDataQualityEvaluationTaskInstanceResponseBody struct {
	// The details of the monitor instance.
	DataQualityEvaluationTaskInstance *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance `json:"DataQualityEvaluationTaskInstance,omitempty" xml:"DataQualityEvaluationTaskInstance,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBody) GetDataQualityEvaluationTaskInstance() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	return s.DataQualityEvaluationTaskInstance
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBody) SetDataQualityEvaluationTaskInstance(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) *GetDataQualityEvaluationTaskInstanceResponseBody {
	s.DataQualityEvaluationTaskInstance = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBody) SetRequestId(v string) *GetDataQualityEvaluationTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBody) Validate() error {
	if s.DataQualityEvaluationTaskInstance != nil {
		if err := s.DataQualityEvaluationTaskInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance struct {
	// The creation time.
	//
	// example:
	//
	// 1716344665000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end time of the instance.
	//
	// example:
	//
	// 1716344665000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the data quality monitoring instance.
	//
	// example:
	//
	// 7234231689
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Data quality verification execution parameters in JSON format. The available keys are as follows:
	//
	// - triggerTime: the millisecond timestamp of the trigger time. The baseline time of the $[yyyymmdd] expression in the data range of data quality monitoring. Required.
	//
	// example:
	//
	// { "triggerTime": 1733284062000 }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 98330
	ProjectId *int64                                                                                      `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Results   []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The status of the data quality monitoring instance.
	//
	// - Running: Verifying
	//
	// - Error: A rule verification Error occurred.
	//
	// - Passed: all rules are verified
	//
	// - Warned: normal alarm threshold triggered by rules
	//
	// - Critical: Threshold for serious alerts triggered by rules
	//
	// example:
	//
	// Passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The monitor.
	Task *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// The context information when the instance is triggered, in JSON format. The possible keys are as follows:
	//
	// - TriggerClient: the trigger source of the data quality monitoring instance, such as CWF2 (scheduling system), may be added later.
	//
	// - TriggerClientId: associated with a specific business resource in the source system. For example, if TriggerClient is CWF2, the ID of the scheduling task is recorded here.
	//
	// example:
	//
	// { "triggerClient": "CWF2", "triggerClientId": 70001238945 }
	TriggerContext *string `json:"TriggerContext,omitempty" xml:"TriggerContext,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetParameters() *string {
	return s.Parameters
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetResults() []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	return s.Results
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetStatus() *string {
	return s.Status
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetTask() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	return s.Task
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) GetTriggerContext() *string {
	return s.TriggerContext
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetCreateTime(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.CreateTime = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetFinishTime(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.FinishTime = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.Id = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetParameters(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.Parameters = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetProjectId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetResults(v []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.Results = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetStatus(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.Status = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetTask(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.Task = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) SetTriggerContext(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance {
	s.TriggerContext = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstance) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults struct {
	CreateTime     *int64                                                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Details        []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	Id             *int64                                                                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Rule           *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule      `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	Sample         *string                                                                                            `json:"Sample,omitempty" xml:"Sample,omitempty"`
	Status         *string                                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskInstanceId *int64                                                                                             `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GetDetails() []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails {
	return s.Details
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GetRule() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	return s.Rule
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GetSample() *string {
	return s.Sample
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GetStatus() *string {
	return s.Status
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) GetTaskInstanceId() *int64 {
	return s.TaskInstanceId
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) SetCreateTime(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	s.CreateTime = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) SetDetails(v []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	s.Details = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) SetId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	s.Id = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) SetRule(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	s.Rule = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) SetSample(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	s.Sample = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) SetStatus(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	s.Status = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) SetTaskInstanceId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults {
	s.TaskInstanceId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResults) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails struct {
	CheckedValue    *string `json:"CheckedValue,omitempty" xml:"CheckedValue,omitempty"`
	ReferencedValue *string `json:"ReferencedValue,omitempty" xml:"ReferencedValue,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) GetCheckedValue() *string {
	return s.CheckedValue
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) GetReferencedValue() *string {
	return s.ReferencedValue
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) GetStatus() *string {
	return s.Status
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) SetCheckedValue(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails {
	s.CheckedValue = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) SetReferencedValue(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails {
	s.ReferencedValue = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) SetStatus(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails {
	s.Status = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsDetails) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule struct {
	CheckingConfig *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig  `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	Description    *string                                                                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled        *bool                                                                                                        `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ErrorHandlers  []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	Id             *int64                                                                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Name           *string                                                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId      *int64                                                                                                       `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig  `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	Severity       *string                                                                                                      `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Target         *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget          `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	TemplateCode   *string                                                                                                      `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetCheckingConfig() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig {
	return s.CheckingConfig
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetDescription() *string {
	return s.Description
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetErrorHandlers() []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers {
	return s.ErrorHandlers
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetName() *string {
	return s.Name
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetSamplingConfig() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig {
	return s.SamplingConfig
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetSeverity() *string {
	return s.Severity
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetTarget() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget {
	return s.Target
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetCheckingConfig(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.CheckingConfig = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetDescription(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.Description = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetEnabled(v bool) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.Enabled = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetErrorHandlers(v []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.ErrorHandlers = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.Id = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetName(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.Name = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetProjectId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetSamplingConfig(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.SamplingConfig = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetSeverity(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.Severity = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetTarget(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.Target = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) SetTemplateCode(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule {
	s.TemplateCode = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRule) Validate() error {
	if s.CheckingConfig != nil {
		if err := s.CheckingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorHandlers != nil {
		for _, item := range s.ErrorHandlers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SamplingConfig != nil {
		if err := s.SamplingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig struct {
	ReferencedSamplesFilter *string                                                                                                               `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	Thresholds              *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	Type                    *string                                                                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) GetThresholds() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds {
	return s.Thresholds
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) SetReferencedSamplesFilter(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) SetThresholds(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) SetType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfig) Validate() error {
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds struct {
	Critical *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Expected *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	Warned   *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned   `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) GetCritical() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) GetExpected() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) GetWarned() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) SetCritical(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) SetExpected(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) SetWarned(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholds) Validate() error {
	if s.Critical != nil {
		if err := s.Critical.Validate(); err != nil {
			return err
		}
	}
	if s.Expected != nil {
		if err := s.Expected.Validate(); err != nil {
			return err
		}
	}
	if s.Warned != nil {
		if err := s.Warned.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical struct {
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) SetExpression(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) SetOperator(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) SetValue(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected struct {
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) SetExpression(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) SetOperator(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) SetValue(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned struct {
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) SetExpression(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) SetOperator(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) SetValue(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers struct {
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) SetErrorDataFilter(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) SetType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig struct {
	Metric           *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	SamplingFilter   *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	SettingConfig    *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) SetMetric(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) SetMetricParameters(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) SetSamplingFilter(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) SetSettingConfig(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget struct {
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	TableGuid    *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) SetDatabaseType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) SetTableGuid(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) SetType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceResultsRuleTarget) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask struct {
	// The description of the monitor.
	//
	// example:
	//
	// OpenAPI quality monitoring test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook.
	Hooks []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 28544990
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitor.
	//
	// example:
	//
	// Data quality OpenAPI monitoring test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of alert notifications.
	Notifications *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// The ID of the workspace.
	//
	// example:
	//
	// 20629
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Extended configuration, JSON-formatted string, takes effect only for EMR-type data quality monitoring.
	//
	// - queue: the yarn queue used when performing EMR data quality verification. The default queue is the queue configured for this project.
	//
	// - sqlEngine: SQL engine used when performing EMR data verification
	//
	//   - HIVE_ SQL
	//
	//   - SPARK_ SQL
	//
	// example:
	//
	// { "queue": "default" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitored object of the monitor.
	Target *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the monitor.
	Trigger *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetDescription() *string {
	return s.Description
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetHooks() []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks {
	return s.Hooks
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetName() *string {
	return s.Name
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetNotifications() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications {
	return s.Notifications
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetTarget() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget {
	return s.Target
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) GetTrigger() *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger {
	return s.Trigger
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetDescription(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.Description = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetHooks(v []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.Hooks = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.Id = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetName(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.Name = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetNotifications(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.Notifications = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetProjectId(v int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetRuntimeConf(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.RuntimeConf = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetTarget(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.Target = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) SetTrigger(v *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask {
	s.Trigger = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTask) Validate() error {
	if s.Hooks != nil {
		for _, item := range s.Hooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Notifications != nil {
		if err := s.Notifications.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks struct {
	// The hook trigger condition. When this condition is met, the hook action is triggered. Only two conditional expressions are supported:
	//
	// 	- Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical.
	//
	// 	- Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// Hook type. Currently, only one type is supported:
	//
	// - BlockTaskInstance: the blocking scheduling task continues to run. Data quality monitoring is triggered by the scheduling task. After the data quality monitoring is completed, the Hook.Condition is used to determine whether the blocking scheduling task continues to run.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) GetCondition() *string {
	return s.Condition
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) SetCondition(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks {
	s.Condition = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) SetType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskHooks) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications struct {
	// The notification trigger condition. When this condition is met, the alert notification is triggered. Only two conditional expressions are supported:
	//
	// 	- Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical.
	//
	// 	- Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The alert notification methods.
	Notifications []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) GetCondition() *string {
	return s.Condition
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) GetNotifications() []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	return s.Notifications
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) SetCondition(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications {
	s.Condition = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) SetNotifications(v []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications {
	s.Notifications = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotifications) Validate() error {
	if s.Notifications != nil {
		for _, item := range s.Notifications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications struct {
	// The notification method.
	NotificationChannels []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The value of the receiver.
	NotificationReceivers []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) GetNotificationChannels() []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) GetNotificationReceivers() []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) SetNotificationChannels(v []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) SetNotificationReceivers(v []*GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotifications) Validate() error {
	if s.NotificationChannels != nil {
		for _, item := range s.NotificationChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NotificationReceivers != nil {
		for _, item := range s.NotificationReceivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels struct {
	// The notification method.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers struct {
	// Additional parameter settings for sending alerts in json format. The supported keys are as follows:
	//
	// - atAll: when sending DingTalk alerts, do you need to @ everyone in the group. It takes effect when ReceiverType is DingdingUrl.
	//
	// example:
	//
	// { "atAll": true }
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of alert recipient.
	//
	// example:
	//
	// DingdingUrl
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The recipient of the alert.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) SetExtension(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) SetReceiverType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) SetReceiverValues(v []*string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget struct {
	// The type of the database to which the table belongs.
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The partition range monitored.
	//
	// example:
	//
	// pt=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The unique ID of the table in the data map.
	//
	// example:
	//
	// odps.api_trace.ods_d_api_log
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The type of the monitoring object.
	//
	// - Table: Table
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) SetDatabaseType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget {
	s.DatabaseType = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) SetPartitionSpec(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget {
	s.PartitionSpec = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) SetTableGuid(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget {
	s.TableGuid = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) SetType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTarget) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger struct {
	// The Id list of the scheduled task, which is valid when the Type is ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger type of the monitor. Valid values:
	//
	// 	- ByManual (default): The monitor is manually triggered.
	//
	// 	- ByScheduledTaskInstance: The monitor is triggered by the associated scheduling tasks.
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) SetTaskIds(v []*int64) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) SetType(v string) *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponseBodyDataQualityEvaluationTaskInstanceTaskTrigger) Validate() error {
	return dara.Validate(s)
}
