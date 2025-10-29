// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamSnapshotInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamSnapshotInfoList(v *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList) *DescribeLiveStreamSnapshotInfoResponseBody
	GetLiveStreamSnapshotInfoList() *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList
	SetNextStartTime(v string) *DescribeLiveStreamSnapshotInfoResponseBody
	GetNextStartTime() *string
	SetRequestId(v string) *DescribeLiveStreamSnapshotInfoResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamSnapshotInfoResponseBody struct {
	// The snapshots.
	LiveStreamSnapshotInfoList *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList `json:"LiveStreamSnapshotInfoList,omitempty" xml:"LiveStreamSnapshotInfoList,omitempty" type:"Struct"`
	// The time when the next call occurred. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// >  If the number of snapshots that were captured within the specified time period exceeds the value of the Limit parameter, this parameter is returned. It indicates the time when the DescribeLiveStreamSnapshotInfo operation was called again. If this parameter is not returned, the number of snapshots that are captured within the specified time period does not exceed the specified limit.
	//
	// example:
	//
	// 2015-12-01T17:36:00Z
	NextStartTime *string `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 62136AE6-7793-45ED-B14A-60D19A9486D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamSnapshotInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoResponseBody) GetLiveStreamSnapshotInfoList() *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList {
	return s.LiveStreamSnapshotInfoList
}

func (s *DescribeLiveStreamSnapshotInfoResponseBody) GetNextStartTime() *string {
	return s.NextStartTime
}

func (s *DescribeLiveStreamSnapshotInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamSnapshotInfoResponseBody) SetLiveStreamSnapshotInfoList(v *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList) *DescribeLiveStreamSnapshotInfoResponseBody {
	s.LiveStreamSnapshotInfoList = v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBody) SetNextStartTime(v string) *DescribeLiveStreamSnapshotInfoResponseBody {
	s.NextStartTime = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBody) SetRequestId(v string) *DescribeLiveStreamSnapshotInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBody) Validate() error {
	if s.LiveStreamSnapshotInfoList != nil {
		if err := s.LiveStreamSnapshotInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList struct {
	LiveStreamSnapshotInfo []*DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo `json:"LiveStreamSnapshotInfo,omitempty" xml:"LiveStreamSnapshotInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList) GetLiveStreamSnapshotInfo() []*DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	return s.LiveStreamSnapshotInfo
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList) SetLiveStreamSnapshotInfo(v []*DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList {
	s.LiveStreamSnapshotInfo = v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoList) Validate() error {
	if s.LiveStreamSnapshotInfo != nil {
		for _, item := range s.LiveStreamSnapshotInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo struct {
	// The time when the snapshot was captured. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-01T17:36:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp when the snapshot file was created. Unit: milliseconds.
	//
	// example:
	//
	// 1653641526637
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The snapshot mode. Valid values:
	//
	// 	- **true**: overwrite mode, which means that a new snapshot overwrites the previous snapshot.
	//
	// 	- **false**: sequence mode, which means that a new snapshot does not overwrite the previous snapshot.
	//
	// example:
	//
	// false
	IsOverlay *bool `json:"IsOverlay,omitempty" xml:"IsOverlay,omitempty"`
	// The name of the OSS bucket.
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
	// The name of the snapshot stored in Object Storage Service (OSS).
	//
	// example:
	//
	// {liveApp****}/{liveStream****}.jpg
	OssObject *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
}

func (s DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GetIsOverlay() *bool {
	return s.IsOverlay
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) GetOssObject() *string {
	return s.OssObject
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetCreateTime(v string) *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetCreateTimestamp(v int64) *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetIsOverlay(v bool) *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.IsOverlay = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetOssBucket(v string) *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetOssEndpoint(v string) *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) SetOssObject(v string) *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveStreamSnapshotInfoResponseBodyLiveStreamSnapshotInfoListLiveStreamSnapshotInfo) Validate() error {
	return dara.Validate(s)
}
