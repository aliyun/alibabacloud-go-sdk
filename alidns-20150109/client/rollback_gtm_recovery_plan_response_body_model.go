// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackGtmRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackGtmRecoveryPlanResponseBody
	GetRequestId() *string
}

type RollbackGtmRecoveryPlanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 853805EA-3D47-47D5-9A1A-A45C24313ABD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackGtmRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackGtmRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackGtmRecoveryPlanResponseBody) SetRequestId(v string) *RollbackGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackGtmRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
