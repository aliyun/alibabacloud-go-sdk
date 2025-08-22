// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangingAffectTime(v string) *DescribeDcdnServiceResponseBody
	GetChangingAffectTime() *string
	SetChangingChargeType(v string) *DescribeDcdnServiceResponseBody
	GetChangingChargeType() *string
	SetInstanceId(v string) *DescribeDcdnServiceResponseBody
	GetInstanceId() *string
	SetInternetChargeType(v string) *DescribeDcdnServiceResponseBody
	GetInternetChargeType() *string
	SetOpeningTime(v string) *DescribeDcdnServiceResponseBody
	GetOpeningTime() *string
	SetOperationLocks(v *DescribeDcdnServiceResponseBodyOperationLocks) *DescribeDcdnServiceResponseBody
	GetOperationLocks() *DescribeDcdnServiceResponseBodyOperationLocks
	SetRequestId(v string) *DescribeDcdnServiceResponseBody
	GetRequestId() *string
	SetWebsocketChangingTime(v string) *DescribeDcdnServiceResponseBody
	GetWebsocketChangingTime() *string
	SetWebsocketChangingType(v string) *DescribeDcdnServiceResponseBody
	GetWebsocketChangingType() *string
	SetWebsocketType(v string) *DescribeDcdnServiceResponseBody
	GetWebsocketType() *string
}

type DescribeDcdnServiceResponseBody struct {
	// The time when the renewed secure DCDN takes effect. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-03-31T16:00:00Z
	ChangingAffectTime *string `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	// The new metering method for the renewed secure DCDN. Valid values:
	//
	// 	- **PayByTraffic**: pay by data transfer
	//
	// 	- **PayByBandwidth**: pay by bandwidth
	//
	// 	- **PayByBandwidth95**: pay by 95th percentile bandwidth
	//
	// 	- **PayByBandwidth_monthavg**: pay by monthly average bandwidth
	//
	// 	- **PayByBandwidth_month4th**: pay by fourth peak bandwidth per month
	//
	// 	- **PayByBandwidth_monthday95avg**: pay by monthly average 95th percentile bandwidth
	//
	// 	- **PayByBandwidth_nighthalf95**: pay by 95th percentile bandwidth (50% off during nighttime)
	//
	// example:
	//
	// PayByBandwidth
	ChangingChargeType *string `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// FP-mkqgwxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The current metering method. Valid values:
	//
	// 	- **PayByTraffic**: pay by data transfer
	//
	// 	- **PayByBandwidth**: pay by bandwidth
	//
	// 	- **PayByBandwidth95**: pay by 95th percentile bandwidth
	//
	// 	- **PayByBandwidth_monthavg**: pay by monthly average bandwidth
	//
	// 	- **PayByBandwidth_month4th**: pay by fourth peak bandwidth per month
	//
	// 	- **PayByBandwidth_monthday95avg**: pay by monthly average 95th percentile bandwidth
	//
	// 	- **PayByBandwidth_nighthalf95**: pay by 95th percentile bandwidth (50% off during nighttime)
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The time when the DCDN service was activated. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2018-03-19T11:16:11Z
	OpeningTime *string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	// The lock status of DCDN.
	OperationLocks *DescribeDcdnServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EF2AEBC2-EDBD-41CF-BF64-7E095D42D6EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the changes of the WebSocket configuration take effect. The value is the same as that of the ChangingAffectTime parameter. This parameter can be displayed in the console only if the specified time is later than the current time.
	//
	// example:
	//
	// 2018-03-19T11:16:11Z
	WebsocketChangingTime *string `json:"WebsocketChangingTime,omitempty" xml:"WebsocketChangingTime,omitempty"`
	// The next effective billing method of WebSocket. Valid values: **websockettraffic*	- and **websocketbps**. A value of websockettraffic indicates that you are billed based on the traffic volume. A value of websocketbps indicates that you are billed based on the bandwidth.
	//
	// example:
	//
	// websocketbps
	WebsocketChangingType *string `json:"WebsocketChangingType,omitempty" xml:"WebsocketChangingType,omitempty"`
	// The current billing method of WebSocket. Valid values: **websockettraffic*	- and **websocketbps**. A value of websockettraffic indicates that you are billed based on the traffic volume. A value of websocketbps indicates that you are billed based on the bandwidth.
	//
	// example:
	//
	// websocketbps
	WebsocketType *string `json:"WebsocketType,omitempty" xml:"WebsocketType,omitempty"`
}

func (s DescribeDcdnServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponseBody) GetChangingAffectTime() *string {
	return s.ChangingAffectTime
}

func (s *DescribeDcdnServiceResponseBody) GetChangingChargeType() *string {
	return s.ChangingChargeType
}

func (s *DescribeDcdnServiceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDcdnServiceResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeDcdnServiceResponseBody) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *DescribeDcdnServiceResponseBody) GetOperationLocks() *DescribeDcdnServiceResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *DescribeDcdnServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnServiceResponseBody) GetWebsocketChangingTime() *string {
	return s.WebsocketChangingTime
}

func (s *DescribeDcdnServiceResponseBody) GetWebsocketChangingType() *string {
	return s.WebsocketChangingType
}

func (s *DescribeDcdnServiceResponseBody) GetWebsocketType() *string {
	return s.WebsocketType
}

func (s *DescribeDcdnServiceResponseBody) SetChangingAffectTime(v string) *DescribeDcdnServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetChangingChargeType(v string) *DescribeDcdnServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetInstanceId(v string) *DescribeDcdnServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetInternetChargeType(v string) *DescribeDcdnServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetOpeningTime(v string) *DescribeDcdnServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetOperationLocks(v *DescribeDcdnServiceResponseBodyOperationLocks) *DescribeDcdnServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetRequestId(v string) *DescribeDcdnServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetWebsocketChangingTime(v string) *DescribeDcdnServiceResponseBody {
	s.WebsocketChangingTime = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetWebsocketChangingType(v string) *DescribeDcdnServiceResponseBody {
	s.WebsocketChangingType = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) SetWebsocketType(v string) *DescribeDcdnServiceResponseBody {
	s.WebsocketType = &v
	return s
}

func (s *DescribeDcdnServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeDcdnServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeDcdnServiceResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponseBodyOperationLocks) GetLockReason() []*DescribeDcdnServiceResponseBodyOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeDcdnServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeDcdnServiceResponseBodyOperationLocksLockReason) *DescribeDcdnServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeDcdnServiceResponseBodyOperationLocks) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnServiceResponseBodyOperationLocksLockReason struct {
	// The reason why secure DCDN was locked. For example, a value of financial indicates that an overdue payment exists.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDcdnServiceResponseBodyOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceResponseBodyOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDcdnServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeDcdnServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeDcdnServiceResponseBodyOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}
