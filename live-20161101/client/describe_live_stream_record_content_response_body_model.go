// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordContentInfoList(v *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList) *DescribeLiveStreamRecordContentResponseBody
	GetRecordContentInfoList() *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList
	SetRequestId(v string) *DescribeLiveStreamRecordContentResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamRecordContentResponseBody struct {
	RecordContentInfoList *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList `json:"RecordContentInfoList,omitempty" xml:"RecordContentInfoList,omitempty" type:"Struct"`
	// The end of the time range to query. The time range that is specified by the StartTime and EndTime parameters cannot exceed 4 days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 62136AE6-7793-45ED-B14A-60D19A9486D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamRecordContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentResponseBody) GetRecordContentInfoList() *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList {
	return s.RecordContentInfoList
}

func (s *DescribeLiveStreamRecordContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamRecordContentResponseBody) SetRecordContentInfoList(v *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList) *DescribeLiveStreamRecordContentResponseBody {
	s.RecordContentInfoList = v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBody) SetRequestId(v string) *DescribeLiveStreamRecordContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBody) Validate() error {
	if s.RecordContentInfoList != nil {
		if err := s.RecordContentInfoList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList struct {
	RecordContentInfo []*DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo `json:"RecordContentInfo,omitempty" xml:"RecordContentInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList) GetRecordContentInfo() []*DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo {
	return s.RecordContentInfo
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList) SetRecordContentInfo(v []*DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList {
	s.RecordContentInfo = v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoList) Validate() error {
	if s.RecordContentInfo != nil {
		for _, item := range s.RecordContentInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo struct {
	Duration        *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime         *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OssBucket       *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint     *string  `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObjectPrefix *string  `json:"OssObjectPrefix,omitempty" xml:"OssObjectPrefix,omitempty"`
	StartTime       *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) GetDuration() *float32 {
	return s.Duration
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) GetOssObjectPrefix() *string {
	return s.OssObjectPrefix
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) SetDuration(v float32) *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo {
	s.Duration = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) SetEndTime(v string) *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) SetOssBucket(v string) *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) SetOssEndpoint(v string) *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) SetOssObjectPrefix(v string) *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo {
	s.OssObjectPrefix = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) SetStartTime(v string) *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponseBodyRecordContentInfoListRecordContentInfo) Validate() error {
	return dara.Validate(s)
}
