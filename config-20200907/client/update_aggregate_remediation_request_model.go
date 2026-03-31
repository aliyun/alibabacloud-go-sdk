// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *UpdateAggregateRemediationRequest
	GetAggregatorId() *string
	SetInvokeType(v string) *UpdateAggregateRemediationRequest
	GetInvokeType() *string
	SetParams(v string) *UpdateAggregateRemediationRequest
	GetParams() *string
	SetRemediationId(v string) *UpdateAggregateRemediationRequest
	GetRemediationId() *string
	SetRemediationTemplateId(v string) *UpdateAggregateRemediationRequest
	GetRemediationTemplateId() *string
	SetRemediationType(v string) *UpdateAggregateRemediationRequest
	GetRemediationType() *string
	SetSourceType(v string) *UpdateAggregateRemediationRequest
	GetSourceType() *string
}

type UpdateAggregateRemediationRequest struct {
	// The ID of the account group.
	//
	// You can the [ListAggregators](https://help.aliyun.com/document_detail/255797.html) operation to obtain the ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-6b4a626622af0012****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The execution mode of the remediation. Valid values:
	//
	// 	- NON_EXECUTION: The remediation is not executed.
	//
	// 	- AUTO_EXECUTION: The remediation is automatically executed.
	//
	// 	- MANUAL_EXECUTION: The remediation is manually executed.
	//
	// 	- NOT_CONFIG: The execution mode is not specified.
	//
	// example:
	//
	// AUTO_EXECUTION
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// The desired parameter values of the remediation setting.
	//
	// example:
	//
	// {"properties":[{"name":"regionId","type":"STRING","value":"{regionId}","allowedValues":[]},{"name":"bucketName","type":"STRING","value":"{resourceId}","allowedValues":[],"description":"OSS Bucket Name."},{"name":"permissionName","type":"STRING","value":"private","allowedValues":["public-read-write","public-read","private"],"description":"ACL Permission Name."}]}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The ID of the remediation setting.
	//
	// You can call the [ListAggregateRemediations](https://help.aliyun.com/document_detail/270036.html) operation to obtain the ID of the remediation setting.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-909ba2d4716700eb****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// The ID of the remediation template.
	//
	// You can call the [ListRemediationTemplates](https://help.aliyun.com/document_detail/270066.html) operation to obtain the ID of the remediation template.
	//
	// example:
	//
	// ACS-OSS-PutBucketAcl
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	// The type of the remediation template. Valid values:
	//
	// 	- OOS: Operation Orchestration Service (OOS)
	//
	// 	- FC: Function Compute. You can use Function Compute to configure custom remediation settings.
	//
	// example:
	//
	// OOS
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The type of the rule for which the remediation template is configured. Valid values:
	//
	// 	- ALIYUN: managed rule.
	//
	// 	- CUSTOM: custom rule.
	//
	// 	- NONE: The rule is not specified.
	//
	// example:
	//
	// ALIYUN
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s UpdateAggregateRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateRemediationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateRemediationRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregateRemediationRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *UpdateAggregateRemediationRequest) GetParams() *string {
	return s.Params
}

func (s *UpdateAggregateRemediationRequest) GetRemediationId() *string {
	return s.RemediationId
}

func (s *UpdateAggregateRemediationRequest) GetRemediationTemplateId() *string {
	return s.RemediationTemplateId
}

func (s *UpdateAggregateRemediationRequest) GetRemediationType() *string {
	return s.RemediationType
}

func (s *UpdateAggregateRemediationRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateAggregateRemediationRequest) SetAggregatorId(v string) *UpdateAggregateRemediationRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetInvokeType(v string) *UpdateAggregateRemediationRequest {
	s.InvokeType = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetParams(v string) *UpdateAggregateRemediationRequest {
	s.Params = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetRemediationId(v string) *UpdateAggregateRemediationRequest {
	s.RemediationId = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetRemediationTemplateId(v string) *UpdateAggregateRemediationRequest {
	s.RemediationTemplateId = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetRemediationType(v string) *UpdateAggregateRemediationRequest {
	s.RemediationType = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetSourceType(v string) *UpdateAggregateRemediationRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) Validate() error {
	return dara.Validate(s)
}
