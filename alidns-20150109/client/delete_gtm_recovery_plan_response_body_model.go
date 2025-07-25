// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGtmRecoveryPlanResponseBody
	GetRequestId() *string
}

type DeleteGtmRecoveryPlanResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGtmRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGtmRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGtmRecoveryPlanResponseBody) SetRequestId(v string) *DeleteGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGtmRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
