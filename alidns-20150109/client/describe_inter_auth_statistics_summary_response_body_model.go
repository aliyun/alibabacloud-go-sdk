// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurPage(v int32) *DescribeInterAuthStatisticsSummaryResponseBody
	GetCurPage() *int32
	SetPageData(v []*DescribeInterAuthStatisticsSummaryResponseBodyPageData) *DescribeInterAuthStatisticsSummaryResponseBody
	GetPageData() []*DescribeInterAuthStatisticsSummaryResponseBodyPageData
	SetPageSize(v int32) *DescribeInterAuthStatisticsSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInterAuthStatisticsSummaryResponseBody
	GetRequestId() *string
	SetTotalPage(v int32) *DescribeInterAuthStatisticsSummaryResponseBody
	GetTotalPage() *int32
	SetTotalSize(v int32) *DescribeInterAuthStatisticsSummaryResponseBody
	GetTotalSize() *int32
}

type DescribeInterAuthStatisticsSummaryResponseBody struct {
	// example:
	//
	// 1
	CurPage  *int32                                                    `json:"CurPage,omitempty" xml:"CurPage,omitempty"`
	PageData []*DescribeInterAuthStatisticsSummaryResponseBodyPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 853805EA-3D47-47D5-9A1A-A45C24313ABD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// example:
	//
	// 48
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeInterAuthStatisticsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) GetCurPage() *int32 {
	return s.CurPage
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) GetPageData() []*DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	return s.PageData
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) SetCurPage(v int32) *DescribeInterAuthStatisticsSummaryResponseBody {
	s.CurPage = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) SetPageData(v []*DescribeInterAuthStatisticsSummaryResponseBodyPageData) *DescribeInterAuthStatisticsSummaryResponseBody {
	s.PageData = v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) SetPageSize(v int32) *DescribeInterAuthStatisticsSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) SetRequestId(v string) *DescribeInterAuthStatisticsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) SetTotalPage(v int32) *DescribeInterAuthStatisticsSummaryResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) SetTotalSize(v int32) *DescribeInterAuthStatisticsSummaryResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBody) Validate() error {
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

type DescribeInterAuthStatisticsSummaryResponseBodyPageData struct {
	// example:
	//
	// 1776754800000
	AggrTimestamp *int64 `json:"AggrTimestamp,omitempty" xml:"AggrTimestamp,omitempty"`
	// example:
	//
	// 20
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// nervermsf.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 19
	FluctuationValue *int32 `json:"FluctuationValue,omitempty" xml:"FluctuationValue,omitempty"`
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// 2227
	PreviousCount *int64 `json:"PreviousCount,omitempty" xml:"PreviousCount,omitempty"`
	// example:
	//
	// TCP%DF\\"
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// A
	Qtype *string `json:"Qtype,omitempty" xml:"Qtype,omitempty"`
	// example:
	//
	// 20
	Ratio *int64 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// example:
	//
	// 0
	Rcode *string `json:"Rcode,omitempty" xml:"Rcode,omitempty"`
	// example:
	//
	// 112.16.17.203
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// telecom
	SourceIsp *string `json:"SourceIsp,omitempty" xml:"SourceIsp,omitempty"`
	// example:
	//
	// cn-shenzhen
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// example:
	//
	// longzi.xyz
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeInterAuthStatisticsSummaryResponseBodyPageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsSummaryResponseBodyPageData) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetAggrTimestamp() *int64 {
	return s.AggrTimestamp
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetCount() *int64 {
	return s.Count
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetFluctuationValue() *int32 {
	return s.FluctuationValue
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetLine() *string {
	return s.Line
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetPreviousCount() *int64 {
	return s.PreviousCount
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetQtype() *string {
	return s.Qtype
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetRatio() *int64 {
	return s.Ratio
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetRcode() *string {
	return s.Rcode
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetSourceIsp() *string {
	return s.SourceIsp
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetAggrTimestamp(v int64) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.AggrTimestamp = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetCount(v int64) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.Count = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetDomainName(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetFluctuationValue(v int32) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.FluctuationValue = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetLine(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.Line = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetPreviousCount(v int64) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.PreviousCount = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetProtocol(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.Protocol = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetQtype(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.Qtype = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetRatio(v int64) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.Ratio = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetRcode(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.Rcode = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetSourceIp(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.SourceIp = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetSourceIsp(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.SourceIsp = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetSourceRegion(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.SourceRegion = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) SetZoneName(v string) *DescribeInterAuthStatisticsSummaryResponseBodyPageData {
	s.ZoneName = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponseBodyPageData) Validate() error {
	return dara.Validate(s)
}
