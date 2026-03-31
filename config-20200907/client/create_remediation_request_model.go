// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRemediationRequest
	GetClientToken() *string
	SetConfigRuleId(v string) *CreateRemediationRequest
	GetConfigRuleId() *string
	SetInvokeType(v string) *CreateRemediationRequest
	GetInvokeType() *string
	SetParams(v string) *CreateRemediationRequest
	GetParams() *string
	SetRemediationTemplateId(v string) *CreateRemediationRequest
	GetRemediationTemplateId() *string
	SetRemediationType(v string) *CreateRemediationRequest
	GetRemediationType() *string
	SetSourceType(v string) *CreateRemediationRequest
	GetSourceType() *string
}

type CreateRemediationRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-8a973ac2e2be00a2****
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
	// For more information about how to obtain the remediation template configuration, see the `TemplateDefinition` response parameter provided in [ListRemediationTemplates](https://help.aliyun.com/document_detail/416781.html).
	//
	// This parameter is required.
	//
	// if can be null:
	// true
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
	// 	- OOS: Operation Orchestration Service (official remediation)
	//
	// 	- FC: Function Compute (custom remediation)
	//
	// This parameter is required.
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The source of remediation. Valid values:
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

func (s CreateRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRemediationRequest) GoString() string {
	return s.String()
}

func (s *CreateRemediationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRemediationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *CreateRemediationRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *CreateRemediationRequest) GetParams() *string {
	return s.Params
}

func (s *CreateRemediationRequest) GetRemediationTemplateId() *string {
	return s.RemediationTemplateId
}

func (s *CreateRemediationRequest) GetRemediationType() *string {
	return s.RemediationType
}

func (s *CreateRemediationRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRemediationRequest) SetClientToken(v string) *CreateRemediationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRemediationRequest) SetConfigRuleId(v string) *CreateRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateRemediationRequest) SetInvokeType(v string) *CreateRemediationRequest {
	s.InvokeType = &v
	return s
}

func (s *CreateRemediationRequest) SetParams(v string) *CreateRemediationRequest {
	s.Params = &v
	return s
}

func (s *CreateRemediationRequest) SetRemediationTemplateId(v string) *CreateRemediationRequest {
	s.RemediationTemplateId = &v
	return s
}

func (s *CreateRemediationRequest) SetRemediationType(v string) *CreateRemediationRequest {
	s.RemediationType = &v
	return s
}

func (s *CreateRemediationRequest) SetSourceType(v string) *CreateRemediationRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRemediationRequest) Validate() error {
	return dara.Validate(s)
}
