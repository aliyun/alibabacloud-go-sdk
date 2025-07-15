// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserTrafficLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails) *DescribeLiveUserTrafficLogResponseBody
	GetDomainLogDetails() *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails
	SetDomainName(v string) *DescribeLiveUserTrafficLogResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeLiveUserTrafficLogResponseBody
	GetRequestId() *string
}

type DescribeLiveUserTrafficLogResponseBody struct {
	DomainLogDetails *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	DomainName       *string                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveUserTrafficLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogResponseBody) GetDomainLogDetails() *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeLiveUserTrafficLogResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveUserTrafficLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveUserTrafficLogResponseBody) SetDomainLogDetails(v *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails) *DescribeLiveUserTrafficLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBody) SetDomainName(v string) *DescribeLiveUserTrafficLogResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBody) SetRequestId(v string) *DescribeLiveUserTrafficLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserTrafficLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	LogCount  *int64                                                                          `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeLiveUserTrafficLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
