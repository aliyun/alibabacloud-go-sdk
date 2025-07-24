// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEstimateCostShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetServiceEstimateCostShrinkRequest
	GetClientToken() *string
	SetCommodityShrink(v string) *GetServiceEstimateCostShrinkRequest
	GetCommodityShrink() *string
	SetOperationName(v string) *GetServiceEstimateCostShrinkRequest
	GetOperationName() *string
	SetParametersShrink(v string) *GetServiceEstimateCostShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *GetServiceEstimateCostShrinkRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceEstimateCostShrinkRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *GetServiceEstimateCostShrinkRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *GetServiceEstimateCostShrinkRequest
	GetServiceVersion() *string
	SetSpecificationName(v string) *GetServiceEstimateCostShrinkRequest
	GetSpecificationName() *string
	SetTemplateName(v string) *GetServiceEstimateCostShrinkRequest
	GetTemplateName() *string
	SetTrialType(v string) *GetServiceEstimateCostShrinkRequest
	GetTrialType() *string
}

type GetServiceEstimateCostShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// qwertyuiop
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the subscription duration.
	CommodityShrink *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The name of the configuration change operation.
	//
	// example:
	//
	// Parameter change
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// { \\"RegionId\\": \\"cn-hangzhou\\", \\"InstanceType\\": \\"ecs.g5.large\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
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
	// service-12xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17xxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the package specification.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceEstimateCostShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEstimateCostShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetServiceEstimateCostShrinkRequest) GetCommodityShrink() *string {
	return s.CommodityShrink
}

func (s *GetServiceEstimateCostShrinkRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *GetServiceEstimateCostShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *GetServiceEstimateCostShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceEstimateCostShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceEstimateCostShrinkRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceEstimateCostShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceEstimateCostShrinkRequest) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *GetServiceEstimateCostShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceEstimateCostShrinkRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *GetServiceEstimateCostShrinkRequest) SetClientToken(v string) *GetServiceEstimateCostShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetCommodityShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.CommodityShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetOperationName(v string) *GetServiceEstimateCostShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetParametersShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetRegionId(v string) *GetServiceEstimateCostShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceVersion(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetSpecificationName(v string) *GetServiceEstimateCostShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetTemplateName(v string) *GetServiceEstimateCostShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetTrialType(v string) *GetServiceEstimateCostShrinkRequest {
	s.TrialType = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) Validate() error {
	return dara.Validate(s)
}
