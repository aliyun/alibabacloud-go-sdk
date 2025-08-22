// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRegionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainRegionDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainRegionDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRegionDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainRegionDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainRegionDataResponseBody
	GetStartTime() *string
	SetValue(v *DescribeDcdnDomainRegionDataResponseBodyValue) *DescribeDcdnDomainRegionDataResponseBody
	GetValue() *DescribeDcdnDomainRegionDataResponseBodyValue
}

type DescribeDcdnDomainRegionDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 86400
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-07T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2E5AD83F-BD7B-462E-8319-2E30E305519A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-05T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The proportions of requests that were initiated from each region.
	Value *DescribeDcdnDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainRegionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainRegionDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRegionDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRegionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRegionDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRegionDataResponseBody) GetValue() *DescribeDcdnDomainRegionDataResponseBodyValue {
	return s.Value
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) SetValue(v *DescribeDcdnDomainRegionDataResponseBodyValue) *DescribeDcdnDomainRegionDataResponseBody {
	s.Value = v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRegionDataResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValue) GetRegionProportionData() []*DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	return s.RegionProportionData
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) *DescribeDcdnDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData struct {
	// The average response size. Unit: bytes.
	//
	// example:
	//
	// 0
	AvgObjectSize *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	// The average response speed. Unit: byte/s.
	//
	// example:
	//
	// 0
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	// The average response time. Unit: milliseconds.
	//
	// example:
	//
	// 0
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	// The bandwidth.
	//
	// example:
	//
	// 0
	Bps *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The proportion of network traffic. For example, a value of 90 indicates that 90% of network traffic was coming from the specified ISP.
	//
	// example:
	//
	// 0.003544181046236794
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	// The proportion of requests from the specified region based on the total number of requests in percentile. For example, a value of 90 indicates that 90% of the requests were coming from the specified region.
	//
	// example:
	//
	// 0
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	// The number of queries per second (QPS).
	//
	// example:
	//
	// 0
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The information of the regions.
	//
	// example:
	//
	// Chongqing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// chongqing
	RegionEname *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	// The total amount of network traffic.
	//
	// example:
	//
	// 0
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The total number of requests that are destined for your website.
	//
	// example:
	//
	// 0
	TotalQuery *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetAvgObjectSize() *string {
	return s.AvgObjectSize
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseRate() *string {
	return s.AvgResponseRate
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseTime() *string {
	return s.AvgResponseTime
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetBps() *string {
	return s.Bps
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetBytesProportion() *string {
	return s.BytesProportion
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetQps() *string {
	return s.Qps
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetRegion() *string {
	return s.Region
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetRegionEname() *string {
	return s.RegionEname
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) GetTotalQuery() *string {
	return s.TotalQuery
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeDcdnDomainRegionDataResponseBodyValueRegionProportionData) Validate() error {
	return dara.Validate(s)
}
