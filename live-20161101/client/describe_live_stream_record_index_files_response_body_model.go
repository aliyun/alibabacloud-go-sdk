// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordIndexFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *DescribeLiveStreamRecordIndexFilesResponseBody
	GetOrder() *string
	SetPageNum(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody
	GetPageSize() *int32
	SetRecordIndexInfoList(v *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList) *DescribeLiveStreamRecordIndexFilesResponseBody
	GetRecordIndexInfoList() *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList
	SetRequestId(v string) *DescribeLiveStreamRecordIndexFilesResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody
	GetTotalPage() *int32
}

type DescribeLiveStreamRecordIndexFilesResponseBody struct {
	// The sort order.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize            *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RecordIndexInfoList *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList `json:"RecordIndexInfoList,omitempty" xml:"RecordIndexInfoList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DE24625C-7C0F-4020-8448-9C31A50C1556
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 12
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 20
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) GetRecordIndexInfoList() *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList {
	return s.RecordIndexInfoList
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) SetOrder(v string) *DescribeLiveStreamRecordIndexFilesResponseBody {
	s.Order = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) SetPageNum(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) SetPageSize(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) SetRecordIndexInfoList(v *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList) *DescribeLiveStreamRecordIndexFilesResponseBody {
	s.RecordIndexInfoList = v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) SetRequestId(v string) *DescribeLiveStreamRecordIndexFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) SetTotalNum(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) SetTotalPage(v int32) *DescribeLiveStreamRecordIndexFilesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBody) Validate() error {
	if s.RecordIndexInfoList != nil {
		if err := s.RecordIndexInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList struct {
	RecordIndexInfo []*DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo `json:"RecordIndexInfo,omitempty" xml:"RecordIndexInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList) GetRecordIndexInfo() []*DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	return s.RecordIndexInfo
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList) SetRecordIndexInfo(v []*DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList {
	s.RecordIndexInfo = v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoList) Validate() error {
	if s.RecordIndexInfo != nil {
		for _, item := range s.RecordIndexInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo struct {
	AppName     *string  `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime  *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainName  *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Duration    *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime     *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Format      *string  `json:"Format,omitempty" xml:"Format,omitempty"`
	Height      *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	OssBucket   *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string  `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObject   *string  `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	RecordId    *string  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RecordUrl   *string  `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	StartTime   *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName  *string  `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	Width       *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetDuration() *float32 {
	return s.Duration
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetFormat() *string {
	return s.Format
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetOssObject() *string {
	return s.OssObject
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetRecordUrl() *string {
	return s.RecordUrl
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) GetWidth() *int32 {
	return s.Width
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetAppName(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetCreateTime(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetDomainName(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetDuration(v float32) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.Duration = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetEndTime(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetFormat(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.Format = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetHeight(v int32) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetOssBucket(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetOssEndpoint(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetOssObject(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetRecordId(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.RecordId = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetRecordUrl(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.RecordUrl = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetStartTime(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetStreamName(v string) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) SetWidth(v int32) *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo {
	s.Width = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponseBodyRecordIndexInfoListRecordIndexInfo) Validate() error {
	return dara.Validate(s)
}
