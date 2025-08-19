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
	// example:
	//
	// serverless
	FilterTag *string `json:"filterTag,omitempty" xml:"filterTag,omitempty"`
	// example:
	//
	// JSON
	NotifyContentFormat *string `json:"notifyContentFormat,omitempty" xml:"notifyContentFormat,omitempty"`
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
