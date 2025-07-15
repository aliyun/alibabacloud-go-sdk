// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLiveRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableLiveRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type DisableLiveRealtimeLogDeliveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLiveRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableLiveRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLiveRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableLiveRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DisableLiveRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableLiveRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
