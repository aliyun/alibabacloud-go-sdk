// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeliveryChannelResponseBody
	GetRequestId() *string
}

type DeleteDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AD5F848D-CCDC-5464-93E1-4BA50A482***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeliveryChannelResponseBody) SetRequestId(v string) *DeleteDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
