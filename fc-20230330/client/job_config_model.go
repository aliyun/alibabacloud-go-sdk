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
	// The maximum number of retries. Valid values: 0 to 100. This parameter specifies the maximum number of retries allowed when Simple Log Service triggers a function based on the trigger interval if an error occurs, such as insufficient permissions, network failures, and function execution exceptions. If the trigger fails after the maximum number of retries is reached, Simple Log Service triggers the function again when the next trigger interval arrives.
	//
	// example:
	//
	// 3
	MaxRetryTime *int32 `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	// The time interval at which Simple Log Service triggers function execution. For example, you can retrieve data from the Logstore within the last 120 seconds to Function Compute every 120 seconds to perform custom computing.
	//
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
