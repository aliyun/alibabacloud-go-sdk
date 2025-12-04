// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferVersionResponseBody
	GetRequestId() *string
}

type TransferVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7084CDB5-308F-5D0B-8F9B-5F7D83E09738
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferVersionResponseBody) GoString() string {
	return s.String()
}

func (s *TransferVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferVersionResponseBody) SetRequestId(v string) *TransferVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
