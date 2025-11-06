// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueRateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetQueueRateRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetQueueRateRequest
	GetInstanceId() *string
	SetQueueNames(v map[string]interface{}) *GetQueueRateRequest
	GetQueueNames() map[string]interface{}
	SetVhostName(v string) *GetQueueRateRequest
	GetVhostName() *string
}

type GetQueueRateRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	QueueNames map[string]interface{} `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s GetQueueRateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueueRateRequest) GoString() string {
	return s.String()
}

func (s *GetQueueRateRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetQueueRateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueueRateRequest) GetQueueNames() map[string]interface{} {
	return s.QueueNames
}

func (s *GetQueueRateRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *GetQueueRateRequest) SetConsoleSessionId(v string) *GetQueueRateRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetQueueRateRequest) SetInstanceId(v string) *GetQueueRateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQueueRateRequest) SetQueueNames(v map[string]interface{}) *GetQueueRateRequest {
	s.QueueNames = v
	return s
}

func (s *GetQueueRateRequest) SetVhostName(v string) *GetQueueRateRequest {
	s.VhostName = &v
	return s
}

func (s *GetQueueRateRequest) Validate() error {
	return dara.Validate(s)
}
