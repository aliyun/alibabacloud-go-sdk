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
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-6b4a626622af0012****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-6b7c626622af00b4****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The execution mode of the remediation template. Valid values:
	//
	// 	- NON_EXECUTION: The remediation template is not executed.
	//
	// 	- AUTO_EXECUTION: The remediation template is automatically executed.
	//
	// 	- MANUAL_EXECUTION: The remediation template is manually executed.
	//
	// 	- NOT_CONFIG: The execution mode is not specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// MANUAL_EXECUTION
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// The configuration of the remediation template.
	//
	// For more information about how to obtain the configuration of the remediation template, see [ListRemediationTemplates](https://help.aliyun.com/document_detail/416781.html). You can view the `TemplateDefinition` response parameter to obtain the configuration of the remediation template.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"bucketName": "{resourceId}", "regionId": "{regionId}", "permissionName": "private"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The ID of the remediation template.
	//
	// 	- If you set the `RemediationType` parameter to `OOS`, set this parameter to the identifier of the relevant official remediation template, such as `ACS-OSS-PutBucketAcl`. For more information about how to obtain the remediation template identifier, see [ListRemediationTemplates](https://help.aliyun.com/document_detail/416781.html).
	//
	// 	- If you set the `RemediationType` parameter to `FC`, set this parameter to the Alibaba Cloud Resource Name (ARN) of the relevant Function Compute resource, such as `acs:fc:cn-hangzhou:100931896542****:services/ConfigService.LATEST/functions/test-php`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS-OSS-PutBucketAcl
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	// The type of the remediation template. Valid values:
	//
	// 	- OOS: stands for Operation Orchestration Service and indicates official remediation.
	//
	// 	- FC: stands for Function Compute and indicates custom remediation.
	//
	// This parameter is required.
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The source of remediation template. Valid values:
	//
	// 	- ALIYUN (default): official template.
	//
	// 	- CUSTOM: custom template.
	//
	// 	- NONE: none.
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
