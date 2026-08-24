// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetDeadLockDetailRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetDeadLockDetailRequest
	GetNodeId() *string
	SetSource(v string) *GetDeadLockDetailRequest
	GetSource() *string
	SetTextId(v string) *GetDeadLockDetailRequest
	GetTextId() *string
}

type GetDeadLockDetailRequest struct {
	// The ID of the database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// > Required for PolarDB for MySQL cluster instances.
	//
	// example:
	//
	// pi-bp16v3824rt73****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The source of the analysis task:
	//
	// - **MANUAL*	- or **not specified**: queries the recent deadlock analysis task.
	//
	// - **AUTO**: queries the full deadlock analysis task.
	//
	// example:
	//
	// AUTO
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the deadlock text. This value is returned from the GetDeadLockHistory operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// a0e390cd5aca9ae964448f040153****
	TextId *string `json:"TextId,omitempty" xml:"TextId,omitempty"`
}

func (s GetDeadLockDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDeadLockDetailRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetDeadLockDetailRequest) GetSource() *string {
	return s.Source
}

func (s *GetDeadLockDetailRequest) GetTextId() *string {
	return s.TextId
}

func (s *GetDeadLockDetailRequest) SetInstanceId(v string) *GetDeadLockDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDeadLockDetailRequest) SetNodeId(v string) *GetDeadLockDetailRequest {
	s.NodeId = &v
	return s
}

func (s *GetDeadLockDetailRequest) SetSource(v string) *GetDeadLockDetailRequest {
	s.Source = &v
	return s
}

func (s *GetDeadLockDetailRequest) SetTextId(v string) *GetDeadLockDetailRequest {
	s.TextId = &v
	return s
}

func (s *GetDeadLockDetailRequest) Validate() error {
	return dara.Validate(s)
}
