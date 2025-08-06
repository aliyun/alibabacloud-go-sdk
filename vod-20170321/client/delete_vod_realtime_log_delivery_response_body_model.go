// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVodRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type DeleteVodRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVodRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DeleteVodRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
