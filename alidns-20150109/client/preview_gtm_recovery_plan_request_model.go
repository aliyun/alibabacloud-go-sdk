// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewGtmRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *PreviewGtmRecoveryPlanRequest
	GetLang() *string
	SetPageNumber(v int32) *PreviewGtmRecoveryPlanRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *PreviewGtmRecoveryPlanRequest
	GetPageSize() *int32
	SetRecoveryPlanId(v int64) *PreviewGtmRecoveryPlanRequest
	GetRecoveryPlanId() *int64
}

type PreviewGtmRecoveryPlanRequest struct {
	// The language of the response. Valid values:
	//
	// zh: Chinese
	//
	// en: English
	//
	// Default: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from **1**. Default: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **20**. Default: **5**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the disaster recovery plan.<props="china">You can call the [DescribeGtmRecoveryPlans](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describegtmrecoveryplans?spm=a2c4g.11186623.help-menu-29697.d_0_5_1_3_13_5.6dd83618vW4yD7) operation to obtain the ID.<props="intl">You can call the [DescribeGtmRecoveryPlans](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describegtmrecoveryplans?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10**
	RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s PreviewGtmRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *PreviewGtmRecoveryPlanRequest) GetLang() *string {
	return s.Lang
}

func (s *PreviewGtmRecoveryPlanRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *PreviewGtmRecoveryPlanRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PreviewGtmRecoveryPlanRequest) GetRecoveryPlanId() *int64 {
	return s.RecoveryPlanId
}

func (s *PreviewGtmRecoveryPlanRequest) SetLang(v string) *PreviewGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) SetPageNumber(v int32) *PreviewGtmRecoveryPlanRequest {
	s.PageNumber = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) SetPageSize(v int32) *PreviewGtmRecoveryPlanRequest {
	s.PageSize = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *PreviewGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

func (s *PreviewGtmRecoveryPlanRequest) Validate() error {
	return dara.Validate(s)
}
