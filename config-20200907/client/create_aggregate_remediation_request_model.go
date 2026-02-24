// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *CreateAggregateRemediationRequest
	GetAggregatorId() *string
	SetClientToken(v string) *CreateAggregateRemediationRequest
	GetClientToken() *string
	SetConfigRuleId(v string) *CreateAggregateRemediationRequest
	GetConfigRuleId() *string
	SetInvokeType(v string) *CreateAggregateRemediationRequest
	GetInvokeType() *string
	SetParams(v string) *CreateAggregateRemediationRequest
	GetParams() *string
	SetRemediationTemplateId(v string) *CreateAggregateRemediationRequest
	GetRemediationTemplateId() *string
	SetRemediationType(v string) *CreateAggregateRemediationRequest
	GetRemediationType() *string
	SetSourceType(v string) *CreateAggregateRemediationRequest
	GetSourceType() *string
}

type CreateAggregateRemediationRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-6b4a626622af0012****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// A client token. It is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The `ClientToken` parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The rule ID.
	//
	// For more information about how to obtain the rule ID, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
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
	// - NOT_CONFIG: The execution mode is not set.
	//
	// This parameter is required.
	//
	// example:
	//
	// MANUAL_EXECUTION
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// The parameters of the remediation.
	//
	// For more information about how to obtain the parameters of the remediation, see the `TemplateDefinition` parameter in [ListRemediationTemplates](https://help.aliyun.com/document_detail/416781.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"bucketName": "{resourceId}", "regionId": "{regionId}", "permissionName": "private"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The remediation template ID.
	//
	// - If you set `RemediationType` to `OOS`, set this parameter to `ACS-OSS-PutBucketAcl`. For more information about how to obtain the remediation template ID, see [ListRemediationTemplates](https://help.aliyun.com/document_detail/416781.html).
	//
	// - If you set `RemediationType` to `FC`, set this parameter to the Alibaba Cloud Resource Name (ARN) of the function in Function Compute. Example: `acs:fc:cn-hangzhou:100931896542****:services/ConfigService.LATEST/functions/test-php`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS-OSS-PutBucketAcl
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	// The remediation type. Valid values:
	//
	// - OOS: OOS (template-based remediation).
	//
	// - FC: FC (custom remediation).
	//
	// This parameter is required.
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The source of the remediation template. Valid values:
	//
	// - ALIYUN (default): official template.
	//
	// - CUSTOM: custom template. This value must be specified for custom FC remediations.
	//
	// - NONE: none.
	//
	// example:
	//
	// ALIYUN
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateAggregateRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateRemediationRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateRemediationRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregateRemediationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregateRemediationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *CreateAggregateRemediationRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *CreateAggregateRemediationRequest) GetParams() *string {
	return s.Params
}

func (s *CreateAggregateRemediationRequest) GetRemediationTemplateId() *string {
	return s.RemediationTemplateId
}

func (s *CreateAggregateRemediationRequest) GetRemediationType() *string {
	return s.RemediationType
}

func (s *CreateAggregateRemediationRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateAggregateRemediationRequest) SetAggregatorId(v string) *CreateAggregateRemediationRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetClientToken(v string) *CreateAggregateRemediationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetConfigRuleId(v string) *CreateAggregateRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetInvokeType(v string) *CreateAggregateRemediationRequest {
	s.InvokeType = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetParams(v string) *CreateAggregateRemediationRequest {
	s.Params = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetRemediationTemplateId(v string) *CreateAggregateRemediationRequest {
	s.RemediationTemplateId = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetRemediationType(v string) *CreateAggregateRemediationRequest {
	s.RemediationType = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetSourceType(v string) *CreateAggregateRemediationRequest {
	s.SourceType = &v
	return s
}

func (s *CreateAggregateRemediationRequest) Validate() error {
	return dara.Validate(s)
}
