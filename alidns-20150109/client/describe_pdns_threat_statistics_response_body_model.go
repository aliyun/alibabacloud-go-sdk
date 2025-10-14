// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribePdnsThreatStatisticsResponseBodyData) *DescribePdnsThreatStatisticsResponseBody
	GetData() []*DescribePdnsThreatStatisticsResponseBodyData
	SetPageNumber(v int64) *DescribePdnsThreatStatisticsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsThreatStatisticsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePdnsThreatStatisticsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePdnsThreatStatisticsResponseBody
	GetTotalCount() *int64
}

type DescribePdnsThreatStatisticsResponseBody struct {
	Data       []*DescribePdnsThreatStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int64                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePdnsThreatStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticsResponseBody) GetData() []*DescribePdnsThreatStatisticsResponseBodyData {
	return s.Data
}

func (s *DescribePdnsThreatStatisticsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsThreatStatisticsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsThreatStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsThreatStatisticsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsThreatStatisticsResponseBody) SetData(v []*DescribePdnsThreatStatisticsResponseBodyData) *DescribePdnsThreatStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBody) SetPageNumber(v int64) *DescribePdnsThreatStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBody) SetPageSize(v int64) *DescribePdnsThreatStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBody) SetRequestId(v string) *DescribePdnsThreatStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBody) SetTotalCount(v int64) *DescribePdnsThreatStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBody) Validate() error {
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

type DescribePdnsThreatStatisticsResponseBodyData struct {
	DohTotalCount    *int64  `json:"DohTotalCount,omitempty" xml:"DohTotalCount,omitempty"`
	DomainCount      *int64  `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LatestThreatTime *int64  `json:"LatestThreatTime,omitempty" xml:"LatestThreatTime,omitempty"`
	MaxThreatLevel   *string `json:"MaxThreatLevel,omitempty" xml:"MaxThreatLevel,omitempty"`
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SubDomain        *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	ThreatLevel      *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatType       *string `json:"ThreatType,omitempty" xml:"ThreatType,omitempty"`
	TotalCount       *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UdpTotalCount    *int64  `json:"UdpTotalCount,omitempty" xml:"UdpTotalCount,omitempty"`
}

func (s DescribePdnsThreatStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetDohTotalCount() *int64 {
	return s.DohTotalCount
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetLatestThreatTime() *int64 {
	return s.LatestThreatTime
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetMaxThreatLevel() *string {
	return s.MaxThreatLevel
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetThreatType() *string {
	return s.ThreatType
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) GetUdpTotalCount() *int64 {
	return s.UdpTotalCount
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetDohTotalCount(v int64) *DescribePdnsThreatStatisticsResponseBodyData {
	s.DohTotalCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetDomainCount(v int64) *DescribePdnsThreatStatisticsResponseBodyData {
	s.DomainCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetDomainName(v string) *DescribePdnsThreatStatisticsResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetLatestThreatTime(v int64) *DescribePdnsThreatStatisticsResponseBodyData {
	s.LatestThreatTime = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetMaxThreatLevel(v string) *DescribePdnsThreatStatisticsResponseBodyData {
	s.MaxThreatLevel = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetSourceIp(v string) *DescribePdnsThreatStatisticsResponseBodyData {
	s.SourceIp = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetSubDomain(v string) *DescribePdnsThreatStatisticsResponseBodyData {
	s.SubDomain = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetThreatLevel(v string) *DescribePdnsThreatStatisticsResponseBodyData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetThreatType(v string) *DescribePdnsThreatStatisticsResponseBodyData {
	s.ThreatType = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetTotalCount(v int64) *DescribePdnsThreatStatisticsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) SetUdpTotalCount(v int64) *DescribePdnsThreatStatisticsResponseBodyData {
	s.UdpTotalCount = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
