// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaultAddrPool(v string) *AddGtmRecoveryPlanRequest
	GetFaultAddrPool() *string
	SetLang(v string) *AddGtmRecoveryPlanRequest
	GetLang() *string
	SetName(v string) *AddGtmRecoveryPlanRequest
	GetName() *string
	SetRemark(v string) *AddGtmRecoveryPlanRequest
	GetRemark() *string
}

type AddGtmRecoveryPlanRequest struct {
	// The IDs of faulty address pools.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["hra0or"]
	FaultAddrPool *string `json:"FaultAddrPool,omitempty" xml:"FaultAddrPool,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the disaster recovery plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// name-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the disaster recovery plan.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddGtmRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *AddGtmRecoveryPlanRequest) GetFaultAddrPool() *string {
	return s.FaultAddrPool
}

func (s *AddGtmRecoveryPlanRequest) GetLang() *string {
	return s.Lang
}

func (s *AddGtmRecoveryPlanRequest) GetName() *string {
	return s.Name
}

func (s *AddGtmRecoveryPlanRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddGtmRecoveryPlanRequest) SetFaultAddrPool(v string) *AddGtmRecoveryPlanRequest {
	s.FaultAddrPool = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) SetLang(v string) *AddGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) SetName(v string) *AddGtmRecoveryPlanRequest {
	s.Name = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) SetRemark(v string) *AddGtmRecoveryPlanRequest {
	s.Remark = &v
	return s
}

func (s *AddGtmRecoveryPlanRequest) Validate() error {
	return dara.Validate(s)
}
