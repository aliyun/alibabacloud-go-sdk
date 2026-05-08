// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTaskResult interface {
	dara.Model
	String() string
	GoString() string
	SetTextTask(v *TextTask) *TextTaskResult
	GetTextTask() *TextTask
}

type TextTaskResult struct {
	TextTask *TextTask `json:"textTask,omitempty" xml:"textTask,omitempty"`
}

func (s TextTaskResult) String() string {
	return dara.Prettify(s)
}

func (s TextTaskResult) GoString() string {
	return s.String()
}

func (s *TextTaskResult) GetTextTask() *TextTask {
	return s.TextTask
}

func (s *TextTaskResult) SetTextTask(v *TextTask) *TextTaskResult {
	s.TextTask = v
	return s
}

func (s *TextTaskResult) Validate() error {
	if s.TextTask != nil {
		if err := s.TextTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}
