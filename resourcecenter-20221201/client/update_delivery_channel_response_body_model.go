// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDeliveryChannelResponseBody
	GetRequestId() *string
}

type UpdateDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AD5F848D-CCDC-5464-93E1-4BA50A482***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDeliveryChannelResponseBody) SetRequestId(v string) *UpdateDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
