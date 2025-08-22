// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnsecServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangingAffectTime(v string) *DescribeDcdnsecServiceResponseBody
	GetChangingAffectTime() *string
	SetChangingChargeType(v string) *DescribeDcdnsecServiceResponseBody
	GetChangingChargeType() *string
	SetDomainNum(v string) *DescribeDcdnsecServiceResponseBody
	GetDomainNum() *string
	SetEndTime(v string) *DescribeDcdnsecServiceResponseBody
	GetEndTime() *string
	SetFlowType(v string) *DescribeDcdnsecServiceResponseBody
	GetFlowType() *string
	SetInstanceId(v string) *DescribeDcdnsecServiceResponseBody
	GetInstanceId() *string
	SetInternetChargeType(v string) *DescribeDcdnsecServiceResponseBody
	GetInternetChargeType() *string
	SetOperationLocks(v *DescribeDcdnsecServiceResponseBodyOperationLocks) *DescribeDcdnsecServiceResponseBody
	GetOperationLocks() *DescribeDcdnsecServiceResponseBodyOperationLocks
	SetRequestId(v string) *DescribeDcdnsecServiceResponseBody
	GetRequestId() *string
	SetRequestType(v string) *DescribeDcdnsecServiceResponseBody
	GetRequestType() *string
	SetStartTime(v string) *DescribeDcdnsecServiceResponseBody
	GetStartTime() *string
	SetVersion(v string) *DescribeDcdnsecServiceResponseBody
	GetVersion() *string
}

type DescribeDcdnsecServiceResponseBody struct {
	// The time when the renewed service takes effect. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-09-30T16:00:00Z
	ChangingAffectTime *string `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	// The new metering method for the renewed DCDN. Valid values:
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
	// The number of accelerated domain names that use DCDN.
	//
	// example:
	//
	// 130
	DomainNum *string `json:"DomainNum,omitempty" xml:"DomainNum,omitempty"`
	// The service expiration time.
	//
	// example:
	//
	// 2021-09-26T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metering method for traffic.
	//
	// example:
	//
	// PayBySecTraffic
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// dcdn_dcdnsec_public_cn-123***
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
	// The lock status of DCDN.
	OperationLocks *DescribeDcdnsecServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E20B46E1-9BCD-10E5-AAEF-6D7B737467A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metering method for requests.
	//
	// example:
	//
	// PayBySecRequest
	RequestType *string `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	// The service activation time.
	//
	// example:
	//
	// 2021-08-26T02:52:08Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// enterprise
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDcdnsecServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnsecServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponseBody) GetChangingAffectTime() *string {
	return s.ChangingAffectTime
}

func (s *DescribeDcdnsecServiceResponseBody) GetChangingChargeType() *string {
	return s.ChangingChargeType
}

func (s *DescribeDcdnsecServiceResponseBody) GetDomainNum() *string {
	return s.DomainNum
}

func (s *DescribeDcdnsecServiceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnsecServiceResponseBody) GetFlowType() *string {
	return s.FlowType
}

func (s *DescribeDcdnsecServiceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDcdnsecServiceResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeDcdnsecServiceResponseBody) GetOperationLocks() *DescribeDcdnsecServiceResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *DescribeDcdnsecServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnsecServiceResponseBody) GetRequestType() *string {
	return s.RequestType
}

func (s *DescribeDcdnsecServiceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnsecServiceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeDcdnsecServiceResponseBody) SetChangingAffectTime(v string) *DescribeDcdnsecServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetChangingChargeType(v string) *DescribeDcdnsecServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetDomainNum(v string) *DescribeDcdnsecServiceResponseBody {
	s.DomainNum = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetEndTime(v string) *DescribeDcdnsecServiceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetFlowType(v string) *DescribeDcdnsecServiceResponseBody {
	s.FlowType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetInstanceId(v string) *DescribeDcdnsecServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetInternetChargeType(v string) *DescribeDcdnsecServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetOperationLocks(v *DescribeDcdnsecServiceResponseBodyOperationLocks) *DescribeDcdnsecServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetRequestId(v string) *DescribeDcdnsecServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetRequestType(v string) *DescribeDcdnsecServiceResponseBody {
	s.RequestType = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetStartTime(v string) *DescribeDcdnsecServiceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) SetVersion(v string) *DescribeDcdnsecServiceResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnsecServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeDcdnsecServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocks) GetLockReason() []*DescribeDcdnsecServiceResponseBodyOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) *DescribeDcdnsecServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocks) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnsecServiceResponseBodyOperationLocksLockReason struct {
	// The reason why the instance was locked.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeDcdnsecServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeDcdnsecServiceResponseBodyOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}
