// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionLogBackupFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeCrossRegionLogBackupFilesResponseBody
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeCrossRegionLogBackupFilesResponseBody
	GetEndTime() *string
	SetItems(v *DescribeCrossRegionLogBackupFilesResponseBodyItems) *DescribeCrossRegionLogBackupFilesResponseBody
	GetItems() *DescribeCrossRegionLogBackupFilesResponseBodyItems
	SetPageNumber(v int32) *DescribeCrossRegionLogBackupFilesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeCrossRegionLogBackupFilesResponseBody
	GetPageRecordCount() *int32
	SetRegionId(v string) *DescribeCrossRegionLogBackupFilesResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeCrossRegionLogBackupFilesResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeCrossRegionLogBackupFilesResponseBody
	GetStartTime() *string
	SetTotalRecordCount(v int32) *DescribeCrossRegionLogBackupFilesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeCrossRegionLogBackupFilesResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-15T12:10:00Z
	EndTime *string                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Items   *DescribeCrossRegionLogBackupFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number. Pages start from page 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of cross-region backup files on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DAC241E8-28E6-49DA-BFB0-B2DD090885C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-05-30T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 100
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeCrossRegionLogBackupFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionLogBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetItems() *DescribeCrossRegionLogBackupFilesResponseBodyItems {
	return s.Items
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetDBInstanceId(v string) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetEndTime(v string) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetItems(v *DescribeCrossRegionLogBackupFilesResponseBodyItems) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetPageNumber(v int32) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetPageRecordCount(v int32) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetRegionId(v string) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetRequestId(v string) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetStartTime(v string) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) SetTotalRecordCount(v int32) *DescribeCrossRegionLogBackupFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCrossRegionLogBackupFilesResponseBodyItems struct {
	Item []*DescribeCrossRegionLogBackupFilesResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeCrossRegionLogBackupFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionLogBackupFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItems) GetItem() []*DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItems) SetItem(v []*DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) *DescribeCrossRegionLogBackupFilesResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCrossRegionLogBackupFilesResponseBodyItemsItem struct {
	CrossBackupRegion         *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	CrossDownloadLink         *string `json:"CrossDownloadLink,omitempty" xml:"CrossDownloadLink,omitempty"`
	CrossIntranetDownloadLink *string `json:"CrossIntranetDownloadLink,omitempty" xml:"CrossIntranetDownloadLink,omitempty"`
	CrossLogBackupId          *int32  `json:"CrossLogBackupId,omitempty" xml:"CrossLogBackupId,omitempty"`
	CrossLogBackupSize        *int64  `json:"CrossLogBackupSize,omitempty" xml:"CrossLogBackupSize,omitempty"`
	InstanceId                *int32  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LinkExpiredTime           *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
	LogBeginTime              *string `json:"LogBeginTime,omitempty" xml:"LogBeginTime,omitempty"`
	LogEndTime                *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
	LogFileName               *string `json:"LogFileName,omitempty" xml:"LogFileName,omitempty"`
}

func (s DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetCrossDownloadLink() *string {
	return s.CrossDownloadLink
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetCrossIntranetDownloadLink() *string {
	return s.CrossIntranetDownloadLink
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetCrossLogBackupId() *int32 {
	return s.CrossLogBackupId
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetCrossLogBackupSize() *int64 {
	return s.CrossLogBackupSize
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetLinkExpiredTime() *string {
	return s.LinkExpiredTime
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetLogBeginTime() *string {
	return s.LogBeginTime
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetLogEndTime() *string {
	return s.LogEndTime
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) GetLogFileName() *string {
	return s.LogFileName
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetCrossBackupRegion(v string) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetCrossDownloadLink(v string) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.CrossDownloadLink = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetCrossIntranetDownloadLink(v string) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.CrossIntranetDownloadLink = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetCrossLogBackupId(v int32) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.CrossLogBackupId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetCrossLogBackupSize(v int64) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.CrossLogBackupSize = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetInstanceId(v int32) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.InstanceId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetLinkExpiredTime(v string) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetLogBeginTime(v string) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.LogBeginTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetLogEndTime(v string) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.LogEndTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) SetLogFileName(v string) *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem {
	s.LogFileName = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponseBodyItemsItem) Validate() error {
	return dara.Validate(s)
}
