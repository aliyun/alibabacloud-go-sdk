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
	// The request ID.
	//
	// example:
	//
	// ******58-5876-****-83CA-B56278******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	TaskInfo *DescribeRtcCloudRecordingFilesResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
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
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcCloudRecordingFilesResponseBodyTaskInfo struct {
	// The list of recording files.
	RecordFileList *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList `json:"RecordFileList,omitempty" xml:"RecordFileList,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// - RUNNING
	//
	// - RECOVERING
	//
	// - STOPPING
	//
	// - STOPPED.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID passed in the request.
	//
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
	if s.RecordFileList != nil {
		if err := s.RecordFileList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList struct {
	// The array of HLS recording file names.
	HlsFileList []*string `json:"HlsFileList,omitempty" xml:"HlsFileList,omitempty" type:"Repeated"`
	// The array of MP3 recording file names.
	Mp3FileList []*string `json:"Mp3FileList,omitempty" xml:"Mp3FileList,omitempty" type:"Repeated"`
	// The array of MP4 recording file names.
	Mp4FileList []*string `json:"Mp4FileList,omitempty" xml:"Mp4FileList,omitempty" type:"Repeated"`
	// The array of VOD media resources. When recording to VOD, this is the collection of recording files for each subscribed stream, where each item corresponds to a subscribed stream.
	VodMediaList []*DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList `json:"VodMediaList,omitempty" xml:"VodMediaList,omitempty" type:"Repeated"`
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

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) GetVodMediaList() []*DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList {
	return s.VodMediaList
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

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) SetVodMediaList(v []*DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList {
	s.VodMediaList = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileList) Validate() error {
	if s.VodMediaList != nil {
		for _, item := range s.VodMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList struct {
	// The array of media resource IDs generated during recording.
	MediaIds []*string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty" type:"Repeated"`
	// The array of automatically merged media resource IDs generated after recording ends.
	MergedIds []*string `json:"MergedIds,omitempty" xml:"MergedIds,omitempty" type:"Repeated"`
	// The subscribed stream.
	//
	//  - For stream mixing recording, the value is always Mix.
	//
	//  - For single-stream recording, the value is Single::{UserId}::{Suffix}.
	//
	//    - UserId is the UserId corresponding to this stream.
	//
	//    - Suffix depends on the StreamType and SourceType specified during subscription.
	//
	//      - When StreamType is 0: if SourceType is 0, Suffix is AV::C. If SourceType is 1, Suffix is AV::S.
	//
	//      - When StreamType is 1: Suffix can only be A.
	//
	//      - When StreamType is 2 (not supported for single-stream recording): if SourceType is 0, Suffix is V::C. If SourceType is 1, Suffix is V::S.
	//
	// example:
	//
	// Single::UserA::AV::C
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) GetMediaIds() []*string {
	return s.MediaIds
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) GetMergedIds() []*string {
	return s.MergedIds
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) GetStream() *string {
	return s.Stream
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) SetMediaIds(v []*string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList {
	s.MediaIds = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) SetMergedIds(v []*string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList {
	s.MergedIds = v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) SetStream(v string) *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList {
	s.Stream = &v
	return s
}

func (s *DescribeRtcCloudRecordingFilesResponseBodyTaskInfoRecordFileListVodMediaList) Validate() error {
	return dara.Validate(s)
}
