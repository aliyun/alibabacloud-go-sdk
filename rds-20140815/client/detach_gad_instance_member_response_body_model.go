// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGadInstanceMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachGadInstanceMemberResponseBody
	GetRequestId() *string
}

type DetachGadInstanceMemberResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachGadInstanceMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachGadInstanceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DetachGadInstanceMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachGadInstanceMemberResponseBody) SetRequestId(v string) *DetachGadInstanceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachGadInstanceMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
