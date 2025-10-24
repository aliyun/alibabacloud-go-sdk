// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaInfoJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaInfoJob(v *SubmitMediaInfoJobResponseBodyMediaInfoJob) *SubmitMediaInfoJobResponseBody
	GetMediaInfoJob() *SubmitMediaInfoJobResponseBodyMediaInfoJob
	SetRequestId(v string) *SubmitMediaInfoJobResponseBody
	GetRequestId() *string
}

type SubmitMediaInfoJobResponseBody struct {
	// The details of the media information analysis job.
	MediaInfoJob *SubmitMediaInfoJobResponseBodyMediaInfoJob `json:"MediaInfoJob,omitempty" xml:"MediaInfoJob,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 62D9BE16-B7D5-550C-A482-7A0F60E09877
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaInfoJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBody) GetMediaInfoJob() *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	return s.MediaInfoJob
}

func (s *SubmitMediaInfoJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaInfoJobResponseBody) SetMediaInfoJob(v *SubmitMediaInfoJobResponseBodyMediaInfoJob) *SubmitMediaInfoJobResponseBody {
	s.MediaInfoJob = v
	return s
}

func (s *SubmitMediaInfoJobResponseBody) SetRequestId(v string) *SubmitMediaInfoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBody) Validate() error {
	if s.MediaInfoJob != nil {
		if err := s.MediaInfoJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitMediaInfoJobResponseBodyMediaInfoJob struct {
	// Indicates whether the job is run in asynchronous mode.
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
	// The information about the input media file.
	Input *SubmitMediaInfoJobResponseBodyMediaInfoJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the job.
	//
	// example:
	//
	// 23ca1d184c0e4341e5b665e2a12****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The message sent by Message Service (MNS) to notify users of the job result.
	MNSMessageResult *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The error message returned if the job fails.
	//
	// example:
	//
	// The parameter ”*” does not conform to the JSON Object specification
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the MPS queue to which the analysis job is submitted.
	//
	// example:
	//
	// 88c6ca184c0e432bbf5b665e2a15****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The properties of the input media file.
	Properties *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The status of the job. Valid values:
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job fails.
	//
	// 	- **Analyzing**: The job is being run.
	//
	// example:
	//
	// Analyzing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The custom data.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJob) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetAsync() *bool {
	return s.Async
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetCode() *string {
	return s.Code
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetInput() *SubmitMediaInfoJobResponseBodyMediaInfoJobInput {
	return s.Input
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetMNSMessageResult() *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetProperties() *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	return s.Properties
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetState() *string {
	return s.State
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetAsync(v bool) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Async = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetCode(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Code = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetCreationTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetInput(v *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Input = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetJobId(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.JobId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetMNSMessageResult(v *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.MNSMessageResult = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetMessage(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Message = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetPipelineId(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetProperties(v *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.Properties = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetState(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.State = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) SetUserData(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJob {
	s.UserData = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJob) Validate() error {
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

type SubmitMediaInfoJobResponseBodyMediaInfoJobInput struct {
	// The name of the OSS bucket in which the input media file is stored.
	//
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The region in which the OSS bucket that stores the input media file resides.
	//
	// example:
	//
	// example-location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input media file.
	//
	// example:
	//
	// example.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobInput) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) GetLocation() *string {
	return s.Location
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) GetObject() *string {
	return s.Object
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) SetBucket(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Bucket = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) SetLocation(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Location = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) SetObject(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobInput {
	s.Object = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobInput) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult struct {
	// The error code that is returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The parameter \\"Input\\" does not conform to the JSON Object specification
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// InvalidParameter.JsonObjectFormatInvalid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the message that is returned if the job is successful. This parameter is not returned if the job fails.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) SetErrorCode(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) SetErrorMessage(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) SetMessageId(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobProperties struct {
	// The bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 1630.045
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration of the input media file. Unit: seconds.
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
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 3509895
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The format information.
	Format *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	// The frame rate.
	//
	// example:
	//
	// 25
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The height of the video. Unit: pixel.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	MD5    *string `json:"MD5,omitempty" xml:"MD5,omitempty"`
	// The media streams that are contained in the input media file.
	Streams *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
	// The width of the video. Unit: pixel.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetFileFormat() *string {
	return s.FileFormat
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetFileSize() *string {
	return s.FileSize
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetFormat() *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	return s.Format
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetFps() *string {
	return s.Fps
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetHeight() *string {
	return s.Height
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetMD5() *string {
	return s.MD5
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetStreams() *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams {
	return s.Streams
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) GetWidth() *string {
	return s.Width
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetBitrate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.Bitrate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetFileFormat(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.FileFormat = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetFileSize(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.FileSize = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetFormat(v *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.Format = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetFps(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.Fps = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetHeight(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.Height = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetMD5(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.MD5 = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetStreams(v *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.Streams = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) SetWidth(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties {
	s.Width = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobProperties) Validate() error {
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

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat struct {
	// The total bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 1630.045
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The duration of the input media file. Unit: seconds.
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
	// The short name of the container format. For more information about the parameters, see [Parameter details](https://help.aliyun.com/document_detail/29253.html).
	//
	// example:
	//
	// mov
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The total number of program streams.
	//
	// example:
	//
	// 0
	NumPrograms *string `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	// The total number of media streams.
	//
	// example:
	//
	// 2
	NumStreams *string `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 3509895
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time.
	//
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetFormatName() *string {
	return s.FormatName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetNumPrograms() *string {
	return s.NumPrograms
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetNumStreams() *string {
	return s.NumStreams
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetSize() *string {
	return s.Size
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetBitrate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.Bitrate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetFormatLongName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.FormatLongName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetFormatName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.FormatName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetNumPrograms(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.NumPrograms = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetNumStreams(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.NumStreams = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetSize(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.Size = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) SetStartTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat {
	s.StartTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesFormat) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams struct {
	// The audio streams. A media file can contain up to four audio streams.
	AudioStreamList *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	// The subtitle streams. A media file can contain up to four subtitle streams.
	SubtitleStreamList *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	// The video streams. A media file can contain up to four video streams.
	VideoStreamList *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) GetAudioStreamList() *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList {
	return s.AudioStreamList
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) GetSubtitleStreamList() *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList {
	return s.SubtitleStreamList
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) GetVideoStreamList() *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList {
	return s.VideoStreamList
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) SetAudioStreamList(v *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams {
	s.AudioStreamList = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) SetSubtitleStreamList(v *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams {
	s.SubtitleStreamList = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) SetVideoStreamList(v *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams {
	s.VideoStreamList = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreams) Validate() error {
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

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList struct {
	AudioStream []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty" type:"Repeated"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList) GetAudioStream() []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	return s.AudioStream
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList) SetAudioStream(v []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList {
	s.AudioStream = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamList) Validate() error {
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

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream struct {
	// The bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 128.806
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The output layout of the sound channels.
	//
	// example:
	//
	// stereo
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// The number of sound channels.
	//
	// example:
	//
	// 2
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// The full name of the encoding format.
	//
	// example:
	//
	// AAC (Advanced Audio Coding)
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// The short name of the encoding format. Default value: acc. Valid values:
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
	// aac
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// The tag of the encoding format.
	//
	// example:
	//
	// 0x6134706d
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// mp4a
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// The codec time base.
	//
	// example:
	//
	// 1/44100
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// The duration of the audio stream. Unit: seconds.
	//
	// example:
	//
	// 17.159546
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationInaccurate *string `json:"DurationInaccurate,omitempty" xml:"DurationInaccurate,omitempty"`
	// The sequence number of the audio stream. The value indicates the position of the audio stream in all audio streams.
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
	// The sampling rate. Unit: Hz.
	//
	// example:
	//
	// 44100
	Samplerate *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	// The start time of the audio stream.
	//
	// example:
	//
	// 0.000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/44100
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetChannelLayout() *string {
	return s.ChannelLayout
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetChannels() *string {
	return s.Channels
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetDurationInaccurate() *string {
	return s.DurationInaccurate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetIndex() *string {
	return s.Index
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetLang() *string {
	return s.Lang
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetSampleFmt() *string {
	return s.SampleFmt
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetSamplerate() *string {
	return s.Samplerate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetBitrate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Bitrate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetChannelLayout(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.ChannelLayout = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetChannels(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Channels = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecLongName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecLongName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecTag(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTag = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecTagString(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTagString = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetCodecTimeBase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetDurationInaccurate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.DurationInaccurate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetIndex(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Index = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetLang(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Lang = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetNumFrames(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.NumFrames = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetSampleFmt(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.SampleFmt = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetSamplerate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Samplerate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetStartTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.StartTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) SetTimebase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream {
	s.Timebase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsAudioStreamListAudioStream) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList struct {
	SubtitleStream []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream `json:"SubtitleStream,omitempty" xml:"SubtitleStream,omitempty" type:"Repeated"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList) GetSubtitleStream() []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	return s.SubtitleStream
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList) SetSubtitleStream(v []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList {
	s.SubtitleStream = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamList) Validate() error {
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

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream struct {
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
	// The duration of the audio stream. Unit: seconds.
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
	// The start time of the subtitle stream.
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

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetIndex() *string {
	return s.Index
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetLang() *string {
	return s.Lang
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecLongName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecLongName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecTag(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecTag = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecTagString(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecTagString = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetCodecTimeBase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetIndex(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Index = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetLang(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Lang = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetStartTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.StartTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) SetTimebase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream {
	s.Timebase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsSubtitleStreamListSubtitleStream) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList struct {
	VideoStream []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty" type:"Repeated"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList) GetVideoStream() []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	return s.VideoStream
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList) SetVideoStream(v []*SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList {
	s.VideoStream = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamList) Validate() error {
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

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream struct {
	// The average frame rate.
	//
	// example:
	//
	// 23.976025
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// The bitrate. Unit: Kbit/s.
	//
	// example:
	//
	// 1496.46
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
	// 0x31637661
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// The tag string of the encoding format.
	//
	// example:
	//
	// avc1
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
	// The display aspect ratio (DAR). DAR is the proportional relationship between the width and the height of a video. The value is used to determine whether the video is in portrait mode or landscape mode.
	//
	// example:
	//
	// 16:9
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// The duration of the video stream. Unit: seconds.
	//
	// example:
	//
	// 17.225542
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationInaccurate *string `json:"DurationInaccurate,omitempty" xml:"DurationInaccurate,omitempty"`
	// The frame rate.
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
	// The height of the video. Unit: pixel.
	//
	// example:
	//
	// 1080
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The sequence number of the video stream. The value indicates the position of the video stream in all video streams. The sequence number of the first video stream to be played can be specified in some players. Default value: 1.
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
	NetworkCost *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
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
	// 90
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The sample aspect ratio (SAR).
	//
	// example:
	//
	// 1:1
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// The start time of the video stream.
	//
	// example:
	//
	// 0.042000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time base.
	//
	// example:
	//
	// 1/24000
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// The width of the video. Unit: pixel.
	//
	// example:
	//
	// 1920
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetAvgFPS() *string {
	return s.AvgFPS
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetBitrate() *string {
	return s.Bitrate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecLongName() *string {
	return s.CodecLongName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecName() *string {
	return s.CodecName
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecTag() *string {
	return s.CodecTag
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecTagString() *string {
	return s.CodecTagString
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetCodecTimeBase() *string {
	return s.CodecTimeBase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetColorPrimaries() *string {
	return s.ColorPrimaries
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetColorRange() *string {
	return s.ColorRange
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetColorTransfer() *string {
	return s.ColorTransfer
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetDar() *string {
	return s.Dar
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetDuration() *string {
	return s.Duration
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetDurationInaccurate() *string {
	return s.DurationInaccurate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetFps() *string {
	return s.Fps
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetHasBFrames() *string {
	return s.HasBFrames
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetHeight() *string {
	return s.Height
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetIndex() *string {
	return s.Index
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetLang() *string {
	return s.Lang
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetLevel() *string {
	return s.Level
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetNetworkCost() *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	return s.NetworkCost
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetNumFrames() *string {
	return s.NumFrames
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetPixFmt() *string {
	return s.PixFmt
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetProfile() *string {
	return s.Profile
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetRotate() *string {
	return s.Rotate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetSar() *string {
	return s.Sar
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetStartTime() *string {
	return s.StartTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetTimebase() *string {
	return s.Timebase
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetWidth() *string {
	return s.Width
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetAvgFPS(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.AvgFPS = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetBitrate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Bitrate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecLongName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecLongName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecName(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecName = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecTag(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTag = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecTagString(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTagString = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetCodecTimeBase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.CodecTimeBase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetColorPrimaries(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.ColorPrimaries = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetColorRange(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.ColorRange = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetColorTransfer(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.ColorTransfer = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetDar(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Dar = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetDuration(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Duration = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetDurationInaccurate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.DurationInaccurate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetFps(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Fps = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetHasBFrames(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.HasBFrames = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetHeight(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Height = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetIndex(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Index = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetLang(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Lang = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetLevel(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Level = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetNetworkCost(v *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.NetworkCost = v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetNumFrames(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.NumFrames = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetPixFmt(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.PixFmt = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetProfile(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Profile = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetRotate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Rotate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetSar(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Sar = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetStartTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.StartTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetTimebase(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Timebase = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetWidth(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.Width = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) Validate() error {
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost struct {
	// The average bitrate. Unit: Kbit/s.
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

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetAvgBitrate() *string {
	return s.AvgBitrate
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetCostBandwidth() *string {
	return s.CostBandwidth
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) GetPreloadTime() *string {
	return s.PreloadTime
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetAvgBitrate(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.AvgBitrate = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetCostBandwidth(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.CostBandwidth = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) SetPreloadTime(v string) *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost {
	s.PreloadTime = &v
	return s
}

func (s *SubmitMediaInfoJobResponseBodyMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost) Validate() error {
	return dara.Validate(s)
}
