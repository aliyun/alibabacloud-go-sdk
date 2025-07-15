// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveCenterTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveCenterTransferResponseBody
	GetRequestId() *string
}

type AddLiveCenterTransferResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7908F2FF-44F8-120F-9FD6-85AE4B6C19EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveCenterTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveCenterTransferResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveCenterTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveCenterTransferResponseBody) SetRequestId(v string) *AddLiveCenterTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveCenterTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
