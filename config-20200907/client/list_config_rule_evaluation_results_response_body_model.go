// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRuleEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationResults(v *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) *ListConfigRuleEvaluationResultsResponseBody
	GetEvaluationResults() *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults
	SetRequestId(v string) *ListConfigRuleEvaluationResultsResponseBody
	GetRequestId() *string
}

type ListConfigRuleEvaluationResultsResponseBody struct {
	// The rule evaluation results.
	EvaluationResults *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2A4A33BD-8186-4D60-91B9-42174EED75B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBody) GetEvaluationResults() *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	return s.EvaluationResults
}

func (s *ListConfigRuleEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigRuleEvaluationResultsResponseBody) SetEvaluationResults(v *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) *ListConfigRuleEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBody) SetRequestId(v string) *ListConfigRuleEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBody) Validate() error {
	if s.EvaluationResults != nil {
		if err := s.EvaluationResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResults struct {
	// The list of rule evaluation results.
	EvaluationResultList []*ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) GetEvaluationResultList() []*ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	return s.EvaluationResultList
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetMaxResults(v int32) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.MaxResults = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetNextToken(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.NextToken = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) Validate() error {
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

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	// The supplementary information about the non-compliant resource. This may include the following information:
	//
	// - `configuration`: The current configuration of the resource, which is the non-compliant configuration.
	//
	// - `desiredValue`: The expected configuration of the resource, which is the compliant configuration.
	//
	// - `operator`: The comparison operator used to compare the current configuration with the expected configuration.
	//
	// - `property`: The JSON path of the current configuration in the resource property struct.
	//
	// - `reason`: The reason why the resource is non-compliant.
	//
	// example:
	//
	// {\\"configuration\\":\\"\\",\\"desiredValue\\":\\"\\",\\"operator\\":\\"IsNotStringEmpty\\",\\"property\\":\\"$.KeyPairName\\",\\"reason\\":\\"No property contains.\\"}
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
	// The UNIX timestamp when the rule was triggered for evaluation. Unit: milliseconds.
	//
	// example:
	//
	// 1622802307081
	ConfigRuleInvokedTimestamp *int64 `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	// The unique ID of the evaluation result.
	//
	// example:
	//
	// 00000089-4e0d-58b5-a96a-8e54112110f3
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The identifier of the rule evaluation result.
	EvaluationResultIdentifier *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
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
	// ConfigurationItemChangeNotification
	InvokingEventMessageType *string `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	// The time when the resource was last remediated to a compliant state. This value is not recorded when a new resource or rule is evaluated as compliant for the first time.
	//
	// example:
	//
	// 1768788515723
	LastCompliantFixedTimestamp *int64 `json:"LastCompliantFixedTimestamp,omitempty" xml:"LastCompliantFixedTimestamp,omitempty"`
	// The start time of the last non-compliance.
	//
	// example:
	//
	// 1744696665000
	LastNonCompliantRecordTimestamp *int64 `json:"LastNonCompliantRecordTimestamp,omitempty" xml:"LastNonCompliantRecordTimestamp,omitempty"`
	// Indicates whether the remediation setting is enabled. Valid values:
	//
	// - true: The remediation setting is enabled.
	//
	// - false: The remediation setting is disabled.
	//
	// example:
	//
	// false
	RemediationEnabled *bool `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	// The UNIX timestamp when the resource evaluation result was generated. Unit: milliseconds.
	//
	// example:
	//
	// 1622802307150
	ResultRecordedTimestamp *int64 `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetAnnotation() *string {
	return s.Annotation
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetConfigRuleInvokedTimestamp() *int64 {
	return s.ConfigRuleInvokedTimestamp
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetEvaluationId() *string {
	return s.EvaluationId
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetEvaluationResultIdentifier() *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	return s.EvaluationResultIdentifier
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetInvokingEventMessageType() *string {
	return s.InvokingEventMessageType
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetLastCompliantFixedTimestamp() *int64 {
	return s.LastCompliantFixedTimestamp
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetLastNonCompliantRecordTimestamp() *int64 {
	return s.LastNonCompliantRecordTimestamp
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetRemediationEnabled() *bool {
	return s.RemediationEnabled
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetResultRecordedTimestamp() *int64 {
	return s.ResultRecordedTimestamp
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetLastCompliantFixedTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.LastCompliantFixedTimestamp = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetLastNonCompliantRecordTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.LastNonCompliantRecordTimestamp = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) Validate() error {
	if s.EvaluationResultIdentifier != nil {
		if err := s.EvaluationResultIdentifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	// The resource information in the rule evaluation result.
	EvaluationResultQualifier *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	// The UNIX timestamp displayed on the timeline. Unit: milliseconds.
	//
	// example:
	//
	// 1622802307081
	OrderingTimestamp *int64 `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GetEvaluationResultQualifier() *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	return s.EvaluationResultQualifier
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GetOrderingTimestamp() *int64 {
	return s.OrderingTimestamp
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) Validate() error {
	if s.EvaluationResultQualifier != nil {
		if err := s.EvaluationResultQualifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	// The ID of the compliance package to which the rule belongs.
	//
	// example:
	//
	// cp-bcc33457e0d900d5****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the rule.
	//
	// example:
	//
	// acs:config::120886317861****:rule/cr-cac56457e0d900d3****
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cr-cac56457e0d900d3****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// ECS实例CPU核数满足最低要求
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The date when the ignored evaluation result is automatically resumed.
	//
	// > If this parameter is empty, the result is not automatically resumed. You must manually resume it.
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
	// rg-aek3tprgnnc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-hp3e4kvhzqn2s11t****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// iZuf6j91r34rnwawoox****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 120886317861****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetIgnoreDate() *string {
	return s.IgnoreDate
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetRegionId() *string {
	return s.RegionId
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetCompliancePackId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetIgnoreDate(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.IgnoreDate = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceGroupId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceGroupId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceOwnerId(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) Validate() error {
	return dara.Validate(s)
}
