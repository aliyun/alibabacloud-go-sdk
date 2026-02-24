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
	// A client token to ensure the idempotence of the request. The token must be unique for each request. The `ClientToken` parameter contains only ASCII characters and must not exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAdDWBF2****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The rule ID.
	//
	// For more information, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-8a973ac2e2be00a2****
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
	// The remediation parameters.
	//
	// For more information, see the `TemplateDefinition` parameter in [ListRemediationTemplates](https://help.aliyun.com/document_detail/416781.html).
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
	// - If `RemediationType` is set to `OOS`, set this parameter to `ACS-OSS-PutBucketAcl`. For more information, see [ListRemediationTemplates](https://help.aliyun.com/document_detail/416781.html).
	//
	// - If `RemediationType` is set to `FC`, set this parameter to the Alibaba Cloud Resource Name (ARN) of the function in Function Compute. Example: `acs:fc:cn-hangzhou:100931896542****:services/ConfigService.LATEST/functions/test-php`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS-OSS-PutBucketAcl
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	// The type of the remediation. Valid values:
	//
	// - OOS: template-based remediation using OOS.
	//
	// - FC: custom remediation using FC.
	//
	// This parameter is required.
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The source of the remediation template. Valid values:
	//
	// - ALIYUN (default): an official template.
	//
	// - CUSTOM: a custom template. This value is required for custom FC remediations.
	//
	// - NONE: no source.
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
