// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEstimateCostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetServiceEstimateCostRequest
	GetClientToken() *string
	SetCommodity(v *GetServiceEstimateCostRequestCommodity) *GetServiceEstimateCostRequest
	GetCommodity() *GetServiceEstimateCostRequestCommodity
	SetParameters(v map[string]interface{}) *GetServiceEstimateCostRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *GetServiceEstimateCostRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceEstimateCostRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *GetServiceEstimateCostRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *GetServiceEstimateCostRequest
	GetServiceVersion() *string
	SetSpecificationName(v string) *GetServiceEstimateCostRequest
	GetSpecificationName() *string
	SetTemplateName(v string) *GetServiceEstimateCostRequest
	GetTemplateName() *string
}

type GetServiceEstimateCostRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// mRdxWuW2ts
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration information about the purchase order of Alibaba Cloud Marketplace.
	Commodity *GetServiceEstimateCostRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"PayType\\":\\"PostPaid\\",\\"InstancePassword\\":\\"xxxxxxxxxx\\",\\"EcsInstanceType\\":\\"ecs.g6.large\\",\\"VSwitchId\\":\\"vsw-0jlueyydpuekoxxxxxxxx\\",\\"VpcId\\":\\"vpc-0jlps6mjbgvpqxxxxxxxx\\",\\"ZoneId\\":\\"cn-wulanchabu-a\\",\\"Enable\\":false,\\"RegionId\\":\\"cn-wulanchabu\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s GetServiceEstimateCostRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetServiceEstimateCostRequest) GetCommodity() *GetServiceEstimateCostRequestCommodity {
	return s.Commodity
}

func (s *GetServiceEstimateCostRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetServiceEstimateCostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceEstimateCostRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceEstimateCostRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceEstimateCostRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceEstimateCostRequest) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *GetServiceEstimateCostRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceEstimateCostRequest) SetClientToken(v string) *GetServiceEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetCommodity(v *GetServiceEstimateCostRequestCommodity) *GetServiceEstimateCostRequest {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetParameters(v map[string]interface{}) *GetServiceEstimateCostRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetRegionId(v string) *GetServiceEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceVersion(v string) *GetServiceEstimateCostRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetSpecificationName(v string) *GetServiceEstimateCostRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetTemplateName(v string) *GetServiceEstimateCostRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceEstimateCostRequest) Validate() error {
	return dara.Validate(s)
}

type GetServiceEstimateCostRequestCommodity struct {
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int32 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- Year
	//
	// 	- Month
	//
	// 	- Day
	//
	// example:
	//
	// Month
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s GetServiceEstimateCostRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEstimateCostRequestCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequestCommodity) GetPayPeriod() *int32 {
	return s.PayPeriod
}

func (s *GetServiceEstimateCostRequestCommodity) GetPayPeriodUnit() *string {
	return s.PayPeriodUnit
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriod(v int32) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriodUnit(v string) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) Validate() error {
	return dara.Validate(s)
}
