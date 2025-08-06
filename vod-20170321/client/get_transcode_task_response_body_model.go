// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistJobIds(v []*string) *GetTranscodeTaskResponseBody
	GetNonExistJobIds() []*string
	SetRequestId(v string) *GetTranscodeTaskResponseBody
	GetRequestId() *string
	SetTranscodeJobInfoList(v []*GetTranscodeTaskResponseBodyTranscodeJobInfoList) *GetTranscodeTaskResponseBody
	GetTranscodeJobInfoList() []*GetTranscodeTaskResponseBodyTranscodeJobInfoList
	SetTranscodeTask(v *GetTranscodeTaskResponseBodyTranscodeTask) *GetTranscodeTaskResponseBody
	GetTranscodeTask() *GetTranscodeTaskResponseBodyTranscodeTask
}

type GetTranscodeTaskResponseBody struct {
	// The nonexistent job ID.
	NonExistJobIds []*string `json:"NonExistJobIds,omitempty" xml:"NonExistJobIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F4C6D5BE-BF13-45*****6C-516EA8906DCD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of transcode job information.
	TranscodeJobInfoList []*GetTranscodeTaskResponseBodyTranscodeJobInfoList `json:"TranscodeJobInfoList,omitempty" xml:"TranscodeJobInfoList,omitempty" type:"Repeated"`
	// Details about transcoding tasks.
	TranscodeTask *GetTranscodeTaskResponseBodyTranscodeTask `json:"TranscodeTask,omitempty" xml:"TranscodeTask,omitempty" type:"Struct"`
}

func (s GetTranscodeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBody) GetNonExistJobIds() []*string {
	return s.NonExistJobIds
}

func (s *GetTranscodeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranscodeTaskResponseBody) GetTranscodeJobInfoList() []*GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	return s.TranscodeJobInfoList
}

func (s *GetTranscodeTaskResponseBody) GetTranscodeTask() *GetTranscodeTaskResponseBodyTranscodeTask {
	return s.TranscodeTask
}

func (s *GetTranscodeTaskResponseBody) SetNonExistJobIds(v []*string) *GetTranscodeTaskResponseBody {
	s.NonExistJobIds = v
	return s
}

func (s *GetTranscodeTaskResponseBody) SetRequestId(v string) *GetTranscodeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeTaskResponseBody) SetTranscodeJobInfoList(v []*GetTranscodeTaskResponseBodyTranscodeJobInfoList) *GetTranscodeTaskResponseBody {
	s.TranscodeJobInfoList = v
	return s
}

func (s *GetTranscodeTaskResponseBody) SetTranscodeTask(v *GetTranscodeTaskResponseBodyTranscodeTask) *GetTranscodeTaskResponseBody {
	s.TranscodeTask = v
	return s
}

func (s *GetTranscodeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeTaskResponseBodyTranscodeJobInfoList struct {
	// The complete time of the transcoding job. The format is yyyy-MM-dd\\"T\\"HH:mm:ssZ (UTC time).
	//
	// example:
	//
	// 2019-02-26T08:30:16Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The creation time of the transcoding job. The format is yyyy-MM-dd\\"T\\"HH:mm:ssZ (UTC time).
	//
	// example:
	//
	// 2019-02-26T08:27:16Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The clarity and audio quality types are defined as follows:
	//
	// - SD: Standard Definition.
	//
	// - HD: High Definition.
	//
	// - FHD: Full High Definition.
	//
	// - OD: Original Definition.
	//
	// - 2K: 2K.
	//
	// - 4K: 4K.
	//
	// - SQ: Standard Audio Quality.
	//
	// - HQ: High Audio Quality.
	//
	// - AUTO: Adaptive Bitrate.  This is only available when the transcoding template is configured with packaging settings. Please refer to [the Transcoding Template Configuration - Package Setting](https://api.aliyun-inc.com/~~52839~~?spm=openapi-amp.newDocPublishment.0.0.65b0281fNUFIXC) for more details.
	//
	// > This value represents the clarity label configured in the transcoding template and does not indicate the actual resolution range of the transcoded output file.
	//
	// example:
	//
	// LD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// ErrorMessage
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The OSS address of the source file for transcoding.
	//
	// example:
	//
	// http://outin-40564*****e1403e7.oss-cn-shanghai.aliyuncs.com/customerTrans/5b95e568f8e*****47f38e/31f1184c-*****b2a2-f94-c213f.wmv
	InputFileUrl *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	// Information about the transcoded output files.
	OutputFile *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	// The priority of the transcoding task.
	//
	// example:
	//
	// 6
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the transcode job.
	//
	// example:
	//
	// 38f0e513c88*****85515f9d50be188
	TranscodeJobId *string `json:"TranscodeJobId,omitempty" xml:"TranscodeJobId,omitempty"`
	// The status of the transcoding job:
	//
	// Transcoding: Transcoding in progress.
	//
	// TranscodeSuccess: Transcoding successful.
	//
	// TranscodeFail: Transcoding failed.
	//
	// example:
	//
	// Transcoding
	TranscodeJobStatus *string `json:"TranscodeJobStatus,omitempty" xml:"TranscodeJobStatus,omitempty"`
	// The processing progress of the transcoding job. The value range is [0, 100].
	//
	// example:
	//
	// 80
	TranscodeProgress *int64 `json:"TranscodeProgress,omitempty" xml:"TranscodeProgress,omitempty"`
	// The ID of the template used for the transcode job.
	//
	// example:
	//
	// 174b0534fea3*****b51c8f0ad1374
	TranscodeTemplateId *string `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeJobInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeJobInfoList) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetDefinition() *string {
	return s.Definition
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetInputFileUrl() *string {
	return s.InputFileUrl
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetOutputFile() *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	return s.OutputFile
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetPriority() *string {
	return s.Priority
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetTranscodeJobId() *string {
	return s.TranscodeJobId
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetTranscodeJobStatus() *string {
	return s.TranscodeJobStatus
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetTranscodeProgress() *int64 {
	return s.TranscodeProgress
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) GetTranscodeTemplateId() *string {
	return s.TranscodeTemplateId
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetCompleteTime(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetCreationTime(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetDefinition(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.Definition = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetErrorCode(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.ErrorCode = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetErrorMessage(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetInputFileUrl(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.InputFileUrl = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetOutputFile(v *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.OutputFile = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetPriority(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.Priority = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetTranscodeJobId(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.TranscodeJobId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetTranscodeJobStatus(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.TranscodeJobStatus = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetTranscodeProgress(v int64) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.TranscodeProgress = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) SetTranscodeTemplateId(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoList {
	s.TranscodeTemplateId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoList) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile struct {
	// List of audio streams.
	//
	// example:
	//
	// "AudioStreamList": "[{\\"Bitrate\\":\\"64.533\\",\\"ChannelLayout\\":\\"stereo\\",\\"Channels\\":\\"2\\",\\"CodecLongName\\":\\"AAC (Advanced Audio Coding)\\",\\"CodecName\\":\\"aac\\",\\"CodecTag\\":\\"0x6134706d\\",\\"CodecTagString\\":\\"mp4a\\",\\"CodecTimeBase\\":\\"1/44100\\",\\"Duration\\":\\"12.615533\\",\\"Index\\":\\"1\\",\\"Lang\\":\\"und\\",\\"SampleFmt\\":\\"fltp\\",\\"Samplerate\\":\\"44100\\",\\"StartTime\\":\\"-0.046440\\",\\"Timebase\\":\\"1/44100\\"}]
	AudioStreamList *string `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty"`
	// Average bitrate of the transcoded output file. Unit: Kbps.
	//
	// example:
	//
	// 964
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// Duration of the transcoded output file. Unit: seconds (s).
	//
	// example:
	//
	// 12
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Encryption configuration used for the transcoded output file. Values:
	//
	// - AliyunVoDEncryption: Alibaba Cloud Video Encryption (private encryption).
	//
	// - HLSEncryption: HLS standard encryption.
	//
	// example:
	//
	// {\\"EncryptType\\":\\"AliyunVoDEncryption\\"}
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// Size of the transcoded output file. Unit: bytes (B).
	//
	// example:
	//
	// 851076
	Filesize *int64 `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	// Container format of the transcoded output file.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Frame rate of the transcoded output file. Unit: frames per second (fps).
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Height of the video frame in the transcoded output file. Unit: pixels (px).
	//
	// example:
	//
	// 360
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// OSS address of the transcoded output file.
	//
	// example:
	//
	// http://outin-40564*****e1403e7.oss-cn-shanghai.aliyuncs.com/883f5d*****f20aaa352f/c3be4f073*****7d5193ec8-{DestMd5}-od-S00000001-200000.mp4
	OutputFileUrl *string `json:"OutputFileUrl,omitempty" xml:"OutputFileUrl,omitempty"`
	// List of subtitle streams.
	//
	// example:
	//
	// []
	SubtitleStreamList *string `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty"`
	// List of video streams.
	//
	// example:
	//
	// [{\\"AvgFPS\\":\\"30.0\\",\\"Bitrate\\":\\"933.814\\",\\"CodecLongName\\":\\"H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10\\",\\"CodecName\\":\\"h264\\",\\"CodecTag\\":\\"0x31637661\\",\\"CodecTagString\\":\\"avc1\\",\\"CodecTimeBase\\":\\"1/60\\",\\"Dar\\":\\"9:16\\",\\"Duration\\":\\"12.033333\\",\\"Fps\\":\\"30.0\\",\\"HasBFrames\\":\\"2\\",\\"Height\\":\\"360\\",\\"Index\\":\\"0\\",\\"Lang\\":\\"und\\",\\"Level\\":\\"30\\",\\"PixFmt\\":\\"yuv420p\\",\\"Profile\\":\\"High\\",\\"Sar\\":\\"81:256\\",\\"StartTime\\":\\"0.000000\\",\\"Timebase\\":\\"1/15360\\",\\"Width\\":\\"640\\"}]
	VideoStreamList *string `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty"`
	// List of watermarks used for transcoding.
	WatermarkIdList []*string `json:"WatermarkIdList,omitempty" xml:"WatermarkIdList,omitempty" type:"Repeated"`
	// Width of the video frame in the transcoded output file. Unit: pixels (px).
	//
	// example:
	//
	// 640
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetAudioStreamList() *string {
	return s.AudioStreamList
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetEncryption() *string {
	return s.Encryption
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetFilesize() *int64 {
	return s.Filesize
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetFormat() *string {
	return s.Format
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetFps() *string {
	return s.Fps
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetOutputFileUrl() *string {
	return s.OutputFileUrl
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetSubtitleStreamList() *string {
	return s.SubtitleStreamList
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetVideoStreamList() *string {
	return s.VideoStreamList
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetWatermarkIdList() []*string {
	return s.WatermarkIdList
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetAudioStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.AudioStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetBitrate(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetDuration(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Duration = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetEncryption(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Encryption = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetFilesize(v int64) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Filesize = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetFormat(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Format = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetFps(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Fps = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetHeight(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Height = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetOutputFileUrl(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.OutputFileUrl = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetSubtitleStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.SubtitleStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetVideoStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.VideoStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetWatermarkIdList(v []*string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.WatermarkIdList = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) SetWidth(v string) *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile {
	s.Width = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeJobInfoListOutputFile) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeTaskResponseBodyTranscodeTask struct {
	// The time when the transcoding task was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-23T12:40:12Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the transcoding task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-23T12:35:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The status of the transcoding task. Valid values:
	//
	// 	- **Processing**: In progress.
	//
	// 	- **Partial**: Some transcoding jobs were complete.
	//
	// 	- **CompleteAllSucc**: All transcoding jobs were successful.
	//
	// 	- **CompleteAllFail**: All transcoding jobs failed. If an exception occurs in the source file, no transcoding job is initiated and the transcoding task fails.
	//
	// 	- **CompletePartialSucc**: All transcoding jobs were complete but only some were successful.
	//
	// example:
	//
	// Processing
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Details about transcoding jobs.
	TranscodeJobInfoList []*GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList `json:"TranscodeJobInfoList,omitempty" xml:"TranscodeJobInfoList,omitempty" type:"Repeated"`
	// The ID of the transcoding task.
	//
	// example:
	//
	// b1b65ab107e14*****3dbb900f6c1fe0
	TranscodeTaskId *string `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// b500c7094bd241*****3e9900752d7c3
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The mode in which the transcoding task is triggered. Valid values:
	//
	// 	- **Auto**: The transcoding task is automatically triggered when the video is uploaded.
	//
	// 	- **Manual**: The transcoding task is triggered by calling the SubmitTranscodeJobs operation.
	//
	// example:
	//
	// Auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// The ID of the audio or video file.
	//
	// example:
	//
	// 883f5d98107*****b7f20aaa352f
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeTask) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeTask) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetTranscodeJobInfoList() []*GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	return s.TranscodeJobInfoList
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetTranscodeTaskId() *string {
	return s.TranscodeTaskId
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetTrigger() *string {
	return s.Trigger
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) GetVideoId() *string {
	return s.VideoId
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetCompleteTime(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetCreationTime(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTaskStatus(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TaskStatus = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTranscodeJobInfoList(v []*GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TranscodeJobInfoList = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTranscodeTaskId(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TranscodeTaskId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTranscodeTemplateGroupId(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTrigger(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.Trigger = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetVideoId(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.VideoId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList struct {
	// The time when the transcoding job was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-26T08:30:16Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the transcoding job was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-26T08:27:16Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The video resolution. Valid values:
	//
	// 	- **LD**: low definition
	//
	// 	- **SD**: standard definition
	//
	// 	- **HD**: high definition
	//
	// 	- **FHD**: ultra high definition
	//
	// 	- **OD**: original definition
	//
	// 	- **2K**: 2K
	//
	// 	- **4K**: 4K
	//
	// 	- **SQ**: standard sound quality
	//
	// 	- **HQ**: high sound quality
	//
	// 	- **AUTO**: adaptive bitrate Adaptive bitrate streams are returned only if PackageSetting is set in the transcoding template. For more information, see [Basic structures](https://help.aliyun.com/document_detail/52839.html).
	//
	// > This parameter indicates the definition that is configured in the transcoding template and does not indicate the actual resolution of the output video.
	//
	// example:
	//
	// LD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The error code returned when the transcoding job failed.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the transcoding job failed.
	//
	// example:
	//
	// ErrorMessage
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The Object Storage Service (OSS) URL of the input file.
	//
	// example:
	//
	// http://outin-40564*****e1403e7.oss-cn-shanghai.aliyuncs.com/customerTrans/5b95e568f8e*****47f38e/31f1184c-*****b2a2-f94-c213f.wmv
	InputFileUrl *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	// The information about the output file.
	OutputFile *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	// The priority of the transcoding job.
	//
	// example:
	//
	// 6
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the transcoding job.
	//
	// example:
	//
	// 38f0e513c88*****85515f9d50be188
	TranscodeJobId *string `json:"TranscodeJobId,omitempty" xml:"TranscodeJobId,omitempty"`
	// The status of the transcoding job.
	//
	// 	- **Transcoding**
	//
	// 	- **TranscodeSuccess**
	//
	// 	- **TranscodeFail**
	//
	// example:
	//
	// Transcoding
	TranscodeJobStatus *string `json:"TranscodeJobStatus,omitempty" xml:"TranscodeJobStatus,omitempty"`
	// The progress of the transcoding job. Valid values: `[0,100]`.
	//
	// example:
	//
	// 100
	TranscodeProgress *int64 `json:"TranscodeProgress,omitempty" xml:"TranscodeProgress,omitempty"`
	// The ID of the transcoding template.
	//
	// example:
	//
	// 174b0534fea3*****b51c8f0ad1374
	TranscodeTemplateId *string `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetDefinition() *string {
	return s.Definition
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetInputFileUrl() *string {
	return s.InputFileUrl
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetOutputFile() *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	return s.OutputFile
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetPriority() *string {
	return s.Priority
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetTranscodeJobId() *string {
	return s.TranscodeJobId
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetTranscodeJobStatus() *string {
	return s.TranscodeJobStatus
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetTranscodeProgress() *int64 {
	return s.TranscodeProgress
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GetTranscodeTemplateId() *string {
	return s.TranscodeTemplateId
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetCompleteTime(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetCreationTime(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetDefinition(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.Definition = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetErrorCode(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.ErrorCode = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetErrorMessage(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetInputFileUrl(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.InputFileUrl = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetOutputFile(v *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.OutputFile = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetPriority(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.Priority = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeJobId(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeJobId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeJobStatus(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeJobStatus = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeProgress(v int64) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeProgress = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeTemplateId(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeTemplateId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) Validate() error {
	return dara.Validate(s)
}

type GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile struct {
	// The audio streams.
	//
	// example:
	//
	// "AudioStreamList": "[{\\"Bitrate\\":\\"64.533\\",\\"ChannelLayout\\":\\"stereo\\",\\"Channels\\":\\"2\\",\\"CodecLongName\\":\\"AAC (Advanced Audio Coding)\\",\\"CodecName\\":\\"aac\\",\\"CodecTag\\":\\"0x6134706d\\",\\"CodecTagString\\":\\"mp4a\\",\\"CodecTimeBase\\":\\"1/44100\\",\\"Duration\\":\\"12.615533\\",\\"Index\\":\\"1\\",\\"Lang\\":\\"und\\",\\"SampleFmt\\":\\"fltp\\",\\"Samplerate\\":\\"44100\\",\\"StartTime\\":\\"-0.046440\\",\\"Timebase\\":\\"1/44100\\"}]
	AudioStreamList *string `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty"`
	// The average bitrate of the output file. Unit: Kbit/s.
	//
	// example:
	//
	// 964
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The length of the output file. Unit: seconds.
	//
	// example:
	//
	// 12
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The encryption method of the output file. Valid values:
	//
	// 	- **AliyunVoDEncryption**: Alibaba Cloud proprietary cryptography
	//
	// 	- **HLSEncryption**: HTTP Live Streaming (HLS) encryption
	//
	// example:
	//
	// {\\"EncryptType\\":\\"AliyunVoDEncryption\\"}
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// The size of the output file. Unit: byte.
	//
	// example:
	//
	// 851076
	Filesize *int64 `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	// The container format of the output file.
	//
	// example:
	//
	// m3u8
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The frame rate of the output file. Unit: frames per second.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the output video. Unit: pixels.
	//
	// example:
	//
	// 360
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The OSS URL of the output file.
	//
	// example:
	//
	// http://outin-40564*****e1403e7.oss-cn-shanghai.aliyuncs.com/883f5d*****f20aaa352f/c3be4f073*****7d5193ec8-{DestMd5}-od-S00000001-200000.mp4
	OutputFileUrl *string `json:"OutputFileUrl,omitempty" xml:"OutputFileUrl,omitempty"`
	// The subtitle streams.
	//
	// example:
	//
	// []
	SubtitleStreamList *string `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty"`
	// The video streams.
	//
	// example:
	//
	// [{\\"AvgFPS\\":\\"30.0\\",\\"Bitrate\\":\\"933.814\\",\\"CodecLongName\\":\\"H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10\\",\\"CodecName\\":\\"h264\\",\\"CodecTag\\":\\"0x31637661\\",\\"CodecTagString\\":\\"avc1\\",\\"CodecTimeBase\\":\\"1/60\\",\\"Dar\\":\\"9:16\\",\\"Duration\\":\\"12.033333\\",\\"Fps\\":\\"30.0\\",\\"HasBFrames\\":\\"2\\",\\"Height\\":\\"360\\",\\"Index\\":\\"0\\",\\"Lang\\":\\"und\\",\\"Level\\":\\"30\\",\\"PixFmt\\":\\"yuv420p\\",\\"Profile\\":\\"High\\",\\"Sar\\":\\"81:256\\",\\"StartTime\\":\\"0.000000\\",\\"Timebase\\":\\"1/15360\\",\\"Width\\":\\"640\\"}]
	VideoStreamList *string `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty"`
	// The IDs of the watermarks used by the output file.
	WatermarkIdList []*string `json:"WatermarkIdList,omitempty" xml:"WatermarkIdList,omitempty" type:"Repeated"`
	// The width of the output video. Unit: pixels.
	//
	// example:
	//
	// 640
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetAudioStreamList() *string {
	return s.AudioStreamList
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetBitrate() *string {
	return s.Bitrate
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetDuration() *string {
	return s.Duration
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetEncryption() *string {
	return s.Encryption
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetFilesize() *int64 {
	return s.Filesize
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetFormat() *string {
	return s.Format
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetFps() *string {
	return s.Fps
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetHeight() *string {
	return s.Height
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetOutputFileUrl() *string {
	return s.OutputFileUrl
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetSubtitleStreamList() *string {
	return s.SubtitleStreamList
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetVideoStreamList() *string {
	return s.VideoStreamList
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetWatermarkIdList() []*string {
	return s.WatermarkIdList
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GetWidth() *string {
	return s.Width
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetAudioStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.AudioStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetBitrate(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetDuration(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Duration = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetEncryption(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Encryption = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetFilesize(v int64) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Filesize = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetFormat(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Format = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetFps(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Fps = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetHeight(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Height = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetOutputFileUrl(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.OutputFileUrl = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetSubtitleStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.SubtitleStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetVideoStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.VideoStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetWatermarkIdList(v []*string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.WatermarkIdList = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetWidth(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Width = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) Validate() error {
	return dara.Validate(s)
}
