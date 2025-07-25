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
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on per page. Maximum value: **20**. Default value: **5**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the disaster recovery plan that you want to preview.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
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
