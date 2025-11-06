// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueRateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetQueueRateShrinkRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetQueueRateShrinkRequest
	GetInstanceId() *string
	SetQueueNamesShrink(v string) *GetQueueRateShrinkRequest
	GetQueueNamesShrink() *string
	SetVhostName(v string) *GetQueueRateShrinkRequest
	GetVhostName() *string
}

type GetQueueRateShrinkRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetQueueRateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueRateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetQueueRateShrinkRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetQueueRateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueueRateShrinkRequest) GetQueueNamesShrink() *string {
	return s.QueueNamesShrink
}

func (s *GetQueueRateShrinkRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetQueueRateShrinkRequest) SetConsoleSessionId(v string) *GetQueueRateShrinkRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetQueueRateShrinkRequest) SetInstanceId(v string) *GetQueueRateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueueRateShrinkRequest) SetQueueNamesShrink(v string) *GetQueueRateShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

func (s *GetQueueRateShrinkRequest) SetVhostName(v string) *GetQueueRateShrinkRequest {
	s.VhostName = &v
	return s
}

func (s *GetQueueRateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
