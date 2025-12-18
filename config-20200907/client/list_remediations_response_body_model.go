// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListRemediationsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRemediationsResponseBody
	GetPageSize() *int64
	SetRemediations(v []*ListRemediationsResponseBodyRemediations) *ListRemediationsResponseBody
	GetRemediations() []*ListRemediationsResponseBodyRemediations
	SetRequestId(v string) *ListRemediationsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListRemediationsResponseBody
	GetTotalCount() *string
}

type ListRemediationsResponseBody struct {
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The converted configuration of the remediation template. This parameter is returned only for an OOS remediation template.
	Remediations []*ListRemediationsResponseBodyRemediations `json:"Remediations,omitempty" xml:"Remediations,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0146963A-20C0-4E75-B93A-7D622B5FD7C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of remediation settings.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRemediationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemediationsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRemediationsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRemediationsResponseBody) GetRemediations() []*ListRemediationsResponseBodyRemediations {
	return s.Remediations
}

func (s *ListRemediationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRemediationsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRemediationsResponseBody) SetPageNumber(v int64) *ListRemediationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRemediationsResponseBody) SetPageSize(v int64) *ListRemediationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRemediationsResponseBody) SetRemediations(v []*ListRemediationsResponseBodyRemediations) *ListRemediationsResponseBody {
	s.Remediations = v
	return s
}

func (s *ListRemediationsResponseBody) SetRequestId(v string) *ListRemediationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRemediationsResponseBody) SetTotalCount(v string) *ListRemediationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRemediationsResponseBody) Validate() error {
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

type ListRemediationsResponseBodyRemediations struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cr-6b7c626622af00b4****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The execution mode of the remediation template. Valid values:
	//
	// 	- NON_EXECUTION: The remediation template was not executed.
	//
	// 	- AUTO_EXECUTION: The remediation template was automatically executed.
	//
	// 	- MANUAL_EXECUTION: The remediation template was manually executed.
	//
	// 	- NOT_CONFIG: The execution mode was not specified.
	//
	// example:
	//
	// AUTO_EXECUTION
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// The ID of the last successful execution of the remediation template.
	//
	// example:
	//
	// bd7629fb-cac8-42fe-bcb1-e362c5a6****
	LastSuccessfulInvocationId *string `json:"LastSuccessfulInvocationId,omitempty" xml:"LastSuccessfulInvocationId,omitempty"`
	// The timestamp of the last successful execution of the remediation template Unit: milliseconds.
	//
	// example:
	//
	// 1625451393589
	LastSuccessfulInvocationTime *int64 `json:"LastSuccessfulInvocationTime,omitempty" xml:"LastSuccessfulInvocationTime,omitempty"`
	// The mode of the last successful execution of the remediation template. Valid values:
	//
	// 	- NON_EXECUTION: The remediation template was not executed.
	//
	// 	- AUTO_EXECUTION: The remediation template was automatically executed.
	//
	// 	- MANUAL_EXECUTION: The remediation template was manually executed.
	//
	// 	- NOT_CONFIG: The execution mode was not specified.
	//
	// example:
	//
	// AUTO_EXECUTION
	LastSuccessfulInvocationType *string `json:"LastSuccessfulInvocationType,omitempty" xml:"LastSuccessfulInvocationType,omitempty"`
	// The ID of the remediation template.
	//
	// example:
	//
	// crr-6b7c626622af0026****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// The converted configuration of the remediation template. This parameter is available only for an OOS remediation template.
	//
	// example:
	//
	// {"bucketName": "{resourceId}", "regionId": "{regionId}", "permissionName": "private"}
	RemediationOriginParams *string `json:"RemediationOriginParams,omitempty" xml:"RemediationOriginParams,omitempty"`
	// The source of remediation. Valid values:
	//
	// 	- ALIYUN: official template.
	//
	// 	- CUSTOM: custom template.
	//
	// 	- NONE: none.
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
	// The type of the remediation template. Valid values:
	//
	// 	- OOS: Operation Orchestration Service (official remediation)
	//
	// 	- FC: Function Compute (custom remediation)
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListRemediationsResponseBodyRemediations) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationsResponseBodyRemediations) GoString() string {
	return s.String()
}

func (s *ListRemediationsResponseBodyRemediations) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListRemediationsResponseBodyRemediations) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListRemediationsResponseBodyRemediations) GetInvokeType() *string {
	return s.InvokeType
}

func (s *ListRemediationsResponseBodyRemediations) GetLastSuccessfulInvocationId() *string {
	return s.LastSuccessfulInvocationId
}

func (s *ListRemediationsResponseBodyRemediations) GetLastSuccessfulInvocationTime() *int64 {
	return s.LastSuccessfulInvocationTime
}

func (s *ListRemediationsResponseBodyRemediations) GetLastSuccessfulInvocationType() *string {
	return s.LastSuccessfulInvocationType
}

func (s *ListRemediationsResponseBodyRemediations) GetRemediationId() *string {
	return s.RemediationId
}

func (s *ListRemediationsResponseBodyRemediations) GetRemediationOriginParams() *string {
	return s.RemediationOriginParams
}

func (s *ListRemediationsResponseBodyRemediations) GetRemediationSourceType() *string {
	return s.RemediationSourceType
}

func (s *ListRemediationsResponseBodyRemediations) GetRemediationTemplateId() *string {
	return s.RemediationTemplateId
}

func (s *ListRemediationsResponseBodyRemediations) GetRemediationType() *string {
	return s.RemediationType
}

func (s *ListRemediationsResponseBodyRemediations) SetAccountId(v int64) *ListRemediationsResponseBodyRemediations {
	s.AccountId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetConfigRuleId(v string) *ListRemediationsResponseBodyRemediations {
	s.ConfigRuleId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetInvokeType(v string) *ListRemediationsResponseBodyRemediations {
	s.InvokeType = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationId(v string) *ListRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationTime(v int64) *ListRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationTime = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationType(v string) *ListRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationType = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationId(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationOriginParams(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationOriginParams = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationSourceType(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationSourceType = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationTemplateId(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationTemplateId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationType(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationType = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) Validate() error {
	return dara.Validate(s)
}
