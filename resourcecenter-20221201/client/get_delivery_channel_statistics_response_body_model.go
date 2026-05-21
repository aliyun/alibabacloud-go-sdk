// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryChannelStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelStatistics(v *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) *GetDeliveryChannelStatisticsResponseBody
	GetDeliveryChannelStatistics() *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics
	SetRequestId(v string) *GetDeliveryChannelStatisticsResponseBody
	GetRequestId() *string
}

type GetDeliveryChannelStatisticsResponseBody struct {
	// The statistics on the delivery channel.
	DeliveryChannelStatistics *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics `json:"DeliveryChannelStatistics,omitempty" xml:"DeliveryChannelStatistics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 80DF0610-504C-56D7-BDCF-7C92FD687***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeliveryChannelStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsResponseBody) GetDeliveryChannelStatistics() *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	return s.DeliveryChannelStatistics
}

func (s *GetDeliveryChannelStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeliveryChannelStatisticsResponseBody) SetDeliveryChannelStatistics(v *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) *GetDeliveryChannelStatisticsResponseBody {
	s.DeliveryChannelStatistics = v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBody) SetRequestId(v string) *GetDeliveryChannelStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBody) Validate() error {
	if s.DeliveryChannelStatistics != nil {
		if err := s.DeliveryChannelStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics struct {
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The last delivery time of resource configuration change events.
	//
	// example:
	//
	// 2025-06-03T16:05:15Z
	LatestChangeDeliveryTime *string `json:"LatestChangeDeliveryTime,omitempty" xml:"LatestChangeDeliveryTime,omitempty"`
	// The last delivery time of scheduled resource snapshots.
	//
	// example:
	//
	// 2025-06-03T16:00:00Z
	LatestSnapshotDeliveryTime *string `json:"LatestSnapshotDeliveryTime,omitempty" xml:"LatestSnapshotDeliveryTime,omitempty"`
}

func (s GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetLatestChangeDeliveryTime() *string {
	return s.LatestChangeDeliveryTime
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetLatestSnapshotDeliveryTime() *string {
	return s.LatestSnapshotDeliveryTime
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelId(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelName(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestChangeDeliveryTime(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestChangeDeliveryTime = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestSnapshotDeliveryTime(v string) *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestSnapshotDeliveryTime = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) Validate() error {
	return dara.Validate(s)
}
