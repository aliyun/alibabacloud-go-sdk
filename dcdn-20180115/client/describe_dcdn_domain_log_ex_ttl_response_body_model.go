// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainLogExTtlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails) *DescribeDcdnDomainLogExTtlResponseBody
	GetDomainLogDetails() *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails
	SetRequestId(v string) *DescribeDcdnDomainLogExTtlResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainLogExTtlResponseBody struct {
	DomainLogDetails *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainLogExTtlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlResponseBody) GetDomainLogDetails() *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeDcdnDomainLogExTtlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainLogExTtlResponseBody) SetDomainLogDetails(v *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails) *DescribeDcdnDomainLogExTtlResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBody) SetRequestId(v string) *DescribeDcdnDomainLogExTtlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail struct {
	DomainName *string                                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LogCount   *int64                                                                          `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos   *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos  *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetDomainName(v string) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
