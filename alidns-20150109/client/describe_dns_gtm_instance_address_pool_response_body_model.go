// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrCount(v int32) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetAddrCount() *int32
	SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetAddrPoolId() *string
	SetAddrs(v *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetAddrs() *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs
	SetCreateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetCreateTimestamp() *int64
	SetLbaStrategy(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetLbaStrategy() *string
	SetMonitorConfigId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetMonitorConfigId() *string
	SetMonitorStatus(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetMonitorStatus() *string
	SetName(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetRequestId() *string
	SetType(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetType() *string
	SetUpdateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeDnsGtmInstanceAddressPoolResponseBody struct {
	// The number of addresses in the address pool.
	//
	// example:
	//
	// 1
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// testpool1
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The addresses in the address pool.
	Addrs *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs `json:"Addrs,omitempty" xml:"Addrs,omitempty" type:"Struct"`
	// The time when the address pool was created.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates the time when the address pool was created.
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The load balancing policy for the address pool. Valid values:
	//
	// 	- ALL_RR: returns all addresses.
	//
	// 	- RATIO: returns addresses by weight.
	//
	// example:
	//
	// all_rr
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The ID of the health check configuration.
	//
	// example:
	//
	// test1
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// Indicates the status of the health check. Valid values:
	//
	// 	- OPEN: The health check is enabled.
	//
	// 	- CLOSE: The health check is disabled.
	//
	// 	- UNCONFIGURED: The health check is not configured.
	//
	// example:
	//
	// open
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- IPV4: IPv4 address
	//
	// 	- IPV6: IPv6 address
	//
	// 	- DOMAIN: domain name
	//
	// example:
	//
	// ipv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the address pool was last updated.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The timestamp that indicates the time when the address pool was last updated.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetAddrs() *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	return s.Addrs
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetLbaStrategy() *string {
	return s.LbaStrategy
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetAddrCount(v int32) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetAddrs(v *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.Addrs = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetCreateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetLbaStrategy(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetMonitorConfigId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetMonitorStatus(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetName(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetType(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetUpdateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) SetUpdateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBody) Validate() error {
	if s.Addrs != nil {
		if err := s.Addrs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs struct {
	Addr []*DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) GetAddr() []*DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	return s.Addr
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) SetAddr(v []*DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs {
	s.Addr = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrs) Validate() error {
	if s.Addr != nil {
		for _, item := range s.Addr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr struct {
	// The address.
	//
	// example:
	//
	// 1.1.1.1
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// The status of the last health check on the address. Valid values:
	//
	// 	- OK: No active alerts are triggered.
	//
	// 	- ALERT: Alerts are triggered based on the alert rules.
	//
	// example:
	//
	// ok
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The source region of the address.
	//
	// 	- lineCode: the line code of the source region of the address. This parameter is deprecated, and lineCodes prevails.
	//
	// 	- lineName: the line name of the source region of the address. This parameter is deprecated.
	//
	// 	- lineCodes: the line codes of the source regions of the address.
	//
	// example:
	//
	// "lineCode":"aliyun_r_cn-zhangjiakou", "lineName": "Alibaba Cloud_China (Zhangjiakou)", "lineCodes": ["aliyun_r_cn-zhangjiakou"]
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The time when the address was added into the address pool.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates the time when the address was added into the address pool.
	//
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
	// 	- SMART: smart return
	//
	// 	- ONLINE: always online
	//
	// 	- OFFLINE: always offline
	//
	// example:
	//
	// online
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The description of the address.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The time when the address was last updated.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The timestamp that indicates the time when the address was last updated.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetAddr() *string {
	return s.Addr
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetMode() *string {
	return s.Mode
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetRemark() *string {
	return s.Remark
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetAddr(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.Addr = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetAlertStatus(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.AlertStatus = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetAttributeInfo(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.AttributeInfo = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetCreateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetLbaWeight(v int32) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetMode(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.Mode = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetRemark(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.Remark = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetUpdateTime(v string) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) SetUpdateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponseBodyAddrsAddr) Validate() error {
	return dara.Validate(s)
}
