// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeDcdnDomainLogResponseBodyDomainLogDetails) *DescribeDcdnDomainLogResponseBody
	GetDomainLogDetails() *DescribeDcdnDomainLogResponseBodyDomainLogDetails
	SetDomainName(v string) *DescribeDcdnDomainLogResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDcdnDomainLogResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainLogResponseBody struct {
	// The log information.
	DomainLogDetails *DescribeDcdnDomainLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 95594003-CAC5-5636-AF72-2A094364****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBody) GetDomainLogDetails() *DescribeDcdnDomainLogResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeDcdnDomainLogResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainLogResponseBody) SetDomainLogDetails(v *DescribeDcdnDomainLogResponseBodyDomainLogDetails) *DescribeDcdnDomainLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeDcdnDomainLogResponseBody) SetDomainName(v string) *DescribeDcdnDomainLogResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBody) SetRequestId(v string) *DescribeDcdnDomainLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeDcdnDomainLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	// The total number of entries returned on the current page.
	//
	// example:
	//
	// 4
	LogCount *int64 `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	// Details about the logs.
	LogInfos *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	// The page information.
	PageInfos *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2021-11-07T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the log file.
	//
	// example:
	//
	// example.com_2021_11_08_010000_020000.gz
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The path of the log file.
	//
	// Take note of the Expires field (expiration timestamp) in the response parameter LogPath. If the log download URL expires, you must obtain it again. For more information, see [LogPath field](https://help.aliyun.com/document_detail/31952.html).
	//
	// example:
	//
	// example.aliyundoc.com /v1.l1cache/105252530/example.com/2021_11_08/example.com_2021_11_08_010000_020000.gz?Expires=1636963354&OSSAccessKeyId=yourAccessKeyID&Signature=u0V6foRfZniHE8i%2BHUdxGOhZsK****
	LogPath *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	// The size of the log file. Unit: bytes.
	//
	// example:
	//
	// 192
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2021-11-07T17:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeDcdnDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
