// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstanceInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRunId(v string) *RunInstanceInspectionResponseBody
	GetRunId() *string
}

type RunInstanceInspectionResponseBody struct {
	// example:
	//
	// inspection-run-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
}

func (s RunInstanceInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunInstanceInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *RunInstanceInspectionResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *RunInstanceInspectionResponseBody) SetRunId(v string) *RunInstanceInspectionResponseBody {
	s.RunId = &v
	return s
}

func (s *RunInstanceInspectionResponseBody) Validate() error {
	return dara.Validate(s)
}
