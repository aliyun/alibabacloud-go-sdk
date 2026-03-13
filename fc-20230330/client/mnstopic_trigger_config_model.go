// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMNSTopicTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFilterTag(v string) *MNSTopicTriggerConfig
	GetFilterTag() *string
	SetNotifyContentFormat(v string) *MNSTopicTriggerConfig
	GetNotifyContentFormat() *string
	SetNotifyStrategy(v string) *MNSTopicTriggerConfig
	GetNotifyStrategy() *string
}

type MNSTopicTriggerConfig struct {
	// The filtering tag. Function execution is triggered only when a message that contains the specified filter tag is received.
	//
	// example:
	//
	// serverless
	FilterTag *string `json:"filterTag,omitempty" xml:"filterTag,omitempty"`
	// The format of the event content. The following two formats are supported:
	//
	// 	- **JSON**
	//
	// 	- **STREAM**
	//
	// >  The default format is STREAM.
	//
	// example:
	//
	// JSON
	NotifyContentFormat *string `json:"notifyContentFormat,omitempty" xml:"notifyContentFormat,omitempty"`
	// The retry policy.
	//
	// 	- **BACKOFF_RETRY**: a backoff retry policy. A total of 3 retries are made. The interval between 2 retries is a random value between 10 and 20 seconds. This is the default value.
	//
	// 	- **EXPONENTIAL_DECAY_RETRY**: an exponential decay retry policy. A total of 176 retries are made, with the interval of each retry increases exponentially to 512 seconds, and the total retry period is 1 day. The interval between two consecutive retries exponentially increases to a maximum of 512 seconds. For example, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 512... 512. The number of 512s is 167.
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
}

func (s MNSTopicTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s MNSTopicTriggerConfig) GoString() string {
	return s.String()
}

func (s *MNSTopicTriggerConfig) GetFilterTag() *string {
	return s.FilterTag
}

func (s *MNSTopicTriggerConfig) GetNotifyContentFormat() *string {
	return s.NotifyContentFormat
}

func (s *MNSTopicTriggerConfig) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *MNSTopicTriggerConfig) SetFilterTag(v string) *MNSTopicTriggerConfig {
	s.FilterTag = &v
	return s
}

func (s *MNSTopicTriggerConfig) SetNotifyContentFormat(v string) *MNSTopicTriggerConfig {
	s.NotifyContentFormat = &v
	return s
}

func (s *MNSTopicTriggerConfig) SetNotifyStrategy(v string) *MNSTopicTriggerConfig {
	s.NotifyStrategy = &v
	return s
}

func (s *MNSTopicTriggerConfig) Validate() error {
	return dara.Validate(s)
}
