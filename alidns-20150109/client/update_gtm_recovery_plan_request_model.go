// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaultAddrPool(v string) *UpdateGtmRecoveryPlanRequest
	GetFaultAddrPool() *string
	SetLang(v string) *UpdateGtmRecoveryPlanRequest
	GetLang() *string
	SetName(v string) *UpdateGtmRecoveryPlanRequest
	GetName() *string
	SetRecoveryPlanId(v int64) *UpdateGtmRecoveryPlanRequest
	GetRecoveryPlanId() *int64
	SetRemark(v string) *UpdateGtmRecoveryPlanRequest
	GetRemark() *string
}

type UpdateGtmRecoveryPlanRequest struct {
	// The list of faulty address pools.
	//
	// example:
	//
	// ["hra0or"]
	FaultAddrPool *string `json:"FaultAddrPool,omitempty" xml:"FaultAddrPool,omitempty"`
	// The language in which you want the values of some response parameters to be returned. These response parameters support multiple languages.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the disaster recovery plan.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the disaster recovery plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	// The remarks about the disaster recovery plan.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateGtmRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmRecoveryPlanRequest) GetFaultAddrPool() *string {
	return s.FaultAddrPool
}

func (s *UpdateGtmRecoveryPlanRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateGtmRecoveryPlanRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGtmRecoveryPlanRequest) GetRecoveryPlanId() *int64 {
	return s.RecoveryPlanId
}

func (s *UpdateGtmRecoveryPlanRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateGtmRecoveryPlanRequest) SetFaultAddrPool(v string) *UpdateGtmRecoveryPlanRequest {
	s.FaultAddrPool = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetLang(v string) *UpdateGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetName(v string) *UpdateGtmRecoveryPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *UpdateGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) SetRemark(v string) *UpdateGtmRecoveryPlanRequest {
	s.Remark = &v
	return s
}

func (s *UpdateGtmRecoveryPlanRequest) Validate() error {
	return dara.Validate(s)
}
