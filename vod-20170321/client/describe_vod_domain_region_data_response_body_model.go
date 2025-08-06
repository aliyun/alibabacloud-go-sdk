// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRegionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainRegionDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainRegionDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRegionDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainRegionDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainRegionDataResponseBody
	GetStartTime() *string
	SetValue(v *DescribeVodDomainRegionDataResponseBodyValue) *DescribeVodDomainRegionDataResponseBody
	GetValue() *DescribeVodDomainRegionDataResponseBodyValue
}

type DescribeVodDomainRegionDataResponseBody struct {
	DataInterval *string                                       `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                       `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeVodDomainRegionDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeVodDomainRegionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRegionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRegionDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainRegionDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRegionDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRegionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRegionDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRegionDataResponseBody) GetValue() *DescribeVodDomainRegionDataResponseBodyValue {
	return s.Value
}

func (s *DescribeVodDomainRegionDataResponseBody) SetDataInterval(v string) *DescribeVodDomainRegionDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBody) SetDomainName(v string) *DescribeVodDomainRegionDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBody) SetEndTime(v string) *DescribeVodDomainRegionDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBody) SetRequestId(v string) *DescribeVodDomainRegionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBody) SetStartTime(v string) *DescribeVodDomainRegionDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBody) SetValue(v *DescribeVodDomainRegionDataResponseBodyValue) *DescribeVodDomainRegionDataResponseBody {
	s.Value = v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRegionDataResponseBodyValue struct {
	RegionProportionData []*DescribeVodDomainRegionDataResponseBodyValueRegionProportionData `json:"RegionProportionData,omitempty" xml:"RegionProportionData,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRegionDataResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRegionDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRegionDataResponseBodyValue) GetRegionProportionData() []*DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	return s.RegionProportionData
}

func (s *DescribeVodDomainRegionDataResponseBodyValue) SetRegionProportionData(v []*DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) *DescribeVodDomainRegionDataResponseBodyValue {
	s.RegionProportionData = v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainRegionDataResponseBodyValueRegionProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionEname     *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetAvgObjectSize() *string {
	return s.AvgObjectSize
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseRate() *string {
	return s.AvgResponseRate
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetAvgResponseTime() *string {
	return s.AvgResponseTime
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetBps() *string {
	return s.Bps
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetBytesProportion() *string {
	return s.BytesProportion
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetQps() *string {
	return s.Qps
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetRegion() *string {
	return s.Region
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetRegionEname() *string {
	return s.RegionEname
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetReqErrRate() *string {
	return s.ReqErrRate
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) GetTotalQuery() *string {
	return s.TotalQuery
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetAvgObjectSize(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseRate(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetAvgResponseTime(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetBps(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetBytesProportion(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetProportion(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetQps(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetRegion(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.Region = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetRegionEname(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.RegionEname = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetReqErrRate(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetTotalBytes(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) SetTotalQuery(v string) *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeVodDomainRegionDataResponseBodyValueRegionProportionData) Validate() error {
	return dara.Validate(s)
}
