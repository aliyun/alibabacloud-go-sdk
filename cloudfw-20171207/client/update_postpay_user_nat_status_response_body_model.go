// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserNatStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePostpayUserNatStatusResponseBody
	GetRequestId() *string
}

type UpdatePostpayUserNatStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePostpayUserNatStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserNatStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserNatStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePostpayUserNatStatusResponseBody) SetRequestId(v string) *UpdatePostpayUserNatStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePostpayUserNatStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
