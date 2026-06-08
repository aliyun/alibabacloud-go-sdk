// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTemplatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGenerateOptions(v []*string) *GenerateTemplatePolicyRequest
	GetGenerateOptions() []*string
	SetOperationTypes(v []*string) *GenerateTemplatePolicyRequest
	GetOperationTypes() []*string
	SetParameters(v []*GenerateTemplatePolicyRequestParameters) *GenerateTemplatePolicyRequest
	GetParameters() []*GenerateTemplatePolicyRequestParameters
	SetTemplateBody(v string) *GenerateTemplatePolicyRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *GenerateTemplatePolicyRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *GenerateTemplatePolicyRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GenerateTemplatePolicyRequest
	GetTemplateVersion() *string
}

type GenerateTemplatePolicyRequest struct {
	GenerateOptions []*string `json:"GenerateOptions,omitempty" xml:"GenerateOptions,omitempty" type:"Repeated"`
	// The type of operation N for which you want to generate the policy information.
	//
	// Valid values:
	//
	// 	- CreateStack: creates a stack by calling the CreateStack operation.
	//
	// 	- UpdateStack: updates a stack by calling the UpdateStack operation.
	//
	// 	- DeleteStack: deletes a stack by calling the DeleteStack operation.
	//
	// 	- DetectStackDrift: detects drifts on a stack by calling the DelectStackDrift operation.
	//
	// 	- ListStackOperationRisks: lists the risks of a deletion operation on a stack by setting the OperationType parameter to DeleteStack in the ListStackOperationRisks operation.
	//
	// 	- GetTemplateEstimateCost: queries the estimated prices of resources that you want to use in the template by calling the GetTemplateEstimateCost operation.
	//
	// 	- GetTemplateParameterConstraints: queries the values of parameters in the template by calling the GetTemplateParameterConstraints operation.
	//
	// 	- ImportResourcesToStack: imports resources to a stack by setting the ChangeSetType parameter to IMPORT in the CreateChangeSet operation.
	//
	// 	- SignalResource: sends a signal to a stack.
	//
	// >  The default value is the combination of all valid values.
	OperationTypes []*string                                  `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	Parameters     []*GenerateTemplatePolicyRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.
	//
	// If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion":"2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared templates and private templates.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// >  If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// The URL can be up to 1,024 bytes in length.
	//
	// example:
	//
	// oss://ros/template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when the TemplateId parameter is specified.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GenerateTemplatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyRequest) GetGenerateOptions() []*string {
	return s.GenerateOptions
}

func (s *GenerateTemplatePolicyRequest) GetOperationTypes() []*string {
	return s.OperationTypes
}

func (s *GenerateTemplatePolicyRequest) GetParameters() []*GenerateTemplatePolicyRequestParameters {
	return s.Parameters
}

func (s *GenerateTemplatePolicyRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GenerateTemplatePolicyRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GenerateTemplatePolicyRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GenerateTemplatePolicyRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GenerateTemplatePolicyRequest) SetGenerateOptions(v []*string) *GenerateTemplatePolicyRequest {
	s.GenerateOptions = v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetOperationTypes(v []*string) *GenerateTemplatePolicyRequest {
	s.OperationTypes = v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetParameters(v []*GenerateTemplatePolicyRequestParameters) *GenerateTemplatePolicyRequest {
	s.Parameters = v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateBody(v string) *GenerateTemplatePolicyRequest {
	s.TemplateBody = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateId(v string) *GenerateTemplatePolicyRequest {
	s.TemplateId = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateURL(v string) *GenerateTemplatePolicyRequest {
	s.TemplateURL = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateVersion(v string) *GenerateTemplatePolicyRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateTemplatePolicyRequestParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GenerateTemplatePolicyRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyRequestParameters) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GenerateTemplatePolicyRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GenerateTemplatePolicyRequestParameters) SetParameterKey(v string) *GenerateTemplatePolicyRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GenerateTemplatePolicyRequestParameters) SetParameterValue(v string) *GenerateTemplatePolicyRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GenerateTemplatePolicyRequestParameters) Validate() error {
	return dara.Validate(s)
}
