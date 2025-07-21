// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMailAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMailAddressResponseBody
	GetRequestId() *string
}

type ModifyMailAddressResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMailAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMailAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMailAddressResponseBody) SetRequestId(v string) *ModifyMailAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMailAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
