// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSiteLogsResponseBody
	GetRequestId() *string
	SetSiteLogDetails(v []*DescribeSiteLogsResponseBodySiteLogDetails) *DescribeSiteLogsResponseBody
	GetSiteLogDetails() []*DescribeSiteLogsResponseBodySiteLogDetails
}

type DescribeSiteLogsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteLogDetails []*DescribeSiteLogsResponseBodySiteLogDetails `json:"SiteLogDetails,omitempty" xml:"SiteLogDetails,omitempty" type:"Repeated"`
}

func (s DescribeSiteLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteLogsResponseBody) GetSiteLogDetails() []*DescribeSiteLogsResponseBodySiteLogDetails {
	return s.SiteLogDetails
}

func (s *DescribeSiteLogsResponseBody) SetRequestId(v string) *DescribeSiteLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteLogsResponseBody) SetSiteLogDetails(v []*DescribeSiteLogsResponseBodySiteLogDetails) *DescribeSiteLogsResponseBody {
	s.SiteLogDetails = v
	return s
}

func (s *DescribeSiteLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteLogsResponseBodySiteLogDetails struct {
	// example:
	//
	// 300
	LogCount  *int32                                                `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  []*DescribeSiteLogsResponseBodySiteLogDetailsLogInfos `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Repeated"`
	PageInfos *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos  `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
	// example:
	//
	// 123456***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DescribeSiteLogsResponseBodySiteLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteLogsResponseBodySiteLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) GetLogCount() *int32 {
	return s.LogCount
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) GetLogInfos() []*DescribeSiteLogsResponseBodySiteLogDetailsLogInfos {
	return s.LogInfos
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) GetPageInfos() *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos {
	return s.PageInfos
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) GetSiteName() *string {
	return s.SiteName
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) SetLogCount(v int32) *DescribeSiteLogsResponseBodySiteLogDetails {
	s.LogCount = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) SetLogInfos(v []*DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) *DescribeSiteLogsResponseBodySiteLogDetails {
	s.LogInfos = v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) SetPageInfos(v *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) *DescribeSiteLogsResponseBodySiteLogDetails {
	s.PageInfos = v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) SetSiteId(v int64) *DescribeSiteLogsResponseBodySiteLogDetails {
	s.SiteId = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) SetSiteName(v string) *DescribeSiteLogsResponseBodySiteLogDetails {
	s.SiteName = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteLogsResponseBodySiteLogDetailsLogInfos struct {
	// example:
	//
	// 2022-11-06T17:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// example.com_2022_11_07_000000_020000.gz.xxxxxx
	LogName *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	// example:
	//
	// example.aliyundoc.com /v1.l1cache/105252530/example.com/2022_11_07/example.com_2022_11_07_000000_020000.gz.xxxxxx?Expires=1636963354&OSSAccessKeyId=LTAIviCc6zy8****&Signature=u0V6foRfZniHE8i%2BHUdxGOhZsK****
	LogPath *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	// example:
	//
	// 438304768
	LogSize *int32 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// example:
	//
	// 2022-11-06T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) GetLogName() *string {
	return s.LogName
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) GetLogSize() *int32 {
	return s.LogSize
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) SetEndTime(v string) *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) SetLogName(v string) *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos {
	s.LogName = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) SetLogPath(v string) *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos {
	s.LogPath = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) SetLogSize(v int32) *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos {
	s.LogSize = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) SetStartTime(v string) *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsLogInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteLogsResponseBodySiteLogDetailsPageInfos struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 47
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) SetPageIndex(v int32) *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) SetPageSize(v int32) *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) SetTotalCount(v int32) *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos {
	s.TotalCount = &v
	return s
}

func (s *DescribeSiteLogsResponseBodySiteLogDetailsPageInfos) Validate() error {
	return dara.Validate(s)
}
