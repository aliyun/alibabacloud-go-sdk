// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioList(v []*QueryMinutesResponseBodyAudioList) *QueryMinutesResponseBody
	GetAudioList() []*QueryMinutesResponseBodyAudioList
	SetRequestId(v string) *QueryMinutesResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *QueryMinutesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryMinutesResponseBody
	GetVendorType() *string
}

type QueryMinutesResponseBody struct {
	AudioList []*QueryMinutesResponseBodyAudioList `json:"audioList,omitempty" xml:"audioList,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s QueryMinutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMinutesResponseBody) GetAudioList() []*QueryMinutesResponseBodyAudioList {
	return s.AudioList
}

func (s *QueryMinutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMinutesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryMinutesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryMinutesResponseBody) SetAudioList(v []*QueryMinutesResponseBodyAudioList) *QueryMinutesResponseBody {
	s.AudioList = v
	return s
}

func (s *QueryMinutesResponseBody) SetRequestId(v string) *QueryMinutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMinutesResponseBody) SetVendorRequestId(v string) *QueryMinutesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryMinutesResponseBody) SetVendorType(v string) *QueryMinutesResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryMinutesResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesResponseBodyAudioList struct {
	// example:
	//
	// 1000
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1000000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1127942
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// https://xxx-hangzhou.oss-cn-hangzhou.aliyuncs.com/record_xxxx.mp3?Expires=1718045081&OSSAccessKeyId=TMP.3KdwHtvZxopmwacMZEdyb4WHLVmbArrNRB9CTKnR1MaJgmRjdmZczs6Rip66cgKgk2HhQon1yygvBnbY3uqEaZNeHBLcBa&Signature=OFWyAIY%2FdlzfwM9wIfEaKoAudkxxxxx
	PlayUrl *string `json:"PlayUrl,omitempty" xml:"PlayUrl,omitempty"`
	// example:
	//
	// 123
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 1000000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryMinutesResponseBodyAudioList) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesResponseBodyAudioList) GoString() string {
	return s.String()
}

func (s *QueryMinutesResponseBodyAudioList) GetDuration() *int64 {
	return s.Duration
}

func (s *QueryMinutesResponseBodyAudioList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMinutesResponseBodyAudioList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *QueryMinutesResponseBodyAudioList) GetPlayUrl() *string {
	return s.PlayUrl
}

func (s *QueryMinutesResponseBodyAudioList) GetRecordId() *string {
	return s.RecordId
}

func (s *QueryMinutesResponseBodyAudioList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMinutesResponseBodyAudioList) GetUserId() *string {
	return s.UserId
}

func (s *QueryMinutesResponseBodyAudioList) SetDuration(v int64) *QueryMinutesResponseBodyAudioList {
	s.Duration = &v
	return s
}

func (s *QueryMinutesResponseBodyAudioList) SetEndTime(v int64) *QueryMinutesResponseBodyAudioList {
	s.EndTime = &v
	return s
}

func (s *QueryMinutesResponseBodyAudioList) SetFileSize(v int64) *QueryMinutesResponseBodyAudioList {
	s.FileSize = &v
	return s
}

func (s *QueryMinutesResponseBodyAudioList) SetPlayUrl(v string) *QueryMinutesResponseBodyAudioList {
	s.PlayUrl = &v
	return s
}

func (s *QueryMinutesResponseBodyAudioList) SetRecordId(v string) *QueryMinutesResponseBodyAudioList {
	s.RecordId = &v
	return s
}

func (s *QueryMinutesResponseBodyAudioList) SetStartTime(v int64) *QueryMinutesResponseBodyAudioList {
	s.StartTime = &v
	return s
}

func (s *QueryMinutesResponseBodyAudioList) SetUserId(v string) *QueryMinutesResponseBodyAudioList {
	s.UserId = &v
	return s
}

func (s *QueryMinutesResponseBodyAudioList) Validate() error {
	return dara.Validate(s)
}
