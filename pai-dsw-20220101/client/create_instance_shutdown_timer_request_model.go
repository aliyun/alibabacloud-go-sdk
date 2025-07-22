// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceShutdownTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDueTime(v string) *CreateInstanceShutdownTimerRequest
	GetDueTime() *string
	SetRemainingTimeInMs(v int64) *CreateInstanceShutdownTimerRequest
	GetRemainingTimeInMs() *int64
}

type CreateInstanceShutdownTimerRequest struct {
	// The scheduled stop time.
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// The time duration before the instance is stopped. Unit: milliseconds.
	//
	// example:
	//
	// 3600000
	RemainingTimeInMs *int64 `json:"RemainingTimeInMs,omitempty" xml:"RemainingTimeInMs,omitempty"`
}

func (s CreateInstanceShutdownTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceShutdownTimerRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerRequest) GetDueTime() *string {
	return s.DueTime
}

func (s *CreateInstanceShutdownTimerRequest) GetRemainingTimeInMs() *int64 {
	return s.RemainingTimeInMs
}

func (s *CreateInstanceShutdownTimerRequest) SetDueTime(v string) *CreateInstanceShutdownTimerRequest {
	s.DueTime = &v
	return s
}

func (s *CreateInstanceShutdownTimerRequest) SetRemainingTimeInMs(v int64) *CreateInstanceShutdownTimerRequest {
	s.RemainingTimeInMs = &v
	return s
}

func (s *CreateInstanceShutdownTimerRequest) Validate() error {
	return dara.Validate(s)
}
