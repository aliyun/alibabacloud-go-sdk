// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildPipelineRunLogLine interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *BuildPipelineRunLogLine
	GetContent() *string
	SetOffset(v int64) *BuildPipelineRunLogLine
	GetOffset() *int64
}

type BuildPipelineRunLogLine struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Offset  *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
}

func (s BuildPipelineRunLogLine) String() string {
	return dara.Prettify(s)
}

func (s BuildPipelineRunLogLine) GoString() string {
	return s.String()
}

func (s *BuildPipelineRunLogLine) GetContent() *string {
	return s.Content
}

func (s *BuildPipelineRunLogLine) GetOffset() *int64 {
	return s.Offset
}

func (s *BuildPipelineRunLogLine) SetContent(v string) *BuildPipelineRunLogLine {
	s.Content = &v
	return s
}

func (s *BuildPipelineRunLogLine) SetOffset(v int64) *BuildPipelineRunLogLine {
	s.Offset = &v
	return s
}

func (s *BuildPipelineRunLogLine) Validate() error {
	return dara.Validate(s)
}
