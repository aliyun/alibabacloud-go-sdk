// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGtmRecoveryPlanResponseBody
	GetRequestId() *string
}

type UpdateGtmRecoveryPlanResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGtmRecoveryPlanResponseBody) SetRequestId(v string) *UpdateGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGtmRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
