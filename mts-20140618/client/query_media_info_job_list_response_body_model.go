// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaInfoJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfoJobList(v *QueryMediaInfoJobListResponseBodyMediaInfoJobList) *QueryMediaInfoJobListResponseBody
	GetMediaInfoJobList() *QueryMediaInfoJobListResponseBodyMediaInfoJobList
	SetNonExistMediaInfoJobIds(v *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds) *QueryMediaInfoJobListResponseBody
	GetNonExistMediaInfoJobIds() *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds
	SetRequestId(v string) *QueryMediaInfoJobListResponseBody
	GetRequestId() *string
}

type QueryMediaInfoJobListResponseBody struct {
	// The details of each returned media information analysis job.
	MediaInfoJobList *QueryMediaInfoJobListResponseBodyMediaInfoJobList `json:"MediaInfoJobList,omitempty" xml:"MediaInfoJobList,omitempty" type:"Struct"`
	// Nonexistent media information analysis jobs.
	NonExistMediaInfoJobIds *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds `json:"NonExistMediaInfoJobIds,omitempty" xml:"NonExistMediaInfoJobIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 46A04AA5-B119-41BB-B750-7C5327AC3E7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMediaInfoJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBody) GetMediaInfoJobList() *QueryMediaInfoJobListResponseBodyMediaInfoJobList {
	return s.MediaInfoJobList
}

func (s *QueryMediaInfoJobListResponseBody) GetNonExistMediaInfoJobIds() *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds {
	return s.NonExistMediaInfoJobIds
}

func (s *QueryMediaInfoJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaInfoJobListResponseBody) SetMediaInfoJobList(v *QueryMediaInfoJobListResponseBodyMediaInfoJobList) *QueryMediaInfoJobListResponseBody {
	s.MediaInfoJobList = v
	return s
}

func (s *QueryMediaInfoJobListResponseBody) SetNonExistMediaInfoJobIds(v *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds) *QueryMediaInfoJobListResponseBody {
	s.NonExistMediaInfoJobIds = v
	return s
}

func (s *QueryMediaInfoJobListResponseBody) SetRequestId(v string) *QueryMediaInfoJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBody) Validate() error {
	if s.MediaInfoJobList != nil {
		if err := s.MediaInfoJobList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistMediaInfoJobIds != nil {
		if err := s.NonExistMediaInfoJobIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobList struct {
	MediaInfoJob []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob `json:"MediaInfoJob,omitempty" xml:"MediaInfoJob,omitempty" type:"Repeated"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobList) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobList) GetMediaInfoJob() []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	return s.MediaInfoJob
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobList) SetMediaInfoJob(v []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) *QueryMediaInfoJobListResponseBodyMediaInfoJobList {
	s.MediaInfoJob = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobList) Validate() error {
	if s.MediaInfoJob != nil {
		for _, item := range s.MediaInfoJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob struct {
	// Indicates whether the job is in asynchronous mode.
	//
	// example:
	//
	// true
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The error code returned if the job fails.
	//
	// example:
	//
	// InvalidParameter.JsonObjectFormatInvalid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2014-01-10T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The information about the job input.
	Input *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the job.
	//
	// example:
	//
	// 23ca1d184c0e4341e5b665e2a12****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The message sent by MNS to notify you of the job result.
	MNSMessageResult *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The error message returned if the job fails.
	//
	// example:
	//
	// The parameter ”*” does not conform to the JSON Object specification
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the MPS queue.
	//
	// example:
	//
	// 88c6ca184c0e432bbf5b665e2a15****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The information about the input file. For more information, see [AliyunProperties](https://help.aliyun.com/document_detail/29251.html).
	Properties *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The status of the job.
	//
	// 	- **Analyzing**: The job is being run.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job fails.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The custom data.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetAsync() *bool {
	return s.Async
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetCode() *string {
	return s.Code
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetInput() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput {
	return s.Input
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetJobId() *string {
	return s.JobId
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetMNSMessageResult() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetMessage() *string {
	return s.Message
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetProperties() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	return s.Properties
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetState() *string {
	return s.State
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) GetUserData() *string {
	return s.UserData
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetAsync(v bool) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.Async = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetCode(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.Code = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetCreationTime(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.CreationTime = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetInput(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.Input = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetJobId(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.JobId = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetMNSMessageResult(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.MNSMessageResult = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetMessage(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.Message = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetPipelineId(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.PipelineId = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetProperties(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.Properties = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetState(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.State = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) SetUserData(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob {
	s.UserData = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJob) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.MNSMessageResult != nil {
		if err := s.MNSMessageResult.Validate(); err != nil {
			return err
		}
	}
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput struct {
	// The OSS bucket in which the input file is stored.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the OSS region.
	//
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the Object Storage Service (OSS) object that is used as the input file.
	//
	// example:
	//
	// example.mp4
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) GetObject() *string {
	return s.Object
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) SetBucket(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput {
	s.Bucket = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) SetLocation(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput {
	s.Location = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) SetObject(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput {
	s.Object = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult struct {
	// The error code returned if the job failed. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// InvalidParameter.JsonObjectFormatInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the job failed. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The parameter \\"Input\\" does not conform to the JSON Object specification
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the message returned if the job was successful. This parameter is not returned if the job fails.
	//
	// example:
	//
	// 123
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) SetErrorCode(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) SetErrorMessage(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) SetMessageId(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties struct {
	// The bitrate of the media file.
	//
	// example:
	//
	// 1630.045
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration of the media file.
	//
	// example:
	//
	// 17.226000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The format of the input media file.
	//
	// example:
	//
	// QuickTime/MOV
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	FileMd5    *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The size of the image file.
	//
	// example:
	//
	// 3509895
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The format information.
	Format *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	// The frame rate of the media file.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the video. Unit: pixel.
	//
	// example:
	//
	// 720
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The media streams that are contained in the input media file.
	Streams *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
	// The width of the video. Unit: pixel.
	//
	// example:
	//
	// 1280
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetFileFormat() *string {
	return s.FileFormat
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetFileMd5() *string {
	return s.FileMd5
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetFileSize() *string {
	return s.FileSize
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetFormat() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	return s.Format
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetStreams() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams {
	return s.Streams
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetBitrate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetDuration(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.Duration = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetFileFormat(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.FileFormat = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetFileMd5(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.FileMd5 = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetFileSize(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.FileSize = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetFormat(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.Format = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetFps(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.Fps = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetHeight(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.Height = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetStreams(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.Streams = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) SetWidth(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties {
	s.Width = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties) Validate() error {
	if s.Format != nil {
		if err := s.Format.Validate(); err != nil {
			return err
		}
	}
	if s.Streams != nil {
		if err := s.Streams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat struct {
	// The total bitrate.
	//
	// example:
	//
	// 1630.045
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The total duration.
	//
	// example:
	//
	// 17.226000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The full name of the container format.
	//
	// example:
	//
	// QuickTime/MOV
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	// The short name of the container format.
	//
	// example:
	//
	// mov
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The total number of program streams.
	//
	// example:
	//
	// 2
	NumPrograms *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	// The total number of media streams.
	//
	// example:
	//
	// 1
	NumStreams *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	// The size of the image file.
	//
	// example:
	//
	// 3509895
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.042000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetFormatName() *string {
	return s.FormatName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetNumPrograms() *string {
	return s.NumPrograms
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetNumStreams() *string {
	return s.NumStreams
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetSize() *string {
	return s.Size
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetBitrate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetDuration(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.Duration = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetFormatLongName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.FormatLongName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetFormatName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.FormatName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetNumPrograms(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.NumPrograms = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetNumStreams(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.NumStreams = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetSize(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.Size = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetStartTime(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.StartTime = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams struct {
	// The information about each audio stream.
	AudioStreamList *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	// The information about each subtitle stream.
	SubtitleStreamList *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	// The information about each video stream.
	VideoStreamList *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) GetAudioStreamList() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList {
	return s.AudioStreamList
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) GetSubtitleStreamList() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList {
	return s.SubtitleStreamList
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) GetVideoStreamList() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList {
	return s.VideoStreamList
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) SetAudioStreamList(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams {
	s.AudioStreamList = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) SetSubtitleStreamList(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams {
	s.SubtitleStreamList = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) SetVideoStreamList(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams {
	s.VideoStreamList = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams) Validate() error {
	if s.AudioStreamList != nil {
		if err := s.AudioStreamList.Validate(); err != nil {
			return err
		}
	}
	if s.SubtitleStreamList != nil {
		if err := s.SubtitleStreamList.Validate(); err != nil {
			return err
		}
	}
	if s.VideoStreamList != nil {
		if err := s.VideoStreamList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList struct {
	AudioStream []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty" type:"Repeated"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList) GetAudioStream() []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	return s.AudioStream
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList) SetAudioStream(v []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList {
	s.AudioStream = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList) Validate() error {
	if s.AudioStream != nil {
		for _, item := range s.AudioStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream struct {
	// The bitrate of the media file.
	//
	// example:
	//
	// 1536000
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The number of sound channels.
	//
	// example:
	//
	// 5.1(side)
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The output layout of the sound channels.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// DCA (DTS Coherent Acoustics)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Valid values:
	//
	// 	- **acc**
	//
	// 	- **mp3**
	//
	// 	- **mp4**
	//
	// 	- **ogg**
	//
	// 	- **flac**
	//
	// example:
	//
	// acc
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x0000
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// [0][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/48000
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration of the media file.
	//
	// example:
	//
	// 123
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationInaccurate *string `json:"DurationInaccurate,omitempty" xml:"DurationInaccurate,omitempty"`
	// The sequence number of the audio stream. The value indicates the position of the audio stream in all audio streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language. For more information, see [FFmpeg documentation](https://www.ffmpeg.org/ffmpeg-all.html?spm=a2c4g.11186623.2.66.243851cd2SntfN#Metadata).
	//
	// example:
	//
	// eng
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The total number of frames.
	//
	// example:
	//
	// 123
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The sampling format.
	//
	// example:
	//
	// fltp
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 48000
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.042000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/1000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetChannels() *string {
	return s.Channels
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetDurationInaccurate() *string {
	return s.DurationInaccurate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetSamplerate() *string {
	return s.Samplerate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetBitrate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetChannelLayout(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetChannels(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Channels = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecLongName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecTag(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTag = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecTagString(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecTimeBase(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetDuration(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Duration = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetDurationInaccurate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.DurationInaccurate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetIndex(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Index = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetLang(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetNumFrames(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.NumFrames = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetSampleFmt(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.SampleFmt = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetSamplerate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Samplerate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetStartTime(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.StartTime = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetTimebase(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Timebase = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList struct {
	SubtitleStream []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream `json:"SubtitleStream,omitempty" xml:"SubtitleStream,omitempty" type:"Repeated"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList) GetSubtitleStream() []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	return s.SubtitleStream
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList) SetSubtitleStream(v []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList {
	s.SubtitleStream = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList) Validate() error {
	if s.SubtitleStream != nil {
		for _, item := range s.SubtitleStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream struct {
	// The full name of the encoding format.
	//
	// example:
	//
	// ASS (Advanced SSA) subtitle
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Valid values:
	//
	// 	- **srt**
	//
	// 	- **ass**
	//
	// example:
	//
	// ass
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x0000
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// [0][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 0/1
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration. Unit: seconds.
	//
	// example:
	//
	// 1370.116000
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The sequence number of the subtitle stream. The value indicates the position of the subtitle stream in all subtitle streams.
	//
	// example:
	//
	// 3
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// eng
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/1000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecLongName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecTag(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecTag = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecTagString(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecTimeBase(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetDuration(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Duration = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetIndex(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Index = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetLang(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetStartTime(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.StartTime = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetTimebase(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Timebase = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList struct {
	VideoStream []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty" type:"Repeated"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList) GetVideoStream() []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	return s.VideoStream
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList) SetVideoStream(v []*QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList {
	s.VideoStream = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList) Validate() error {
	if s.VideoStream != nil {
		for _, item := range s.VideoStream {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream struct {
	// The average frame rate.
	//
	// example:
	//
	// 24000/1001
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate of the media file.
	//
	// example:
	//
	// 30541090
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// H.264/AVC/MPEG-4 AVC/MPEG-4 part 10
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Valid values:
	//
	// 	- **h264**
	//
	// 	- **h265**
	//
	// 	- **gif**
	//
	// 	- **webp**
	//
	// example:
	//
	// h264
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x0000
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// [0][0][0][0]
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1001/48000
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The level of color reconstruction.
	//
	// example:
	//
	// 700
	ColorPrimaries *string `json:"ColorPrimaries,omitempty" xml:"ColorPrimaries,omitempty"`
	// The color range.
	//
	// example:
	//
	// 700
	ColorRange *string `json:"ColorRange,omitempty" xml:"ColorRange,omitempty"`
	// The color channel.
	//
	// example:
	//
	// R255 G83 B170
	ColorTransfer *string `json:"ColorTransfer,omitempty" xml:"ColorTransfer,omitempty"`
	// The display aspect ratio (DAR).
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration of the media file.
	//
	// example:
	//
	// 100
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationInaccurate *string `json:"DurationInaccurate,omitempty" xml:"DurationInaccurate,omitempty"`
	// The frame rate of the media file.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// Indicates whether the video stream contains bidirectional frames (B-frames). A value of 1 indicates that the video stream contains B-frames. A value of 0 indicates that the video stream does not contain B-frames.
	//
	// example:
	//
	// 0
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// The height of the video stream in pixels.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the video stream. The value indicates the position of the video stream in all video streams.
	//
	// example:
	//
	// 1
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// The language.
	//
	// example:
	//
	// eng
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The codec level.
	//
	// example:
	//
	// 41
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network bandwidth that is consumed.
	NetworkCost *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	// The total number of frames.
	//
	// example:
	//
	// 100
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// The pixel format.
	//
	// example:
	//
	// yuv420p
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// The codec profile.
	//
	// example:
	//
	// High
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The rotation angle of the video.
	//
	// example:
	//
	// 180
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR).
	//
	// example:
	//
	// 1:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.042000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/1000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The width of the video in pixels.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetColorPrimaries() *string {
	return s.ColorPrimaries
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetColorRange() *string {
	return s.ColorRange
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetColorTransfer() *string {
	return s.ColorTransfer
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetDar() *string {
	return s.Dar
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetDurationInaccurate() *string {
	return s.DurationInaccurate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetFps() *string {
	return s.Fps
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetHeight() *string {
	return s.Height
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetIndex() *string {
	return s.Index
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetLang() *string {
	return s.Lang
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetLevel() *string {
	return s.Level
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetNetworkCost() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	return s.NetworkCost
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetPixFmt() *string {
	return s.PixFmt
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetProfile() *string {
	return s.Profile
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetRotate() *string {
	return s.Rotate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetSar() *string {
	return s.Sar
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetTimebase() *string {
	return s.Timebase
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetWidth() *string {
	return s.Width
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetAvgFPS(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.AvgFPS = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetBitrate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Bitrate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecLongName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecLongName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecName(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecName = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecTag(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTag = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecTagString(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTagString = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecTimeBase(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetColorPrimaries(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.ColorPrimaries = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetColorRange(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.ColorRange = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetColorTransfer(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.ColorTransfer = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetDar(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Dar = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetDuration(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Duration = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetDurationInaccurate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.DurationInaccurate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetFps(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Fps = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetHasBFrames(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.HasBFrames = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetHeight(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Height = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetIndex(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Index = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetLang(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Lang = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetLevel(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Level = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetNetworkCost(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.NetworkCost = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetNumFrames(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.NumFrames = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetPixFmt(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.PixFmt = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetProfile(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Profile = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetRotate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Rotate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetSar(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Sar = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetStartTime(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.StartTime = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetTimebase(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Timebase = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetWidth(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Width = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) Validate() error {
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost struct {
	// The average bitrate of the video stream.
	//
	// example:
	//
	// 300.34
	AvgBitrate *string `json:"AvgBitrate,omitempty" xml:"AvgBitrate,omitempty"`
	// The maximum bandwidth that is consumed.
	//
	// example:
	//
	// 10
	CostBandwidth *string `json:"CostBandwidth,omitempty" xml:"CostBandwidth,omitempty"`
	// The time consumed to preload the video.
	//
	// example:
	//
	// 8
	PreloadTime *string `json:"PreloadTime,omitempty" xml:"PreloadTime,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetAvgBitrate() *string {
	return s.AvgBitrate
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetCostBandwidth() *string {
	return s.CostBandwidth
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetPreloadTime() *string {
	return s.PreloadTime
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetAvgBitrate(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.AvgBitrate = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetCostBandwidth(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.CostBandwidth = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetPreloadTime(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.PreloadTime = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds) GetString_() []*string {
	return s.String_
}

func (s *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds) SetString_(v []*string) *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds {
	s.String_ = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyNonExistMediaInfoJobIds) Validate() error {
	return dara.Validate(s)
}
