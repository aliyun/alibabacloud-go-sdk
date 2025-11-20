// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCloudRecordVideoResponseBody
	GetRequestId() *string
	SetVideoList(v []*QueryCloudRecordVideoResponseBodyVideoList) *QueryCloudRecordVideoResponseBody
	GetVideoList() []*QueryCloudRecordVideoResponseBodyVideoList
}

type QueryCloudRecordVideoResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	VideoList []*QueryCloudRecordVideoResponseBodyVideoList `json:"videoList,omitempty" xml:"videoList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCloudRecordVideoResponseBody) GetVideoList() []*QueryCloudRecordVideoResponseBodyVideoList {
	return s.VideoList
}

func (s *QueryCloudRecordVideoResponseBody) SetRequestId(v string) *QueryCloudRecordVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBody) SetVideoList(v []*QueryCloudRecordVideoResponseBodyVideoList) *QueryCloudRecordVideoResponseBody {
	s.VideoList = v
	return s
}

func (s *QueryCloudRecordVideoResponseBody) Validate() error {
	if s.VideoList != nil {
		for _, item := range s.VideoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCloudRecordVideoResponseBodyVideoList struct {
	// example:
	//
	// 59886
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1631172094000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1127942
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// faa1566c5bc24f21821ae2394f82db2e
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 290882268xxx1172033231
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 1
	RecordType *int64 `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1631172094000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryCloudRecordVideoResponseBodyVideoList) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoResponseBodyVideoList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetRecordId() *string {
	return s.RecordId
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetRecordType() *int64 {
	return s.RecordType
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) GetUserId() *string {
	return s.UserId
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetDuration(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.Duration = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetEndTime(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetFileSize(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.FileSize = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetMediaId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.MediaId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRecordId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RecordId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRecordType(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RecordType = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRegionId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RegionId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetStartTime(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetUserId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.UserId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) Validate() error {
	return dara.Validate(s)
}
