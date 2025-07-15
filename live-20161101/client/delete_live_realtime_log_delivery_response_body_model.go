// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type DeleteLiveRealtimeLogDeliveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DeleteLiveRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
