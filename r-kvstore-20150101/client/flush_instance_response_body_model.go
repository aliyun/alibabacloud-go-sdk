// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FlushInstanceResponseBody
	GetRequestId() *string
}

type FlushInstanceResponseBody struct {
	// The operation that you want to perform. Set the value to **FlushInstance**.
	//
	// example:
	//
	// 8D0C0AFC-E9CD-47A4-8395-5C31BF9B3E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FlushInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlushInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *FlushInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlushInstanceResponseBody) SetRequestId(v string) *FlushInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlushInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
