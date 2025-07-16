// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type DisableRealtimeLogDeliveryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DisableRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
