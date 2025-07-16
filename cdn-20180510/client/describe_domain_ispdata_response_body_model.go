// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainISPDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainISPDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainISPDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainISPDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainISPDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainISPDataResponseBody
	GetStartTime() *string
	SetValue(v *DescribeDomainISPDataResponseBodyValue) *DescribeDomainISPDataResponseBody
	GetValue() *DescribeDomainISPDataResponseBodyValue
}

type DescribeDomainISPDataResponseBody struct {
	// The time interval. Unit: seconds.
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
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DE81639B-DAC1-4C76-AB72-F34B836837D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2019-11-29T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The access statistics by ISP.
	Value *DescribeDomainISPDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainISPDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainISPDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainISPDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainISPDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainISPDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainISPDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainISPDataResponseBody) GetValue() *DescribeDomainISPDataResponseBodyValue {
	return s.Value
}

func (s *DescribeDomainISPDataResponseBody) SetDataInterval(v string) *DescribeDomainISPDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetDomainName(v string) *DescribeDomainISPDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetEndTime(v string) *DescribeDomainISPDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetRequestId(v string) *DescribeDomainISPDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetStartTime(v string) *DescribeDomainISPDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBody) SetValue(v *DescribeDomainISPDataResponseBodyValue) *DescribeDomainISPDataResponseBody {
	s.Value = v
	return s
}

func (s *DescribeDomainISPDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainISPDataResponseBodyValue struct {
	ISPProportionData []*DescribeDomainISPDataResponseBodyValueISPProportionData `json:"ISPProportionData,omitempty" xml:"ISPProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainISPDataResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainISPDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBodyValue) GetISPProportionData() []*DescribeDomainISPDataResponseBodyValueISPProportionData {
	return s.ISPProportionData
}

func (s *DescribeDomainISPDataResponseBodyValue) SetISPProportionData(v []*DescribeDomainISPDataResponseBodyValueISPProportionData) *DescribeDomainISPDataResponseBodyValue {
	s.ISPProportionData = v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainISPDataResponseBodyValueISPProportionData struct {
	// The average response size. Unit: bytes.
	//
	// example:
	//
	// 7081884.7
	AvgObjectSize *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	// The average response speed. Unit: byte/ms.
	//
	// example:
	//
	// 88.92594866772144
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	// The average response time. Unit: milliseconds.
	//
	// example:
	//
	// 79638.0
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	// The bandwidth.
	//
	// example:
	//
	// 1311.4601296296296
	Bps *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The proportion of network traffic.
	//
	// example:
	//
	// 0.012220518530445479
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	// The information about the ISP.
	//
	// example:
	//
	// Alibaba
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The name of the ISP.
	//
	// example:
	//
	// alibaba
	IspEname *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 0.004509176173513099
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	// The QPS.
	//
	// example:
	//
	// 2.3148148148148147E-5
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The request error rate.
	//
	// example:
	//
	// 0.0
	ReqErrRate *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	// The total volume of traffic.
	//
	// example:
	//
	// 7081884
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 1
	TotalQuery *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeDomainISPDataResponseBodyValueISPProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainISPDataResponseBodyValueISPProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetAvgObjectSize() *string {
	return s.AvgObjectSize
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetAvgResponseRate() *string {
	return s.AvgResponseRate
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetAvgResponseTime() *string {
	return s.AvgResponseTime
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetBps() *string {
	return s.Bps
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetBytesProportion() *string {
	return s.BytesProportion
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetISP() *string {
	return s.ISP
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetIspEname() *string {
	return s.IspEname
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetQps() *string {
	return s.Qps
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetReqErrRate() *string {
	return s.ReqErrRate
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) GetTotalQuery() *string {
	return s.TotalQuery
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgObjectSize(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseTime(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetBps(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetBytesProportion(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetISP(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ISP = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetIspEname(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.IspEname = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetProportion(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetQps(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetReqErrRate(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetTotalBytes(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) SetTotalQuery(v string) *DescribeDomainISPDataResponseBodyValueISPProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeDomainISPDataResponseBodyValueISPProportionData) Validate() error {
	return dara.Validate(s)
}
