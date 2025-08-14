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
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime *string                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime *string                         `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Name       *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Result     *MediaConvertOutputDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Status     *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string                         `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	return dara.Validate(s)
}

type MediaConvertOutputDetailResult struct {
	OutFileMeta *MediaConvertOutputDetailFileMeta         `json:"OutFileMeta,omitempty" xml:"OutFileMeta,omitempty"`
	OutputFile  *MediaConvertOutputDetailResultOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type MediaConvertOutputDetailResultOutputFile struct {
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
