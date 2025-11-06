// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInRefetchWhoisEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferInRefetchWhoisEmailResponseBody
	GetRequestId() *string
}

type TransferInRefetchWhoisEmailResponseBody struct {
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferInRefetchWhoisEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferInRefetchWhoisEmailResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInRefetchWhoisEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferInRefetchWhoisEmailResponseBody) SetRequestId(v string) *TransferInRefetchWhoisEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferInRefetchWhoisEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
