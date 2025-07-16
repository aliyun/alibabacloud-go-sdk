// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangingAffectTime(v string) *DescribeCdnServiceResponseBody
	GetChangingAffectTime() *string
	SetChangingChargeType(v string) *DescribeCdnServiceResponseBody
	GetChangingChargeType() *string
	SetChangingDynamicBillingType(v string) *DescribeCdnServiceResponseBody
	GetChangingDynamicBillingType() *string
	SetDynamicBillingType(v string) *DescribeCdnServiceResponseBody
	GetDynamicBillingType() *string
	SetInstanceId(v string) *DescribeCdnServiceResponseBody
	GetInstanceId() *string
	SetInternetChargeType(v string) *DescribeCdnServiceResponseBody
	GetInternetChargeType() *string
	SetOpeningTime(v string) *DescribeCdnServiceResponseBody
	GetOpeningTime() *string
	SetOperationLocks(v *DescribeCdnServiceResponseBodyOperationLocks) *DescribeCdnServiceResponseBody
	GetOperationLocks() *DescribeCdnServiceResponseBodyOperationLocks
	SetRequestId(v string) *DescribeCdnServiceResponseBody
	GetRequestId() *string
}

type DescribeCdnServiceResponseBody struct {
	// The time when the metering method for the next cycle takes effect. The time is displayed in GMT.
	//
	// example:
	//
	// 2019-11-27T16:00:00Z
	ChangingAffectTime *string `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	// The metering method for the next cycle. Valid values:
	//
	// 	- **PayByTraffic**: pay-by-data-transfer
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// example:
	//
	// PayByTraffic
	ChangingChargeType         *string `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	ChangingDynamicBillingType *string `json:"ChangingDynamicBillingType,omitempty" xml:"ChangingDynamicBillingType,omitempty"`
	DynamicBillingType         *string `json:"DynamicBillingType,omitempty" xml:"DynamicBillingType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// aliuidxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The current metering method. Valid values:
	//
	// 	- **PayByTraffic**: pay-by-data-transfer
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The time when the service was activated. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2019-02-28T13:11:49Z
	OpeningTime *string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	// The lock status.
	OperationLocks *DescribeCdnServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBody) GetChangingAffectTime() *string {
	return s.ChangingAffectTime
}

func (s *DescribeCdnServiceResponseBody) GetChangingChargeType() *string {
	return s.ChangingChargeType
}

func (s *DescribeCdnServiceResponseBody) GetChangingDynamicBillingType() *string {
	return s.ChangingDynamicBillingType
}

func (s *DescribeCdnServiceResponseBody) GetDynamicBillingType() *string {
	return s.DynamicBillingType
}

func (s *DescribeCdnServiceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCdnServiceResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeCdnServiceResponseBody) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *DescribeCdnServiceResponseBody) GetOperationLocks() *DescribeCdnServiceResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *DescribeCdnServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnServiceResponseBody) SetChangingAffectTime(v string) *DescribeCdnServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetChangingChargeType(v string) *DescribeCdnServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetChangingDynamicBillingType(v string) *DescribeCdnServiceResponseBody {
	s.ChangingDynamicBillingType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetDynamicBillingType(v string) *DescribeCdnServiceResponseBody {
	s.DynamicBillingType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetInstanceId(v string) *DescribeCdnServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetInternetChargeType(v string) *DescribeCdnServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetOpeningTime(v string) *DescribeCdnServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetOperationLocks(v *DescribeCdnServiceResponseBodyOperationLocks) *DescribeCdnServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeCdnServiceResponseBody) SetRequestId(v string) *DescribeCdnServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeCdnServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeCdnServiceResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBodyOperationLocks) GetLockReason() []*DescribeCdnServiceResponseBodyOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeCdnServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeCdnServiceResponseBodyOperationLocksLockReason) *DescribeCdnServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeCdnServiceResponseBodyOperationLocks) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnServiceResponseBodyOperationLocksLockReason struct {
	// The reason why the service is locked. A value of financial indicates that the service is locked due to overdue payments.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeCdnServiceResponseBodyOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceResponseBodyOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeCdnServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeCdnServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeCdnServiceResponseBodyOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}
