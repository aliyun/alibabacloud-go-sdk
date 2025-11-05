// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeCdnDomainLogsResponseBodyDomainLogDetails) *DescribeCdnDomainLogsResponseBody
	GetDomainLogDetails() *DescribeCdnDomainLogsResponseBodyDomainLogDetails
	SetRequestId(v string) *DescribeCdnDomainLogsResponseBody
	GetRequestId() *string
}

type DescribeCdnDomainLogsResponseBody struct {
	// A set of DomainLogDetail data.
	DomainLogDetails *DescribeCdnDomainLogsResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBody) GetDomainLogDetails() *DescribeCdnDomainLogsResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeCdnDomainLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainLogsResponseBody) SetDomainLogDetails(v *DescribeCdnDomainLogsResponseBodyDomainLogDetails) *DescribeCdnDomainLogsResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetRequestId(v string) *DescribeCdnDomainLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) Validate() error {
	if s.DomainLogDetails != nil {
		if err := s.DomainLogDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) *DescribeCdnDomainLogsResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetails) Validate() error {
	if s.DomainLogDetail != nil {
		for _, item := range s.DomainLogDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail struct {
	// The accelerated domain name.
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
	// A set of LogInfoDetail data.
	LogInfos *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	// A set of PageInfoDetail data.
	PageInfos *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetDomainName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
	if s.LogInfos != nil {
		if err := s.LogInfos.Validate(); err != nil {
			return err
		}
	}
	if s.PageInfos != nil {
		if err := s.PageInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
	if s.LogInfoDetail != nil {
		for _, item := range s.LogInfoDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
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
	// demo.aliyundoc.com_2015_05_23_2100_2200.gz
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// The path of the log file.
	//
	// example:
	//
	// guide.aliyundoc.com-hangzhou.xxx
	LogPath *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	// The size of the log file.
	//
	// example:
	//
	// 258
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2015-05-23T13:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
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

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
