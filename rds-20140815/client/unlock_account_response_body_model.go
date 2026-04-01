// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockAccountResponseBody
	GetRequestId() *string
}

type UnlockAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AA65190D-852A-4C9B-88DA-E92698CAA350
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockAccountResponseBody) SetRequestId(v string) *UnlockAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
