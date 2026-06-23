// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopScheduledPreloadExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopScheduledPreloadExecutionRequest
	GetId() *string
}

type StopScheduledPreloadExecutionRequest struct {
	// The prefetch plan ID. You can obtain this value from the SuccessExecutions[].Id field returned by CreateScheduledPreloadExecutions. Before calling this operation, create a prefetch job by calling CreateScheduledPreloadJob, and then create an execution plan by calling CreateScheduledPreloadExecutions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 66599bd7397885b43804901c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopScheduledPreloadExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StopScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *StopScheduledPreloadExecutionRequest) GetId() *string {
	return s.Id
}

func (s *StopScheduledPreloadExecutionRequest) SetId(v string) *StopScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

func (s *StopScheduledPreloadExecutionRequest) Validate() error {
	return dara.Validate(s)
}
