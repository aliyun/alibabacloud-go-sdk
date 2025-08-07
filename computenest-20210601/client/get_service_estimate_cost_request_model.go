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
	SetOperationName(v string) *GetServiceEstimateCostRequest
	GetOperationName() *string
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
	SetTrialType(v string) *GetServiceEstimateCostRequest
	GetTrialType() *string
}

type GetServiceEstimateCostRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// qwertyuiop
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the subscription duration.
	Commodity *GetServiceEstimateCostRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s *GetServiceEstimateCostRequest) GetOperationName() *string {
	return s.OperationName
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

func (s *GetServiceEstimateCostRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *GetServiceEstimateCostRequest) SetClientToken(v string) *GetServiceEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetCommodity(v *GetServiceEstimateCostRequestCommodity) *GetServiceEstimateCostRequest {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetOperationName(v string) *GetServiceEstimateCostRequest {
	s.OperationName = &v
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

func (s *GetServiceEstimateCostRequest) SetTrialType(v string) *GetServiceEstimateCostRequest {
	s.TrialType = &v
	return s
}

func (s *GetServiceEstimateCostRequest) Validate() error {
	return dara.Validate(s)
}

type GetServiceEstimateCostRequestCommodity struct {
	// The coupon ID.
	//
	// example:
	//
	// 302070970220
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The subscription duration of the instance.
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
	// Year
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
	// Marketplace private offer Id
	//
	// example:
	//
	// 17cde2e8-2f5d-xxxx-xxxx-5120dd215d66
	QuotationId *string `json:"QuotationId,omitempty" xml:"QuotationId,omitempty"`
}

func (s GetServiceEstimateCostRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEstimateCostRequestCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequestCommodity) GetCouponId() *string {
	return s.CouponId
}

func (s *GetServiceEstimateCostRequestCommodity) GetPayPeriod() *int32 {
	return s.PayPeriod
}

func (s *GetServiceEstimateCostRequestCommodity) GetPayPeriodUnit() *string {
	return s.PayPeriodUnit
}

func (s *GetServiceEstimateCostRequestCommodity) GetQuotationId() *string {
	return s.QuotationId
}

func (s *GetServiceEstimateCostRequestCommodity) SetCouponId(v string) *GetServiceEstimateCostRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriod(v int32) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriodUnit(v string) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetQuotationId(v string) *GetServiceEstimateCostRequestCommodity {
	s.QuotationId = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) Validate() error {
	return dara.Validate(s)
}
