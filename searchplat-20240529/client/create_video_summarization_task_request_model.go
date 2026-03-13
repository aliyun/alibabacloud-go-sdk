// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSummarizationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *CreateVideoSummarizationTaskRequestInput) *CreateVideoSummarizationTaskRequest
	GetInput() *CreateVideoSummarizationTaskRequestInput
	SetOutput(v *CreateVideoSummarizationTaskRequestOutput) *CreateVideoSummarizationTaskRequest
	GetOutput() *CreateVideoSummarizationTaskRequestOutput
	SetParameters(v map[string]interface{}) *CreateVideoSummarizationTaskRequest
	GetParameters() map[string]interface{}
}

type CreateVideoSummarizationTaskRequest struct {
	Input      *CreateVideoSummarizationTaskRequestInput  `json:"input,omitempty" xml:"input,omitempty" type:"Struct"`
	Output     *CreateVideoSummarizationTaskRequestOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Parameters map[string]interface{}                     `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s CreateVideoSummarizationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskRequest) GetInput() *CreateVideoSummarizationTaskRequestInput {
	return s.Input
}

func (s *CreateVideoSummarizationTaskRequest) GetOutput() *CreateVideoSummarizationTaskRequestOutput {
	return s.Output
}

func (s *CreateVideoSummarizationTaskRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateVideoSummarizationTaskRequest) SetInput(v *CreateVideoSummarizationTaskRequestInput) *CreateVideoSummarizationTaskRequest {
	s.Input = v
	return s
}

func (s *CreateVideoSummarizationTaskRequest) SetOutput(v *CreateVideoSummarizationTaskRequestOutput) *CreateVideoSummarizationTaskRequest {
	s.Output = v
	return s
}

func (s *CreateVideoSummarizationTaskRequest) SetParameters(v map[string]interface{}) *CreateVideoSummarizationTaskRequest {
	s.Parameters = v
	return s
}

func (s *CreateVideoSummarizationTaskRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVideoSummarizationTaskRequestInput struct {
	Chunks   []*CreateVideoSummarizationTaskRequestInputChunks `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	FileName *string                                           `json:"file_name,omitempty" xml:"file_name,omitempty"`
	Oss      *string                                           `json:"oss,omitempty" xml:"oss,omitempty"`
	Url      *string                                           `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateVideoSummarizationTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskRequestInput) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskRequestInput) GetChunks() []*CreateVideoSummarizationTaskRequestInputChunks {
	return s.Chunks
}

func (s *CreateVideoSummarizationTaskRequestInput) GetFileName() *string {
	return s.FileName
}

func (s *CreateVideoSummarizationTaskRequestInput) GetOss() *string {
	return s.Oss
}

func (s *CreateVideoSummarizationTaskRequestInput) GetUrl() *string {
	return s.Url
}

func (s *CreateVideoSummarizationTaskRequestInput) SetChunks(v []*CreateVideoSummarizationTaskRequestInputChunks) *CreateVideoSummarizationTaskRequestInput {
	s.Chunks = v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInput) SetFileName(v string) *CreateVideoSummarizationTaskRequestInput {
	s.FileName = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInput) SetOss(v string) *CreateVideoSummarizationTaskRequestInput {
	s.Oss = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInput) SetUrl(v string) *CreateVideoSummarizationTaskRequestInput {
	s.Url = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInput) Validate() error {
	if s.Chunks != nil {
		for _, item := range s.Chunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVideoSummarizationTaskRequestInputChunks struct {
	EndTime    *float32                                                   `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Snapshots  []*CreateVideoSummarizationTaskRequestInputChunksSnapshots `json:"snapshots,omitempty" xml:"snapshots,omitempty" type:"Repeated"`
	StartTime  *float32                                                   `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Transcript *string                                                    `json:"transcript,omitempty" xml:"transcript,omitempty"`
}

func (s CreateVideoSummarizationTaskRequestInputChunks) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskRequestInputChunks) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) GetEndTime() *float32 {
	return s.EndTime
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) GetSnapshots() []*CreateVideoSummarizationTaskRequestInputChunksSnapshots {
	return s.Snapshots
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) GetStartTime() *float32 {
	return s.StartTime
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) GetTranscript() *string {
	return s.Transcript
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) SetEndTime(v float32) *CreateVideoSummarizationTaskRequestInputChunks {
	s.EndTime = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) SetSnapshots(v []*CreateVideoSummarizationTaskRequestInputChunksSnapshots) *CreateVideoSummarizationTaskRequestInputChunks {
	s.Snapshots = v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) SetStartTime(v float32) *CreateVideoSummarizationTaskRequestInputChunks {
	s.StartTime = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) SetTranscript(v string) *CreateVideoSummarizationTaskRequestInputChunks {
	s.Transcript = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInputChunks) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVideoSummarizationTaskRequestInputChunksSnapshots struct {
	FrameIndex *int32   `json:"frame_index,omitempty" xml:"frame_index,omitempty"`
	FrameTime  *float32 `json:"frame_time,omitempty" xml:"frame_time,omitempty"`
	Path       *string  `json:"path,omitempty" xml:"path,omitempty"`
}

func (s CreateVideoSummarizationTaskRequestInputChunksSnapshots) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskRequestInputChunksSnapshots) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskRequestInputChunksSnapshots) GetFrameIndex() *int32 {
	return s.FrameIndex
}

func (s *CreateVideoSummarizationTaskRequestInputChunksSnapshots) GetFrameTime() *float32 {
	return s.FrameTime
}

func (s *CreateVideoSummarizationTaskRequestInputChunksSnapshots) GetPath() *string {
	return s.Path
}

func (s *CreateVideoSummarizationTaskRequestInputChunksSnapshots) SetFrameIndex(v int32) *CreateVideoSummarizationTaskRequestInputChunksSnapshots {
	s.FrameIndex = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInputChunksSnapshots) SetFrameTime(v float32) *CreateVideoSummarizationTaskRequestInputChunksSnapshots {
	s.FrameTime = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInputChunksSnapshots) SetPath(v string) *CreateVideoSummarizationTaskRequestInputChunksSnapshots {
	s.Path = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestInputChunksSnapshots) Validate() error {
	return dara.Validate(s)
}

type CreateVideoSummarizationTaskRequestOutput struct {
	Oss  *string `json:"oss,omitempty" xml:"oss,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateVideoSummarizationTaskRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskRequestOutput) GetOss() *string {
	return s.Oss
}

func (s *CreateVideoSummarizationTaskRequestOutput) GetType() *string {
	return s.Type
}

func (s *CreateVideoSummarizationTaskRequestOutput) SetOss(v string) *CreateVideoSummarizationTaskRequestOutput {
	s.Oss = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestOutput) SetType(v string) *CreateVideoSummarizationTaskRequestOutput {
	s.Type = &v
	return s
}

func (s *CreateVideoSummarizationTaskRequestOutput) Validate() error {
	return dara.Validate(s)
}
