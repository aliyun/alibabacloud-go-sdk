// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferNodeResponseBody
	GetRequestId() *string
	SetResult(v bool) *TransferNodeResponseBody
	GetResult() *bool
}

type TransferNodeResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TransferNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferNodeResponseBody) GoString() string {
	return s.String()
}

func (s *TransferNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferNodeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *TransferNodeResponseBody) SetRequestId(v string) *TransferNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferNodeResponseBody) SetResult(v bool) *TransferNodeResponseBody {
	s.Result = &v
	return s
}

func (s *TransferNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
