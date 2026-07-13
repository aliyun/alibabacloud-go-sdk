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
	// The list of IDs of the fault address pools.
	//
	// example:
	//
	// ["hra0**"]
	FaultAddrPool *string `json:"FaultAddrPool,omitempty" xml:"FaultAddrPool,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default value: en
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
	// The ID of the disaster recovery plan.<props="china">You can call the [DescribeGtmRecoveryPlans](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describegtmrecoveryplans?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_13_5.6dd83618vW4yD7) operation to obtain the ID.
	//
	// <props="intl">You can call the [DescribeGtmRecoveryPlans](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describegtmrecoveryplans?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10*******
	RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	// The remarks.
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
