// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiAccountDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMultiAccountDeliveryChannelResponseBody
	GetRequestId() *string
}

type UpdateMultiAccountDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *UpdateMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
