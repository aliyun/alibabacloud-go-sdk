// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTemplateParameterConstraintsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest
	GetClientToken() *string
	SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetDeployRegionId() *string
	SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest
	GetEnablePrivateVpcConnection() *bool
	SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest
	GetParameters() []*GetServiceTemplateParameterConstraintsRequestParameters
	SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest
	GetServiceVersion() *string
	SetSpecificationName(v string) *GetServiceTemplateParameterConstraintsRequest
	GetSpecificationName() *string
	SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest
	GetTemplateName() *string
	SetTrialType(v string) *GetServiceTemplateParameterConstraintsRequest
	GetTrialType() *string
}

type GetServiceTemplateParameterConstraintsRequest struct {
	// The client token that is used to ensure the idempotence of the request. Generate a unique value from your client for each request. The token can contain only ASCII characters and cannot be more than 64 characters long.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The deployment region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// Indicates whether PrivateLink is enabled. Valid values:
	//
	// - true: enabled
	//
	// - false: disabled
	//
	// example:
	//
	// true
	EnablePrivateVpcConnection *bool `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	// The configuration parameters of the service instance.
	Parameters []*GetServiceTemplateParameterConstraintsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
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
	// service-731f788406024axxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-461ee95f46ca46xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The specification name.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The usage type. Valid values:
	//
	// - Trial: The service supports a trial.
	//
	// - NotTrial: The service does not support a trial.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetEnablePrivateVpcConnection() *bool {
	return s.EnablePrivateVpcConnection
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetParameters() []*GetServiceTemplateParameterConstraintsRequestParameters {
	return s.Parameters
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceTemplateParameterConstraintsRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetSpecificationName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTrialType(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TrialType = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) Validate() error {
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

type GetServiceTemplateParameterConstraintsRequestParameters struct {
	// The name of the parameter. If you do not specify the name and value of the parameter, Resource Orchestration Service (ROS) uses the default value that is specified in the template.
	//
	// > The Parameters parameter is optional. If you specify Parameters, ParameterKey is required.
	//
	// example:
	//
	// InstanceType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter that is defined in the template.
	//
	// > The Parameters parameter is optional. If you specify Parameters, ParameterValue is required.
	//
	// example:
	//
	// cn-hangzhou-j
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterValue(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) Validate() error {
	return dara.Validate(s)
}
