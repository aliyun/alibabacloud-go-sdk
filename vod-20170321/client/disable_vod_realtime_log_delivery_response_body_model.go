// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVodRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableVodRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type DisableVodRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVodRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableVodRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVodRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableVodRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DisableVodRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableVodRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
