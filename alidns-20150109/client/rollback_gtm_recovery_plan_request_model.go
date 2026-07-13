// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackGtmRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *RollbackGtmRecoveryPlanRequest
	GetLang() *string
	SetRecoveryPlanId(v int64) *RollbackGtmRecoveryPlanRequest
	GetRecoveryPlanId() *int64
}

type RollbackGtmRecoveryPlanRequest struct {
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the disaster recovery plan.<props="china"> Call the [DescribeGtmRecoveryPlans](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describegtmrecoveryplans?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_13_5.6dd83618vW4yD7) operation to obtain the ID.
	//
	// <props="intl">Call the [DescribeGtmRecoveryPlans](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describegtmrecoveryplans?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10**
	RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s RollbackGtmRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *RollbackGtmRecoveryPlanRequest) GetLang() *string {
	return s.Lang
}

func (s *RollbackGtmRecoveryPlanRequest) GetRecoveryPlanId() *int64 {
	return s.RecoveryPlanId
}

func (s *RollbackGtmRecoveryPlanRequest) SetLang(v string) *RollbackGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *RollbackGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *RollbackGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

func (s *RollbackGtmRecoveryPlanRequest) Validate() error {
	return dara.Validate(s)
}
