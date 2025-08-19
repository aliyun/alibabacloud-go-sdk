// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMaxRetryTime(v int32) *JobConfig
	GetMaxRetryTime() *int32
	SetTriggerInterval(v int32) *JobConfig
	GetTriggerInterval() *int32
}

type JobConfig struct {
	// example:
	//
	// 3
	MaxRetryTime *int32 `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	// example:
	//
	// 60
	TriggerInterval *int32 `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
}

func (s JobConfig) String() string {
	return dara.Prettify(s)
}

func (s JobConfig) GoString() string {
	return s.String()
}

func (s *JobConfig) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *JobConfig) GetTriggerInterval() *int32 {
	return s.TriggerInterval
}

func (s *JobConfig) SetMaxRetryTime(v int32) *JobConfig {
	s.MaxRetryTime = &v
	return s
}

func (s *JobConfig) SetTriggerInterval(v int32) *JobConfig {
	s.TriggerInterval = &v
	return s
}

func (s *JobConfig) Validate() error {
	return dara.Validate(s)
}
