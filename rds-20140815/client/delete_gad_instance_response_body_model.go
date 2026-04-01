// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGadInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGadInstanceResponseBody
	GetRequestId() *string
}

type DeleteGadInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGadInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGadInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGadInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGadInstanceResponseBody) SetRequestId(v string) *DeleteGadInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGadInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
