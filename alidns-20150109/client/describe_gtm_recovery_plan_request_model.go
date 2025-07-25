// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGtmRecoveryPlanRequest
	GetLang() *string
	SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlanRequest
	GetRecoveryPlanId() *int64
}

type DescribeGtmRecoveryPlanRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the disaster recovery plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s DescribeGtmRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmRecoveryPlanRequest) GetRecoveryPlanId() *int64 {
	return s.RecoveryPlanId
}

func (s *DescribeGtmRecoveryPlanRequest) SetLang(v string) *DescribeGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanRequest) Validate() error {
	return dara.Validate(s)
}
