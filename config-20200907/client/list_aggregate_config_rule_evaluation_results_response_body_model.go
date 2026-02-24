// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRuleEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationResults(v *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) *ListAggregateConfigRuleEvaluationResultsResponseBody
	GetEvaluationResults() *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults
	SetRequestId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBody
	GetRequestId() *string
}

type ListAggregateConfigRuleEvaluationResultsResponseBody struct {
	// The collection of evaluation results.
	EvaluationResults *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A6662516-D056-4325-B6A7-CD3E89C97C39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBody) GetEvaluationResults() *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	return s.EvaluationResults
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBody) SetEvaluationResults(v *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) *ListAggregateConfigRuleEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBody) SetRequestId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBody) Validate() error {
	if s.EvaluationResults != nil {
		if err := s.EvaluationResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults struct {
	// A list of evaluation results.
	EvaluationResultList []*ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) GetEvaluationResultList() []*ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	return s.EvaluationResultList
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetMaxResults(v int32) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetNextToken(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.NextToken = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) Validate() error {
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

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	// The annotation of the non-compliant resource. The annotation may include the following information:
	//
	// - `configuration`: the current configuration of the resource, which is the non-compliant configuration.
	//
	// - `desiredValue`: the expected configuration of the resource, which is the compliant configuration.
	//
	// - `operator`: the comparison operator that is used to compare the current configuration with the expected configuration.
	//
	// - `property`: the JSON path of the current configuration in the resource property struct.
	//
	// - `reason`: the reason why the resource is non-compliant.
	//
	// example:
	//
	// {\\"configuration\\":\\"LRS\\",\\"desiredValue\\":\\"ZRS\\",\\"operator\\":\\"StringEquals\\",\\"property\\":\\"$.DataRedundancyType\\"}
	Annotation *string `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// - COMPLIANT: The resource is compliant.
	//
	// - NON_COMPLIANT: The resource is non-compliant.
	//
	// - NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// - INSUFFICIENT_DATA: No data is available.
	//
	// - IGNORED: The evaluation result is ignored.
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The timestamp when the rule was triggered to evaluate the resource. Unit: milliseconds.
	//
	// example:
	//
	// 1624869012713
	ConfigRuleInvokedTimestamp *int64 `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	// The unique ID of the evaluation result.
	//
	// example:
	//
	// 00000089-4e0d-58b5-a96a-8e54112110f3
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The identifier of the evaluation result.
	EvaluationResultIdentifier *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
	// The trigger type of the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule is triggered by a configuration change.
	//
	// - ScheduledNotification: The rule is triggered periodically.
	//
	// example:
	//
	// ScheduledNotification
	InvokingEventMessageType *string `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	// The timestamp when the resource was last remediated to a compliant state. This parameter is not returned if a new resource or rule is evaluated as compliant for the first time.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1768788515723
	LastCompliantFixedTimestamp *int64 `json:"LastCompliantFixedTimestamp,omitempty" xml:"LastCompliantFixedTimestamp,omitempty"`
	// The timestamp when the resource last became non-compliant.
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
	// false
	RemediationEnabled *bool `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	// The timestamp when the evaluation result was recorded. Unit: milliseconds.
	//
	// example:
	//
	// 1624869013065
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

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetAnnotation() *string {
	return s.Annotation
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetConfigRuleInvokedTimestamp() *int64 {
	return s.ConfigRuleInvokedTimestamp
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetEvaluationId() *string {
	return s.EvaluationId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetEvaluationResultIdentifier() *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	return s.EvaluationResultIdentifier
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetInvokingEventMessageType() *string {
	return s.InvokingEventMessageType
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetLastCompliantFixedTimestamp() *int64 {
	return s.LastCompliantFixedTimestamp
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetLastNonCompliantRecordTimestamp() *int64 {
	return s.LastNonCompliantRecordTimestamp
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetRemediationEnabled() *bool {
	return s.RemediationEnabled
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetResultRecordedTimestamp() *int64 {
	return s.ResultRecordedTimestamp
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetLastCompliantFixedTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.LastCompliantFixedTimestamp = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetLastNonCompliantRecordTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.LastNonCompliantRecordTimestamp = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) Validate() error {
	if s.EvaluationResultIdentifier != nil {
		if err := s.EvaluationResultIdentifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	// The information about the resource that is evaluated.
	EvaluationResultQualifier *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	// The timestamp that is displayed on the timeline. Unit: milliseconds.
	//
	// > This is the timestamp when the rule was triggered to evaluate the resource. It is the same as the value of the `ConfigRuleInvokedTimestamp` parameter.
	//
	// example:
	//
	// 1624869012713
	OrderingTimestamp *int64 `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GetEvaluationResultQualifier() *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	return s.EvaluationResultQualifier
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GetOrderingTimestamp() *int64 {
	return s.OrderingTimestamp
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) Validate() error {
	if s.EvaluationResultQualifier != nil {
		if err := s.EvaluationResultQualifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	// The ID of the compliance package to which the rule belongs.
	//
	// example:
	//
	// cr-7263fd26622af00bc****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the rule.
	//
	// example:
	//
	// acs:config::100931896542****:rule/cr-888f626622af00ae****
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cr-888f626622af00ae****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// OSS存储空间开启同城冗余存储
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The date when the ignored evaluation result is automatically restored.
	//
	// > If this parameter is empty, the result is not automatically restored. You must manually restore it.
	//
	// example:
	//
	// 2022-06-01
	IgnoreDate *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	// The ID of the region to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// rg-acfm26cicib****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// Bucket-test
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// Bucket-test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 173808452267****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetCompliancePackId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetIgnoreDate(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.IgnoreDate = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceGroupId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceOwnerId(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) Validate() error {
	return dara.Validate(s)
}
