// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteGtmRecoveryPlanRequest
	GetLang() *string
	SetRecoveryPlanId(v int64) *DeleteGtmRecoveryPlanRequest
	GetRecoveryPlanId() *int64
}

type DeleteGtmRecoveryPlanRequest struct {
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the disaster recovery plan that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
}

func (s DeleteGtmRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteGtmRecoveryPlanRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteGtmRecoveryPlanRequest) GetRecoveryPlanId() *int64 {
	return s.RecoveryPlanId
}

func (s *DeleteGtmRecoveryPlanRequest) SetLang(v string) *DeleteGtmRecoveryPlanRequest {
	s.Lang = &v
	return s
}

func (s *DeleteGtmRecoveryPlanRequest) SetRecoveryPlanId(v int64) *DeleteGtmRecoveryPlanRequest {
	s.RecoveryPlanId = &v
	return s
}

func (s *DeleteGtmRecoveryPlanRequest) Validate() error {
	return dara.Validate(s)
}
