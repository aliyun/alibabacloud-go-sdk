// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *DescribeRecordingResponseBody
	GetFileName() *string
	SetFilePath(v string) *DescribeRecordingResponseBody
	GetFilePath() *string
	SetRequestId(v string) *DescribeRecordingResponseBody
	GetRequestId() *string
	SetVoiceSliceRecordingListJson(v string) *DescribeRecordingResponseBody
	GetVoiceSliceRecordingListJson() *string
}

type DescribeRecordingResponseBody struct {
	// example:
	//
	// 2019080913202222.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// url
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId                   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VoiceSliceRecordingListJson *string `json:"VoiceSliceRecordingListJson,omitempty" xml:"VoiceSliceRecordingListJson,omitempty"`
}

func (s DescribeRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordingResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *DescribeRecordingResponseBody) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordingResponseBody) GetVoiceSliceRecordingListJson() *string {
	return s.VoiceSliceRecordingListJson
}

func (s *DescribeRecordingResponseBody) SetFileName(v string) *DescribeRecordingResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetFilePath(v string) *DescribeRecordingResponseBody {
	s.FilePath = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetRequestId(v string) *DescribeRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordingResponseBody) SetVoiceSliceRecordingListJson(v string) *DescribeRecordingResponseBody {
	s.VoiceSliceRecordingListJson = &v
	return s
}

func (s *DescribeRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}
