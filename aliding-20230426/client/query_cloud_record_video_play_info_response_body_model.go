// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoPlayInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int64) *QueryCloudRecordVideoPlayInfoResponseBody
	GetDuration() *int64
	SetFileSize(v int64) *QueryCloudRecordVideoPlayInfoResponseBody
	GetFileSize() *int64
	SetMp4FileUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody
	GetMp4FileUrl() *string
	SetPlayUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody
	GetPlayUrl() *string
	SetRequestId(v string) *QueryCloudRecordVideoPlayInfoResponseBody
	GetRequestId() *string
	SetStatus(v int64) *QueryCloudRecordVideoPlayInfoResponseBody
	GetStatus() *int64
}

type QueryCloudRecordVideoPlayInfoResponseBody struct {
	// example:
	//
	// 59886
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1127942
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// https://vod.mcs.dingtalk.com/faa1566c5bc24f21821ae2394f82db2e/8bbd1612e686462ab4717919f67bb721-b8531e0d534b2f9747a9fdfxxxxxxxxc-sd.mp4
	Mp4FileUrl *string `json:"mp4FileUrl,omitempty" xml:"mp4FileUrl,omitempty"`
	// example:
	//
	// https://vod.mcs.dingtalk.com/faa1566c5bc24f21821ae2394f82db2e/8bbd1612e686462ab4717919f67bb721-ab85cc044a163568c9485xxxxxxxx76d-sd.m3u8
	PlayUrl *string `json:"playUrl,omitempty" xml:"playUrl,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) GetFileSize() *int64 {
	return s.FileSize
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) GetMp4FileUrl() *string {
	return s.Mp4FileUrl
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) GetPlayUrl() *string {
	return s.PlayUrl
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetDuration(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetFileSize(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.FileSize = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetMp4FileUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Mp4FileUrl = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetPlayUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.PlayUrl = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetRequestId(v string) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetStatus(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Status = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
