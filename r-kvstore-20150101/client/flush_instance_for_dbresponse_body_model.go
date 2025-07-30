// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushInstanceForDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FlushInstanceForDBResponseBody
	GetRequestId() *string
}

type FlushInstanceForDBResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FlushInstanceForDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlushInstanceForDBResponseBody) GoString() string {
	return s.String()
}

func (s *FlushInstanceForDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlushInstanceForDBResponseBody) SetRequestId(v string) *FlushInstanceForDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlushInstanceForDBResponseBody) Validate() error {
	return dara.Validate(s)
}
