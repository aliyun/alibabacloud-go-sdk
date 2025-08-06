// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLogModel(v *DescribeCdnDomainLogsResponseBodyDomainLogModel) *DescribeCdnDomainLogsResponseBody
	GetDomainLogModel() *DescribeCdnDomainLogsResponseBodyDomainLogModel
	SetPageNo(v int64) *DescribeCdnDomainLogsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeCdnDomainLogsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeCdnDomainLogsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeCdnDomainLogsResponseBody
	GetTotal() *int64
}

type DescribeCdnDomainLogsResponseBody struct {
	DomainLogModel *DescribeCdnDomainLogsResponseBodyDomainLogModel `json:"DomainLogModel,omitempty" xml:"DomainLogModel,omitempty" type:"Struct"`
	PageNo         *int64                                           `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total          *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBody) GetDomainLogModel() *DescribeCdnDomainLogsResponseBodyDomainLogModel {
	return s.DomainLogModel
}

func (s *DescribeCdnDomainLogsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeCdnDomainLogsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnDomainLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainLogsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeCdnDomainLogsResponseBody) SetDomainLogModel(v *DescribeCdnDomainLogsResponseBodyDomainLogModel) *DescribeCdnDomainLogsResponseBody {
	s.DomainLogModel = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetPageNo(v int64) *DescribeCdnDomainLogsResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetPageSize(v int64) *DescribeCdnDomainLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetRequestId(v string) *DescribeCdnDomainLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) SetTotal(v int64) *DescribeCdnDomainLogsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsResponseBodyDomainLogModel struct {
	DomainLogDetails *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	DomainName       *string                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModel) GetDomainLogDetails() *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails {
	return s.DomainLogDetails
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModel) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModel) SetDomainLogDetails(v *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) *DescribeCdnDomainLogsResponseBodyDomainLogModel {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModel) SetDomainName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModel {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModel) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails struct {
	DomainLogDetail []*DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) GetDomainLogDetail() []*DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	return s.DomainLogDetail
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) SetDomainLogDetail(v []*DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) GetLogName() *string {
	return s.LogName
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) GetLogPath() *string {
	return s.LogPath
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetEndTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetLogName(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.LogName = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetLogPath(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetLogSize(v int64) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) SetStartTime(v string) *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnDomainLogsResponseBodyDomainLogModelDomainLogDetailsDomainLogDetail) Validate() error {
	return dara.Validate(s)
}
