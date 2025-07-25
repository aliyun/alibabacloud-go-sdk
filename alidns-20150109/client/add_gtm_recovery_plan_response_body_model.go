// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecoveryPlanId(v string) *AddGtmRecoveryPlanResponseBody
	GetRecoveryPlanId() *string
	SetRequestId(v string) *AddGtmRecoveryPlanResponseBody
	GetRequestId() *string
}

type AddGtmRecoveryPlanResponseBody struct {
	// The ID of the disaster recovery plan.
	//
	// example:
	//
	// 100
	RecoveryPlanId *string `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGtmRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *AddGtmRecoveryPlanResponseBody) GetRecoveryPlanId() *string {
	return s.RecoveryPlanId
}

func (s *AddGtmRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGtmRecoveryPlanResponseBody) SetRecoveryPlanId(v string) *AddGtmRecoveryPlanResponseBody {
	s.RecoveryPlanId = &v
	return s
}

func (s *AddGtmRecoveryPlanResponseBody) SetRequestId(v string) *AddGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGtmRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
