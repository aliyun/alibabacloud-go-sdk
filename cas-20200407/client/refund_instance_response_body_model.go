// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefundInstanceResponseBody
	GetRequestId() *string
}

type RefundInstanceResponseBody struct {
	// example:
	//
	// D3F1FA43-1C26-50A2-8F0F-7A03851DBB46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RefundInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundInstanceResponseBody) SetRequestId(v string) *RefundInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
