// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListVodRealtimeLogDeliveryRequest
	GetOwnerId() *int64
}

type ListVodRealtimeLogDeliveryRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVodRealtimeLogDeliveryRequest) SetOwnerId(v int64) *ListVodRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
