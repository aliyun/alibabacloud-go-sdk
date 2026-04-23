// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogDetails(v *DescribeVodDomainLogResponseBodyDomainLogDetails) *DescribeVodDomainLogResponseBody
	GetDomainLogDetails() *DescribeVodDomainLogResponseBodyDomainLogDetails
	SetRequestId(v string) *DescribeVodDomainLogResponseBody
	GetRequestId() *string
}

type DescribeVodDomainLogResponseBody struct {
	DomainLogDetails *DescribeVodDomainLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 077D0284-F041-4A41-4D3C-B48377FD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBody) GetDomainLogDetails() *DescribeVodDomainLogResponseBodyDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeVodDomainLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainLogResponseBody) SetDomainLogDetails(v *DescribeVodDomainLogResponseBodyDomainLogDetails) *DescribeVodDomainLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeVodDomainLogResponseBody) SetRequestId(v string) *DescribeVodDomainLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainLogResponseBody) Validate() error {
	if s.DomainLogDetails != nil {
		if err := s.DomainLogDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetails) GetDomainLogDetail() []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeVodDomainLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetails) Validate() error {
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

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	DomainName *string                                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LogCount   *int64                                                                    `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos   *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos  *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetLogInfos() *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	return s.LogInfos
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GetPageInfos() *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	return s.PageInfos
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetDomainName(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) Validate() error {
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

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GetLogInfoDetail() []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	return s.LogInfoDetail
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) Validate() error {
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

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageNumber(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) Validate() error {
	return dara.Validate(s)
}
