// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *GetUrgentDemandPlanListRequest
	GetCurrent() *int64
	SetPeriod(v string) *GetUrgentDemandPlanListRequest
	GetPeriod() *string
	SetPlanType(v int64) *GetUrgentDemandPlanListRequest
	GetPlanType() *int64
	SetSize(v int64) *GetUrgentDemandPlanListRequest
	GetSize() *int64
	SetUserId(v string) *GetUrgentDemandPlanListRequest
	GetUserId() *string
}

type GetUrgentDemandPlanListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FY2022
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PlanType *int64 `json:"planType,omitempty" xml:"planType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUrgentDemandPlanListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanListRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *GetUrgentDemandPlanListRequest) GetPeriod() *string {
	return s.Period
}

func (s *GetUrgentDemandPlanListRequest) GetPlanType() *int64 {
	return s.PlanType
}

func (s *GetUrgentDemandPlanListRequest) GetSize() *int64 {
	return s.Size
}

func (s *GetUrgentDemandPlanListRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetUrgentDemandPlanListRequest) SetCurrent(v int64) *GetUrgentDemandPlanListRequest {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetPeriod(v string) *GetUrgentDemandPlanListRequest {
	s.Period = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetPlanType(v int64) *GetUrgentDemandPlanListRequest {
	s.PlanType = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetSize(v int64) *GetUrgentDemandPlanListRequest {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetUserId(v string) *GetUrgentDemandPlanListRequest {
	s.UserId = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) Validate() error {
	return dara.Validate(s)
}
