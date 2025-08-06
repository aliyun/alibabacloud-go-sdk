// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangingAffectTime(v string) *DescribeVodServiceResponseBody
	GetChangingAffectTime() *string
	SetChangingChargeType(v string) *DescribeVodServiceResponseBody
	GetChangingChargeType() *string
	SetInstanceId(v string) *DescribeVodServiceResponseBody
	GetInstanceId() *string
	SetInternetChargeType(v string) *DescribeVodServiceResponseBody
	GetInternetChargeType() *string
	SetOpeningTime(v string) *DescribeVodServiceResponseBody
	GetOpeningTime() *string
	SetOperationLocks(v *DescribeVodServiceResponseBodyOperationLocks) *DescribeVodServiceResponseBody
	GetOperationLocks() *DescribeVodServiceResponseBodyOperationLocks
	SetRequestId(v string) *DescribeVodServiceResponseBody
	GetRequestId() *string
}

type DescribeVodServiceResponseBody struct {
	ChangingAffectTime *string                                       `json:"ChangingAffectTime,omitempty" xml:"ChangingAffectTime,omitempty"`
	ChangingChargeType *string                                       `json:"ChangingChargeType,omitempty" xml:"ChangingChargeType,omitempty"`
	InstanceId         *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetChargeType *string                                       `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OpeningTime        *string                                       `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	OperationLocks     *DescribeVodServiceResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodServiceResponseBody) GetChangingAffectTime() *string {
	return s.ChangingAffectTime
}

func (s *DescribeVodServiceResponseBody) GetChangingChargeType() *string {
	return s.ChangingChargeType
}

func (s *DescribeVodServiceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVodServiceResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeVodServiceResponseBody) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *DescribeVodServiceResponseBody) GetOperationLocks() *DescribeVodServiceResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *DescribeVodServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodServiceResponseBody) SetChangingAffectTime(v string) *DescribeVodServiceResponseBody {
	s.ChangingAffectTime = &v
	return s
}

func (s *DescribeVodServiceResponseBody) SetChangingChargeType(v string) *DescribeVodServiceResponseBody {
	s.ChangingChargeType = &v
	return s
}

func (s *DescribeVodServiceResponseBody) SetInstanceId(v string) *DescribeVodServiceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeVodServiceResponseBody) SetInternetChargeType(v string) *DescribeVodServiceResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeVodServiceResponseBody) SetOpeningTime(v string) *DescribeVodServiceResponseBody {
	s.OpeningTime = &v
	return s
}

func (s *DescribeVodServiceResponseBody) SetOperationLocks(v *DescribeVodServiceResponseBodyOperationLocks) *DescribeVodServiceResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeVodServiceResponseBody) SetRequestId(v string) *DescribeVodServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodServiceResponseBodyOperationLocks struct {
	LockReason []*DescribeVodServiceResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeVodServiceResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodServiceResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeVodServiceResponseBodyOperationLocks) GetLockReason() []*DescribeVodServiceResponseBodyOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeVodServiceResponseBodyOperationLocks) SetLockReason(v []*DescribeVodServiceResponseBodyOperationLocksLockReason) *DescribeVodServiceResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeVodServiceResponseBodyOperationLocks) Validate() error {
	return dara.Validate(s)
}

type DescribeVodServiceResponseBodyOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeVodServiceResponseBodyOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodServiceResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeVodServiceResponseBodyOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeVodServiceResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeVodServiceResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeVodServiceResponseBodyOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}
