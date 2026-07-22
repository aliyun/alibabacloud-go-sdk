// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserVpcStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePostpayUserVpcStatusResponseBody
	GetRequestId() *string
}

type UpdatePostpayUserVpcStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 450D47F5-956E-543E-8502-2F71C8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePostpayUserVpcStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserVpcStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserVpcStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePostpayUserVpcStatusResponseBody) SetRequestId(v string) *UpdatePostpayUserVpcStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePostpayUserVpcStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
