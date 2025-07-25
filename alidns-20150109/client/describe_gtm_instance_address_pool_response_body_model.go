// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrCount(v int32) *DescribeGtmInstanceAddressPoolResponseBody
	GetAddrCount() *int32
	SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetAddrPoolId() *string
	SetAddrs(v *DescribeGtmInstanceAddressPoolResponseBodyAddrs) *DescribeGtmInstanceAddressPoolResponseBody
	GetAddrs() *DescribeGtmInstanceAddressPoolResponseBodyAddrs
	SetCreateTime(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBody
	GetCreateTimestamp() *int64
	SetMinAvailableAddrNum(v int32) *DescribeGtmInstanceAddressPoolResponseBody
	GetMinAvailableAddrNum() *int32
	SetMonitorConfigId(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetMonitorConfigId() *string
	SetMonitorStatus(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetMonitorStatus() *string
	SetName(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetStatus() *string
	SetType(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetType() *string
	SetUpdateTime(v string) *DescribeGtmInstanceAddressPoolResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeGtmInstanceAddressPoolResponseBody struct {
	// The number of addresses in the address pool queried.
	//
	// example:
	//
	// 2
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// 1234abc
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The addresses in the address pool.
	Addrs *DescribeGtmInstanceAddressPoolResponseBodyAddrs `json:"Addrs,omitempty" xml:"Addrs,omitempty" type:"Struct"`
	// The time when the address pool was created.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The minimum number of available addresses in the address pool.
	//
	// example:
	//
	// 2
	MinAvailableAddrNum *int32 `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	// The health check ID of the address pool.
	//
	// example:
	//
	// 100abc
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// Indicates whether health check was enabled for the address pool. Valid values:
	//
	// 	- **OPEN**: Enabled
	//
	// 	- **CLOSE**: Disabled
	//
	// 	- **UNCONFIGURED**: Not configured
	//
	// example:
	//
	// OPEN
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// The name of the address pool.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The availability status of the address pool. Valid values:
	//
	// 	- **AVAILABLE**: Available
	//
	// 	- **NOT_AVAILABLE**: Unavailable
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- **IP**: IP address
	//
	// 	- **DOMAIN**: Domain name
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The last time when the address pool was updated.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// A timestamp that indicates the last time the address pool was updated.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetAddrs() *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	return s.Addrs
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetMinAvailableAddrNum() *int32 {
	return s.MinAvailableAddrNum
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetAddrCount(v int32) *DescribeGtmInstanceAddressPoolResponseBody {
	s.AddrCount = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetAddrs(v *DescribeGtmInstanceAddressPoolResponseBodyAddrs) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Addrs = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetCreateTime(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetCreateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetMinAvailableAddrNum(v int32) *DescribeGtmInstanceAddressPoolResponseBody {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetMonitorConfigId(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetMonitorStatus(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetName(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetRequestId(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetStatus(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetType(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetUpdateTime(v string) *DescribeGtmInstanceAddressPoolResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) SetUpdateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmInstanceAddressPoolResponseBodyAddrs struct {
	Addr []*DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
}

func (s DescribeGtmInstanceAddressPoolResponseBodyAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolResponseBodyAddrs) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) GetAddr() []*DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	return s.Addr
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) SetAddr(v []*DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) *DescribeGtmInstanceAddressPoolResponseBodyAddrs {
	s.Addr = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrs) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr struct {
	// The ID of the address.
	//
	// example:
	//
	// 123
	AddrId *int64 `json:"AddrId,omitempty" xml:"AddrId,omitempty"`
	// Indicates whether health check was enabled for the address. Valid values:
	//
	// 	- **OK**: Normal
	//
	// 	- **ALERT**: Alert
	//
	// example:
	//
	// OK
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The time when the address pool was created.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The weight of the address.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The mode of the address. Valid values:
	//
	// 	- **SMART**: Intelligent return
	//
	// 	- **ONLINE**: Always online
	//
	// 	- **OFFLINE**: Always offline
	//
	// example:
	//
	// SMART
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The last time when the address was updated.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// A timestamp that indicates the last time when the address was updated.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The address.
	//
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetAddrId() *int64 {
	return s.AddrId
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetMode() *string {
	return s.Mode
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) GetValue() *string {
	return s.Value
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetAddrId(v int64) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.AddrId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetAlertStatus(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.AlertStatus = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetCreateTime(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetCreateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetLbaWeight(v int32) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.LbaWeight = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetMode(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.Mode = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetUpdateTime(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetUpdateTimestamp(v int64) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) SetValue(v string) *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.Value = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponseBodyAddrsAddr) Validate() error {
	return dara.Validate(s)
}
