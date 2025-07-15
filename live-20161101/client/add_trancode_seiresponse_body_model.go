// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrancodeSEIResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTrancodeSEIResponseBody
	GetRequestId() *string
}

type AddTrancodeSEIResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4E*****43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTrancodeSEIResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTrancodeSEIResponseBody) GoString() string {
	return s.String()
}

func (s *AddTrancodeSEIResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTrancodeSEIResponseBody) SetRequestId(v string) *AddTrancodeSEIResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTrancodeSEIResponseBody) Validate() error {
	return dara.Validate(s)
}
