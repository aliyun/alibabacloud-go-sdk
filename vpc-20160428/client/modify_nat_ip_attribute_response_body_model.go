// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatIpAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNatIpAttributeResponseBody
	GetRequestId() *string
}

type ModifyNatIpAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6CC9456C-2E29-452A-9180-B6926E51B5CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatIpAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatIpAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatIpAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatIpAttributeResponseBody) SetRequestId(v string) *ModifyNatIpAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatIpAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
