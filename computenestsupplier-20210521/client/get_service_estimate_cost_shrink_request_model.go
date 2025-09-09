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
}

type GetServiceEstimateCostShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// mRdxWuW2ts
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration information about the purchase order of Alibaba Cloud Marketplace.
	CommodityShrink *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"PayType\\":\\"PostPaid\\",\\"InstancePassword\\":\\"xxxxxxxxxx\\",\\"EcsInstanceType\\":\\"ecs.g6.large\\",\\"VSwitchId\\":\\"vsw-0jlueyydpuekoxxxxxxxx\\",\\"VpcId\\":\\"vpc-0jlps6mjbgvpqxxxxxxxx\\",\\"ZoneId\\":\\"cn-wulanchabu-a\\",\\"Enable\\":false,\\"RegionId\\":\\"cn-wulanchabu\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-16fbd358d75e49xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Custom_Image_Ecs
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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

func (s *GetServiceEstimateCostShrinkRequest) SetClientToken(v string) *GetServiceEstimateCostShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetCommodityShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.CommodityShrink = &v
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

func (s *GetServiceEstimateCostShrinkRequest) Validate() error {
	return dara.Validate(s)
}
