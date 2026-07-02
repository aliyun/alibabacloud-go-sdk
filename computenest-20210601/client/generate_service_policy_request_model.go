// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateServicePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationTypes(v []*string) *GenerateServicePolicyRequest
	GetOperationTypes() []*string
	SetParameters(v map[string]interface{}) *GenerateServicePolicyRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *GenerateServicePolicyRequest
	GetRegionId() *string
	SetServiceId(v string) *GenerateServicePolicyRequest
	GetServiceId() *string
	SetServiceVersion(v string) *GenerateServicePolicyRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *GenerateServicePolicyRequest
	GetTemplateName() *string
	SetTrialType(v string) *GenerateServicePolicyRequest
	GetTrialType() *string
}

type GenerateServicePolicyRequest struct {
	// The operation types for which policy information needs to be generated.
	OperationTypes []*string              `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	Parameters     map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-b3e9ed878b0c4xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial type. Default value: NotTrial. Valid values:
	//
	// - Trial: Trial is supported.
	//
	// - NotTrial: Trial is not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GenerateServicePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyRequest) GetOperationTypes() []*string {
	return s.OperationTypes
}

func (s *GenerateServicePolicyRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GenerateServicePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateServicePolicyRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GenerateServicePolicyRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GenerateServicePolicyRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GenerateServicePolicyRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *GenerateServicePolicyRequest) SetOperationTypes(v []*string) *GenerateServicePolicyRequest {
	s.OperationTypes = v
	return s
}

func (s *GenerateServicePolicyRequest) SetParameters(v map[string]interface{}) *GenerateServicePolicyRequest {
	s.Parameters = v
	return s
}

func (s *GenerateServicePolicyRequest) SetRegionId(v string) *GenerateServicePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetServiceId(v string) *GenerateServicePolicyRequest {
	s.ServiceId = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetServiceVersion(v string) *GenerateServicePolicyRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetTemplateName(v string) *GenerateServicePolicyRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetTrialType(v string) *GenerateServicePolicyRequest {
	s.TrialType = &v
	return s
}

func (s *GenerateServicePolicyRequest) Validate() error {
	return dara.Validate(s)
}
