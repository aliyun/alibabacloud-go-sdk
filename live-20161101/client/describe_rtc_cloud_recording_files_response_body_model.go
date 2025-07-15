// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcCloudRecordingFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRtcCloudRecordingFilesResponseBody
	GetRequestId() *string
	SetTaskInfo(v *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) *DescribeRtcCloudRecordingFilesResponseBody
	GetTaskInfo() *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo
}

type DescribeRtcCloudRecordingFilesResponseBody struct {
	// example:
	//
	// ******58-5876-****-83CA-B56278******
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeRtcCloudRecordingFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudRecordingFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudRecordingFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcCloudRecordingFilesResponseBody) GetTaskInfo() *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *DescribeRtcCloudRecordingFilesResponseBody) SetRequestId(v string) *DescribeRtcCloudRecordingFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBody) SetTaskInfo(v *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) *DescribeRtcCloudRecordingFilesResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcCloudRecordingFilesResponseBodyTaskInfo struct {
	RecordFileList *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList `json:"RecordFileList,omitempty" xml:"RecordFileList,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) GetRecordFileList() *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList {
	return s.RecordFileList
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) SetRecordFileList(v *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo {
	s.RecordFileList = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) SetStatus(v string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) SetTaskId(v string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList struct {
	HlsFileList []*string `json:"HlsFileList,omitempty" xml:"HlsFileList,omitempty" type:"Repeated"`
	Mp3FileList []*string `json:"Mp3FileList,omitempty" xml:"Mp3FileList,omitempty" type:"Repeated"`
	Mp4FileList []*string `json:"Mp4FileList,omitempty" xml:"Mp4FileList,omitempty" type:"Repeated"`
}

func (s DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) GetHlsFileList() []*string {
	return s.HlsFileList
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) GetMp3FileList() []*string {
	return s.Mp3FileList
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) GetMp4FileList() []*string {
	return s.Mp4FileList
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) SetHlsFileList(v []*string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList {
	s.HlsFileList = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) SetMp3FileList(v []*string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList {
	s.Mp3FileList = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) SetMp4FileList(v []*string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList {
	s.Mp4FileList = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) Validate() error {
	return dara.Validate(s)
}
