// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLogExTtlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails) *DescribeLiveDomainLogExTtlResponseBody
	GetDomainLogDetails() *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails
	SetDomainName(v string) *DescribeLiveDomainLogExTtlResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeLiveDomainLogExTtlResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainLogExTtlResponseBody struct {
	DomainLogDetails *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	DomainName       *string                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainLogExTtlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlResponseBody) GetDomainLogDetails() *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeLiveDomainLogExTtlResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainLogExTtlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainLogExTtlResponseBody) SetDomainLogDetails(v *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails) *DescribeLiveDomainLogExTtlResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBody) SetDomainName(v string) *DescribeLiveDomainLogExTtlResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBody) SetRequestId(v string) *DescribeLiveDomainLogExTtlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBody) Validate() error {
	if s.DomainLogDetails != nil {
		if err := s.DomainLogDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetails) Validate() error {
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

type DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail struct {
	LogCount  *int64                                                                          `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos  *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
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

type DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
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

type DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageIndex(v int64) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageIndex = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
