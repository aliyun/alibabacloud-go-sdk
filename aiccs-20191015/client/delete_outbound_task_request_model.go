// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOutboundTaskRequest
	GetInstanceId() *string
	SetOutboundTaskId(v int64) *DeleteOutboundTaskRequest
	GetOutboundTaskId() *int64
}

type DeleteOutboundTaskRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	OutboundTaskId *int64 `json:"OutboundTaskId,omitempty" xml:"OutboundTaskId,omitempty"`
}

func (s DeleteOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOutboundTaskRequest) GetOutboundTaskId() *int64 {
	return s.OutboundTaskId
}

func (s *DeleteOutboundTaskRequest) SetInstanceId(v string) *DeleteOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOutboundTaskRequest) SetOutboundTaskId(v int64) *DeleteOutboundTaskRequest {
	s.OutboundTaskId = &v
	return s
}

func (s *DeleteOutboundTaskRequest) Validate() error {
	return dara.Validate(s)
}
