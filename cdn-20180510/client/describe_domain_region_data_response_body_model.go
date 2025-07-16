// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRegionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainRegionDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainRegionDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRegionDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainRegionDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainRegionDataResponseBody
	GetStartTime() *string
	SetValue(v *DescribeDomainRegionDataResponseBodyValue) *DescribeDomainRegionDataResponseBody
	GetValue() *DescribeDomainRegionDataResponseBodyValue
}

type DescribeDomainRegionDataResponseBody struct {
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
	// The end of the time range for which the data was queried.
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
	// The beginning of the time range for which the data was queried.
	//
	// example:
	//
	// 2015-12-05T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The proportions of requests initiated from each region.
	Value *DescribeDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainRegionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainRegionDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRegionDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRegionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRegionDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRegionDataResponseBody) GetValue() *DescribeDomainRegionDataResponseBodyValue {
	return s.Value
}

func (s *DescribeDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetDomainName(v string) *DescribeDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetEndTime(v string) *DescribeDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetRequestId(v string) *DescribeDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetStartTime(v string) *DescribeDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) SetValue(v *DescribeDomainRegionDataResponseBodyValue) *DescribeDomainRegionDataResponseBody {
	s.Value = v
	return s
}

func (s *DescribeDomainRegionDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRegionDataResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBodyValue) GetRegionProportionData() []*DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	return s.RegionProportionData
}

func (s *DescribeDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeDomainRegionDataResponseBodyValueRegionProportionData) *DescribeDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainRegionDataResponseBodyValueRegionProportionData struct {
	// The average response size. Unit: bytes.
	//
	// example:
	//
	// 800019.0
	AvgObjectSize *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	// The average response speed. Unit: bit/s.
	//
	// example:
	//
	// 154.3345765545624
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	// The average response time. Unit: milliseconds.
	//
	// example:
	//
	// 5183.666666666667
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	// The bandwidth.
	//
	// example:
	//
	// 380.9614285714286
	Bps *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The proportion of traffic from the region. For example, a value of 90 indicates that 90% of the traffic is from the region.
	//
	// example:
	//
	// 0.003544181046236794
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	// The proportion of visits from the region. For example, a value of 90 indicates that 90% of the visits are from the region.
	//
	// example:
	//
	// 0.01155980271270037
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	// The number of queries per second.
	//
	// example:
	//
	// 5.9523809523809524E-5
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The information about the region.
	//
	// example:
	//
	// Japan
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// japan
	RegionEname *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	// The request error rate. A value of 90 indicates that 90% of the requests encountered errors.
	//
	// example:
	//
	// 0.0
	ReqErrRate *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	// The total traffic. Unit: bytes.
	//
	// example:
	//
	// 2400057
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 3
	TotalQuery *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetAvgObjectSize() *string {
	return s.AvgObjectSize
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseRate() *string {
	return s.AvgResponseRate
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseTime() *string {
	return s.AvgResponseTime
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetBps() *string {
	return s.Bps
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetBytesProportion() *string {
	return s.BytesProportion
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetQps() *string {
	return s.Qps
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetRegion() *string {
	return s.Region
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetRegionEname() *string {
	return s.RegionEname
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetReqErrRate() *string {
	return s.ReqErrRate
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) GetTotalQuery() *string {
	return s.TotalQuery
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetReqErrRate(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeDomainRegionDataResponseBodyValueRegionProportionData) Validate() error {
	return dara.Validate(s)
}
