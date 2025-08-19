// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetPushRetryStrategy(v string) *RetryStrategy
	GetPushRetryStrategy() *string
}

type RetryStrategy struct {
	// example:
	//
	// BACKOFF_RETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s RetryStrategy) String() string {
	return dara.Prettify(s)
}

func (s RetryStrategy) GoString() string {
	return s.String()
}

func (s *RetryStrategy) GetPushRetryStrategy() *string {
	return s.PushRetryStrategy
}

func (s *RetryStrategy) SetPushRetryStrategy(v string) *RetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

func (s *RetryStrategy) Validate() error {
	return dara.Validate(s)
}
