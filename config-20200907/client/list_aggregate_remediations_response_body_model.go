// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRemediationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediations(v []*ListAggregateRemediationsResponseBodyRemediations) *ListAggregateRemediationsResponseBody
	GetRemediations() []*ListAggregateRemediationsResponseBodyRemediations
	SetRequestId(v string) *ListAggregateRemediationsResponseBody
	GetRequestId() *string
}

type ListAggregateRemediationsResponseBody struct {
	// The remediation settings.
	Remediations []*ListAggregateRemediationsResponseBodyRemediations `json:"Remediations,omitempty" xml:"Remediations,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0146963A-20C0-4E75-B93A-7D622B5FD7C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateRemediationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsResponseBody) GetRemediations() []*ListAggregateRemediationsResponseBodyRemediations {
	return s.Remediations
}

func (s *ListAggregateRemediationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateRemediationsResponseBody) SetRemediations(v []*ListAggregateRemediationsResponseBodyRemediations) *ListAggregateRemediationsResponseBody {
	s.Remediations = v
	return s
}

func (s *ListAggregateRemediationsResponseBody) SetRequestId(v string) *ListAggregateRemediationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBody) Validate() error {
	if s.Remediations != nil {
		for _, item := range s.Remediations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateRemediationsResponseBodyRemediations struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the account group.
	//
	// example:
	//
	// ca-6b4a626622af0012****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cr-6b7c626622af00b4****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The execution mode of the remediation. Valid values:
	//
	// - NON_EXECUTION: The remediation is not executed.
	//
	// - AUTO_EXECUTION: The remediation is automatically executed.
	//
	// - MANUAL_EXECUTION: The remediation is manually executed.
	//
	// - NOT_CONFIG: The remediation is not configured.
	//
	// example:
	//
	// AUTO_EXECUTION
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// The ID of the last successful remediation.
	//
	// example:
	//
	// bd7629fb-cac8-42fe-bcb1-e362c5a6****
	LastSuccessfulInvocationId *string `json:"LastSuccessfulInvocationId,omitempty" xml:"LastSuccessfulInvocationId,omitempty"`
	// The timestamp of the last successful remediation. Unit: milliseconds.
	//
	// example:
	//
	// 1625451393589
	LastSuccessfulInvocationTime *int64 `json:"LastSuccessfulInvocationTime,omitempty" xml:"LastSuccessfulInvocationTime,omitempty"`
	// The mode of the last successful remediation. Valid values:
	//
	// - NON_EXECUTION: The remediation was not executed.
	//
	// - AUTO_EXECUTION: The remediation was automatically executed.
	//
	// - MANUAL_EXECUTION: The remediation was manually executed.
	//
	// - NOT_CONFIG: The remediation was not configured.
	//
	// example:
	//
	// AUTO_EXECUTION
	LastSuccessfulInvocationType *string `json:"LastSuccessfulInvocationType,omitempty" xml:"LastSuccessfulInvocationType,omitempty"`
	// The converted format of the remediation setting parameters. This parameter is used only to convert the parameters of an OOS template.
	//
	// example:
	//
	// {"bucketName": "{resourceId}", "regionId": "{regionId}", "permissionName": "private"}
	RemediaitonOriginParams *string `json:"RemediaitonOriginParams,omitempty" xml:"RemediaitonOriginParams,omitempty"`
	// The ID of the remediation setting.
	//
	// example:
	//
	// crr-6b7c626622af0026****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// The source of the remediation template. Valid values:
	//
	// - ALIYUN: official template.
	//
	// - CUSTOM: custom template.
	//
	// - NONE: none.
	//
	// example:
	//
	// ALIYUN
	RemediationSourceType *string `json:"RemediationSourceType,omitempty" xml:"RemediationSourceType,omitempty"`
	// The ID of the remediation template.
	//
	// example:
	//
	// ACS-OSS-PutBucketAcl
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	// The type of the remediation. Valid values:
	//
	// - OOS: Operation Orchestration Service (official remediation).
	//
	// - FC: Function Compute (custom remediation).
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListAggregateRemediationsResponseBodyRemediations) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationsResponseBodyRemediations) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetInvokeType() *string {
	return s.InvokeType
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetLastSuccessfulInvocationId() *string {
	return s.LastSuccessfulInvocationId
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetLastSuccessfulInvocationTime() *int64 {
	return s.LastSuccessfulInvocationTime
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetLastSuccessfulInvocationType() *string {
	return s.LastSuccessfulInvocationType
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetRemediaitonOriginParams() *string {
	return s.RemediaitonOriginParams
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetRemediationId() *string {
	return s.RemediationId
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetRemediationSourceType() *string {
	return s.RemediationSourceType
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetRemediationTemplateId() *string {
	return s.RemediationTemplateId
}

func (s *ListAggregateRemediationsResponseBodyRemediations) GetRemediationType() *string {
	return s.RemediationType
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetAccountId(v int64) *ListAggregateRemediationsResponseBodyRemediations {
	s.AccountId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetAggregatorId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetConfigRuleId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetInvokeType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.InvokeType = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationTime(v int64) *ListAggregateRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationTime = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationType = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediaitonOriginParams(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediaitonOriginParams = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationSourceType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationSourceType = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationTemplateId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationTemplateId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationType = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) Validate() error {
	return dara.Validate(s)
}
