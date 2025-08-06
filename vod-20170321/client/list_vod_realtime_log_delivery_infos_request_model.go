// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListVodRealtimeLogDeliveryInfosRequest
	GetOwnerId() *int64
}

type ListVodRealtimeLogDeliveryInfosRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosRequest) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVodRealtimeLogDeliveryInfosRequest) SetOwnerId(v int64) *ListVodRealtimeLogDeliveryInfosRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosRequest) Validate() error {
	return dara.Validate(s)
}
