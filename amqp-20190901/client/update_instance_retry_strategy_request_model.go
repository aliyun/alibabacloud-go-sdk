// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRetryStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *UpdateInstanceRetryStrategyRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *UpdateInstanceRetryStrategyRequest
	GetInstanceId() *string
	SetRetryInterval(v int32) *UpdateInstanceRetryStrategyRequest
	GetRetryInterval() *int32
	SetRetryTimes(v int32) *UpdateInstanceRetryStrategyRequest
	GetRetryTimes() *int32
}

type UpdateInstanceRetryStrategyRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RetryInterval    *int32  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	RetryTimes       *int32  `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
}

func (s UpdateInstanceRetryStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRetryStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRetryStrategyRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *UpdateInstanceRetryStrategyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRetryStrategyRequest) GetRetryInterval() *int32 {
	return s.RetryInterval
}

func (s *UpdateInstanceRetryStrategyRequest) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *UpdateInstanceRetryStrategyRequest) SetConsoleSessionId(v string) *UpdateInstanceRetryStrategyRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *UpdateInstanceRetryStrategyRequest) SetInstanceId(v string) *UpdateInstanceRetryStrategyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRetryStrategyRequest) SetRetryInterval(v int32) *UpdateInstanceRetryStrategyRequest {
	s.RetryInterval = &v
	return s
}

func (s *UpdateInstanceRetryStrategyRequest) SetRetryTimes(v int32) *UpdateInstanceRetryStrategyRequest {
	s.RetryTimes = &v
	return s
}

func (s *UpdateInstanceRetryStrategyRequest) Validate() error {
	return dara.Validate(s)
}
