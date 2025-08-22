// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangingAffectTime(v string) *DescribeDcdnIpaServiceResponseBody
	GetChangingAffectTime() *string
	SetChangingChargeType(v string) *DescribeDcdnIpaServiceResponseBody
	GetChangingChargeType() *string
	SetInstanceId(v string) *DescribeDcdnIpaServiceResponseBody
	GetInstanceId() *string
	SetInternetChargeType(v string) *DescribeDcdnIpaServiceResponseBody
	GetInternetChargeType() *string
	SetOpeningTime(v string) *DescribeDcdnIpaServiceResponseBody
	GetOpeningTime() *string
	SetOperationLocks(v *DescribeDcdnIpaServiceResponseBodyOperationLocks) *DescribeDcdnIpaServiceResponseBody
	GetOperationLocks() *DescribeDcdnIpaServiceResponseBodyOperationLocks
	SetRequestId(v string) *DescribeDcdnIpaServiceResponseBody
	GetRequestId() *string
}

type DescribeDcdnIpaServiceResponseBody struct {
	// The time when the change of the billing method starts to take effect. The time is in GMT. This time appears on the frontend only when it is later than the current time.
	//
	// example:
	//
	// 2018-03-31T16:00:00Z
	ChangingAffectTime *string `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	// The new billing method to take effect. Valid values:
	//
	// 	- **PayByTraffic**: pay-by-data-transfer
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// 	- **PayByBandwidth95**: pay-by-95th percentile bandwidth
	//
	// 	- **PayByBandwidth_monthavg**: pay-by-monthly average bandwidth
	//
	// 	- **PayByBandwidth_month4th**: pay-by-fourth peak bandwidth per month
	//
	// 	- **PayByBandwidth_monthday95avg**: pay-by-monthly average 95th percentile bandwidth
	//
	// 	- **PayByBandwidth_nighthalf95**: pay-by-95th percentile bandwidth (50% off during nighttime)
	//
	// example:
	//
	// PayByBandwidth
	ChangingChargeType *string `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 1883927335936173
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PayByTraffic**: pay-by-data-transfer
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// 	- **PayByBandwidth95**: pay-by-95th percentile bandwidth
	//
	// 	- **PayByBandwidth_monthavg**: pay-by-monthly average bandwidth
	//
	// 	- **PayByBandwidth_month4th**: pay-by-fourth peak bandwidth per month
	//
	// 	- **PayByBandwidth_monthday95avg**: pay-by-monthly average 95th percentile bandwidth
	//
	// 	- **PayByBandwidth_nighthalf95**: pay-by-95th percentile bandwidth (50% off during nighttime)
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The time when the DCDN service was activated. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2018-03-19T11:16:11Z
	OpeningTime *string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	// The lock status of secure DCDN.
	OperationLocks *DescribeDcdnIpaServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EF2AEBC2-EDBD-41CF-BF64-7E095D42D6EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpaServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponseBody) GetChangingAffectTime() *string {
	return s.ChangingAffectTime
}

func (s *DescribeDcdnIpaServiceResponseBody) GetChangingChargeType() *string {
	return s.ChangingChargeType
}

func (s *DescribeDcdnIpaServiceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDcdnIpaServiceResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeDcdnIpaServiceResponseBody) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *DescribeDcdnIpaServiceResponseBody) GetOperationLocks() *DescribeDcdnIpaServiceResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *DescribeDcdnIpaServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnIpaServiceResponseBody) SetChangingAffectTime(v string) *DescribeDcdnIpaServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetChangingChargeType(v string) *DescribeDcdnIpaServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetInstanceId(v string) *DescribeDcdnIpaServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetInternetChargeType(v string) *DescribeDcdnIpaServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetOpeningTime(v string) *DescribeDcdnIpaServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetOperationLocks(v *DescribeDcdnIpaServiceResponseBodyOperationLocks) *DescribeDcdnIpaServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) SetRequestId(v string) *DescribeDcdnIpaServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocks) GetLockReason() []*DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) *DescribeDcdnIpaServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocks) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason struct {
	// The reason why the instance is locked. For example, a value of **financial*	- indicates that an overdue payment exists.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponseBodyOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}
