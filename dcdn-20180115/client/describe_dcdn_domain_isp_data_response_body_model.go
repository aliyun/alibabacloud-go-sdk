// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIspDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainIspDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainIspDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainIspDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainIspDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainIspDataResponseBody
	GetStartTime() *string
	SetValue(v *DescribeDcdnDomainIspDataResponseBodyValue) *DescribeDcdnDomainIspDataResponseBody
	GetValue() *DescribeDcdnDomainIspDataResponseBodyValue
}

type DescribeDcdnDomainIspDataResponseBody struct {
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
	// 2019-12-06T12:00:00Z
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
	// 2019-12-05T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The access statistics by ISP.
	Value *DescribeDcdnDomainIspDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainIspDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainIspDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainIspDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIspDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainIspDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIspDataResponseBody) GetValue() *DescribeDcdnDomainIspDataResponseBodyValue {
	return s.Value
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainIspDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) SetValue(v *DescribeDcdnDomainIspDataResponseBodyValue) *DescribeDcdnDomainIspDataResponseBody {
	s.Value = v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIspDataResponseBodyValue struct {
	IspProportionData []*DescribeDcdnDomainIspDataResponseBodyValueIspProportionData `json:"IspProportionData,omitempty" xml:"IspProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainIspDataResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponseBodyValue) GetIspProportionData() []*DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	return s.IspProportionData
}

func (s *DescribeDcdnDomainIspDataResponseBodyValue) SetIspProportionData(v []*DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) *DescribeDcdnDomainIspDataResponseBodyValue {
	s.IspProportionData = v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIspDataResponseBodyValueIspProportionData struct {
	// The average response size. Unit: bytes.
	//
	// example:
	//
	// 800019.0
	AvgObjectSize *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	// The average response speed. Unit: byte/ms.
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
	// The proportion of network traffic. For example, a value of 90 indicates that 90% of network traffic was coming from the specified ISP.
	//
	// example:
	//
	// 0.003544181046236794
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	// The information about the ISP.
	//
	// example:
	//
	// China Unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the ISP.
	//
	// example:
	//
	// unicom
	IspEname *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 0.01155980271270037
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	// The number of queries per second (QPS).
	//
	// example:
	//
	// 5.9523809523809524E-5
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The total volume of traffic.
	//
	// example:
	//
	// 2400057
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The total number of requests that are destined for your website.
	//
	// example:
	//
	// 3
	TotalQuery *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetAvgObjectSize() *string {
	return s.AvgObjectSize
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetAvgResponseRate() *string {
	return s.AvgResponseRate
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetAvgResponseTime() *string {
	return s.AvgResponseTime
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetBps() *string {
	return s.Bps
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetBytesProportion() *string {
	return s.BytesProportion
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetIsp() *string {
	return s.Isp
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetIspEname() *string {
	return s.IspEname
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetQps() *string {
	return s.Qps
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) GetTotalQuery() *string {
	return s.TotalQuery
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetAvgObjectSize(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetAvgResponseRate(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetAvgResponseTime(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetBps(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetBytesProportion(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetIsp(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Isp = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetIspEname(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.IspEname = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetProportion(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetQps(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetTotalBytes(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) SetTotalQuery(v string) *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponseBodyValueIspProportionData) Validate() error {
	return dara.Validate(s)
}
