// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsExTtlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails) *DescribeCdnDomainLogsExTtlResponseBody
	GetDomainLogDetails() *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails
	SetRequestId(v string) *DescribeCdnDomainLogsExTtlResponseBody
	GetRequestId() *string
}

type DescribeCdnDomainLogsExTtlResponseBody struct {
	DomainLogDetails *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainLogsExTtlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlResponseBody) GetDomainLogDetails() *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeCdnDomainLogsExTtlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainLogsExTtlResponseBody) SetDomainLogDetails(v *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails) *DescribeCdnDomainLogsExTtlResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBody) SetRequestId(v string) *DescribeCdnDomainLogsExTtlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 10
	LogCount  *int64                                                                          `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetDomainName(v string) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	// example:
	//
	// 2023-09-23T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// demo.aliyundoc.com_2015_05_23_2100_2200.gz
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// example:
	//
	// guide.aliyundoc.com-hangzhou.xxx
	LogPath *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	// example:
	//
	// 258
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// example:
	//
	// 2023-09-23T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
