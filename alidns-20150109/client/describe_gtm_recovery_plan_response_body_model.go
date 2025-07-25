// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeGtmRecoveryPlanResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody
	GetCreateTimestamp() *int64
	SetFaultAddrPoolNum(v int32) *DescribeGtmRecoveryPlanResponseBody
	GetFaultAddrPoolNum() *int32
	SetFaultAddrPools(v *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) *DescribeGtmRecoveryPlanResponseBody
	GetFaultAddrPools() *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools
	SetLastExecuteTime(v string) *DescribeGtmRecoveryPlanResponseBody
	GetLastExecuteTime() *string
	SetLastExecuteTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody
	GetLastExecuteTimestamp() *int64
	SetLastRollbackTime(v string) *DescribeGtmRecoveryPlanResponseBody
	GetLastRollbackTime() *string
	SetLastRollbackTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody
	GetLastRollbackTimestamp() *int64
	SetName(v string) *DescribeGtmRecoveryPlanResponseBody
	GetName() *string
	SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlanResponseBody
	GetRecoveryPlanId() *int64
	SetRemark(v string) *DescribeGtmRecoveryPlanResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeGtmRecoveryPlanResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeGtmRecoveryPlanResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *DescribeGtmRecoveryPlanResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeGtmRecoveryPlanResponseBody struct {
	// The time when the disaster recovery plan was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-11T05:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the disaster recovery plan was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1565499867000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The number of faulty address pools.
	//
	// example:
	//
	// 2
	FaultAddrPoolNum *int32 `json:"FaultAddrPoolNum,omitempty" xml:"FaultAddrPoolNum,omitempty"`
	// The faulty address pools.
	FaultAddrPools *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools `json:"FaultAddrPools,omitempty" xml:"FaultAddrPools,omitempty" type:"Struct"`
	// The time when the disaster recovery plan was last executed. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-11T05:04Z
	LastExecuteTime *string `json:"LastExecuteTime,omitempty" xml:"LastExecuteTime,omitempty"`
	// The time when the disaster recovery plan was last executed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1565505898000
	LastExecuteTimestamp *int64 `json:"LastExecuteTimestamp,omitempty" xml:"LastExecuteTimestamp,omitempty"`
	// The time when the disaster recovery plan was last rolled back. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 2019-08-11T06:45Z
	LastRollbackTime *string `json:"LastRollbackTime,omitempty" xml:"LastRollbackTime,omitempty"`
	// The time when the disaster recovery plan was last rolled back. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1565505919000
	LastRollbackTimestamp *int64 `json:"LastRollbackTimestamp,omitempty" xml:"LastRollbackTimestamp,omitempty"`
	// The name of the disaster recovery plan.
	//
	// example:
	//
	// name-example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the disaster recovery plan.
	//
	// example:
	//
	// 55
	RecoveryPlanId *int64 `json:"RecoveryPlanId,omitempty" xml:"RecoveryPlanId,omitempty"`
	// The description of the disaster recovery plan.
	//
	// example:
	//
	// remark-example
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A5F4315-D6E8-435E-82DF-24F4C97D6999
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the disaster recovery plan.
	//
	// example:
	//
	// UNEXECUTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the disaster recovery plan was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-11T06:45Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the disaster recovery plan was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1565499867000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeGtmRecoveryPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetFaultAddrPoolNum() *int32 {
	return s.FaultAddrPoolNum
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetFaultAddrPools() *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools {
	return s.FaultAddrPools
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetLastExecuteTime() *string {
	return s.LastExecuteTime
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetLastExecuteTimestamp() *int64 {
	return s.LastExecuteTimestamp
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetLastRollbackTime() *string {
	return s.LastRollbackTime
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetLastRollbackTimestamp() *int64 {
	return s.LastRollbackTimestamp
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetRecoveryPlanId() *int64 {
	return s.RecoveryPlanId
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeGtmRecoveryPlanResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetCreateTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetCreateTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetFaultAddrPoolNum(v int32) *DescribeGtmRecoveryPlanResponseBody {
	s.FaultAddrPoolNum = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetFaultAddrPools(v *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) *DescribeGtmRecoveryPlanResponseBody {
	s.FaultAddrPools = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastExecuteTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.LastExecuteTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastExecuteTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.LastExecuteTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastRollbackTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.LastRollbackTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetLastRollbackTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.LastRollbackTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetName(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetRecoveryPlanId(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.RecoveryPlanId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetRemark(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetRequestId(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetStatus(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetUpdateTime(v string) *DescribeGtmRecoveryPlanResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) SetUpdateTimestamp(v int64) *DescribeGtmRecoveryPlanResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmRecoveryPlanResponseBodyFaultAddrPools struct {
	FaultAddrPool []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool `json:"FaultAddrPool,omitempty" xml:"FaultAddrPool,omitempty" type:"Repeated"`
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) GetFaultAddrPool() []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool {
	return s.FaultAddrPool
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) SetFaultAddrPool(v []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools {
	s.FaultAddrPool = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPools) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool struct {
	// The address pool ID.
	//
	// example:
	//
	// hra0oq
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The address pool name.
	AddrPoolName *string                                                              `json:"AddrPoolName,omitempty" xml:"AddrPoolName,omitempty"`
	Addrs        *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs `json:"Addrs,omitempty" xml:"Addrs,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// instance-zwy-38
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) GetAddrPoolName() *string {
	return s.AddrPoolName
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) GetAddrs() *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs {
	return s.Addrs
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) SetAddrPoolId(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) SetAddrPoolName(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool {
	s.AddrPoolName = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) SetAddrs(v *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool {
	s.Addrs = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) SetInstanceId(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs struct {
	Addr []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs) GetAddr() []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr {
	return s.Addr
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs) SetAddr(v []*DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs {
	s.Addr = v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrs) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr struct {
	// The address ID.
	//
	// example:
	//
	// 739
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The address mode.
	//
	// example:
	//
	// OFFLINE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The address.
	//
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) GetId() *int64 {
	return s.Id
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) GetMode() *string {
	return s.Mode
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) GetValue() *string {
	return s.Value
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) SetId(v int64) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr {
	s.Id = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) SetMode(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr {
	s.Mode = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) SetValue(v string) *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr {
	s.Value = &v
	return s
}

func (s *DescribeGtmRecoveryPlanResponseBodyFaultAddrPoolsFaultAddrPoolAddrsAddr) Validate() error {
	return dara.Validate(s)
}
