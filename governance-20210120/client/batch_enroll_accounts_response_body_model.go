// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEnrollAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchEnrollAccountsResponseBody
	GetRequestId() *string
}

type BatchEnrollAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16B208DD-86BD-5E7D-AC93-FFD44B6FBDF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchEnrollAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchEnrollAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchEnrollAccountsResponseBody) SetRequestId(v string) *BatchEnrollAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchEnrollAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
