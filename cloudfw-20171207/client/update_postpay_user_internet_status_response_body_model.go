// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostpayUserInternetStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePostpayUserInternetStatusResponseBody
	GetRequestId() *string
}

type UpdatePostpayUserInternetStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-000901******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePostpayUserInternetStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostpayUserInternetStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePostpayUserInternetStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePostpayUserInternetStatusResponseBody) SetRequestId(v string) *UpdatePostpayUserInternetStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePostpayUserInternetStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
