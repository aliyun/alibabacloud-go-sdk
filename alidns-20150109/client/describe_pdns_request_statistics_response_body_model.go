// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsRequestStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribePdnsRequestStatisticsResponseBodyData) *DescribePdnsRequestStatisticsResponseBody
	GetData() []*DescribePdnsRequestStatisticsResponseBodyData
	SetPageNumber(v int64) *DescribePdnsRequestStatisticsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsRequestStatisticsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePdnsRequestStatisticsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePdnsRequestStatisticsResponseBody
	GetTotalCount() *int64
}

type DescribePdnsRequestStatisticsResponseBody struct {
	// The statistics on the DNS requests.
	Data []*DescribePdnsRequestStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **20**. Valid values: **1 to 100**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 49
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePdnsRequestStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticsResponseBody) GetData() []*DescribePdnsRequestStatisticsResponseBodyData {
	return s.Data
}

func (s *DescribePdnsRequestStatisticsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsRequestStatisticsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsRequestStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsRequestStatisticsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsRequestStatisticsResponseBody) SetData(v []*DescribePdnsRequestStatisticsResponseBodyData) *DescribePdnsRequestStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBody) SetPageNumber(v int64) *DescribePdnsRequestStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBody) SetPageSize(v int64) *DescribePdnsRequestStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBody) SetRequestId(v string) *DescribePdnsRequestStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBody) SetTotalCount(v int64) *DescribePdnsRequestStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePdnsRequestStatisticsResponseBodyData struct {
	// The total number of DoH requests, including the HTTP and HTTPS requests.
	//
	// example:
	//
	// 0
	DohTotalCount *int64 `json:"DohTotalCount,omitempty" xml:"DohTotalCount,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The number of HTTP requests.
	//
	// example:
	//
	// 0
	HttpCount *int64 `json:"HttpCount,omitempty" xml:"HttpCount,omitempty"`
	// The number of HTTPS requests. On the Traffic Analysis tab of the public DNS console, the value of this parameter includes the number of DNS over HTTPs (DoH) requests. Therefore, the number of DoH requests is not separately displayed in the console.
	//
	// example:
	//
	// 0
	HttpsCount *int64 `json:"HttpsCount,omitempty" xml:"HttpsCount,omitempty"`
	// The number of source IP addresses.
	//
	// example:
	//
	// 10
	IpCount *int64 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The current version does not support this parameter.
	//
	// example:
	//
	// -
	MaxThreatLevel *string `json:"MaxThreatLevel,omitempty" xml:"MaxThreatLevel,omitempty"`
	// The subdomain name.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The current version does not support this parameter.
	//
	// example:
	//
	// -
	ThreatCount *int64 `json:"ThreatCount,omitempty" xml:"ThreatCount,omitempty"`
	// The current version does not support this parameter.
	ThreatInfo []*DescribePdnsRequestStatisticsResponseBodyDataThreatInfo `json:"ThreatInfo,omitempty" xml:"ThreatInfo,omitempty" type:"Repeated"`
	// The total number of requests.
	//
	// example:
	//
	// 500
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of UDP requests.
	//
	// example:
	//
	// 500
	UdpTotalCount *int64 `json:"UdpTotalCount,omitempty" xml:"UdpTotalCount,omitempty"`
	// The number of IPv4-based requests.
	//
	// example:
	//
	// 0
	V4Count *int64 `json:"V4Count,omitempty" xml:"V4Count,omitempty"`
	// The number of IPv4-based HTTP requests.
	//
	// example:
	//
	// 0
	V4HttpCount *int64 `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	// The number of IPv4-based HTTPS requests.
	//
	// example:
	//
	// 0
	V4HttpsCount *int64 `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	// The number of IPv6-based requests.
	//
	// example:
	//
	// 0
	V6Count *int64 `json:"V6Count,omitempty" xml:"V6Count,omitempty"`
	// The number of IPv6-based HTTP requests.
	//
	// example:
	//
	// 0
	V6HttpCount *int64 `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	// The number of IPv6-based HTTPS requests.
	//
	// example:
	//
	// 0
	V6HttpsCount *int64 `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribePdnsRequestStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetDohTotalCount() *int64 {
	return s.DohTotalCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetHttpCount() *int64 {
	return s.HttpCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetHttpsCount() *int64 {
	return s.HttpsCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetIpCount() *int64 {
	return s.IpCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetMaxThreatLevel() *string {
	return s.MaxThreatLevel
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetThreatCount() *int64 {
	return s.ThreatCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetThreatInfo() []*DescribePdnsRequestStatisticsResponseBodyDataThreatInfo {
	return s.ThreatInfo
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetUdpTotalCount() *int64 {
	return s.UdpTotalCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetV4Count() *int64 {
	return s.V4Count
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetV4HttpCount() *int64 {
	return s.V4HttpCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetV4HttpsCount() *int64 {
	return s.V4HttpsCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetV6Count() *int64 {
	return s.V6Count
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetV6HttpCount() *int64 {
	return s.V6HttpCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) GetV6HttpsCount() *int64 {
	return s.V6HttpsCount
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetDohTotalCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.DohTotalCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetDomainName(v string) *DescribePdnsRequestStatisticsResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetHttpCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.HttpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetHttpsCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.HttpsCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetIpCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.IpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetMaxThreatLevel(v string) *DescribePdnsRequestStatisticsResponseBodyData {
	s.MaxThreatLevel = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetSubDomain(v string) *DescribePdnsRequestStatisticsResponseBodyData {
	s.SubDomain = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetThreatCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.ThreatCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetThreatInfo(v []*DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) *DescribePdnsRequestStatisticsResponseBodyData {
	s.ThreatInfo = v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetTotalCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetUdpTotalCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.UdpTotalCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetV4Count(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.V4Count = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetV4HttpCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.V4HttpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetV4HttpsCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetV6Count(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.V6Count = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetV6HttpCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.V6HttpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) SetV6HttpsCount(v int64) *DescribePdnsRequestStatisticsResponseBodyData {
	s.V6HttpsCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyData) Validate() error {
	if s.ThreatInfo != nil {
		for _, item := range s.ThreatInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePdnsRequestStatisticsResponseBodyDataThreatInfo struct {
	// The current version does not support this parameter.
	//
	// example:
	//
	// -
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	// The current version does not support this parameter.
	//
	// example:
	//
	// -
	ThreatType *string `json:"ThreatType,omitempty" xml:"ThreatType,omitempty"`
}

func (s DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) GetThreatType() *string {
	return s.ThreatType
}

func (s *DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) SetThreatLevel(v string) *DescribePdnsRequestStatisticsResponseBodyDataThreatInfo {
	s.ThreatLevel = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) SetThreatType(v string) *DescribePdnsRequestStatisticsResponseBodyDataThreatInfo {
	s.ThreatType = &v
	return s
}

func (s *DescribePdnsRequestStatisticsResponseBodyDataThreatInfo) Validate() error {
	return dara.Validate(s)
}
