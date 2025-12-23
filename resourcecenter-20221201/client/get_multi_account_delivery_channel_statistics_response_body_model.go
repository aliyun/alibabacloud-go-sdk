// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountDeliveryChannelStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelStatistics(v *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) *GetMultiAccountDeliveryChannelStatisticsResponseBody
	GetDeliveryChannelStatistics() *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics
	SetRequestId(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBody
	GetRequestId() *string
}

type GetMultiAccountDeliveryChannelStatisticsResponseBody struct {
	// The statistics on the delivery channel.
	DeliveryChannelStatistics *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics `json:"DeliveryChannelStatistics,omitempty" xml:"DeliveryChannelStatistics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 80DF0610-504C-56D7-BDCF-7C92FD687***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultiAccountDeliveryChannelStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBody) GetDeliveryChannelStatistics() *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	return s.DeliveryChannelStatistics
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBody) SetDeliveryChannelStatistics(v *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) *GetMultiAccountDeliveryChannelStatisticsResponseBody {
	s.DeliveryChannelStatistics = v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBody) SetRequestId(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBody) Validate() error {
	if s.DeliveryChannelStatistics != nil {
		if err := s.DeliveryChannelStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics struct {
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
	// test-multi-account-delivery
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

func (s GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetLatestChangeDeliveryTime() *string {
	return s.LatestChangeDeliveryTime
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) GetLatestSnapshotDeliveryTime() *string {
	return s.LatestSnapshotDeliveryTime
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetDeliveryChannelName(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestChangeDeliveryTime(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestChangeDeliveryTime = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) SetLatestSnapshotDeliveryTime(v string) *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics {
	s.LatestSnapshotDeliveryTime = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponseBodyDeliveryChannelStatistics) Validate() error {
	return dara.Validate(s)
}
