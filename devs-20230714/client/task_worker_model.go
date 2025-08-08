// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskWorker interface {
	dara.Model
	String() string
	GoString() string
	SetPresetWorker(v string) *TaskWorker
	GetPresetWorker() *string
}

type TaskWorker struct {
	// example:
	//
	// serverless-runner
	PresetWorker *string `json:"presetWorker,omitempty" xml:"presetWorker,omitempty"`
}

func (s TaskWorker) String() string {
	return dara.Prettify(s)
}

func (s TaskWorker) GoString() string {
	return s.String()
}

func (s *TaskWorker) GetPresetWorker() *string {
	return s.PresetWorker
}

func (s *TaskWorker) SetPresetWorker(v string) *TaskWorker {
	s.PresetWorker = &v
	return s
}

func (s *TaskWorker) Validate() error {
	return dara.Validate(s)
}
