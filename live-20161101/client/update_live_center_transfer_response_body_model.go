// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveCenterTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveCenterTransferResponseBody
	GetRequestId() *string
}

type UpdateLiveCenterTransferResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7908F2FF-44F8-120F-9FD6-85AE4B6C19EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveCenterTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveCenterTransferResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveCenterTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveCenterTransferResponseBody) SetRequestId(v string) *UpdateLiveCenterTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveCenterTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
