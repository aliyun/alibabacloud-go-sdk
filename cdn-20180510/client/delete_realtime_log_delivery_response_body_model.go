// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type DeleteRealtimeLogDeliveryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DeleteRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
