// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertOutputDetail interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MediaConvertOutputDetail
	GetCode() *string
	SetCreateTime(v string) *MediaConvertOutputDetail
	GetCreateTime() *string
	SetFinishTime(v string) *MediaConvertOutputDetail
	GetFinishTime() *string
	SetMessage(v string) *MediaConvertOutputDetail
	GetMessage() *string
	SetName(v string) *MediaConvertOutputDetail
	GetName() *string
	SetResult(v *MediaConvertOutputDetailResult) *MediaConvertOutputDetail
	GetResult() *MediaConvertOutputDetailResult
	SetStatus(v string) *MediaConvertOutputDetail
	GetStatus() *string
	SetTaskId(v string) *MediaConvertOutputDetail
	GetTaskId() *string
}

type MediaConvertOutputDetail struct {
	// The error code for a failed task.
	//
	// example:
	//
	// InvalidParameter.ResourceContentBad
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time the output task was created, in UTC format (*yyyy-MM-dd*T*HH:mm:ss*Z)
	//
	// example:
	//
	// 2024-12-07T06:06:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time the output task finished, in UTC format (*yyyy-MM-dd*T*HH:mm:ss*Z)
	//
	// example:
	//
	// 2024-12-07T13:01:07Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The reason for a task failure.
	//
	// example:
	//
	// The resource operated InputFile is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the output.
	//
	// example:
	//
	// 720P-mp4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detailed output results.
	Result *MediaConvertOutputDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// 	- Init: Initializing the task.
	//
	// 	- Scheduled: The task is scheduled for processing.
	//
	// 	- Success: The task is completed.
	//
	// 	- Failed: The task failed.
	//
	// 	- Skipped: The task was skipped.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ******4215e042b3966ca5441e******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s MediaConvertOutputDetail) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputDetail) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputDetail) GetCode() *string {
	return s.Code
}

func (s *MediaConvertOutputDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *MediaConvertOutputDetail) GetFinishTime() *string {
	return s.FinishTime
}

func (s *MediaConvertOutputDetail) GetMessage() *string {
	return s.Message
}

func (s *MediaConvertOutputDetail) GetName() *string {
	return s.Name
}

func (s *MediaConvertOutputDetail) GetResult() *MediaConvertOutputDetailResult {
	return s.Result
}

func (s *MediaConvertOutputDetail) GetStatus() *string {
	return s.Status
}

func (s *MediaConvertOutputDetail) GetTaskId() *string {
	return s.TaskId
}

func (s *MediaConvertOutputDetail) SetCode(v string) *MediaConvertOutputDetail {
	s.Code = &v
	return s
}

func (s *MediaConvertOutputDetail) SetCreateTime(v string) *MediaConvertOutputDetail {
	s.CreateTime = &v
	return s
}

func (s *MediaConvertOutputDetail) SetFinishTime(v string) *MediaConvertOutputDetail {
	s.FinishTime = &v
	return s
}

func (s *MediaConvertOutputDetail) SetMessage(v string) *MediaConvertOutputDetail {
	s.Message = &v
	return s
}

func (s *MediaConvertOutputDetail) SetName(v string) *MediaConvertOutputDetail {
	s.Name = &v
	return s
}

func (s *MediaConvertOutputDetail) SetResult(v *MediaConvertOutputDetailResult) *MediaConvertOutputDetail {
	s.Result = v
	return s
}

func (s *MediaConvertOutputDetail) SetStatus(v string) *MediaConvertOutputDetail {
	s.Status = &v
	return s
}

func (s *MediaConvertOutputDetail) SetTaskId(v string) *MediaConvertOutputDetail {
	s.TaskId = &v
	return s
}

func (s *MediaConvertOutputDetail) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaConvertOutputDetailResult struct {
	// The metadata of the audio and video streams.
	OutFileMeta *MediaConvertOutputDetailFileMeta `json:"OutFileMeta,omitempty" xml:"OutFileMeta,omitempty"`
	// Details about the generated output file.
	OutputFile *MediaConvertOutputDetailResultOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
}

func (s MediaConvertOutputDetailResult) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputDetailResult) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputDetailResult) GetOutFileMeta() *MediaConvertOutputDetailFileMeta {
	return s.OutFileMeta
}

func (s *MediaConvertOutputDetailResult) GetOutputFile() *MediaConvertOutputDetailResultOutputFile {
	return s.OutputFile
}

func (s *MediaConvertOutputDetailResult) SetOutFileMeta(v *MediaConvertOutputDetailFileMeta) *MediaConvertOutputDetailResult {
	s.OutFileMeta = v
	return s
}

func (s *MediaConvertOutputDetailResult) SetOutputFile(v *MediaConvertOutputDetailResultOutputFile) *MediaConvertOutputDetailResult {
	s.OutputFile = v
	return s
}

func (s *MediaConvertOutputDetailResult) Validate() error {
	if s.OutFileMeta != nil {
		if err := s.OutFileMeta.Validate(); err != nil {
			return err
		}
	}
	if s.OutputFile != nil {
		if err := s.OutputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MediaConvertOutputDetailResultOutputFile struct {
	// The value depends on the Type field:
	//
	// 	- If Type is set to OSS, the value is the URL of the output file. The following formats are supported: oss://... and https://...
	//
	// 	- If Type is set to Media, the value is the ID of the media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output file. Valid values:
	//
	// 	- OSS: an Object Storage Service (OSS) object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// If Type is set to Media, this field provides the actual storage URL of the media asset.
	//
	// example:
	//
	// oss://bucket/path/to/video.mp4
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MediaConvertOutputDetailResultOutputFile) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputDetailResultOutputFile) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputDetailResultOutputFile) GetMedia() *string {
	return s.Media
}

func (s *MediaConvertOutputDetailResultOutputFile) GetType() *string {
	return s.Type
}

func (s *MediaConvertOutputDetailResultOutputFile) GetUrl() *string {
	return s.Url
}

func (s *MediaConvertOutputDetailResultOutputFile) SetMedia(v string) *MediaConvertOutputDetailResultOutputFile {
	s.Media = &v
	return s
}

func (s *MediaConvertOutputDetailResultOutputFile) SetType(v string) *MediaConvertOutputDetailResultOutputFile {
	s.Type = &v
	return s
}

func (s *MediaConvertOutputDetailResultOutputFile) SetUrl(v string) *MediaConvertOutputDetailResultOutputFile {
	s.Url = &v
	return s
}

func (s *MediaConvertOutputDetailResultOutputFile) Validate() error {
	return dara.Validate(s)
}
