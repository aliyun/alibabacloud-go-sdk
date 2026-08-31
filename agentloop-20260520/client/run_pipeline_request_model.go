// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromTime(v int64) *RunPipelineRequest
	GetFromTime() *int64
	SetOutput(v *RunPipelineRequestOutput) *RunPipelineRequest
	GetOutput() *RunPipelineRequestOutput
	SetToTime(v int64) *RunPipelineRequest
	GetToTime() *int64
}

type RunPipelineRequest struct {
	// The start of the data processing window, in UNIX seconds. This parameter is required for SQL Pipeline and time window-based WorkItem Source. This parameter is optional for checkpoint-based WorkItem Source types such as Dataset and Logstore Pull.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The output control settings. SQL Pipeline supports overriding. WorkItem Pipeline uses the saved Sink configuration and must keep the default value.
	Output *RunPipelineRequestOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// The end of the data processing window, in UNIX seconds. This parameter must be provided together with fromTime and must be greater than fromTime. This parameter is optional for checkpoint-based WorkItem Source types.
	//
	// example:
	//
	// 1735747200
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
}

func (s RunPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s RunPipelineRequest) GoString() string {
	return s.String()
}

func (s *RunPipelineRequest) GetFromTime() *int64 {
	return s.FromTime
}

func (s *RunPipelineRequest) GetOutput() *RunPipelineRequestOutput {
	return s.Output
}

func (s *RunPipelineRequest) GetToTime() *int64 {
	return s.ToTime
}

func (s *RunPipelineRequest) SetFromTime(v int64) *RunPipelineRequest {
	s.FromTime = &v
	return s
}

func (s *RunPipelineRequest) SetOutput(v *RunPipelineRequestOutput) *RunPipelineRequest {
	s.Output = v
	return s
}

func (s *RunPipelineRequest) SetToTime(v int64) *RunPipelineRequest {
	s.ToTime = &v
	return s
}

func (s *RunPipelineRequest) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunPipelineRequestOutput struct {
	// Specifies whether to write to the target Dataset. Default value: true.
	//
	// example:
	//
	// true
	Dataset *bool `json:"dataset,omitempty" xml:"dataset,omitempty"`
	// Specifies whether to return inline results in the response. Default value: false.
	//
	// example:
	//
	// false
	Inline *bool `json:"inline,omitempty" xml:"inline,omitempty"`
}

func (s RunPipelineRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s RunPipelineRequestOutput) GoString() string {
	return s.String()
}

func (s *RunPipelineRequestOutput) GetDataset() *bool {
	return s.Dataset
}

func (s *RunPipelineRequestOutput) GetInline() *bool {
	return s.Inline
}

func (s *RunPipelineRequestOutput) SetDataset(v bool) *RunPipelineRequestOutput {
	s.Dataset = &v
	return s
}

func (s *RunPipelineRequestOutput) SetInline(v bool) *RunPipelineRequestOutput {
	s.Inline = &v
	return s
}

func (s *RunPipelineRequestOutput) Validate() error {
	return dara.Validate(s)
}
