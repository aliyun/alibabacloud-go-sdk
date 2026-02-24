// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRemediationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediation(v *DescribeRemediationResponseBodyRemediation) *DescribeRemediationResponseBody
	GetRemediation() *DescribeRemediationResponseBodyRemediation
	SetRequestId(v string) *DescribeRemediationResponseBody
	GetRequestId() *string
}

type DescribeRemediationResponseBody struct {
	Remediation *DescribeRemediationResponseBodyRemediation `json:"Remediation,omitempty" xml:"Remediation,omitempty" type:"Struct"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRemediationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRemediationResponseBody) GetRemediation() *DescribeRemediationResponseBodyRemediation {
	return s.Remediation
}

func (s *DescribeRemediationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRemediationResponseBody) SetRemediation(v *DescribeRemediationResponseBodyRemediation) *DescribeRemediationResponseBody {
	s.Remediation = v
	return s
}

func (s *DescribeRemediationResponseBody) SetRequestId(v string) *DescribeRemediationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRemediationResponseBody) Validate() error {
	if s.Remediation != nil {
		if err := s.Remediation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRemediationResponseBodyRemediation struct {
	AccountId                    *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ConfigRuleId                 *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	InvokeType                   *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	LastSuccessfulInvocationId   *string `json:"LastSuccessfulInvocationId,omitempty" xml:"LastSuccessfulInvocationId,omitempty"`
	LastSuccessfulInvocationTime *int64  `json:"LastSuccessfulInvocationTime,omitempty" xml:"LastSuccessfulInvocationTime,omitempty"`
	LastSuccessfulInvocationType *string `json:"LastSuccessfulInvocationType,omitempty" xml:"LastSuccessfulInvocationType,omitempty"`
	RemediationId                *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// This parameter is required.
	RemediationOriginParams *string `json:"RemediationOriginParams,omitempty" xml:"RemediationOriginParams,omitempty"`
	RemediationSourceType   *string `json:"RemediationSourceType,omitempty" xml:"RemediationSourceType,omitempty"`
	RemediationTemplateId   *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	RemediationType         *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s DescribeRemediationResponseBodyRemediation) String() string {
	return dara.Prettify(s)
}

func (s DescribeRemediationResponseBodyRemediation) GoString() string {
	return s.String()
}

func (s *DescribeRemediationResponseBodyRemediation) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DescribeRemediationResponseBodyRemediation) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DescribeRemediationResponseBodyRemediation) GetInvokeType() *string {
	return s.InvokeType
}

func (s *DescribeRemediationResponseBodyRemediation) GetLastSuccessfulInvocationId() *string {
	return s.LastSuccessfulInvocationId
}

func (s *DescribeRemediationResponseBodyRemediation) GetLastSuccessfulInvocationTime() *int64 {
	return s.LastSuccessfulInvocationTime
}

func (s *DescribeRemediationResponseBodyRemediation) GetLastSuccessfulInvocationType() *string {
	return s.LastSuccessfulInvocationType
}

func (s *DescribeRemediationResponseBodyRemediation) GetRemediationId() *string {
	return s.RemediationId
}

func (s *DescribeRemediationResponseBodyRemediation) GetRemediationOriginParams() *string {
	return s.RemediationOriginParams
}

func (s *DescribeRemediationResponseBodyRemediation) GetRemediationSourceType() *string {
	return s.RemediationSourceType
}

func (s *DescribeRemediationResponseBodyRemediation) GetRemediationTemplateId() *string {
	return s.RemediationTemplateId
}

func (s *DescribeRemediationResponseBodyRemediation) GetRemediationType() *string {
	return s.RemediationType
}

func (s *DescribeRemediationResponseBodyRemediation) SetAccountId(v int64) *DescribeRemediationResponseBodyRemediation {
	s.AccountId = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetConfigRuleId(v string) *DescribeRemediationResponseBodyRemediation {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetInvokeType(v string) *DescribeRemediationResponseBodyRemediation {
	s.InvokeType = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetLastSuccessfulInvocationId(v string) *DescribeRemediationResponseBodyRemediation {
	s.LastSuccessfulInvocationId = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetLastSuccessfulInvocationTime(v int64) *DescribeRemediationResponseBodyRemediation {
	s.LastSuccessfulInvocationTime = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetLastSuccessfulInvocationType(v string) *DescribeRemediationResponseBodyRemediation {
	s.LastSuccessfulInvocationType = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetRemediationId(v string) *DescribeRemediationResponseBodyRemediation {
	s.RemediationId = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetRemediationOriginParams(v string) *DescribeRemediationResponseBodyRemediation {
	s.RemediationOriginParams = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetRemediationSourceType(v string) *DescribeRemediationResponseBodyRemediation {
	s.RemediationSourceType = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetRemediationTemplateId(v string) *DescribeRemediationResponseBodyRemediation {
	s.RemediationTemplateId = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) SetRemediationType(v string) *DescribeRemediationResponseBodyRemediation {
	s.RemediationType = &v
	return s
}

func (s *DescribeRemediationResponseBodyRemediation) Validate() error {
	return dara.Validate(s)
}
