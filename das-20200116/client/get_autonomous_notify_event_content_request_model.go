// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutonomousNotifyEventContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAutonomousNotifyEventContentRequest
	GetInstanceId() *string
	SetSpanId(v string) *GetAutonomousNotifyEventContentRequest
	GetSpanId() *string
	SetContext(v string) *GetAutonomousNotifyEventContentRequest
	GetContext() *string
}

type GetAutonomousNotifyEventContentRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The unique identifier of the event. You can call the [GetAutonomousNotifyEventsInRange](https://help.aliyun.com/document_detail/288371.html) operation to query the unique identifier returned by the SpanId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7e7b2774-95b8-4fa3-bd9c-0ab47cb7****
	SpanId *string `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Context *string `json:"__context,omitempty" xml:"__context,omitempty"`
}

func (s GetAutonomousNotifyEventContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventContentRequest) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventContentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutonomousNotifyEventContentRequest) GetSpanId() *string {
	return s.SpanId
}

func (s *GetAutonomousNotifyEventContentRequest) GetContext() *string {
	return s.Context
}

func (s *GetAutonomousNotifyEventContentRequest) SetInstanceId(v string) *GetAutonomousNotifyEventContentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutonomousNotifyEventContentRequest) SetSpanId(v string) *GetAutonomousNotifyEventContentRequest {
	s.SpanId = &v
	return s
}

func (s *GetAutonomousNotifyEventContentRequest) SetContext(v string) *GetAutonomousNotifyEventContentRequest {
	s.Context = &v
	return s
}

func (s *GetAutonomousNotifyEventContentRequest) Validate() error {
	return dara.Validate(s)
}
