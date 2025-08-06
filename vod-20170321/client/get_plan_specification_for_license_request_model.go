// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlanSpecificationForLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromApp(v string) *GetPlanSpecificationForLicenseRequest
	GetFromApp() *string
	SetNeedItemSpecification(v bool) *GetPlanSpecificationForLicenseRequest
	GetNeedItemSpecification() *bool
	SetPlanCode(v string) *GetPlanSpecificationForLicenseRequest
	GetPlanCode() *string
	SetPlanId(v string) *GetPlanSpecificationForLicenseRequest
	GetPlanId() *string
}

type GetPlanSpecificationForLicenseRequest struct {
	FromApp               *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	NeedItemSpecification *bool   `json:"NeedItemSpecification,omitempty" xml:"NeedItemSpecification,omitempty"`
	PlanCode              *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	PlanId                *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s GetPlanSpecificationForLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlanSpecificationForLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetPlanSpecificationForLicenseRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *GetPlanSpecificationForLicenseRequest) GetNeedItemSpecification() *bool {
	return s.NeedItemSpecification
}

func (s *GetPlanSpecificationForLicenseRequest) GetPlanCode() *string {
	return s.PlanCode
}

func (s *GetPlanSpecificationForLicenseRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *GetPlanSpecificationForLicenseRequest) SetFromApp(v string) *GetPlanSpecificationForLicenseRequest {
	s.FromApp = &v
	return s
}

func (s *GetPlanSpecificationForLicenseRequest) SetNeedItemSpecification(v bool) *GetPlanSpecificationForLicenseRequest {
	s.NeedItemSpecification = &v
	return s
}

func (s *GetPlanSpecificationForLicenseRequest) SetPlanCode(v string) *GetPlanSpecificationForLicenseRequest {
	s.PlanCode = &v
	return s
}

func (s *GetPlanSpecificationForLicenseRequest) SetPlanId(v string) *GetPlanSpecificationForLicenseRequest {
	s.PlanId = &v
	return s
}

func (s *GetPlanSpecificationForLicenseRequest) Validate() error {
	return dara.Validate(s)
}
