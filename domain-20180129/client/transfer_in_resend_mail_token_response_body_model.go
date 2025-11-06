// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInResendMailTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferInResendMailTokenResponseBody
	GetRequestId() *string
}

type TransferInResendMailTokenResponseBody struct {
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferInResendMailTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferInResendMailTokenResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInResendMailTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferInResendMailTokenResponseBody) SetRequestId(v string) *TransferInResendMailTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferInResendMailTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
