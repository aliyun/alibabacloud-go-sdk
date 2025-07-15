// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveOpenapiReserve(v string) *ListLiveRealtimeLogDeliveryRequest
	GetLiveOpenapiReserve() *string
	SetOwnerId(v int64) *ListLiveRealtimeLogDeliveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListLiveRealtimeLogDeliveryRequest
	GetRegionId() *string
}

type ListLiveRealtimeLogDeliveryRequest struct {
	// This parameter has no practical significance.
	//
	// example:
	//
	// 1
	LiveOpenapiReserve *string `json:"LiveOpenapiReserve,omitempty" xml:"LiveOpenapiReserve,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryRequest) GetLiveOpenapiReserve() *string {
	return s.LiveOpenapiReserve
}

func (s *ListLiveRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListLiveRealtimeLogDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLiveRealtimeLogDeliveryRequest) SetLiveOpenapiReserve(v string) *ListLiveRealtimeLogDeliveryRequest {
	s.LiveOpenapiReserve = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryRequest) SetOwnerId(v int64) *ListLiveRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryRequest) SetRegionId(v string) *ListLiveRealtimeLogDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
