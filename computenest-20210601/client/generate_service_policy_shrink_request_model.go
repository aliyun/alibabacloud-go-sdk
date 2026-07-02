// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateServicePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationTypes(v []*string) *GenerateServicePolicyShrinkRequest
	GetOperationTypes() []*string
	SetParametersShrink(v string) *GenerateServicePolicyShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *GenerateServicePolicyShrinkRequest
	GetRegionId() *string
	SetServiceId(v string) *GenerateServicePolicyShrinkRequest
	GetServiceId() *string
	SetServiceVersion(v string) *GenerateServicePolicyShrinkRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *GenerateServicePolicyShrinkRequest
	GetTemplateName() *string
	SetTrialType(v string) *GenerateServicePolicyShrinkRequest
	GetTrialType() *string
}

type GenerateServicePolicyShrinkRequest struct {
	// The operation types for which policy information needs to be generated.
	OperationTypes   []*string `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	ParametersShrink *string   `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s GenerateServicePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateServicePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyShrinkRequest) GetOperationTypes() []*string {
	return s.OperationTypes
}

func (s *GenerateServicePolicyShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *GenerateServicePolicyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateServicePolicyShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GenerateServicePolicyShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GenerateServicePolicyShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GenerateServicePolicyShrinkRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *GenerateServicePolicyShrinkRequest) SetOperationTypes(v []*string) *GenerateServicePolicyShrinkRequest {
	s.OperationTypes = v
	return s
}

func (s *GenerateServicePolicyShrinkRequest) SetParametersShrink(v string) *GenerateServicePolicyShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GenerateServicePolicyShrinkRequest) SetRegionId(v string) *GenerateServicePolicyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateServicePolicyShrinkRequest) SetServiceId(v string) *GenerateServicePolicyShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GenerateServicePolicyShrinkRequest) SetServiceVersion(v string) *GenerateServicePolicyShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GenerateServicePolicyShrinkRequest) SetTemplateName(v string) *GenerateServicePolicyShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateServicePolicyShrinkRequest) SetTrialType(v string) *GenerateServicePolicyShrinkRequest {
	s.TrialType = &v
	return s
}

func (s *GenerateServicePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
