// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsCacheDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*DescribeDnsCacheDomainsResponseBodyDomains) *DescribeDnsCacheDomainsResponseBody
	GetDomains() []*DescribeDnsCacheDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeDnsCacheDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDnsCacheDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDnsCacheDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDnsCacheDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDnsCacheDomainsResponseBody struct {
	// The domain names.
	Domains []*DescribeDnsCacheDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 57121A9C-CDCF-541F-AD39-275D89099420
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDnsCacheDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsCacheDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsCacheDomainsResponseBody) GetDomains() []*DescribeDnsCacheDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDnsCacheDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDnsCacheDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDnsCacheDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsCacheDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDnsCacheDomainsResponseBody) SetDomains(v []*DescribeDnsCacheDomainsResponseBodyDomains) *DescribeDnsCacheDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBody) SetPageNumber(v int64) *DescribeDnsCacheDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBody) SetPageSize(v int64) *DescribeDnsCacheDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBody) SetRequestId(v string) *DescribeDnsCacheDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBody) SetTotalCount(v int64) *DescribeDnsCacheDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBody) Validate() error {
	if s.Domains != nil {
		for _, item := range s.Domains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsCacheDomainsResponseBodyDomains struct {
	// The maximum time-to-live (TTL) period of the cached data retrieved from the origin DNS server. Unit: seconds. Valid values: 30 to 86400.
	//
	// example:
	//
	// 86400
	CacheTtlMax *int32 `json:"CacheTtlMax,omitempty" xml:"CacheTtlMax,omitempty"`
	// The minimum TTL period of the cached data retrieved from the origin DNS server. Unit: seconds. Valid values: 30 to 86400.
	//
	// example:
	//
	// 30
	CacheTtlMin *int32 `json:"CacheTtlMin,omitempty" xml:"CacheTtlMin,omitempty"`
	// The time when the domain name was added. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-17T06:13Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the domain name was added. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1660546144000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The ID of the cache-accelerated domain name.
	//
	// example:
	//
	// 00efd71a-770e-4255-b54e-6fe5659baffe
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The cache-accelerated domain name.
	//
	// example:
	//
	// dns-example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-12-02T16:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1692374400000
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	// The instance ID of the cache-accelerated domain name.
	//
	// example:
	//
	// i-7bg
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The description of the domain name.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The origin DNS servers.
	SourceDnsServers []*DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers `json:"SourceDnsServers,omitempty" xml:"SourceDnsServers,omitempty" type:"Repeated"`
	// Specifies whether the origin Domain Name System (DNS) server supports Extension Mechanisms for DNS (EDNS). Valid values: NOT_SUPPORT and SUPPORT.
	//
	// example:
	//
	// SUPPORT
	SourceEdns *string `json:"SourceEdns,omitempty" xml:"SourceEdns,omitempty"`
	// The origin protocol policy. Valid values: TCP and UDP. Default value: UDP.
	//
	// example:
	//
	// UDP
	SourceProtocol *string `json:"SourceProtocol,omitempty" xml:"SourceProtocol,omitempty"`
	// The time when the configurations of the domain name were updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-22T03:40Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the configurations of the domain name were updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1671690491000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The edition code of Alibaba Cloud DNS.
	//
	// example:
	//
	// ultimate
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s DescribeDnsCacheDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsCacheDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetCacheTtlMax() *int32 {
	return s.CacheTtlMax
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetCacheTtlMin() *int32 {
	return s.CacheTtlMin
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetRemark() *string {
	return s.Remark
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetSourceDnsServers() []*DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers {
	return s.SourceDnsServers
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetSourceEdns() *string {
	return s.SourceEdns
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetSourceProtocol() *string {
	return s.SourceProtocol
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetCacheTtlMax(v int32) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.CacheTtlMax = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetCacheTtlMin(v int32) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.CacheTtlMin = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetCreateTime(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetCreateTimestamp(v int64) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetDomainId(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetDomainName(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.DomainName = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetExpireTime(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetExpireTimestamp(v int64) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetInstanceId(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetRemark(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.Remark = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetSourceDnsServers(v []*DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.SourceDnsServers = v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetSourceEdns(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.SourceEdns = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetSourceProtocol(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.SourceProtocol = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetUpdateTime(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetUpdateTimestamp(v int64) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) SetVersionCode(v string) *DescribeDnsCacheDomainsResponseBodyDomains {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomains) Validate() error {
	if s.SourceDnsServers != nil {
		for _, item := range s.SourceDnsServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers struct {
	// The domain name or IP address of the origin DNS server.
	//
	// example:
	//
	// ns8.alidns.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port of the origin DNS server.
	//
	// example:
	//
	// 53
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) GoString() string {
	return s.String()
}

func (s *DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) GetHost() *string {
	return s.Host
}

func (s *DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) GetPort() *string {
	return s.Port
}

func (s *DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) SetHost(v string) *DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers {
	s.Host = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) SetPort(v string) *DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers {
	s.Port = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponseBodyDomainsSourceDnsServers) Validate() error {
	return dara.Validate(s)
}
