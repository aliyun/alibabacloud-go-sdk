// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateDemandPlanRequest
	GetAccountId() *string
	SetDemandType(v string) *CreateDemandPlanRequest
	GetDemandType() *string
	SetDescription(v string) *CreateDemandPlanRequest
	GetDescription() *string
	SetName(v string) *CreateDemandPlanRequest
	GetName() *string
	SetPeriod(v string) *CreateDemandPlanRequest
	GetPeriod() *string
	SetSource(v string) *CreateDemandPlanRequest
	GetSource() *string
	SetTargetCid(v int64) *CreateDemandPlanRequest
	GetTargetCid() *int64
	SetType(v string) *CreateDemandPlanRequest
	GetType() *string
	SetUserId(v string) *CreateDemandPlanRequest
	GetUserId() *string
}

type CreateDemandPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1065737167271819
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	DemandType  *string `json:"demandType,omitempty" xml:"demandType,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FY2022
	Period    *string `json:"period,omitempty" xml:"period,omitempty"`
	Source    *string `json:"source,omitempty" xml:"source,omitempty"`
	TargetCid *int64  `json:"targetCid,omitempty" xml:"targetCid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// URGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateDemandPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateDemandPlanRequest) GetDemandType() *string {
	return s.DemandType
}

func (s *CreateDemandPlanRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDemandPlanRequest) GetName() *string {
	return s.Name
}

func (s *CreateDemandPlanRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDemandPlanRequest) GetSource() *string {
	return s.Source
}

func (s *CreateDemandPlanRequest) GetTargetCid() *int64 {
	return s.TargetCid
}

func (s *CreateDemandPlanRequest) GetType() *string {
	return s.Type
}

func (s *CreateDemandPlanRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateDemandPlanRequest) SetAccountId(v string) *CreateDemandPlanRequest {
	s.AccountId = &v
	return s
}

func (s *CreateDemandPlanRequest) SetDemandType(v string) *CreateDemandPlanRequest {
	s.DemandType = &v
	return s
}

func (s *CreateDemandPlanRequest) SetDescription(v string) *CreateDemandPlanRequest {
	s.Description = &v
	return s
}

func (s *CreateDemandPlanRequest) SetName(v string) *CreateDemandPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateDemandPlanRequest) SetPeriod(v string) *CreateDemandPlanRequest {
	s.Period = &v
	return s
}

func (s *CreateDemandPlanRequest) SetSource(v string) *CreateDemandPlanRequest {
	s.Source = &v
	return s
}

func (s *CreateDemandPlanRequest) SetTargetCid(v int64) *CreateDemandPlanRequest {
	s.TargetCid = &v
	return s
}

func (s *CreateDemandPlanRequest) SetType(v string) *CreateDemandPlanRequest {
	s.Type = &v
	return s
}

func (s *CreateDemandPlanRequest) SetUserId(v string) *CreateDemandPlanRequest {
	s.UserId = &v
	return s
}

func (s *CreateDemandPlanRequest) Validate() error {
	return dara.Validate(s)
}
