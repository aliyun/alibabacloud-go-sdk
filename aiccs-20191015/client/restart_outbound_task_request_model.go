// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RestartOutboundTaskRequest
	GetInstanceId() *string
	SetOutboundTaskId(v int64) *RestartOutboundTaskRequest
	GetOutboundTaskId() *int64
}

type RestartOutboundTaskRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	OutboundTaskId *int64 `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
}

func (s RestartOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *RestartOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartOutboundTaskRequest) GetOutboundTaskId() *int64 {
	return s.OutboundTaskId
}

func (s *RestartOutboundTaskRequest) SetInstanceId(v string) *RestartOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartOutboundTaskRequest) SetOutboundTaskId(v int64) *RestartOutboundTaskRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *RestartOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}
