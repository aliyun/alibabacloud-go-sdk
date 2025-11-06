// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInReenterTransferAuthorizationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferInReenterTransferAuthorizationCodeResponseBody
	GetRequestId() *string
}

type TransferInReenterTransferAuthorizationCodeResponseBody struct {
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferInReenterTransferAuthorizationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferInReenterTransferAuthorizationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInReenterTransferAuthorizationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferInReenterTransferAuthorizationCodeResponseBody) SetRequestId(v string) *TransferInReenterTransferAuthorizationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
