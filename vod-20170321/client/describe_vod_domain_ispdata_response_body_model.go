// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainISPDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainISPDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainISPDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainISPDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainISPDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainISPDataResponseBody
	GetStartTime() *string
	SetValue(v *DescribeVodDomainISPDataResponseBodyValue) *DescribeVodDomainISPDataResponseBody
	GetValue() *DescribeVodDomainISPDataResponseBodyValue
}

type DescribeVodDomainISPDataResponseBody struct {
	DataInterval *string                                    `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Value        *DescribeVodDomainISPDataResponseBodyValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeVodDomainISPDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainISPDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainISPDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainISPDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainISPDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainISPDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainISPDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainISPDataResponseBody) GetValue() *DescribeVodDomainISPDataResponseBodyValue {
	return s.Value
}

func (s *DescribeVodDomainISPDataResponseBody) SetDataInterval(v string) *DescribeVodDomainISPDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBody) SetDomainName(v string) *DescribeVodDomainISPDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBody) SetEndTime(v string) *DescribeVodDomainISPDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBody) SetRequestId(v string) *DescribeVodDomainISPDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBody) SetStartTime(v string) *DescribeVodDomainISPDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBody) SetValue(v *DescribeVodDomainISPDataResponseBodyValue) *DescribeVodDomainISPDataResponseBody {
	s.Value = v
	return s
}

func (s *DescribeVodDomainISPDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainISPDataResponseBodyValue struct {
	ISPProportionData []*DescribeVodDomainISPDataResponseBodyValueISPProportionData `json:"ISPProportionData,omitempty" xml:"ISPProportionData,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainISPDataResponseBodyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainISPDataResponseBodyValue) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainISPDataResponseBodyValue) GetISPProportionData() []*DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	return s.ISPProportionData
}

func (s *DescribeVodDomainISPDataResponseBodyValue) SetISPProportionData(v []*DescribeVodDomainISPDataResponseBodyValueISPProportionData) *DescribeVodDomainISPDataResponseBodyValue {
	s.ISPProportionData = v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainISPDataResponseBodyValueISPProportionData struct {
	AvgObjectSize   *string `json:"AvgObjectSize,omitempty" xml:"AvgObjectSize,omitempty"`
	AvgResponseRate *string `json:"AvgResponseRate,omitempty" xml:"AvgResponseRate,omitempty"`
	AvgResponseTime *string `json:"AvgResponseTime,omitempty" xml:"AvgResponseTime,omitempty"`
	Bps             *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BytesProportion *string `json:"BytesProportion,omitempty" xml:"BytesProportion,omitempty"`
	ISP             *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	IspEname        *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	Proportion      *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
	Qps             *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	ReqErrRate      *string `json:"ReqErrRate,omitempty" xml:"ReqErrRate,omitempty"`
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalQuery      *string `json:"TotalQuery,omitempty" xml:"TotalQuery,omitempty"`
}

func (s DescribeVodDomainISPDataResponseBodyValueISPProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainISPDataResponseBodyValueISPProportionData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetAvgObjectSize() *string {
	return s.AvgObjectSize
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetAvgResponseRate() *string {
	return s.AvgResponseRate
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetAvgResponseTime() *string {
	return s.AvgResponseTime
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetBps() *string {
	return s.Bps
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetBytesProportion() *string {
	return s.BytesProportion
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetISP() *string {
	return s.ISP
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetIspEname() *string {
	return s.IspEname
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetQps() *string {
	return s.Qps
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetReqErrRate() *string {
	return s.ReqErrRate
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) GetTotalQuery() *string {
	return s.TotalQuery
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetAvgObjectSize(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.AvgObjectSize = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseRate(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseRate = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetAvgResponseTime(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.AvgResponseTime = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetBps(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.Bps = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetBytesProportion(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.BytesProportion = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetISP(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.ISP = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetIspEname(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.IspEname = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetProportion(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetQps(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.Qps = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetReqErrRate(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.ReqErrRate = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetTotalBytes(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) SetTotalQuery(v string) *DescribeVodDomainISPDataResponseBodyValueISPProportionData {
	s.TotalQuery = &v
	return s
}

func (s *DescribeVodDomainISPDataResponseBodyValueISPProportionData) Validate() error {
	return dara.Validate(s)
}
