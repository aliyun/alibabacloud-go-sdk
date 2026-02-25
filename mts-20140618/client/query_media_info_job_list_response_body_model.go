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
	MediaInfoJobList        *QueryMediaInfoJobListResponseBodyMediaInfoJobList        `json:"MediaInfoJobList,omitempty" xml:"MediaInfoJobList,omitempty" type:"Struct"`
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
	Async            *bool                                                                          `json:"Async,omitempty" xml:"Async,omitempty"`
	Code             *string                                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime     *string                                                                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Input            *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobInput            `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	JobId            *string                                                                        `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MNSMessageResult *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	Message          *string                                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PipelineId       *string                                                                        `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Properties       *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobProperties       `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	State            *string                                                                        `json:"State,omitempty" xml:"State,omitempty"`
	UserData         *string                                                                        `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MessageId    *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
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
	Bitrate    *string                                                                         `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration   *string                                                                         `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileFormat *string                                                                         `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	FileMd5    *string                                                                         `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	FileSize   *string                                                                         `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	Format     *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat  `json:"Format,omitempty" xml:"Format,omitempty" type:"Struct"`
	Fps        *string                                                                         `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height     *string                                                                         `json:"Height,omitempty" xml:"Height,omitempty"`
	Streams    *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams `json:"Streams,omitempty" xml:"Streams,omitempty" type:"Struct"`
	Width      *string                                                                         `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Bitrate        *string                `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration       *string                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FormatLongName *string                `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName     *string                `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	NumPrograms    *string                `json:"NumPrograms,omitempty" xml:"NumPrograms,omitempty"`
	NumStreams     *string                `json:"NumStreams,omitempty" xml:"NumStreams,omitempty"`
	Size           *string                `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime      *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tags           map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) GetTags() map[string]interface{} {
	return s.Tags
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

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) SetTags(v map[string]interface{}) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat {
	s.Tags = v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesFormat) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreams struct {
	AudioStreamList    *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsAudioStreamList    `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Struct"`
	SubtitleStreamList *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsSubtitleStreamList `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty" type:"Struct"`
	VideoStreamList    *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamList    `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Struct"`
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
	Bitrate            *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	ChannelLayout      *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	Channels           *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	CodecLongName      *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName          *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag           *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString     *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase      *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Duration           *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationInaccurate *string `json:"DurationInaccurate,omitempty" xml:"DurationInaccurate,omitempty"`
	Index              *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NumFrames          *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	SampleFmt          *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	Samplerate         *string `json:"Samplerate,omitempty" xml:"Samplerate,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase           *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
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
	CodecLongName  *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Index          *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase       *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
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
	AvgFPS             *string                                                                                                              `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	Bitrate            *string                                                                                                              `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName      *string                                                                                                              `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName          *string                                                                                                              `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag           *string                                                                                                              `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString     *string                                                                                                              `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase      *string                                                                                                              `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	ColorPrimaries     *string                                                                                                              `json:"ColorPrimaries,omitempty" xml:"ColorPrimaries,omitempty"`
	ColorRange         *string                                                                                                              `json:"ColorRange,omitempty" xml:"ColorRange,omitempty"`
	ColorTransfer      *string                                                                                                              `json:"ColorTransfer,omitempty" xml:"ColorTransfer,omitempty"`
	Dar                *string                                                                                                              `json:"Dar,omitempty" xml:"Dar,omitempty"`
	DolbyVision        *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision `json:"DolbyVision,omitempty" xml:"DolbyVision,omitempty" type:"Struct"`
	Duration           *string                                                                                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DurationInaccurate *string                                                                                                              `json:"DurationInaccurate,omitempty" xml:"DurationInaccurate,omitempty"`
	Fps                *string                                                                                                              `json:"Fps,omitempty" xml:"Fps,omitempty"`
	HasBFrames         *string                                                                                                              `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height             *string                                                                                                              `json:"Height,omitempty" xml:"Height,omitempty"`
	Index              *string                                                                                                              `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang               *string                                                                                                              `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Level              *string                                                                                                              `json:"Level,omitempty" xml:"Level,omitempty"`
	NetworkCost        *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost `json:"NetworkCost,omitempty" xml:"NetworkCost,omitempty" type:"Struct"`
	NumFrames          *string                                                                                                              `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	PixFmt             *string                                                                                                              `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Profile            *string                                                                                                              `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Rotate             *string                                                                                                              `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	Sar                *string                                                                                                              `json:"Sar,omitempty" xml:"Sar,omitempty"`
	StartTime          *string                                                                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase           *string                                                                                                              `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	Width              *string                                                                                                              `json:"Width,omitempty" xml:"Width,omitempty"`
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

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) GetDolbyVision() *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision {
	return s.DolbyVision
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

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream) SetDolbyVision(v *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStream {
	s.DolbyVision = v
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
	if s.DolbyVision != nil {
		if err := s.DolbyVision.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkCost != nil {
		if err := s.NetworkCost.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision struct {
	Level   *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) GoString() string {
	return s.String()
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) GetLevel() *string {
	return s.Level
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) GetProfile() *string {
	return s.Profile
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) SetLevel(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision {
	s.Level = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) SetProfile(v string) *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision {
	s.Profile = &v
	return s
}

func (s *QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamDolbyVision) Validate() error {
	return dara.Validate(s)
}

type QueryMediaInfoJobListResponseBodyMediaInfoJobListMediaInfoJobPropertiesStreamsVideoStreamListVideoStreamNetworkCost struct {
	AvgBitrate    *string `json:"AvgBitrate,omitempty" xml:"AvgBitrate,omitempty"`
	CostBandwidth *string `json:"CostBandwidth,omitempty" xml:"CostBandwidth,omitempty"`
	PreloadTime   *string `json:"PreloadTime,omitempty" xml:"PreloadTime,omitempty"`
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
