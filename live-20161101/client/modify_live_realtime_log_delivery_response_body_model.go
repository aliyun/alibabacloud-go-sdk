// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLiveRealtimeLogDeliveryResponseBody
	GetRequestId() *string
}

type ModifyLiveRealtimeLogDeliveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveRealtimeLogDeliveryResponseBody) SetRequestId(v string) *ModifyLiveRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
