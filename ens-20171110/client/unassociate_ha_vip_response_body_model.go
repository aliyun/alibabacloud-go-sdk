// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateHaVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateHaVipResponseBody
	GetRequestId() *string
}

type UnassociateHaVipResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateHaVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateHaVipResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateHaVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateHaVipResponseBody) SetRequestId(v string) *UnassociateHaVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateHaVipResponseBody) Validate() error {
	return dara.Validate(s)
}
