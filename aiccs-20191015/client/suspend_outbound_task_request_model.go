// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SuspendOutboundTaskRequest
	GetInstanceId() *string
	SetOutboundTaskId(v int64) *SuspendOutboundTaskRequest
	GetOutboundTaskId() *int64
}

type SuspendOutboundTaskRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	OutboundTaskId *int64 `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
}

func (s SuspendOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *SuspendOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SuspendOutboundTaskRequest) GetOutboundTaskId() *int64 {
	return s.OutboundTaskId
}

func (s *SuspendOutboundTaskRequest) SetInstanceId(v string) *SuspendOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendOutboundTaskRequest) SetOutboundTaskId(v int64) *SuspendOutboundTaskRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *SuspendOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}
