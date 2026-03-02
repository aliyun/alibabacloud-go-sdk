// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobRequestBody interface {
	dara.Model
	String() string
	GoString() string
	SetStopStrategy(v string) *StopJobRequestBody
	GetStopStrategy() *string
}

type StopJobRequestBody struct {
	// The strategy that is used to stop a job. Valid values:
	//
	// 	- `NONE`: Directly stop the job.
	//
	// 	- `STOP_WITH_SAVEPOINT`: Stop the job after a savepoint is generated.
	//
	// 	- `STOP_WITH_DRAIN`: Stop the job and drain the pipeline. Use this strategy only if you want to permanently terminate the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// NONE
	StopStrategy *string `json:"stopStrategy,omitempty" xml:"stopStrategy,omitempty"`
}

func (s StopJobRequestBody) String() string {
	return dara.Prettify(s)
}

func (s StopJobRequestBody) GoString() string {
	return s.String()
}

func (s *StopJobRequestBody) GetStopStrategy() *string {
	return s.StopStrategy
}

func (s *StopJobRequestBody) SetStopStrategy(v string) *StopJobRequestBody {
	s.StopStrategy = &v
	return s
}

func (s *StopJobRequestBody) Validate() error {
	return dara.Validate(s)
}
