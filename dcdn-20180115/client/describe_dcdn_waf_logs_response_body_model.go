// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v []*DescribeDcdnWafLogsResponseBodyDomainLogDetails) *DescribeDcdnWafLogsResponseBody
	GetDomainLogDetails() []*DescribeDcdnWafLogsResponseBodyDomainLogDetails
	SetRequestId(v string) *DescribeDcdnWafLogsResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafLogsResponseBody struct {
	// Details about logs returned.
	DomainLogDetails []*DescribeDcdnWafLogsResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0985A362-C81E-5A56-891D-90226BEECA7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafLogsResponseBody) GetDomainLogDetails() []*DescribeDcdnWafLogsResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeDcdnWafLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafLogsResponseBody) SetDomainLogDetails(v []*DescribeDcdnWafLogsResponseBodyDomainLogDetails) *DescribeDcdnWafLogsResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeDcdnWafLogsResponseBody) SetRequestId(v string) *DescribeDcdnWafLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafLogsResponseBodyDomainLogDetails struct {
	// The WAF domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The total number of entries returned on the current page.
	//
	// example:
	//
	// 10
	LogCount *int64 `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	// The log information.
	LogInfos []*DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Repeated"`
	// The page information.
	PageInfos *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeDcdnWafLogsResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafLogsResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) GetLogInfos() []*DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos {
	return s.LogInfos
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) GetPageInfos() *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos {
	return s.PageInfos
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) SetDomainName(v string) *DescribeDcdnWafLogsResponseBodyDomainLogDetails {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) SetLogCount(v int64) *DescribeDcdnWafLogsResponseBodyDomainLogDetails {
	s.LogCount = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) SetLogInfos(v []*DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) *DescribeDcdnWafLogsResponseBodyDomainLogDetails {
	s.LogInfos = v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) SetPageInfos(v *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) *DescribeDcdnWafLogsResponseBodyDomainLogDetails {
	s.PageInfos = v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-05-23T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the log file.
	//
	// example:
	//
	// demo.aliyundoc.com_2015_05_23_2100_2200.xxxxxx.gz
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The path of the log file.
	//
	// example:
	//
	// guide.aliyundoc.com-hangzhou.xxx
	LogPath *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	// The size of the log file. Unit: bytes.
	//
	// example:
	//
	// 258
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2015-05-23T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) GetLogName() *string {
	return s.LogName
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) SetEndTime(v string) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) SetLogName(v string) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos {
	s.LogName = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) SetLogPath(v string) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos {
	s.LogPath = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) SetLogSize(v int64) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos {
	s.LogSize = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) SetStartTime(v string) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos struct {
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
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) SetPageIndex(v int64) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) SetPageSize(v int64) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) SetTotal(v int64) *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeDcdnWafLogsResponseBodyDomainLogDetailsPageInfos) Validate() error {
	return dara.Validate(s)
}
