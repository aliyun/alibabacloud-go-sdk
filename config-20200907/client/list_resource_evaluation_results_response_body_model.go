// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationResults(v *ListResourceEvaluationResultsResponseBodyEvaluationResults) *ListResourceEvaluationResultsResponseBody
	GetEvaluationResults() *ListResourceEvaluationResultsResponseBodyEvaluationResults
	SetRequestId(v string) *ListResourceEvaluationResultsResponseBody
	GetRequestId() *string
}

type ListResourceEvaluationResultsResponseBody struct {
	// The resource evaluation results.
	EvaluationResults *ListResourceEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 25C89DDB-BB79-487D-88C3-4A561F21EFC4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBody) GetEvaluationResults() *ListResourceEvaluationResultsResponseBodyEvaluationResults {
	return s.EvaluationResults
}

func (s *ListResourceEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceEvaluationResultsResponseBody) SetEvaluationResults(v *ListResourceEvaluationResultsResponseBodyEvaluationResults) *ListResourceEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBody) SetRequestId(v string) *ListResourceEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBody) Validate() error {
	if s.EvaluationResults != nil {
		if err := s.EvaluationResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceEvaluationResultsResponseBodyEvaluationResults struct {
	// The list of resource evaluation results.
	EvaluationResultList []*ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResults) String() string {
	return dara.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) GetEvaluationResultList() []*ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	return s.EvaluationResultList
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *ListResourceEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) SetMaxResults(v int32) *ListResourceEvaluationResultsResponseBodyEvaluationResults {
	s.MaxResults = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) SetNextToken(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResults {
	s.NextToken = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) Validate() error {
	if s.EvaluationResultList != nil {
		for _, item := range s.EvaluationResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	// Additional information about the non-compliant resource. The value of this parameter can contain the following information:
	//
	// - `configuration`: The current configuration of the resource, which is the non-compliant configuration.
	//
	// - `desiredValue`: The expected configuration of the resource, which is the compliant configuration.
	//
	// - `operator`: The comparison operator that is used to compare the current configuration with the expected configuration.
	//
	// - `property`: The JSON path of the current configuration in the resource property struct.
	//
	// - `reason`: The reason why the resource is non-compliant.
	//
	// example:
	//
	// {\\"configuration\\":\\"false\\",\\"desiredValue\\":\\"True\\",\\"operator\\":\\"StringEquals\\",\\"property\\":\\"$.LoginProfile.MFABindRequired\\"}
	Annotation *string `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// - COMPLIANT: The resource is compliant.
	//
	// - NON_COMPLIANT: The resource is non-compliant.
	//
	// - NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// - INSUFFICIENT_DATA: The data is insufficient.
	//
	// - IGNORED: The evaluation result is ignored.
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The timestamp when the rule was triggered for evaluation. Unit: milliseconds.
	//
	// example:
	//
	// 1624932227157
	ConfigRuleInvokedTimestamp *int64 `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	// The unique ID of the evaluation result.
	//
	// example:
	//
	// 00000089-4e0d-58b5-a96a-8e54112110f3
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The identifier of the resource evaluation result.
	EvaluationResultIdentifier *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
	// The trigger type of the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule is triggered by a configuration change.
	//
	// - ScheduledNotification: The rule is triggered periodically.
	//
	// - Manual: The rule is triggered manually.
	//
	// example:
	//
	// ScheduledNotification
	InvokingEventMessageType *string `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	// The start time of the last non-compliance.
	//
	// example:
	//
	// 1744696393000
	LastNonCompliantRecordTimestamp *int64 `json:"LastNonCompliantRecordTimestamp,omitempty" xml:"LastNonCompliantRecordTimestamp,omitempty"`
	// Indicates whether the remediation setting is enabled. Valid values:
	//
	// - true: The remediation setting is enabled.
	//
	// - false: The remediation setting is not enabled.
	//
	// example:
	//
	// true
	RemediationEnabled *bool `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	// The timestamp when the resource evaluation result was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1624932227595
	ResultRecordedTimestamp *int64 `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: high risk.
	//
	// - 2: medium risk.
	//
	// - 3: low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetAnnotation() *string {
	return s.Annotation
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetConfigRuleInvokedTimestamp() *int64 {
	return s.ConfigRuleInvokedTimestamp
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetEvaluationId() *string {
	return s.EvaluationId
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetEvaluationResultIdentifier() *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	return s.EvaluationResultIdentifier
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetInvokingEventMessageType() *string {
	return s.InvokingEventMessageType
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetLastNonCompliantRecordTimestamp() *int64 {
	return s.LastNonCompliantRecordTimestamp
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetRemediationEnabled() *bool {
	return s.RemediationEnabled
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetResultRecordedTimestamp() *int64 {
	return s.ResultRecordedTimestamp
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationId(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetLastNonCompliantRecordTimestamp(v int64) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.LastNonCompliantRecordTimestamp = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) Validate() error {
	if s.EvaluationResultIdentifier != nil {
		if err := s.EvaluationResultIdentifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	// The resource information in the evaluation result.
	EvaluationResultQualifier *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	// The timestamp that is displayed on the timeline. Unit: milliseconds.
	//
	// example:
	//
	// 1624932227157
	OrderingTimestamp *int64 `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return dara.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GetEvaluationResultQualifier() *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	return s.EvaluationResultQualifier
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GetOrderingTimestamp() *int64 {
	return s.OrderingTimestamp
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) Validate() error {
	if s.EvaluationResultQualifier != nil {
		if err := s.EvaluationResultQualifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	// The ARN of the rule.
	//
	// example:
	//
	// acs:config::100931896542****:rule/cr-7f7d626622af0041****
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cr-7f7d626622af0041****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// ram-user-mfa-check
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The date when the ignored evaluation result is automatically resumed.
	//
	// > If this parameter is empty, the evaluation result is not automatically resumed. You must manually resume the result.
	//
	// example:
	//
	// 2022-06-01
	IgnoreDate *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	// The ID of the region where the resource resides.
	//
	// example:
	//
	// global
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 23642660635396****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// Alice
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::RAM::User
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return dara.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetIgnoreDate(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.IgnoreDate = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) Validate() error {
	return dara.Validate(s)
}
