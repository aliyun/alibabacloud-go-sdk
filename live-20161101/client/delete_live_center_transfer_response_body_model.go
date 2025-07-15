// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveCenterTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveCenterTransferResponseBody
	GetRequestId() *string
}

type DeleteLiveCenterTransferResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7908F2FF-44F8-120F-9FD6-85A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveCenterTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveCenterTransferResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveCenterTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveCenterTransferResponseBody) SetRequestId(v string) *DeleteLiveCenterTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveCenterTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
