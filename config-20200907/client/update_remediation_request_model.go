// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRemediationRequest
	GetClientToken() *string
	SetInvokeType(v string) *UpdateRemediationRequest
	GetInvokeType() *string
	SetParams(v string) *UpdateRemediationRequest
	GetParams() *string
	SetRemediationId(v string) *UpdateRemediationRequest
	GetRemediationId() *string
	SetRemediationTemplateId(v string) *UpdateRemediationRequest
	GetRemediationTemplateId() *string
	SetRemediationType(v string) *UpdateRemediationRequest
	GetRemediationType() *string
	SetSourceType(v string) *UpdateRemediationRequest
	GetSourceType() *string
}

type UpdateRemediationRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 2014295338-e7361358-5822-4276-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// You can call the [ListRemediations](https://help.aliyun.com/document_detail/270772.html) operation to obtain the ID of the remediation setting.
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
	// The source of the remediation setting. Valid values:
	//
	// 	- ALIYUN: the default remediation setting of Alibaba Cloud.
	//
	// 	- CUSTOM: a custom remediation setting.
	//
	// 	- NONE: The source is not specified.
	//
	// example:
	//
	// ALIYUN
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s UpdateRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRemediationRequest) GoString() string {
	return s.String()
}

func (s *UpdateRemediationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRemediationRequest) GetInvokeType() *string {
	return s.InvokeType
}

func (s *UpdateRemediationRequest) GetParams() *string {
	return s.Params
}

func (s *UpdateRemediationRequest) GetRemediationId() *string {
	return s.RemediationId
}

func (s *UpdateRemediationRequest) GetRemediationTemplateId() *string {
	return s.RemediationTemplateId
}

func (s *UpdateRemediationRequest) GetRemediationType() *string {
	return s.RemediationType
}

func (s *UpdateRemediationRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateRemediationRequest) SetClientToken(v string) *UpdateRemediationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRemediationRequest) SetInvokeType(v string) *UpdateRemediationRequest {
	s.InvokeType = &v
	return s
}

func (s *UpdateRemediationRequest) SetParams(v string) *UpdateRemediationRequest {
	s.Params = &v
	return s
}

func (s *UpdateRemediationRequest) SetRemediationId(v string) *UpdateRemediationRequest {
	s.RemediationId = &v
	return s
}

func (s *UpdateRemediationRequest) SetRemediationTemplateId(v string) *UpdateRemediationRequest {
	s.RemediationTemplateId = &v
	return s
}

func (s *UpdateRemediationRequest) SetRemediationType(v string) *UpdateRemediationRequest {
	s.RemediationType = &v
	return s
}

func (s *UpdateRemediationRequest) SetSourceType(v string) *UpdateRemediationRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateRemediationRequest) Validate() error {
	return dara.Validate(s)
}
