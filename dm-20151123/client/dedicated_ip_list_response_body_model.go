// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DedicatedIpListResponseBody
	GetCurrentPage() *int32
	SetHasMore(v bool) *DedicatedIpListResponseBody
	GetHasMore() *bool
	SetIps(v []*DedicatedIpListResponseBodyIps) *DedicatedIpListResponseBody
	GetIps() []*DedicatedIpListResponseBodyIps
	SetPageSize(v int32) *DedicatedIpListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DedicatedIpListResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *DedicatedIpListResponseBody
	GetTotalCounts() *int32
}

type DedicatedIpListResponseBody struct {
	// Current page
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Whether there is a next page
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// IP list
	Ips []*DedicatedIpListResponseBodyIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total amount of purchased IP data
	//
	// example:
	//
	// 5
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s DedicatedIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DedicatedIpListResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *DedicatedIpListResponseBody) GetIps() []*DedicatedIpListResponseBodyIps {
	return s.Ips
}

func (s *DedicatedIpListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DedicatedIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpListResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *DedicatedIpListResponseBody) SetCurrentPage(v int32) *DedicatedIpListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DedicatedIpListResponseBody) SetHasMore(v bool) *DedicatedIpListResponseBody {
	s.HasMore = &v
	return s
}

func (s *DedicatedIpListResponseBody) SetIps(v []*DedicatedIpListResponseBodyIps) *DedicatedIpListResponseBody {
	s.Ips = v
	return s
}

func (s *DedicatedIpListResponseBody) SetPageSize(v int32) *DedicatedIpListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DedicatedIpListResponseBody) SetRequestId(v string) *DedicatedIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpListResponseBody) SetTotalCounts(v int32) *DedicatedIpListResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *DedicatedIpListResponseBody) Validate() error {
	if s.Ips != nil {
		for _, item := range s.Ips {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DedicatedIpListResponseBodyIps struct {
	// Expiration time
	//
	// example:
	//
	// 2025-06-12T09:19:20Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// IP ID, consistent with the purchased instance ID
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Purchased instance ID
	//
	// example:
	//
	// xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IP address
	//
	// example:
	//
	// xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Extended information
	IpExt *DedicatedIpListResponseBodyIpsIpExt `json:"IpExt,omitempty" xml:"IpExt,omitempty" type:"Struct"`
	// Name of the IP pool
	//
	// example:
	//
	// xxx
	IpPoolName *string `json:"IpPoolName,omitempty" xml:"IpPoolName,omitempty"`
	// Purchase time
	//
	// example:
	//
	// 2025-05-12T09:19:20Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// IP status
	//
	// example:
	//
	// sold
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Warm-up status
	//
	// example:
	//
	// finish
	WarmupStatus *string `json:"WarmupStatus,omitempty" xml:"WarmupStatus,omitempty"`
	// Warm-up method
	//
	// example:
	//
	// cusSelfManager
	WarmupType *string `json:"WarmupType,omitempty" xml:"WarmupType,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DedicatedIpListResponseBodyIps) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpListResponseBodyIps) GoString() string {
	return s.String()
}

func (s *DedicatedIpListResponseBodyIps) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DedicatedIpListResponseBodyIps) GetId() *string {
	return s.Id
}

func (s *DedicatedIpListResponseBodyIps) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DedicatedIpListResponseBodyIps) GetIp() *string {
	return s.Ip
}

func (s *DedicatedIpListResponseBodyIps) GetIpExt() *DedicatedIpListResponseBodyIpsIpExt {
	return s.IpExt
}

func (s *DedicatedIpListResponseBodyIps) GetIpPoolName() *string {
	return s.IpPoolName
}

func (s *DedicatedIpListResponseBodyIps) GetStartTime() *string {
	return s.StartTime
}

func (s *DedicatedIpListResponseBodyIps) GetStatus() *string {
	return s.Status
}

func (s *DedicatedIpListResponseBodyIps) GetWarmupStatus() *string {
	return s.WarmupStatus
}

func (s *DedicatedIpListResponseBodyIps) GetWarmupType() *string {
	return s.WarmupType
}

func (s *DedicatedIpListResponseBodyIps) GetZoneId() *string {
	return s.ZoneId
}

func (s *DedicatedIpListResponseBodyIps) SetExpiredTime(v string) *DedicatedIpListResponseBodyIps {
	s.ExpiredTime = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetId(v string) *DedicatedIpListResponseBodyIps {
	s.Id = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetInstanceId(v string) *DedicatedIpListResponseBodyIps {
	s.InstanceId = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetIp(v string) *DedicatedIpListResponseBodyIps {
	s.Ip = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetIpExt(v *DedicatedIpListResponseBodyIpsIpExt) *DedicatedIpListResponseBodyIps {
	s.IpExt = v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetIpPoolName(v string) *DedicatedIpListResponseBodyIps {
	s.IpPoolName = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetStartTime(v string) *DedicatedIpListResponseBodyIps {
	s.StartTime = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetStatus(v string) *DedicatedIpListResponseBodyIps {
	s.Status = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetWarmupStatus(v string) *DedicatedIpListResponseBodyIps {
	s.WarmupStatus = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetWarmupType(v string) *DedicatedIpListResponseBodyIps {
	s.WarmupType = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) SetZoneId(v string) *DedicatedIpListResponseBodyIps {
	s.ZoneId = &v
	return s
}

func (s *DedicatedIpListResponseBodyIps) Validate() error {
	if s.IpExt != nil {
		if err := s.IpExt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DedicatedIpListResponseBodyIpsIpExt struct {
	// Whether auto-renewal is enabled
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// Whether an email has been sent
	//
	// example:
	//
	// true
	HasSendMail               *bool   `json:"HasSendMail,omitempty" xml:"HasSendMail,omitempty"`
	LastWarmUpTypeChangedTime *string `json:"LastWarmUpTypeChangedTime,omitempty" xml:"LastWarmUpTypeChangedTime,omitempty"`
}

func (s DedicatedIpListResponseBodyIpsIpExt) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpListResponseBodyIpsIpExt) GoString() string {
	return s.String()
}

func (s *DedicatedIpListResponseBodyIpsIpExt) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DedicatedIpListResponseBodyIpsIpExt) GetHasSendMail() *bool {
	return s.HasSendMail
}

func (s *DedicatedIpListResponseBodyIpsIpExt) GetLastWarmUpTypeChangedTime() *string {
	return s.LastWarmUpTypeChangedTime
}

func (s *DedicatedIpListResponseBodyIpsIpExt) SetAutoRenewal(v bool) *DedicatedIpListResponseBodyIpsIpExt {
	s.AutoRenewal = &v
	return s
}

func (s *DedicatedIpListResponseBodyIpsIpExt) SetHasSendMail(v bool) *DedicatedIpListResponseBodyIpsIpExt {
	s.HasSendMail = &v
	return s
}

func (s *DedicatedIpListResponseBodyIpsIpExt) SetLastWarmUpTypeChangedTime(v string) *DedicatedIpListResponseBodyIpsIpExt {
	s.LastWarmUpTypeChangedTime = &v
	return s
}

func (s *DedicatedIpListResponseBodyIpsIpExt) Validate() error {
	return dara.Validate(s)
}
