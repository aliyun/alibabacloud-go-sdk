// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTrafficDomainLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails) *DescribeLiveTrafficDomainLogResponseBody
	GetDomainLogDetails() *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails
	SetDomainName(v string) *DescribeLiveTrafficDomainLogResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeLiveTrafficDomainLogResponseBody
	GetRequestId() *string
}

type DescribeLiveTrafficDomainLogResponseBody struct {
	DomainLogDetails *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	DomainName       *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId        *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveTrafficDomainLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogResponseBody) GetDomainLogDetails() *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeLiveTrafficDomainLogResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveTrafficDomainLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveTrafficDomainLogResponseBody) SetDomainLogDetails(v *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails) *DescribeLiveTrafficDomainLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBody) SetDomainName(v string) *DescribeLiveTrafficDomainLogResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBody) SetRequestId(v string) *DescribeLiveTrafficDomainLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	LogCount  *int64                                                                            `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeLiveTrafficDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
