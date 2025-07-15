// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamSnapshotConfigList(v *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList) *DescribeLiveSnapshotConfigResponseBody
	GetLiveStreamSnapshotConfigList() *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList
	SetOrder(v string) *DescribeLiveSnapshotConfigResponseBody
	GetOrder() *string
	SetPageNum(v int32) *DescribeLiveSnapshotConfigResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveSnapshotConfigResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveSnapshotConfigResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveSnapshotConfigResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveSnapshotConfigResponseBody
	GetTotalPage() *int32
}

type DescribeLiveSnapshotConfigResponseBody struct {
	// The snapshot configurations.
	LiveStreamSnapshotConfigList *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList `json:"LiveStreamSnapshotConfigList,omitempty" xml:"LiveStreamSnapshotConfigList,omitempty" type:"Struct"`
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
	// 2
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 11
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3136B58-5876-4168-83CA-B562781981A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 6
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveSnapshotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigResponseBody) GetLiveStreamSnapshotConfigList() *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList {
	return s.LiveStreamSnapshotConfigList
}

func (s *DescribeLiveSnapshotConfigResponseBody) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveSnapshotConfigResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveSnapshotConfigResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveSnapshotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveSnapshotConfigResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveSnapshotConfigResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveSnapshotConfigResponseBody) SetLiveStreamSnapshotConfigList(v *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList) *DescribeLiveSnapshotConfigResponseBody {
	s.LiveStreamSnapshotConfigList = v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBody) SetOrder(v string) *DescribeLiveSnapshotConfigResponseBody {
	s.Order = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBody) SetPageNum(v int32) *DescribeLiveSnapshotConfigResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBody) SetPageSize(v int32) *DescribeLiveSnapshotConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBody) SetRequestId(v string) *DescribeLiveSnapshotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBody) SetTotalNum(v int32) *DescribeLiveSnapshotConfigResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBody) SetTotalPage(v int32) *DescribeLiveSnapshotConfigResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList struct {
	LiveStreamSnapshotConfig []*DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig `json:"LiveStreamSnapshotConfig,omitempty" xml:"LiveStreamSnapshotConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList) GetLiveStreamSnapshotConfig() []*DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	return s.LiveStreamSnapshotConfig
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList) SetLiveStreamSnapshotConfig(v []*DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList {
	s.LiveStreamSnapshotConfig = v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The callback URL that is used to receive notifications about snapshot capture.
	//
	// example:
	//
	// https://learn.aliyundoc.com
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The time when the configuration was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-05-20T01:33:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// liveBucket****
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The endpoint of the OSS bucket.
	//
	// example:
	//
	// cn-oss-****.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The naming format of snapshots that are stored in the overwrite mode, which means that a new snapshot overwrites the previous snapshot.
	//
	// example:
	//
	// {liveApp****}/{liveStream****}.jpg
	OverwriteOssObject *string `json:"OverwriteOssObject,omitempty" xml:"OverwriteOssObject,omitempty"`
	// The naming format of snapshots that are stored in sequence, which means that a new snapshot does not overwrite the previous snapshot.
	//
	// example:
	//
	// snapshot/{liveApp****}/{liveStream****}/{UnixTimestamp}.jpg
	SequenceOssObject *string `json:"SequenceOssObject,omitempty" xml:"SequenceOssObject,omitempty"`
	// The interval at which snapshots are captured. Unit: seconds.
	//
	// example:
	//
	// 10
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetCallback() *string {
	return s.Callback
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetOverwriteOssObject() *string {
	return s.OverwriteOssObject
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetSequenceOssObject() *string {
	return s.SequenceOssObject
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetAppName(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetCallback(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.Callback = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetCreateTime(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetDomainName(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetOssBucket(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetOssEndpoint(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetOverwriteOssObject(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.OverwriteOssObject = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetSequenceOssObject(v string) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.SequenceOssObject = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) SetTimeInterval(v int32) *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig {
	s.TimeInterval = &v
	return s
}

func (s *DescribeLiveSnapshotConfigResponseBodyLiveStreamSnapshotConfigListLiveStreamSnapshotConfig) Validate() error {
	return dara.Validate(s)
}
