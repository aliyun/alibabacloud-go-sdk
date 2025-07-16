// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type ModifyRealtimeLogDeliveryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRealtimeLogDeliveryResponseBody) SetRequestId(v string) *ModifyRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
