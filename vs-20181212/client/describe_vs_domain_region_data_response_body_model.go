// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainRegionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVsDomainRegionDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainRegionDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainRegionDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVsDomainRegionDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainRegionDataResponseBody
	GetStartTime() *string
	SetValue(v *DescribeVsDomainRegionDataResponseBodyValue) *DescribeVsDomainRegionDataResponseBody
	GetValue() *DescribeVsDomainRegionDataResponseBodyValue
}

type DescribeVsDomainRegionDataResponseBody struct {
	// example:
	//
	// 3600
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-10-31T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-10-30T16:00:00Z
	StartTime *string                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value     *DescribeVsDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeVsDomainRegionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainRegionDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainRegionDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainRegionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainRegionDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainRegionDataResponseBody) GetValue() *DescribeVsDomainRegionDataResponseBodyValue {
	return s.Value
}

func (s *DescribeVsDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeVsDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetDomainName(v string) *DescribeVsDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetEndTime(v string) *DescribeVsDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetRequestId(v string) *DescribeVsDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetStartTime(v string) *DescribeVsDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) SetValue(v *DescribeVsDomainRegionDataResponseBodyValue) *DescribeVsDomainRegionDataResponseBody {
	s.Value = v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBody) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeVsDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainRegionDataResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponseBodyValue) GetRegionProportionData() []*DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	return s.RegionProportionData
}

func (s *DescribeVsDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) *DescribeVsDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValue) Validate() error {
	if s.RegionProportionData != nil {
		for _, item := range s.RegionProportionData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVsDomainRegionDataResponseBodyValueRegionProportionData struct {
	// example:
	//
	// 2888253.7875
	AvgObjectSize *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	// example:
	//
	// 154.3345765545624
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	// example:
	//
	// 5183.666666666667
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	// example:
	//
	// 380.9614285714286
	Bps *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// example:
	//
	// 0.003544181046236794
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	// example:
	//
	// 0.01155980271270037
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	// example:
	//
	// 0.001746031746031746
	Qps    *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// chongqing
	RegionEname *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	// example:
	//
	// 0
	ReqErrRate *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	// example:
	//
	// 2400057
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// example:
	//
	// 3
	TotalQuery *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetAvgObjectSize() *string {
	return s.AvgObjectSize
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseRate() *string {
	return s.AvgResponseRate
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseTime() *string {
	return s.AvgResponseTime
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetBps() *string {
	return s.Bps
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetBytesProportion() *string {
	return s.BytesProportion
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetQps() *string {
	return s.Qps
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetRegion() *string {
	return s.Region
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetRegionEname() *string {
	return s.RegionEname
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetReqErrRate() *string {
	return s.ReqErrRate
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) GetTotalQuery() *string {
	return s.TotalQuery
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetReqErrRate(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeVsDomainRegionDataResponseBodyValueRegionProportionData) Validate() error {
	return dara.Validate(s)
}
