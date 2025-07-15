// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveOpenapiReserve(v string) *ListLiveRealtimeLogDeliveryInfosRequest
	GetLiveOpenapiReserve() *string
	SetOwnerId(v int64) *ListLiveRealtimeLogDeliveryInfosRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListLiveRealtimeLogDeliveryInfosRequest
	GetRegionId() *string
}

type ListLiveRealtimeLogDeliveryInfosRequest struct {
	// This parameter has no practical significance.
	//
	// example:
	//
	// 1
	LiveOpenapiReserve *string `json:"LiveOpenapiReserve,omitempty" xml:"LiveOpenapiReserve,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryInfosRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) GetLiveOpenapiReserve() *string {
	return s.LiveOpenapiReserve
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) SetLiveOpenapiReserve(v string) *ListLiveRealtimeLogDeliveryInfosRequest {
	s.LiveOpenapiReserve = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) SetOwnerId(v int64) *ListLiveRealtimeLogDeliveryInfosRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) SetRegionId(v string) *ListLiveRealtimeLogDeliveryInfosRequest {
	s.RegionId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryInfosRequest) Validate() error {
	return dara.Validate(s)
}
