// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDirectConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocateDirectConnectionResponseBody
	GetRequestId() *string
}

type AllocateDirectConnectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateDirectConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateDirectConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateDirectConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateDirectConnectionResponseBody) SetRequestId(v string) *AllocateDirectConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateDirectConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
