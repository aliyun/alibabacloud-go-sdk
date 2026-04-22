// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePvtzStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurPage(v int32) *DescribePvtzStatisticsSummaryResponseBody
	GetCurPage() *int32
	SetPageData(v []*DescribePvtzStatisticsSummaryResponseBodyPageData) *DescribePvtzStatisticsSummaryResponseBody
	GetPageData() []*DescribePvtzStatisticsSummaryResponseBodyPageData
	SetPageSize(v int32) *DescribePvtzStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePvtzStatisticsSummaryResponseBody
	GetRequestId() *string
	SetTotalPage(v int32) *DescribePvtzStatisticsSummaryResponseBody
	GetTotalPage() *int32
	SetTotalSize(v int32) *DescribePvtzStatisticsSummaryResponseBody
	GetTotalSize() *int32
}

type DescribePvtzStatisticsSummaryResponseBody struct {
	// example:
	//
	// 1
	CurPage  *int32                                               `json:"CurPage,omitempty" xml:"CurPage,omitempty"`
	PageData []*DescribePvtzStatisticsSummaryResponseBodyPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F0FCB52A-D512-41A0-8595-40234EDCFD8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// example:
	//
	// 11
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribePvtzStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsSummaryResponseBody) GetCurPage() *int32 {
	return s.CurPage
}

func (s *DescribePvtzStatisticsSummaryResponseBody) GetPageData() []*DescribePvtzStatisticsSummaryResponseBodyPageData {
	return s.PageData
}

func (s *DescribePvtzStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePvtzStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePvtzStatisticsSummaryResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribePvtzStatisticsSummaryResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribePvtzStatisticsSummaryResponseBody) SetCurPage(v int32) *DescribePvtzStatisticsSummaryResponseBody {
	s.CurPage = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBody) SetPageData(v []*DescribePvtzStatisticsSummaryResponseBodyPageData) *DescribePvtzStatisticsSummaryResponseBody {
	s.PageData = v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribePvtzStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBody) SetRequestId(v string) *DescribePvtzStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBody) SetTotalPage(v int32) *DescribePvtzStatisticsSummaryResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBody) SetTotalSize(v int32) *DescribePvtzStatisticsSummaryResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBody) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePvtzStatisticsSummaryResponseBodyPageData struct {
	// example:
	//
	// 1776774900000
	AggrTimestamp *int64 `json:"AggrTimestamp,omitempty" xml:"AggrTimestamp,omitempty"`
	// example:
	//
	// 4
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// lb-ni1iadds-2c8uyzvgrm5ftsnq.clb.gz-tencentclb.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 19
	FluctuationValue *int32 `json:"FluctuationValue,omitempty" xml:"FluctuationValue,omitempty"`
	// example:
	//
	// 30
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// AUTHORITY
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// 2227
	PreviousCount *int64 `json:"PreviousCount,omitempty" xml:"PreviousCount,omitempty"`
	// example:
	//
	// UDP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// A
	Qtype *string `json:"Qtype,omitempty" xml:"Qtype,omitempty"`
	// example:
	//
	// 30
	Ratio *int64 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// example:
	//
	// 0
	Rcode *string `json:"Rcode,omitempty" xml:"Rcode,omitempty"`
	// example:
	//
	// 110.19.60.72
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// telecom
	SourceIsp *string `json:"SourceIsp,omitempty" xml:"SourceIsp,omitempty"`
	// example:
	//
	// cn-beijing
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-bp1uvv79h1t8unnzdh3nq
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 479226c2db1f9bed449b0502c13bcd9d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// example:
	//
	// bwcj.biz
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribePvtzStatisticsSummaryResponseBodyPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribePvtzStatisticsSummaryResponseBodyPageData) GoString() string {
	return s.String()
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetAggrTimestamp() *int64 {
	return s.AggrTimestamp
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetCount() *int64 {
	return s.Count
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetFluctuationValue() *int32 {
	return s.FluctuationValue
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetLatency() *int64 {
	return s.Latency
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetLine() *string {
	return s.Line
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetModule() *string {
	return s.Module
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetPreviousCount() *int64 {
	return s.PreviousCount
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetQtype() *string {
	return s.Qtype
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetRatio() *int64 {
	return s.Ratio
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetRcode() *string {
	return s.Rcode
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetSourceIsp() *string {
	return s.SourceIsp
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetAggrTimestamp(v int64) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.AggrTimestamp = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetCount(v int64) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Count = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetDomainName(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.DomainName = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetFluctuationValue(v int32) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.FluctuationValue = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetLatency(v int64) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Latency = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetLine(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Line = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetModule(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Module = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetPreviousCount(v int64) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.PreviousCount = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetProtocol(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Protocol = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetQtype(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Qtype = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetRatio(v int64) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Ratio = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetRcode(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.Rcode = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetSourceIp(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.SourceIp = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetSourceIsp(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.SourceIsp = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetSourceRegion(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.SourceRegion = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetVpcId(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.VpcId = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetZoneId(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.ZoneId = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) SetZoneName(v string) *DescribePvtzStatisticsSummaryResponseBodyPageData {
	s.ZoneName = &v
	return s
}

func (s *DescribePvtzStatisticsSummaryResponseBodyPageData) Validate() error {
	return dara.Validate(s)
}
