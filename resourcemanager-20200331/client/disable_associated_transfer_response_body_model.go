// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAssociatedTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAssociatedTransferResponseBody
	GetRequestId() *string
}

type DisableAssociatedTransferResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7CE0AE54-6F27-5522-A429-4C5EE8FD40C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAssociatedTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAssociatedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAssociatedTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAssociatedTransferResponseBody) SetRequestId(v string) *DisableAssociatedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAssociatedTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
