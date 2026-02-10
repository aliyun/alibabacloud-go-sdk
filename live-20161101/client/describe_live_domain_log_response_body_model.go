// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeLiveDomainLogResponseBodyDomainLogDetails) *DescribeLiveDomainLogResponseBody
	GetDomainLogDetails() *DescribeLiveDomainLogResponseBodyDomainLogDetails
	SetDomainName(v string) *DescribeLiveDomainLogResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeLiveDomainLogResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainLogResponseBody struct {
	DomainLogDetails *DescribeLiveDomainLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	// The streaming domain or ingest domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogResponseBody) GetDomainLogDetails() *DescribeLiveDomainLogResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeLiveDomainLogResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainLogResponseBody) SetDomainLogDetails(v *DescribeLiveDomainLogResponseBodyDomainLogDetails) *DescribeLiveDomainLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeLiveDomainLogResponseBody) SetDomainName(v string) *DescribeLiveDomainLogResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBody) SetRequestId(v string) *DescribeLiveDomainLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBody) Validate() error {
	if s.DomainLogDetails != nil {
		if err := s.DomainLogDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeLiveDomainLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetails) Validate() error {
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

type DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	LogCount  *int64                                                                     `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
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

type DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
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

type DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeLiveDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
