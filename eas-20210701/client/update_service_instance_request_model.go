// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsReplica(v bool) *UpdateServiceInstanceRequest
	GetIsReplica() *bool
	SetDetach(v bool) *UpdateServiceInstanceRequest
	GetDetach() *bool
	SetHibernate(v bool) *UpdateServiceInstanceRequest
	GetHibernate() *bool
	SetIsolate(v bool) *UpdateServiceInstanceRequest
	GetIsolate() *bool
}

type UpdateServiceInstanceRequest struct {
	// Specifies whether the current instance is an instance unit.
	//
	// example:
	//
	// false
	IsReplica *bool `json:"IsReplica,omitempty" xml:"IsReplica,omitempty"`
	// Specifies whether to isolate the service instance. After an instance is isolated, it is no longer managed by the controller, and a new instance is created. The isolated instance is used to preserve the state or for debugging. Note: An isolated instance cannot be unisolated. Valid values:
	//
	// - true: The instance is isolated.
	//
	// example:
	//
	// true
	Detach *bool `json:"Detach,omitempty" xml:"Detach,omitempty"`
	// >This parameter is in invitational preview and is not supported.
	//
	// example:
	//
	// 0
	Hibernate *bool `json:"Hibernate,omitempty" xml:"Hibernate,omitempty"`
	// Specifies whether the instance receives traffic. Valid values:
	//
	// - true: The instance receives traffic.
	//
	// - false: The instance does not receive traffic.
	//
	// example:
	//
	// true
	Isolate *bool `json:"Isolate,omitempty" xml:"Isolate,omitempty"`
}

func (s UpdateServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceRequest) GetIsReplica() *bool {
	return s.IsReplica
}

func (s *UpdateServiceInstanceRequest) GetDetach() *bool {
	return s.Detach
}

func (s *UpdateServiceInstanceRequest) GetHibernate() *bool {
	return s.Hibernate
}

func (s *UpdateServiceInstanceRequest) GetIsolate() *bool {
	return s.Isolate
}

func (s *UpdateServiceInstanceRequest) SetIsReplica(v bool) *UpdateServiceInstanceRequest {
	s.IsReplica = &v
	return s
}

func (s *UpdateServiceInstanceRequest) SetDetach(v bool) *UpdateServiceInstanceRequest {
	s.Detach = &v
	return s
}

func (s *UpdateServiceInstanceRequest) SetHibernate(v bool) *UpdateServiceInstanceRequest {
	s.Hibernate = &v
	return s
}

func (s *UpdateServiceInstanceRequest) SetIsolate(v bool) *UpdateServiceInstanceRequest {
	s.Isolate = &v
	return s
}

func (s *UpdateServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
