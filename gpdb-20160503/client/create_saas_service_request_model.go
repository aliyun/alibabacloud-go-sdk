// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaasServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v int32) *CreateSaasServiceRequest
	GetCu() *int32
	SetPayType(v string) *CreateSaasServiceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateSaasServiceRequest
	GetPeriod() *string
	SetPlan(v string) *CreateSaasServiceRequest
	GetPlan() *string
	SetRegionId(v string) *CreateSaasServiceRequest
	GetRegionId() *string
	SetServiceType(v string) *CreateSaasServiceRequest
	GetServiceType() *string
	SetUsedTime(v string) *CreateSaasServiceRequest
	GetUsedTime() *string
	SetWorkspaceId(v string) *CreateSaasServiceRequest
	GetWorkspaceId() *string
}

type CreateSaasServiceRequest struct {
	// The compute resources of the SaaS service.
	//
	// example:
	//
	// 1
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// > - If you leave this parameter empty, a Free service is created by default.
	//
	// > - The subscription billing method offers discounts for purchases of one year or longer. Select a billing method as needed.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the duration for which you want to purchase the resource. Valid values:
	//
	// - **Month**: month.
	//
	// - **Year**: year.
	//
	// > This parameter is required when you create a subscription instance.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// deprecated
	Plan *string `json:"Plan,omitempty" xml:"Plan,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service type. Valid values:
	//
	// - **memroy**
	//
	// - **drama**.
	//
	// This parameter is required.
	//
	// example:
	//
	// drama
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The duration for which you want to purchase the resource. Valid values:
	//
	// - If **Period*	- is set to **Month**, the valid values are 1 to 11.
	//
	// - If **Period*	- is set to **Year**, the valid values are 1 to 3.
	//
	// > This parameter is required when you create a subscription instance.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The workspace of the SaaS service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateSaasServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSaasServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateSaasServiceRequest) GetCu() *int32 {
	return s.Cu
}

func (s *CreateSaasServiceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateSaasServiceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateSaasServiceRequest) GetPlan() *string {
	return s.Plan
}

func (s *CreateSaasServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSaasServiceRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateSaasServiceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateSaasServiceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateSaasServiceRequest) SetCu(v int32) *CreateSaasServiceRequest {
	s.Cu = &v
	return s
}

func (s *CreateSaasServiceRequest) SetPayType(v string) *CreateSaasServiceRequest {
	s.PayType = &v
	return s
}

func (s *CreateSaasServiceRequest) SetPeriod(v string) *CreateSaasServiceRequest {
	s.Period = &v
	return s
}

func (s *CreateSaasServiceRequest) SetPlan(v string) *CreateSaasServiceRequest {
	s.Plan = &v
	return s
}

func (s *CreateSaasServiceRequest) SetRegionId(v string) *CreateSaasServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSaasServiceRequest) SetServiceType(v string) *CreateSaasServiceRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateSaasServiceRequest) SetUsedTime(v string) *CreateSaasServiceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateSaasServiceRequest) SetWorkspaceId(v string) *CreateSaasServiceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateSaasServiceRequest) Validate() error {
	return dara.Validate(s)
}
