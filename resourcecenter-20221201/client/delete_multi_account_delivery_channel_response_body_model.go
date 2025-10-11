// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiAccountDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMultiAccountDeliveryChannelResponseBody
	GetRequestId() *string
}

type DeleteMultiAccountDeliveryChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3C5CDBF6-4DB7-53E9-ADDC-5919E3FAC***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMultiAccountDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultiAccountDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *DeleteMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMultiAccountDeliveryChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
