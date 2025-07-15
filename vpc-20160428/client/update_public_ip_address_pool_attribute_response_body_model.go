// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicIpAddressPoolAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePublicIpAddressPoolAttributeResponseBody
	GetRequestId() *string
}

type UpdatePublicIpAddressPoolAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePublicIpAddressPoolAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicIpAddressPoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePublicIpAddressPoolAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePublicIpAddressPoolAttributeResponseBody) SetRequestId(v string) *UpdatePublicIpAddressPoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePublicIpAddressPoolAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
